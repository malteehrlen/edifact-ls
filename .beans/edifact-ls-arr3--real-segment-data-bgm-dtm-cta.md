---
# edifact-ls-arr3
title: 'Real segment data: BGM, DTM, CTA'
status: completed
type: feature
priority: normal
created_at: 2026-09-01T19:56:59Z
updated_at: 2026-09-01T20:01:07Z
parent: edifact-ls-9ger
blocked_by:
    - edifact-ls-92qy
---

# Description

Transcribe real UNSD segment definitions for BGM, DTM, and CTA into the
schema format from the engine story, and register all three. Sources
(fetched via the Wayback Machine, same Cloudflare-403-on-direct-fetch
caveat as iftmcs_d21a.go):
- BGM: https://service.unece.org/trade/untdid/d21a/trsd/trsdbgm.htm
- DTM: https://service.unece.org/trade/untdid/d21a/trsd/trsddtm.htm
- CTA: https://service.unece.org/trade/untdid/d21a/trsd/trsdcta.htm

DTM is the one with real mandatory structure (its element C507 is
mandatory, and within it component 2005 is mandatory); BGM and CTA are
entirely conditional, proving the validator doesn't false-positive on
that shape.

# Acceptance Criteria

[x] BGM, DTM, CTA element/component structures transcribed accurately
from the cited sources and registered
[x] Unit tests: a conformant DTM passes; a DTM missing its mandatory
component 2005 produces the expected violation; BGM and CTA with
minimal/no elements produce no violations (nothing mandatory)
[x] e2e check: opening a fixture with a real content violation (e.g. a
DTM missing its date qualifier) shows the diagnostic in nvim
[x] Source URLs cited in the data file's source comment

## Summary of Changes

internal/edifact/segment_elements_data.go (new): BGM, DTM, CTA
transcribed from the real UNSD pages and registered via
init()/RegisterSegmentElementSchema. Source URLs and the
Cloudflare/Wayback caveat cited in the file's doc comment.

internal/edifact/segment_elements_data_test.go: 6 unit tests against
the real registered data -- all three registered, a conformant DTM
passes, a DTM missing its qualifier component is flagged, a DTM
missing entirely is flagged, and minimal BGM/CTA (no elements at all)
never false-positive since nothing in either is actually mandatory.

testdata/content-violation.edi + scripts/e2e_check.lua: a new e2e
check -- a DTM missing its function-code-qualifier component --
confirms the diagnostic reaches a real nvim + LSP session:

    PASS: diagnostics for .../testdata/content-violation.edi include a message containing "function code qualifier" with severity ERROR

Also manually confirmed via the CLI:

    4:5: error: segment "DTM" element 1 (Date/time/period) is missing its mandatory component 1 (Date or time or period function code qualifier)

Full suite (`make test`) and e2e harness (`make test-e2e`, now 13
checks) green. Existing fixtures' real BGM/DTM/CTA usages (e.g.
testdata/minimal.edi's DTM) were already genuinely conformant, so
nothing regressed.
