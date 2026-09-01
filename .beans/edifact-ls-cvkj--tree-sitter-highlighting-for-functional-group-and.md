---
# edifact-ls-cvkj
title: Tree-sitter highlighting for functional group and section-control segments
status: completed
type: feature
priority: normal
created_at: 2026-09-01T16:14:14Z
updated_at: 2026-09-01T16:26:39Z
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
- [x] A tree-sitter corpus test covering a functional-group-wrapped
      interchange (UNB, UNG, UNH/UNT x2, UNE, UNZ) and one using UNS,
      parsing with zero ERROR/MISSING nodes
- [x] highlights.scm distinguishes service segment tags (UN*) from user
      data segment tags as a deliberate choice, not an accident of the
      existing single @keyword capture -- decide and document whichever
      way we go
- [x] e2e check confirms highlighting still activates without tree-sitter
      errors on a fixture using functional groups

## Summary of Changes
`tree-sitter-edifact/queries/highlights.scm`: service segment tags
(matching `^UN` via a `#match?` predicate on the existing `segment_tag`
node -- a query-level distinction, no grammar changes needed) get
`@keyword.directive` in addition to the base `@keyword`; confirmed in a
real nvim session (`vim.treesitter.get_captures_at_pos`) that the
later-declared `keyword.directive` pattern actually wins the render for
UN*-prefixed tags while ordinary tags (e.g. BGM) only carry `@keyword`.

2 new corpus tests (a functional-group-wrapped interchange with two
messages, a UNS-using message) parse with zero ERROR/MISSING nodes.
`check_treesitter` in the e2e harness gained an optional fixture parameter
and now also runs against a new `testdata/functional-group.edi` fixture
through the real headless-nvim pipeline.
