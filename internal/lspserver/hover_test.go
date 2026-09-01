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
