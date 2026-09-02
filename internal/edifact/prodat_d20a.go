package edifact

// PRODAT D.20A message specification: the structural (segment/group
// presence, order, cardinality) branching diagram for this message,
// UN/EDIFACT directory release D.20A. Scope is structural only -- not
// element-level content -- per edifact-ls-3uzr's epic non-goal.
// 28 segment groups, max nesting depth 4.
//
// Source: https://service.unece.org/trade/untdid/d20a/trmd/prodat_c.htm
// section 4.3.1 "Segment table". That URL currently returns HTTP 403 from
// Cloudflare when fetched directly; this data was transcribed from the
// Wayback Machine's archived copy instead:
// http://web.archive.org/web/20231207013932/https://service.unece.org/trade/untdid/d20a/trmd/prodat_c.htm
// -- re-check the direct URL first if cross-checking later, in case the
// block has lifted.
//
// Transcription note: generated in bulk (edifact-ls-13gu) the same way
// as every other message-type schema -- a script parsed the source's
// exact rail-art column positions mechanically and verified the result
// balances (all 28 groups open and close consistently)
// before being transcribed here, rather than reading the ASCII nesting
// by eye. Not individually spot-checked against the raw source the way
// the first several message types in this project were -- see
// edifact-ls-13gu for which representative samples were spot-checked
// for this batch.

func init() {
	RegisterSchema(
		MessageID{Type: "PRODAT", Version: "D", Release: "20A", Agency: "UN"},
		prodatD20aSchema,
		"https://service.unece.org/trade/untdid/d20a/trmd/prodat_c.htm",
	)
}

