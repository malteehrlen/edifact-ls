---
# edifact-ls-l8pq
title: 'grammar.js: segments/elements/components/delimiters'
status: todo
type: feature
priority: normal
created_at: 2026-09-01T13:03:03Z
updated_at: 2026-09-01T13:03:03Z
parent: edifact-ls-nr8g
---

# Description
Write the core tree-sitter grammar for EDIFACT syntax: interchange envelope,
segments, data elements, composite/component elements, and delimiters
(including handling a custom `UNA`-defined delimiter set, if feasible within
tree-sitter's parsing model — otherwise document the limitation).

# Acceptance Criteria
- [ ] `grammar.js` builds cleanly with `tree-sitter generate`
- [ ] Parses the same sample interchanges used by the parser core's unit
      tests with zero `ERROR`/`MISSING` nodes
- [ ] Tree-sitter test corpus under `corpus/` covering default delimiters,
      composite elements, and at least one realistic full interchange
- [ ] `tree-sitter test` passes locally
