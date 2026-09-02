package edifact

import (
	"fmt"
	"strings"
)

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
// ValidateEnvelopes checks. src is the original interchange text; needed
// (only) to let the redundant-UNA Fix consume a trailing line terminator
// so removing it doesn't leave a blank line behind -- see trailingNewline.
func Lint(ic *Interchange, src string) ErrorList {
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

	if ic.UNA != nil && hasFunctionalDefaultDelimiters(ic.Delimiters) {
		// UNA can only ever appear at the very start of src (detectUNA
		// requires src[:3] == "UNA"), so there's never a preceding
		// newline to worry about -- only a trailing one.
		end := ic.UNA.Pos.Offset + len(ic.UNA.Raw)
		oldText := ic.UNA.Raw + trailingNewline(src, end)
		errs.AddFixable(ic.UNA.Pos, SeverityInfo, "redundant-una", &Fix{
			Title:   "Remove redundant UNA service string advice",
			Pos:     ic.UNA.Pos,
			OldText: oldText,
			NewText: "",
		}, "UNA service string advice defines exactly the default component/element/release/terminator delimiters%s; it can be safely omitted", decimalMarkNote(ic.Delimiters.Decimal))
	}

	return errs
}

// trailingNewline returns the line terminator ("\r\n", "\n", or bare "\r")
// immediately following byte offset end in src, or "" if there isn't one.
// Real-world interchanges commonly put each segment on its own line purely
// for readability (see lexer.go's skipInterSegmentWhitespace); a Fix that
// deletes a whole such segment needs to consume that line's terminator
// too, or the deletion leaves a blank line behind.
func trailingNewline(src string, end int) string {
	if end >= len(src) {
		return ""
	}
	if src[end] == '\r' {
		if end+1 < len(src) && src[end+1] == '\n' {
			return "\r\n"
		}
		return "\r"
	}
	if src[end] == '\n' {
		return "\n"
	}
	return ""
}

// hasFunctionalDefaultDelimiters reports whether d matches the documented
// defaults on the four delimiters this parser actually treats as
// structurally significant (component, element, release, terminator).
// The decimal mark and reserved character are deliberately excluded: this
// parser never uses either to affect parsing, and per Wikipedia's EDIFACT
// article, ISO 9735 itself has been inconsistent about the decimal mark
// (versions 1-3 default to a comma; version 4 states the position is to be
// ignored entirely, with comma and dot usable interchangeably in numeric
// data) -- so treating it as part of "the" default would be asserting a
// precision the standard doesn't actually have.
func hasFunctionalDefaultDelimiters(d Delimiters) bool {
	def := DefaultDelimiters()
	return d.Component == def.Component &&
		d.Element == def.Element &&
		d.Release == def.Release &&
		d.Terminator == def.Terminator
}

// decimalMarkNote describes a UNA's decimal mark character in terms of
// which ISO 9735 convention it matches, per the version history above.
func decimalMarkNote(decimal byte) string {
	switch decimal {
	case ',':
		return " (using a comma decimal mark, the ISO 9735 version 1-3 default)"
	case '.':
		return " (using a dot decimal mark, the common convention since version 4, which says the decimal mark position is not significant)"
	default:
		return fmt.Sprintf(" (using %q as the decimal mark, itself not structurally significant per ISO 9735 version 4)", string(rune(decimal)))
	}
}
