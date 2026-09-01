---
# edifact-ls-dyrz
title: Standalone nvim-lspconfig server definition + docs
status: todo
type: feature
priority: low
created_at: 2026-09-01T13:03:42Z
updated_at: 2026-09-01T13:03:42Z
parent: edifact-ls-hlup
---

# Description
A server definition suitable for `nvim-lspconfig`'s config table (or a PR to
lspconfig itself), independent of the in-repo dev harness, plus install/
config documentation for end users.

# Acceptance Criteria
- [ ] Server definition (cmd, filetypes, root_dir detection) documented in
      README and tested against a release/locally-built binary
- [ ] Filetype detection documented (`.edi`/`.edifact`)
- [ ] Manual verification: a fresh nvim config following only the README can
      attach to edifact-ls on a sample file
