---
# edifact-ls-vxf5
title: Mason package definition
status: todo
type: feature
priority: low
created_at: 2026-09-01T13:03:42Z
updated_at: 2026-09-01T13:03:42Z
parent: edifact-ls-hlup
blocked_by:
    - edifact-ls-dyrz
---

# Description
Package edifact-ls so it can be installed via Mason (`:MasonInstall
edifact-ls`), including release binary distribution if not already set up.

# Acceptance Criteria
- [ ] Release binaries are built/published in a way Mason can consume (e.g.
      GitHub releases per platform)
- [ ] Mason package/registry definition added (or PR opened against
      mason-registry) and installs successfully locally
- [ ] README documents the Mason install path
