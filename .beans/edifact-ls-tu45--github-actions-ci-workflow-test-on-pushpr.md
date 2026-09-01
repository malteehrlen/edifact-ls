---
# edifact-ls-tu45
title: GitHub Actions CI workflow (test on push/PR)
status: completed
type: feature
priority: normal
created_at: 2026-09-01T15:38:35Z
updated_at: 2026-09-01T15:55:23Z
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
- [x] Verified: a pushed branch/PR actually triggers the workflow and it
      passes

## Summary of Changes
`.github/workflows/ci.yml` runs `make test` + `make test-e2e` on every push
and PR, using GitHub's own setup-go/setup-node/apt-get rather than
bootstrapping Hermit, with Neovim and npm dependencies cached across runs.

Verified for real: pushed to `main` twice (once triggering the initial
workflow, once after fixing stale Action versions -- see below), both runs
completed successfully in ~1 minute.

Found via testing (real bug, not local): the first successful run still
logged a GitHub-side deprecation warning -- the pinned Action majors
(checkout@v4, setup-go@v5, setup-node@v4, cache@v4) all still targeted the
Node.js 20 runtime GitHub is deprecating on Actions runners. Confirmed via
each action's release notes and bumped to checkout@v7, setup-go@v7,
setup-node@v7, cache@v6 (all in this same CI story since it shares the
workflow files with the release story's actions) -- also applied to
goreleaser-action@v7 in the release workflow for the same reason.
