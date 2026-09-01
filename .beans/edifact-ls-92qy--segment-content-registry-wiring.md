---
# edifact-ls-92qy
title: Segment content registry + wiring
status: completed
type: feature
priority: normal
created_at: 2026-09-01T19:56:59Z
updated_at: 2026-09-01T19:59:31Z
parent: edifact-ls-9ger
blocked_by:
    - edifact-ls-ails
---

# Description

Wire the engine into a generic, message-type-independent pass: a
registry mapping segment tag -> SegmentElementSchema (RegisterSegmentElementSchema),
and ValidateSegmentContent(ic *Interchange) ErrorList that walks every
segment in the interchange and validates it against a registered
schema for its tag, if one exists. No MessageID/schema-registry
involvement at all -- confirmed in the epic that segment structure is
intrinsic to the tag, not per-message-type. Wire into edifact.Validate
alongside the existing checks, so it reaches diagnostics and the CLI
`check` command automatically, same as everything else.

# Acceptance Criteria

[x] In-memory registry keyed by plain segment tag (not MessageID),
with a registration API a later story uses to add BGM/DTM/CTA
[x] ValidateSegmentContent walks every segment in an Interchange and
validates any with a registered schema; segments with no registered
schema produce no diagnostic
[x] edifact.Validate calls it alongside Parse/ValidateEnvelopes/Lint/
ValidateMessageSchemas
[x] Unit tests covering: a segment with a registered schema and a
violation, a segment with no registered schema (silent), and
registering two unrelated tags back to back (plug-and-play, same
guardrail as edifact-ls-ogqj)

## Summary of Changes

internal/edifact/segment_content.go: added segmentElementSchemas
(map[string]SegmentElementSchema), RegisterSegmentElementSchema(tag,
schema), and ValidateSegmentContent(ic) which walks every segment and
validates it against a registered schema for its tag, if any.

internal/edifact/edifact.go: Validate now also calls
ValidateSegmentContent, appended after ValidateMessageSchemas -- so it
reaches LSP diagnostics and `edifact-ls check` automatically, same as
every other check in the pipeline.

3 new unit tests, including one registering two unrelated tags back to
back and validating a message using both without any shared-code
changes -- the same plug-and-play guardrail edifact-ls-ogqj
established for the message-schema registry. Full suite (`make test`)
and e2e harness (`make test-e2e`) green; no behavior change yet since
no real segment tags are registered until edifact-ls-arr3.
