---
# edifact-ls-dk3k
title: APERAK D.00B schema data
status: completed
type: feature
priority: normal
created_at: 2026-09-02T14:10:37Z
updated_at: 2026-09-02T14:10:37Z
parent: edifact-ls-gdt6
---

# Description

User request: they're working with real APERAK D.00B messages and
saved the spec page themselves (`tmp/UN_EDIFACT Message APERAK
Release_ 00B.html`) -- unlike almost every other source in this
project, fetched directly from their own browser rather than needing
the Wayback Machine (service.unece.org otherwise 403s via Cloudflare
for this session's own fetches).

Transcribe APERAK's real D.00B branching diagram into SchemaNode data
and register it for (Type: "APERAK", Version: "D", Release: "00B",
Agency: "UN"), using the same mechanical extraction approach as every
other message type in this project (parse the source's exact rail-art
column positions programmatically, verify the result balances before
transcribing) rather than reading the ASCII nesting by eye.

Source: https://service.unece.org/trade/untdid/d00b/trmd/aperak_c.htm

# Acceptance Criteria

[x] APERAK D.00B's real branching diagram transcribed accurately
(position, tag, mandatory/conditional, max repeat, nesting) from the
user's saved source, verified to balance before transcription
[x] Registered for the exact tuple (APERAK, D, 00B, UN)
[x] Unit tests: a conformant message passes with no structural
violations; a real violation is flagged
[x] e2e check: opening a fixture with a structural violation shows the
diagnostic in nvim
[x] Source URL cited in the schema data's source comment, noting this
one didn't need the Cloudflare/Wayback workaround

## Summary of Changes

internal/edifact/aperak_d00b.go: extracted from the user's saved page
via the same script used for every prior message type (5 segment
groups, max nesting depth 2, balanced). Notable finding: this tree is
structurally identical, tag-for-tag and repeat-for-repeat, to
aperak_d20a.go's D.20A schema -- APERAK's structure hasn't changed
across 20 directory releases and revisions (D.00B revision 4,
2000-06-28, vs. D.20A). Registered as its own tuple anyway, since a
real message self-reporting "APERAK:D:00B:UN" still needs an exact
match per this project's release-specific design, even though the
underlying tree happens to match a different release's.

internal/edifact/aperak_d00b_test.go: registered, minimal conformant
pass, missing mandatory BGM, BGM exceeding its own cap of 1 -- same
shape as aperak_d20a_test.go.

testdata/aperak-d00b-violation.edi + scripts/e2e_check.lua: e2e check
(BGM repeated twice) confirms the diagnostic reaches a real nvim
session.

docs/SUPPORTED_MESSAGES.md regenerated. Full suite (`make test`) and
e2e harness (`make test-e2e`) green.

## Retro

- The user's own direct browser access continuing to be a real,
  legitimate unblock (not just a fallback) when Cloudflare blocks this
  session's own fetches -- consistent with what edifact-ls-9117's
  retro already noted for D.96A/D.96B. Worth treating a user-supplied
  saved page as at least as trustworthy a source as a Wayback capture,
  and reaching for it directly rather than re-attempting a fetch this
  session already knows will 403.
- Finding APERAK D.00B structurally identical to D.20A is a genuinely
  useful data point for anyone maintaining this registry: it suggests
  checking whether a newly-requested release of an already-registered
  message type might be a quick structural diff against the existing
  data, before assuming a full independent transcription is needed --
  though transcribing independently and comparing after (as done here)
  is still the safer default, since assuming similarity in advance
  risks missing a real difference.