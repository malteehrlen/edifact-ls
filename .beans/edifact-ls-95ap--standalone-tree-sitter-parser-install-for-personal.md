---
# edifact-ls-95ap
title: Standalone tree-sitter parser install for personal nvim config
status: todo
type: feature
priority: low
created_at: 2026-09-01T14:27:34Z
updated_at: 2026-09-01T14:27:39Z
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
- [ ] Documented steps to build `edifact.so` locally (`npm install` +
      `tree-sitter build`) and install it to a stable location outside a
      throwaway repo checkout
- [ ] Copy-pasteable nvim snippet (`vim.treesitter.language.add` +
      `vim.treesitter.query.set`) for `.edi`/`.edifact` filetypes,
      independent of the in-repo dev harness and without requiring the
      `nvim-treesitter` plugin
- [ ] README covers this alongside the LSP install path (edifact-ls-dyrz)
- [ ] Manual verification: a fresh nvim config (not this repo's dev
      harness) following only the README shows highlighting on a sample file
