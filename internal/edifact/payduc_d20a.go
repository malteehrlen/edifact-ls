package edifact

// PAYDUC D.20A message specification: the structural (segment/group
// presence, order, cardinality) branching diagram for this message,
// UN/EDIFACT directory release D.20A. Scope is structural only -- not
// element-level content -- per edifact-ls-3uzr's epic non-goal.
// 7 segment groups, max nesting depth 4.
//
// Source: https://service.unece.org/trade/untdid/d20a/trmd/payduc_c.htm
// section 4.3.1 "Segment table". That URL currently returns HTTP 403 from
// Cloudflare when fetched directly; this data was transcribed from the
// Wayback Machine's archived copy instead:
// http://web.archive.org/web/20231208125057/https://service.unece.org/trade/untdid/d20a/trmd/payduc_c.htm
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
		MessageID{Type: "PAYDUC", Version: "D", Release: "20A", Agency: "UN"},
		payducD20aSchema,
		"https://service.unece.org/trade/untdid/d20a/trmd/payduc_c.htm",
	)
}

var payducD20aSchema = Schema{Nodes: []SchemaNode{
	{Segment: "BGM", Mandatory: true, MaxRepeat: 1},  // Beginning of message
	{Segment: "PAI", Mandatory: true, MaxRepeat: 1},  // Payment instructions
	{Segment: "FII", Mandatory: true, MaxRepeat: 2},  // Financial institution information
	{Segment: "DTM", Mandatory: true, MaxRepeat: 4},  // Date/time/period
	{Segment: "CUX", Mandatory: false, MaxRepeat: 1}, // Currencies
	{Segment: "PYT", Mandatory: false, MaxRepeat: 1}, // Payment terms
	{Segment: "FTX", Mandatory: false, MaxRepeat: 5}, // Free text
	{ // Segment group 1
		Group: []SchemaNode{
			{Segment: "RFF", Mandatory: true, MaxRepeat: 1},  // Reference
			{Segment: "DTM", Mandatory: false, MaxRepeat: 1}, // Date/time/period
		},
		Mandatory: false, MaxRepeat: 99,
	},
	{ // Segment group 2
		Group: []SchemaNode{
			{Segment: "NAD", Mandatory: true, MaxRepeat: 1}, // Name and address
			{ // Segment group 3
				Group: []SchemaNode{
					{Segment: "CTA", Mandatory: true, MaxRepeat: 1},  // Contact information
					{Segment: "COM", Mandatory: false, MaxRepeat: 1}, // Communication contact
				},
				Mandatory: false, MaxRepeat: 5,
			},
		},
		Mandatory: false, MaxRepeat: 6,
	},
	{ // Segment group 4
		Group: []SchemaNode{
			{Segment: "GEI", Mandatory: true, MaxRepeat: 1},  // Processing information
			{Segment: "RFF", Mandatory: true, MaxRepeat: 1},  // Reference
			{Segment: "MOA", Mandatory: true, MaxRepeat: 9},  // Monetary amount
			{Segment: "BUS", Mandatory: false, MaxRepeat: 1}, // Business function
			{Segment: "CUX", Mandatory: false, MaxRepeat: 1}, // Currencies
			{ // Segment group 5
				Group: []SchemaNode{
					{Segment: "DTM", Mandatory: true, MaxRepeat: 1}, // Date/time/period
					{ // Segment group 6
						Group: []SchemaNode{
							{Segment: "UGH", Mandatory: true, MaxRepeat: 1}, // Anti-collision segment group header
							{ // Segment group 7
								Group: []SchemaNode{
									{Segment: "NAD", Mandatory: true, MaxRepeat: 1},  // Name and address
									{Segment: "RFF", Mandatory: false, MaxRepeat: 9}, // Reference
									{Segment: "MOA", Mandatory: true, MaxRepeat: 9},  // Monetary amount
									{Segment: "AJT", Mandatory: false, MaxRepeat: 9}, // Adjustment details
									{Segment: "PYT", Mandatory: false, MaxRepeat: 1}, // Payment terms
									{Segment: "FTX", Mandatory: false, MaxRepeat: 3}, // Free text
									{Segment: "DTM", Mandatory: false, MaxRepeat: 9}, // Date/time/period
								},
								Mandatory: false, MaxRepeat: 999999,
							},
							{Segment: "UGT", Mandatory: true, MaxRepeat: 1}, // Anti-collision segment group trailer
						},
						Mandatory: false, MaxRepeat: 1,
					},
				},
				Mandatory: true, MaxRepeat: 99,
			},
		},
		Mandatory: true, MaxRepeat: 9,
	},
	{Segment: "UNS", Mandatory: true, MaxRepeat: 1},  // Section control
	{Segment: "MOA", Mandatory: true, MaxRepeat: 1},  // Monetary amount
	{Segment: "CNT", Mandatory: false, MaxRepeat: 9}, // Control total
	{Segment: "AUT", Mandatory: false, MaxRepeat: 1}, // Authentication result
}}
