package edifact

// INVRPT D.20A message specification: the structural (segment/group
// presence, order, cardinality) branching diagram for the Inventory Report
// message, UN/EDIFACT directory release D.20A. Scope is structural only --
// not element-level content -- per edifact-ls-3uzr's epic non-goal.
// 19 segment groups, max nesting depth 6.
//
// Source: https://service.unece.org/trade/untdid/d20a/trmd/invrpt_c.htm
// section 4.3.1 "Segment table". That URL currently returns HTTP 403 from
// Cloudflare when fetched directly; this data was transcribed from the
// Wayback Machine's archived copy instead:
// http://web.archive.org/web/20240518232655/https://service.unece.org/trade/untdid/d20a/trmd/invrpt_c.htm
// -- re-check the direct URL first if cross-checking later, in case the
// block has lifted.
//
// Transcription note: generated the same way as iftmcs_d21a.go -- a
// one-off script parsed the source's exact rail-art column positions
// mechanically and verified the result balances (all 19 groups
// open and close consistently) before being transcribed here, rather than
// reading the ASCII nesting by eye.

func init() {
	RegisterSchema(
		MessageID{Type: "INVRPT", Version: "D", Release: "20A", Agency: "UN"},
		invrptD20aSchema,
		"https://service.unece.org/trade/untdid/d20a/trmd/invrpt_c.htm",
	)
}

