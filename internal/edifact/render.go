package edifact

import "strings"

// Render serializes an Interchange back to EDIFACT text, either as the
// human-readable multiline form (one segment per line -- what
// textDocument/formatting produces) or as compact "wire" format (segments
// joined by their terminator only, no inserted whitespace -- how an
// interchange actually looks in transit).
//
// Both modes reproduce component data exactly as parsed, via each
// Component's Raw field (which already preserves any release-character
// escaping verbatim), so re-parsing the output always yields the same data
// as the original input: Render is a pure re-layout, never a re-encoding.
func Render(ic *Interchange, multiline bool) string {
	var b strings.Builder
	d := ic.Delimiters

	if ic.UNA != nil {
		b.WriteString(ic.UNA.Raw)
		if multiline {
			b.WriteByte('\n')
		}
	}

	for _, seg := range ic.Segments {
		renderSegment(&b, seg, d)
		if multiline {
			b.WriteByte('\n')
		}
	}

	return b.String()
}

func renderSegment(b *strings.Builder, seg Segment, d Delimiters) {
	b.WriteString(seg.Tag)
	for _, el := range seg.Elements {
		b.WriteByte(d.Element)
		for i, c := range el.Components {
			if i > 0 {
				b.WriteByte(d.Component)
			}
			b.WriteString(c.Raw)
		}
	}
	b.WriteByte(d.Terminator)
}
