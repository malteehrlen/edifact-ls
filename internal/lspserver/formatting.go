package lspserver

import (
	"github.com/malteehrlen/edifact-ls/internal/edifact"
	"github.com/tliron/glsp"
	protocol "github.com/tliron/glsp/protocol_3_16"
)

func (st *state) textDocumentFormatting(context *glsp.Context, params *protocol.DocumentFormattingParams) ([]protocol.TextEdit, error) {
	st.docsMu.Lock()
	text, ok := st.documents[params.TextDocument.URI]
	st.docsMu.Unlock()
	if !ok {
		return nil, nil
	}

	ic, errs := edifact.Parse(text)
	if errs.HasErrors() {
		// Don't touch a document we can't fully make sense of.
		return nil, nil
	}

	formatted := edifact.Render(ic, true)
	if formatted == text {
		return nil, nil
	}

	return []protocol.TextEdit{wholeDocumentReplace(text, formatted)}, nil
}

// wholeDocumentReplace builds a TextEdit that replaces the entire document
// (from its start to its end) with newText.
func wholeDocumentReplace(oldText, newText string) protocol.TextEdit {
	return protocol.TextEdit{
		Range: protocol.Range{
			Start: protocol.Position{Line: 0, Character: 0},
			End:   offsetToLSPPosition(oldText, len(oldText)),
		},
		NewText: newText,
	}
}
