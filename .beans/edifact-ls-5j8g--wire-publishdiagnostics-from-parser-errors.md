---
# edifact-ls-5j8g
title: Wire publishDiagnostics from parser errors
status: completed
type: feature
priority: normal
created_at: 2026-09-01T13:03:16Z
updated_at: 2026-09-01T13:40:05Z
parent: edifact-ls-z9te
---

# Description
On document open/change, run the parser core over the document text and
translate its structured errors into LSP `Diagnostic` objects, published via
`textDocument/publishDiagnostics`.

# Acceptance Criteria
- [x] Diagnostics published on open and on every change (no debounce needed:
      files this small make re-parsing on every keystroke a non-issue)
- [x] Byte/line/column positions from the parser correctly map to LSP
      `Range` (UTF-16 code unit columns per spec)
- [x] Previously-published diagnostics are cleared/replaced (not accumulated)
      on each new publish
- [x] Unit test driving the server with a known-bad document and asserting
      the exact diagnostics published

## Summary of Changes
`internal/lspserver/diagnostics.go`: wires `textDocument/didOpen`,
`didChange` (full-sync), and `didClose` to run `edifact.Parse` +
`edifact.ValidateEnvelopes` over the current document text and publish the
result via `textDocument/publishDiagnostics`. Positions are converted from
the parser's byte offsets to LSP's 0-based-line/UTF-16-character `Range` via
`offsetToLSPPosition` (widens to a 1-character-wide range where possible so
clients render a visible squiggle rather than a zero-width marker).
`initialize` now advertises `TextDocumentSyncKindFull`. Covered by
`diagnostics_test.go`: pure unit tests on the offset/diagnostic translation,
plus a full stdio-transport test asserting the actual published notification
on open and that it's replaced (not appended to) after an edit.

Found via testing: my first two test fixtures were bare fragments (no
UNB/UNZ), so envelope validation correctly added *extra* diagnostics on top
of the one syntax error each test meant to isolate -- not a bug in the
implementation, a bug in the fixtures. Fixed by using complete, otherwise-
valid envelopes with just the one targeted error.
