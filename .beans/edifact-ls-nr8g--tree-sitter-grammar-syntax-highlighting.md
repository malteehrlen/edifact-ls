---
# edifact-ls-nr8g
title: Tree-sitter grammar & syntax highlighting
status: completed
type: epic
priority: normal
created_at: 2026-09-01T13:02:49Z
updated_at: 2026-09-01T14:14:10Z
parent: edifact-ls-gdt6
---

# Goal
A tree-sitter grammar for EDIFACT so nvim (and other tree-sitter-based
editors) can highlight `.edi`/`.edifact` files directly, independent of the
LSP server's own parser.

# Acceptance Criteria
- [x] `grammar.js` parses well-formed EDIFACT interchanges into a concrete
      syntax tree distinguishing segments, elements, components, tags, and
      delimiters, with no `ERROR` nodes on valid input
- [x] Highlight query (`queries/highlights.scm`) maps node types to sensible
      highlight groups (segment tags, delimiters, data values, etc.)
- [x] Grammar has a tree-sitter test corpus (`test/corpus/*.txt` --
      tree-sitter's standard location) covering the same syntax variations as
      the parser core's unit tests
- [x] Verified working inside nvim using the e2e harness from the scaffolding
      epic (highlighting renders without tree-sitter parse errors on the
      sample fixture)

## Summary of Changes
New `tree-sitter-edifact/` grammar (default delimiters only -- a custom
UNA-redefined delimiter set isn't honored, since that would require an
external C scanner; documented as a deliberate scope limit rather than a
gap). 6 corpus tests pass; zero ERROR/MISSING nodes across all `testdata/`
fixtures except the deliberately-malformed one. `highlights.scm` maps
tags/delimiters/data to standard capture names. Wired into the dev harness
behind an opt-in env var so plain LSP testing still needs nothing beyond Go;
`scripts/e2e.sh` builds and verifies it automatically when Node/npm/a C
compiler are available.

## Retro
- Environment setup was the bulk of the friction here, not the grammar
  logic itself: needed `build-essential` (no C compiler present), and
  needed to work out tree-sitter-cli 0.27's config requirements (a
  `tree-sitter.json` with a `metadata` block, which isn't obvious from
  `tree-sitter generate`'s own error messages until you hit them one at a
  time) and its corpus test format (`---` separating source from expected
  tree, which isn't the first thing `tree-sitter test` tells you when it
  silently reports "0 parses" instead of an error).
- The grammar itself needed two real fixes before it would generate at all:
  tree-sitter disallows a named rule that can match the empty string, which
  bit both the "empty component" and "empty element" cases. Solved by
  pushing `optional(...)` up to each *call site* instead of into the rule
  definition, and, for elements, splitting into two alternatives (starts
  with a component vs. starts with a bare ':') so every derivation is
  guaranteed non-empty. Worth remembering if this grammar gets extended.
- Same test-fixture-mistake pattern recurred once more (a corpus case
  missing a valid segment tag) and once again in the e2e check (reusing an
  already-dirtied buffer) -- both fixed the same way as prior epics: give
  the test real, complete input, and force `:edit` reloads since checks
  intentionally leave buffers unsaved.
- This was the last open, unblocked epic besides the deliberately-deferred
  packaging one (editor packaging & distribution, edifact-ls-hlup, still
  low priority and blocked-by the scaffolding epic only -- ready whenever
  you want a real end-user install path). No adjustments needed before
  continuing.
