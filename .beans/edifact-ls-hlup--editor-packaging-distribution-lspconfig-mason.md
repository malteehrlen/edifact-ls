---
# edifact-ls-hlup
title: Editor packaging & distribution (lspconfig + Mason)
status: todo
type: epic
priority: low
created_at: 2026-09-01T13:03:42Z
updated_at: 2026-09-01T13:03:42Z
parent: edifact-ls-gdt6
blocked_by:
    - edifact-ls-untk
---

# Goal
Make edifact-ls easy for an end user (not just this repo's dev harness) to
install and wire up: an upstreamable `nvim-lspconfig` server definition and,
optionally, Mason packaging so it can be installed with `:MasonInstall`.

# Acceptance Criteria
- [ ] Documented, copy-pasteable `nvim-lspconfig` server definition (or a
      minimal plugin) usable without the in-repo dev harness config
- [ ] Mason registry entry/package definition allowing installation via
      Mason, tested locally
- [ ] README covers: build from source, install release binary, install via
      Mason, and configure lspconfig for `.edi`/`.edifact` files
- [ ] Not started before there's a meaningfully complete server to package
      (diagnostics/formatting land first) — revisit priority at that point
