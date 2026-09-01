---
# edifact-ls-z9te
title: LSP diagnostics & linting
status: todo
type: epic
priority: normal
created_at: 2026-09-01T13:03:16Z
updated_at: 2026-09-01T13:03:16Z
parent: edifact-ls-gdt6
blocked_by:
    - edifact-ls-btas
---

# Goal
Surface the parser core's syntax and semantic errors as LSP diagnostics, so
opening a malformed EDIFACT file in nvim shows inline errors.

# Acceptance Criteria
- [ ] Opening/editing a document triggers `textDocument/publishDiagnostics`
      with correct ranges derived from the parser's positioned errors
- [ ] Diagnostics clear/update correctly as the document is edited (no stale
      diagnostics left behind)
- [ ] Covers both syntax errors (lexer/parser) and semantic errors (envelope
      validation) from the parser core epic
- [ ] Verified end-to-end via the e2e harness: opening a known-bad fixture in
      headless nvim shows the expected diagnostic(s)
