---
# edifact-ls-0m41
title: Registry source metadata + ListRegisteredSchemas
status: completed
type: feature
priority: normal
created_at: 2026-09-02T08:44:03Z
updated_at: 2026-09-02T08:47:11Z
parent: edifact-ls-9117
---

# Description

Add source-URL metadata to the schema registry and a query API, laying
the foundation for the CLI command and generated docs that come next.

`RegisterSchema` gains a third parameter: `source string`, the
canonical spec URL. The registry's internal value type changes from a
bare `Schema` to a small struct carrying both the schema and its
source. Add `ListRegisteredSchemas() []SchemaInfo` returning every
registered `MessageID` paired with its source, sorted by
Type/Version/Release/Agency for deterministic output.

Update all 12 existing `RegisterSchema` call sites (one per message-
type file) to pass the source URL already cited in each file's doc
comment.

# Acceptance Criteria

[x] `RegisterSchema(id MessageID, schema Schema, source string)` --
signature updated, all 12 existing call sites updated
[x] `SchemaInfo{ID MessageID, Source string}` and
`ListRegisteredSchemas() []SchemaInfo` exist, sorted deterministically
[x] `validateOneMessage` (and anywhere else reading the registry)
updated for the new internal value type
[x] Unit tests: ListRegisteredSchemas reflects registered schemas
correctly (using the plug-and-play pattern -- register a throwaway
schema, assert it appears with its source, clean up)
[x] Full suite (`make test`) still green -- this story changes no
externally-visible validation behavior, only adds queryable metadata

## Summary of Changes

internal/edifact/schema_registry.go: schemaRegistry's value type
changed from Schema to a private registeredSchema{Schema, Source}
struct; RegisterSchema takes a third source string argument;
SchemaInfo{ID, Source} + ListRegisteredSchemas() (sorted by
Type/Version/Release/Agency) added. validateOneMessage updated for the
new value type (rs.Schema).

All 12 existing message-type files updated to pass the source URL
already cited in each file's own doc comment as the new third argument
-- kept in sync, not duplicated data, since both now say the same
thing (the doc comment stays for the fuller context/caveats each file
carries; the argument is what's actually queryable at runtime).

internal/edifact/schema_registry_test.go: 5 existing throwaway
RegisterSchema calls updated with placeholder https://example.test/...
sources; 2 new tests -- ListRegisteredSchemas reflects a newly
registered schema's ID+Source, and sorts deterministically (using two
types that would land in reverse order under Go's deliberately-
randomized map iteration, to actually exercise the sort rather than
pass by coincidence).

Full suite (`make test`) and e2e harness (`make test-e2e`) green -- no
externally-visible validation behavior changed, purely additive
metadata.
