package edifact

// RECECO D.20A message specification: the structural (segment/group
// presence, order, cardinality) branching diagram for this message,
// UN/EDIFACT directory release D.20A. Scope is structural only -- not
// element-level content -- per edifact-ls-3uzr's epic non-goal.
// 2 segment groups, max nesting depth 1.
//
// Source: https://service.unece.org/trade/untdid/d20a/trmd/receco_c.htm
// section 4.3.1 "Segment table". That URL currently returns HTTP 403 from
// Cloudflare when fetched directly; this data was transcribed from the
// Wayback Machine's archived copy instead:
// http://web.archive.org/web/20231208120637/https://service.unece.org/trade/untdid/d20a/trmd/receco_c.htm
// -- re-check the direct URL first if cross-checking later, in case the
// block has lifted.
//
// Transcription note: generated in bulk (edifact-ls-13gu) the same way
// as every other message-type schema -- a script parsed the source's
// exact rail-art column positions mechanically and verified the result
// balances (all 2 groups open and close consistently)
// before being transcribed here, rather than reading the ASCII nesting
// by eye. Not individually spot-checked against the raw source the way
// the first several message types in this project were -- see
// edifact-ls-13gu for which representative samples were spot-checked
// for this batch.

func init() {
	RegisterSchema(
		MessageID{Type: "RECECO", Version: "D", Release: "20A", Agency: "UN"},
		rececoD20aSchema,
		"https://service.unece.org/trade/untdid/d20a/trmd/receco_c.htm",
	)
}

var rececoD20aSchema = Schema{Nodes: []SchemaNode{
	{Segment: "BGM", Mandatory: true, MaxRepeat: 1},  // Beginning of message
	{Segment: "RFF", Mandatory: false, MaxRepeat: 9}, // Reference
	{Segment: "DTM", Mandatory: false, MaxRepeat: 9}, // Date/time/period
	{ // Segment group 1
		Group: []SchemaNode{
			{Segment: "NAD", Mandatory: true, MaxRepeat: 1},  // Name and address
			{Segment: "CTA", Mandatory: false, MaxRepeat: 1}, // Contact information
			{Segment: "COM", Mandatory: false, MaxRepeat: 5}, // Communication contact
			{Segment: "FII", Mandatory: false, MaxRepeat: 1}, // Financial institution information
		},
		Mandatory: false, MaxRepeat: 99,
	},
	{ // Segment group 2
		Group: []SchemaNode{
			{Segment: "DOC", Mandatory: true, MaxRepeat: 1},  // Document/message details
			{Segment: "MOA", Mandatory: false, MaxRepeat: 1}, // Monetary amount
			{Segment: "CUX", Mandatory: false, MaxRepeat: 1}, // Currencies
			{Segment: "DTM", Mandatory: false, MaxRepeat: 5}, // Date/time/period
			{Segment: "CCD", Mandatory: false, MaxRepeat: 1}, // Credit cover details
			{Segment: "PCD", Mandatory: false, MaxRepeat: 1}, // Percentage details
			{Segment: "PYT", Mandatory: false, MaxRepeat: 1}, // Payment terms
			{Segment: "FTX", Mandatory: false, MaxRepeat: 1}, // Free text
			{Segment: "RFF", Mandatory: false, MaxRepeat: 1}, // Reference
		},
		Mandatory: false, MaxRepeat: 99,
	},
}}
