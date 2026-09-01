---
# edifact-ls-dyrz
title: Standalone nvim-lspconfig server definition + docs
status: completed
type: feature
priority: low
created_at: 2026-09-01T13:03:42Z
updated_at: 2026-09-01T14:30:38Z
parent: edifact-ls-hlup
---

# Description
LSP half of a real local install: a way to get the `edifact-ls` binary onto
your `$PATH` (not just `dist/edifact-ls`), plus a server definition suitable
for `nvim-lspconfig`'s config table (or a PR to lspconfig itself),
independent of the in-repo dev harness, plus install/config documentation
for end users. Syntax highlighting's local-install story is separate
(edifact-ls-95ap) -- this one is LSP-only.

# Acceptance Criteria
- [x] A documented local-install path for the binary onto `$PATH`/`$GOBIN`
      (e.g. `go install ./cmd/edifact-ls`, or a `make install` wrapper)
- [x] Server definition (cmd, filetypes, root_dir detection) documented in
      README and tested against that installed binary
- [x] Filetype detection documented (`.edi`/`.edifact`)
- [x] Manual verification: a fresh nvim config (not this repo's dev harness)
      following only the README can attach to edifact-ls on a sample file

## Summary of Changes
`Makefile`: `make install` (wraps `go install ./cmd/edifact-ls`).
`editors/nvim/standalone_lsp.lua`: a real, self-contained config snippet for
a user's own Neovim config -- filetype detection, LSP registration via
`vim.fn.exepath("edifact-ls")` (installed-on-PATH lookup, not the dev
harness's env var), a `vim.fs.root(".git")`-based root_dir with a
same-directory fallback, and the `:EdifactMinify` command. README gained an
"Installing for your own Neovim config" section, including a documented
gotcha: under this repo's Hermit-activated shell, `go install`'s `$GOBIN`
points at `.hermit/go/bin/`, only on `$PATH` while that shell is active --
not useful for a persistent config, so the README calls out installing with
a non-Hermit Go toolchain or an explicit `GOBIN` instead.

Manually verified end-to-end: built+installed the binary, launched nvim with
`-u` pointed at a config that does nothing but `dofile` the standalone
snippet (no dev-harness env vars at all) against `testdata/minimal.edi`,
confirmed the LSP client attaches and `:EdifactMinify` actually collapses
the buffer.
