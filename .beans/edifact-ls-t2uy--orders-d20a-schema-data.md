---
# edifact-ls-t2uy
title: ORDERS D.20A schema data
status: completed
type: feature
priority: normal
created_at: 2026-09-01T20:13:55Z
updated_at: 2026-09-01T20:42:53Z
parent: edifact-ls-oton
---

# Description

Transcribe ORDERS's (Purchase Order) real D.20A branching diagram into
SchemaNode data and register it for (Type: "ORDERS", Version: "D",
Release: "20A", Agency: "UN"), following the exact approach
edifact-ls-7uhx used for IFTMCS (see that epic's parent for the
approach: fetch the real segment table -- 403s directly via Cloudflare,
use the Wayback Machine archive -- then parse its ASCII rail-art
nesting mechanically rather than by eye, verify the result balances,
and generate the Go data from the verified tree rather than
hand-typing it).

Source: https://service.unece.org/trade/untdid/d20a/trmd/orders_c.htm

# Acceptance Criteria

[x] ORDERS's real branching diagram transcribed accurately (position,
tag, mandatory/conditional, max repeat, nesting) from the cited
source, verified to balance before transcription
[x] Registered for the exact tuple (ORDERS, D, 20A, UN)
[x] Unit tests: a conformant ORDERS message passes with no structural
violations; at least one fixture produces a real violation the
actual fetched structure supports (don't assume which kind --
IFTMCS and BGM/CTA both turned out to have no mandatory
groups/elements at some levels)
[x] e2e check: opening a fixture with a structural ORDERS violation
shows the diagnostic in nvim
[x] Source URL(s) cited in the schema data's source comment, including
the Cloudflare/Wayback caveat

## Summary of Changes

internal/edifact/orders_d20a.go: 63 segment groups (max nesting depth
3 -- shallower but far wider than IFTMCS). Three mandatory top-level
nodes: BGM, DTM, and UNS (section control) -- more than the earlier
mid-sized batch, where at most two were mandatory.

internal/edifact/orders_d20a_test.go: registered, minimal conformant
pass (BGM+DTM+UNS), missing mandatory DTM, and PAI exceeding its cap
of 1. Caught and fixed a real test-construction mistake here: the
matcher is order-sensitive, so a conditional segment placed *after* a
later mandatory segment in the wire (rather than in its real schema
position) doesn't get attributed as that segment's own overflow -- it
just falls through as an unrelated "unexpected trailing segment"
instead. Fixed by keeping wire order consistent with schema order.

testdata/orders-violation.edi + scripts/e2e_check.lua: e2e check (PAI
repeated twice, correctly ordered) confirms the diagnostic reaches a
real nvim session.

Full suite (`make test`) and e2e harness (`make test-e2e`) green.
