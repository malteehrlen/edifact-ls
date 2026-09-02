package edifact

// CNTCND D.20A message specification: the structural (segment/group
// presence, order, cardinality) branching diagram for this message,
// UN/EDIFACT directory release D.20A. Scope is structural only -- not
// element-level content -- per edifact-ls-3uzr's epic non-goal.
// 40 segment groups, max nesting depth 3.
//
// Source: https://service.unece.org/trade/untdid/d20a/trmd/cntcnd_c.htm
// section 4.3.1 "Segment table". That URL currently returns HTTP 403 from
// Cloudflare when fetched directly; this data was transcribed from the
// Wayback Machine's archived copy instead:
// http://web.archive.org/web/20231002215916/https://service.unece.org/trade/untdid/d20a/trmd/cntcnd_c.htm
// -- re-check the direct URL first if cross-checking later, in case the
// block has lifted.
//
// Transcription note: generated in bulk (edifact-ls-13gu) the same way
// as every other message-type schema -- a script parsed the source's
// exact rail-art column positions mechanically and verified the result
// balances (all 40 groups open and close consistently)
// before being transcribed here, rather than reading the ASCII nesting
// by eye. Not individually spot-checked against the raw source the way
// the first several message types in this project were -- see
// edifact-ls-13gu for which representative samples were spot-checked
// for this batch.

func init() {
	RegisterSchema(
		MessageID{Type: "CNTCND", Version: "D", Release: "20A", Agency: "UN"},
		cntcndD20aSchema,
		"https://service.unece.org/trade/untdid/d20a/trmd/cntcnd_c.htm",
	)
}

