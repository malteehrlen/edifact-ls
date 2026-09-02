---
# edifact-ls-13gu
title: Broad D.20A message-type schema coverage
status: completed
type: epic
priority: normal
created_at: 2026-09-02T09:27:29Z
updated_at: 2026-09-02T09:27:29Z
parent: edifact-ls-gdt6
---

# Goal

Exploratory, broad expansion of structural schema coverage, per the
user's explicit request: "do some exploratory work in the wayback
machine and add all schemas you can find... the process is pretty
streamlined now."

# Scope: D.20A, the entire archived message catalog for that release

A CDX survey of `service.unece.org/trade/untdid/d20a/trmd/` found 194
real message-type pages archived (all several KB+, comfortably above
the ~1KB stub size that made D.96A unusable), of which 10 are already
registered (APERAK, DELFOR, DESADV, IFTMIN, IFTSTA, INVOIC, INVRPT,
ORDERS, ORDRSP, PRICAT). D.20A is also this project's best-proven
release: 11 message types already sourced from it with zero format
surprises. That makes it the highest-value, lowest-risk target for a
broad pass -- 182 more real, archived, previously-unregistered message
types.

Other releases (D.21A, D.99B, D.01B, etc.) each likely have their own
large catalogs too, but are out of scope for this pass -- flagged as a
natural follow-up, not pursued here to keep this batch's scope
coherent.

# Approach: mechanical extraction at scale, generic tests instead of
# per-type bespoke ones

At this scale (182 message types), the earlier per-type ceremony
(hand-written conformant/missing-mandatory/exceeded-repeat tests, a
dedicated e2e fixture, a full retro-worthy story each) doesn't scale --
literally hundreds of near-identical test functions would be worse for
maintainability than the coverage is worth. Instead:

1. Fetch every page via the Wayback Machine (batched, rate-limited).
2. Run extract_msgtype.py's logic on every fetched page. Its own
   balance verification (every segment group's rails must open and
   close consistently) already catches structural transcription errors
   mechanically, without needing to eyeball each one -- this has been
   the core integrity mechanism since edifact-ls-7uhx (IFTMCS). Any
   page that fails extraction (a new format quirk, an unparseable
   table) is logged and skipped rather than blocking the batch.
3. Manually spot-check a sample (not all 182) against their raw source
   for sanity, the same way CUSCAR's segment group 14 was spot-checked
   -- covering a range of sizes/shapes, not just the easy cases.
4. Register every successfully-extracted schema.
5. Add three *generic* tests (in schema_registry_generic_test.go) that
   automatically validate *every* registered schema, current and
   future, by deriving test fixtures mechanically from each schema's
   own tree rather than hand-writing them:
   - TestAllRegisteredSchemasAcceptMinimalConformantMessage: the
     message built from exactly a schema's own unconditionally-
     mandatory tags validates clean.
   - TestAllRegisteredSchemasFlagMissingMandatoryTag: dropping the
     last tag from that minimal sequence is flagged.
   - TestAllRegisteredSchemasFlagExceededTopLevelRepeat: for schemas
     whose first top-level node has a small enough cap to synthesize,
     repeating it past that cap is flagged.
   These don't replace transcription-fidelity checks (that's the
   balance verification + spot-checks) -- they instead assure the
   *engine* correctly accepts/rejects what each schema claims, for
   every registered schema without per-type authoring.
6. Spot-check a handful (not all) of the new types with real e2e
   fixtures in a live nvim session, same as every prior epic -- proving
   the wiring reaches the editor, which doesn't need re-proving 182
   times once proven once per release-format.
7. Regenerate docs/SUPPORTED_MESSAGES.md.

# Acceptance Criteria

[x] Every archived D.20A message-type page fetched and run through
extraction; failures logged (not silently dropped) with a reason
[x] Every successfully-extracted schema registered with real,
balance-verified data and a cited source
[x] A representative sample manually spot-checked against raw source
[x] The three generic tests exist, cover every registered schema, and
pass
[x] A handful of new types e2e-verified in a real nvim session
[x] docs/SUPPORTED_MESSAGES.md regenerated and accurate
[x] Full suite (`make test`) and e2e harness (`make test-e2e`) green

## Summary of Changes

