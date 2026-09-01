---
# edifact-ls-6uvf
title: Info and warning severity diagnostics
status: todo
type: feature
priority: normal
created_at: 2026-09-01T16:13:54Z
updated_at: 2026-09-01T16:13:54Z
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
- [ ] `edifact.SeverityInfo` added alongside the existing Error/Warning
- [ ] `internal/lspserver`'s diagnostic-severity mapping covers all three
      (maps to LSP's Error/Warning/Information)
- [ ] Warning: a segment tag starting with "UN" that isn't a recognized
      service segment tag
- [ ] Info: a `UNA` present whose 6 delimiter characters exactly match
      `DefaultDelimiters()` (redundant, safe to omit)
- [ ] Unit tests for both new checks, plus an e2e check confirming a
      warning/info-level diagnostic actually renders with the right
      severity in nvim (not just error-level, which is already covered)
