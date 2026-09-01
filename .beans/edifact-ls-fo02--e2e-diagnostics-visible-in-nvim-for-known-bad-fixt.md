---
# edifact-ls-fo02
title: 'e2e: diagnostics visible in nvim for known-bad fixtures'
status: completed
type: feature
priority: normal
created_at: 2026-09-01T13:03:16Z
updated_at: 2026-09-01T13:40:59Z
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
- [x] At least two new fixtures under `testdata/`: one with a syntax error,
      one with a semantic/envelope error
- [x] e2e harness extended with a check that opens each and asserts the
      expected diagnostic message/range is present
- [x] Passes as part of the single e2e entrypoint command

## Summary of Changes
Added `testdata/syntax-error.edi` (invalid segment tag) and
`testdata/envelope-error.edi` (missing UNZ). Extended
`scripts/e2e_check.lua` with `check_diagnostic(fixture, expect_substring)`,
which opens the fixture, waits (via `vim.wait`) for `vim.diagnostic.get(0)`
to be non-empty, and asserts one of the messages contains the expected
substring. Verified both the pass path and, by temporarily replacing a
fixture with valid content, that the harness fails non-zero with a clear
`FAIL:` reason when an expected diagnostic doesn't show up.
