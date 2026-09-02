---
# edifact-ls-n3yd
title: textDocument/codeAction handler
status: completed
type: feature
priority: normal
created_at: 2026-09-02T11:38:49Z
updated_at: 2026-09-02T11:46:09Z
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

[x] `textDocument/codeAction` wired into the LSP server, capability
auto-advertised
[x] Returns a quickfix `CodeAction` (title from `Fix.Title`,
`IsPreferred: true`, correct `WorkspaceEdit`) for each diagnostic
overlapping the requested range that carries a non-nil `Fix`
[x] Returns no actions for a range with no fixable diagnostic (not an
error -- just an empty/nil result)
[x] Skips (does not emit) an action whose `Fix.OldText` no longer
matches the current buffer text at `Fix.Pos`, rather than emitting an
edit that would corrupt the document
[x] Unit tests (direct handler calls, matching `hover_test.go`'s
pattern): a known-fixable document returns the right `[]CodeAction`;
a document with only non-fixable diagnostics returns none
[x] e2e check: open a fixture with a redundant UNA, request
`textDocument/codeAction` via `buf_request_sync`, apply the returned
edit, assert the buffer changed as expected, in a real nvim session;
same for one envelope-mismatch fixture
[x] Full suite (`make test`) and e2e harness (`make test-e2e`) green

## Summary of Changes

internal/lspserver/codeaction.go (new): `textDocumentCodeAction`
re-validates the current buffer text (same `edifact.Validate` pipeline
diagnostics already use), and for each `Error` with a non-nil `Fix`
whose diagnostic range (`errorRange(text, e.Pos)` -- the same range the
client sees the squiggle at) overlaps the requested range, emits a
quickfix `CodeAction` with the `Fix`'s title, `IsPreferred: true`, and
a single-file `WorkspaceEdit`. Before emitting, it re-slices the
buffer at `Fix.Pos`/`len(Fix.OldText)` and confirms it still equals
`Fix.OldText`, skipping the action rather than risk splicing a stale
span if that assumption is ever violated. `rangesOverlap` treats
touching endpoints as overlapping, since the request range is often a
zero-width cursor position right at a diagnostic's edge.

internal/lspserver/diagnostics.go: factored the per-error ->
`protocol.Diagnostic` translation out into `toDiagnostic`, reused by
`codeaction.go` so a `CodeAction.Diagnostics` entry is built the exact
same way the client already received it, rather than a second,
possibly-drifting translation.

internal/lspserver/lspserver.go: one line, `TextDocumentCodeAction:
st.textDocumentCodeAction`, in the handler struct -- capability
advertisement is derived automatically from the handler being set, per
this project's established pattern (see the comment above the handler
struct).

7 new unit tests in codeaction_test.go, including one that applies the
returned edit and re-validates the result to confirm the fix actually
resolves the diagnostic (not just that some CodeAction exists), one
confirming no actions for a non-fixable diagnostic, and one confirming
range-filtering (no action when the requested range doesn't overlap
the diagnostic).

Two new e2e checks (a new `check_code_action` helper in
scripts/e2e_check.lua, applying the action via
`vim.lsp.util.apply_workspace_edit` and then re-checking
`vim.diagnostic.get(0)` for errors) against a reused fixture
(testdata/lint-info.edi, redundant UNA) and a new one
(testdata/envelope-count-mismatch.edi). Both confirm, in a real nvim
session, that requesting a code action and applying it actually
resolves the underlying problem. Full suite (`make test`) and e2e
harness (`make test-e2e`, 37 checks) pass.
