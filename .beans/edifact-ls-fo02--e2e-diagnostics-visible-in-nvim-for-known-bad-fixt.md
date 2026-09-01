---
# edifact-ls-fo02
title: 'e2e: diagnostics visible in nvim for known-bad fixtures'
status: todo
type: feature
priority: normal
created_at: 2026-09-01T13:03:16Z
updated_at: 2026-09-01T13:03:16Z
parent: edifact-ls-z9te
blocked_by:
    - edifact-ls-5j8g
    - edifact-ls-vwk2
---

# Description
Add known-bad EDIFACT fixtures (syntax error, semantic/envelope error) and
extend the headless e2e harness to open them and assert the expected
diagnostics appear via `vim.diagnostic.get()`.

# Acceptance Criteria
- [ ] At least two new fixtures under `testdata/`: one with a syntax error,
      one with a semantic/envelope error
- [ ] e2e harness extended with a check that opens each and asserts the
      expected diagnostic message/range is present
- [ ] Passes as part of the single e2e entrypoint command
