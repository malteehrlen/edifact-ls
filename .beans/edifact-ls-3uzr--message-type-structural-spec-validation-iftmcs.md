---
# edifact-ls-3uzr
title: Message-type structural spec validation (IFTMCS)
status: completed
type: epic
priority: normal
created_at: 2026-09-01T17:03:46Z
updated_at: 2026-09-01T18:59:51Z
parent: edifact-ls-gdt6
---

# Goal

UN/EDIFACT messages self-report their type via UNH's S009 composite
(element 1: data elements 0065 message type, 0052 version, 0054 release,
0051 controlling agency -- e.g. `IFTMCS:D:21A:UN`). Everything built so
far (`Lint`, `ValidateEnvelopes`) validates generic interchange syntax
only; per-message-type *structure* -- which segments/groups a specific
message type requires, in what order, how many times each may repeat --
was an explicit non-goal of edifact-ls-0d7g. This epic closes that gap
for one concrete, real message type: IFTMCS.

The UN/EDIFACT branching diagram fully specifies this (position, tag,
mandatory/conditional status, max repeat, nesting) -- it's data to
transcribe, not a new parsing problem. Source for IFTMCS D.21A (revision
13, 2021-06-10): https://service.unece.org/trade/untdid/d21a/trmd/iftmcs_c.htm
(currently returns HTTP 403 from Cloudflare when fetched directly; content
was sourced via the Wayback Machine archive at
http://web.archive.org/web/20240303212338/https://service.unece.org/trade/untdid/d21a/trmd/iftmcs_c.htm
-- re-check the direct URL first if cross-checking later, in case the
block has lifted).

Design: a generic, message-type-agnostic schema tree (segment/group nodes
with mandatory/conditional + max-repeat + children) and a walker that
validates a parsed message's segment stream against it. Schemas are
looked up by the *full* (type, version, release, agency) tuple from
UNH's S009 -- not just the type code -- since the branching diagram is
release-specific and matching only on type risks validating against the
wrong release. It's fine that matching the full tuple is more upfront
work than matching on type alone.

Surfaces as ordinary diagnostics (same mechanism as `Lint`/
`ValidateEnvelopes` -- ties into `didOpen`/`didChange`, no new nvim
command needed), plus a CLI subcommand on the `edifact-ls` binary for
scripted/CI use, since spec compliance is naturally a batch/pipeline
check as much as an editing-time one.

# Explicit non-goal

Element/composite-level content *within* each segment (which sub-fields
are mandatory, code list values, data types) is a different, deeper kind
of validation -- tracked separately as its own epic (message content
validation), not started, intentionally deferred until this epic ships.

# Acceptance Criteria

[x] A generic schema type + validator exist that can check any parsed
message's segment/group sequence against a branching-diagram-shaped
schema (mandatory/conditional presence, order, max-repeat),
independent of any specific message type
[x] A schema registry looks up a schema by the full (message type,
version, release, agency) tuple read from UNH's S009 composite; a
recognized type with no schema for its exact version/release
produces an informational diagnostic naming what is available,
rather than silently skipping or validating against the wrong release
[x] IFTMCS D.21A's real branching diagram (41 segment groups) is encoded
as schema data and registered for (IFTMCS, D, 21A, UN), with tests
against both conformant and non-conformant fixtures
[x] A `edifact-ls check <file>` CLI subcommand runs the same validation
path and exits non-zero on any structural violation, for use outside
the editor (CI, scripts)

## Summary of Changes

All four stories completed: a generic, message-type-agnostic schema
engine (internal/edifact/schema.go); a schema registry keyed by UNH's
full self-reported (type, version, release, agency) tuple, wired into
the existing diagnostics pipeline automatically (schema_registry.go);
IFTMCS D.21A's real 41-group branching diagram, mechanically parsed
from the source's ASCII rail art and code-generated into Go data
(iftmcs_d21a.go) rather than hand-transcribed; and a `edifact-ls check
<file>` CLI subcommand sharing one canonical validation pipeline
(edifact.Validate) with the LSP server. Verified end-to-end: opening a
real IFTMCS message with a genuine structural violation shows the
diagnostic in a live nvim + LSP session, and the CLI reports the same
violation with the right exit code.

## Retro

- The plug-and-play design goal (stated explicitly by the user before
  implementation started) held up in practice: adding IFTMCS required
  touching only a new schema-data file plus one registration call, no
  changes to the engine or registry -- verified directly by a test
  that registers a second, unrelated hand-built message type and
  exercises both without any shared-code changes.
- Writing the nested-group unit test surfaced a real engine bug before
  any real data existed to hide it in: a repeating group's first child
  shares its own leading tag, so naively checking that child's repeat
  cap locally misfires whenever an outer occurrence of the group
  legitimately starts right after. Fixed by resolving that ambiguity
  one level up. Building the generic engine first, against small
  adversarial hand-built schemas, before wiring any real message type,
  paid for itself here.
- Sourcing the real spec data was the epic's biggest risk, and it was
  justified: an earlier general-purpose fetch of the same table (via a
  third-party mirror, used only for a rough scale estimate before this
  epic existed) had genuinely mischaracterized the nesting -- claiming
  segment group 18 contained groups 19-34, when the real table has
  SG19 and SG36 as top-level siblings of SG18. Hand-reading ASCII rail
  art across 41 groups and 4 nesting levels is exactly the kind of
  task that's easy to get subtly wrong and hard to notice; writing a
  small script to parse the source's exact column positions
  mechanically, and checking that the result balances (all 41 opens
  matched by closes) before transcribing it, caught this rather than
  propagating it into the shipped schema.
- Reality also corrected the plan mid-epic in a smaller way: the AC
  for edifact-ls-7uhx assumed a "missing mandatory group" test case
  would exist; the real data revealed every one of IFTMCS's 41 groups
  is conditional with its only mandatory child being its own
  leading/detecting segment, so that scenario can't actually occur.
  Adjusted the AC and tests to match reality rather than force an
  artificial case -- worth remembering that AC written before reading
  the real source data is a plan, not a guarantee.
- The "e2e check" AC on edifact-ls-ogqj couldn't be honestly satisfied
  when written -- there was nothing registered yet for it to exercise.
  Moving it to edifact-ls-7uhx (where real data existed) rather than
  faking a check now, or silently dropping it, kept the backlog
  honest about what was actually verified at each step.
- No adjustments needed before starting the next epic. The remaining
  backlog items this session created alongside this epic (message
  content/element-level validation, and the three hover tiers) are
  unblocked or already were; none require anything further from this
  epic's design beyond what's already in place.