var cntcndD20aSchema = Schema{Nodes: []SchemaNode{
	{Segment: "BGM", Mandatory: true, MaxRepeat: 1},   // Beginning of message
	{Segment: "DTM", Mandatory: true, MaxRepeat: 99},  // Date/time/period
	{Segment: "ALI", Mandatory: false, MaxRepeat: 99}, // Additional information
	{Segment: "AGR", Mandatory: false, MaxRepeat: 99}, // Agreement identification
	{Segment: "MOA", Mandatory: false, MaxRepeat: 99}, // Monetary amount
	{Segment: "GEI", Mandatory: false, MaxRepeat: 99}, // Processing information
	{Segment: "PAI", Mandatory: false, MaxRepeat: 99}, // Payment instructions
	{Segment: "TOD", Mandatory: false, MaxRepeat: 99}, // Terms of delivery or transport
	{Segment: "FTX", Mandatory: false, MaxRepeat: 99}, // Free text
	{ // Segment group 1
		Group: []SchemaNode{
			{Segment: "RFF", Mandatory: true, MaxRepeat: 1},   // Reference
			{Segment: "DTM", Mandatory: false, MaxRepeat: 99}, // Date/time/period
		},
		Mandatory: false, MaxRepeat: 99,
	},
	{ // Segment group 2
		Group: []SchemaNode{
			{Segment: "PRI", Mandatory: true, MaxRepeat: 1},   // Price details
			{Segment: "RNG", Mandatory: false, MaxRepeat: 99}, // Range details
		},
		Mandatory: false, MaxRepeat: 99,
	},
	{ // Segment group 3
		Group: []SchemaNode{
			{Segment: "QTY", Mandatory: true, MaxRepeat: 1},   // Quantity
			{Segment: "DTM", Mandatory: false, MaxRepeat: 99}, // Date/time/period
		},
		Mandatory: false, MaxRepeat: 99,
	},
	{ // Segment group 4
		Group: []SchemaNode{
			{Segment: "DOC", Mandatory: true, MaxRepeat: 1},   // Document/message details
			{Segment: "DTM", Mandatory: false, MaxRepeat: 99}, // Date/time/period
			{Segment: "RFF", Mandatory: false, MaxRepeat: 99}, // Reference
			{Segment: "FTX", Mandatory: false, MaxRepeat: 99}, // Free text
		},
		Mandatory: false, MaxRepeat: 99,
	},
	{ // Segment group 5
		Group: []SchemaNode{
			{Segment: "CUX", Mandatory: true, MaxRepeat: 1},   // Currencies
			{Segment: "DTM", Mandatory: false, MaxRepeat: 99}, // Date/time/period
		},
		Mandatory: false, MaxRepeat: 99,
	},
	{ // Segment group 6
		Group: []SchemaNode{
			{Segment: "TAX", Mandatory: true, MaxRepeat: 1},   // Duty/tax/fee details
			{Segment: "MOA", Mandatory: false, MaxRepeat: 99}, // Monetary amount
		},
		Mandatory: false, MaxRepeat: 99,
	},
	{ // Segment group 7
		Group: []SchemaNode{
			{Segment: "PYT", Mandatory: true, MaxRepeat: 1},   // Payment terms
			{Segment: "DTM", Mandatory: false, MaxRepeat: 99}, // Date/time/period
			{Segment: "PCD", Mandatory: false, MaxRepeat: 99}, // Percentage details
			{Segment: "MOA", Mandatory: false, MaxRepeat: 99}, // Monetary amount
		},
		Mandatory: false, MaxRepeat: 99,
	},
	{ // Segment group 8
		Group: []SchemaNode{
			{Segment: "SEQ", Mandatory: true, MaxRepeat: 1},   // Sequence details
			{Segment: "NAD", Mandatory: true, MaxRepeat: 1},   // Name and address
			{Segment: "GEI", Mandatory: false, MaxRepeat: 99}, // Processing information
			{Segment: "FII", Mandatory: false, MaxRepeat: 9},  // Financial institution information
			{Segment: "FTX", Mandatory: false, MaxRepeat: 99}, // Free text
			{ // Segment group 9
				Group: []SchemaNode{
					{Segment: "RFF", Mandatory: true, MaxRepeat: 1},   // Reference
					{Segment: "DTM", Mandatory: false, MaxRepeat: 99}, // Date/time/period
				},
				Mandatory: false, MaxRepeat: 99,
			},
			{ // Segment group 10
				Group: []SchemaNode{
					{Segment: "CTA", Mandatory: true, MaxRepeat: 1},   // Contact information
					{Segment: "COM", Mandatory: false, MaxRepeat: 99}, // Communication contact
				},
				Mandatory: false, MaxRepeat: 99,
			},
			{ // Segment group 11
				Group: []SchemaNode{
					{Segment: "REL", Mandatory: true, MaxRepeat: 1},   // Relationship
					{Segment: "NAD", Mandatory: false, MaxRepeat: 99}, // Name and address
					{Segment: "RFF", Mandatory: false, MaxRepeat: 99}, // Reference
				},
				Mandatory: false, MaxRepeat: 99,
			},
		},
		Mandatory: true, MaxRepeat: 99,
	},
	{ // Segment group 12
		Group: []SchemaNode{
			{Segment: "ALC", Mandatory: true, MaxRepeat: 1},   // Allowance or charge
			{Segment: "ALI", Mandatory: false, MaxRepeat: 99}, // Additional information
			{Segment: "DTM", Mandatory: false, MaxRepeat: 99}, // Date/time/period
			{ // Segment group 13
				Group: []SchemaNode{
					{Segment: "RFF", Mandatory: true, MaxRepeat: 1},   // Reference
					{Segment: "DTM", Mandatory: false, MaxRepeat: 99}, // Date/time/period
				},
				Mandatory: false, MaxRepeat: 99,
			},
			{ // Segment group 14
				Group: []SchemaNode{
					{Segment: "QTY", Mandatory: true, MaxRepeat: 1},   // Quantity
					{Segment: "RNG", Mandatory: false, MaxRepeat: 99}, // Range details
				},
				Mandatory: false, MaxRepeat: 99,
			},
			{ // Segment group 15
				Group: []SchemaNode{
					{Segment: "PCD", Mandatory: true, MaxRepeat: 1},   // Percentage details
					{Segment: "RNG", Mandatory: false, MaxRepeat: 99}, // Range details
				},
				Mandatory: false, MaxRepeat: 99,
			},
			{ // Segment group 16
				Group: []SchemaNode{
					{Segment: "RTE", Mandatory: true, MaxRepeat: 1},   // Rate details
					{Segment: "RNG", Mandatory: false, MaxRepeat: 99}, // Range details
				},
				Mandatory: false, MaxRepeat: 99,
			},
			{ // Segment group 17
				Group: []SchemaNode{
					{Segment: "MOA", Mandatory: true, MaxRepeat: 1},   // Monetary amount
					{Segment: "RNG", Mandatory: false, MaxRepeat: 99}, // Range details
					{Segment: "DTM", Mandatory: false, MaxRepeat: 99}, // Date/time/period
				},
				Mandatory: false, MaxRepeat: 99,
			},
			{ // Segment group 18
				Group: []SchemaNode{
					{Segment: "TAX", Mandatory: true, MaxRepeat: 1},   // Duty/tax/fee details
					{Segment: "MOA", Mandatory: false, MaxRepeat: 99}, // Monetary amount
				},
				Mandatory: false, MaxRepeat: 99,
			},
		},
		Mandatory: false, MaxRepeat: 99,
	},
	{ // Segment group 19
		Group: []SchemaNode{
			{Segment: "FOR", Mandatory: true, MaxRepeat: 1},   // Formula
			{Segment: "GEI", Mandatory: false, MaxRepeat: 99}, // Processing information
			{Segment: "ALI", Mandatory: false, MaxRepeat: 99}, // Additional information
			{Segment: "RFF", Mandatory: false, MaxRepeat: 99}, // Reference
			{Segment: "DTM", Mandatory: false, MaxRepeat: 99}, // Date/time/period
			{ // Segment group 20
				Group: []SchemaNode{
					{Segment: "FSQ", Mandatory: true, MaxRepeat: 1},   // Formula sequence
					{Segment: "RTE", Mandatory: false, MaxRepeat: 99}, // Rate details
					{Segment: "PCD", Mandatory: false, MaxRepeat: 99}, // Percentage details
					{Segment: "MEA", Mandatory: false, MaxRepeat: 99}, // Measurements
					{Segment: "MOA", Mandatory: false, MaxRepeat: 99}, // Monetary amount
					{Segment: "RFF", Mandatory: false, MaxRepeat: 99}, // Reference
					{Segment: "DTM", Mandatory: false, MaxRepeat: 99}, // Date/time/period
					{ // Segment group 21
						Group: []SchemaNode{
							{Segment: "IND", Mandatory: true, MaxRepeat: 1},   // Index details
							{Segment: "ALI", Mandatory: false, MaxRepeat: 99}, // Additional information
							{Segment: "DTM", Mandatory: false, MaxRepeat: 99}, // Date/time/period
							{Segment: "QTY", Mandatory: false, MaxRepeat: 99}, // Quantity
						},
						Mandatory: false, MaxRepeat: 99,
					},
				},
				Mandatory: false, MaxRepeat: 99,
			},
		},
		Mandatory: false, MaxRepeat: 99,
	},
	{ // Segment group 22
		Group: []SchemaNode{
			{Segment: "LIN", Mandatory: true, MaxRepeat: 1},   // Line item
			{Segment: "PIA", Mandatory: false, MaxRepeat: 99}, // Additional product id
			{Segment: "IMD", Mandatory: false, MaxRepeat: 99}, // Item description
			{Segment: "MOA", Mandatory: false, MaxRepeat: 99}, // Monetary amount
			{Segment: "DTM", Mandatory: false, MaxRepeat: 99}, // Date/time/period
			{Segment: "FTX", Mandatory: false, MaxRepeat: 99}, // Free text
			{ // Segment group 23
				Group: []SchemaNode{
					{Segment: "RFF", Mandatory: true, MaxRepeat: 1},   // Reference
					{Segment: "DTM", Mandatory: false, MaxRepeat: 99}, // Date/time/period
				},
				Mandatory: false, MaxRepeat: 99,
			},
			{ // Segment group 24
				Group: []SchemaNode{
					{Segment: "TAX", Mandatory: true, MaxRepeat: 1},   // Duty/tax/fee details
					{Segment: "MOA", Mandatory: false, MaxRepeat: 99}, // Monetary amount
				},
				Mandatory: false, MaxRepeat: 99,
			},
			{ // Segment group 25
				Group: []SchemaNode{
					{Segment: "GEI", Mandatory: true, MaxRepeat: 1},   // Processing information
					{Segment: "RTE", Mandatory: false, MaxRepeat: 99}, // Rate details
					{Segment: "MOA", Mandatory: false, MaxRepeat: 99}, // Monetary amount
					{Segment: "TAX", Mandatory: false, MaxRepeat: 99}, // Duty/tax/fee details
					{Segment: "PCD", Mandatory: false, MaxRepeat: 99}, // Percentage details
					{Segment: "DTM", Mandatory: false, MaxRepeat: 99}, // Date/time/period
				},
				Mandatory: false, MaxRepeat: 99,
			},
			{ // Segment group 26
				Group: []SchemaNode{
					{Segment: "QTY", Mandatory: true, MaxRepeat: 1},   // Quantity
					{Segment: "DTM", Mandatory: false, MaxRepeat: 99}, // Date/time/period
				},
				Mandatory: false, MaxRepeat: 99,
			},
			{ // Segment group 27
				Group: []SchemaNode{
					{Segment: "FOR", Mandatory: true, MaxRepeat: 1},   // Formula
					{Segment: "ALI", Mandatory: false, MaxRepeat: 99}, // Additional information
					{Segment: "MOA", Mandatory: false, MaxRepeat: 99}, // Monetary amount
					{Segment: "RFF", Mandatory: false, MaxRepeat: 99}, // Reference
					{Segment: "DTM", Mandatory: false, MaxRepeat: 99}, // Date/time/period
					{Segment: "GEI", Mandatory: false, MaxRepeat: 99}, // Processing information
					{ // Segment group 28
						Group: []SchemaNode{
							{Segment: "FSQ", Mandatory: true, MaxRepeat: 1},   // Formula sequence
							{Segment: "RTE", Mandatory: false, MaxRepeat: 99}, // Rate details
							{Segment: "PCD", Mandatory: false, MaxRepeat: 99}, // Percentage details
							{Segment: "MEA", Mandatory: false, MaxRepeat: 99}, // Measurements
							{Segment: "MOA", Mandatory: false, MaxRepeat: 99}, // Monetary amount
							{Segment: "RFF", Mandatory: false, MaxRepeat: 99}, // Reference
							{Segment: "DTM", Mandatory: false, MaxRepeat: 99}, // Date/time/period
							{Segment: "GEI", Mandatory: false, MaxRepeat: 99}, // Processing information
							{Segment: "QTY", Mandatory: false, MaxRepeat: 99}, // Quantity
						},
						Mandatory: false, MaxRepeat: 99,
					},
				},
				Mandatory: false, MaxRepeat: 99,
			},
			{ // Segment group 29
				Group: []SchemaNode{
					{Segment: "ALC", Mandatory: true, MaxRepeat: 1},   // Allowance or charge
					{Segment: "ALI", Mandatory: false, MaxRepeat: 99}, // Additional information
					{Segment: "DTM", Mandatory: false, MaxRepeat: 99}, // Date/time/period
					{ // Segment group 30
						Group: []SchemaNode{
							{Segment: "RFF", Mandatory: true, MaxRepeat: 1},   // Reference
							{Segment: "DTM", Mandatory: false, MaxRepeat: 99}, // Date/time/period
						},
						Mandatory: false, MaxRepeat: 99,
					},
					{ // Segment group 31
						Group: []SchemaNode{
							{Segment: "QTY", Mandatory: true, MaxRepeat: 1},   // Quantity
							{Segment: "RNG", Mandatory: false, MaxRepeat: 99}, // Range details
						},
						Mandatory: false, MaxRepeat: 99,
					},
					{ // Segment group 32
						Group: []SchemaNode{
							{Segment: "PCD", Mandatory: true, MaxRepeat: 1},   // Percentage details
							{Segment: "RNG", Mandatory: false, MaxRepeat: 99}, // Range details
						},
						Mandatory: false, MaxRepeat: 99,
					},
					{ // Segment group 33
						Group: []SchemaNode{
							{Segment: "RTE", Mandatory: true, MaxRepeat: 1},   // Rate details
							{Segment: "RNG", Mandatory: false, MaxRepeat: 99}, // Range details
						},
						Mandatory: false, MaxRepeat: 99,
					},
					{ // Segment group 34
						Group: []SchemaNode{
							{Segment: "MOA", Mandatory: true, MaxRepeat: 1},   // Monetary amount
							{Segment: "RNG", Mandatory: false, MaxRepeat: 99}, // Range details
							{Segment: "DTM", Mandatory: false, MaxRepeat: 99}, // Date/time/period
						},
						Mandatory: false, MaxRepeat: 99,
					},
					{ // Segment group 35
						Group: []SchemaNode{
							{Segment: "TAX", Mandatory: true, MaxRepeat: 1},   // Duty/tax/fee details
							{Segment: "MOA", Mandatory: false, MaxRepeat: 99}, // Monetary amount
						},
						Mandatory: false, MaxRepeat: 99,
					},
				},
				Mandatory: false, MaxRepeat: 99,
			},
			{ // Segment group 36
				Group: []SchemaNode{
					{Segment: "PRI", Mandatory: true, MaxRepeat: 1},   // Price details
					{Segment: "APR", Mandatory: false, MaxRepeat: 99}, // Additional price information
					{Segment: "RNG", Mandatory: false, MaxRepeat: 99}, // Range details
					{Segment: "CUX", Mandatory: false, MaxRepeat: 99}, // Currencies
					{Segment: "DTM", Mandatory: false, MaxRepeat: 99}, // Date/time/period
				},
				Mandatory: false, MaxRepeat: 99,
			},
			{ // Segment group 37
				Group: []SchemaNode{
					{Segment: "TOD", Mandatory: true, MaxRepeat: 1},   // Terms of delivery or transport
					{Segment: "LOC", Mandatory: false, MaxRepeat: 99}, // Place/location identification
				},
				Mandatory: false, MaxRepeat: 99,
			},
			{ // Segment group 38
				Group: []SchemaNode{
					{Segment: "PAC", Mandatory: true, MaxRepeat: 1},   // Package
					{Segment: "MEA", Mandatory: false, MaxRepeat: 99}, // Measurements
					{ // Segment group 39
						Group: []SchemaNode{
							{Segment: "PCI", Mandatory: true, MaxRepeat: 1},   // Package identification
							{Segment: "RFF", Mandatory: false, MaxRepeat: 99}, // Reference
							{Segment: "DTM", Mandatory: false, MaxRepeat: 99}, // Date/time/period
							{Segment: "GIN", Mandatory: false, MaxRepeat: 99}, // Goods identity number
						},
						Mandatory: false, MaxRepeat: 99,
					},
				},
				Mandatory: false, MaxRepeat: 99,
			},
			{ // Segment group 40
				Group: []SchemaNode{
					{Segment: "NAD", Mandatory: true, MaxRepeat: 1},   // Name and address
					{Segment: "GEI", Mandatory: false, MaxRepeat: 99}, // Processing information
					{Segment: "FII", Mandatory: false, MaxRepeat: 99}, // Financial institution information
					{Segment: "MOA", Mandatory: false, MaxRepeat: 99}, // Monetary amount
					{Segment: "TAX", Mandatory: false, MaxRepeat: 99}, // Duty/tax/fee details
					{Segment: "RFF", Mandatory: false, MaxRepeat: 99}, // Reference
				},
				Mandatory: false, MaxRepeat: 99,
			},
		},
		Mandatory: false, MaxRepeat: 99,
	},
}}
