package lspserver

import (
	"fmt"

	"github.com/malteehrlen/edifact-ls/internal/edifact"
	"github.com/tliron/glsp"
	protocol "github.com/tliron/glsp/protocol_3_16"
)

// textDocumentHover looks up which segment (or the UNA service string
// advice) the given position falls within and, if its tag has a known
// description, returns that as markdown. A position outside any segment,
// or a segment whose tag isn't in the description table, returns nil --
// no hover, not an empty/placeholder one.
func (st *state) textDocumentHover(context *glsp.Context, params *protocol.HoverParams) (*protocol.Hover, error) {
	st.docsMu.Lock()
	text, ok := st.documents[params.TextDocument.URI]
	st.docsMu.Unlock()
	if !ok {
		return nil, nil
	}

	offset := lspPositionToOffset(text, params.Position)
	ic, _ := edifact.Parse(text)

	tag, found := segmentTagAt(ic, offset)
	if !found {
		return nil, nil
	}

	info, ok := edifact.SegmentDescription(tag)
	if !ok {
		return nil, nil
	}

	return &protocol.Hover{
		Contents: protocol.MarkupContent{
			Kind:  protocol.MarkupKindMarkdown,
			Value: fmt.Sprintf("**%s** -- %s\n\n%s", tag, info.Name, info.Description),
		},
	}, nil
}

// segmentTagAt returns the tag of whichever segment -- or the UNA service
// string advice, which isn't itself a Segment -- spans the given byte
// offset, if any. Matching the whole segment's span (not just its tag
// characters) means hovering anywhere on the segment's line resolves it,
// not only the exact three tag characters.
func segmentTagAt(ic *edifact.Interchange, offset int) (string, bool) {
	if ic.UNA != nil {
		start := ic.UNA.Pos.Offset
		if offset >= start && offset < start+len(ic.UNA.Raw) {
			return "UNA", true
		}
	}
	for _, seg := range ic.Segments {
		if offset >= seg.Pos.Offset && offset < seg.EndPos.Offset {
			return seg.Tag, true
		}
	}
	return "", false
}
