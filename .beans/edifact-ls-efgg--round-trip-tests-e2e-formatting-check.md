---
# edifact-ls-efgg
title: Round-trip tests + e2e formatting check
status: todo
type: feature
priority: normal
created_at: 2026-09-01T13:03:29Z
updated_at: 2026-09-01T13:03:29Z
parent: edifact-ls-s200
blocked_by:
    - edifact-ls-ct07
    - edifact-ls-vwk2
---

# Description
Property-style round-trip tests proving formatting never changes parsed
semantic content, plus an e2e check that formatting works through the actual
LSP/nvim path.

# Acceptance Criteria
- [ ] Round-trip test: for each fixture, `parse(format(x))` yields the same
      logical AST (ignoring position/whitespace) as `parse(x)`
- [ ] e2e harness extended with a check that triggers `:lua
      vim.lsp.buf.format()` (or equivalent) on a sample fixture in headless
      nvim and asserts the resulting buffer content
- [ ] Passes as part of the single e2e entrypoint command
