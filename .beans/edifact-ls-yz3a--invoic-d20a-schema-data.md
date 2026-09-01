---
# edifact-ls-yz3a
title: INVOIC D.20A schema data
status: completed
type: feature
priority: normal
created_at: 2026-09-01T20:13:55Z
updated_at: 2026-09-01T20:42:53Z
parent: edifact-ls-oton
---

# Description

Transcribe INVOIC's (Invoice) real D.20A branching diagram into
SchemaNode data and register it for (Type: "INVOIC", Version: "D",
Release: "20A", Agency: "UN"), following the exact approach
edifact-ls-7uhx used for IFTMCS (see that epic's parent for the
approach: fetch the real segment table -- 403s directly via Cloudflare,
use the Wayback Machine archive -- then parse its ASCII rail-art
nesting mechanically rather than by eye, verify the result balances,
and generate the Go data from the verified tree rather than
hand-typing it).

Source: https://service.unece.org/trade/untdid/d20a/trmd/invoic_c.htm

# Acceptance Criteria

[x] INVOIC's real branching diagram transcribed accurately (position,
tag, mandatory/conditional, max repeat, nesting) from the cited
source, verified to balance before transcription
[x] Registered for the exact tuple (INVOIC, D, 20A, UN)
[x] Unit tests: a conformant INVOIC message passes with no structural
violations; at least one fixture produces a real violation the
actual fetched structure supports (don't assume which kind --
IFTMCS and BGM/CTA both turned out to have no mandatory
groups/elements at some levels)
[x] e2e check: opening a fixture with a structural INVOIC violation
shows the diagnostic in nvim
[x] Source URL(s) cited in the schema data's source comment, including
the Cloudflare/Wayback caveat

## Summary of Changes

internal/edifact/invoic_d20a.go: 55 segment groups (max nesting depth
3). Genuinely different from every message before it in this whole
project: it has a *mandatory top-level group* (segment group 52,
leading with MOA, max repeat 100) in addition to mandatory top-level
segments (BGM, DTM, UNS) -- previously only leaf segments had ever
been mandatory at the top level.

internal/edifact/invoic_d20a_test.go: registered, minimal conformant
pass (BGM+DTM+UNS+MOA), missing mandatory DTM, and PAI exceeding its
cap of 1 -- with PAI, UNS, and MOA kept in real schema order in the
wire (PAI before UNS before segment group 52).

testdata/invoic-violation.edi + scripts/e2e_check.lua: e2e check
confirms the diagnostic reaches a real nvim session.

Full suite (`make test`) and e2e harness (`make test-e2e`) green.
