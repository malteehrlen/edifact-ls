---
# edifact-ls-1cwj
title: Segment description data
status: completed
type: feature
priority: normal
created_at: 2026-09-01T17:53:47Z
updated_at: 2026-09-01T19:39:39Z
parent: edifact-ls-tnp9
---

# Description

Build a static Go lookup table mapping segment tag -> (name, one-line
description), sourced from the UN Segment Directory (UNSD). Cover at
least the service segments this project already recognizes
(UNA/UNB/UNZ/UNG/UNE/UNH/UNT/UNS) plus the common business segments
already appearing in this project's fixtures (BGM, DTM, NAD, LOC, RFF,
MOA, FTX, TSR, CUX, CNT, GDS). The table is meant to be extended
incrementally, not exhaustive from day one.

# Acceptance Criteria

[x] Lookup table type + data defined, keyed by 3-letter tag
[x] Covers all currently-recognized service segments plus the business
segments listed above
[x] Source cited in a comment for future cross-checking/extension
[x] Unit test confirming a few known tags resolve to the expected text

## Summary of Changes

internal/edifact/segments.go: SegmentInfo{Name, Description} +
segmentDescriptions map + SegmentDescription(tag) (SegmentInfo, bool)
accessor. Covers all 8 recognized service segments plus BGM, DTM, NAD,
LOC, RFF, MOA, FTX, TSR, CUX, CNT, GDS, and CTA (already used in this
project's own IFTMCS fixtures) -- 20 tags total.

Source citation notes something concrete rather than a bare URL: each
segment has its own UNSD definition page at
.../trsd/trsd<tag-lowercase>.htm, and I'd actually seen 15 of these 20
tags' real URLs directly (as hyperlinks) while sourcing IFTMCS's schema
data earlier -- the other 6 (UNA/UNB/UNZ/UNG/UNE/UNS) are envelope
segments that don't appear in IFTMCS's own segment table, so their URLs
follow the identical pattern but weren't individually confirmed the
same way. Flagged that distinction in the comment so a future
cross-check knows which ones to verify first.

3 unit tests: known tags resolve correctly, every currently-recognized
service segment (knownServiceSegmentTags in lint.go) has a description
so the two lists can't silently drift apart, and an unknown tag
correctly reports not-found. Full suite (`make test`) green.
