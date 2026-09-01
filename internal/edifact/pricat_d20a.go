package edifact

// PRICAT D.20A message specification: the structural (segment/group
// presence, order, cardinality) branching diagram for the Price/Sales Catalogue
// message, UN/EDIFACT directory release D.20A. Scope is structural only --
// not element-level content -- per edifact-ls-3uzr's epic non-goal.
// 60 segment groups, max nesting depth 5.
//
// Source: https://service.unece.org/trade/untdid/d20a/trmd/pricat_c.htm
// section 4.3.1 "Segment table". That URL currently returns HTTP 403 from
// Cloudflare when fetched directly; this data was transcribed from the
// Wayback Machine's archived copy instead:
// http://web.archive.org/web/20250908085614/https://service.unece.org/trade/untdid/d20a/trmd/pricat_c.htm
// -- re-check the direct URL first if cross-checking later, in case the
// block has lifted.
//
// Transcription note: generated the same way as iftmcs_d21a.go -- a
// one-off script parsed the source's exact rail-art column positions
// mechanically and verified the result balances (all 60 groups
// open and close consistently) before being transcribed here, rather than
// reading the ASCII nesting by eye.

func init() {
	RegisterSchema(
		MessageID{Type: "PRICAT", Version: "D", Release: "20A", Agency: "UN"},
		pricatD20aSchema,
	)
}

