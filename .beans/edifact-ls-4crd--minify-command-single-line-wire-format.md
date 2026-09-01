---
# edifact-ls-4crd
title: Minify command (single-line wire format)
status: todo
type: feature
priority: normal
created_at: 2026-09-01T13:36:02Z
updated_at: 2026-09-01T13:36:02Z
parent: edifact-ls-s200
blocked_by:
    - edifact-ls-ct07
---

# Description
Expose a custom LSP command, `edifact-ls.minify`, that collapses the current
document to true "wire" EDIFACT: one continuous line, segments joined by
their terminator only, with no inserted newlines — how an interchange
actually looks in transit. This is the one direction not already covered:
the multiline, human-readable direction is `textDocument/formatting`
(edifact-ls-ct07) itself, so there's no separate "unpack"/"expand" command —
`minify` should reuse that same renderer's parse step and AST, just with a
different (single-line) render mode, rather than duplicating logic.

# Acceptance Criteria
- [ ] Server advertises `ExecuteCommandProvider` with the `edifact-ls.minify`
      command ID, handled via `workspace/executeCommand`
- [ ] `edifact-ls.minify` collapses the document to single-line wire format:
      segments joined by their terminator only, no inserted whitespace/
      newlines between segments
- [ ] Semantically lossless: `parse(minify(x))` carries the same data as
      `parse(x)` (same guarantee as `textDocument/formatting`)
- [ ] Applied to the document via a workspace edit (consistent with how
      `textDocument/formatting` applies its edit)
- [ ] nvim side: a user-facing command (e.g. `:EdifactMinify`, wired to
      `vim.lsp.buf.execute_command`) documented in README; e2e harness gets
      a check exercising it on the sample fixture
- [ ] Unit tests: minify a multiline fixture -> single-line wire output;
      round-trip `format(minify(x))` is semantically equivalent to `x` (and,
      given `textDocument/formatting`'s own idempotency, equal to `format(x)`)
