---
# edifact-ls-dyrz
title: Standalone nvim-lspconfig server definition + docs
status: todo
type: feature
priority: low
created_at: 2026-09-01T13:03:42Z
updated_at: 2026-09-01T14:27:39Z
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
- [ ] A documented local-install path for the binary onto `$PATH`/`$GOBIN`
      (e.g. `go install ./cmd/edifact-ls`, or a `make install` wrapper)
- [ ] Server definition (cmd, filetypes, root_dir detection) documented in
      README and tested against that installed binary
- [ ] Filetype detection documented (`.edi`/`.edifact`)
- [ ] Manual verification: a fresh nvim config (not this repo's dev harness)
      following only the README can attach to edifact-ls on a sample file
