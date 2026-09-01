package edifact

// tokenKind classifies a single lexer token.
type tokenKind int

const (
	tokenEOF tokenKind = iota
	tokenData
	tokenComponentSep
	tokenElementSep
	tokenSegmentTerminator
)

// token is one lexical unit: either a run of data or a single structural
// separator/terminator character.
type token struct {
	Kind tokenKind
	Raw  string // exact source text; only meaningful for tokenData
	Pos  Position
}

// lexer scans EDIFACT source byte-by-byte (the UNOA/UNOB/UNOC character sets
// EDIFACT is defined over are single-byte), tracking line/column as it goes
// so callers get positions relative to the original interchange text
// regardless of any UNA prefix already skipped via skip().
type lexer struct {
	src    string
	delims Delimiters
	errs   *ErrorList

	pos  int
	line int
	col  int
}

func newLexer(src string, delims Delimiters, errs *ErrorList) *lexer {
	return &lexer{src: src, delims: delims, errs: errs, pos: 0, line: 1, col: 1}
}

// detectUNA inspects the start of src for a "UNA" service string advice
// segment (a fixed 9-byte layout: "UNA" followed by the 6 delimiter
// characters, in order: component separator, element separator, decimal
// mark, release character, reserved, segment terminator). If present, it
// returns the delimiters it defines and consumed=9; otherwise it returns
// the documented defaults and consumed=0.
func detectUNA(src string) (delims Delimiters, consumed int, present bool) {
	if len(src) < 9 || src[:3] != "UNA" {
		return DefaultDelimiters(), 0, false
	}
	chars := src[3:9]
	return Delimiters{
		Component:  chars[0],
		Element:    chars[1],
		Decimal:    chars[2],
		Release:    chars[3],
		Reserved:   chars[4],
		Terminator: chars[5],
	}, 9, true
}

func (l *lexer) eof() bool { return l.pos >= len(l.src) }

func (l *lexer) position() Position {
	return Position{Offset: l.pos, Line: l.line, Column: l.col}
}

// advance consumes and returns the next byte, updating line/column.
func (l *lexer) advance() byte {
	c := l.src[l.pos]
	l.pos++
	if c == '\n' {
		l.line++
		l.col = 1
	} else {
		l.col++
	}
	return c
}

// skip consumes exactly n bytes (or up to EOF), updating position tracking.
// Used to step over a UNA service string advice segment once its delimiters
// have been read directly out of l.src.
func (l *lexer) skip(n int) {
	for i := 0; i < n && !l.eof(); i++ {
		l.advance()
	}
}

// skipInterSegmentWhitespace skips CR/LF bytes between segments. Real-world
// interchanges commonly put each segment on its own line purely for human
// readability; EDIFACT itself has no concept of significant newlines.
func (l *lexer) skipInterSegmentWhitespace() {
	for !l.eof() {
		c := l.src[l.pos]
		if c != '\r' && c != '\n' {
			return
		}
		l.advance()
	}
}

// next returns the next token. At EOF it returns tokenEOF repeatedly without
// erroring, so callers can loop on it safely.
func (l *lexer) next() token {
	if l.eof() {
		return token{Kind: tokenEOF, Pos: l.position()}
	}

	start := l.position()
	c := l.src[l.pos]

	switch c {
	case l.delims.Component:
		l.advance()
		return token{Kind: tokenComponentSep, Pos: start}
	case l.delims.Element:
		l.advance()
		return token{Kind: tokenElementSep, Pos: start}
	case l.delims.Terminator:
		l.advance()
		return token{Kind: tokenSegmentTerminator, Pos: start}
	}

	// Data run: everything up to (but not including) the next unescaped
	// separator/terminator, EOF, or literal newline.
	startOffset := l.pos
	for !l.eof() {
		c := l.src[l.pos]
		if c == l.delims.Release {
			relPos := l.position()
			l.advance() // consume the release character itself
			if l.eof() {
				l.errs.Add(relPos, SeverityError, "release character %q at end of input with no character to escape", l.delims.Release)
				break
			}
			l.advance() // consume the escaped character literally
			continue
		}
		if c == l.delims.Component || c == l.delims.Element || c == l.delims.Terminator || c == '\r' || c == '\n' {
			break
		}
		l.advance()
	}

	return token{Kind: tokenData, Raw: l.src[startOffset:l.pos], Pos: start}
}
