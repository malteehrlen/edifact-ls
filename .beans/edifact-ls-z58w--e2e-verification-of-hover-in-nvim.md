---
# edifact-ls-z58w
title: e2e verification of hover in nvim
status: completed
type: feature
priority: normal
created_at: 2026-09-01T17:53:47Z
updated_at: 2026-09-01T19:45:03Z
parent: edifact-ls-tnp9
blocked_by:
    - edifact-ls-fo5r
---

# Description

Add a headless e2e check (scripts/e2e_check.lua, same pattern as the
existing check_* functions) that opens a fixture, requests hover at a
known segment tag's position via `vim.lsp.buf.hover()` (or a direct
`textDocument/hover` request), and asserts the response contains the
expected description text.

# Acceptance Criteria

[x] New check function added to scripts/e2e_check.lua following the
existing fail()/pass() convention
[x] Verifies hover content for at least one service segment and one
business segment
[x] Wired into the check sequence at the bottom of the script
[x] Full e2e suite (`make test-e2e`) passes

## Summary of Changes

scripts/e2e_check.lua: check_hover(fixture, line, character,
expect_substring) makes a direct, synchronous `textDocument/hover`
request via `vim.lsp.buf_request_sync` with explicit position
parameters, rather than `vim.lsp.buf.hover()` -- that API drives an
async floating window off the *current cursor position*, which is a
worse fit for scripted assertions than a direct request with an
explicit position. Wired in against testdata/minimal.edi: UNB (service
segment, line 0) and BGM (business segment, line 2), matching the
epic's "at least one of each" bar.

Confirmed for real, not just plausible: `make test-e2e` runs this
against the actual built binary in a real headless nvim + LSP session,
and both checks pass --

    PASS: hover at .../testdata/minimal.edi:0:1 includes a message containing "Interchange header"
    PASS: hover at .../testdata/minimal.edi:2:1 includes a message containing "Beginning of message"

Full e2e suite (now 12 checks) green.
