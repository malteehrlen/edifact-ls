---
# edifact-ls-ct07
title: Implement textDocument/formatting from AST
status: completed
type: feature
priority: normal
created_at: 2026-09-01T13:03:29Z
updated_at: 2026-09-01T13:49:31Z
parent: edifact-ls-s200
---

# Description
Implement the `textDocument/formatting` LSP request: parse the document,
render the AST back out with a consistent pretty-printing style (segments
each on their own line, consistent delimiter spacing/casing), and return the
edit as LSP `TextEdit`s.

# Acceptance Criteria
- [x] Handles `textDocument/formatting` and returns `TextEdit[]` (whole-
      document replace is acceptable initially)
- [x] Idempotent on its own output (unit test: format twice, same result)
- [x] Returns no edits (or a clear no-op) when the document fails to parse,
      rather than corrupting the file
- [x] Unit tests covering at least: a minimal interchange, one with composite
      elements, one already correctly formatted (no-op)

## Summary of Changes
`internal/edifact/render.go`: `Render(ic, multiline)` re-serializes an
Interchange either as one-segment-per-line (formatting) or true single-line
wire format (shared with the upcoming minify story), reproducing each
component's original `Raw` text verbatim so it's a pure re-layout, never a
re-encoding -- losslessness follows from that by construction, not from
separately re-implemented escaping logic.
`internal/lspserver/formatting.go`: wires `textDocument/formatting`,
returning a single whole-document `TextEdit` (or no edits at all if the
document doesn't parse cleanly, or is already formatted). Also refactored
`initialize` to build capabilities via `handler.CreateServerCapabilities()`
instead of a hand-maintained list, so newly wired handler funcs (like this
one) can't be forgotten from the advertised capabilities -- overriding only
`TextDocumentSyncKind` to Full, since we track whole-document text rather
than incremental ranges.
