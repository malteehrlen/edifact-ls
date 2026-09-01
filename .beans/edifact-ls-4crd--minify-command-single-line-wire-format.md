---
# edifact-ls-4crd
title: Minify command (single-line wire format)
status: completed
type: feature
priority: normal
created_at: 2026-09-01T13:36:02Z
updated_at: 2026-09-01T14:01:21Z
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
- [x] Server advertises `ExecuteCommandProvider` with the `edifact-ls.minify`
      command ID, handled via `workspace/executeCommand`
- [x] `edifact-ls.minify` collapses the document to single-line wire format:
      segments joined by their terminator only, no inserted whitespace/
      newlines between segments
- [x] Semantically lossless: `parse(minify(x))` carries the same data as
      `parse(x)` (same guarantee as `textDocument/formatting`)
- [x] Applied to the document via a workspace edit (consistent with how
      `textDocument/formatting` applies its edit)
- [x] nvim side: a user-facing command (`:EdifactMinify`, wired to
      `client:exec_cmd` -- `vim.lsp.buf.execute_command` is deprecated as of
      nvim 0.12) documented in README; e2e harness gets a check exercising
      it on the sample fixture
- [x] Unit tests: minify a multiline fixture -> single-line wire output;
      round-trip `format(minify(x))` is semantically equivalent to `x` (and,
      given `textDocument/formatting`'s own idempotency, equal to `format(x)`)

## Summary of Changes
`internal/lspserver/commands.go`: wires `workspace/executeCommand`,
dispatching `edifact-ls.minify` (expects a document-URI argument) to a
handler that parses the document, renders it via `edifact.Render(ic,
false)` (the wire-format mode added alongside the multiline one in
edifact-ls-ct07), and applies the result via a server-initiated
`workspace/applyEdit` request. `ExecuteCommandProvider.Commands` is filled
in manually in `initialize`, since glsp's `CreateServerCapabilities` sets
the capability but leaves `Commands` empty. nvim side: `:EdifactMinify` user
command in `editors/nvim/init.lua`, using `client:exec_cmd` (the
non-deprecated replacement for `vim.lsp.buf.execute_command` as of nvim
0.12 -- confirmed via the runtime's own deprecation notice during e2e
testing). Documented in README.

Found via testing (real bug, not a fixture mistake this time): calling
`workspace/applyEdit` synchronously from inside the `workspace/executeCommand`
handler deadlocked the connection. This jsonrpc2 version dispatches incoming
requests inline on its single read-loop goroutine rather than spawning one
per request, so a handler blocking on its own nested outbound call blocks
the very loop that would read that call's response. Fixed by firing the
`applyEdit` call from a separate goroutine and only logging (not returning)
an error if the client doesn't apply it -- executeCommand's own response no
longer waits on the edit landing. Worth remembering for any future feature
that needs a server-initiated request (not just a notification) from within
a request handler.
