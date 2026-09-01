---
# edifact-ls-oton
title: Additional message-type schemas (ORDERS, ORDRSP, INVOIC, DESADV, IFTMIN, IFTSTA, PRICAT, INVRPT, DELFOR, CONTRL, APERAK)
status: in-progress
type: epic
priority: normal
created_at: 2026-09-01T20:13:55Z
updated_at: 2026-09-01T20:22:52Z
parent: edifact-ls-gdt6
---

# Goal

Extend the structural schema registry (built for IFTMCS in
edifact-ls-3uzr) with 11 more common message types, using the exact
same plug-and-play mechanism that epic's own tests verified: no engine
or registry changes, only a new schema-data file + one RegisterSchema
call per message type. Sources are the D.20A UNECE message pages the
user supplied:

| Type | Spec URL |
| --- | --- |
| ORDERS | https://service.unece.org/trade/untdid/d20a/trmd/orders_c.htm |
| ORDRSP | https://service.unece.org/trade/untdid/d20a/trmd/ordrsp_c.htm |
| INVOIC | https://service.unece.org/trade/untdid/d20a/trmd/invoic_c.htm |
| DESADV | https://service.unece.org/trade/untdid/d20a/trmd/desadv_c.htm |
| IFTMIN | https://service.unece.org/trade/untdid/d20a/trmd/iftmin_c.htm |
| IFTSTA | https://service.unece.org/trade/untdid/d20a/trmd/iftsta_c.htm |
| PRICAT | https://service.unece.org/trade/untdid/d20a/trmd/pricat_c.htm |
| INVRPT | https://service.unece.org/trade/untdid/d20a/trmd/invrpt_c.htm |
| DELFOR | https://service.unece.org/trade/untdid/d20a/trmd/delfor_c.htm |
| CONTRL | https://service.unece.org/trade/untdid/d20a/trmd/cntrl_c.htm |
| APERAK | https://service.unece.org/trade/untdid/d20a/trmd/aperak_c.htm |

Note the release is **D.20A**, not D.21A like IFTMCS -- each schema
registers as MessageID{Type, Version: "D", Release: "20A", Agency:
"UN"}, matching what these specific pages actually declare, not
IFTMCS's release.

# Approach (per story, mirrors edifact-ls-7uhx)

1. Fetch the real segment table (the direct URL will very likely 403
   via Cloudflare, same as every other service.unece.org page hit so
   far this project -- check the Wayback Machine archive instead).
2. Don't hand-transcribe the ASCII rail-art nesting by eye -- iftmcs_d21a.go's
   doc comment explains why (an earlier informal summary of IFTMCS's
   own page got the nesting wrong). Reuse the same approach: parse the
   source's exact column positions mechanically, verify the bracket
   structure balances, generate the Go SchemaNode literal from the
   verified tree.
3. Register for (TYPE, D, 20A, UN). Cite the source URL (and Wayback
   URL actually used) in the data file's doc comment.
4. Unit tests: at minimum a conformant message and one real violation
   the actual fetched structure supports (mandatory-element-missing or
   exceeded-repeat -- whichever the real data actually has; IFTMCS
   and BGM/CTA both turned out to have no mandatory groups/elements at
   some levels, so don't assume before checking).
5. e2e check: a fixture with a real violation shows the diagnostic in
   nvim (scripts/e2e_check.lua, same pattern as
   testdata/iftmcs-violation.edi).

# Explicit non-goal

Element/component-level content validation for these message types'
segments (mirroring edifact-ls-9ger) is not in scope here -- this epic
only adds structural (segment/group presence, order, cardinality)
schemas, matching edifact-ls-3uzr's own scope boundary.

# Acceptance Criteria

[ ] All 11 message types have real, sourced schema data registered for
(TYPE, D, 20A, UN)
[ ] Each has unit tests against real fetched data (conformant + at
least one violation)
[ ] Each has an e2e check confirming a real violation surfaces in nvim
[ ] Full suite (`make test`) and e2e harness (`make test-e2e`) green
throughout
