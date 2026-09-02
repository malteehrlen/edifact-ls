package lspserver

import (
	"strings"
	"testing"

	"github.com/malteehrlen/edifact-ls/internal/edifact"
	protocol "github.com/tliron/glsp/protocol_3_16"
)

func codeActionsAt(t *testing.T, text string, startLine, startChar, endLine, endChar protocol.UInteger) []protocol.CodeAction {
	t.Helper()
	st := &state{documents: map[protocol.DocumentUri]string{}}
	const uri = "file:///t.edi"
	st.documents[uri] = text

	result, err := st.textDocumentCodeAction(nil, &protocol.CodeActionParams{
		TextDocument: protocol.TextDocumentIdentifier{URI: uri},
		Range: protocol.Range{
			Start: protocol.Position{Line: startLine, Character: startChar},
			End:   protocol.Position{Line: endLine, Character: endChar},
		},
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if result == nil {
		return nil
	}
	actions, ok := result.([]protocol.CodeAction)
	if !ok {
		t.Fatalf("result = %T, want []protocol.CodeAction", result)
	}
	return actions
}

// applyWorkspaceEdit applies a's single-file edit to text the way a real
// client would, so tests can assert on the resulting document rather than
// just on the edit's shape.
func applyWorkspaceEdit(t *testing.T, text string, a protocol.CodeAction) string {
	t.Helper()
	if a.Edit == nil {
		t.Fatal("CodeAction has no Edit")
	}
	edits := a.Edit.Changes["file:///t.edi"]
	if len(edits) != 1 {
		t.Fatalf("got %d edits, want 1", len(edits))
	}
	e := edits[0]
	start := lspPositionToOffset(text, e.Range.Start)
	end := lspPositionToOffset(text, e.Range.End)
	return text[:start] + e.NewText + text[end:]
}

func TestTextDocumentCodeActionRedundantUNA(t *testing.T) {
	text := "UNA:+.? 'UNB+UNOA:1+S+R+201001:1200+1'UNH+1+ORDERS:D:96A:UN'BGM+220'UNT+3+1'UNZ+1+1'"
	actions := codeActionsAt(t, text, 0, 0, 0, 9)
	if len(actions) != 1 {
		t.Fatalf("got %d actions, want 1: %+v", len(actions), actions)
	}
	a := actions[0]
	if a.Kind == nil || *a.Kind != protocol.CodeActionKindQuickFix {
		t.Errorf("Kind = %v, want quickfix", a.Kind)
	}
	if a.IsPreferred == nil || !*a.IsPreferred {
		t.Errorf("IsPreferred = %v, want true", a.IsPreferred)
	}
	if !strings.Contains(a.Title, "UNA") {
		t.Errorf("Title = %q, want it to mention UNA", a.Title)
	}
	fixed := applyWorkspaceEdit(t, text, a)
	if strings.HasPrefix(fixed, "UNA") {
		t.Errorf("fixed text still starts with UNA: %q", fixed)
	}
	_, errs := edifact.Validate(fixed)
	if errs.HasErrors() {
		t.Errorf("unexpected errors after applying fix: %v", errs)
	}
}

func TestTextDocumentCodeActionEnvelopeCountMismatch(t *testing.T) {
	text := "UNB+UNOA:1+S+R+201001:1200+1'UNH+1+ORDERS:D:96A:UN'BGM+220'UNT+99+1'UNZ+1+1'"
	// The diagnostic (and so its displayed range) anchors on the UNT
	// segment's own position, not the wrong value's -- find "UNT" rather
	// than hand-counting, so the test doesn't rot if the fixture changes.
	idx := strings.Index(text, "UNT+99")
	char := protocol.UInteger(idx)
	actions := codeActionsAt(t, text, 0, char, 0, char+6)
	if len(actions) != 1 {
		t.Fatalf("got %d actions, want 1: %+v", len(actions), actions)
	}
	a := actions[0]
	fixed := applyWorkspaceEdit(t, text, a)
	_, errs := edifact.Validate(fixed)
	if errs.HasErrors() {
		t.Errorf("unexpected errors after applying fix: %v (fixed text: %q)", errs, fixed)
	}
	if strings.Contains(fixed, "UNT+99") {
		t.Errorf("fixed text still contains the wrong count: %q", fixed)
	}
}

func TestTextDocumentCodeActionUNAAvailableAnywhereInItsSpan(t *testing.T) {
	// Regression: the request range used to be checked only against the
	// diagnostic's displayed range, which for a segment-anchored
	// diagnostic like this is a single byte -- making the action
	// effectively unreachable unless the cursor sat on that exact
	// character. It must now resolve from any position across the whole
	// 9-byte UNA advice this action actually removes.
	text := "UNA:+.? 'UNB+UNOA:1+S+R+201001:1200+1'UNH+1+ORDERS:D:96A:UN'BGM+220'UNT+3+1'UNZ+1+1'"
	for _, char := range []protocol.UInteger{0, 3, 5, 8} {
		actions := codeActionsAt(t, text, 0, char, 0, char+1)
		if len(actions) != 1 {
			t.Errorf("character=%d: got %d actions, want 1", char, len(actions))
		}
	}
}

func TestTextDocumentCodeActionEnvelopeMismatchAvailableAtWrongValue(t *testing.T) {
	// The Fix's own span is the wrong value itself (e.g. "99" in
	// "UNT+99"), not the segment tag the diagnostic is anchored at -- a
	// request range sitting on the wrong value, not just on "UNT", must
	// also resolve the action.
	text := "UNB+UNOA:1+S+R+201001:1200+1'UNH+1+ORDERS:D:96A:UN'BGM+220'UNT+99+1'UNZ+1+1'"
	idx := strings.Index(text, "99")
	char := protocol.UInteger(idx)
	actions := codeActionsAt(t, text, 0, char, 0, char+1)
	if len(actions) != 1 {
		t.Fatalf("got %d actions, want 1 when the cursor is on the wrong value itself: %+v", len(actions), actions)
	}
}

func TestTextDocumentCodeActionNoActionsForNonFixableDiagnostic(t *testing.T) {
	// Missing UNT: a real diagnostic, but nothing to safely insert.
	text := "UNB+UNOA:1+S+R+201001:1200+1'UNH+1+ORDERS:D:96A:UN'BGM+220'UNZ+1+1'"
	actions := codeActionsAt(t, text, 0, 0, 0, 100)
	if len(actions) != 0 {
		t.Fatalf("got %d actions, want 0: %+v", len(actions), actions)
	}
}

func TestTextDocumentCodeActionNoActionsOutsideRequestedRange(t *testing.T) {
	text := "UNA:+.? 'UNB+UNOA:1+S+R+201001:1200+1'UNH+1+ORDERS:D:96A:UN'BGM+220'UNT+3+1'UNZ+1+1'"
	// A range far past the UNA advice's own position.
	actions := codeActionsAt(t, text, 0, 70, 0, 80)
	if len(actions) != 0 {
		t.Fatalf("got %d actions, want 0 outside the UNA diagnostic's range: %+v", len(actions), actions)
	}
}

func TestTextDocumentCodeActionUnknownDocumentReturnsNil(t *testing.T) {
	st := &state{documents: map[protocol.DocumentUri]string{}}
	result, err := st.textDocumentCodeAction(nil, &protocol.CodeActionParams{
		TextDocument: protocol.TextDocumentIdentifier{URI: "file:///never-opened.edi"},
		Range:        protocol.Range{},
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if result != nil {
		t.Fatalf("got %+v, want nil for a document that was never opened", result)
	}
}
