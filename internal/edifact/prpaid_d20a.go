package edifact

// PRPAID D.20A message specification: the structural (segment/group
// presence, order, cardinality) branching diagram for this message,
// UN/EDIFACT directory release D.20A. Scope is structural only -- not
// element-level content -- per edifact-ls-3uzr's epic non-goal.
// 7 segment groups, max nesting depth 2.
//
// Source: https://service.unece.org/trade/untdid/d20a/trmd/prpaid_c.htm
// section 4.3.1 "Segment table". That URL currently returns HTTP 403 from
// Cloudflare when fetched directly; this data was transcribed from the
// Wayback Machine's archived copy instead:
// http://web.archive.org/web/20231208130650/https://service.unece.org/trade/untdid/d20a/trmd/prpaid_c.htm
// -- re-check the direct URL first if cross-checking later, in case the
// block has lifted.
//
// Transcription note: generated in bulk (edifact-ls-13gu) the same way
// as every other message-type schema -- a script parsed the source's
// exact rail-art column positions mechanically and verified the result
// balances (all 7 groups open and close consistently)
// before being transcribed here, rather than reading the ASCII nesting
// by eye. Not individually spot-checked against the raw source the way
// the first several message types in this project were -- see
// edifact-ls-13gu for which representative samples were spot-checked
// for this batch.

func init() {
	RegisterSchema(
		MessageID{Type: "PRPAID", Version: "D", Release: "20A", Agency: "UN"},
		prpaidD20aSchema,
		"https://service.unece.org/trade/untdid/d20a/trmd/prpaid_c.htm",
	)
}

var prpaidD20aSchema = Schema{Nodes: []SchemaNode{
	{Segment: "BGM", Mandatory: false, MaxRepeat: 1}, // Beginning of message
	{Segment: "GEI", Mandatory: true, MaxRepeat: 1},  // Processing information
	{Segment: "RFF", Mandatory: true, MaxRepeat: 9},  // Reference
	{Segment: "DTM", Mandatory: true, MaxRepeat: 9},  // Date/time/period
	{Segment: "ICD", Mandatory: false, MaxRepeat: 1}, // Insurance cover description
	{ // Segment group 1
		Group: []SchemaNode{
			{Segment: "NAD", Mandatory: true, MaxRepeat: 1},  // Name and address
			{Segment: "CTA", Mandatory: false, MaxRepeat: 1}, // Contact information
			{Segment: "COM", Mandatory: false, MaxRepeat: 3}, // Communication contact
			{Segment: "RFF", Mandatory: false, MaxRepeat: 9}, // Reference
		},
		Mandatory: true, MaxRepeat: 9,
	},
	{ // Segment group 2
		Group: []SchemaNode{
			{Segment: "DOC", Mandatory: true, MaxRepeat: 1}, // Document/message details
			{ // Segment group 3
				Group: []SchemaNode{
					{Segment: "GEI", Mandatory: true, MaxRepeat: 1},  // Processing information
					{Segment: "FTX", Mandatory: false, MaxRepeat: 1}, // Free text
				},
				Mandatory: false, MaxRepeat: 1,
			},
			{Segment: "DTM", Mandatory: false, MaxRepeat: 9}, // Date/time/period
			{Segment: "NAD", Mandatory: false, MaxRepeat: 9}, // Name and address
			{ // Segment group 4
				Group: []SchemaNode{
					{Segment: "MOA", Mandatory: true, MaxRepeat: 1},  // Monetary amount
					{Segment: "DTM", Mandatory: false, MaxRepeat: 1}, // Date/time/period
					{Segment: "PCD", Mandatory: false, MaxRepeat: 1}, // Percentage details
				},
				Mandatory: false, MaxRepeat: 99,
			},
			{Segment: "RFF", Mandatory: true, MaxRepeat: 9}, // Reference
			{ // Segment group 5
				Group: []SchemaNode{
					{Segment: "ICD", Mandatory: true, MaxRepeat: 1},   // Insurance cover description
					{Segment: "MOA", Mandatory: false, MaxRepeat: 99}, // Monetary amount
					{Segment: "RFF", Mandatory: false, MaxRepeat: 9},  // Reference
					{Segment: "PCD", Mandatory: false, MaxRepeat: 1},  // Percentage details
				},
				Mandatory: false, MaxRepeat: 99,
			},
			{ // Segment group 6
				Group: []SchemaNode{
					{Segment: "ATT", Mandatory: true, MaxRepeat: 1},  // Attribute
					{Segment: "PCD", Mandatory: false, MaxRepeat: 1}, // Percentage details
				},
				Mandatory: false, MaxRepeat: 9,
			},
			{Segment: "CUX", Mandatory: false, MaxRepeat: 1}, // Currencies
		},
		Mandatory: false, MaxRepeat: 9999,
	},
	{Segment: "UNS", Mandatory: true, MaxRepeat: 1}, // Section control
	{ // Segment group 7
		Group: []SchemaNode{
			{Segment: "MOA", Mandatory: true, MaxRepeat: 1},  // Monetary amount
			{Segment: "GEI", Mandatory: false, MaxRepeat: 1}, // Processing information
			{Segment: "DTM", Mandatory: false, MaxRepeat: 1}, // Date/time/period
			{Segment: "RFF", Mandatory: false, MaxRepeat: 1}, // Reference
			{Segment: "ATT", Mandatory: false, MaxRepeat: 1}, // Attribute
			{Segment: "ICD", Mandatory: false, MaxRepeat: 1}, // Insurance cover description
			{Segment: "CUX", Mandatory: false, MaxRepeat: 1}, // Currencies
		},
		Mandatory: false, MaxRepeat: 99,
	},
}}
