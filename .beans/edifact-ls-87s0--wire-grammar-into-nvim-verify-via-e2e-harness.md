---
# edifact-ls-87s0
title: Wire grammar into nvim + verify via e2e harness
status: completed
type: feature
priority: normal
created_at: 2026-09-01T13:03:03Z
updated_at: 2026-09-01T14:13:48Z
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
- [x] Minimal nvim config (from the scaffolding epic) registers the EDIFACT
      tree-sitter parser and enables highlighting for `.edi`/`.edifact` files
- [x] Headless e2e check extended to assert the buffer's tree-sitter parser
      is active and produces no `ERROR` nodes on the sample fixture
- [x] README updated with how to install/build the parser locally

## Summary of Changes
`editors/nvim/init.lua`: conditionally registers the tree-sitter parser via
`vim.treesitter.language.add`/`vim.treesitter.query.set` and starts
highlighting on `FileType edifact`, gated on `EDIFACT_TS_PARSER` being set
so plain LSP-only testing doesn't require the tree-sitter toolchain.
`scripts/e2e.sh` builds the grammar (`npm install` on first run, then
`tree-sitter build`) and exports `EDIFACT_TS_PARSER`/`EDIFACT_TS_HIGHLIGHTS`.
`scripts/e2e_check.lua` gained `check_treesitter`, asserting
`vim.treesitter.get_parser` succeeds and the root node has no errors on
`testdata/minimal.edi` (skips gracefully, doesn't fail, if the env var isn't
set). README updated (main README's new "Syntax highlighting" section, plus
a dedicated `tree-sitter-edifact/README.md`); manually verified the exact
documented command sequence works standalone, not just through the e2e
script.

Found via testing (same class of bug as the formatting epic, not a new
one): `check_treesitter` re-`:edit`s the same `minimal.edi` that
`check_minify` had just left modified-but-unsaved, which nvim refuses
(E37) since it's a *reload* of the current dirty buffer, not a switch to a
different one -- that distinction is why this didn't surface until a check
reused a fixture another check had already dirtied. Fixed by making every
`vim.cmd.edit(...)` call in the script forceful (`{ path, bang = true }`),
not just the one that happened to trigger it, since any check could
plausibly revisit a fixture in the future.
