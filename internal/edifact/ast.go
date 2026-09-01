package edifact

// Component is one component of a (possibly composite) data element. Raw
// preserves the exact source text, including any release-character escape
// sequences, so formatting can reproduce it losslessly; use Value to get
// the unescaped data.
type Component struct {
	Raw string
	Pos Position
}

// Value returns the component's data with release-character escapes
// resolved (e.g. "12?+34" with release '?' becomes "12+34").
func (c Component) Value(d Delimiters) string {
	return Unescape(c.Raw, d.Release)
}

// Unescape resolves release-character escape sequences in raw source text:
// each release byte plus the byte following it is replaced by that
// following byte alone.
func Unescape(raw string, release byte) string {
	i := indexByte(raw, release)
	if i < 0 {
		return raw
	}
	b := make([]byte, 0, len(raw))
	for i := 0; i < len(raw); i++ {
		if raw[i] == release && i+1 < len(raw) {
			i++
			b = append(b, raw[i])
			continue
		}
		b = append(b, raw[i])
	}
	return string(b)
}

func indexByte(s string, b byte) int {
	for i := 0; i < len(s); i++ {
		if s[i] == b {
			return i
		}
	}
	return -1
}

// Element is one data element of a segment: a single component for a
// simple data element, or several for a composite data element.
type Element struct {
	Components []Component
	Pos        Position
}

// Segment is one EDIFACT segment: a tag followed by zero or more data
// elements, originally terminated by the segment terminator.
type Segment struct {
	Tag    string
	TagPos Position

	// TagControlNumbers holds the segment tag's optional, component-
	// separator-delimited control-number components (e.g. the "1" in
	// "GDS:1+..."), used for "explicit representation" of repeating
	// segments -- rare in modern usage (current UNSMs use "implicit"
	// representation, i.e. no control numbers), but part of the formal
	// syntax. Empty when absent, which is the common case.
	TagControlNumbers []string

	Elements []Element

	Pos    Position // position of the segment's first byte (the tag)
	EndPos Position // position just after the terminator (or EOF, if unterminated)
}

// Element0 returns the segment's element at the given 0-based index, or nil
// if the segment has no such element. Convenience for envelope validation
// and diagnostics, which frequently need to look at "element N" of a known
// segment shape.
func (s Segment) Element0(i int) *Element {
	if i < 0 || i >= len(s.Elements) {
		return nil
	}
	return &s.Elements[i]
}

// Component0 returns the first component's unescaped value of the element
// at the given 0-based index, or "" if absent.
func (s Segment) Component0(elementIndex int, d Delimiters) string {
	el := s.Element0(elementIndex)
	if el == nil || len(el.Components) == 0 {
		return ""
	}
	return el.Components[0].Value(d)
}

// UNAAdvice is the optional service string advice segment that redefines
// an interchange's delimiters. It has a fixed 9-byte layout ("UNA" plus 6
// delimiter characters) and, unlike every other segment, is not itself
// terminated by the segment terminator.
type UNAAdvice struct {
	Pos Position
	Raw string // always exactly 9 bytes: "UNA" + the 6 delimiter characters
}

// Interchange is the syntactic parse of a whole EDIFACT interchange: an
// optional UNA advice, delimiters (defaulted or from that advice), and the
// flat sequence of segments that followed. This layer is purely syntactic —
// it does not know that "UNB"/"UNZ"/"UNH"/"UNT" are structurally special;
// see ValidateEnvelopes for that.
type Interchange struct {
	UNA        *UNAAdvice
	Segments   []Segment
	Delimiters Delimiters
}
