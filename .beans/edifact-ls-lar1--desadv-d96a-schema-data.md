---
# edifact-ls-lar1
title: DESADV D.96A schema data
status: todo
type: feature
priority: normal
created_at: 2026-09-02T08:44:04Z
updated_at: 2026-09-02T08:44:04Z
parent: edifact-ls-9117
blocked_by:
    - edifact-ls-0m41
---

# Description

Transcribe DESADV's (Despatch Advice) real D.96A branching diagram into
SchemaNode data and register it alongside the existing D.20A schema
for (Type: "DESADV", Version: "D", Release: "96A", Agency: "UN"),
following the same approach as every prior message-type story: fetch
the real segment table (403s directly via Cloudflare, use the Wayback
Machine archive), parse its ASCII rail-art nesting mechanically rather
than by eye, verify the result balances, generate the Go data from the
verified tree. Uses the new 3-argument RegisterSchema(id, schema,
source) from edifact-ls-<registry-story>.

Source: https://service.unece.org/trade/untdid/d96a/trmd/desadv_c.htm

# Acceptance Criteria

[ ] DESADV D.96A's real branching diagram transcribed accurately
(position, tag, mandatory/conditional, max repeat, nesting) from the
cited source, verified to balance before transcription
[ ] Registered for the exact tuple (DESADV, D, 96A, UN), alongside the
existing (registered) D.20A schema for the same type -- both must
remain independently correct (existing D.20A tests still pass)
[ ] Unit tests: a conformant DESADV D.96A message passes with no
structural violations; at least one fixture produces a real
violation the actual fetched structure supports
[ ] e2e check: opening a fixture with a structural DESADV D.96A
violation shows the diagnostic in nvim
[ ] Source URL cited via the RegisterSchema source argument and in the
schema data's doc comment, including the Cloudflare/Wayback caveat
