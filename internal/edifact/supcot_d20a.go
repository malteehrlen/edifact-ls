package edifact

// SUPCOT D.20A message specification: the structural (segment/group
// presence, order, cardinality) branching diagram for this message,
// UN/EDIFACT directory release D.20A. Scope is structural only -- not
// element-level content -- per edifact-ls-3uzr's epic non-goal.
// 5 segment groups, max nesting depth 3.
//
// Source: https://service.unece.org/trade/untdid/d20a/trmd/supcot_c.htm
// section 4.3.1 "Segment table". That URL currently returns HTTP 403 from
// Cloudflare when fetched directly; this data was transcribed from the
// Wayback Machine's archived copy instead:
// http://web.archive.org/web/20231201133357/https://service.unece.org/trade/untdid/d20a/trmd/supcot_c.htm
// -- re-check the direct URL first if cross-checking later, in case the
// block has lifted.
//
// Transcription note: generated in bulk (edifact-ls-13gu) the same way
// as every other message-type schema -- a script parsed the source's
// exact rail-art column positions mechanically and verified the result
// balances (all 5 groups open and close consistently)
// before being transcribed here, rather than reading the ASCII nesting
// by eye. Not individually spot-checked against the raw source the way
// the first several message types in this project were -- see
// edifact-ls-13gu for which representative samples were spot-checked
// for this batch.

func init() {
	RegisterSchema(
		MessageID{Type: "SUPCOT", Version: "D", Release: "20A", Agency: "UN"},
		supcotD20aSchema,
		"https://service.unece.org/trade/untdid/d20a/trmd/supcot_c.htm",
	)
}

var supcotD20aSchema = Schema{Nodes: []SchemaNode{
	{Segment: "BGM", Mandatory: true, MaxRepeat: 1},  // Beginning of message
	{Segment: "RFF", Mandatory: true, MaxRepeat: 6},  // Reference
	{Segment: "PAI", Mandatory: false, MaxRepeat: 1}, // Payment instructions
	{Segment: "CUX", Mandatory: false, MaxRepeat: 1}, // Currencies
	{Segment: "FII", Mandatory: false, MaxRepeat: 2}, // Financial institution information
	{Segment: "DTM", Mandatory: false, MaxRepeat: 9}, // Date/time/period
	{Segment: "PYT", Mandatory: false, MaxRepeat: 1}, // Payment terms
	{Segment: "FTX", Mandatory: false, MaxRepeat: 5}, // Free text
	{ // Segment group 1
		Group: []SchemaNode{
			{Segment: "NAD", Mandatory: true, MaxRepeat: 1}, // Name and address
			{ // Segment group 2
				Group: []SchemaNode{
					{Segment: "CTA", Mandatory: true, MaxRepeat: 1},  // Contact information
					{Segment: "COM", Mandatory: false, MaxRepeat: 1}, // Communication contact
				},
				Mandatory: false, MaxRepeat: 5,
			},
		},
		Mandatory: false, MaxRepeat: 6,
	},
	{Segment: "UNS", Mandatory: true, MaxRepeat: 1}, // Section control
	{ // Segment group 3
		Group: []SchemaNode{
			{Segment: "NAD", Mandatory: true, MaxRepeat: 1},  // Name and address
			{Segment: "ATT", Mandatory: false, MaxRepeat: 1}, // Attribute
			{Segment: "RFF", Mandatory: false, MaxRepeat: 1}, // Reference
			{Segment: "DTM", Mandatory: false, MaxRepeat: 1}, // Date/time/period
			{Segment: "EMP", Mandatory: false, MaxRepeat: 1}, // Employment details
			{Segment: "LOC", Mandatory: false, MaxRepeat: 1}, // Place/location identification
			{Segment: "FTX", Mandatory: false, MaxRepeat: 3}, // Free text
			{ // Segment group 4
				Group: []SchemaNode{
					{Segment: "PYT", Mandatory: true, MaxRepeat: 1},  // Payment terms
					{Segment: "MOA", Mandatory: false, MaxRepeat: 1}, // Monetary amount
					{Segment: "CUX", Mandatory: false, MaxRepeat: 1}, // Currencies
					{Segment: "DTM", Mandatory: false, MaxRepeat: 3}, // Date/time/period
					{Segment: "FTX", Mandatory: false, MaxRepeat: 3}, // Free text
					{ // Segment group 5
						Group: []SchemaNode{
							{Segment: "COT", Mandatory: true, MaxRepeat: 1},  // Contribution details
							{Segment: "MOA", Mandatory: true, MaxRepeat: 1},  // Monetary amount
							{Segment: "DTM", Mandatory: false, MaxRepeat: 3}, // Date/time/period
							{Segment: "FTX", Mandatory: false, MaxRepeat: 3}, // Free text
						},
						Mandatory: true, MaxRepeat: 99,
					},
				},
				Mandatory: true, MaxRepeat: 9,
			},
		},
		Mandatory: true, MaxRepeat: 999999,
	},
	{Segment: "UNS", Mandatory: true, MaxRepeat: 1},  // Section control
	{Segment: "MOA", Mandatory: true, MaxRepeat: 1},  // Monetary amount
	{Segment: "CNT", Mandatory: false, MaxRepeat: 9}, // Control total
	{Segment: "AUT", Mandatory: false, MaxRepeat: 1}, // Authentication result
}}
