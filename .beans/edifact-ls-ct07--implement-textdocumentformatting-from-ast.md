---
# edifact-ls-ct07
title: Implement textDocument/formatting from AST
status: todo
type: feature
priority: normal
created_at: 2026-09-01T13:03:29Z
updated_at: 2026-09-01T13:03:29Z
parent: edifact-ls-s200
---

# Description
Implement the `textDocument/formatting` LSP request: parse the document,
render the AST back out with a consistent pretty-printing style (segments
each on their own line, consistent delimiter spacing/casing), and return the
edit as LSP `TextEdit`s.

# Acceptance Criteria
- [ ] Handles `textDocument/formatting` and returns `TextEdit[]` (whole-
      document replace is acceptable initially)
- [ ] Idempotent on its own output (unit test: format twice, same result)
- [ ] Returns no edits (or a clear no-op) when the document fails to parse,
      rather than corrupting the file
- [ ] Unit tests covering at least: a minimal interchange, one with composite
      elements, one already correctly formatted (no-op)
