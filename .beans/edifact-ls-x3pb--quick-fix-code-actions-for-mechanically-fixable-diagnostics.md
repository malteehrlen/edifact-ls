---
# edifact-ls-x3pb
title: Quick-fix code actions for mechanically-fixable diagnostics
status: completed
type: epic
priority: normal
created_at: 2026-09-02T11:38:49Z
updated_at: 2026-09-02T11:46:09Z
parent: edifact-ls-gdt6
---

# Goal

Address a subset of the diagnostics this project already produces
(`Lint`, `ValidateEnvelopes`, `ValidateMessageSchemas`,
`ValidateSegmentContent`, `Parse`) with LSP quick-fix code actions,
so the user can apply a fix directly from the editor instead of
hand-editing.

Not every diagnostic is a candidate: a code action needs a fix that's
mechanically derivable from information the validator already has, with
no guessing about user intent. Surveying every diagnostic source, only
two kinds qualify:

- **Redundant `UNA` service string advice** (`Lint`, info) -- the fix is
  simply deleting the segment.
- **Envelope count/reference mismatches** (`ValidateEnvelopes`: `UNT`,
  `UNE`, `UNZ` count and/or reference fields) -- the fix is replacing
  the wrong value with the value the validator already computed as
  correct while detecting the mismatch.

Everything else (reserved-tag warnings, missing mandatory
segments/groups/elements, exceeded repeats, syntax errors, unrecognized
tags) requires inserting or renaming something the validator cannot
derive on its own, so those stay diagnostics-only.

# Approach

1. Give `edifact.Error` two new optional fields: `Code` (a stable
   string id like `redundant-una` or `envelope-count-mismatch`, for
   `Diagnostic.Code` and for a code-action handler to match against
   without regex-parsing message prose) and `Fix` (`*Fix`, nil when not
   mechanically fixable) holding the exact `Pos`/`OldText`/`NewText` of
   the edit. Populate both at the exact call sites in `lint.go` and
   `envelope.go` that already know the fix -- no new derivation logic
   needed elsewhere. `envelope.go`'s six count/reference comparisons
   need a small mechanical change first: they currently only have the
   *value* string, not the component's `Pos`/`Raw`, so they need
   `Element0(i).Components[0]` instead of the current accessor to get
   an exact replaceable span.
2. Wire `textDocument/codeAction` into the LSP server the same way
   hover/formatting are wired (auto-derived capability, no manual
   advertisement): for each diagnostic overlapping the requested range
   that carries a `Fix`, emit a `CodeAction` with
   `Kind: CodeActionKindQuickFix`, `IsPreferred: true`, and a
   `WorkspaceEdit` built from the `Fix`.

# Explicit non-goals

- Any fix that requires inserting new content whose value isn't already
  known (missing mandatory segments/elements, exceeded repeats,
  unrecognized tags) -- deferred indefinitely, not just out of scope
  for this epic, since there's no way to do it without guessing.
- Refactoring-kind code actions (anything under `refactor.*` /
  `source.*`) -- this epic is quickfix-only.

# Acceptance Criteria

[x] `edifact.Error` carries optional `Code` and `Fix` fields; `Fix` is
populated for the redundant-`UNA` case and all six `UNT`/`UNE`/`UNZ`
count-or-reference mismatch cases, and left nil everywhere else
[x] `Diagnostic.Code` is populated in the LSP translation layer for
every diagnostic that has a `Code`, independent of whether code
actions exist yet
[x] `textDocument/codeAction` returns a quickfix `CodeAction` with a
correct `WorkspaceEdit` for each fixable diagnostic overlapping the
requested range, and returns nothing for ranges with no fixable
diagnostic
[x] Unit tests confirm the redundant-UNA fix and each envelope
mismatch fix independently
[x] e2e check: opening a fixture with a redundant UNA (or a wrong
envelope count), requesting `textDocument/codeAction` via
`buf_request_sync`, applying the returned edit, and asserting the
buffer changed as expected, in a real nvim session
[x] Full suite (`make test`) and e2e harness (`make test-e2e`) green

## Summary of Changes

Both stories completed as scoped: edifact-ls-338y (`Error.Code`/`Fix`
data model, populated at the redundant-UNA and six envelope
count/reference-mismatch call sites) and edifact-ls-n3yd
(`textDocument/codeAction` handler consuming it). See each story's own
Summary of Changes for the file-level detail. Net result: opening a
document with a redundant UNA or a wrong UNT/UNE/UNZ count/reference
now offers a quickfix code action in the editor that, when applied,
genuinely resolves the diagnostic -- verified both by unit tests that
apply the fix and re-validate, and by two new e2e checks that do the
same in a real nvim session via `vim.lsp.util.apply_workspace_edit`.

## Retro

- Splitting the data-model story from the handler story paid off
  exactly as planned: story 1 was pure, low-risk plumbing (no
  LSP-facing behavior change, verified by the full suite staying green
  with zero e2e deltas), so story 2 could focus entirely on the
  handler and range-filtering logic without also re-deriving which
  diagnostics are fixable.
- The design discussion done well before implementation (mapping every
  diagnostic source against "is the fix mechanically derivable")
  turned out to be accurate on first implementation -- no diagnostic
  kind was discovered mid-story to need reclassifying either way. That
  discussion, not skipped or compressed, is why the scope held.
- One real subtlety only became visible while writing the handler
  test, not the design doc: the diagnostic's displayed range
  (`errorRange(text, e.Pos)`) and the Fix's own replacement span
  (`Fix.Pos`/`OldText`) are not the same span for the envelope-mismatch
  cases -- e.Pos is the segment's position (e.g. UNT's), while Fix.Pos
  is the specific wrong component's position. Code-action range
  filtering correctly uses the former (matching what the client
  actually sees highlighted), while the edit itself correctly uses the
  latter (the precise text to replace). Worth remembering for any
  future diagnostic-with-fix: these two positions can differ, and
  conflating them would either miss legitimate requests or replace the
  wrong span.
- Reusing `toDiagnostic` (factored out of `diagnostics.go`) so the
  code action's `Diagnostics` field is built through the exact same
  path as what's already published, rather than a second hand-rolled
  translation, is a small thing but avoids a class of "the code action
  references a diagnostic that doesn't quite match what the client has"
  bug before it could exist.
- No adjustments needed before starting further work in this area. A
  natural, not-yet-requested follow-up would be surfacing `Code` as a
  stable id clients can build keybindings/quickfix-all-in-file
  behavior around, but that's speculative until asked for.
