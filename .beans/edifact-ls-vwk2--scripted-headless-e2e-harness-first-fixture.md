---
# edifact-ls-vwk2
title: Scripted headless e2e harness + first fixture
status: completed
type: feature
priority: high
created_at: 2026-09-01T13:02:21Z
updated_at: 2026-09-01T13:11:42Z
parent: edifact-ls-untk
blocked_by:
    - edifact-ls-upnf
    - edifact-ls-1pz0
---

# Description
A non-interactive, scriptable e2e test that proves the whole chain works:
build server -> launch nvim headless -> open a sample EDIFACT file -> LSP
client attaches. This becomes the harness every later epic (highlighting,
diagnostics, formatting) plugs its own e2e checks into.

# Acceptance Criteria
- [x] At least one minimal valid EDIFACT sample fixture checked in (e.g.
      `testdata/minimal.edi`)
- [x] Headless nvim script (Lua, run via `nvim --headless`) opens the fixture
      using the config from the previous story and asserts the LSP client
      attached (e.g. via `vim.lsp.get_clients()`), exiting non-zero on failure
- [x] Single entrypoint command (e.g. `make test-e2e` or `scripts/e2e.sh`)
      builds the server and runs the headless check
- [x] Documented in README: how to run it, how to extend it for new checks
- [x] Runs successfully in this environment (or CI) with no manual steps

## Summary of Changes
Added `testdata/minimal.edi`, `scripts/e2e_check.lua` (headless checks,
extensible via `check_*` functions + `fail`/`pass` helpers), and
`scripts/e2e.sh` (locates or downloads a pinned Neovim into `.tools/` if
none is on PATH, then runs the checks against `$EDIFACT_LS_BIN`). Wired up
as `make test-e2e`. Verified both the pass path and the fail path (pointing
`EDIFACT_LS_BIN` at a non-LSP binary correctly exits 1 with a clear
message).