var prodatD20aSchema = Schema{Nodes: []SchemaNode{
	{Segment: "BGM", Mandatory: true, MaxRepeat: 1},   // Beginning of message
	{Segment: "DTM", Mandatory: true, MaxRepeat: 10},  // Date/time/period
	{Segment: "ALI", Mandatory: false, MaxRepeat: 5},  // Additional information
	{Segment: "IMD", Mandatory: false, MaxRepeat: 10}, // Item description
	{Segment: "FTX", Mandatory: false, MaxRepeat: 5},  // Free text
	{Segment: "PGI", Mandatory: false, MaxRepeat: 10}, // Product group information
	{Segment: "CUX", Mandatory: false, MaxRepeat: 9},  // Currencies
	{ // Segment group 1
		Group: []SchemaNode{
			{Segment: "TRU", Mandatory: true, MaxRepeat: 1},  // Technical rules
			{Segment: "DTM", Mandatory: false, MaxRepeat: 1}, // Date/time/period
		},
		Mandatory: false, MaxRepeat: 10,
	},
	{ // Segment group 2
		Group: []SchemaNode{
			{Segment: "RCS", Mandatory: true, MaxRepeat: 1},  // Requirements and conditions
			{Segment: "PIA", Mandatory: false, MaxRepeat: 5}, // Additional product id
		},
		Mandatory: false, MaxRepeat: 10,
	},
	{ // Segment group 3
		Group: []SchemaNode{
			{Segment: "RFF", Mandatory: true, MaxRepeat: 1},  // Reference
			{Segment: "DTM", Mandatory: false, MaxRepeat: 5}, // Date/time/period
		},
		Mandatory: false, MaxRepeat: 99,
	},
	{ // Segment group 4
		Group: []SchemaNode{
			{Segment: "NAD", Mandatory: true, MaxRepeat: 1}, // Name and address
			{ // Segment group 5
				Group: []SchemaNode{
					{Segment: "CTA", Mandatory: true, MaxRepeat: 1},   // Contact information
					{Segment: "COM", Mandatory: false, MaxRepeat: 10}, // Communication contact
				},
				Mandatory: false, MaxRepeat: 5,
			},
			{ // Segment group 6
				Group: []SchemaNode{
					{Segment: "RFF", Mandatory: true, MaxRepeat: 1},  // Reference
					{Segment: "DTM", Mandatory: false, MaxRepeat: 5}, // Date/time/period
				},
				Mandatory: false, MaxRepeat: 5,
			},
		},
		Mandatory: false, MaxRepeat: 99,
	},
	{ // Segment group 7
		Group: []SchemaNode{
			{Segment: "CCI", Mandatory: true, MaxRepeat: 1},   // Characteristic/class id
			{Segment: "CAV", Mandatory: false, MaxRepeat: 10}, // Characteristic value
			{Segment: "MEA", Mandatory: false, MaxRepeat: 10}, // Measurements
		},
		Mandatory: false, MaxRepeat: 999,
	},
	{ // Segment group 8
		Group: []SchemaNode{
			{Segment: "EFI", Mandatory: true, MaxRepeat: 1},   // External file link identification
			{Segment: "CED", Mandatory: false, MaxRepeat: 99}, // Computer environment details
			{Segment: "COM", Mandatory: false, MaxRepeat: 9},  // Communication contact
			{Segment: "RFF", Mandatory: false, MaxRepeat: 9},  // Reference
			{Segment: "DTM", Mandatory: false, MaxRepeat: 9},  // Date/time/period
			{Segment: "QTY", Mandatory: false, MaxRepeat: 9},  // Quantity
		},
		Mandatory: false, MaxRepeat: 99,
	},
	{ // Segment group 9
		Group: []SchemaNode{
			{Segment: "LIN", Mandatory: true, MaxRepeat: 1},   // Line item
			{Segment: "PIA", Mandatory: false, MaxRepeat: 10}, // Additional product id
			{Segment: "DTM", Mandatory: false, MaxRepeat: 99}, // Date/time/period
			{Segment: "MEA", Mandatory: false, MaxRepeat: 10}, // Measurements
			{Segment: "HAN", Mandatory: false, MaxRepeat: 5},  // Handling instructions
			{Segment: "DOC", Mandatory: false, MaxRepeat: 99}, // Document/message details
			{Segment: "FTX", Mandatory: false, MaxRepeat: 99}, // Free text
			{Segment: "PGI", Mandatory: false, MaxRepeat: 10}, // Product group information
			{ // Segment group 10
				Group: []SchemaNode{
					{Segment: "IMD", Mandatory: true, MaxRepeat: 1},   // Item description
					{Segment: "FTX", Mandatory: false, MaxRepeat: 99}, // Free text
				},
				Mandatory: false, MaxRepeat: 10,
			},
			{ // Segment group 11
				Group: []SchemaNode{
					{Segment: "TRU", Mandatory: true, MaxRepeat: 1},  // Technical rules
					{Segment: "DTM", Mandatory: false, MaxRepeat: 1}, // Date/time/period
				},
				Mandatory: false, MaxRepeat: 10,
			},
			{ // Segment group 12
				Group: []SchemaNode{
					{Segment: "RCS", Mandatory: true, MaxRepeat: 1},  // Requirements and conditions
					{Segment: "PIA", Mandatory: false, MaxRepeat: 5}, // Additional product id
				},
				Mandatory: false, MaxRepeat: 10,
			},
			{ // Segment group 13
				Group: []SchemaNode{
					{Segment: "QTY", Mandatory: true, MaxRepeat: 1},  // Quantity
					{Segment: "DTM", Mandatory: false, MaxRepeat: 5}, // Date/time/period
					{Segment: "STS", Mandatory: false, MaxRepeat: 5}, // Status
				},
				Mandatory: false, MaxRepeat: 10,
			},
			{ // Segment group 14
				Group: []SchemaNode{
					{Segment: "PRI", Mandatory: true, MaxRepeat: 1},  // Price details
					{Segment: "CUX", Mandatory: false, MaxRepeat: 1}, // Currencies
					{Segment: "RNG", Mandatory: false, MaxRepeat: 1}, // Range details
				},
				Mandatory: false, MaxRepeat: 5,
			},
			{ // Segment group 15
				Group: []SchemaNode{
					{Segment: "CCI", Mandatory: true, MaxRepeat: 1},   // Characteristic/class id
					{Segment: "CAV", Mandatory: false, MaxRepeat: 10}, // Characteristic value
					{Segment: "MEA", Mandatory: false, MaxRepeat: 10}, // Measurements
				},
				Mandatory: false, MaxRepeat: 99999,
			},
			{ // Segment group 16
				Group: []SchemaNode{
					{Segment: "ALI", Mandatory: true, MaxRepeat: 1},  // Additional information
					{Segment: "PCD", Mandatory: false, MaxRepeat: 5}, // Percentage details
				},
				Mandatory: false, MaxRepeat: 999,
			},
			{ // Segment group 17
				Group: []SchemaNode{
					{Segment: "RFF", Mandatory: true, MaxRepeat: 1},  // Reference
					{Segment: "DTM", Mandatory: false, MaxRepeat: 5}, // Date/time/period
				},
				Mandatory: false, MaxRepeat: 99,
			},
			{ // Segment group 18
				Group: []SchemaNode{
					{Segment: "NAD", Mandatory: true, MaxRepeat: 1},   // Name and address
					{Segment: "RFF", Mandatory: false, MaxRepeat: 99}, // Reference
					{Segment: "QTY", Mandatory: false, MaxRepeat: 5},  // Quantity
					{ // Segment group 19
						Group: []SchemaNode{
							{Segment: "CTA", Mandatory: true, MaxRepeat: 1},   // Contact information
							{Segment: "COM", Mandatory: false, MaxRepeat: 10}, // Communication contact
						},
						Mandatory: false, MaxRepeat: 5,
					},
				},
				Mandatory: false, MaxRepeat: 99,
			},
			{ // Segment group 20
				Group: []SchemaNode{
					{Segment: "DGS", Mandatory: true, MaxRepeat: 1},  // Dangerous goods
					{Segment: "QTY", Mandatory: false, MaxRepeat: 1}, // Quantity
					{Segment: "FTX", Mandatory: false, MaxRepeat: 5}, // Free text
				},
				Mandatory: false, MaxRepeat: 10,
			},
			{ // Segment group 21
				Group: []SchemaNode{
					{Segment: "PAC", Mandatory: true, MaxRepeat: 1},   // Package
					{Segment: "MEA", Mandatory: false, MaxRepeat: 10}, // Measurements
					{Segment: "QTY", Mandatory: false, MaxRepeat: 5},  // Quantity
					{Segment: "HAN", Mandatory: false, MaxRepeat: 5},  // Handling instructions
					{Segment: "PCI", Mandatory: false, MaxRepeat: 5},  // Package identification
					{ // Segment group 22
						Group: []SchemaNode{
							{Segment: "COD", Mandatory: true, MaxRepeat: 1},  // Component details
							{Segment: "MEA", Mandatory: false, MaxRepeat: 9}, // Measurements
							{Segment: "QTY", Mandatory: false, MaxRepeat: 9}, // Quantity
							{Segment: "PCD", Mandatory: false, MaxRepeat: 9}, // Percentage details
						},
						Mandatory: false, MaxRepeat: 99,
					},
				},
				Mandatory: false, MaxRepeat: 5,
			},
			{ // Segment group 23
				Group: []SchemaNode{
					{Segment: "HYN", Mandatory: true, MaxRepeat: 1},   // Hierarchy information
					{Segment: "PIA", Mandatory: false, MaxRepeat: 10}, // Additional product id
					{Segment: "QTY", Mandatory: false, MaxRepeat: 5},  // Quantity
					{Segment: "FTX", Mandatory: false, MaxRepeat: 99}, // Free text
					{ // Segment group 24
						Group: []SchemaNode{
							{Segment: "RFF", Mandatory: true, MaxRepeat: 1},  // Reference
							{Segment: "DTM", Mandatory: false, MaxRepeat: 5}, // Date/time/period
						},
						Mandatory: false, MaxRepeat: 999,
					},
					{ // Segment group 25
						Group: []SchemaNode{
							{Segment: "CCI", Mandatory: true, MaxRepeat: 1},   // Characteristic/class id
							{Segment: "CAV", Mandatory: false, MaxRepeat: 10}, // Characteristic value
							{Segment: "MEA", Mandatory: false, MaxRepeat: 10}, // Measurements
						},
						Mandatory: false, MaxRepeat: 99,
					},
					{ // Segment group 26
						Group: []SchemaNode{
							{Segment: "NAD", Mandatory: true, MaxRepeat: 1},   // Name and address
							{Segment: "PIA", Mandatory: false, MaxRepeat: 10}, // Additional product id
							{Segment: "QTY", Mandatory: false, MaxRepeat: 5},  // Quantity
							{ // Segment group 27
								Group: []SchemaNode{
									{Segment: "CCI", Mandatory: true, MaxRepeat: 1},   // Characteristic/class id
									{Segment: "CAV", Mandatory: false, MaxRepeat: 99}, // Characteristic value
									{Segment: "MEA", Mandatory: false, MaxRepeat: 10}, // Measurements
								},
								Mandatory: false, MaxRepeat: 99,
							},
						},
						Mandatory: false, MaxRepeat: 99,
					},
				},
				Mandatory: false, MaxRepeat: 999999,
			},
			{ // Segment group 28
				Group: []SchemaNode{
					{Segment: "EFI", Mandatory: true, MaxRepeat: 1},   // External file link identification
					{Segment: "CED", Mandatory: false, MaxRepeat: 99}, // Computer environment details
					{Segment: "COM", Mandatory: false, MaxRepeat: 9},  // Communication contact
					{Segment: "RFF", Mandatory: false, MaxRepeat: 9},  // Reference
					{Segment: "DTM", Mandatory: false, MaxRepeat: 9},  // Date/time/period
					{Segment: "QTY", Mandatory: false, MaxRepeat: 9},  // Quantity
				},
				Mandatory: false, MaxRepeat: 99,
			},
		},
		Mandatory: false, MaxRepeat: 999999,
	},
}}
