---
# edifact-ls-1pz0
title: Minimal nvim plugin config via lspconfig
status: completed
type: feature
priority: high
created_at: 2026-09-01T13:02:21Z
updated_at: 2026-09-01T13:11:33Z
parent: edifact-ls-untk
blocked_by:
    - edifact-ls-upnf
---

# Description
A minimal, in-repo nvim configuration that registers edifact-ls as an LSP
server for EDIFACT files, pointed at a local dev build. This is what the
e2e harness (and any human) uses to try the server in a real editor.

# Acceptance Criteria
- [x] In-repo nvim config (`editors/nvim/init.lua`) launches edifact-ls for
      `*.edi`/`*.edifact` filetypes via built-in `vim.lsp.config`/`vim.lsp.enable`
      (nvim >=0.11) rather than the nvim-lspconfig plugin, to keep the dev/test
      harness dependency-free; nvim-lspconfig itself now builds on this same
      API. A real nvim-lspconfig-based setup for end users is Epic
      edifact-ls-hlup's job.
- [x] Filetype detection for `.edi`/`.edifact` added (nvim autocmd or ftdetect)
- [x] Config is parameterized to point at a locally built binary (path env var
      or similar), not a system-installed one
- [x] README section explaining how to launch nvim with this config manually
      for interactive testing

## Summary of Changes
Added `editors/nvim/init.lua`: registers `.edi`/`.edifact` filetype
detection and an `edifact_ls` server config via `vim.lsp.config`/`vim.lsp.enable`,
pointed at a binary path from `$EDIFACT_LS_BIN`. Documented manual usage in
README.md.
