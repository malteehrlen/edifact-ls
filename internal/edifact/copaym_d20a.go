package edifact

// COPAYM D.20A message specification: the structural (segment/group
// presence, order, cardinality) branching diagram for this message,
// UN/EDIFACT directory release D.20A. Scope is structural only -- not
// element-level content -- per edifact-ls-3uzr's epic non-goal.
// 7 segment groups, max nesting depth 3.
//
// Source: https://service.unece.org/trade/untdid/d20a/trmd/copaym_c.htm
// section 4.3.1 "Segment table". That URL currently returns HTTP 403 from
// Cloudflare when fetched directly; this data was transcribed from the
// Wayback Machine's archived copy instead:
// http://web.archive.org/web/20231208130955/https://service.unece.org/trade/untdid/d20a/trmd/copaym_c.htm
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
		MessageID{Type: "COPAYM", Version: "D", Release: "20A", Agency: "UN"},
		copaymD20aSchema,
		"https://service.unece.org/trade/untdid/d20a/trmd/copaym_c.htm",
	)
}

var copaymD20aSchema = Schema{Nodes: []SchemaNode{
	{Segment: "BGM", Mandatory: true, MaxRepeat: 1},  // Beginning of message
	{Segment: "GEI", Mandatory: true, MaxRepeat: 1},  // Processing information
	{Segment: "DTM", Mandatory: true, MaxRepeat: 9},  // Date/time/period
	{Segment: "QTY", Mandatory: false, MaxRepeat: 1}, // Quantity
	{Segment: "RFF", Mandatory: false, MaxRepeat: 9}, // Reference
	{Segment: "CUX", Mandatory: false, MaxRepeat: 1}, // Currencies
	{ // Segment group 1
		Group: []SchemaNode{
			{Segment: "PNA", Mandatory: true, MaxRepeat: 1},  // Party identification
			{Segment: "ADR", Mandatory: false, MaxRepeat: 1}, // Address
			{Segment: "CTA", Mandatory: false, MaxRepeat: 1}, // Contact information
			{Segment: "COM", Mandatory: false, MaxRepeat: 9}, // Communication contact
			{Segment: "RFF", Mandatory: false, MaxRepeat: 9}, // Reference
			{Segment: "ATT", Mandatory: false, MaxRepeat: 9}, // Attribute
		},
		Mandatory: true, MaxRepeat: 9,
	},
	{Segment: "UNS", Mandatory: true, MaxRepeat: 1}, // Section control
	{ // Segment group 2
		Group: []SchemaNode{
			{Segment: "PNA", Mandatory: true, MaxRepeat: 1},   // Party identification
			{Segment: "ADR", Mandatory: false, MaxRepeat: 9},  // Address
			{Segment: "CTA", Mandatory: false, MaxRepeat: 1},  // Contact information
			{Segment: "COM", Mandatory: false, MaxRepeat: 9},  // Communication contact
			{Segment: "DTM", Mandatory: false, MaxRepeat: 9},  // Date/time/period
			{Segment: "MOA", Mandatory: false, MaxRepeat: 9},  // Monetary amount
			{Segment: "QTY", Mandatory: false, MaxRepeat: 99}, // Quantity
			{Segment: "RFF", Mandatory: false, MaxRepeat: 9},  // Reference
			{Segment: "ATT", Mandatory: false, MaxRepeat: 9},  // Attribute
			{ // Segment group 3
				Group: []SchemaNode{
					{Segment: "COT", Mandatory: true, MaxRepeat: 1},  // Contribution details
					{Segment: "DLI", Mandatory: false, MaxRepeat: 1}, // Document line identification
					{Segment: "DTM", Mandatory: false, MaxRepeat: 1}, // Date/time/period
					{Segment: "PCD", Mandatory: false, MaxRepeat: 9}, // Percentage details
					{Segment: "MOA", Mandatory: false, MaxRepeat: 9}, // Monetary amount
					{Segment: "QTY", Mandatory: false, MaxRepeat: 9}, // Quantity
					{Segment: "RFF", Mandatory: false, MaxRepeat: 9}, // Reference
					{Segment: "ATT", Mandatory: false, MaxRepeat: 9}, // Attribute
				},
				Mandatory: false, MaxRepeat: 999,
			},
			{ // Segment group 4
				Group: []SchemaNode{
					{Segment: "LOC", Mandatory: true, MaxRepeat: 1}, // Place/location identification
					{ // Segment group 5
						Group: []SchemaNode{
							{Segment: "PCD", Mandatory: true, MaxRepeat: 1}, // Percentage details
							{Segment: "MOA", Mandatory: true, MaxRepeat: 9}, // Monetary amount
						},
						Mandatory: true, MaxRepeat: 9,
					},
				},
				Mandatory: false, MaxRepeat: 99,
			},
		},
		Mandatory: true, MaxRepeat: 9999,
	},
	{Segment: "UNS", Mandatory: true, MaxRepeat: 1},  // Section control
	{Segment: "MOA", Mandatory: false, MaxRepeat: 9}, // Monetary amount
	{Segment: "CNT", Mandatory: false, MaxRepeat: 1}, // Control total
	{ // Segment group 6
		Group: []SchemaNode{
			{Segment: "PAI", Mandatory: true, MaxRepeat: 1},  // Payment instructions
			{Segment: "FII", Mandatory: false, MaxRepeat: 1}, // Financial institution information
			{Segment: "MOA", Mandatory: false, MaxRepeat: 9}, // Monetary amount
			{Segment: "RFF", Mandatory: false, MaxRepeat: 9}, // Reference
			{Segment: "DTM", Mandatory: false, MaxRepeat: 9}, // Date/time/period
		},
		Mandatory: false, MaxRepeat: 9,
	},
	{ // Segment group 7
		Group: []SchemaNode{
			{Segment: "AUT", Mandatory: true, MaxRepeat: 1},  // Authentication result
			{Segment: "DTM", Mandatory: false, MaxRepeat: 9}, // Date/time/period
			{Segment: "RFF", Mandatory: false, MaxRepeat: 9}, // Reference
		},
		Mandatory: false, MaxRepeat: 9,
	},
}}
