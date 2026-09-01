---
# edifact-ls-efgg
title: Round-trip tests + e2e formatting check
status: completed
type: feature
priority: normal
created_at: 2026-09-01T13:03:29Z
updated_at: 2026-09-01T13:54:51Z
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
- [x] Round-trip test: for each fixture, `parse(format(x))` yields the same
      logical AST (ignoring position/whitespace) as `parse(x)`
- [x] e2e harness extended with a check that triggers `:lua
      vim.lsp.buf.format()` (or equivalent) on a sample fixture in headless
      nvim and asserts the resulting buffer content
- [x] Passes as part of the single e2e entrypoint command

## Summary of Changes
`internal/edifact/render_test.go`: `TestRenderRoundTripsAllCleanFixtures`
sweeps every fixture under `testdata/` that parses without syntax errors and
asserts `parse(Render(parse(x)))` carries the same data as `parse(x)`, in
both multiline and wire render modes (fixtures with deliberate syntax
errors are skipped by design -- error-recovery discards the malformed
segment's content, so re-rendering it is expected to lose data, which is
exactly why `textDocument/formatting` itself refuses to touch unparseable
documents).

Added `testdata/unformatted.edi` (single-line wire-style) and a
`check_formatting` e2e check that opens it, calls `vim.lsp.buf.format()`,
and asserts the resulting buffer matches the expected multiline layout.

Found via testing: the harness's final `qa` failed with E37 ("no write
since last change") once a check started leaving buffers modified without
saving -- hung nvim indefinitely since a headless session can't answer the
interactive retry prompt. Changed to `qa!`. A second hang on re-run came
from a leftover swapfile from that killed session blocking the next run
behind an unanswerable "ATTENTION" prompt; fixed by passing `--cmd "set
noswapfile"` in scripts/e2e.sh (checks intentionally leave buffers modified
without saving, so a persisted swapfile serves no purpose here).
