---
# edifact-ls-tu45
title: GitHub Actions CI workflow (test on push/PR)
status: in-progress
type: feature
priority: normal
created_at: 2026-09-01T15:38:35Z
updated_at: 2026-09-01T15:45:56Z
parent: edifact-ls-u9x8
---

# Description
Run make test and make test-e2e automatically on every push and PR.
Independent of the release workflow (S2) -- can be built in parallel.

Uses GitHub's standard actions/setup-go + actions/setup-node + apt-get for
a C compiler, rather than bootstrapping Hermit in CI: Hermit's value is
reproducible *local* dev environments, but CI runners already provide
well-supported, cached setup actions for exactly these toolchains, so
routing through Hermit there would add fragility (network-dependent
bootstrap, cache-unfriendly) for no real benefit.

# Acceptance Criteria
- [x] `.github/workflows/ci.yml` triggers on push and pull_request
- [x] Sets up Go (matching go.mod's version) and Node.js, installs a C
      compiler (build-essential), and runs `make test` and `make test-e2e`
- [x] The e2e job's Neovim download/tree-sitter build steps are cached
      appropriately so CI runs aren't slower than necessary
- [ ] Verified: a pushed branch/PR actually triggers the workflow and it
      passes
