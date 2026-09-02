---
# edifact-ls-3enb
title: ORDERS D.99B schema data
status: completed
type: feature
priority: normal
created_at: 2026-09-02T08:44:04Z
updated_at: 2026-09-02T09:18:38Z
parent: edifact-ls-9117
blocked_by:
    - edifact-ls-0m41
---

# Description

Transcribe ORDERS's (Purchase Order) real D.99B branching diagram into
SchemaNode data and register it alongside the existing D.20A schema,
for (Type: "ORDERS", Version: "D", Release: "99B", Agency: "UN"), using
the 3-argument RegisterSchema(id, schema, source) from
edifact-ls-0m41.

# Release: D.99B, not D.96A

Originally scoped as D.96A (see the parent epic). That fell through
for every message type, not just this one: every D.96A trmd page is
the identical ~1KB placeholder stub ("There is some standard text
here"), each with exactly one Wayback capture ever (confirmed via CDX
across ORDERS/ORDRSP/INVOIC/DESADV/IFTMCS/CUSCAR/INVRPT). Releases
D.95B through D.98B have no archived captures at all for these message
types.

The user then found and shared five real D.96B pages directly from the
live site -- but they turned out to be the wrong section ("Boilerplate
text of X", the segment-clarification narrative/notes page, ending in
a plain alphabetical segment index, not the branching diagram). Several
follow-up guesses at the real page's URL suffix (_d.htm, _s.htm) 
couldn't be confirmed archived or accessible from this environment, and
no UNTDID index/table-of-contents page surfaced a working link either.

Falling back to **D.99B** -- the same release already sourced
successfully for CUSCAR (edifact-ls-076u), with real, complete,
verified segment tables. The user's own real files (declaring D:96B)
still get the honest "recognized type, different release registered"
info diagnostic rather than silence or a false match, consistent with
CUSCAR's precedent.

Source: https://service.unece.org/trade/untdid/d99b/trmd/orders_c.htm
(403s directly via Cloudflare; archived via the Wayback Machine, same
as every other message-type story).

# Acceptance Criteria

[x] ORDERS D.99B's real branching diagram transcribed accurately
(position, tag, mandatory/conditional, max repeat, nesting) from the
cited source, verified to balance before transcription
[x] Registered for the exact tuple (ORDERS, D, 99B, UN), alongside the
existing (registered) D.20A schema for the same type -- both must
remain independently correct (existing D.20A tests still pass)
[x] Unit tests: a conformant ORDERS D.99B message passes with no
structural violations; at least one fixture produces a real
violation the actual fetched structure supports
[x] e2e check: opening a fixture with a structural ORDERS D.99B
violation shows the diagnostic in nvim
[x] Source URL cited via the RegisterSchema source argument and in the
schema data's doc comment, including the Cloudflare/Wayback caveat

## Summary of Changes

internal/edifact/orders_d99b.go: 60 segment groups, max depth 3. Same
shape as D.20A's ORDERS: BGM, DTM, and UNS all mandatory at the top
level.

internal/edifact/orders_d99b_test.go: registered, minimal conformant
pass, missing mandatory DTM, BGM exceeding its cap of 1, and a
dedicated test (TestORDERSBothReleasesIndependentlyCorrect) validating
both the D.20A and D.99B registrations against their own conformant
fixtures in the same test -- directly proving the "multiple releases
of a popular message type coexist correctly" goal this whole epic was
about, not just asserting it.

testdata/orders-d99b-violation.edi + scripts/e2e_check.lua: e2e check
confirms the diagnostic reaches a real nvim session.

Full suite (`make test`) and e2e harness (`make test-e2e`) green; `make
docs` regenerated to include the new registration.
