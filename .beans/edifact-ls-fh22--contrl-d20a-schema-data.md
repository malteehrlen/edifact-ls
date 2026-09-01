---
# edifact-ls-fh22
title: CONTRL D.20A schema data
status: todo
type: feature
priority: normal
created_at: 2026-09-01T20:13:55Z
updated_at: 2026-09-01T20:13:55Z
parent: edifact-ls-oton
---

# Description

Transcribe CONTRL's (Control (syntax acknowledgment)) real D.20A branching diagram into
SchemaNode data and register it for (Type: "CONTRL", Version: "D",
Release: "20A", Agency: "UN"), following the exact approach
edifact-ls-7uhx used for IFTMCS (see that epic's parent for the
approach: fetch the real segment table -- 403s directly via Cloudflare,
use the Wayback Machine archive -- then parse its ASCII rail-art
nesting mechanically rather than by eye, verify the result balances,
and generate the Go data from the verified tree rather than
hand-typing it).

Source: https://service.unece.org/trade/untdid/d20a/trmd/cntrl_c.htm

# Acceptance Criteria

[ ] CONTRL's real branching diagram transcribed accurately (position,
tag, mandatory/conditional, max repeat, nesting) from the cited
source, verified to balance before transcription
[ ] Registered for the exact tuple (CONTRL, D, 20A, UN)
[ ] Unit tests: a conformant CONTRL message passes with no structural
violations; at least one fixture produces a real violation the
actual fetched structure supports (don't assume which kind --
IFTMCS and BGM/CTA both turned out to have no mandatory
groups/elements at some levels)
[ ] e2e check: opening a fixture with a structural CONTRL violation
shows the diagnostic in nvim
[ ] Source URL(s) cited in the schema data's source comment, including
the Cloudflare/Wayback caveat
