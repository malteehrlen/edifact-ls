---
# edifact-ls-cvkj
title: Tree-sitter highlighting for functional group and section-control segments
status: todo
type: feature
priority: normal
created_at: 2026-09-01T16:14:14Z
updated_at: 2026-09-01T16:14:14Z
parent: edifact-ls-0d7g
blocked_by:
    - edifact-ls-zc0r
---

# Description
Now that UNG/UNE/UNS are recognized service segments (edifact-ls-zc0r),
make sure a real interchange using them highlights sensibly -- they're
already syntactically just ordinary segments to the grammar
(tree-sitter-edifact/grammar.js has no per-tag special-casing today, so
this is mostly a verification + possibly a distinct highlight group for
service segments generally so UNG/UNE/UNS read as visually related to
UNB/UNZ/UNH/UNT rather than indistinguishable from user data segments).

Reference: https://unece.org/DAM/trade/untdid/texts/d423.htm, sections 8.3.6/8.3.7/8.3.11.

# Acceptance Criteria
- [ ] A tree-sitter corpus test covering a functional-group-wrapped
      interchange (UNB, UNG, UNH/UNT x2, UNE, UNZ) and one using UNS,
      parsing with zero ERROR/MISSING nodes
- [ ] highlights.scm distinguishes service segment tags (UN*) from user
      data segment tags as a deliberate choice, not an accident of the
      existing single @keyword capture -- decide and document whichever
      way we go
- [ ] e2e check confirms highlighting still activates without tree-sitter
      errors on a fixture using functional groups
