---
# edifact-ls-6uvf
title: Info and warning severity diagnostics
status: completed
type: feature
priority: normal
created_at: 2026-09-01T16:13:54Z
updated_at: 2026-09-01T16:21:07Z
parent: edifact-ls-0d7g
blocked_by:
    - edifact-ls-zc0r
---

# Description
`edifact.Severity` currently only distinguishes Error/Warning, and in
practice everything we emit today is an Error. Add a real Info tier, and
add at least one genuine check at each of the two non-error levels, so the
LSP diagnostics pipeline (already UTF-16-range-aware, see
internal/lspserver/diagnostics.go) has real severity variety to map to
DiagnosticSeverityInformation/DiagnosticSeverityWarning.

Grounded in https://unece.org/DAM/trade/untdid/texts/d423.htm, section 6.2: "User data segments must not be
created with the first two letters of the tag 'UN', as these are reserved
for use in service segments." We can't know for certain a UN-prefixed tag
is wrong (it could be a service segment we don't recognize, e.g. from
CONTRL/APERAK), but a tag starting with "UN" that isn't one of the service
segments we do recognize (UNA/UNB/UNZ/UNG/UNE/UNH/UNT/UNS, after
edifact-ls-zc0r) is worth a **warning**.

For an **info**-level example: a `UNA` service string advice segment
present but defining exactly the documented default delimiters (section
8.3.2) is valid but redundant -- worth a soft note, not a warning.

# Acceptance Criteria
- [x] `edifact.SeverityInfo` added alongside the existing Error/Warning
- [x] `internal/lspserver`'s diagnostic-severity mapping covers all three
      (maps to LSP's Error/Warning/Information)
- [x] Warning: a segment tag starting with "UN" that isn't a recognized
      service segment tag
- [x] Info: a `UNA` present whose 6 delimiter characters exactly match
      `DefaultDelimiters()` (redundant, safe to omit)
- [x] Unit tests for both new checks, plus an e2e check confirming a
      warning/info-level diagnostic actually renders with the right
      severity in nvim (not just error-level, which is already covered)

## Summary of Changes
`edifact.SeverityInfo` added; `internal/lspserver`'s severity mapping now
covers Error/Warning/Information. New `internal/edifact/lint.go`
(`Lint(*Interchange) ErrorList`, separate from `ValidateEnvelopes` since
these are advisory, not structural correctness checks): a warning for a
well-formed segment tag starting with "UN" that isn't one of the
recognized service segments, and an info note for a `UNA` that defines
exactly the default delimiters (redundant). Wired into
`diagnosticsForText` alongside the existing parse/envelope errors.

Two new fixtures (`testdata/lint-warning.edi`, `testdata/lint-info.edi`)
and `check_diagnostic` in the e2e harness gained an optional expected-
severity parameter, now exercised for all four severity/fixture
combinations (error x2, warning, info) through the real headless-nvim
pipeline -- not just asserted at the Go unit-test level.