var pricatD20aSchema = Schema{Nodes: []SchemaNode{
	{Segment: "BGM", Mandatory: true, MaxRepeat: 1},   // Beginning of message
	{Segment: "DTM", Mandatory: true, MaxRepeat: 35},  // Date/time/period
	{Segment: "ALI", Mandatory: false, MaxRepeat: 5},  // Additional information
	{Segment: "FTX", Mandatory: false, MaxRepeat: 99}, // Free text
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
			{ // Segment group 3
				Group: []SchemaNode{
					{Segment: "RFF", Mandatory: true, MaxRepeat: 1},  // Reference
					{Segment: "DTM", Mandatory: false, MaxRepeat: 5}, // Date/time/period
				},
				Mandatory: false, MaxRepeat: 10,
			},
			{ // Segment group 4
				Group: []SchemaNode{
					{Segment: "CTA", Mandatory: true, MaxRepeat: 1},  // Contact information
					{Segment: "COM", Mandatory: false, MaxRepeat: 5}, // Communication contact
				},
				Mandatory: false, MaxRepeat: 5,
			},
		},
		Mandatory: false, MaxRepeat: 99,
	},
	{ // Segment group 5
		Group: []SchemaNode{
			{Segment: "TAX", Mandatory: true, MaxRepeat: 1},  // Duty/tax/fee details
			{Segment: "MOA", Mandatory: false, MaxRepeat: 1}, // Monetary amount
			{Segment: "DTM", Mandatory: false, MaxRepeat: 1}, // Date/time/period
		},
		Mandatory: false, MaxRepeat: 5,
	},
	{ // Segment group 6
		Group: []SchemaNode{
			{Segment: "CUX", Mandatory: true, MaxRepeat: 1},  // Currencies
			{Segment: "DTM", Mandatory: false, MaxRepeat: 5}, // Date/time/period
		},
		Mandatory: false, MaxRepeat: 20,
	},
	{ // Segment group 7
		Group: []SchemaNode{
			{Segment: "PYT", Mandatory: true, MaxRepeat: 1},  // Payment terms
			{Segment: "DTM", Mandatory: false, MaxRepeat: 5}, // Date/time/period
			{Segment: "PCD", Mandatory: false, MaxRepeat: 1}, // Percentage details
			{Segment: "MOA", Mandatory: false, MaxRepeat: 1}, // Monetary amount
		},
		Mandatory: false, MaxRepeat: 10,
	},
	{ // Segment group 8
		Group: []SchemaNode{
			{Segment: "TDT", Mandatory: true, MaxRepeat: 1},   // Transport information
			{Segment: "LOC", Mandatory: false, MaxRepeat: 10}, // Place/location identification
		},
		Mandatory: false, MaxRepeat: 10,
	},
	{ // Segment group 9
		Group: []SchemaNode{
			{Segment: "TOD", Mandatory: true, MaxRepeat: 1},  // Terms of delivery or transport
			{Segment: "LOC", Mandatory: false, MaxRepeat: 2}, // Place/location identification
		},
		Mandatory: false, MaxRepeat: 5,
	},
	{ // Segment group 10
		Group: []SchemaNode{
			{Segment: "ALC", Mandatory: true, MaxRepeat: 1},  // Allowance or charge
			{Segment: "ALI", Mandatory: false, MaxRepeat: 5}, // Additional information
			{Segment: "DTM", Mandatory: false, MaxRepeat: 9}, // Date/time/period
			{ // Segment group 11
				Group: []SchemaNode{
					{Segment: "QTY", Mandatory: true, MaxRepeat: 1},  // Quantity
					{Segment: "RNG", Mandatory: false, MaxRepeat: 1}, // Range details
				},
				Mandatory: false, MaxRepeat: 10,
			},
			{ // Segment group 12
				Group: []SchemaNode{
					{Segment: "PCD", Mandatory: true, MaxRepeat: 1},  // Percentage details
					{Segment: "RNG", Mandatory: false, MaxRepeat: 1}, // Range details
				},
				Mandatory: false, MaxRepeat: 10,
			},
			{ // Segment group 13
				Group: []SchemaNode{
					{Segment: "MOA", Mandatory: true, MaxRepeat: 1},  // Monetary amount
					{Segment: "RNG", Mandatory: false, MaxRepeat: 1}, // Range details
				},
				Mandatory: false, MaxRepeat: 10,
			},
			{ // Segment group 14
				Group: []SchemaNode{
					{Segment: "RTE", Mandatory: true, MaxRepeat: 1},  // Rate details
					{Segment: "RNG", Mandatory: false, MaxRepeat: 1}, // Range details
				},
				Mandatory: false, MaxRepeat: 10,
			},
			{ // Segment group 15
				Group: []SchemaNode{
					{Segment: "TAX", Mandatory: true, MaxRepeat: 1},  // Duty/tax/fee details
					{Segment: "MOA", Mandatory: false, MaxRepeat: 1}, // Monetary amount
					{Segment: "DTM", Mandatory: false, MaxRepeat: 1}, // Date/time/period
				},
				Mandatory: false, MaxRepeat: 5,
			},
		},
		Mandatory: false, MaxRepeat: 10,
	},
	{ // Segment group 16
		Group: []SchemaNode{
			{Segment: "TRU", Mandatory: true, MaxRepeat: 1},  // Technical rules
			{Segment: "DTM", Mandatory: false, MaxRepeat: 9}, // Date/time/period
		},
		Mandatory: false, MaxRepeat: 9,
	},
	{ // Segment group 17
		Group: []SchemaNode{
			{Segment: "PGI", Mandatory: true, MaxRepeat: 1},   // Product group information
			{Segment: "DTM", Mandatory: false, MaxRepeat: 15}, // Date/time/period
			{Segment: "QTY", Mandatory: false, MaxRepeat: 10}, // Quantity
			{Segment: "ALI", Mandatory: false, MaxRepeat: 5},  // Additional information
			{Segment: "FTX", Mandatory: false, MaxRepeat: 5},  // Free text
			{ // Segment group 18
				Group: []SchemaNode{
					{Segment: "CUX", Mandatory: true, MaxRepeat: 1},  // Currencies
					{Segment: "DTM", Mandatory: false, MaxRepeat: 5}, // Date/time/period
				},
				Mandatory: false, MaxRepeat: 10,
			},
			{ // Segment group 19
				Group: []SchemaNode{
					{Segment: "PRI", Mandatory: true, MaxRepeat: 1},   // Price details
					{Segment: "CUX", Mandatory: false, MaxRepeat: 1},  // Currencies
					{Segment: "APR", Mandatory: false, MaxRepeat: 99}, // Additional price information
					{Segment: "RNG", Mandatory: false, MaxRepeat: 1},  // Range details
					{Segment: "DTM", Mandatory: false, MaxRepeat: 5},  // Date/time/period
				},
				Mandatory: false, MaxRepeat: 100,
			},
			{ // Segment group 20
				Group: []SchemaNode{
					{Segment: "TAX", Mandatory: true, MaxRepeat: 1},  // Duty/tax/fee details
					{Segment: "MOA", Mandatory: false, MaxRepeat: 1}, // Monetary amount
					{Segment: "DTM", Mandatory: false, MaxRepeat: 1}, // Date/time/period
				},
				Mandatory: false, MaxRepeat: 10,
			},
			{ // Segment group 21
				Group: []SchemaNode{
					{Segment: "ALC", Mandatory: true, MaxRepeat: 1},  // Allowance or charge
					{Segment: "ALI", Mandatory: false, MaxRepeat: 5}, // Additional information
					{ // Segment group 22
						Group: []SchemaNode{
							{Segment: "QTY", Mandatory: true, MaxRepeat: 1},  // Quantity
							{Segment: "RNG", Mandatory: false, MaxRepeat: 1}, // Range details
						},
						Mandatory: false, MaxRepeat: 1,
					},
					{ // Segment group 23
						Group: []SchemaNode{
							{Segment: "PCD", Mandatory: true, MaxRepeat: 1},  // Percentage details
							{Segment: "RNG", Mandatory: false, MaxRepeat: 1}, // Range details
						},
						Mandatory: false, MaxRepeat: 1,
					},
					{ // Segment group 24
						Group: []SchemaNode{
							{Segment: "MOA", Mandatory: true, MaxRepeat: 1},  // Monetary amount
							{Segment: "RNG", Mandatory: false, MaxRepeat: 1}, // Range details
						},
						Mandatory: false, MaxRepeat: 2,
					},
					{ // Segment group 25
						Group: []SchemaNode{
							{Segment: "RTE", Mandatory: true, MaxRepeat: 1},  // Rate details
							{Segment: "RNG", Mandatory: false, MaxRepeat: 1}, // Range details
						},
						Mandatory: false, MaxRepeat: 1,
					},
					{ // Segment group 26
						Group: []SchemaNode{
							{Segment: "TAX", Mandatory: true, MaxRepeat: 1},  // Duty/tax/fee details
							{Segment: "MOA", Mandatory: false, MaxRepeat: 1}, // Monetary amount
							{Segment: "DTM", Mandatory: false, MaxRepeat: 1}, // Date/time/period
						},
						Mandatory: false, MaxRepeat: 5,
					},
				},
				Mandatory: false, MaxRepeat: 99,
			},
			{ // Segment group 27
				Group: []SchemaNode{
					{Segment: "NAD", Mandatory: true, MaxRepeat: 1},  // Name and address
					{Segment: "LOC", Mandatory: false, MaxRepeat: 5}, // Place/location identification
					{ // Segment group 28
						Group: []SchemaNode{
							{Segment: "RFF", Mandatory: true, MaxRepeat: 1},  // Reference
							{Segment: "DTM", Mandatory: false, MaxRepeat: 5}, // Date/time/period
						},
						Mandatory: false, MaxRepeat: 10,
					},
					{ // Segment group 29
						Group: []SchemaNode{
							{Segment: "CTA", Mandatory: true, MaxRepeat: 1},  // Contact information
							{Segment: "COM", Mandatory: false, MaxRepeat: 5}, // Communication contact
						},
						Mandatory: false, MaxRepeat: 5,
					},
				},
				Mandatory: false, MaxRepeat: 20,
			},
			{ // Segment group 30
				Group: []SchemaNode{
					{Segment: "PYT", Mandatory: true, MaxRepeat: 1},  // Payment terms
					{Segment: "DTM", Mandatory: false, MaxRepeat: 5}, // Date/time/period
					{Segment: "PCD", Mandatory: false, MaxRepeat: 1}, // Percentage details
					{Segment: "MOA", Mandatory: false, MaxRepeat: 1}, // Monetary amount
				},
				Mandatory: false, MaxRepeat: 10,
			},
			{ // Segment group 31
				Group: []SchemaNode{
					{Segment: "TDT", Mandatory: true, MaxRepeat: 1},   // Transport information
					{Segment: "LOC", Mandatory: false, MaxRepeat: 10}, // Place/location identification
				},
				Mandatory: false, MaxRepeat: 10,
			},
			{ // Segment group 32
				Group: []SchemaNode{
					{Segment: "TOD", Mandatory: true, MaxRepeat: 1},  // Terms of delivery or transport
					{Segment: "LOC", Mandatory: false, MaxRepeat: 2}, // Place/location identification
				},
				Mandatory: false, MaxRepeat: 5,
			},
			{ // Segment group 33
				Group: []SchemaNode{
					{Segment: "PAC", Mandatory: true, MaxRepeat: 1},   // Package
					{Segment: "MEA", Mandatory: false, MaxRepeat: 10}, // Measurements
					{Segment: "HAN", Mandatory: false, MaxRepeat: 5},  // Handling instructions
				},
				Mandatory: false, MaxRepeat: 10,
			},
			{ // Segment group 34
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
			{ // Segment group 35
				Group: []SchemaNode{
					{Segment: "TRU", Mandatory: true, MaxRepeat: 1},  // Technical rules
					{Segment: "DTM", Mandatory: false, MaxRepeat: 9}, // Date/time/period
				},
				Mandatory: false, MaxRepeat: 9,
			},
			{ // Segment group 36
				Group: []SchemaNode{
					{Segment: "LIN", Mandatory: true, MaxRepeat: 1},    // Line item
					{Segment: "PIA", Mandatory: false, MaxRepeat: 99},  // Additional product id
					{Segment: "IMD", Mandatory: false, MaxRepeat: 999}, // Item description
					{Segment: "MEA", Mandatory: false, MaxRepeat: 99},  // Measurements
					{Segment: "QTY", Mandatory: false, MaxRepeat: 10},  // Quantity
					{Segment: "HAN", Mandatory: false, MaxRepeat: 5},   // Handling instructions
					{Segment: "ALI", Mandatory: false, MaxRepeat: 5},   // Additional information
					{Segment: "DTM", Mandatory: false, MaxRepeat: 10},  // Date/time/period
					{Segment: "RFF", Mandatory: false, MaxRepeat: 99},  // Reference
					{Segment: "LOC", Mandatory: false, MaxRepeat: 99},  // Place/location identification
					{Segment: "DOC", Mandatory: false, MaxRepeat: 1},   // Document/message details
					{Segment: "PTY", Mandatory: false, MaxRepeat: 99},  // Priority
					{Segment: "FTX", Mandatory: false, MaxRepeat: 999}, // Free text
					{ // Segment group 37
						Group: []SchemaNode{
							{Segment: "CCI", Mandatory: true, MaxRepeat: 1},   // Characteristic/class id
							{Segment: "CAV", Mandatory: false, MaxRepeat: 10}, // Characteristic value
							{Segment: "MEA", Mandatory: false, MaxRepeat: 10}, // Measurements
						},
						Mandatory: false, MaxRepeat: 999,
					},
					{ // Segment group 38
						Group: []SchemaNode{
							{Segment: "TAX", Mandatory: true, MaxRepeat: 1},  // Duty/tax/fee details
							{Segment: "MOA", Mandatory: false, MaxRepeat: 1}, // Monetary amount
							{Segment: "DTM", Mandatory: false, MaxRepeat: 1}, // Date/time/period
						},
						Mandatory: false, MaxRepeat: 10,
					},
					{ // Segment group 39
						Group: []SchemaNode{
							{Segment: "CUX", Mandatory: true, MaxRepeat: 1},  // Currencies
							{Segment: "DTM", Mandatory: false, MaxRepeat: 5}, // Date/time/period
						},
						Mandatory: false, MaxRepeat: 5,
					},
					{ // Segment group 40
						Group: []SchemaNode{
							{Segment: "PRI", Mandatory: true, MaxRepeat: 1},   // Price details
							{Segment: "CUX", Mandatory: false, MaxRepeat: 1},  // Currencies
							{Segment: "APR", Mandatory: false, MaxRepeat: 99}, // Additional price information
							{Segment: "RNG", Mandatory: false, MaxRepeat: 1},  // Range details
							{Segment: "DTM", Mandatory: false, MaxRepeat: 5},  // Date/time/period
							{Segment: "PCD", Mandatory: false, MaxRepeat: 5},  // Percentage details
							{Segment: "LOC", Mandatory: false, MaxRepeat: 99}, // Place/location identification
						},
						Mandatory: false, MaxRepeat: 100,
					},
					{ // Segment group 41
						Group: []SchemaNode{
							{Segment: "ALC", Mandatory: true, MaxRepeat: 1},   // Allowance or charge
							{Segment: "ALI", Mandatory: false, MaxRepeat: 5},  // Additional information
							{Segment: "DTM", Mandatory: false, MaxRepeat: 9},  // Date/time/period
							{Segment: "LOC", Mandatory: false, MaxRepeat: 99}, // Place/location identification
							{ // Segment group 42
								Group: []SchemaNode{
									{Segment: "QTY", Mandatory: true, MaxRepeat: 1},  // Quantity
									{Segment: "RNG", Mandatory: false, MaxRepeat: 1}, // Range details
								},
								Mandatory: false, MaxRepeat: 10,
							},
							{ // Segment group 43
								Group: []SchemaNode{
									{Segment: "PCD", Mandatory: true, MaxRepeat: 1},  // Percentage details
									{Segment: "RNG", Mandatory: false, MaxRepeat: 1}, // Range details
								},
								Mandatory: false, MaxRepeat: 99,
							},
							{ // Segment group 44
								Group: []SchemaNode{
									{Segment: "MOA", Mandatory: true, MaxRepeat: 1},  // Monetary amount
									{Segment: "RNG", Mandatory: false, MaxRepeat: 1}, // Range details
								},
								Mandatory: false, MaxRepeat: 10,
							},
							{ // Segment group 45
								Group: []SchemaNode{
									{Segment: "RTE", Mandatory: true, MaxRepeat: 1},  // Rate details
									{Segment: "RNG", Mandatory: false, MaxRepeat: 1}, // Range details
								},
								Mandatory: false, MaxRepeat: 10,
							},
							{ // Segment group 46
								Group: []SchemaNode{
									{Segment: "TAX", Mandatory: true, MaxRepeat: 1},  // Duty/tax/fee details
									{Segment: "MOA", Mandatory: false, MaxRepeat: 1}, // Monetary amount
									{Segment: "DTM", Mandatory: false, MaxRepeat: 1}, // Date/time/period
								},
								Mandatory: false, MaxRepeat: 5,
							},
						},
						Mandatory: false, MaxRepeat: 99,
					},
					{ // Segment group 47
						Group: []SchemaNode{
							{Segment: "PAC", Mandatory: true, MaxRepeat: 1},   // Package
							{Segment: "MEA", Mandatory: false, MaxRepeat: 10}, // Measurements
							{Segment: "QTY", Mandatory: false, MaxRepeat: 99}, // Quantity
							{Segment: "HAN", Mandatory: false, MaxRepeat: 5},  // Handling instructions
						},
						Mandatory: false, MaxRepeat: 999,
					},
					{ // Segment group 48
						Group: []SchemaNode{
							{Segment: "PYT", Mandatory: true, MaxRepeat: 1},  // Payment terms
							{Segment: "DTM", Mandatory: false, MaxRepeat: 5}, // Date/time/period
							{Segment: "PCD", Mandatory: false, MaxRepeat: 1}, // Percentage details
							{Segment: "MOA", Mandatory: false, MaxRepeat: 1}, // Monetary amount
						},
						Mandatory: false, MaxRepeat: 10,
					},
					{ // Segment group 49
						Group: []SchemaNode{
							{Segment: "TDT", Mandatory: true, MaxRepeat: 1},   // Transport information
							{Segment: "LOC", Mandatory: false, MaxRepeat: 10}, // Place/location identification
						},
						Mandatory: false, MaxRepeat: 10,
					},
					{ // Segment group 50
						Group: []SchemaNode{
							{Segment: "TOD", Mandatory: true, MaxRepeat: 1},  // Terms of delivery or transport
							{Segment: "LOC", Mandatory: false, MaxRepeat: 2}, // Place/location identification
						},
						Mandatory: false, MaxRepeat: 5,
					},
					{ // Segment group 51
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
					{ // Segment group 52
						Group: []SchemaNode{
							{Segment: "NAD", Mandatory: true, MaxRepeat: 1}, // Name and address
							{ // Segment group 53
								Group: []SchemaNode{
									{Segment: "CTA", Mandatory: true, MaxRepeat: 1},   // Contact information
									{Segment: "COM", Mandatory: false, MaxRepeat: 99}, // Communication contact
								},
								Mandatory: false, MaxRepeat: 99,
							},
						},
						Mandatory: false, MaxRepeat: 99,
					},
					{ // Segment group 54
						Group: []SchemaNode{
							{Segment: "TRU", Mandatory: true, MaxRepeat: 1},  // Technical rules
							{Segment: "DTM", Mandatory: false, MaxRepeat: 9}, // Date/time/period
						},
						Mandatory: false, MaxRepeat: 9,
					},
					{ // Segment group 55
						Group: []SchemaNode{
							{Segment: "DGS", Mandatory: true, MaxRepeat: 1},    // Dangerous goods
							{Segment: "QTY", Mandatory: false, MaxRepeat: 9},   // Quantity
							{Segment: "FTX", Mandatory: false, MaxRepeat: 999}, // Free text
						},
						Mandatory: false, MaxRepeat: 9,
					},
					{ // Segment group 56
						Group: []SchemaNode{
							{Segment: "HYN", Mandatory: true, MaxRepeat: 1},   // Hierarchy information
							{Segment: "PIA", Mandatory: false, MaxRepeat: 99}, // Additional product id
							{Segment: "QTY", Mandatory: false, MaxRepeat: 9},  // Quantity
							{Segment: "FTX", Mandatory: false, MaxRepeat: 99}, // Free text
							{ // Segment group 57
								Group: []SchemaNode{
									{Segment: "RFF", Mandatory: true, MaxRepeat: 1},  // Reference
									{Segment: "DTM", Mandatory: false, MaxRepeat: 9}, // Date/time/period
								},
								Mandatory: false, MaxRepeat: 999,
							},
							{ // Segment group 58
								Group: []SchemaNode{
									{Segment: "CCI", Mandatory: true, MaxRepeat: 1},   // Characteristic/class id
									{Segment: "CAV", Mandatory: false, MaxRepeat: 9},  // Characteristic value
									{Segment: "MEA", Mandatory: false, MaxRepeat: 99}, // Measurements
								},
								Mandatory: false, MaxRepeat: 99,
							},
							{ // Segment group 59
								Group: []SchemaNode{
									{Segment: "NAD", Mandatory: true, MaxRepeat: 1},   // Name and address
									{Segment: "PIA", Mandatory: false, MaxRepeat: 99}, // Additional product id
									{Segment: "QTY", Mandatory: false, MaxRepeat: 9},  // Quantity
									{ // Segment group 60
										Group: []SchemaNode{
											{Segment: "CCI", Mandatory: true, MaxRepeat: 1},   // Characteristic/class id
											{Segment: "CAV", Mandatory: false, MaxRepeat: 99}, // Characteristic value
											{Segment: "MEA", Mandatory: false, MaxRepeat: 9},  // Measurements
										},
										Mandatory: false, MaxRepeat: 99,
									},
								},
								Mandatory: false, MaxRepeat: 99,
							},
						},
						Mandatory: false, MaxRepeat: 999,
					},
				},
				Mandatory: false, MaxRepeat: 999999,
			},
		},
		Mandatory: false, MaxRepeat: 1000,
	},
}}
