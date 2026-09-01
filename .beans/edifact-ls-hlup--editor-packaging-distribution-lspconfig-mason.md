---
# edifact-ls-hlup
title: Editor packaging & distribution (lspconfig + Mason)
status: in-progress
type: epic
priority: low
created_at: 2026-09-01T13:03:42Z
updated_at: 2026-09-01T14:31:16Z
parent: edifact-ls-gdt6
blocked_by:
    - edifact-ls-untk
---

# Goal
Make edifact-ls easy for an end user (not just this repo's dev harness) to
install and wire up: an upstreamable `nvim-lspconfig` server definition and,
optionally, Mason packaging so it can be installed with `:MasonInstall`.

# Acceptance Criteria
- [x] Documented, copy-pasteable `nvim-lspconfig` server definition (or a
      minimal plugin) usable without the in-repo dev harness config
      (`editors/nvim/standalone_lsp.lua` + `standalone_treesitter.lua`;
      built-in `vim.lsp.config`, not the `nvim-lspconfig` plugin -- same
      dependency-free choice as the dev harness)
- [ ] Mason registry entry/package definition allowing installation via
      Mason, tested locally
- [ ] README covers: build from source, install release binary, install via
      Mason, and configure lspconfig for `.edi`/`.edifact` files
- [x] Not started before there's a meaningfully complete server to package
      (diagnostics/formatting land first) — revisit priority at that point

## Interim note (2026-09-01)
Local-install half done (both edifact-ls-dyrz and edifact-ls-95ap
completed): binary via `go install`/`make install`, standalone LSP +
tree-sitter config snippets, both manually verified against a genuinely
fresh (non-dev-harness) nvim config. Remaining: edifact-ls-vxf5 (Mason
packaging, needs published GitHub release binaries first) -- deliberately
paused here per an explicit request to discuss distribution/"getting listed
elsewhere" as a separate conversation before starting it. Epic stays
in-progress, not completed, until that's done.
