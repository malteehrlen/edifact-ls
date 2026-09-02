---
# edifact-ls-n3yd
title: textDocument/codeAction handler
status: todo
type: feature
priority: normal
created_at: 2026-09-02T11:38:49Z
updated_at: 2026-09-02T11:38:49Z
parent: edifact-ls-x3pb
blocked_by:
    - edifact-ls-338y
---

# Description

Wire `textDocument/codeAction` into the LSP server, same shape as
hover/formatting: a handler method on `state`, capability auto-derived
from the registered handler (no manual capability advertisement).

```go
// internal/lspserver/codeaction.go
func (st *state) textDocumentCodeAction(ctx *glsp.Context, params *protocol.CodeActionParams) (any, error) {
    _, errs := edifact.Validate(text)
    var actions []protocol.CodeAction
    for _, e := range errs {
        if e.Fix == nil || !rangesOverlap(params.Range, errorRange(text, e.Pos)) {
            continue
        }
        actions = append(actions, protocol.CodeAction{
            Title:       e.Fix.Title,
            Kind:        &protocol.CodeActionKindQuickFix,
            Diagnostics: []protocol.Diagnostic{ /* the matching one */ },
            IsPreferred: ptr(true),
            Edit: &protocol.WorkspaceEdit{Changes: map[string][]protocol.TextEdit{
                uri: {{Range: fixRange(text, e.Fix), NewText: e.Fix.NewText}},
            }},
        })
    }
    return actions, nil
}
```

Before applying `Fix.NewText`, sanity-check `Fix.OldText` still matches
the current buffer text at `Fix.Pos` (the diagnostic could be stale if
the buffer changed since the last publish) -- skip that action rather
than emit a corrupting edit if it doesn't match.

# Acceptance Criteria

[ ] `textDocument/codeAction` wired into the LSP server, capability
auto-advertised
[ ] Returns a quickfix `CodeAction` (title from `Fix.Title`,
`IsPreferred: true`, correct `WorkspaceEdit`) for each diagnostic
overlapping the requested range that carries a non-nil `Fix`
[ ] Returns no actions for a range with no fixable diagnostic (not an
error -- just an empty/nil result)
[ ] Skips (does not emit) an action whose `Fix.OldText` no longer
matches the current buffer text at `Fix.Pos`, rather than emitting an
edit that would corrupt the document
[ ] Unit tests (direct handler calls, matching `hover_test.go`'s
pattern): a known-fixable document returns the right `[]CodeAction`;
a document with only non-fixable diagnostics returns none
[ ] e2e check: open a fixture with a redundant UNA, request
`textDocument/codeAction` via `buf_request_sync`, apply the returned
edit, assert the buffer changed as expected, in a real nvim session;
same for one envelope-mismatch fixture
[ ] Full suite (`make test`) and e2e harness (`make test-e2e`) green
