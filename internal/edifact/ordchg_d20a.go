package edifact

// ORDCHG D.20A message specification: the structural (segment/group
// presence, order, cardinality) branching diagram for this message,
// UN/EDIFACT directory release D.20A. Scope is structural only -- not
// element-level content -- per edifact-ls-3uzr's epic non-goal.
// 60 segment groups, max nesting depth 3.
//
// Source: https://service.unece.org/trade/untdid/d20a/trmd/ordchg_c.htm
// section 4.3.1 "Segment table". That URL currently returns HTTP 403 from
// Cloudflare when fetched directly; this data was transcribed from the
// Wayback Machine's archived copy instead:
// http://web.archive.org/web/20231002203159/https://service.unece.org/trade/untdid/d20a/trmd/ordchg_c.htm
// -- re-check the direct URL first if cross-checking later, in case the
// block has lifted.
//
// Transcription note: generated in bulk (edifact-ls-13gu) the same way
// as every other message-type schema -- a script parsed the source's
// exact rail-art column positions mechanically and verified the result
// balances (all 60 groups open and close consistently)
// before being transcribed here, rather than reading the ASCII nesting
// by eye. Not individually spot-checked against the raw source the way
// the first several message types in this project were -- see
// edifact-ls-13gu for which representative samples were spot-checked
// for this batch.

func init() {
	RegisterSchema(
		MessageID{Type: "ORDCHG", Version: "D", Release: "20A", Agency: "UN"},
		ordchgD20aSchema,
		"https://service.unece.org/trade/untdid/d20a/trmd/ordchg_c.htm",
	)
}

