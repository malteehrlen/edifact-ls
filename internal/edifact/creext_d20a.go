package edifact

// CREEXT D.20A message specification: the structural (segment/group
// presence, order, cardinality) branching diagram for this message,
// UN/EDIFACT directory release D.20A. Scope is structural only -- not
// element-level content -- per edifact-ls-3uzr's epic non-goal.
// 18 segment groups, max nesting depth 4.
//
// Source: https://service.unece.org/trade/untdid/d20a/trmd/creext_c.htm
// section 4.3.1 "Segment table". That URL currently returns HTTP 403 from
// Cloudflare when fetched directly; this data was transcribed from the
// Wayback Machine's archived copy instead:
// http://web.archive.org/web/20230608023249/https://service.unece.org/trade/untdid/d20a/trmd/creext_c.htm
// -- re-check the direct URL first if cross-checking later, in case the
// block has lifted.
//
// Transcription note: generated in bulk (edifact-ls-13gu) the same way
// as every other message-type schema -- a script parsed the source's
// exact rail-art column positions mechanically and verified the result
// balances (all 18 groups open and close consistently)
// before being transcribed here, rather than reading the ASCII nesting
// by eye. Not individually spot-checked against the raw source the way
// the first several message types in this project were -- see
// edifact-ls-13gu for which representative samples were spot-checked
// for this batch.

func init() {
	RegisterSchema(
		MessageID{Type: "CREEXT", Version: "D", Release: "20A", Agency: "UN"},
		creextD20aSchema,
		"https://service.unece.org/trade/untdid/d20a/trmd/creext_c.htm",
	)
}

