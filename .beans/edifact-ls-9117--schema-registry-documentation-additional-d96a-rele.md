---
# edifact-ls-9117
title: Schema registry documentation + additional D.96A releases
status: completed
type: epic
priority: normal
created_at: 2026-09-02T08:43:30Z
updated_at: 2026-09-02T09:19:07Z
parent: edifact-ls-gdt6
---

# Goal

Two related asks from the user, both about the structural schema
registry (edifact-ls-3uzr onward):

1. Document exactly which message specifications (type/version/release/
   agency) are actually supported -- currently discoverable only by
   reading source or triggering the "no schema registered" info
   diagnostic by trial and error.
2. Register multiple releases of the same "popular" message type, not
   just whichever single release happened to get sourced first.

# Design: the registry is already the right mechanism for #2

`schemaRegistry` is keyed on the *full* MessageID tuple (type, version,
release, agency), not just type -- adding a second release of ORDERS is
already just "another file, another RegisterSchema call, different
Release string," the same plug-and-play mechanism verified back in
edifact-ls-ogqj. No registry/validator/diagnostics changes needed for
breadth; #2 is purely sourcing more pages through the existing pipeline.

# Design: generate documentation from the registry, don't hand-write it

Hand-maintained docs drift from the actual registry contents (we just
hit a case -- CUSCAR -- where even *we* didn't know a message's real
release until opening the file). The registry has all the data needed;
it just isn't queryable or exposed yet:

1. `RegisterSchema` gains a third parameter, `source string` (the
   canonical spec URL) -- carried as data instead of living only in a
   Go doc comment. Touches all 12 existing call sites (mechanical,
   one line each) plus the 4 new D.96A ones land using the new
   signature from day one.
2. `ListRegisteredSchemas() []SchemaInfo` -- the single source of truth
   for "what's supported," sorted for determinism.
3. A `edifact-ls schemas` CLI subcommand using it directly.
4. A generated `docs/SUPPORTED_MESSAGES.md`, produced by a small
   `tools/gendocs` program (`make docs`) from the same
   `ListRegisteredSchemas()` data, with a test
   (`TestSupportedMessagesDocIsUpToDate`) that fails if the checked-in
   file doesn't match what regenerating it would produce -- the same
   idea as a `gofmt -l` check, so docs and code structurally cannot
   disagree.
5. One line in the README's Features section pointing at the generated
   doc.

# Scope for #2 (this round): D.99B, not D.96A as originally picked

Originally scoped as D.96A for ORDERS, ORDRSP, INVOIC, DESADV -- picked
by the user as a widely-used real-world release (same era as their own
CUSCAR file's D:96B). D.96A turned out to be universally unusable: every
message page under that release is the identical ~1KB placeholder stub,
each with exactly one Wayback capture ever, confirmed across every
message type checked (not just these four). D.95B-D.98B have no
archived captures at all. The user then fetched real D.96B pages
directly from the live site, but they turned out to be the wrong
section (segment-clarification narrative "boilerplate text," not the
branching diagram); further guesses at the real page's URL suffix
couldn't be confirmed reachable from this environment, and no UNTDID
index page surfaced a working link either. Fell back to **D.99B** --
the same release already proven to work for CUSCAR (edifact-ls-076u) --
by mutual agreement. Not exhaustive; more releases/types (including
another attempt at D.96B, if a working source turns up later) can
follow the same pattern.

# Acceptance Criteria

[x] RegisterSchema carries a source URL; all existing + new call sites
updated
[x] ListRegisteredSchemas() and a `edifact-ls schemas` CLI subcommand
exist and reflect the real registry
[x] docs/SUPPORTED_MESSAGES.md is generated (via `make docs`) from the
registry, with a test enforcing it stays in sync
[x] README links to the generated doc
[x] ORDERS, ORDRSP, INVOIC, DESADV each also registered for D.99B
(real sourced data, tested, e2e-verified), alongside their existing
D.20A registrations

## Summary of Changes

Both original asks delivered. Documentation: RegisterSchema now carries
a source URL, ListRegisteredSchemas() is the single source of truth,
surfaced via both `edifact-ls schemas` and a generated
docs/SUPPORTED_MESSAGES.md with a test that fails if the two ever
disagree (verified directly by deliberately desyncing the doc and
confirming the test catches it). Multiple releases: ORDERS, ORDRSP,
INVOIC, and DESADV are now each registered under two releases (D.20A
and D.99B) -- proven independently correct via a dedicated test
(TestORDERSBothReleasesIndependentlyCorrect) rather than just asserted.
16 message-type registrations total, up from 12.

The release actually used (D.99B) isn't what was originally asked for
(D.96A) or what the user's own real files declare (D.96B) -- both
turned out to be unreachable after a genuinely thorough search
(archived-page CDX checks across every nearby release, direct fetches,
the user pulling pages from the live site themselves, several URL-
suffix guesses). Landed on D.99B by mutual agreement as the best
available real alternative, consistent with the CUSCAR precedent.

## Retro

- Sourcing turned out to be the dominant cost of this epic, not the
  mechanism -- the registry/documentation half (story 1+2) went exactly
  to plan in one sitting, while the "just add D.96A" half turned into a
  multi-message back-and-forth spanning several releases and both sides
  fetching pages. Worth remembering for scoping future message-type
  work: verifying a release is actually sourceable is not a formality,
  it can be the majority of the real effort.
- The user's own live-site access ended up being the actual unblock,
  twice: once for the (ultimately wrong-section) D.96B pages, and
  implicitly by confirming D.96A truly doesn't exist live either
  (not just an archive gap). When Cloudflare blocks direct fetches and
  Wayback has nothing usable, asking the user to fetch directly is a
  real, legitimate escalation path -- not just a fallback to reach for
  reluctantly. Doing so early (rather than continuing to guess URL
  suffixes solo) would likely have saved a round or two.
- Caught my own overcorrection: when the user first said "_d, not _c",
  I found an old archived index page that only linked to "_c.htm" and
  used that to second-guess them instead of trusting their direct,
  current observation of the live site. Their own fetched page then
  confirmed *they* were right about the underlying pattern (_c.htm
  ≠ the branching diagram) even though "_d.htm" itself didn't pan out
  either. Old archived data is evidence, not a trump card over what
  someone is looking at right now -- worth weighting current, direct
  observation appropriately next time this kind of conflict comes up.
- Knowing when to stop was as important as knowing how to search: after
  the _s.htm guess also came up empty and no index page surfaced a
  working link, proposing the concrete, already-proven D.99B fallback
  (rather than proposing yet another guess) let the user close the loop
  quickly once they'd had enough of the hunt themselves.
- The generic extraction tooling (now used across 16 registrations
  spanning D.20A/D.21A/D.99B) continues to hold up without further
  changes this round -- no new page-format quirks hit this time,
  unlike CUSCAR's 4-digit-position/change-indicator surprise.
