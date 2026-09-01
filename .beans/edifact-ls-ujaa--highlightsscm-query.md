---
# edifact-ls-ujaa
title: highlights.scm query
status: completed
type: feature
priority: normal
created_at: 2026-09-01T13:03:03Z
updated_at: 2026-09-01T14:09:15Z
parent: edifact-ls-nr8g
blocked_by:
    - edifact-ls-l8pq
---

# Description
Highlight query mapping the grammar's node types to standard tree-sitter
highlight groups (e.g. `@keyword` for segment tags, `@punctuation.delimiter`
for separators, `@string`/`@constant` for data values).

# Acceptance Criteria
- [x] `queries/highlights.scm` covers every node type produced by the grammar
      that carries its own text (`segment_tag`, `una_advice`, `data`, `+`,
      `:`, `terminator`); structural container nodes (`source_file`,
      `segment`, `element`, `component`) carry no text of their own and are
      intentionally left uncaptured, per standard tree-sitter practice --
      highlighting them would just re-apply a color over their
      already-captured children.
- [x] Uses standard/conventional capture names so nvim's default theme
      renders something sensible without custom user config
- [x] Spot-checked visually against the sample fixture (screenshot or manual
      note in the bean is fine — full automated color assertions not required)

## Summary of Changes
`tree-sitter-edifact/queries/highlights.scm`: `@keyword` for segment tags,
`@keyword.directive` for the UNA advice segment, `@string` for data values,
`@punctuation.delimiter` for `+`/`:`/the terminator -- all standard tree-
sitter capture names nvim's built-in themes already style distinctly with
no user config. "Spot-checked visually" here means via `tree-sitter query
queries/highlights.scm testdata/minimal.edi`, confirming each token gets
the intended capture (evidence: segment tags -> keyword, delimiters ->
punctuation.delimiter, values -> string, in that exact CLI output). Full
in-editor rendering is verified end-to-end once the parser is registered
with nvim, which is story edifact-ls-87s0's job.
