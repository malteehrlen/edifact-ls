---
# edifact-ls-87s0
title: Wire grammar into nvim + verify via e2e harness
status: todo
type: feature
priority: normal
created_at: 2026-09-01T13:03:03Z
updated_at: 2026-09-01T13:03:03Z
parent: edifact-ls-nr8g
blocked_by:
    - edifact-ls-ujaa
    - edifact-ls-vwk2
---

# Description
Register the EDIFACT tree-sitter parser with nvim's tree-sitter integration
(e.g. `nvim-treesitter` parser config pointing at this repo, or a manual
`vim.treesitter.language.add`) inside the in-repo minimal nvim config, and
extend the e2e harness to assert highlighting actually activates.

# Acceptance Criteria
- [ ] Minimal nvim config (from the scaffolding epic) registers the EDIFACT
      tree-sitter parser and enables highlighting for `.edi`/`.edifact` files
- [ ] Headless e2e check extended to assert the buffer's tree-sitter parser
      is active and produces no `ERROR` nodes on the sample fixture
- [ ] README updated with how to install/build the parser locally
