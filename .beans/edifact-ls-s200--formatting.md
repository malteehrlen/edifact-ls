---
# edifact-ls-s200
title: Formatting
status: completed
type: epic
priority: normal
created_at: 2026-09-01T13:03:29Z
updated_at: 2026-09-01T14:01:40Z
parent: edifact-ls-gdt6
blocked_by:
    - edifact-ls-btas
---

# Goal
LSP document formatting for EDIFACT files — a stable, idempotent pretty-
printer built on top of the parser core's AST.

# Acceptance Criteria
- [x] `textDocument/formatting` implemented, producing a consistent,
      readable layout (e.g. one segment per line) from the parsed AST
- [x] Idempotent: formatting already-formatted output produces no changes
- [x] Semantically lossless: `parse(format(x))` carries the same data as
      `parse(x)` (formatting never changes segment/element content)
- [x] Verified end-to-end via the e2e harness: formatting a sample fixture
      in headless nvim produces the expected output

## Summary of Changes
`internal/edifact/render.go`: one `Render(ic, multiline)` function backs both
`textDocument/formatting` (multiline) and the `edifact-ls.minify` command
(single-line wire format), reproducing each component's original escaped
text verbatim so losslessness follows from the implementation rather than
needing separate proof. `internal/lspserver/formatting.go` and
`commands.go` wire both into the server; `initialize` now derives
capabilities via `handler.CreateServerCapabilities()` (extended with the two
gaps that helper leaves: `TextDocumentSyncKind` defaults to Incremental, and
`ExecuteCommandOptions.Commands` is left empty) instead of a hand-maintained
list. All three stories' e2e checks (formatting, minify) run through the
real headless-nvim harness alongside the existing diagnostics ones.

## Retro
- Same test-fixture discipline as the diagnostics epic paid off again: the
  fixture-sweeping round-trip test caught nothing new, but having it
  meant I wasn't relying on hand-picked cases to prove losslessness.
- One real bug this time, not a test mistake: a server-initiated
  `workspace/applyEdit` call made synchronously from inside the
  `workspace/executeCommand` handler deadlocked the connection, because this
  jsonrpc2 version dispatches incoming requests inline on its one read-loop
  goroutine rather than per-request goroutines. Fixed by firing the nested
  call from a goroutine. This is a general constraint on this library, not
  specific to minify -- worth remembering for any future server-initiated
  request (client/registerCapability, window/showMessageRequest, etc.) made
  from within a request handler, not just this one command.
- Also hit (and fixed) two e2e-harness robustness issues along the way: the
  harness's final `qa` needed to become `qa!` once checks started leaving
  buffers modified without saving, and a leftover swapfile from a killed run
  needed `--cmd "set noswapfile"` to stop blocking future runs behind an
  unanswerable prompt. Both are now handled for any future check that edits
  a buffer.
- No adjustments needed before continuing. Tree-sitter grammar & syntax
  highlighting is the last open, unblocked epic (packaging stays
  deliberately deferred); ready whenever you'd like to green-light it.
