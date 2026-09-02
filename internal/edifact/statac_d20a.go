package edifact

// STATAC D.20A message specification: the structural (segment/group
// presence, order, cardinality) branching diagram for this message,
// UN/EDIFACT directory release D.20A. Scope is structural only -- not
// element-level content -- per edifact-ls-3uzr's epic non-goal.
// 3 segment groups, max nesting depth 2.
//
// Source: https://service.unece.org/trade/untdid/d20a/trmd/statac_c.htm
// section 4.3.1 "Segment table". That URL currently returns HTTP 403 from
// Cloudflare when fetched directly; this data was transcribed from the
// Wayback Machine's archived copy instead:
// http://web.archive.org/web/20231208132925/https://service.unece.org/trade/untdid/d20a/trmd/statac_c.htm
// -- re-check the direct URL first if cross-checking later, in case the
// block has lifted.
//
// Transcription note: generated in bulk (edifact-ls-13gu) the same way
// as every other message-type schema -- a script parsed the source's
// exact rail-art column positions mechanically and verified the result
// balances (all 3 groups open and close consistently)
// before being transcribed here, rather than reading the ASCII nesting
// by eye. Not individually spot-checked against the raw source the way
// the first several message types in this project were -- see
// edifact-ls-13gu for which representative samples were spot-checked
// for this batch.

func init() {
	RegisterSchema(
		MessageID{Type: "STATAC", Version: "D", Release: "20A", Agency: "UN"},
		statacD20aSchema,
		"https://service.unece.org/trade/untdid/d20a/trmd/statac_c.htm",
	)
}

var statacD20aSchema = Schema{Nodes: []SchemaNode{
	{Segment: "BGM", Mandatory: true, MaxRepeat: 1},  // Beginning of message
	{Segment: "DTM", Mandatory: true, MaxRepeat: 5},  // Date/time/period
	{Segment: "RFF", Mandatory: false, MaxRepeat: 5}, // Reference
	{Segment: "CUX", Mandatory: false, MaxRepeat: 1}, // Currencies
	{ // Segment group 1
		Group: []SchemaNode{
			{Segment: "NAD", Mandatory: true, MaxRepeat: 1}, // Name and address
			{ // Segment group 2
				Group: []SchemaNode{
					{Segment: "CTA", Mandatory: true, MaxRepeat: 1},  // Contact information
					{Segment: "COM", Mandatory: false, MaxRepeat: 5}, // Communication contact
				},
				Mandatory: false, MaxRepeat: 5,
			},
		},
		Mandatory: true, MaxRepeat: 99,
	},
	{ // Segment group 3
		Group: []SchemaNode{
			{Segment: "DOC", Mandatory: true, MaxRepeat: 1},  // Document/message details
			{Segment: "MOA", Mandatory: true, MaxRepeat: 5},  // Monetary amount
			{Segment: "DTM", Mandatory: false, MaxRepeat: 5}, // Date/time/period
			{Segment: "RFF", Mandatory: false, MaxRepeat: 5}, // Reference
		},
		Mandatory: true, MaxRepeat: 200000,
	},
	{Segment: "UNS", Mandatory: true, MaxRepeat: 1},   // Section control
	{Segment: "MOA", Mandatory: true, MaxRepeat: 9},   // Monetary amount
	{Segment: "FTX", Mandatory: false, MaxRepeat: 99}, // Free text
}}