var ordchgD20aSchema = Schema{Nodes: []SchemaNode{
	{Segment: "BGM", Mandatory: true, MaxRepeat: 1},    // Beginning of message
	{Segment: "DTM", Mandatory: true, MaxRepeat: 35},   // Date/time/period
	{Segment: "PAI", Mandatory: false, MaxRepeat: 1},   // Payment instructions
	{Segment: "ALI", Mandatory: false, MaxRepeat: 5},   // Additional information
	{Segment: "IMD", Mandatory: false, MaxRepeat: 999}, // Item description
	{Segment: "FTX", Mandatory: false, MaxRepeat: 99},  // Free text
	{Segment: "GIR", Mandatory: false, MaxRepeat: 10},  // Related identification numbers
	{ // Segment group 1
		Group: []SchemaNode{
			{Segment: "RFF", Mandatory: true, MaxRepeat: 1},  // Reference
			{Segment: "DTM", Mandatory: false, MaxRepeat: 5}, // Date/time/period
		},
		Mandatory: false, MaxRepeat: 9999,
	},
	{ // Segment group 2
		Group: []SchemaNode{
			{Segment: "AJT", Mandatory: true, MaxRepeat: 1},  // Adjustment details
			{Segment: "FTX", Mandatory: false, MaxRepeat: 5}, // Free text
		},
		Mandatory: false, MaxRepeat: 1,
	},
	{ // Segment group 3
		Group: []SchemaNode{
			{Segment: "NAD", Mandatory: true, MaxRepeat: 1},   // Name and address
			{Segment: "LOC", Mandatory: false, MaxRepeat: 99}, // Place/location identification
			{Segment: "FII", Mandatory: false, MaxRepeat: 5},  // Financial institution information
			{ // Segment group 4
				Group: []SchemaNode{
					{Segment: "RFF", Mandatory: true, MaxRepeat: 1},  // Reference
					{Segment: "DTM", Mandatory: false, MaxRepeat: 5}, // Date/time/period
				},
				Mandatory: false, MaxRepeat: 99,
			},
			{ // Segment group 5
				Group: []SchemaNode{
					{Segment: "DOC", Mandatory: true, MaxRepeat: 1},  // Document/message details
					{Segment: "DTM", Mandatory: false, MaxRepeat: 5}, // Date/time/period
				},
				Mandatory: false, MaxRepeat: 5,
			},
			{ // Segment group 6
				Group: []SchemaNode{
					{Segment: "CTA", Mandatory: true, MaxRepeat: 1},  // Contact information
					{Segment: "COM", Mandatory: false, MaxRepeat: 5}, // Communication contact
				},
				Mandatory: false, MaxRepeat: 5,
			},
		},
		Mandatory: false, MaxRepeat: 99,
	},
	{ // Segment group 7
		Group: []SchemaNode{
			{Segment: "TAX", Mandatory: true, MaxRepeat: 1},  // Duty/tax/fee details
			{Segment: "MOA", Mandatory: false, MaxRepeat: 1}, // Monetary amount
			{Segment: "LOC", Mandatory: false, MaxRepeat: 5}, // Place/location identification
		},
		Mandatory: false, MaxRepeat: 5,
	},
	{ // Segment group 8
		Group: []SchemaNode{
			{Segment: "CUX", Mandatory: true, MaxRepeat: 1},  // Currencies
			{Segment: "PCD", Mandatory: false, MaxRepeat: 5}, // Percentage details
			{Segment: "DTM", Mandatory: false, MaxRepeat: 5}, // Date/time/period
		},
		Mandatory: false, MaxRepeat: 5,
	},
	{ // Segment group 9
		Group: []SchemaNode{
			{Segment: "PYT", Mandatory: true, MaxRepeat: 1},  // Payment terms
			{Segment: "DTM", Mandatory: false, MaxRepeat: 5}, // Date/time/period
			{Segment: "PCD", Mandatory: false, MaxRepeat: 1}, // Percentage details
			{ // Segment group 10
				Group: []SchemaNode{
					{Segment: "MOA", Mandatory: true, MaxRepeat: 1},   // Monetary amount
					{Segment: "GIR", Mandatory: false, MaxRepeat: 9},  // Related identification numbers
					{Segment: "RJL", Mandatory: false, MaxRepeat: 99}, // Accounting journal identification
				},
				Mandatory: false, MaxRepeat: 9999,
			},
		},
		Mandatory: false, MaxRepeat: 10,
	},
	{ // Segment group 11
		Group: []SchemaNode{
			{Segment: "TDT", Mandatory: true, MaxRepeat: 1}, // Transport information
			{ // Segment group 12
				Group: []SchemaNode{
					{Segment: "LOC", Mandatory: true, MaxRepeat: 1},  // Place/location identification
					{Segment: "DTM", Mandatory: false, MaxRepeat: 5}, // Date/time/period
				},
				Mandatory: false, MaxRepeat: 10,
			},
		},
		Mandatory: false, MaxRepeat: 10,
	},
	{ // Segment group 13
		Group: []SchemaNode{
			{Segment: "TOD", Mandatory: true, MaxRepeat: 1},  // Terms of delivery or transport
			{Segment: "LOC", Mandatory: false, MaxRepeat: 2}, // Place/location identification
		},
		Mandatory: false, MaxRepeat: 5,
	},
	{ // Segment group 14
		Group: []SchemaNode{
			{Segment: "PAC", Mandatory: true, MaxRepeat: 1},  // Package
			{Segment: "MEA", Mandatory: false, MaxRepeat: 5}, // Measurements
			{ // Segment group 15
				Group: []SchemaNode{
					{Segment: "PCI", Mandatory: true, MaxRepeat: 1},   // Package identification
					{Segment: "RFF", Mandatory: false, MaxRepeat: 1},  // Reference
					{Segment: "DTM", Mandatory: false, MaxRepeat: 5},  // Date/time/period
					{Segment: "GIN", Mandatory: false, MaxRepeat: 10}, // Goods identity number
				},
				Mandatory: false, MaxRepeat: 5,
			},
		},
		Mandatory: false, MaxRepeat: 99,
	},
	{ // Segment group 16
		Group: []SchemaNode{
			{Segment: "EQD", Mandatory: true, MaxRepeat: 1},  // Equipment details
			{Segment: "HAN", Mandatory: false, MaxRepeat: 5}, // Handling instructions
			{Segment: "MEA", Mandatory: false, MaxRepeat: 5}, // Measurements
			{Segment: "FTX", Mandatory: false, MaxRepeat: 5}, // Free text
		},
		Mandatory: false, MaxRepeat: 10,
	},
	{ // Segment group 17
		Group: []SchemaNode{
			{Segment: "SCC", Mandatory: true, MaxRepeat: 1},  // Scheduling conditions
			{Segment: "FTX", Mandatory: false, MaxRepeat: 5}, // Free text
			{Segment: "RFF", Mandatory: false, MaxRepeat: 5}, // Reference
			{ // Segment group 18
				Group: []SchemaNode{
					{Segment: "QTY", Mandatory: true, MaxRepeat: 1},  // Quantity
					{Segment: "DTM", Mandatory: false, MaxRepeat: 5}, // Date/time/period
				},
				Mandatory: false, MaxRepeat: 10,
			},
		},
		Mandatory: false, MaxRepeat: 10,
	},
	{ // Segment group 19
		Group: []SchemaNode{
			{Segment: "APR", Mandatory: true, MaxRepeat: 1},  // Additional price information
			{Segment: "DTM", Mandatory: false, MaxRepeat: 5}, // Date/time/period
			{Segment: "RNG", Mandatory: false, MaxRepeat: 1}, // Range details
		},
		Mandatory: false, MaxRepeat: 25,
	},
	{ // Segment group 20
		Group: []SchemaNode{
			{Segment: "ALC", Mandatory: true, MaxRepeat: 1},  // Allowance or charge
			{Segment: "ALI", Mandatory: false, MaxRepeat: 5}, // Additional information
			{Segment: "DTM", Mandatory: false, MaxRepeat: 5}, // Date/time/period
			{ // Segment group 21
				Group: []SchemaNode{
					{Segment: "QTY", Mandatory: true, MaxRepeat: 1},  // Quantity
					{Segment: "RNG", Mandatory: false, MaxRepeat: 1}, // Range details
				},
				Mandatory: false, MaxRepeat: 1,
			},
			{ // Segment group 22
				Group: []SchemaNode{
					{Segment: "PCD", Mandatory: true, MaxRepeat: 1},  // Percentage details
					{Segment: "RNG", Mandatory: false, MaxRepeat: 1}, // Range details
				},
				Mandatory: false, MaxRepeat: 1,
			},
			{ // Segment group 23
				Group: []SchemaNode{
					{Segment: "MOA", Mandatory: true, MaxRepeat: 1},  // Monetary amount
					{Segment: "RNG", Mandatory: false, MaxRepeat: 1}, // Range details
				},
				Mandatory: false, MaxRepeat: 2,
			},
			{ // Segment group 24
				Group: []SchemaNode{
					{Segment: "RTE", Mandatory: true, MaxRepeat: 1},  // Rate details
					{Segment: "RNG", Mandatory: false, MaxRepeat: 1}, // Range details
				},
				Mandatory: false, MaxRepeat: 1,
			},
			{ // Segment group 25
				Group: []SchemaNode{
					{Segment: "TAX", Mandatory: true, MaxRepeat: 1},  // Duty/tax/fee details
					{Segment: "MOA", Mandatory: false, MaxRepeat: 1}, // Monetary amount
				},
				Mandatory: false, MaxRepeat: 5,
			},
		},
		Mandatory: false, MaxRepeat: 99,
	},
	{ // Segment group 26
		Group: []SchemaNode{
			{Segment: "RCS", Mandatory: true, MaxRepeat: 1},      // Requirements and conditions
			{Segment: "RFF", Mandatory: false, MaxRepeat: 5},     // Reference
			{Segment: "DTM", Mandatory: false, MaxRepeat: 5},     // Date/time/period
			{Segment: "FTX", Mandatory: false, MaxRepeat: 99999}, // Free text
		},
		Mandatory: false, MaxRepeat: 999,
	},
	{ // Segment group 27
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
	{ // Segment group 28
		Group: []SchemaNode{
			{Segment: "LIN", Mandatory: true, MaxRepeat: 1},     // Line item
			{Segment: "PIA", Mandatory: false, MaxRepeat: 25},   // Additional product id
			{Segment: "IMD", Mandatory: false, MaxRepeat: 99},   // Item description
			{Segment: "MEA", Mandatory: false, MaxRepeat: 99},   // Measurements
			{Segment: "QTY", Mandatory: false, MaxRepeat: 99},   // Quantity
			{Segment: "PCD", Mandatory: false, MaxRepeat: 5},    // Percentage details
			{Segment: "ALI", Mandatory: false, MaxRepeat: 5},    // Additional information
			{Segment: "DTM", Mandatory: false, MaxRepeat: 35},   // Date/time/period
			{Segment: "MOA", Mandatory: false, MaxRepeat: 10},   // Monetary amount
			{Segment: "GEI", Mandatory: false, MaxRepeat: 99},   // Processing information
			{Segment: "GIN", Mandatory: false, MaxRepeat: 1000}, // Goods identity number
			{Segment: "GIR", Mandatory: false, MaxRepeat: 1000}, // Related identification numbers
			{Segment: "QVR", Mandatory: false, MaxRepeat: 1},    // Quantity variances
			{Segment: "DOC", Mandatory: false, MaxRepeat: 99},   // Document/message details
			{Segment: "PAI", Mandatory: false, MaxRepeat: 1},    // Payment instructions
			{Segment: "MTD", Mandatory: false, MaxRepeat: 99},   // Maintenance operation details
			{Segment: "FTX", Mandatory: false, MaxRepeat: 99},   // Free text
			{ // Segment group 29
				Group: []SchemaNode{
					{Segment: "CCI", Mandatory: true, MaxRepeat: 1},   // Characteristic/class id
					{Segment: "CAV", Mandatory: false, MaxRepeat: 10}, // Characteristic value
					{Segment: "MEA", Mandatory: false, MaxRepeat: 10}, // Measurements
				},
				Mandatory: false, MaxRepeat: 999,
			},
			{ // Segment group 30
				Group: []SchemaNode{
					{Segment: "PYT", Mandatory: true, MaxRepeat: 1},  // Payment terms
					{Segment: "DTM", Mandatory: false, MaxRepeat: 5}, // Date/time/period
					{Segment: "PCD", Mandatory: false, MaxRepeat: 1}, // Percentage details
					{ // Segment group 31
						Group: []SchemaNode{
							{Segment: "MOA", Mandatory: true, MaxRepeat: 1},  // Monetary amount
							{Segment: "GIR", Mandatory: false, MaxRepeat: 9}, // Related identification numbers
						},
						Mandatory: false, MaxRepeat: 9999,
					},
				},
				Mandatory: false, MaxRepeat: 10,
			},
			{ // Segment group 32
				Group: []SchemaNode{
					{Segment: "AJT", Mandatory: true, MaxRepeat: 1},  // Adjustment details
					{Segment: "FTX", Mandatory: false, MaxRepeat: 5}, // Free text
				},
				Mandatory: false, MaxRepeat: 1,
			},
			{ // Segment group 33
				Group: []SchemaNode{
					{Segment: "PRI", Mandatory: true, MaxRepeat: 1},  // Price details
					{Segment: "CUX", Mandatory: false, MaxRepeat: 1}, // Currencies
					{Segment: "APR", Mandatory: false, MaxRepeat: 1}, // Additional price information
					{Segment: "RNG", Mandatory: false, MaxRepeat: 1}, // Range details
					{Segment: "DTM", Mandatory: false, MaxRepeat: 5}, // Date/time/period
				},
				Mandatory: false, MaxRepeat: 25,
			},
			{ // Segment group 34
				Group: []SchemaNode{
					{Segment: "RFF", Mandatory: true, MaxRepeat: 1},  // Reference
					{Segment: "DTM", Mandatory: false, MaxRepeat: 5}, // Date/time/period
				},
				Mandatory: false, MaxRepeat: 9999,
			},
			{ // Segment group 35
				Group: []SchemaNode{
					{Segment: "PAC", Mandatory: true, MaxRepeat: 1},  // Package
					{Segment: "MEA", Mandatory: false, MaxRepeat: 5}, // Measurements
					{Segment: "QTY", Mandatory: false, MaxRepeat: 5}, // Quantity
					{Segment: "DTM", Mandatory: false, MaxRepeat: 5}, // Date/time/period
					{ // Segment group 36
						Group: []SchemaNode{
							{Segment: "RFF", Mandatory: true, MaxRepeat: 1},  // Reference
							{Segment: "DTM", Mandatory: false, MaxRepeat: 5}, // Date/time/period
						},
						Mandatory: false, MaxRepeat: 1,
					},
					{ // Segment group 37
						Group: []SchemaNode{
							{Segment: "PCI", Mandatory: true, MaxRepeat: 1},   // Package identification
							{Segment: "RFF", Mandatory: false, MaxRepeat: 1},  // Reference
							{Segment: "DTM", Mandatory: false, MaxRepeat: 5},  // Date/time/period
							{Segment: "GIN", Mandatory: false, MaxRepeat: 10}, // Goods identity number
						},
						Mandatory: false, MaxRepeat: 5,
					},
				},
				Mandatory: false, MaxRepeat: 99,
			},
			{ // Segment group 38
				Group: []SchemaNode{
					{Segment: "LOC", Mandatory: true, MaxRepeat: 1},  // Place/location identification
					{Segment: "QTY", Mandatory: false, MaxRepeat: 1}, // Quantity
					{Segment: "DTM", Mandatory: false, MaxRepeat: 5}, // Date/time/period
				},
				Mandatory: false, MaxRepeat: 9999,
			},
			{ // Segment group 39
				Group: []SchemaNode{
					{Segment: "TAX", Mandatory: true, MaxRepeat: 1},  // Duty/tax/fee details
					{Segment: "MOA", Mandatory: false, MaxRepeat: 1}, // Monetary amount
					{Segment: "LOC", Mandatory: false, MaxRepeat: 5}, // Place/location identification
				},
				Mandatory: false, MaxRepeat: 10,
			},
			{ // Segment group 40
				Group: []SchemaNode{
					{Segment: "NAD", Mandatory: true, MaxRepeat: 1},  // Name and address
					{Segment: "LOC", Mandatory: false, MaxRepeat: 5}, // Place/location identification
					{Segment: "FII", Mandatory: false, MaxRepeat: 5}, // Financial institution information
					{ // Segment group 41
						Group: []SchemaNode{
							{Segment: "RFF", Mandatory: true, MaxRepeat: 1},  // Reference
							{Segment: "DTM", Mandatory: false, MaxRepeat: 5}, // Date/time/period
						},
						Mandatory: false, MaxRepeat: 99,
					},
					{ // Segment group 42
						Group: []SchemaNode{
							{Segment: "DOC", Mandatory: true, MaxRepeat: 1},  // Document/message details
							{Segment: "DTM", Mandatory: false, MaxRepeat: 5}, // Date/time/period
						},
						Mandatory: false, MaxRepeat: 5,
					},
					{ // Segment group 43
						Group: []SchemaNode{
							{Segment: "CTA", Mandatory: true, MaxRepeat: 1},  // Contact information
							{Segment: "COM", Mandatory: false, MaxRepeat: 5}, // Communication contact
						},
						Mandatory: false, MaxRepeat: 5,
					},
				},
				Mandatory: false, MaxRepeat: 999,
			},
			{ // Segment group 44
				Group: []SchemaNode{
					{Segment: "ALC", Mandatory: true, MaxRepeat: 1},  // Allowance or charge
					{Segment: "ALI", Mandatory: false, MaxRepeat: 5}, // Additional information
					{Segment: "DTM", Mandatory: false, MaxRepeat: 5}, // Date/time/period
					{ // Segment group 45
						Group: []SchemaNode{
							{Segment: "QTY", Mandatory: true, MaxRepeat: 1},  // Quantity
							{Segment: "RNG", Mandatory: false, MaxRepeat: 1}, // Range details
						},
						Mandatory: false, MaxRepeat: 1,
					},
					{ // Segment group 46
						Group: []SchemaNode{
							{Segment: "PCD", Mandatory: true, MaxRepeat: 1},  // Percentage details
							{Segment: "RNG", Mandatory: false, MaxRepeat: 1}, // Range details
						},
						Mandatory: false, MaxRepeat: 1,
					},
					{ // Segment group 47
						Group: []SchemaNode{
							{Segment: "MOA", Mandatory: true, MaxRepeat: 1},  // Monetary amount
							{Segment: "RNG", Mandatory: false, MaxRepeat: 1}, // Range details
						},
						Mandatory: false, MaxRepeat: 2,
					},
					{ // Segment group 48
						Group: []SchemaNode{
							{Segment: "RTE", Mandatory: true, MaxRepeat: 1},  // Rate details
							{Segment: "RNG", Mandatory: false, MaxRepeat: 1}, // Range details
						},
						Mandatory: false, MaxRepeat: 1,
					},
					{ // Segment group 49
						Group: []SchemaNode{
							{Segment: "TAX", Mandatory: true, MaxRepeat: 1},  // Duty/tax/fee details
							{Segment: "MOA", Mandatory: false, MaxRepeat: 1}, // Monetary amount
						},
						Mandatory: false, MaxRepeat: 5,
					},
				},
				Mandatory: false, MaxRepeat: 99,
			},
			{ // Segment group 50
				Group: []SchemaNode{
					{Segment: "TDT", Mandatory: true, MaxRepeat: 1}, // Transport information
					{ // Segment group 51
						Group: []SchemaNode{
							{Segment: "LOC", Mandatory: true, MaxRepeat: 1},  // Place/location identification
							{Segment: "DTM", Mandatory: false, MaxRepeat: 5}, // Date/time/period
						},
						Mandatory: false, MaxRepeat: 10,
					},
				},
				Mandatory: false, MaxRepeat: 10,
			},
			{ // Segment group 52
				Group: []SchemaNode{
					{Segment: "TOD", Mandatory: true, MaxRepeat: 1},  // Terms of delivery or transport
					{Segment: "LOC", Mandatory: false, MaxRepeat: 2}, // Place/location identification
				},
				Mandatory: false, MaxRepeat: 5,
			},
			{ // Segment group 53
				Group: []SchemaNode{
					{Segment: "EQD", Mandatory: true, MaxRepeat: 1},  // Equipment details
					{Segment: "HAN", Mandatory: false, MaxRepeat: 5}, // Handling instructions
					{Segment: "MEA", Mandatory: false, MaxRepeat: 5}, // Measurements
					{Segment: "FTX", Mandatory: false, MaxRepeat: 5}, // Free text
				},
				Mandatory: false, MaxRepeat: 10,
			},
			{ // Segment group 54
				Group: []SchemaNode{
					{Segment: "SCC", Mandatory: true, MaxRepeat: 1},  // Scheduling conditions
					{Segment: "FTX", Mandatory: false, MaxRepeat: 5}, // Free text
					{Segment: "RFF", Mandatory: false, MaxRepeat: 5}, // Reference
					{ // Segment group 55
						Group: []SchemaNode{
							{Segment: "QTY", Mandatory: true, MaxRepeat: 1},  // Quantity
							{Segment: "DTM", Mandatory: false, MaxRepeat: 5}, // Date/time/period
						},
						Mandatory: false, MaxRepeat: 10,
					},
				},
				Mandatory: false, MaxRepeat: 100,
			},
			{ // Segment group 56
				Group: []SchemaNode{
					{Segment: "RCS", Mandatory: true, MaxRepeat: 1},      // Requirements and conditions
					{Segment: "RFF", Mandatory: false, MaxRepeat: 5},     // Reference
					{Segment: "DTM", Mandatory: false, MaxRepeat: 5},     // Date/time/period
					{Segment: "FTX", Mandatory: false, MaxRepeat: 99999}, // Free text
				},
				Mandatory: false, MaxRepeat: 999,
			},
			{ // Segment group 57
				Group: []SchemaNode{
					{Segment: "STG", Mandatory: true, MaxRepeat: 1}, // Stages
					{ // Segment group 58
						Group: []SchemaNode{
							{Segment: "QTY", Mandatory: true, MaxRepeat: 1},  // Quantity
							{Segment: "MOA", Mandatory: false, MaxRepeat: 1}, // Monetary amount
						},
						Mandatory: false, MaxRepeat: 3,
					},
				},
				Mandatory: false, MaxRepeat: 10,
			},
			{ // Segment group 59
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
		Mandatory: false, MaxRepeat: 200000,
	},
	{Segment: "UNS", Mandatory: true, MaxRepeat: 1},   // Section control
	{Segment: "MOA", Mandatory: false, MaxRepeat: 12}, // Monetary amount
	{Segment: "CNT", Mandatory: false, MaxRepeat: 10}, // Control total
	{ // Segment group 60
		Group: []SchemaNode{
			{Segment: "ALC", Mandatory: true, MaxRepeat: 1},  // Allowance or charge
			{Segment: "ALI", Mandatory: false, MaxRepeat: 1}, // Additional information
			{Segment: "MOA", Mandatory: true, MaxRepeat: 2},  // Monetary amount
		},
		Mandatory: false, MaxRepeat: 10,
	},
}}
