---
# edifact-ls-vxf5
title: Mason package definition
status: todo
type: feature
priority: low
created_at: 2026-09-01T13:03:42Z
updated_at: 2026-09-01T15:38:48Z
parent: edifact-ls-hlup
blocked_by:
    - edifact-ls-dyrz
    - edifact-ls-f31a
---

# Description
Package edifact-ls so it can be installed via Mason (`:MasonInstall
edifact-ls`), using a custom/private Mason registry pointed at this repo --
not a PR against the official `mason-registry` (that stays a "maybe
eventually", decoupled from whether Mason installation actually works).
Depends on real release binaries existing (edifact-ls-f31a).

# Acceptance Criteria
- [ ] A Mason package spec (`package.yaml`, mason-registry's format)
      referencing the GitHub release asset naming from edifact-ls-f31a,
      living in a small registry repo/directory
- [ ] Documented how to add it as an extra `registries` entry in a user's
      own `mason.nvim` config
- [ ] Verified locally: `:MasonInstall edifact-ls` actually installs and
      the resulting binary works with the standalone lspconfig snippet
- [ ] README documents this install path alongside the existing
      `go install` one
- [ ] Explicitly out of scope here: opening a PR against the official
      `mason-registry` -- revisit only if asked
