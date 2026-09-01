package lspserver

import (
	"testing"

	protocol "github.com/tliron/glsp/protocol_3_16"
)

func TestTextDocumentFormattingMinimalInterchange(t *testing.T) {
	st := &state{documents: map[protocol.DocumentUri]string{}}
	const uri = "file:///t.edi"
	st.documents[uri] = "UNH+1'"

	edits, err := st.textDocumentFormatting(nil, &protocol.DocumentFormattingParams{
		TextDocument: protocol.TextDocumentIdentifier{URI: uri},
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(edits) != 1 {
		t.Fatalf("got %d edits, want 1: %+v", len(edits), edits)
	}
	if want := "UNH+1'\n"; edits[0].NewText != want {
		t.Errorf("NewText = %q, want %q", edits[0].NewText, want)
	}
}

func TestTextDocumentFormattingCompositeElements(t *testing.T) {
	st := &state{documents: map[protocol.DocumentUri]string{}}
	const uri = "file:///t.edi"
	st.documents[uri] = "UNH+1+ORDERS:D:96A:UN'DTM+137:20100101:102'"

	edits, err := st.textDocumentFormatting(nil, &protocol.DocumentFormattingParams{
		TextDocument: protocol.TextDocumentIdentifier{URI: uri},
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(edits) != 1 {
		t.Fatalf("got %d edits, want 1: %+v", len(edits), edits)
	}
	want := "UNH+1+ORDERS:D:96A:UN'\nDTM+137:20100101:102'\n"
	if edits[0].NewText != want {
		t.Errorf("NewText = %q, want %q", edits[0].NewText, want)
	}
}

func TestTextDocumentFormattingAlreadyFormattedIsNoOp(t *testing.T) {
	st := &state{documents: map[protocol.DocumentUri]string{}}
	const uri = "file:///t.edi"
	st.documents[uri] = "UNH+1'\nBGM+220'\n"

	edits, err := st.textDocumentFormatting(nil, &protocol.DocumentFormattingParams{
		TextDocument: protocol.TextDocumentIdentifier{URI: uri},
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if edits != nil {
		t.Fatalf("got %d edits for already-formatted input, want none: %+v", len(edits), edits)
	}
}

func TestTextDocumentFormattingIdempotentViaTwoCalls(t *testing.T) {
	st := &state{documents: map[protocol.DocumentUri]string{}}
	const uri = "file:///t.edi"
	st.documents[uri] = "UNH+1+ORDERS:D:96A:UN'BGM+220+ORDER123+9'DTM+137:20100101:102'"

	edits, err := st.textDocumentFormatting(nil, &protocol.DocumentFormattingParams{
		TextDocument: protocol.TextDocumentIdentifier{URI: uri},
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(edits) != 1 {
		t.Fatalf("got %d edits, want 1: %+v", len(edits), edits)
	}

	st.documents[uri] = edits[0].NewText
	edits2, err := st.textDocumentFormatting(nil, &protocol.DocumentFormattingParams{
		TextDocument: protocol.TextDocumentIdentifier{URI: uri},
	})
	if err != nil {
		t.Fatalf("unexpected error on second format: %v", err)
	}
	if edits2 != nil {
		t.Fatalf("formatting already-formatted output produced edits: %+v", edits2)
	}
}

func TestTextDocumentFormattingNoOpOnUnparseableDocument(t *testing.T) {
	st := &state{documents: map[protocol.DocumentUri]string{}}
	const uri = "file:///t.edi"
	// A dangling release character makes this a lexical error, not just an
	// envelope one -- Parse itself reports it.
	st.documents[uri] = "UNH+1?"

	edits, err := st.textDocumentFormatting(nil, &protocol.DocumentFormattingParams{
		TextDocument: protocol.TextDocumentIdentifier{URI: uri},
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if edits != nil {
		t.Fatalf("got edits for an unparseable document, want none: %+v", edits)
	}
}
