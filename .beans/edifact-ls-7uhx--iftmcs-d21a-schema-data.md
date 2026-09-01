---
# edifact-ls-7uhx
title: IFTMCS D.21A schema data
status: completed
type: feature
priority: normal
created_at: 2026-09-01T17:03:46Z
updated_at: 2026-09-01T18:56:00Z
parent: edifact-ls-3uzr
blocked_by:
    - edifact-ls-ogqj
---

# Description

Transcribe the real IFTMCS D.21A branching diagram (revision 13,
2021-06-10; 41 segment groups, up to 4 levels of nesting, some groups
repeating up to 999 times) into the schema format from "Generic
segment/group schema engine", and register it for the tuple
(IFTMCS, D, 21A, UN). Scope is structural only (segment/group presence,
order, cardinality) -- not element-level content, per the epic's
non-goal.

Source: https://service.unece.org/trade/untdid/d21a/trmd/iftmcs_c.htm
section 4.3.1 "Segment table" (currently 403s directly via Cloudflare;
sourced via
http://web.archive.org/web/20240303212338/https://service.unece.org/trade/untdid/d21a/trmd/iftmcs_c.htm).

# Acceptance Criteria

[x] All 41 segment groups + top-level segments transcribed accurately
(position, tag, mandatory/conditional, max repeat, nesting) from the
cited source
[x] Registered for the exact tuple (IFTMCS, D, 21A, UN)
[~] Unit tests: a minimal conformant IFTMCS message passes with no
structural violations; fixtures for at least a missing-mandatory-group
case and an exceeded-repeat case each produce the expected diagnostic
-- adjusted after transcribing the real data (see summary): every one
of IFTMCS's 41 groups is itself conditional, and each group's sole
mandatory child is always its own leading/detecting segment, so
"missing mandatory group" and "missing mandatory child within a
started group" can never actually occur against this schema by
construction. Tested instead: clean/minimal pass, missing mandatory
BGM (the one real "missing mandatory" case in this schema), and CTA
exceeding its repeat cap.
[x] Source URL(s) cited in the schema data's source comment for future
cross-checking, including the Cloudflare caveat
[x] e2e check (carried over from edifact-ls-ogqj, which added the
registry/dispatch but had no real message type to exercise it
against): opening a fixture with a structural IFTMCS violation shows
the diagnostic in nvim

## Summary of Changes

internal/edifact/iftmcs_d21a.go (new): the full IFTMCS D.21A branching
diagram as SchemaNode data (41 groups, ~190 segment/group entries),
registered via init()/RegisterSchema for (IFTMCS, D, 21A, UN). Source
cited in the file's doc comment, including the Cloudflare-403 caveat
and the Wayback Machine URL actually used.

Transcription method, worth recording: the source renders the
branching diagram as ASCII rail art (nested "|"/"+" characters marking
each group's start/extent/end by fixed column position). Hand-reading
that reliably across 41 groups and 4 nesting levels is genuinely
error-prone -- and an earlier general-purpose fetch of this same table
(via a third-party mirror, used only for a rough complexity estimate
before this story) had in fact mischaracterized the nesting (claimed
segment group 18 contains groups 19-34; the real table has SG19 and
SG36 as top-level siblings of SG18, not children of it). So rather
than transcribing by eye, I wrote a one-off script that parsed the
source's exact column positions to mechanically determine each rail's
open/close points, then verified the result: the bracket stack ends
perfectly balanced across all 41 groups, and the derived group count
(41) and max nesting depth (4) independently match what that earlier
rough estimate got right. The verified tree was then code-generated
into the Go literal above, rather than hand-typed, so the transcription
step itself couldn't introduce new errors. Documented this approach in
the file's doc comment so a future correction is re-derived the same
way rather than hand-edited.

A genuine finding while writing tests: every one of the 41 groups is
conditional (S=C) at the group level, and in every group the *only*
mandatory child is its own leading segment -- which my matching
algorithm requires to be present anyway just to recognize the group
occurred at all. So "missing mandatory group" and "missing mandatory
child of a started group" are structurally unreachable against this
real schema, even though the engine handles both correctly in the
abstract (already covered by schema_test.go's hand-built cases). Only
BGM (the lone non-leading mandatory node in the whole schema) can ever
produce "missing mandatory"; adjusted the AC and this story's tests to
match reality instead of forcing an artificial case.

internal/edifact/iftmcs_d21a_test.go: 4 unit tests against the real
registered schema (registration check, clean pass, missing BGM,
exceeded repeat).

testdata/iftmcs-violation.edi + scripts/e2e_check.lua: a new e2e check
-- an envelope-clean IFTMCS message (valid UNB/UNH/UNT/UNZ) with CTA
repeated twice against its cap of 1 -- confirms the diagnostic appears
in a real running nvim + LSP session, verified via `make test-e2e`
(now 10 checks, all passing). This is the first real end-to-end
confirmation of the feature the user actually asked for.

Full suite (`make test`) and e2e harness (`make test-e2e`) green.
