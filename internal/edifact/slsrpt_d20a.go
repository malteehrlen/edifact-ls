package edifact

// SLSRPT D.20A message specification: the structural (segment/group
// presence, order, cardinality) branching diagram for this message,
// UN/EDIFACT directory release D.20A. Scope is structural only -- not
// element-level content -- per edifact-ls-3uzr's epic non-goal.
// 14 segment groups, max nesting depth 4.
//
// Source: https://service.unece.org/trade/untdid/d20a/trmd/slsrpt_c.htm
// section 4.3.1 "Segment table". That URL currently returns HTTP 403 from
// Cloudflare when fetched directly; this data was transcribed from the
// Wayback Machine's archived copy instead:
// http://web.archive.org/web/20231002204827/https://service.unece.org/trade/untdid/d20a/trmd/slsrpt_c.htm
// -- re-check the direct URL first if cross-checking later, in case the
// block has lifted.
//
// Transcription note: generated in bulk (edifact-ls-13gu) the same way
// as every other message-type schema -- a script parsed the source's
// exact rail-art column positions mechanically and verified the result
// balances (all 14 groups open and close consistently)
// before being transcribed here, rather than reading the ASCII nesting
// by eye. Not individually spot-checked against the raw source the way
// the first several message types in this project were -- see
// edifact-ls-13gu for which representative samples were spot-checked
// for this batch.

func init() {
	RegisterSchema(
		MessageID{Type: "SLSRPT", Version: "D", Release: "20A", Agency: "UN"},
		slsrptD20aSchema,
		"https://service.unece.org/trade/untdid/d20a/trmd/slsrpt_c.htm",
	)
}

