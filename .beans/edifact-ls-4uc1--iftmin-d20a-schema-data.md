---
# edifact-ls-4uc1
title: IFTMIN D.20A schema data
status: completed
type: feature
priority: normal
created_at: 2026-09-01T20:13:55Z
updated_at: 2026-09-01T20:42:53Z
parent: edifact-ls-oton
---

# Description

Transcribe IFTMIN's (Instruction to Transport) real D.20A branching diagram into
SchemaNode data and register it for (Type: "IFTMIN", Version: "D",
Release: "20A", Agency: "UN"), following the exact approach
edifact-ls-7uhx used for IFTMCS (see that epic's parent for the
approach: fetch the real segment table -- 403s directly via Cloudflare,
use the Wayback Machine archive -- then parse its ASCII rail-art
nesting mechanically rather than by eye, verify the result balances,
and generate the Go data from the verified tree rather than
hand-typing it).

Source: https://service.unece.org/trade/untdid/d20a/trmd/iftmin_c.htm

# Acceptance Criteria

[x] IFTMIN's real branching diagram transcribed accurately (position,
tag, mandatory/conditional, max repeat, nesting) from the cited
source, verified to balance before transcription
[x] Registered for the exact tuple (IFTMIN, D, 20A, UN)
[x] Unit tests: a conformant IFTMIN message passes with no structural
violations; at least one fixture produces a real violation the
actual fetched structure supports (don't assume which kind --
IFTMCS and BGM/CTA both turned out to have no mandatory
groups/elements at some levels)
[x] e2e check: opening a fixture with a structural IFTMIN violation
shows the diagnostic in nvim
[x] Source URL(s) cited in the schema data's source comment, including
the Cloudflare/Wayback caveat

## Summary of Changes

internal/edifact/iftmin_d20a.go: 45 segment groups (max nesting depth
4). Only BGM is mandatory among top-level *segments*, but segment
group 12 (leading with NAD, max repeat 99) is itself a mandatory
*group* -- the same shape INVOIC turned out to have, discovered
independently here.

internal/edifact/iftmin_d20a_test.go: registered, minimal conformant
pass (BGM+NAD), missing mandatory BGM, and CTA exceeding its cap of 1.
Caught a real mistake writing these: initially wrote the
minimal-conformant and missing-BGM tests without NAD at all (an
oversight, not fixed at the same time as the exceeded-repeat test's
ordering fix), so they failed against the real schema's actual
mandatory-group requirement -- fixed by including NAD in both, in its
correct schema-relative position after CTA.

testdata/iftmin-violation.edi + scripts/e2e_check.lua: e2e check
confirms the diagnostic reaches a real nvim session.

Full suite (`make test`) and e2e harness (`make test-e2e`) green.
