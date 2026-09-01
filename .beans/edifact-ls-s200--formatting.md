---
# edifact-ls-s200
title: Formatting
status: todo
type: epic
priority: normal
created_at: 2026-09-01T13:03:29Z
updated_at: 2026-09-01T13:03:29Z
parent: edifact-ls-gdt6
blocked_by:
    - edifact-ls-btas
---

# Goal
LSP document formatting for EDIFACT files — a stable, idempotent pretty-
printer built on top of the parser core's AST.

# Acceptance Criteria
- [ ] `textDocument/formatting` implemented, producing a consistent,
      readable layout (e.g. one segment per line) from the parsed AST
- [ ] Idempotent: formatting already-formatted output produces no changes
- [ ] Semantically lossless: `parse(format(x))` carries the same data as
      `parse(x)` (formatting never changes segment/element content)
- [ ] Verified end-to-end via the e2e harness: formatting a sample fixture
      in headless nvim produces the expected output
