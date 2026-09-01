package edifact

import "strings"

// knownServiceSegmentTags are the syntax-level service segments this
// package recognizes. Per section 6.2 of
// https://unece.org/DAM/trade/untdid/texts/d423.htm: "All service segment
// tags start with the two letters 'UN' which are reserved for this
// purpose... User data segments must not be created with the first two
// letters of the tag 'UN'."
var knownServiceSegmentTags = map[string]bool{
	"UNA": true,
	"UNB": true,
	"UNZ": true,
	"UNG": true,
	"UNE": true,
	"UNH": true,
	"UNT": true,
	"UNS": true,
}

// Lint runs advisory checks over an already-parsed Interchange, producing
// warning/info-severity diagnostics for things that are syntactically
// valid but questionable, as opposed to the structural correctness
// ValidateEnvelopes checks.
func Lint(ic *Interchange) ErrorList {
	var errs ErrorList

	for _, seg := range ic.Segments {
		// Only lint tags that are themselves well-formed (3 uppercase
		// letters) -- an already-malformed tag gets its own syntax error
		// from Parse, and flagging it again here would just be noise.
		if segmentTagPattern.MatchString(seg.Tag) &&
			strings.HasPrefix(seg.Tag, "UN") && !knownServiceSegmentTags[seg.Tag] {
			errs.Add(seg.TagPos, SeverityWarning, "segment tag %q starts with the reserved \"UN\" prefix but isn't a recognized service segment; this prefix is reserved for service segments", seg.Tag)
		}
	}

	if ic.UNA != nil && ic.Delimiters == DefaultDelimiters() {
		errs.Add(ic.UNA.Pos, SeverityInfo, "UNA service string advice defines exactly the default delimiters; it can be safely omitted")
	}

	return errs
}
