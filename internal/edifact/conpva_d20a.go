package edifact

// CONPVA D.20A message specification: the structural (segment/group
// presence, order, cardinality) branching diagram for this message,
// UN/EDIFACT directory release D.20A. Scope is structural only -- not
// element-level content -- per edifact-ls-3uzr's epic non-goal.
// 34 segment groups, max nesting depth 4.
//
// Source: https://service.unece.org/trade/untdid/d20a/trmd/conpva_c.htm
// section 4.3.1 "Segment table". That URL currently returns HTTP 403 from
// Cloudflare when fetched directly; this data was transcribed from the
// Wayback Machine's archived copy instead:
// http://web.archive.org/web/20231202082956/https://service.unece.org/trade/untdid/d20a/trmd/conpva_c.htm
// -- re-check the direct URL first if cross-checking later, in case the
// block has lifted.
//
// Transcription note: generated in bulk (edifact-ls-13gu) the same way
// as every other message-type schema -- a script parsed the source's
// exact rail-art column positions mechanically and verified the result
// balances (all 34 groups open and close consistently)
// before being transcribed here, rather than reading the ASCII nesting
// by eye. Not individually spot-checked against the raw source the way
// the first several message types in this project were -- see
// edifact-ls-13gu for which representative samples were spot-checked
// for this batch.

func init() {
	RegisterSchema(
		MessageID{Type: "CONPVA", Version: "D", Release: "20A", Agency: "UN"},
		conpvaD20aSchema,
		"https://service.unece.org/trade/untdid/d20a/trmd/conpva_c.htm",
	)
}

