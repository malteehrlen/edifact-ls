---
# edifact-ls-076u
title: CUSCAR D.99B schema data
status: completed
type: feature
priority: normal
created_at: 2026-09-02T08:28:58Z
updated_at: 2026-09-02T08:30:29Z
parent: edifact-ls-5wdy
---

# Description

Transcribe CUSCAR's real D.99B branching diagram into SchemaNode data
and register it, following the exact approach edifact-ls-7uhx used for
IFTMCS. See the parent epic for the release-selection story (D.96B not
archived; D.96A archived but a stub; D.99B chosen by the user over
D.01B).

Source: https://service.unece.org/trade/untdid/d99b/trmd/cuscar_c.htm

# Acceptance Criteria

[x] CUSCAR's real branching diagram transcribed accurately (position,
tag, mandatory/conditional, max repeat, nesting) from the cited
source, verified to balance before transcription
[x] Registered for the exact tuple (CUSCAR, D, 99B, UN)
[x] Unit tests: a conformant CUSCAR message passes with no structural
violations; at least one fixture produces a real violation the
actual fetched structure supports
[x] e2e check: opening a fixture with a structural CUSCAR violation
shows the diagnostic in nvim
[x] Source URL(s) cited in the schema data's source comment, including
the Cloudflare/Wayback caveat and the release-selection note

## Summary of Changes

internal/edifact/cuscar_d99b.go: 17 segment groups, max nesting depth
4. Only BGM is unconditionally mandatory at the top level; segment
group 14 (leading GID, mandatory) only applies once segment group 7
(CNI) then segment group 8 (RFF) are actually entered -- both
conditional, so a bare BGM alone is a genuinely clean minimal message.

Extending the reusable extraction script (from edifact-ls-oton) to this
older-vintage page surfaced two real format differences it hadn't hit
before: 4-digit position numbers instead of 5 ("0010" vs "00010"), and
this page actually uses the "change indicator" character column the
page's own legend documents (e.g. "0110 + DTM ..."), which the script's
regexes didn't account for. Fixed by generalizing both regexes rather
than hand-patching this file's data; spot-checked one result (segment
group 14's mandatory flag) directly against the raw source line to
confirm the generalized parsing is still trustworthy.

internal/edifact/cuscar_d99b_test.go: registered, minimal conformant
pass (bare BGM -- corrected after an initial wrong assumption that
segment group 14 was unconditionally mandatory; it's only mandatory
once its enclosing conditional groups are entered), missing mandatory
GID once segment groups 7/8 are entered, and BGM exceeding its own cap
of 1.

testdata/cuscar-violation.edi + scripts/e2e_check.lua: e2e check (BGM
repeated twice) confirms the diagnostic reaches a real nvim session.

Confirmed the actual outcome for the user's real file (test.edi,
declaring CUSCAR:D:96B:UN): correctly produces the "recognized type,
different release registered" info diagnostic rather than either
silence or a false structural match --

    3:1: info: no message specification registered for "CUSCAR" version D release 96B (agency UN); structural validation skipped -- registered for "CUSCAR": D:99B:UN

Full suite (`make test`) and e2e harness (`make test-e2e`, now 25
checks) green.
