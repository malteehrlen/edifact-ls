---
# edifact-ls-ujaa
title: highlights.scm query
status: todo
type: feature
priority: normal
created_at: 2026-09-01T13:03:03Z
updated_at: 2026-09-01T13:03:03Z
parent: edifact-ls-nr8g
blocked_by:
    - edifact-ls-l8pq
---

# Description
Highlight query mapping the grammar's node types to standard tree-sitter
highlight groups (e.g. `@keyword` for segment tags, `@punctuation.delimiter`
for separators, `@string`/`@constant` for data values).

# Acceptance Criteria
- [ ] `queries/highlights.scm` covers every node type produced by the grammar
- [ ] Uses standard/conventional capture names so nvim's default theme
      renders something sensible without custom user config
- [ ] Spot-checked visually against the sample fixture (screenshot or manual
      note in the bean is fine — full automated color assertions not required)