var creextD20aSchema = Schema{Nodes: []SchemaNode{
	{Segment: "BGM", Mandatory: true, MaxRepeat: 1},  // Beginning of message
	{Segment: "BUS", Mandatory: false, MaxRepeat: 1}, // Business function
	{Segment: "DTM", Mandatory: true, MaxRepeat: 4},  // Date/time/period
	{ // Segment group 1
		Group: []SchemaNode{
			{Segment: "RFF", Mandatory: true, MaxRepeat: 1},  // Reference
			{Segment: "DTM", Mandatory: false, MaxRepeat: 1}, // Date/time/period
		},
		Mandatory: false, MaxRepeat: 5,
	},
	{ // Segment group 2
		Group: []SchemaNode{
			{Segment: "MOA", Mandatory: true, MaxRepeat: 1},  // Monetary amount
			{Segment: "CUX", Mandatory: false, MaxRepeat: 1}, // Currencies
			{Segment: "DTM", Mandatory: false, MaxRepeat: 2}, // Date/time/period
			{Segment: "RFF", Mandatory: false, MaxRepeat: 1}, // Reference
		},
		Mandatory: true, MaxRepeat: 4,
	},
	{ // Segment group 3
		Group: []SchemaNode{
			{Segment: "FII", Mandatory: true, MaxRepeat: 1},  // Financial institution information
			{Segment: "CTA", Mandatory: false, MaxRepeat: 1}, // Contact information
			{Segment: "COM", Mandatory: false, MaxRepeat: 5}, // Communication contact
		},
		Mandatory: false, MaxRepeat: 4,
	},
	{ // Segment group 4
		Group: []SchemaNode{
			{Segment: "NAD", Mandatory: true, MaxRepeat: 1},  // Name and address
			{Segment: "CTA", Mandatory: false, MaxRepeat: 1}, // Contact information
			{Segment: "COM", Mandatory: false, MaxRepeat: 5}, // Communication contact
		},
		Mandatory: false, MaxRepeat: 6,
	},
	{ // Segment group 5
		Group: []SchemaNode{
			{Segment: "INP", Mandatory: true, MaxRepeat: 1},  // Parties and instruction
			{Segment: "FTX", Mandatory: false, MaxRepeat: 1}, // Free text
			{Segment: "DTM", Mandatory: false, MaxRepeat: 3}, // Date/time/period
		},
		Mandatory: false, MaxRepeat: 4,
	},
	{ // Segment group 6
		Group: []SchemaNode{
			{Segment: "GEI", Mandatory: true, MaxRepeat: 1},   // Processing information
			{Segment: "MOA", Mandatory: false, MaxRepeat: 1},  // Monetary amount
			{Segment: "LOC", Mandatory: false, MaxRepeat: 2},  // Place/location identification
			{Segment: "NAD", Mandatory: false, MaxRepeat: 1},  // Name and address
			{Segment: "RCS", Mandatory: false, MaxRepeat: 1},  // Requirements and conditions
			{Segment: "FTX", Mandatory: false, MaxRepeat: 10}, // Free text
		},
		Mandatory: false, MaxRepeat: 10,
	},
	{ // Segment group 7
		Group: []SchemaNode{
			{Segment: "FCA", Mandatory: true, MaxRepeat: 1},  // Financial charges allocation
			{Segment: "MOA", Mandatory: false, MaxRepeat: 2}, // Monetary amount
			{ // Segment group 8
				Group: []SchemaNode{
					{Segment: "ALC", Mandatory: true, MaxRepeat: 1},  // Allowance or charge
					{Segment: "PCD", Mandatory: false, MaxRepeat: 1}, // Percentage details
					{Segment: "MOA", Mandatory: false, MaxRepeat: 2}, // Monetary amount
					{Segment: "CUX", Mandatory: false, MaxRepeat: 1}, // Currencies
					{Segment: "DTM", Mandatory: false, MaxRepeat: 2}, // Date/time/period
					{ // Segment group 9
						Group: []SchemaNode{
							{Segment: "TAX", Mandatory: true, MaxRepeat: 1},  // Duty/tax/fee details
							{Segment: "MOA", Mandatory: false, MaxRepeat: 2}, // Monetary amount
							{Segment: "CUX", Mandatory: false, MaxRepeat: 1}, // Currencies
							{Segment: "DTM", Mandatory: false, MaxRepeat: 1}, // Date/time/period
						},
						Mandatory: false, MaxRepeat: 5,
					},
				},
				Mandatory: true, MaxRepeat: 20,
			},
		},
		Mandatory: false, MaxRepeat: 5,
	},
	{ // Segment group 10
		Group: []SchemaNode{
			{Segment: "PRC", Mandatory: true, MaxRepeat: 1},  // Process identification
			{Segment: "FTX", Mandatory: false, MaxRepeat: 5}, // Free text
			{ // Segment group 11
				Group: []SchemaNode{
					{Segment: "DOC", Mandatory: true, MaxRepeat: 1},  // Document/message details
					{Segment: "MOA", Mandatory: false, MaxRepeat: 5}, // Monetary amount
					{Segment: "DTM", Mandatory: false, MaxRepeat: 5}, // Date/time/period
					{Segment: "RFF", Mandatory: false, MaxRepeat: 5}, // Reference
					{Segment: "NAD", Mandatory: false, MaxRepeat: 2}, // Name and address
					{ // Segment group 12
						Group: []SchemaNode{
							{Segment: "CUX", Mandatory: true, MaxRepeat: 1},  // Currencies
							{Segment: "DTM", Mandatory: false, MaxRepeat: 1}, // Date/time/period
						},
						Mandatory: false, MaxRepeat: 5,
					},
					{ // Segment group 13
						Group: []SchemaNode{
							{Segment: "AJT", Mandatory: true, MaxRepeat: 1},  // Adjustment details
							{Segment: "MOA", Mandatory: false, MaxRepeat: 1}, // Monetary amount
							{Segment: "RFF", Mandatory: false, MaxRepeat: 1}, // Reference
							{Segment: "FTX", Mandatory: false, MaxRepeat: 5}, // Free text
						},
						Mandatory: false, MaxRepeat: 100,
					},
					{ // Segment group 14
						Group: []SchemaNode{
							{Segment: "DLI", Mandatory: true, MaxRepeat: 1},  // Document line identification
							{Segment: "MOA", Mandatory: false, MaxRepeat: 5}, // Monetary amount
							{Segment: "PIA", Mandatory: false, MaxRepeat: 5}, // Additional product id
							{Segment: "DTM", Mandatory: false, MaxRepeat: 5}, // Date/time/period
							{ // Segment group 15
								Group: []SchemaNode{
									{Segment: "CUX", Mandatory: true, MaxRepeat: 1},  // Currencies
									{Segment: "DTM", Mandatory: false, MaxRepeat: 1}, // Date/time/period
								},
								Mandatory: false, MaxRepeat: 5,
							},
							{ // Segment group 16
								Group: []SchemaNode{
									{Segment: "AJT", Mandatory: true, MaxRepeat: 1},  // Adjustment details
									{Segment: "MOA", Mandatory: false, MaxRepeat: 1}, // Monetary amount
									{Segment: "RFF", Mandatory: false, MaxRepeat: 1}, // Reference
									{Segment: "FTX", Mandatory: false, MaxRepeat: 5}, // Free text
								},
								Mandatory: false, MaxRepeat: 10,
							},
						},
						Mandatory: false, MaxRepeat: 9999,
					},
				},
				Mandatory: false, MaxRepeat: 9999,
			},
			{ // Segment group 17
				Group: []SchemaNode{
					{Segment: "GEI", Mandatory: true, MaxRepeat: 1},  // Processing information
					{Segment: "MOA", Mandatory: false, MaxRepeat: 5}, // Monetary amount
				},
				Mandatory: false, MaxRepeat: 1,
			},
		},
		Mandatory: false, MaxRepeat: 1,
	},
	{ // Segment group 18
		Group: []SchemaNode{
			{Segment: "AUT", Mandatory: true, MaxRepeat: 1},  // Authentication result
			{Segment: "DTM", Mandatory: false, MaxRepeat: 1}, // Date/time/period
		},
		Mandatory: false, MaxRepeat: 5,
	},
}}