var slsrptD20aSchema = Schema{Nodes: []SchemaNode{
	{Segment: "BGM", Mandatory: true, MaxRepeat: 1},  // Beginning of message
	{Segment: "DTM", Mandatory: true, MaxRepeat: 5},  // Date/time/period
	{Segment: "MKS", Mandatory: false, MaxRepeat: 1}, // Market/sales channel information
	{ // Segment group 1
		Group: []SchemaNode{
			{Segment: "RFF", Mandatory: true, MaxRepeat: 1},  // Reference
			{Segment: "DTM", Mandatory: false, MaxRepeat: 9}, // Date/time/period
		},
		Mandatory: false, MaxRepeat: 9,
	},
	{ // Segment group 2
		Group: []SchemaNode{
			{Segment: "NAD", Mandatory: true, MaxRepeat: 1}, // Name and address
			{ // Segment group 3
				Group: []SchemaNode{
					{Segment: "CTA", Mandatory: true, MaxRepeat: 1},  // Contact information
					{Segment: "COM", Mandatory: false, MaxRepeat: 5}, // Communication contact
				},
				Mandatory: false, MaxRepeat: 5,
			},
			{ // Segment group 4
				Group: []SchemaNode{
					{Segment: "RFF", Mandatory: true, MaxRepeat: 1},  // Reference
					{Segment: "DTM", Mandatory: false, MaxRepeat: 9}, // Date/time/period
				},
				Mandatory: false, MaxRepeat: 9,
			},
		},
		Mandatory: true, MaxRepeat: 5,
	},
	{ // Segment group 5
		Group: []SchemaNode{
			{Segment: "CUX", Mandatory: true, MaxRepeat: 1},  // Currencies
			{Segment: "DTM", Mandatory: false, MaxRepeat: 5}, // Date/time/period
			{Segment: "STS", Mandatory: false, MaxRepeat: 5}, // Status
		},
		Mandatory: false, MaxRepeat: 5,
	},
	{ // Segment group 6
		Group: []SchemaNode{
			{Segment: "LOC", Mandatory: true, MaxRepeat: 1},  // Place/location identification
			{Segment: "DTM", Mandatory: false, MaxRepeat: 5}, // Date/time/period
			{Segment: "FTX", Mandatory: false, MaxRepeat: 5}, // Free text
			{ // Segment group 7
				Group: []SchemaNode{
					{Segment: "RFF", Mandatory: true, MaxRepeat: 1},  // Reference
					{Segment: "DTM", Mandatory: false, MaxRepeat: 5}, // Date/time/period
				},
				Mandatory: false, MaxRepeat: 99,
			},
			{ // Segment group 8
				Group: []SchemaNode{
					{Segment: "LIN", Mandatory: true, MaxRepeat: 1},     // Line item
					{Segment: "PIA", Mandatory: false, MaxRepeat: 5},    // Additional product id
					{Segment: "IMD", Mandatory: false, MaxRepeat: 5},    // Item description
					{Segment: "DTM", Mandatory: false, MaxRepeat: 9},    // Date/time/period
					{Segment: "PAC", Mandatory: false, MaxRepeat: 5},    // Package
					{Segment: "RFF", Mandatory: false, MaxRepeat: 5},    // Reference
					{Segment: "DOC", Mandatory: false, MaxRepeat: 5},    // Document/message details
					{Segment: "ALI", Mandatory: false, MaxRepeat: 5},    // Additional information
					{Segment: "MOA", Mandatory: false, MaxRepeat: 5},    // Monetary amount
					{Segment: "GIN", Mandatory: false, MaxRepeat: 9999}, // Goods identity number
					{ // Segment group 9
						Group: []SchemaNode{
							{Segment: "PRI", Mandatory: true, MaxRepeat: 1},  // Price details
							{Segment: "CUX", Mandatory: false, MaxRepeat: 1}, // Currencies
						},
						Mandatory: false, MaxRepeat: 999,
					},
					{Segment: "PAI", Mandatory: false, MaxRepeat: 9}, // Payment instructions
					{ // Segment group 10
						Group: []SchemaNode{
							{Segment: "QTY", Mandatory: true, MaxRepeat: 1},  // Quantity
							{Segment: "MKS", Mandatory: false, MaxRepeat: 1}, // Market/sales channel information
							{Segment: "NAD", Mandatory: false, MaxRepeat: 1}, // Name and address
							{Segment: "PRI", Mandatory: false, MaxRepeat: 9}, // Price details
							{Segment: "DTM", Mandatory: false, MaxRepeat: 9}, // Date/time/period
							{Segment: "RFF", Mandatory: false, MaxRepeat: 9}, // Reference
							{Segment: "STS", Mandatory: false, MaxRepeat: 9}, // Status
						},
						Mandatory: false, MaxRepeat: 999,
					},
				},
				Mandatory: false, MaxRepeat: 200000,
			},
		},
		Mandatory: false, MaxRepeat: 200000,
	},
	{ // Segment group 11
		Group: []SchemaNode{
			{Segment: "GEI", Mandatory: true, MaxRepeat: 1}, // Processing information
			{ // Segment group 12
				Group: []SchemaNode{
					{Segment: "LIN", Mandatory: true, MaxRepeat: 1},   // Line item
					{Segment: "PIA", Mandatory: false, MaxRepeat: 99}, // Additional product id
					{Segment: "IMD", Mandatory: false, MaxRepeat: 9},  // Item description
					{Segment: "PAC", Mandatory: false, MaxRepeat: 9},  // Package
					{Segment: "MOA", Mandatory: false, MaxRepeat: 9},  // Monetary amount
					{ // Segment group 13
						Group: []SchemaNode{
							{Segment: "LOC", Mandatory: true, MaxRepeat: 1},  // Place/location identification
							{Segment: "RFF", Mandatory: false, MaxRepeat: 9}, // Reference
							{Segment: "DTM", Mandatory: false, MaxRepeat: 9}, // Date/time/period
							{Segment: "MOA", Mandatory: false, MaxRepeat: 9}, // Monetary amount
							{ // Segment group 14
								Group: []SchemaNode{
									{Segment: "QTY", Mandatory: true, MaxRepeat: 1},  // Quantity
									{Segment: "PRI", Mandatory: false, MaxRepeat: 9}, // Price details
									{Segment: "DTM", Mandatory: false, MaxRepeat: 9}, // Date/time/period
									{Segment: "RFF", Mandatory: false, MaxRepeat: 9}, // Reference
									{Segment: "STS", Mandatory: false, MaxRepeat: 9}, // Status
								},
								Mandatory: false, MaxRepeat: 99,
							},
						},
						Mandatory: false, MaxRepeat: 9999,
					},
				},
				Mandatory: false, MaxRepeat: 999999,
			},
		},
		Mandatory: false, MaxRepeat: 1,
	},
	{Segment: "UNS", Mandatory: true, MaxRepeat: 1},  // Section control
	{Segment: "MOA", Mandatory: false, MaxRepeat: 9}, // Monetary amount
}}
