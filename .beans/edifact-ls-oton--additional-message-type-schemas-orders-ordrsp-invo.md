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

# Status: 10 of 11 done; CONTRL blocked

DESADV, IFTSTA, INVRPT, DELFOR, ORDERS, ORDRSP, INVOIC, IFTMIN, PRICAT,
and APERAK are all done -- real data, registered, tested, e2e-verified.
CONTRL (edifact-ls-fh22) is blocked: its page 403s directly and, unlike
every other message here, isn't archived by the Wayback Machine under
any path or release checked. See that story for the full search
performed. Needs a working source before it can be picked up.

# Acceptance Criteria

[x] All 11 message types have real, sourced schema data registered for
(TYPE, D, 20A, UN) -- 10 of 11; CONTRL blocked, see above
[x] Each has unit tests against real fetched data (conformant + at
least one violation) -- for the 10 completed
[x] Each has an e2e check confirming a real violation surfaces in nvim
-- for the 10 completed
[x] Full suite (`make test`) and e2e harness (`make test-e2e`) green
throughout

## Summary of Changes

10 of 11 message types now have real structural schemas: DESADV (28
groups), IFTSTA (29), INVRPT (19, but depth 6 -- deepest despite
fewest groups), DELFOR (32), ORDERS (63), ORDRSP (60), INVOIC (55,
first message with a mandatory top-level *group*, not just mandatory
segments), IFTMIN (45, also a mandatory top-level group), PRICAT (60),
and APERAK (5, the smallest in the project). e2e harness grew from 10
checks (before this epic) to 23, each confirming a real structural
violation surfaces in a live nvim + LSP session.

Generalized IFTMCS's one-off column-parsing script into a reusable
tool, used identically across all 10 -- proving out the "generic
engine, per-message data" design at real scale: zero engine or
registry changes were needed across 10 more message types spanning
5-63 segment groups each.

A real, cross-cutting side effect surfaced and fixed: several existing
tests/fixtures used "ORDERS:D:96A:UN" as a neutral placeholder for "a
message type with no registered schema" -- which stopped being true
once ORDERS was actually registered (for a different release, D:20A).
Fixed by switching those specific tests/fixtures to a safe, clearly
fictional placeholder ("TESTMSG") rather than a real message type,
restoring their original isolation.

## Retro

- The plug-and-play design held up at real scale, not just in a
  single-message proof of concept: 10 message types, zero shared-code
  changes beyond new data files, confirmed by the fact that nothing in
  schema.go or schema_registry.go needed touching at all this epic.
- Registering more real message types surfaced a real, if narrow,
  blast radius: existing tests using an unregistered-type string as a
  neutral placeholder silently stopped being neutral. Worth remembering
  for the future -- any test that depends on a message type staying
  *unregistered* is implicitly coupled to the schema registry's global
  state, even if the test itself never mentions schemas.
- The schema matcher's order-sensitivity bit me while writing tests,
  not while building the engine: placing a conditional segment being
  tested for overflow *after* a later mandatory segment in the wire
  (rather than in its real schema-relative position) made it fall
  through as an unrelated "unexpected trailing segment" instead of
  being attributed as that segment's own overflow. Real UN/EDIFACT
  test fixtures need to respect schema order, not just segment
  presence -- worth remembering for any future message-type story.
- INVOIC and IFTMIN both turned out to have a mandatory *group*, not
  just mandatory segments, at the top level -- the first time that
  shape appeared in this project. Caught by tests failing against the
  real generated schema rather than assumed in advance, consistent
  with this project's pattern all along: read the real diagnostic, fix
  the test's understanding of the actual structure, don't guess.
- CONTRL's missing source is a genuine, currently-unresolved gap, not
  a shortcut -- documented clearly on its own story rather than silently
  dropped or faked with placeholder data.