var invrptD20aSchema = Schema{Nodes: []SchemaNode{
	{Segment: "BGM", Mandatory: true, MaxRepeat: 1},   // Beginning of message
	{Segment: "FTX", Mandatory: false, MaxRepeat: 99}, // Free text
	{Segment: "DTM", Mandatory: true, MaxRepeat: 10},  // Date/time/period
	{ // Segment group 1
		Group: []SchemaNode{
			{Segment: "RFF", Mandatory: true, MaxRepeat: 1},  // Reference
			{Segment: "DTM", Mandatory: false, MaxRepeat: 1}, // Date/time/period
		},
		Mandatory: false, MaxRepeat: 10,
	},
	{ // Segment group 2
		Group: []SchemaNode{
			{Segment: "NAD", Mandatory: true, MaxRepeat: 1},  // Name and address
			{Segment: "LOC", Mandatory: false, MaxRepeat: 5}, // Place/location identification
			{ // Segment group 3
				Group: []SchemaNode{
					{Segment: "RFF", Mandatory: true, MaxRepeat: 1},  // Reference
					{Segment: "DTM", Mandatory: false, MaxRepeat: 1}, // Date/time/period
				},
				Mandatory: false, MaxRepeat: 99,
			},
			{ // Segment group 4
				Group: []SchemaNode{
					{Segment: "CTA", Mandatory: true, MaxRepeat: 1},  // Contact information
					{Segment: "COM", Mandatory: false, MaxRepeat: 5}, // Communication contact
				},
				Mandatory: false, MaxRepeat: 10,
			},
		},
		Mandatory: false, MaxRepeat: 20,
	},
	{ // Segment group 5
		Group: []SchemaNode{
			{Segment: "CUX", Mandatory: true, MaxRepeat: 1},  // Currencies
			{Segment: "DTM", Mandatory: false, MaxRepeat: 1}, // Date/time/period
		},
		Mandatory: false, MaxRepeat: 5,
	},
	{ // Segment group 6
		Group: []SchemaNode{
			{Segment: "CPS", Mandatory: true, MaxRepeat: 1}, // Consignment packing sequence
			{ // Segment group 7
				Group: []SchemaNode{
					{Segment: "PAC", Mandatory: true, MaxRepeat: 1},     // Package
					{Segment: "PCI", Mandatory: false, MaxRepeat: 1000}, // Package identification
					{ // Segment group 8
						Group: []SchemaNode{
							{Segment: "QTY", Mandatory: true, MaxRepeat: 1},     // Quantity
							{Segment: "GIN", Mandatory: false, MaxRepeat: 9999}, // Goods identity number
							{Segment: "DTM", Mandatory: false, MaxRepeat: 5},    // Date/time/period
						},
						Mandatory: false, MaxRepeat: 9999,
					},
				},
				Mandatory: false, MaxRepeat: 999,
			},
		},
		Mandatory: false, MaxRepeat: 9999,
	},
	{ // Segment group 9
		Group: []SchemaNode{
			{Segment: "LIN", Mandatory: true, MaxRepeat: 1},   // Line item
			{Segment: "PIA", Mandatory: false, MaxRepeat: 10}, // Additional product id
			{Segment: "IMD", Mandatory: false, MaxRepeat: 10}, // Item description
			{Segment: "NAD", Mandatory: false, MaxRepeat: 1},  // Name and address
			{Segment: "MEA", Mandatory: false, MaxRepeat: 10}, // Measurements
			{Segment: "ALI", Mandatory: false, MaxRepeat: 10}, // Additional information
			{Segment: "LOC", Mandatory: false, MaxRepeat: 5},  // Place/location identification
			{Segment: "DTM", Mandatory: false, MaxRepeat: 5},  // Date/time/period
			{Segment: "FTX", Mandatory: false, MaxRepeat: 99}, // Free text
			{ // Segment group 10
				Group: []SchemaNode{
					{Segment: "RFF", Mandatory: true, MaxRepeat: 1},  // Reference
					{Segment: "DTM", Mandatory: false, MaxRepeat: 1}, // Date/time/period
				},
				Mandatory: false, MaxRepeat: 10,
			},
			{ // Segment group 11
				Group: []SchemaNode{
					{Segment: "CTA", Mandatory: true, MaxRepeat: 1},  // Contact information
					{Segment: "COM", Mandatory: false, MaxRepeat: 9}, // Communication contact
				},
				Mandatory: false, MaxRepeat: 1,
			},
			{ // Segment group 12
				Group: []SchemaNode{
					{Segment: "INV", Mandatory: true, MaxRepeat: 1},     // Inventory management related details
					{Segment: "QTY", Mandatory: true, MaxRepeat: 9},     // Quantity
					{Segment: "GIN", Mandatory: false, MaxRepeat: 9999}, // Goods identity number
					{Segment: "LOC", Mandatory: false, MaxRepeat: 5},    // Place/location identification
					{Segment: "DTM", Mandatory: false, MaxRepeat: 5},    // Date/time/period
					{Segment: "STS", Mandatory: false, MaxRepeat: 9},    // Status
					{Segment: "FTX", Mandatory: false, MaxRepeat: 9},    // Free text
					{ // Segment group 13
						Group: []SchemaNode{
							{Segment: "NAD", Mandatory: true, MaxRepeat: 1},  // Name and address
							{Segment: "LOC", Mandatory: false, MaxRepeat: 1}, // Place/location identification
						},
						Mandatory: false, MaxRepeat: 5,
					},
					{ // Segment group 14
						Group: []SchemaNode{
							{Segment: "PRI", Mandatory: true, MaxRepeat: 1},  // Price details
							{Segment: "CUX", Mandatory: false, MaxRepeat: 1}, // Currencies
							{Segment: "DTM", Mandatory: false, MaxRepeat: 1}, // Date/time/period
						},
						Mandatory: false, MaxRepeat: 5,
					},
					{ // Segment group 15
						Group: []SchemaNode{
							{Segment: "RFF", Mandatory: true, MaxRepeat: 1},  // Reference
							{Segment: "DTM", Mandatory: false, MaxRepeat: 1}, // Date/time/period
						},
						Mandatory: false, MaxRepeat: 10,
					},
					{ // Segment group 16
						Group: []SchemaNode{
							{Segment: "CPS", Mandatory: true, MaxRepeat: 1}, // Consignment packing sequence
							{ // Segment group 17
								Group: []SchemaNode{
									{Segment: "PAC", Mandatory: true, MaxRepeat: 1},   // Package
									{Segment: "MEA", Mandatory: false, MaxRepeat: 10}, // Measurements
									{Segment: "QTY", Mandatory: false, MaxRepeat: 10}, // Quantity
									{ // Segment group 18
										Group: []SchemaNode{
											{Segment: "PCI", Mandatory: true, MaxRepeat: 1},     // Package identification
											{Segment: "RFF", Mandatory: false, MaxRepeat: 1},    // Reference
											{Segment: "DTM", Mandatory: false, MaxRepeat: 5},    // Date/time/period
											{Segment: "GIN", Mandatory: false, MaxRepeat: 9999}, // Goods identity number
											{ // Segment group 19
												Group: []SchemaNode{
													{Segment: "GIR", Mandatory: true, MaxRepeat: 1},  // Related identification numbers
													{Segment: "DTM", Mandatory: false, MaxRepeat: 5}, // Date/time/period
												},
												Mandatory: false, MaxRepeat: 99,
											},
										},
										Mandatory: false, MaxRepeat: 9999,
									},
								},
								Mandatory: false, MaxRepeat: 9999,
							},
						},
						Mandatory: false, MaxRepeat: 9999,
					},
				},
				Mandatory: false, MaxRepeat: 9999,
			},
		},
		Mandatory: false, MaxRepeat: 9999999,
	},
}}