var conpvaD20aSchema = Schema{Nodes: []SchemaNode{
	{Segment: "BGM", Mandatory: true, MaxRepeat: 1},  // Beginning of message
	{Segment: "DTM", Mandatory: true, MaxRepeat: 9},  // Date/time/period
	{Segment: "AUT", Mandatory: false, MaxRepeat: 2}, // Authentication result
	{Segment: "FTX", Mandatory: false, MaxRepeat: 5}, // Free text
	{ // Segment group 1
		Group: []SchemaNode{
			{Segment: "RFF", Mandatory: true, MaxRepeat: 1},  // Reference
			{Segment: "DTM", Mandatory: false, MaxRepeat: 5}, // Date/time/period
		},
		Mandatory: false, MaxRepeat: 99,
	},
	{ // Segment group 2
		Group: []SchemaNode{
			{Segment: "NAD", Mandatory: true, MaxRepeat: 1},   // Name and address
			{Segment: "LOC", Mandatory: false, MaxRepeat: 25}, // Place/location identification
			{Segment: "FII", Mandatory: false, MaxRepeat: 5},  // Financial institution information
			{ // Segment group 3
				Group: []SchemaNode{
					{Segment: "RFF", Mandatory: true, MaxRepeat: 1},  // Reference
					{Segment: "DTM", Mandatory: false, MaxRepeat: 5}, // Date/time/period
				},
				Mandatory: false, MaxRepeat: 10,
			},
			{ // Segment group 4
				Group: []SchemaNode{
					{Segment: "DOC", Mandatory: true, MaxRepeat: 1},  // Document/message details
					{Segment: "DTM", Mandatory: false, MaxRepeat: 5}, // Date/time/period
				},
				Mandatory: false, MaxRepeat: 5,
			},
			{ // Segment group 5
				Group: []SchemaNode{
					{Segment: "CTA", Mandatory: true, MaxRepeat: 1},  // Contact information
					{Segment: "COM", Mandatory: false, MaxRepeat: 5}, // Communication contact
				},
				Mandatory: false, MaxRepeat: 5,
			},
		},
		Mandatory: true, MaxRepeat: 9,
	},
	{ // Segment group 6
		Group: []SchemaNode{
			{Segment: "CUX", Mandatory: true, MaxRepeat: 1},  // Currencies
			{Segment: "DTM", Mandatory: false, MaxRepeat: 5}, // Date/time/period
		},
		Mandatory: false, MaxRepeat: 5,
	},
	{ // Segment group 7
		Group: []SchemaNode{
			{Segment: "IND", Mandatory: true, MaxRepeat: 1},  // Index details
			{Segment: "RFF", Mandatory: true, MaxRepeat: 1},  // Reference
			{Segment: "DTM", Mandatory: false, MaxRepeat: 1}, // Date/time/period
		},
		Mandatory: false, MaxRepeat: 100,
	},
	{ // Segment group 8
		Group: []SchemaNode{
			{Segment: "BII", Mandatory: true, MaxRepeat: 1},  // Structure identification
			{Segment: "RCS", Mandatory: false, MaxRepeat: 1}, // Requirements and conditions
			{Segment: "PAI", Mandatory: false, MaxRepeat: 1}, // Payment instructions
			{Segment: "PYT", Mandatory: false, MaxRepeat: 1}, // Payment terms
			{Segment: "PCD", Mandatory: false, MaxRepeat: 9}, // Percentage details
			{ // Segment group 9
				Group: []SchemaNode{
					{Segment: "APR", Mandatory: true, MaxRepeat: 1},  // Additional price information
					{Segment: "DTM", Mandatory: false, MaxRepeat: 5}, // Date/time/period
				},
				Mandatory: false, MaxRepeat: 2,
			},
			{ // Segment group 10
				Group: []SchemaNode{
					{Segment: "ARD", Mandatory: true, MaxRepeat: 1},   // Monetary amount function
					{Segment: "MOA", Mandatory: true, MaxRepeat: 6},   // Monetary amount
					{Segment: "FTX", Mandatory: false, MaxRepeat: 10}, // Free text
					{ // Segment group 11
						Group: []SchemaNode{
							{Segment: "TAX", Mandatory: true, MaxRepeat: 1},  // Duty/tax/fee details
							{Segment: "MOA", Mandatory: false, MaxRepeat: 1}, // Monetary amount
							{Segment: "LOC", Mandatory: false, MaxRepeat: 5}, // Place/location identification
						},
						Mandatory: false, MaxRepeat: 5,
					},
					{ // Segment group 12
						Group: []SchemaNode{
							{Segment: "APR", Mandatory: true, MaxRepeat: 1},  // Additional price information
							{Segment: "RFF", Mandatory: false, MaxRepeat: 9}, // Reference
							{Segment: "DTM", Mandatory: false, MaxRepeat: 9}, // Date/time/period
							{Segment: "MOA", Mandatory: false, MaxRepeat: 9}, // Monetary amount
							{Segment: "TAX", Mandatory: false, MaxRepeat: 9}, // Duty/tax/fee details
						},
						Mandatory: false, MaxRepeat: 99,
					},
					{ // Segment group 13
						Group: []SchemaNode{
							{Segment: "ALC", Mandatory: true, MaxRepeat: 1},  // Allowance or charge
							{Segment: "ALI", Mandatory: false, MaxRepeat: 5}, // Additional information
							{Segment: "RFF", Mandatory: false, MaxRepeat: 9}, // Reference
							{Segment: "DTM", Mandatory: false, MaxRepeat: 9}, // Date/time/period
							{Segment: "FTX", Mandatory: false, MaxRepeat: 9}, // Free text
							{ // Segment group 14
								Group: []SchemaNode{
									{Segment: "QTY", Mandatory: true, MaxRepeat: 1},  // Quantity
									{Segment: "RNG", Mandatory: false, MaxRepeat: 1}, // Range details
								},
								Mandatory: false, MaxRepeat: 1,
							},
							{ // Segment group 15
								Group: []SchemaNode{
									{Segment: "PCD", Mandatory: true, MaxRepeat: 1},  // Percentage details
									{Segment: "RNG", Mandatory: false, MaxRepeat: 1}, // Range details
								},
								Mandatory: false, MaxRepeat: 1,
							},
							{ // Segment group 16
								Group: []SchemaNode{
									{Segment: "MOA", Mandatory: true, MaxRepeat: 1},  // Monetary amount
									{Segment: "RNG", Mandatory: false, MaxRepeat: 1}, // Range details
								},
								Mandatory: false, MaxRepeat: 1,
							},
							{ // Segment group 17
								Group: []SchemaNode{
									{Segment: "RTE", Mandatory: true, MaxRepeat: 1},  // Rate details
									{Segment: "RNG", Mandatory: false, MaxRepeat: 1}, // Range details
								},
								Mandatory: false, MaxRepeat: 1,
							},
							{ // Segment group 18
								Group: []SchemaNode{
									{Segment: "TAX", Mandatory: true, MaxRepeat: 1},  // Duty/tax/fee details
									{Segment: "MOA", Mandatory: false, MaxRepeat: 1}, // Monetary amount
								},
								Mandatory: false, MaxRepeat: 1,
							},
						},
						Mandatory: false, MaxRepeat: 99,
					},
				},
				Mandatory: true, MaxRepeat: 100,
			},
			{ // Segment group 19
				Group: []SchemaNode{
					{Segment: "NAD", Mandatory: true, MaxRepeat: 1},   // Name and address
					{Segment: "LOC", Mandatory: false, MaxRepeat: 25}, // Place/location identification
					{Segment: "FII", Mandatory: false, MaxRepeat: 5},  // Financial institution information
					{Segment: "MOA", Mandatory: false, MaxRepeat: 99}, // Monetary amount
					{ // Segment group 20
						Group: []SchemaNode{
							{Segment: "RFF", Mandatory: true, MaxRepeat: 1},  // Reference
							{Segment: "DTM", Mandatory: false, MaxRepeat: 5}, // Date/time/period
						},
						Mandatory: false, MaxRepeat: 10,
					},
					{ // Segment group 21
						Group: []SchemaNode{
							{Segment: "DOC", Mandatory: true, MaxRepeat: 1},  // Document/message details
							{Segment: "DTM", Mandatory: false, MaxRepeat: 5}, // Date/time/period
						},
						Mandatory: false, MaxRepeat: 5,
					},
					{ // Segment group 22
						Group: []SchemaNode{
							{Segment: "CTA", Mandatory: true, MaxRepeat: 1},  // Contact information
							{Segment: "COM", Mandatory: false, MaxRepeat: 5}, // Communication contact
						},
						Mandatory: false, MaxRepeat: 5,
					},
				},
				Mandatory: false, MaxRepeat: 99,
			},
		},
		Mandatory: false, MaxRepeat: 1000,
	},
	{Segment: "UNS", Mandatory: true, MaxRepeat: 1}, // Section control
	{ // Segment group 23
		Group: []SchemaNode{
			{Segment: "BII", Mandatory: true, MaxRepeat: 1},  // Structure identification
			{Segment: "RCS", Mandatory: false, MaxRepeat: 1}, // Requirements and conditions
			{Segment: "QTY", Mandatory: true, MaxRepeat: 6},  // Quantity
			{Segment: "PRI", Mandatory: false, MaxRepeat: 1}, // Price details
			{Segment: "MOA", Mandatory: true, MaxRepeat: 6},  // Monetary amount
			{ // Segment group 24
				Group: []SchemaNode{
					{Segment: "LIN", Mandatory: true, MaxRepeat: 1}, // Line item
					{ // Segment group 25
						Group: []SchemaNode{
							{Segment: "IMD", Mandatory: true, MaxRepeat: 1},  // Item description
							{Segment: "RFF", Mandatory: false, MaxRepeat: 5}, // Reference
							{Segment: "GEI", Mandatory: false, MaxRepeat: 5}, // Processing information
						},
						Mandatory: false, MaxRepeat: 1000,
					},
				},
				Mandatory: true, MaxRepeat: 100,
			},
			{ // Segment group 26
				Group: []SchemaNode{
					{Segment: "APR", Mandatory: true, MaxRepeat: 1},  // Additional price information
					{Segment: "DTM", Mandatory: false, MaxRepeat: 5}, // Date/time/period
				},
				Mandatory: false, MaxRepeat: 2,
			},
			{ // Segment group 27
				Group: []SchemaNode{
					{Segment: "TAX", Mandatory: true, MaxRepeat: 1},  // Duty/tax/fee details
					{Segment: "MOA", Mandatory: false, MaxRepeat: 1}, // Monetary amount
					{Segment: "LOC", Mandatory: false, MaxRepeat: 5}, // Place/location identification
				},
				Mandatory: false, MaxRepeat: 5,
			},
			{ // Segment group 28
				Group: []SchemaNode{
					{Segment: "ALC", Mandatory: true, MaxRepeat: 1},  // Allowance or charge
					{Segment: "ALI", Mandatory: false, MaxRepeat: 1}, // Additional information
					{ // Segment group 29
						Group: []SchemaNode{
							{Segment: "QTY", Mandatory: true, MaxRepeat: 1},  // Quantity
							{Segment: "RNG", Mandatory: false, MaxRepeat: 1}, // Range details
						},
						Mandatory: false, MaxRepeat: 1,
					},
					{ // Segment group 30
						Group: []SchemaNode{
							{Segment: "PCD", Mandatory: true, MaxRepeat: 1},  // Percentage details
							{Segment: "RNG", Mandatory: false, MaxRepeat: 1}, // Range details
						},
						Mandatory: false, MaxRepeat: 1,
					},
					{ // Segment group 31
						Group: []SchemaNode{
							{Segment: "MOA", Mandatory: true, MaxRepeat: 1},  // Monetary amount
							{Segment: "RNG", Mandatory: false, MaxRepeat: 1}, // Range details
						},
						Mandatory: false, MaxRepeat: 1,
					},
					{ // Segment group 32
						Group: []SchemaNode{
							{Segment: "RTE", Mandatory: true, MaxRepeat: 1},  // Rate details
							{Segment: "RNG", Mandatory: false, MaxRepeat: 1}, // Range details
						},
						Mandatory: false, MaxRepeat: 1,
					},
					{ // Segment group 33
						Group: []SchemaNode{
							{Segment: "TAX", Mandatory: true, MaxRepeat: 1},  // Duty/tax/fee details
							{Segment: "MOA", Mandatory: false, MaxRepeat: 1}, // Monetary amount
						},
						Mandatory: false, MaxRepeat: 1,
					},
				},
				Mandatory: false, MaxRepeat: 20,
			},
		},
		Mandatory: false, MaxRepeat: 9999,
	},
	{Segment: "UNS", Mandatory: true, MaxRepeat: 1}, // Section control
	{ // Segment group 34
		Group: []SchemaNode{
			{Segment: "TAX", Mandatory: true, MaxRepeat: 1},  // Duty/tax/fee details
			{Segment: "MOA", Mandatory: false, MaxRepeat: 2}, // Monetary amount
		},
		Mandatory: false, MaxRepeat: 10,
	},
	{Segment: "CNT", Mandatory: false, MaxRepeat: 5}, // Control total
}}