Fetched all 182 previously-unregistered D.20A message-type pages found
by the CDX survey via the Wayback Machine (batched, 2.5s-delayed,
direct-year-redirect fetching), extracted every one with
extract_msgtype.py (182/182 succeeded after a real parser fix, see
below), and registered all of them in
internal/edifact/<tag>_d20a.go (one file per type, same header/doc
format as every prior message-type file, each citing its Wayback
source URL). Total registered schemas: 198 (16 pre-existing + 182 new).

Two real bugs surfaced and were fixed as part of this batch, not just
data added:

1. extract_msgtype.py couldn't parse QALITY's SPS segment, whose
   descriptive name is long enough to wrap onto a continuation line
   that carries the actual M/C/repeat-count columns. Added
   preprocess_wrapped_names() to detect and merge a partial pos+tag+name
   line with its continuation line before the main per-line parse,
   stripping the primary line's own trailing rail characters first so
   they don't leak into the merged name text.
2. internal/edifact/schema.go's matchSequence had a real correctness
   bug: two schema-tree siblings sharing a leading tag (UN/EDIFACT's
   standard double-UNS header/detail/summary boundary convention, also
   common in GOVCBR where many groups all lead with NAD) could produce
   false "exceeded repeat" / "missing mandatory" diagnostics on
   genuinely valid messages. Caught by the new generic tests on
   CONPVA, CUSDEC, and GOVCBR. Fixed by replacing the
   immediate-next-sibling check with laterSiblingSharesLeadingTag,
   which scans all remaining siblings so a conditional sibling sitting
   between the two same-tag nodes (as in CONPVA's real BII group
   between its two UNS nodes) doesn't defeat the check. Three
   regression tests added to schema_test.go, including one confirming
   genuine overflow is still correctly flagged (no overcorrection).

Added internal/edifact/schema_registry_generic_test.go: three tests
that mechanically derive fixtures from every registered schema's own
tree (rather than hand-writing per-type tests, which doesn't scale to
198 types) and validate the engine's accept/reject behavior against
all of them automatically -- this is what caught bug #2 above, and
will keep catching this class of engine bug against every future
registration for free.

Added 6 new e2e fixtures under testdata/ (qality-violation,
cusdec-conformant, genral-violation, mscons-violation,
coarri-violation, baplie-conformant) and a new check_no_error_diagnostic
helper in scripts/e2e_check.lua (the existing check_diagnostic helper
can only assert a diagnostic *appears*; this asserts none of severity
ERROR does, needed to positively confirm the double-UNS fix works
end-to-end). All 35 e2e checks pass.

Regenerated docs/SUPPORTED_MESSAGES.md via `make docs`; the
staleness-enforcing test confirms it matches the registry.

## Retro

- The generic, schema-derived testing strategy was the right call at
  this scale and proved its worth immediately: it caught two real bugs
  (one in the extraction tool, one in the core matching engine) that a
  narrower "spot-check a sample" pass alone would very plausibly have
  missed, since neither QALITY nor the double-UNS messages were
  necessarily going to be in whatever sample got hand-picked for
  spot-checking.
- The double-UNS bug is the more significant find of the two: it's a
  genuine engine-correctness defect that could have produced wrong
  diagnostics on real, valid UN/EDIFACT traffic using the standard
  section-boundary convention, not just a gap in data coverage. Worth
  remembering that broad data-sourcing passes are also a good forcing
  function for exercising engine code paths that narrow, one-type-at-a-
  time work doesn't stress as hard.
- First fix attempt for the double-UNS bug (checking only the
  immediate next sibling) looked sufficient against CUSDEC and GOVCBR
  but wasn't -- CONPVA's real schema has a conditional group sitting
  between its two UNS nodes, which only showed up by directly
  inspecting the extracted tree rather than reasoning abstractly about
  what "adjacent" should mean. Went back to the source data before
  generalizing the fix, which is what caught it.
- Caught and fixed a small fixture-arithmetic bug in my own test data
  (cusdec-conformant.edi's UNT count) via the e2e harness itself,
  exactly the kind of error check_no_error_diagnostic is meant to
  surface -- a good sign the new helper is pulling its weight, and a
  reminder that hand-counted segment counts in fixtures remain an
  easy, recurring mistake to make in this project.
- Scoping this as one batch operation instead of one story per type
  (unlike every prior message-type epic) was the right call for 182
  types -- per-type ceremony would have produced hundreds of
  near-identical stories with no real decision-making in most of them.
  The judgment call to instead invest in generic tests + spot-checks +
  a real engine-level fix is the kind of thing per-type stories
  wouldn't have surfaced as clearly anyway.
