package edifact

import "regexp"

// segmentTagPattern matches a well-formed UN/EDIFACT segment tag: exactly
// three uppercase letters (e.g. "UNB", "BGM", "DTM").
var segmentTagPattern = regexp.MustCompile(`^[A-Z]{3}$`)

// Parse lexes and parses a whole EDIFACT interchange into an Interchange
// AST. It never panics on malformed input: syntax problems are recorded as
// positioned errors in the returned ErrorList, and the parser recovers by
// resyncing at the next segment terminator so it can keep parsing whatever
// follows.
func Parse(src string) (*Interchange, ErrorList) {
	var errs ErrorList

	delims, consumed, present := detectUNA(src)
	lx := newLexer(src, delims, &errs)

	var una *UNAAdvice
	if present {
		start := lx.position()
		lx.skip(consumed)
		una = &UNAAdvice{Pos: start, Raw: src[:consumed]}
	}

	var segments []Segment
	for {
		seg, cont := parseSegment(lx)
		if seg != nil {
			segments = append(segments, *seg)
		}
		if !cont {
			break
		}
	}

	return &Interchange{UNA: una, Segments: segments, Delimiters: delims}, errs
}

// parseSegment parses one segment starting at the lexer's current position.
// It returns the segment (nil if nothing was there to parse, e.g. a stray
// terminator) and whether the caller should keep parsing further segments.
func parseSegment(lx *lexer) (*Segment, bool) {
	lx.skipInterSegmentWhitespace()
	if lx.eof() {
		return nil, false
	}

	startPos := lx.position()
	tag := lx.next()

	if tag.Kind == tokenSegmentTerminator {
		// A stray empty segment (e.g. two terminators in a row); skip it.
		return nil, !lx.eof()
	}
	if tag.Kind != tokenData {
		lx.errs.Add(tag.Pos, SeverityError, "expected a segment tag, found unexpected separator")
		return nil, !lx.eof()
	}

	seg := &Segment{Tag: tag.Raw, TagPos: tag.Pos, Pos: startPos}

	if !segmentTagPattern.MatchString(tag.Raw) {
		lx.errs.Add(tag.Pos, SeverityError, "invalid segment tag %q: expected exactly three uppercase letters", tag.Raw)
		endTok := recoverToNextTerminator(lx)
		seg.EndPos = endTok.Pos
		if endTok.Kind == tokenSegmentTerminator {
			seg.EndPos = lx.position()
		}
		return seg, !lx.eof()
	}

	consumeTagControlNumbers(lx, seg)

	endTok := parseSegmentBody(lx, seg)
	if endTok.Kind == tokenEOF {
		lx.errs.Add(endTok.Pos, SeverityError, "unexpected end of input: segment %q is missing its terminator", seg.Tag)
		seg.EndPos = endTok.Pos
		return seg, false
	}
	seg.EndPos = lx.position()
	return seg, !lx.eof()
}

// consumeTagControlNumbers consumes the segment tag's optional
// component-separator-delimited control-number components (explicit
// representation of repeating segments, e.g. the "1" and "2" in
// "GDS:1:2+..."), appending each to seg.TagControlNumbers. It stops at (and
// rewinds past) the first token that doesn't fit that pattern -- a
// tokenElementSep, tokenSegmentTerminator/EOF, or a componentSep not
// followed by data -- leaving the lexer positioned for parseSegmentBody to
// take over from there, so an ordinary segment with no control numbers
// (the common case) is completely unaffected.
func consumeTagControlNumbers(lx *lexer, seg *Segment) {
	for {
		pos, line, col, errLen := lx.pos, lx.line, lx.col, len(*lx.errs)

		sepTok := lx.next()
		if sepTok.Kind != tokenComponentSep {
			lx.pos, lx.line, lx.col = pos, line, col
			*lx.errs = (*lx.errs)[:errLen]
			return
		}

		numTok := lx.next()
		if numTok.Kind != tokenData {
			lx.pos, lx.line, lx.col = pos, line, col
			*lx.errs = (*lx.errs)[:errLen]
			return
		}

		seg.TagControlNumbers = append(seg.TagControlNumbers, numTok.Raw)
	}
}

// parseSegmentBody consumes elements/components until a segment terminator
// or EOF, populating seg.Elements. It returns the terminating token.
//
// Grammar reminder: a segment is `tag (SEP element)*`, i.e. every element is
// *introduced* by a separator rather than elements being interspersed
// between separators. So the very first separator right after the tag does
// not close a preceding element (there isn't one) — it only starts
// collecting the first one. `first` tracks that so it isn't mistaken for an
// empty leading element.
func parseSegmentBody(lx *lexer, seg *Segment) token {
	var elements []Element
	var curComponents []Component
	elementStart := lx.position()

	pendingData := ""
	pendingPos := lx.position()
	haveData := false
	first := true

	flushComponent := func(pos Position) {
		if haveData {
			curComponents = append(curComponents, Component{Raw: pendingData, Pos: pendingPos})
		} else {
			curComponents = append(curComponents, Component{Raw: "", Pos: pos})
		}
		haveData = false
		pendingData = ""
	}
	flushElement := func() {
		elements = append(elements, Element{Components: curComponents, Pos: elementStart})
		curComponents = nil
	}

	for {
		tok := lx.next()
		switch tok.Kind {
		case tokenData:
			pendingData = tok.Raw
			pendingPos = tok.Pos
			haveData = true
			first = false

		case tokenComponentSep:
			if !first {
				flushComponent(tok.Pos)
			}
			first = false

		case tokenElementSep:
			if !first {
				flushComponent(tok.Pos)
				flushElement()
			}
			elementStart = lx.position()
			first = false

		case tokenSegmentTerminator, tokenEOF:
			if !first {
				flushComponent(tok.Pos)
				flushElement()
			}
			seg.Elements = elements
			return tok
		}
	}
}

// recoverToNextTerminator discards tokens until (and including) the next
// segment terminator, or EOF, so parsing can continue after a malformed
// segment tag. It returns the token it stopped on.
func recoverToNextTerminator(lx *lexer) token {
	for {
		tok := lx.next()
		if tok.Kind == tokenSegmentTerminator || tok.Kind == tokenEOF {
			return tok
		}
	}
}
