package edifact

// INVOIC D.20A message specification: the structural (segment/group
// presence, order, cardinality) branching diagram for the Invoice
// message, UN/EDIFACT directory release D.20A. Scope is structural only --
// not element-level content -- per edifact-ls-3uzr's epic non-goal.
// 55 segment groups, max nesting depth 3.
//
// Source: https://service.unece.org/trade/untdid/d20a/trmd/invoic_c.htm
// section 4.3.1 "Segment table". That URL currently returns HTTP 403 from
// Cloudflare when fetched directly; this data was transcribed from the
// Wayback Machine's archived copy instead:
// http://web.archive.org/web/20240518231218/https://service.unece.org/trade/untdid/d20a/trmd/invoic_c.htm
// -- re-check the direct URL first if cross-checking later, in case the
// block has lifted.
//
// Transcription note: generated the same way as iftmcs_d21a.go -- a
// one-off script parsed the source's exact rail-art column positions
// mechanically and verified the result balances (all 55 groups
// open and close consistently) before being transcribed here, rather than
// reading the ASCII nesting by eye.

func init() {
	RegisterSchema(
		MessageID{Type: "INVOIC", Version: "D", Release: "20A", Agency: "UN"},
		invoicD20aSchema,
	)
}

var invoicD20aSchema = Schema{Nodes: []SchemaNode{
	{Segment: "BGM", Mandatory: true, MaxRepeat: 1},   // Beginning of message
	{Segment: "DTM", Mandatory: true, MaxRepeat: 35},  // Date/time/period
	{Segment: "PAI", Mandatory: false, MaxRepeat: 1},  // Payment instructions
	{Segment: "ALI", Mandatory: false, MaxRepeat: 5},  // Additional information
	{Segment: "IMD", Mandatory: false, MaxRepeat: 1},  // Item description
	{Segment: "FTX", Mandatory: false, MaxRepeat: 99}, // Free text
	{Segment: "LOC", Mandatory: false, MaxRepeat: 10}, // Place/location identification
	{Segment: "GEI", Mandatory: false, MaxRepeat: 10}, // Processing information
	{Segment: "DGS", Mandatory: false, MaxRepeat: 1},  // Dangerous goods
	{Segment: "GIR", Mandatory: false, MaxRepeat: 10}, // Related identification numbers
	{ // Segment group 1
		Group: []SchemaNode{
			{Segment: "RFF", Mandatory: true, MaxRepeat: 1},   // Reference
			{Segment: "DTM", Mandatory: false, MaxRepeat: 5},  // Date/time/period
			{Segment: "GIR", Mandatory: false, MaxRepeat: 5},  // Related identification numbers
			{Segment: "LOC", Mandatory: false, MaxRepeat: 2},  // Place/location identification
			{Segment: "MEA", Mandatory: false, MaxRepeat: 5},  // Measurements
			{Segment: "QTY", Mandatory: false, MaxRepeat: 2},  // Quantity
			{Segment: "FTX", Mandatory: false, MaxRepeat: 5},  // Free text
			{Segment: "MOA", Mandatory: false, MaxRepeat: 2},  // Monetary amount
			{Segment: "RTE", Mandatory: false, MaxRepeat: 99}, // Rate details
		},
		Mandatory: false, MaxRepeat: 99999,
	},
	{ // Segment group 2
		Group: []SchemaNode{
			{Segment: "NAD", Mandatory: true, MaxRepeat: 1},   // Name and address
			{Segment: "LOC", Mandatory: false, MaxRepeat: 25}, // Place/location identification
			{Segment: "FII", Mandatory: false, MaxRepeat: 5},  // Financial institution information
			{Segment: "MOA", Mandatory: false, MaxRepeat: 99}, // Monetary amount
			{ // Segment group 3
				Group: []SchemaNode{
					{Segment: "RFF", Mandatory: true, MaxRepeat: 1},  // Reference
					{Segment: "DTM", Mandatory: false, MaxRepeat: 5}, // Date/time/period
				},
				Mandatory: false, MaxRepeat: 9999,
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
		Mandatory: false, MaxRepeat: 99,
	},
	{ // Segment group 6
		Group: []SchemaNode{
			{Segment: "TAX", Mandatory: true, MaxRepeat: 1},  // Duty/tax/fee details
			{Segment: "MOA", Mandatory: false, MaxRepeat: 1}, // Monetary amount
			{Segment: "LOC", Mandatory: false, MaxRepeat: 5}, // Place/location identification
		},
		Mandatory: false, MaxRepeat: 5,
	},
	{ // Segment group 7
		Group: []SchemaNode{
			{Segment: "CUX", Mandatory: true, MaxRepeat: 1},  // Currencies
			{Segment: "DTM", Mandatory: false, MaxRepeat: 5}, // Date/time/period
		},
		Mandatory: false, MaxRepeat: 99,
	},
	{ // Segment group 8
		Group: []SchemaNode{
			{Segment: "PYT", Mandatory: true, MaxRepeat: 1},  // Payment terms
			{Segment: "DTM", Mandatory: false, MaxRepeat: 5}, // Date/time/period
			{Segment: "PCD", Mandatory: false, MaxRepeat: 1}, // Percentage details
			{Segment: "MOA", Mandatory: false, MaxRepeat: 1}, // Monetary amount
			{Segment: "PAI", Mandatory: false, MaxRepeat: 1}, // Payment instructions
			{Segment: "FII", Mandatory: false, MaxRepeat: 1}, // Financial institution information
		},
		Mandatory: false, MaxRepeat: 10,
	},
	{ // Segment group 9
		Group: []SchemaNode{
			{Segment: "TDT", Mandatory: true, MaxRepeat: 1},  // Transport information
			{Segment: "TSR", Mandatory: false, MaxRepeat: 1}, // Transport service requirements
			{ // Segment group 10
				Group: []SchemaNode{
					{Segment: "LOC", Mandatory: true, MaxRepeat: 1},  // Place/location identification
					{Segment: "DTM", Mandatory: false, MaxRepeat: 5}, // Date/time/period
				},
				Mandatory: false, MaxRepeat: 10,
			},
			{ // Segment group 11
				Group: []SchemaNode{
					{Segment: "RFF", Mandatory: true, MaxRepeat: 1},  // Reference
					{Segment: "DTM", Mandatory: false, MaxRepeat: 5}, // Date/time/period
				},
				Mandatory: false, MaxRepeat: 9999,
			},
		},
		Mandatory: false, MaxRepeat: 10,
	},
	{ // Segment group 12
		Group: []SchemaNode{
			{Segment: "TOD", Mandatory: true, MaxRepeat: 1},  // Terms of delivery or transport
			{Segment: "LOC", Mandatory: false, MaxRepeat: 2}, // Place/location identification
		},
		Mandatory: false, MaxRepeat: 5,
	},
	{ // Segment group 13
		Group: []SchemaNode{
			{Segment: "EQD", Mandatory: true, MaxRepeat: 1},  // Equipment details
			{Segment: "SEL", Mandatory: false, MaxRepeat: 9}, // Seal number
		},
		Mandatory: false, MaxRepeat: 99,
	},
	{ // Segment group 14
		Group: []SchemaNode{
			{Segment: "PAC", Mandatory: true, MaxRepeat: 1},  // Package
			{Segment: "MEA", Mandatory: false, MaxRepeat: 5}, // Measurements
			{Segment: "EQD", Mandatory: false, MaxRepeat: 1}, // Equipment details
			{ // Segment group 15
				Group: []SchemaNode{
					{Segment: "PCI", Mandatory: true, MaxRepeat: 1},  // Package identification
					{Segment: "RFF", Mandatory: false, MaxRepeat: 1}, // Reference
					{Segment: "DTM", Mandatory: false, MaxRepeat: 5}, // Date/time/period
					{Segment: "GIN", Mandatory: false, MaxRepeat: 5}, // Goods identity number
				},
				Mandatory: false, MaxRepeat: 5,
			},
		},
		Mandatory: false, MaxRepeat: 1000,
	},
	{ // Segment group 16
		Group: []SchemaNode{
			{Segment: "ALC", Mandatory: true, MaxRepeat: 1},  // Allowance or charge
			{Segment: "ALI", Mandatory: false, MaxRepeat: 5}, // Additional information
			{Segment: "FTX", Mandatory: false, MaxRepeat: 1}, // Free text
			{ // Segment group 17
				Group: []SchemaNode{
					{Segment: "RFF", Mandatory: true, MaxRepeat: 1},  // Reference
					{Segment: "DTM", Mandatory: false, MaxRepeat: 5}, // Date/time/period
				},
				Mandatory: false, MaxRepeat: 5,
			},
			{ // Segment group 18
				Group: []SchemaNode{
					{Segment: "QTY", Mandatory: true, MaxRepeat: 1},  // Quantity
					{Segment: "RNG", Mandatory: false, MaxRepeat: 1}, // Range details
				},
				Mandatory: false, MaxRepeat: 1,
			},
			{ // Segment group 19
				Group: []SchemaNode{
					{Segment: "PCD", Mandatory: true, MaxRepeat: 1},  // Percentage details
					{Segment: "RNG", Mandatory: false, MaxRepeat: 1}, // Range details
				},
				Mandatory: false, MaxRepeat: 1,
			},
			{ // Segment group 20
				Group: []SchemaNode{
					{Segment: "MOA", Mandatory: true, MaxRepeat: 1},  // Monetary amount
					{Segment: "RNG", Mandatory: false, MaxRepeat: 1}, // Range details
					{Segment: "CUX", Mandatory: false, MaxRepeat: 1}, // Currencies
					{Segment: "DTM", Mandatory: false, MaxRepeat: 1}, // Date/time/period
				},
				Mandatory: false, MaxRepeat: 2,
			},
			{ // Segment group 21
				Group: []SchemaNode{
					{Segment: "RTE", Mandatory: true, MaxRepeat: 1},  // Rate details
					{Segment: "RNG", Mandatory: false, MaxRepeat: 1}, // Range details
				},
				Mandatory: false, MaxRepeat: 1,
			},
			{ // Segment group 22
				Group: []SchemaNode{
					{Segment: "TAX", Mandatory: true, MaxRepeat: 1},  // Duty/tax/fee details
					{Segment: "MOA", Mandatory: false, MaxRepeat: 1}, // Monetary amount
				},
				Mandatory: false, MaxRepeat: 5,
			},
		},
		Mandatory: false, MaxRepeat: 9999,
	},
	{ // Segment group 23
		Group: []SchemaNode{
			{Segment: "RCS", Mandatory: true, MaxRepeat: 1},  // Requirements and conditions
			{Segment: "RFF", Mandatory: false, MaxRepeat: 5}, // Reference
			{Segment: "DTM", Mandatory: false, MaxRepeat: 5}, // Date/time/period
			{Segment: "FTX", Mandatory: false, MaxRepeat: 5}, // Free text
		},
		Mandatory: false, MaxRepeat: 100,
	},
	{ // Segment group 24
		Group: []SchemaNode{
			{Segment: "AJT", Mandatory: true, MaxRepeat: 1},  // Adjustment details
			{Segment: "FTX", Mandatory: false, MaxRepeat: 5}, // Free text
		},
		Mandatory: false, MaxRepeat: 1,
	},
	{ // Segment group 25
		Group: []SchemaNode{
			{Segment: "INP", Mandatory: true, MaxRepeat: 1},  // Parties and instruction
			{Segment: "FTX", Mandatory: false, MaxRepeat: 5}, // Free text
		},
		Mandatory: false, MaxRepeat: 1,
	},
	{ // Segment group 26
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
	{ // Segment group 27
		Group: []SchemaNode{
			{Segment: "LIN", Mandatory: true, MaxRepeat: 1},     // Line item
			{Segment: "PIA", Mandatory: false, MaxRepeat: 25},   // Additional product id
			{Segment: "PGI", Mandatory: false, MaxRepeat: 99},   // Product group information
			{Segment: "IMD", Mandatory: false, MaxRepeat: 99},   // Item description
			{Segment: "MEA", Mandatory: false, MaxRepeat: 5},    // Measurements
			{Segment: "QTY", Mandatory: false, MaxRepeat: 5},    // Quantity
			{Segment: "PCD", Mandatory: false, MaxRepeat: 1},    // Percentage details
			{Segment: "ALI", Mandatory: false, MaxRepeat: 5},    // Additional information
			{Segment: "DTM", Mandatory: false, MaxRepeat: 35},   // Date/time/period
			{Segment: "GIN", Mandatory: false, MaxRepeat: 1000}, // Goods identity number
			{Segment: "GIR", Mandatory: false, MaxRepeat: 1000}, // Related identification numbers
			{Segment: "QVR", Mandatory: false, MaxRepeat: 1},    // Quantity variances
			{Segment: "EQD", Mandatory: false, MaxRepeat: 1},    // Equipment details
			{Segment: "FTX", Mandatory: false, MaxRepeat: 99},   // Free text
			{Segment: "DGS", Mandatory: false, MaxRepeat: 1},    // Dangerous goods
			{ // Segment group 28
				Group: []SchemaNode{
					{Segment: "MOA", Mandatory: true, MaxRepeat: 1},  // Monetary amount
					{Segment: "CUX", Mandatory: false, MaxRepeat: 1}, // Currencies
				},
				Mandatory: false, MaxRepeat: 99,
			},
			{ // Segment group 29
				Group: []SchemaNode{
					{Segment: "PYT", Mandatory: true, MaxRepeat: 1},   // Payment terms
					{Segment: "DTM", Mandatory: false, MaxRepeat: 5},  // Date/time/period
					{Segment: "PCD", Mandatory: false, MaxRepeat: 99}, // Percentage details
					{Segment: "MOA", Mandatory: false, MaxRepeat: 1},  // Monetary amount
				},
				Mandatory: false, MaxRepeat: 10,
			},
			{ // Segment group 30
				Group: []SchemaNode{
					{Segment: "PRI", Mandatory: true, MaxRepeat: 1},  // Price details
					{Segment: "CUX", Mandatory: false, MaxRepeat: 1}, // Currencies
					{Segment: "APR", Mandatory: false, MaxRepeat: 1}, // Additional price information
					{Segment: "RNG", Mandatory: false, MaxRepeat: 1}, // Range details
					{Segment: "DTM", Mandatory: false, MaxRepeat: 5}, // Date/time/period
				},
				Mandatory: false, MaxRepeat: 25,
			},
			{ // Segment group 31
				Group: []SchemaNode{
					{Segment: "RFF", Mandatory: true, MaxRepeat: 1},  // Reference
					{Segment: "DTM", Mandatory: false, MaxRepeat: 5}, // Date/time/period
				},
				Mandatory: false, MaxRepeat: 10,
			},
			{ // Segment group 32
				Group: []SchemaNode{
					{Segment: "PAC", Mandatory: true, MaxRepeat: 1},   // Package
					{Segment: "MEA", Mandatory: false, MaxRepeat: 10}, // Measurements
					{Segment: "EQD", Mandatory: false, MaxRepeat: 1},  // Equipment details
					{ // Segment group 33
						Group: []SchemaNode{
							{Segment: "PCI", Mandatory: true, MaxRepeat: 1},   // Package identification
							{Segment: "RFF", Mandatory: false, MaxRepeat: 1},  // Reference
							{Segment: "DTM", Mandatory: false, MaxRepeat: 5},  // Date/time/period
							{Segment: "GIN", Mandatory: false, MaxRepeat: 10}, // Goods identity number
						},
						Mandatory: false, MaxRepeat: 10,
					},
				},
				Mandatory: false, MaxRepeat: 999,
			},
			{ // Segment group 34
				Group: []SchemaNode{
					{Segment: "LOC", Mandatory: true, MaxRepeat: 1},    // Place/location identification
					{Segment: "QTY", Mandatory: false, MaxRepeat: 100}, // Quantity
					{Segment: "DTM", Mandatory: false, MaxRepeat: 5},   // Date/time/period
				},
				Mandatory: false, MaxRepeat: 9999,
			},
			{ // Segment group 35
				Group: []SchemaNode{
					{Segment: "TAX", Mandatory: true, MaxRepeat: 1},  // Duty/tax/fee details
					{Segment: "MOA", Mandatory: false, MaxRepeat: 2}, // Monetary amount
					{Segment: "LOC", Mandatory: false, MaxRepeat: 5}, // Place/location identification
				},
				Mandatory: false, MaxRepeat: 99,
			},
			{ // Segment group 36
				Group: []SchemaNode{
					{Segment: "NAD", Mandatory: true, MaxRepeat: 1},  // Name and address
					{Segment: "LOC", Mandatory: false, MaxRepeat: 5}, // Place/location identification
					{Segment: "FII", Mandatory: false, MaxRepeat: 5}, // Financial institution information
					{ // Segment group 37
						Group: []SchemaNode{
							{Segment: "RFF", Mandatory: true, MaxRepeat: 1},  // Reference
							{Segment: "DTM", Mandatory: false, MaxRepeat: 5}, // Date/time/period
						},
						Mandatory: false, MaxRepeat: 5,
					},
					{ // Segment group 38
						Group: []SchemaNode{
							{Segment: "DOC", Mandatory: true, MaxRepeat: 1},  // Document/message details
							{Segment: "DTM", Mandatory: false, MaxRepeat: 5}, // Date/time/period
						},
						Mandatory: false, MaxRepeat: 5,
					},
					{ // Segment group 39
						Group: []SchemaNode{
							{Segment: "CTA", Mandatory: true, MaxRepeat: 1},  // Contact information
							{Segment: "COM", Mandatory: false, MaxRepeat: 5}, // Communication contact
						},
						Mandatory: false, MaxRepeat: 5,
					},
				},
				Mandatory: false, MaxRepeat: 99,
			},
			{ // Segment group 40
				Group: []SchemaNode{
					{Segment: "ALC", Mandatory: true, MaxRepeat: 1},  // Allowance or charge
					{Segment: "ALI", Mandatory: false, MaxRepeat: 5}, // Additional information
					{Segment: "DTM", Mandatory: false, MaxRepeat: 5}, // Date/time/period
					{Segment: "FTX", Mandatory: false, MaxRepeat: 1}, // Free text
					{ // Segment group 41
						Group: []SchemaNode{
							{Segment: "QTY", Mandatory: true, MaxRepeat: 1},  // Quantity
							{Segment: "RNG", Mandatory: false, MaxRepeat: 1}, // Range details
						},
						Mandatory: false, MaxRepeat: 1,
					},
					{ // Segment group 42
						Group: []SchemaNode{
							{Segment: "PCD", Mandatory: true, MaxRepeat: 1},  // Percentage details
							{Segment: "RNG", Mandatory: false, MaxRepeat: 1}, // Range details
						},
						Mandatory: false, MaxRepeat: 1,
					},
					{ // Segment group 43
						Group: []SchemaNode{
							{Segment: "MOA", Mandatory: true, MaxRepeat: 1},  // Monetary amount
							{Segment: "RNG", Mandatory: false, MaxRepeat: 1}, // Range details
							{Segment: "CUX", Mandatory: false, MaxRepeat: 1}, // Currencies
							{Segment: "DTM", Mandatory: false, MaxRepeat: 1}, // Date/time/period
						},
						Mandatory: false, MaxRepeat: 2,
					},
					{ // Segment group 44
						Group: []SchemaNode{
							{Segment: "RTE", Mandatory: true, MaxRepeat: 1},  // Rate details
							{Segment: "RNG", Mandatory: false, MaxRepeat: 1}, // Range details
						},
						Mandatory: false, MaxRepeat: 1,
					},
					{ // Segment group 45
						Group: []SchemaNode{
							{Segment: "TAX", Mandatory: true, MaxRepeat: 1},  // Duty/tax/fee details
							{Segment: "MOA", Mandatory: false, MaxRepeat: 2}, // Monetary amount
						},
						Mandatory: false, MaxRepeat: 5,
					},
				},
				Mandatory: false, MaxRepeat: 30,
			},
			{ // Segment group 46
				Group: []SchemaNode{
					{Segment: "TDT", Mandatory: true, MaxRepeat: 1}, // Transport information
					{ // Segment group 47
						Group: []SchemaNode{
							{Segment: "LOC", Mandatory: true, MaxRepeat: 1},  // Place/location identification
							{Segment: "DTM", Mandatory: false, MaxRepeat: 5}, // Date/time/period
						},
						Mandatory: false, MaxRepeat: 10,
					},
				},
				Mandatory: false, MaxRepeat: 10,
			},
			{ // Segment group 48
				Group: []SchemaNode{
					{Segment: "TOD", Mandatory: true, MaxRepeat: 1},  // Terms of delivery or transport
					{Segment: "LOC", Mandatory: false, MaxRepeat: 2}, // Place/location identification
				},
				Mandatory: false, MaxRepeat: 5,
			},
			{ // Segment group 49
				Group: []SchemaNode{
					{Segment: "RCS", Mandatory: true, MaxRepeat: 1},  // Requirements and conditions
					{Segment: "RFF", Mandatory: false, MaxRepeat: 5}, // Reference
					{Segment: "DTM", Mandatory: false, MaxRepeat: 5}, // Date/time/period
					{Segment: "FTX", Mandatory: false, MaxRepeat: 5}, // Free text
				},
				Mandatory: false, MaxRepeat: 100,
			},
			{ // Segment group 50
				Group: []SchemaNode{
					{Segment: "GEI", Mandatory: true, MaxRepeat: 1},   // Processing information
					{Segment: "RFF", Mandatory: false, MaxRepeat: 9},  // Reference
					{Segment: "IMD", Mandatory: false, MaxRepeat: 99}, // Item description
					{Segment: "DTM", Mandatory: false, MaxRepeat: 5},  // Date/time/period
					{Segment: "GIR", Mandatory: false, MaxRepeat: 5},  // Related identification numbers
					{Segment: "LOC", Mandatory: false, MaxRepeat: 2},  // Place/location identification
					{Segment: "MEA", Mandatory: false, MaxRepeat: 5},  // Measurements
					{Segment: "QTY", Mandatory: false, MaxRepeat: 9},  // Quantity
					{Segment: "FTX", Mandatory: false, MaxRepeat: 5},  // Free text
					{Segment: "MOA", Mandatory: false, MaxRepeat: 2},  // Monetary amount
				},
				Mandatory: false, MaxRepeat: 999,
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
		},
		Mandatory: false, MaxRepeat: 9999999,
	},
	{Segment: "UNS", Mandatory: true, MaxRepeat: 1},   // Section control
	{Segment: "CNT", Mandatory: false, MaxRepeat: 10}, // Control total
	{ // Segment group 52
		Group: []SchemaNode{
			{Segment: "MOA", Mandatory: true, MaxRepeat: 1}, // Monetary amount
			{ // Segment group 53
				Group: []SchemaNode{
					{Segment: "RFF", Mandatory: true, MaxRepeat: 1},  // Reference
					{Segment: "DTM", Mandatory: false, MaxRepeat: 5}, // Date/time/period
				},
				Mandatory: false, MaxRepeat: 1,
			},
		},
		Mandatory: true, MaxRepeat: 100,
	},
	{ // Segment group 54
		Group: []SchemaNode{
			{Segment: "TAX", Mandatory: true, MaxRepeat: 1},  // Duty/tax/fee details
			{Segment: "MOA", Mandatory: false, MaxRepeat: 9}, // Monetary amount
		},
		Mandatory: false, MaxRepeat: 10,
	},
	{ // Segment group 55
		Group: []SchemaNode{
			{Segment: "ALC", Mandatory: true, MaxRepeat: 1},  // Allowance or charge
			{Segment: "ALI", Mandatory: false, MaxRepeat: 1}, // Additional information
			{Segment: "MOA", Mandatory: false, MaxRepeat: 2}, // Monetary amount
			{Segment: "FTX", Mandatory: false, MaxRepeat: 1}, // Free text
		},
		Mandatory: false, MaxRepeat: 15,
	},
}}
