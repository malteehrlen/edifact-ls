package lspserver

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/malteehrlen/edifact-ls/internal/edifact"
	"github.com/tliron/glsp"
	protocol "github.com/tliron/glsp/protocol_3_16"
)

// textDocumentHover looks up which segment (or the UNA service string
// advice) the given position falls within, and combines whatever's
// available for it. Three tiers, checked most-specific first:
//   - tier 2: hovering a specific coded value (e.g. "9" in BGM+9) whose
//     component has a known code list -- see edifact-ls-6xaz. Takes
//     priority over the tag-level tiers below when it applies, since
//     it's strictly more specific to the exact position hovered.
//   - tier 1: a tag description (name + one-line description,
//     independent of message type)
//   - tier 3: message context (which segment group(s) this occurrence
//     falls inside, per its message's registered Schema) -- see
//     edifact-ls-pcm0
//
// Returns nil only when nothing at all is available, not a placeholder.
func (st *state) textDocumentHover(context *glsp.Context, params *protocol.HoverParams) (*protocol.Hover, error) {
	st.docsMu.Lock()
	text, ok := st.documents[params.TextDocument.URI]
	st.docsMu.Unlock()
	if !ok {
		return nil, nil
	}

	offset := lspPositionToOffset(text, params.Position)
	ic, _ := edifact.Parse(text)

	if content, ok := codedValueHoverAt(ic, offset); ok {
		return &protocol.Hover{
			Contents: protocol.MarkupContent{Kind: protocol.MarkupKindMarkdown, Value: content},
		}, nil
	}

	tag, found := segmentTagAt(ic, offset)
	if !found {
		return nil, nil
	}

	info, hasInfo := edifact.SegmentDescription(tag)
	groupPath := groupPathAt(ic, offset)

	if !hasInfo && groupPath == nil {
		return nil, nil
	}

	var sb strings.Builder
	if hasInfo {
		fmt.Fprintf(&sb, "**%s** -- %s\n\n%s", tag, info.Name, info.Description)
	} else {
		fmt.Fprintf(&sb, "**%s**", tag)
	}
	if groupPath != nil {
		if sb.Len() > 0 {
			sb.WriteString("\n\n")
		}
		sb.WriteString("Part of " + formatGroupPath(groupPath))
	}

	return &protocol.Hover{
		Contents: protocol.MarkupContent{
			Kind:  protocol.MarkupKindMarkdown,
			Value: sb.String(),
		},
	}, nil
}

// codedValueHoverAt returns markdown content for the coded value under
// offset, if the segment's element/component structure there is known
// (see segment_elements_data.go) and the actual text present is a
// recognized code in that component's registered code list (see
// codelist.go). ok is false for anything short of that -- including a
// component that's coded but whose actual value isn't a recognized code
// -- so a wrong or unrecognized value never gets asserted a false
// meaning; callers fall back to less specific hover content instead.
func codedValueHoverAt(ic *edifact.Interchange, offset int) (string, bool) {
	for _, seg := range ic.Segments {
		if offset < seg.Pos.Offset || offset >= seg.EndPos.Offset {
			continue
		}
		schema, ok := edifact.SegmentElementSchemaFor(seg.Tag)
		if !ok {
			return "", false
		}
		for ei, el := range seg.Elements {
			if ei >= len(schema.Elements) {
				break
			}
			es := schema.Elements[ei]
			for ci, c := range el.Components {
				start := c.Pos.Offset
				if offset < start || offset >= start+len(c.Raw) {
					continue
				}
				if ci >= len(es.Components) || es.Components[ci].CodeList == "" {
					return "", false
				}
				value := c.Value(ic.Delimiters)
				cv, ok := edifact.LookupCode(es.Components[ci].CodeList, value)
				if !ok {
					return "", false
				}
				return fmt.Sprintf("**%s** -- %s\n\n%s", value, cv.Name, cv.Description), true
			}
		}
		return "", false
	}
	return "", false
}

// groupPathAt returns the sequence of segment-group numbers (outermost
// first) that the segment spanning offset falls within, per its
// message's registered Schema, or nil if there's nothing to report --
// offset isn't within any real Segment (including because it's on the
// UNA advice), the segment is UNH/UNT/UNB/UNZ/UNG/UNE itself rather than
// part of a message body, its message type has no registered Schema, or
// the schema doesn't place it inside any group. See edifact-ls-pcm0.
func groupPathAt(ic *edifact.Interchange, offset int) []int {
	segIdx := segmentIndexAt(ic, offset)
	if segIdx < 0 {
		return nil
	}

	unhIdx := -1
	for i := segIdx; i >= 0; i-- {
		if ic.Segments[i].Tag == "UNH" {
			unhIdx = i
			break
		}
		if ic.Segments[i].Tag == "UNT" {
			break // segIdx sits outside any message body (or is UNT itself)
		}
	}
	if unhIdx < 0 || unhIdx == segIdx {
		return nil
	}

	untIdx := -1
	for i := unhIdx + 1; i < len(ic.Segments); i++ {
		if ic.Segments[i].Tag == "UNT" {
			untIdx = i
			break
		}
	}
	if untIdx < 0 || segIdx >= untIdx {
		return nil
	}

	id := edifact.MessageIDOf(&ic.Segments[unhIdx], ic.Delimiters)
	schema, ok := edifact.LookupSchema(id)
	if !ok {
		return nil
	}

	body := ic.Segments[unhIdx+1 : untIdx]
	return edifact.GroupPathAt(schema, body, segIdx-(unhIdx+1))
}

// segmentIndexAt returns the index into ic.Segments of whichever segment
// spans offset, or -1 if none does (including when offset falls on the
// UNA advice, which isn't itself a Segment).
func segmentIndexAt(ic *edifact.Interchange, offset int) int {
	for i, seg := range ic.Segments {
		if offset >= seg.Pos.Offset && offset < seg.EndPos.Offset {
			return i
		}
	}
	return -1
}

// formatGroupPath renders a GroupPathAt result the way UN/EDIFACT specs
// themselves name segment groups, outermost first.
func formatGroupPath(path []int) string {
	parts := make([]string, len(path))
	for i, n := range path {
		parts[i] = "segment group " + strconv.Itoa(n)
	}
	return strings.Join(parts, " > ")
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
