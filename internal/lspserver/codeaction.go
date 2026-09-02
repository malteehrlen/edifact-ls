package lspserver

import (
	"github.com/malteehrlen/edifact-ls/internal/edifact"
	"github.com/tliron/glsp"
	protocol "github.com/tliron/glsp/protocol_3_16"
)

// textDocumentCodeAction returns a quickfix CodeAction for each diagnostic
// overlapping the requested range that carries a mechanical Fix (see
// edifact.Error.Fix -- only a handful of diagnostic kinds have one; most
// diagnostics simply produce no action here, which is a normal "nothing to
// fix" result, not an error).
func (st *state) textDocumentCodeAction(context *glsp.Context, params *protocol.CodeActionParams) (any, error) {
	st.docsMu.Lock()
	text, ok := st.documents[params.TextDocument.URI]
	st.docsMu.Unlock()
	if !ok {
		return nil, nil
	}

	_, errs := edifact.Validate(text)

	var actions []protocol.CodeAction
	for _, e := range errs {
		if e.Fix == nil || !rangesOverlap(params.Range, errorRange(text, e.Pos)) {
			continue
		}

		// The Fix was derived from the same fresh Validate call as the
		// diagnostic it fixes, so this should always hold -- guard anyway
		// so a violated assumption skips the action instead of splicing a
		// stale span into the buffer.
		start := e.Fix.Pos.Offset
		end := start + len(e.Fix.OldText)
		if start < 0 || end > len(text) || text[start:end] != e.Fix.OldText {
			continue
		}

		kind := protocol.CodeActionKindQuickFix
		preferred := true
		actions = append(actions, protocol.CodeAction{
			Title:       e.Fix.Title,
			Kind:        &kind,
			Diagnostics: []protocol.Diagnostic{toDiagnostic(text, e)},
			IsPreferred: &preferred,
			Edit: &protocol.WorkspaceEdit{
				Changes: map[protocol.DocumentUri][]protocol.TextEdit{
					params.TextDocument.URI: {
						{Range: fixRange(text, e.Fix), NewText: e.Fix.NewText},
					},
				},
			},
		})
	}
	return actions, nil
}

// fixRange builds the exact LSP Range a Fix replaces: from Fix.Pos to
// Fix.Pos plus the byte length of Fix.OldText.
func fixRange(text string, fix *edifact.Fix) protocol.Range {
	start := offsetToLSPPosition(text, fix.Pos.Offset)
	end := offsetToLSPPosition(text, fix.Pos.Offset+len(fix.OldText))
	return protocol.Range{Start: start, End: end}
}

// rangesOverlap reports whether a and b share at least one position,
// treating touching endpoints as overlapping (a request range is often a
// zero-width cursor position right at a diagnostic's edge).
func rangesOverlap(a, b protocol.Range) bool {
	return !positionBefore(a.End, b.Start) && !positionBefore(b.End, a.Start)
}

func positionBefore(a, b protocol.Position) bool {
	if a.Line != b.Line {
		return a.Line < b.Line
	}
	return a.Character < b.Character
}
