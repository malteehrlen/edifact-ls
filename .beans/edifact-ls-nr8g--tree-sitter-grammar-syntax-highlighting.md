---
# edifact-ls-nr8g
title: Tree-sitter grammar & syntax highlighting
status: todo
type: epic
priority: normal
created_at: 2026-09-01T13:02:49Z
updated_at: 2026-09-01T13:02:49Z
parent: edifact-ls-gdt6
---

# Goal
A tree-sitter grammar for EDIFACT so nvim (and other tree-sitter-based
editors) can highlight `.edi`/`.edifact` files directly, independent of the
LSP server's own parser.

# Acceptance Criteria
- [ ] `grammar.js` parses well-formed EDIFACT interchanges into a concrete
      syntax tree distinguishing segments, elements, components, tags, and
      delimiters, with no `ERROR` nodes on valid input
- [ ] Highlight query (`queries/highlights.scm`) maps node types to sensible
      highlight groups (segment tags, delimiters, data values, etc.)
- [ ] Grammar has a tree-sitter test corpus (`corpus/*.txt`) covering the
      same syntax variations as the parser core's unit tests
- [ ] Verified working inside nvim using the e2e harness from the scaffolding
      epic (highlighting renders without tree-sitter parse errors on the
      sample fixture)
