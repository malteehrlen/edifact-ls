---
# edifact-ls-95ap
title: Standalone tree-sitter parser install for personal nvim config
status: completed
type: feature
priority: low
created_at: 2026-09-01T14:27:34Z
updated_at: 2026-09-01T14:30:53Z
parent: edifact-ls-hlup
---

# Description
Highlighting half of a real local install: a way to build and install the
`tree-sitter-edifact` grammar outside this repo's dev harness, plus a
copy-pasteable snippet for a user's own `init.lua` -- independent of, and
in parallel with, the LSP-only install story (edifact-ls-dyrz). Uses the
same dependency-free approach as the dev harness (`vim.treesitter.language.add`
+ `vim.treesitter.query.set`) rather than requiring the `nvim-treesitter`
plugin specifically, so it's portable regardless of plugin manager.

# Acceptance Criteria
- [x] Documented steps to build `edifact.so` locally (`npm install` +
      `tree-sitter build`) and install it to a stable location outside a
      throwaway repo checkout
- [x] Copy-pasteable nvim snippet (`vim.treesitter.language.add` +
      `vim.treesitter.query.set`) for `.edi`/`.edifact` filetypes,
      independent of the in-repo dev harness and without requiring the
      `nvim-treesitter` plugin
- [x] README covers this alongside the LSP install path (edifact-ls-dyrz)
- [x] Manual verification: a fresh nvim config (not this repo's dev
      harness) following only the README shows highlighting on a sample file

## Summary of Changes
`editors/nvim/standalone_treesitter.lua`: a real, self-contained config
snippet reading `edifact.so`/`highlights.scm` from a configurable stable
directory (`~/.local/share/edifact-ls` by default), registering the parser
via `vim.treesitter.language.add`/`query.set` and starting highlighting on
`FileType edifact` -- no `nvim-treesitter` plugin dependency, matching the
same approach as the dev harness and the LSP standalone story. README's new
"Installing for your own Neovim config" section (added alongside
edifact-ls-dyrz) covers building and copying the parser files in.

Manually verified end-to-end: built the parser, copied it to
`~/.local/share/edifact-ls/`, launched nvim with a config that does nothing
but `dofile` both standalone snippets (LSP + tree-sitter) against
`testdata/minimal.edi` -- confirmed the tree-sitter parser is active with no
ERROR nodes, alongside the LSP client attaching in the same session.
