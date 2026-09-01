---
# edifact-ls-untk
title: Project scaffolding & E2E test harness
status: completed
type: epic
priority: high
created_at: 2026-09-01T13:02:07Z
updated_at: 2026-09-01T13:12:06Z
parent: edifact-ls-gdt6
---

# Goal
Stand up the minimal skeleton of edifact-ls end to end — a Go LSP server and a
minimal nvim plugin config wired via lspconfig — so that every subsequent
epic has a working, testable harness to build against instead of testing
manually in an editor.

# Acceptance Criteria
- [x] `go.mod`/project layout exists for the LSP server (stdio transport)
- [x] Server implements `initialize`/`initialized`/`shutdown`/`exit` and
      negotiates capabilities (empty/no-op capabilities are fine at this stage)
- [x] A minimal nvim config (`editors/nvim/init.lua`) registers edifact-ls via
      built-in `vim.lsp.config`/`vim.lsp.enable` (see edifact-ls-1pz0 for why,
      not the nvim-lspconfig plugin) for `.edi`/`.edifact` files, pointing at
      a local dev build
- [x] A scripted, non-interactive e2e test exists that: builds the server,
      launches nvim headlessly with the minimal config, opens a sample EDIFACT
      fixture file, and asserts the LSP client attaches successfully
- [x] The e2e harness is runnable via a single command (e.g. `make test-e2e`)
      and documented so later epics can extend it
- [x] CI-friendly: harness exits non-zero on failure, no manual steps required

## Summary of Changes
Go LSP server (`internal/lspserver` + `cmd/edifact-ls`, on `github.com/tliron/glsp`)
implements the initialize/initialized/shutdown/exit lifecycle with a
handshake unit test over a real stdio-shaped in-memory transport. A minimal
in-repo nvim config (`editors/nvim/init.lua`) attaches it to `.edi`/`.edifact`
files via built-in `vim.lsp.config`. `scripts/e2e.sh` + `scripts/e2e_check.lua`
give a single-command (`make test-e2e`), CI-friendly, no-manual-steps headless
check that the two attach correctly — verified both the pass and fail paths.

## Retro
- Went smoothly once the environment was sorted (hermit/go/nvim install,
  network policy). No blockers in the actual implementation.
- Deviation from the milestone's literal wording: used Neovim's built-in
  `vim.lsp.config`/`vim.lsp.enable` instead of the `nvim-lspconfig` plugin for
  this dev/test harness, to avoid a plugin-manager dependency in e2e runs.
  nvim-lspconfig 3.x is itself built on this same API, so this isn't a
  meaningfully different mechanism — but the *actual* nvim-lspconfig plugin
  registration for end users still needs to happen in the packaging epic
  (edifact-ls-hlup). Flagging this explicitly in case that substitution isn't
  acceptable for the harness itself.
- The e2e harness only asserts client attachment so far (by design — that's
  all Epic 1 needed). It's structured (`check_*` functions in
  `scripts/e2e_check.lua`) so diagnostics/formatting/highlighting epics can
  each add their own check without restructuring anything.
- No adjustments needed before continuing to the parser core epic.
