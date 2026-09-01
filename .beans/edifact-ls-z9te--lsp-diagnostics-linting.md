---
# edifact-ls-z9te
title: LSP diagnostics & linting
status: completed
type: epic
priority: normal
created_at: 2026-09-01T13:03:16Z
updated_at: 2026-09-01T13:41:17Z
parent: edifact-ls-gdt6
blocked_by:
    - edifact-ls-btas
---

# Goal
Surface the parser core's syntax and semantic errors as LSP diagnostics, so
opening a malformed EDIFACT file in nvim shows inline errors.

# Acceptance Criteria
- [x] Opening/editing a document triggers `textDocument/publishDiagnostics`
      with correct ranges derived from the parser's positioned errors
- [x] Diagnostics clear/update correctly as the document is edited (no stale
      diagnostics left behind)
- [x] Covers both syntax errors (lexer/parser) and semantic errors (envelope
      validation) from the parser core epic
- [x] Verified end-to-end via the e2e harness: opening a known-bad fixture in
      headless nvim shows the expected diagnostic(s)

## Summary of Changes
Wired the parser core's errors into real LSP diagnostics: `initialize` now
advertises full text-document sync, `didOpen`/`didChange`/`didClose` track
each document's text and publish freshly computed diagnostics (parse +
envelope validation combined) via `textDocument/publishDiagnostics`, with
byte offsets converted to spec-correct UTF-16-based LSP ranges. Verified
both at the unit level (stdio-transport test asserting the actual published
notification, including replace-not-accumulate across an edit) and via two
new e2e fixtures exercised through the real headless-nvim harness.

## Retro
- Straightforward epic -- parser core epic already did the hard part
  (structured, positioned errors); this was mostly plumbing plus getting the
  UTF-16 range conversion right.
- Test-fixture mistake worth repeating from the last retro: two early unit
  tests used bare fragments (missing UNB/UNZ) as "syntax error" fixtures, so
  envelope validation correctly piled on extra diagnostics the tests weren't
  expecting. Fixed by always using complete, otherwise-valid envelopes when
  a test wants to isolate one specific error. Might be worth a shared test
  helper (`validEnvelopeWith(segment)`) if this pattern keeps recurring in
  the formatting epic's tests too.
- Also manually verified the e2e harness's failure path is trustworthy here
  (temporarily broke a fixture, confirmed FAIL + exit 1, restored it) -- same
  discipline as the scaffolding epic.
- No adjustments needed before continuing. Formatting and tree-sitter
  highlighting are both still open and unblocked whenever you'd like to
  green-light one.
