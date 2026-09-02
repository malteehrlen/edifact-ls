package lspserver

import (
	"strings"
	"testing"

	protocol "github.com/tliron/glsp/protocol_3_16"
)

func hoverAt(t *testing.T, text string, line, character protocol.UInteger) *protocol.Hover {
	t.Helper()
	st := &state{documents: map[protocol.DocumentUri]string{}}
	const uri = "file:///t.edi"
	st.documents[uri] = text

	hover, err := st.textDocumentHover(nil, &protocol.HoverParams{
		TextDocumentPositionParams: protocol.TextDocumentPositionParams{
			TextDocument: protocol.TextDocumentIdentifier{URI: uri},
			Position:     protocol.Position{Line: line, Character: character},
		},
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	return hover
}

func TestTextDocumentHoverServiceSegment(t *testing.T) {
	hover := hoverAt(t, "UNH+1+ORDERS:D:96A:UN'\nBGM+220'\n", 0, 1)
	if hover == nil {
		t.Fatal("got nil hover, want content for UNH")
	}
	content, ok := hover.Contents.(protocol.MarkupContent)
	if !ok {
		t.Fatalf("Contents = %T, want protocol.MarkupContent", hover.Contents)
	}
	if content.Kind != protocol.MarkupKindMarkdown {
		t.Errorf("Kind = %v, want markdown", content.Kind)
	}
	if !strings.Contains(content.Value, "UNH") || !strings.Contains(content.Value, "Message header") {
		t.Errorf("Value = %q, want it to mention UNH and its name", content.Value)
	}
}

func TestTextDocumentHoverBusinessSegment(t *testing.T) {
	hover := hoverAt(t, "UNH+1+ORDERS:D:96A:UN'\nBGM+220'\n", 1, 1)
	if hover == nil {
		t.Fatal("got nil hover, want content for BGM")
	}
	content := hover.Contents.(protocol.MarkupContent)
	if !strings.Contains(content.Value, "Beginning of message") {
		t.Errorf("Value = %q, want it to mention BGM's name", content.Value)
	}
}

func TestTextDocumentHoverAnywhereOnSegmentLine(t *testing.T) {
	// Hovering an element, not just the tag itself, still resolves to the
	// segment's description -- see segmentTagAt's doc comment.
	hover := hoverAt(t, "BGM+220'", 0, 6)
	if hover == nil {
		t.Fatal("got nil hover, want content when hovering an element within the BGM segment")
	}
}

func TestTextDocumentHoverUNAServiceStringAdvice(t *testing.T) {
	hover := hoverAt(t, "UNA:+.? 'UNH+1'", 0, 1)
	if hover == nil {
		t.Fatal("got nil hover, want content for UNA")
	}
	content := hover.Contents.(protocol.MarkupContent)
	if !strings.Contains(content.Value, "Service string advice") {
		t.Errorf("Value = %q, want it to mention UNA's name", content.Value)
	}
}

func TestTextDocumentHoverUnrecognizedTagReturnsNil(t *testing.T) {
	// XXX is a syntactically valid tag (3 uppercase letters) but not in the
	// description table.
	if hover := hoverAt(t, "XXX+1'", 0, 1); hover != nil {
		t.Fatalf("got %+v, want nil for an unrecognized tag", hover)
	}
}

func TestTextDocumentHoverOutsideAnySegmentReturnsNil(t *testing.T) {
	// Line 1 is a blank separator line between the two segments -- no
	// segment's span covers it.
	if hover := hoverAt(t, "UNH+1'\n\nBGM+220'\n", 1, 0); hover != nil {
		t.Fatalf("got %+v, want nil between segments", hover)
	}
}

// iftmcsFixture is a real, minimal IFTMCS D.21A message (see
// internal/edifact/iftmcs_d21a.go): BGM is top-level and mandatory; TOD
// is the mandatory leading segment of "segment group 2" (an optional
// group, so a message can go straight from BGM to TOD without segment
// group 1's LOC first). TOD isn't in the tier-1 tag description table
// (segments.go), so it also exercises hover's "group context alone, no
// tier-1 description" fallback.
const iftmcsFixture = "UNH+1+IFTMCS:D:21A:UN'\nBGM+320'\nTOD+1'\nUNT+4+1'\n"

func TestTextDocumentHoverGroupContextOnUngroupedTierOneTag(t *testing.T) {
	hover := hoverAt(t, iftmcsFixture, 1, 1) // BGM
	if hover == nil {
		t.Fatal("got nil hover, want content for BGM")
	}
	content := hover.Contents.(protocol.MarkupContent)
	if !strings.Contains(content.Value, "Beginning of message") {
		t.Errorf("Value = %q, want the tier-1 description", content.Value)
	}
	if strings.Contains(content.Value, "segment group") {
		t.Errorf("Value = %q, want no group context for a top-level segment", content.Value)
	}
}

func TestTextDocumentHoverGroupContextOnUnknownTag(t *testing.T) {
	hover := hoverAt(t, iftmcsFixture, 2, 1) // TOD
	if hover == nil {
		t.Fatal("got nil hover, want group-context content for TOD even with no tier-1 description")
	}
	content := hover.Contents.(protocol.MarkupContent)
	if !strings.Contains(content.Value, "TOD") {
		t.Errorf("Value = %q, want it to mention TOD", content.Value)
	}
	if !strings.Contains(content.Value, "Part of segment group 2") {
		t.Errorf("Value = %q, want it to name segment group 2", content.Value)
	}
}

func TestTextDocumentHoverNoGroupContextForUNH(t *testing.T) {
	// UNH identifies the message but isn't itself part of its own body,
	// so it shouldn't be attributed to any segment group.
	hover := hoverAt(t, iftmcsFixture, 0, 1)
	if hover == nil {
		t.Fatal("got nil hover, want the tier-1 UNH description")
	}
	content := hover.Contents.(protocol.MarkupContent)
	if strings.Contains(content.Value, "segment group") {
		t.Errorf("Value = %q, want no group context for UNH itself", content.Value)
	}
}

func TestTextDocumentHoverNoGroupContextForUnregisteredMessageType(t *testing.T) {
	// ZZZZZZ has no registered schema, so group context can't be computed
	// -- LOC is in the tier-1 table, so hover should still return that
	// alone, not nil.
	text := "UNH+1+ZZZZZZ:D:99Z:UN'\nBGM+320'\nLOC+1'\nUNT+4+1'\n"
	hover := hoverAt(t, text, 2, 1) // LOC
	if hover == nil {
		t.Fatal("got nil hover, want the tier-1 LOC description even with no registered schema")
	}
	content := hover.Contents.(protocol.MarkupContent)
	if !strings.Contains(content.Value, "Place/location identification") {
		t.Errorf("Value = %q, want the tier-1 LOC description", content.Value)
	}
	if strings.Contains(content.Value, "segment group") {
		t.Errorf("Value = %q, want no group context for an unregistered message type", content.Value)
	}
}

func TestTextDocumentHoverUnknownDocumentReturnsNil(t *testing.T) {
	st := &state{documents: map[protocol.DocumentUri]string{}}
	hover, err := st.textDocumentHover(nil, &protocol.HoverParams{
		TextDocumentPositionParams: protocol.TextDocumentPositionParams{
			TextDocument: protocol.TextDocumentIdentifier{URI: "file:///never-opened.edi"},
			Position:     protocol.Position{Line: 0, Character: 0},
		},
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if hover != nil {
		t.Fatalf("got %+v, want nil for a document that was never opened", hover)
	}
}
