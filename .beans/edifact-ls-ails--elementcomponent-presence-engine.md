---
# edifact-ls-ails
title: Element/component presence engine
status: completed
type: feature
priority: normal
created_at: 2026-09-01T19:56:59Z
updated_at: 2026-09-01T19:58:14Z
parent: edifact-ls-9ger
---

# Description

Build a generic element/component presence engine, independent of any
specific segment tag: a SegmentElementSchema describes one segment's
ordered element positions, each with its own Mandatory flag and an
ordered list of ComponentSchema (1 entry for a simple data element, N
for a composite), each component also with its own Mandatory flag. A
validator checks a real parsed Segment against a schema, reporting a
missing mandatory element (the element itself absent) or a missing
mandatory component (the element present but a required component
missing/empty within it).

Not wired to any real segment tag yet -- engine only, exercised by
hand-built test schemas, mirroring how edifact-ls-bygc proved the
structural engine before any real message data existed.

# Acceptance Criteria

[x] SegmentElementSchema / ElementSchema / ComponentSchema types defined
[x] Validator reports missing-mandatory-element and
missing-mandatory-component violations with accurate positions
(element-level for the former, component-level for the latter)
[x] A composite with all-conditional components never false-positives
(mirrors BGM/CTA's real shape)
[x] Unit tests using small hand-built schemas covering: clean pass,
missing mandatory simple element, missing mandatory composite
element, missing mandatory component within a present composite

## Summary of Changes

internal/edifact/segment_content.go: SegmentElementSchema{Elements},
ElementSchema{Name, Mandatory, Components}, ComponentSchema{Name,
Mandatory}. ValidateSegmentElements(schema, seg) walks each element
position: if the schema element is mandatory and absent from the
segment, reports "missing its mandatory element N (name)"; if present,
checks each mandatory component within it, reporting either the same
element-level message (when the element is simple -- exactly one
component) or a component-specific one naming both the element and the
component. No message-type or group nesting concerns at all -- this
engine only ever looks at one segment at a time.

5 unit tests, all against hand-built schemas (not real BGM/DTM/CTA
data yet, per the story's own scope) using buildSegment(t, src) to
build real Segments via Parse rather than hand-constructing AST
structs. Full suite (`make test`) green.
