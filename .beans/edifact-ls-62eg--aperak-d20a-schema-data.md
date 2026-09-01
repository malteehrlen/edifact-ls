---
# edifact-ls-62eg
title: APERAK D.20A schema data
status: completed
type: feature
priority: normal
created_at: 2026-09-01T20:13:55Z
updated_at: 2026-09-01T20:42:53Z
parent: edifact-ls-oton
---

# Description

Transcribe APERAK's (Application Acknowledgment) real D.20A branching diagram into
SchemaNode data and register it for (Type: "APERAK", Version: "D",
Release: "20A", Agency: "UN"), following the exact approach
edifact-ls-7uhx used for IFTMCS (see that epic's parent for the
approach: fetch the real segment table -- 403s directly via Cloudflare,
use the Wayback Machine archive -- then parse its ASCII rail-art
nesting mechanically rather than by eye, verify the result balances,
and generate the Go data from the verified tree rather than
hand-typing it).

Source: https://service.unece.org/trade/untdid/d20a/trmd/aperak_c.htm

# Acceptance Criteria

[x] APERAK's real branching diagram transcribed accurately (position,
tag, mandatory/conditional, max repeat, nesting) from the cited
source, verified to balance before transcription
[x] Registered for the exact tuple (APERAK, D, 20A, UN)
[x] Unit tests: a conformant APERAK message passes with no structural
violations; at least one fixture produces a real violation the
actual fetched structure supports (don't assume which kind --
IFTMCS and BGM/CTA both turned out to have no mandatory
groups/elements at some levels)
[x] e2e check: opening a fixture with a structural APERAK violation
shows the diagnostic in nvim
[x] Source URL(s) cited in the schema data's source comment, including
the Cloudflare/Wayback caveat

## Summary of Changes

internal/edifact/aperak_d20a.go: the smallest message type in this
whole project so far -- 5 segment groups, max nesting depth 2. Only
BGM is mandatory at the top level, and its SG4 (leading with ERC,
"Application error information") matches APERAK's real purpose.

internal/edifact/aperak_d20a_test.go: registered, minimal conformant
pass, missing mandatory BGM, and BGM exceeding its own cap of 1.

testdata/aperak-violation.edi + scripts/e2e_check.lua: e2e check (BGM
repeated twice) confirms the diagnostic reaches a real nvim session.

Full suite (`make test`) and e2e harness (`make test-e2e`) green.
