package edifact

// IFTFCC D.20A message specification: the structural (segment/group
// presence, order, cardinality) branching diagram for this message,
// UN/EDIFACT directory release D.20A. Scope is structural only -- not
// element-level content -- per edifact-ls-3uzr's epic non-goal.
// 45 segment groups, max nesting depth 3.
//
// Source: https://service.unece.org/trade/untdid/d20a/trmd/iftfcc_c.htm
// section 4.3.1 "Segment table". That URL currently returns HTTP 403 from
// Cloudflare when fetched directly; this data was transcribed from the
// Wayback Machine's archived copy instead:
// http://web.archive.org/web/20231202090031/https://service.unece.org/trade/untdid/d20a/trmd/iftfcc_c.htm
// -- re-check the direct URL first if cross-checking later, in case the
// block has lifted.
//
// Transcription note: generated in bulk (edifact-ls-13gu) the same way
// as every other message-type schema -- a script parsed the source's
// exact rail-art column positions mechanically and verified the result
// balances (all 45 groups open and close consistently)
// before being transcribed here, rather than reading the ASCII nesting
// by eye. Not individually spot-checked against the raw source the way
// the first several message types in this project were -- see
// edifact-ls-13gu for which representative samples were spot-checked
// for this batch.

func init() {
	RegisterSchema(
		MessageID{Type: "IFTFCC", Version: "D", Release: "20A", Agency: "UN"},
		iftfccD20aSchema,
		"https://service.unece.org/trade/untdid/d20a/trmd/iftfcc_c.htm",
	)
}

var iftfccD20aSchema = Schema{Nodes: []SchemaNode{
	{Segment: "BGM", Mandatory: true, MaxRepeat: 1},   // Beginning of message
	{Segment: "CTA", Mandatory: false, MaxRepeat: 1},  // Contact information
	{Segment: "COM", Mandatory: false, MaxRepeat: 9},  // Communication contact
	{Segment: "FTX", Mandatory: false, MaxRepeat: 99}, // Free text
	{Segment: "DTM", Mandatory: false, MaxRepeat: 9},  // Date/time/period
	{Segment: "TSR", Mandatory: false, MaxRepeat: 9},  // Transport service requirements
	{Segment: "DOC", Mandatory: false, MaxRepeat: 9},  // Document/message details
	{ // Segment group 1
		Group: []SchemaNode{
			{Segment: "LOC", Mandatory: true, MaxRepeat: 1},  // Place/location identification
			{Segment: "DTM", Mandatory: false, MaxRepeat: 9}, // Date/time/period
		},
		Mandatory: false, MaxRepeat: 9,
	},
	{ // Segment group 2
		Group: []SchemaNode{
			{Segment: "RFF", Mandatory: true, MaxRepeat: 1},  // Reference
			{Segment: "DTM", Mandatory: false, MaxRepeat: 9}, // Date/time/period
		},
		Mandatory: false, MaxRepeat: 9,
	},
	{ // Segment group 3
		Group: []SchemaNode{
			{Segment: "MOA", Mandatory: true, MaxRepeat: 1},  // Monetary amount
			{Segment: "PCD", Mandatory: false, MaxRepeat: 1}, // Percentage details
		},
		Mandatory: false, MaxRepeat: 99,
	},
	{ // Segment group 4
		Group: []SchemaNode{
			{Segment: "TAX", Mandatory: true, MaxRepeat: 1},  // Duty/tax/fee details
			{Segment: "PCD", Mandatory: false, MaxRepeat: 1}, // Percentage details
			{Segment: "MOA", Mandatory: false, MaxRepeat: 2}, // Monetary amount
		},
		Mandatory: false, MaxRepeat: 9,
	},
	{ // Segment group 5
		Group: []SchemaNode{
			{Segment: "CUX", Mandatory: true, MaxRepeat: 1},  // Currencies
			{Segment: "DTM", Mandatory: false, MaxRepeat: 9}, // Date/time/period
		},
		Mandatory: false, MaxRepeat: 9,
	},
	{ // Segment group 6
		Group: []SchemaNode{
			{Segment: "TCC", Mandatory: true, MaxRepeat: 1},  // Charge/rate calculations
			{Segment: "LOC", Mandatory: false, MaxRepeat: 9}, // Place/location identification
			{Segment: "DTM", Mandatory: false, MaxRepeat: 9}, // Date/time/period
			{Segment: "RFF", Mandatory: false, MaxRepeat: 9}, // Reference
			{Segment: "FTX", Mandatory: false, MaxRepeat: 9}, // Free text
			{Segment: "PCD", Mandatory: false, MaxRepeat: 9}, // Percentage details
			{Segment: "QTY", Mandatory: false, MaxRepeat: 9}, // Quantity
			{ // Segment group 7
				Group: []SchemaNode{
					{Segment: "PRI", Mandatory: true, MaxRepeat: 1},  // Price details
					{Segment: "CUX", Mandatory: false, MaxRepeat: 9}, // Currencies
				},
				Mandatory: false, MaxRepeat: 9,
			},
			{ // Segment group 8
				Group: []SchemaNode{
					{Segment: "MOA", Mandatory: true, MaxRepeat: 1}, // Monetary amount
					{ // Segment group 9
						Group: []SchemaNode{
							{Segment: "CUX", Mandatory: true, MaxRepeat: 1},  // Currencies
							{Segment: "DTM", Mandatory: false, MaxRepeat: 1}, // Date/time/period
						},
						Mandatory: false, MaxRepeat: 1,
					},
					{Segment: "PCD", Mandatory: false, MaxRepeat: 1}, // Percentage details
					{ // Segment group 10
						Group: []SchemaNode{
							{Segment: "TAX", Mandatory: true, MaxRepeat: 1},  // Duty/tax/fee details
							{Segment: "PCD", Mandatory: false, MaxRepeat: 1}, // Percentage details
						},
						Mandatory: false, MaxRepeat: 1,
					},
				},
				Mandatory: false, MaxRepeat: 9,
			},
		},
		Mandatory: false, MaxRepeat: 999,
	},
	{ // Segment group 11
		Group: []SchemaNode{
			{Segment: "NAD", Mandatory: true, MaxRepeat: 1},  // Name and address
			{Segment: "FII", Mandatory: false, MaxRepeat: 9}, // Financial institution information
			{Segment: "LOC", Mandatory: false, MaxRepeat: 9}, // Place/location identification
			{ // Segment group 12
				Group: []SchemaNode{
					{Segment: "CTA", Mandatory: true, MaxRepeat: 1},  // Contact information
					{Segment: "COM", Mandatory: false, MaxRepeat: 9}, // Communication contact
				},
				Mandatory: false, MaxRepeat: 9,
			},
			{ // Segment group 13
				Group: []SchemaNode{
					{Segment: "RFF", Mandatory: true, MaxRepeat: 1},  // Reference
					{Segment: "DTM", Mandatory: false, MaxRepeat: 9}, // Date/time/period
				},
				Mandatory: false, MaxRepeat: 9,
			},
		},
		Mandatory: false, MaxRepeat: 99,
	},
	{ // Segment group 14
		Group: []SchemaNode{
			{Segment: "TOD", Mandatory: true, MaxRepeat: 1},  // Terms of delivery or transport
			{Segment: "LOC", Mandatory: false, MaxRepeat: 2}, // Place/location identification
		},
		Mandatory: false, MaxRepeat: 5,
	},
	{ // Segment group 15
		Group: []SchemaNode{
			{Segment: "CPI", Mandatory: true, MaxRepeat: 1},  // Charge payment instructions
			{Segment: "CUX", Mandatory: false, MaxRepeat: 1}, // Currencies
			{Segment: "LOC", Mandatory: false, MaxRepeat: 9}, // Place/location identification
			{Segment: "MOA", Mandatory: false, MaxRepeat: 9}, // Monetary amount
		},
		Mandatory: false, MaxRepeat: 9,
	},
	{ // Segment group 16
		Group: []SchemaNode{
			{Segment: "PYT", Mandatory: true, MaxRepeat: 1},  // Payment terms
			{Segment: "DTM", Mandatory: false, MaxRepeat: 5}, // Date/time/period
			{Segment: "PCD", Mandatory: false, MaxRepeat: 1}, // Percentage details
			{Segment: "MOA", Mandatory: false, MaxRepeat: 1}, // Monetary amount
		},
		Mandatory: false, MaxRepeat: 5,
	},
	{ // Segment group 17
		Group: []SchemaNode{
			{Segment: "TDT", Mandatory: true, MaxRepeat: 1},   // Transport information
			{Segment: "TCC", Mandatory: false, MaxRepeat: 99}, // Charge/rate calculations
			{Segment: "DTM", Mandatory: false, MaxRepeat: 9},  // Date/time/period
			{Segment: "LOC", Mandatory: false, MaxRepeat: 9},  // Place/location identification
			{ // Segment group 18
				Group: []SchemaNode{
					{Segment: "RFF", Mandatory: true, MaxRepeat: 1},  // Reference
					{Segment: "DTM", Mandatory: false, MaxRepeat: 1}, // Date/time/period
				},
				Mandatory: false, MaxRepeat: 9,
			},
		},
		Mandatory: false, MaxRepeat: 99,
	},
	{ // Segment group 19
		Group: []SchemaNode{
			{Segment: "GID", Mandatory: true, MaxRepeat: 1},   // Goods item details
			{Segment: "TCC", Mandatory: false, MaxRepeat: 99}, // Charge/rate calculations
			{Segment: "HAN", Mandatory: false, MaxRepeat: 1},  // Handling instructions
			{Segment: "TMP", Mandatory: false, MaxRepeat: 1},  // Temperature
			{Segment: "TMD", Mandatory: false, MaxRepeat: 1},  // Transport movement details
			{Segment: "LOC", Mandatory: false, MaxRepeat: 9},  // Place/location identification
			{Segment: "PCI", Mandatory: false, MaxRepeat: 9},  // Package identification
			{Segment: "PIA", Mandatory: false, MaxRepeat: 9},  // Additional product id
			{Segment: "FTX", Mandatory: false, MaxRepeat: 9},  // Free text
			{ // Segment group 20
				Group: []SchemaNode{
					{Segment: "MEA", Mandatory: true, MaxRepeat: 1},  // Measurements
					{Segment: "EQN", Mandatory: false, MaxRepeat: 1}, // Number of units
				},
				Mandatory: false, MaxRepeat: 9,
			},
			{ // Segment group 21
				Group: []SchemaNode{
					{Segment: "DIM", Mandatory: true, MaxRepeat: 1},  // Dimensions
					{Segment: "EQN", Mandatory: false, MaxRepeat: 1}, // Number of units
				},
				Mandatory: false, MaxRepeat: 9,
			},
			{ // Segment group 22
				Group: []SchemaNode{
					{Segment: "RFF", Mandatory: true, MaxRepeat: 1},  // Reference
					{Segment: "DTM", Mandatory: false, MaxRepeat: 1}, // Date/time/period
				},
				Mandatory: false, MaxRepeat: 9,
			},
			{ // Segment group 23
				Group: []SchemaNode{
					{Segment: "DGS", Mandatory: true, MaxRepeat: 1},  // Dangerous goods
					{Segment: "FTX", Mandatory: false, MaxRepeat: 1}, // Free text
				},
				Mandatory: false, MaxRepeat: 9,
			},
		},
		Mandatory: false, MaxRepeat: 99,
	},
	{ // Segment group 24
		Group: []SchemaNode{
			{Segment: "EQD", Mandatory: true, MaxRepeat: 1},   // Equipment details
			{Segment: "TCC", Mandatory: false, MaxRepeat: 99}, // Charge/rate calculations
			{Segment: "EQN", Mandatory: false, MaxRepeat: 1},  // Number of units
			{Segment: "TMD", Mandatory: false, MaxRepeat: 1},  // Transport movement details
			{Segment: "MEA", Mandatory: false, MaxRepeat: 9},  // Measurements
			{Segment: "DIM", Mandatory: false, MaxRepeat: 9},  // Dimensions
			{Segment: "SEL", Mandatory: false, MaxRepeat: 99}, // Seal number
			{Segment: "TPL", Mandatory: false, MaxRepeat: 9},  // Transport placement
			{Segment: "FTX", Mandatory: false, MaxRepeat: 9},  // Free text
			{ // Segment group 25
				Group: []SchemaNode{
					{Segment: "RFF", Mandatory: true, MaxRepeat: 1},  // Reference
					{Segment: "DTM", Mandatory: false, MaxRepeat: 9}, // Date/time/period
				},
				Mandatory: false, MaxRepeat: 9,
			},
			{ // Segment group 26
				Group: []SchemaNode{
					{Segment: "NAD", Mandatory: true, MaxRepeat: 1},  // Name and address
					{Segment: "DTM", Mandatory: false, MaxRepeat: 1}, // Date/time/period
				},
				Mandatory: false, MaxRepeat: 9,
			},
			{ // Segment group 27
				Group: []SchemaNode{
					{Segment: "EQA", Mandatory: true, MaxRepeat: 1},  // Attached equipment
					{Segment: "EQN", Mandatory: false, MaxRepeat: 1}, // Number of units
				},
				Mandatory: false, MaxRepeat: 99,
			},
		},
		Mandatory: false, MaxRepeat: 999,
	},
	{ // Segment group 28
		Group: []SchemaNode{
			{Segment: "CNI", Mandatory: true, MaxRepeat: 1},   // Consignment information
			{Segment: "TCC", Mandatory: false, MaxRepeat: 99}, // Charge/rate calculations
			{Segment: "DTM", Mandatory: false, MaxRepeat: 9},  // Date/time/period
			{Segment: "TSR", Mandatory: false, MaxRepeat: 9},  // Transport service requirements
			{Segment: "FTX", Mandatory: false, MaxRepeat: 9},  // Free text
			{Segment: "MOA", Mandatory: false, MaxRepeat: 9},  // Monetary amount
			{ // Segment group 29
				Group: []SchemaNode{
					{Segment: "CUX", Mandatory: true, MaxRepeat: 1},  // Currencies
					{Segment: "DTM", Mandatory: false, MaxRepeat: 9}, // Date/time/period
				},
				Mandatory: false, MaxRepeat: 9,
			},
			{ // Segment group 30
				Group: []SchemaNode{
					{Segment: "LOC", Mandatory: true, MaxRepeat: 1},  // Place/location identification
					{Segment: "DTM", Mandatory: false, MaxRepeat: 9}, // Date/time/period
				},
				Mandatory: false, MaxRepeat: 9,
			},
			{ // Segment group 31
				Group: []SchemaNode{
					{Segment: "RFF", Mandatory: true, MaxRepeat: 1},  // Reference
					{Segment: "DTM", Mandatory: false, MaxRepeat: 9}, // Date/time/period
				},
				Mandatory: false, MaxRepeat: 9,
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
					{Segment: "CPI", Mandatory: true, MaxRepeat: 1},  // Charge payment instructions
					{Segment: "CUX", Mandatory: false, MaxRepeat: 1}, // Currencies
					{Segment: "LOC", Mandatory: false, MaxRepeat: 9}, // Place/location identification
					{Segment: "MOA", Mandatory: false, MaxRepeat: 9}, // Monetary amount
				},
				Mandatory: false, MaxRepeat: 9,
			},
			{ // Segment group 34
				Group: []SchemaNode{
					{Segment: "TDT", Mandatory: true, MaxRepeat: 1},  // Transport information
					{Segment: "DTM", Mandatory: false, MaxRepeat: 9}, // Date/time/period
					{Segment: "LOC", Mandatory: false, MaxRepeat: 9}, // Place/location identification
					{Segment: "RFF", Mandatory: false, MaxRepeat: 9}, // Reference
				},
				Mandatory: false, MaxRepeat: 99,
			},
			{ // Segment group 35
				Group: []SchemaNode{
					{Segment: "NAD", Mandatory: true, MaxRepeat: 1}, // Name and address
					{ // Segment group 36
						Group: []SchemaNode{
							{Segment: "CTA", Mandatory: true, MaxRepeat: 1},  // Contact information
							{Segment: "COM", Mandatory: false, MaxRepeat: 9}, // Communication contact
						},
						Mandatory: false, MaxRepeat: 9,
					},
					{Segment: "RFF", Mandatory: false, MaxRepeat: 9}, // Reference
				},
				Mandatory: false, MaxRepeat: 99,
			},
			{ // Segment group 37
				Group: []SchemaNode{
					{Segment: "GID", Mandatory: true, MaxRepeat: 1},   // Goods item details
					{Segment: "TCC", Mandatory: false, MaxRepeat: 99}, // Charge/rate calculations
					{Segment: "HAN", Mandatory: false, MaxRepeat: 1},  // Handling instructions
					{Segment: "TMP", Mandatory: false, MaxRepeat: 1},  // Temperature
					{Segment: "TMD", Mandatory: false, MaxRepeat: 1},  // Transport movement details
					{Segment: "LOC", Mandatory: false, MaxRepeat: 9},  // Place/location identification
					{Segment: "PCI", Mandatory: false, MaxRepeat: 9},  // Package identification
					{Segment: "PIA", Mandatory: false, MaxRepeat: 9},  // Additional product id
					{Segment: "FTX", Mandatory: false, MaxRepeat: 9},  // Free text
					{ // Segment group 38
						Group: []SchemaNode{
							{Segment: "MEA", Mandatory: true, MaxRepeat: 1},  // Measurements
							{Segment: "EQN", Mandatory: false, MaxRepeat: 1}, // Number of units
						},
						Mandatory: false, MaxRepeat: 9,
					},
					{ // Segment group 39
						Group: []SchemaNode{
							{Segment: "DIM", Mandatory: true, MaxRepeat: 1},  // Dimensions
							{Segment: "EQN", Mandatory: false, MaxRepeat: 1}, // Number of units
						},
						Mandatory: false, MaxRepeat: 9,
					},
					{ // Segment group 40
						Group: []SchemaNode{
							{Segment: "RFF", Mandatory: true, MaxRepeat: 1},  // Reference
							{Segment: "DTM", Mandatory: false, MaxRepeat: 1}, // Date/time/period
						},
						Mandatory: false, MaxRepeat: 9,
					},
					{ // Segment group 41
						Group: []SchemaNode{
							{Segment: "DGS", Mandatory: true, MaxRepeat: 1},  // Dangerous goods
							{Segment: "FTX", Mandatory: false, MaxRepeat: 1}, // Free text
						},
						Mandatory: false, MaxRepeat: 9,
					},
				},
				Mandatory: false, MaxRepeat: 99,
			},
			{ // Segment group 42
				Group: []SchemaNode{
					{Segment: "EQD", Mandatory: true, MaxRepeat: 1},   // Equipment details
					{Segment: "TCC", Mandatory: false, MaxRepeat: 99}, // Charge/rate calculations
					{Segment: "EQN", Mandatory: false, MaxRepeat: 1},  // Number of units
					{Segment: "TMD", Mandatory: false, MaxRepeat: 1},  // Transport movement details
					{Segment: "MEA", Mandatory: false, MaxRepeat: 9},  // Measurements
					{Segment: "DIM", Mandatory: false, MaxRepeat: 9},  // Dimensions
					{Segment: "SEL", Mandatory: false, MaxRepeat: 99}, // Seal number
					{Segment: "TPL", Mandatory: false, MaxRepeat: 9},  // Transport placement
					{Segment: "FTX", Mandatory: false, MaxRepeat: 9},  // Free text
					{ // Segment group 43
						Group: []SchemaNode{
							{Segment: "RFF", Mandatory: true, MaxRepeat: 1},  // Reference
							{Segment: "DTM", Mandatory: false, MaxRepeat: 9}, // Date/time/period
						},
						Mandatory: false, MaxRepeat: 9,
					},
					{ // Segment group 44
						Group: []SchemaNode{
							{Segment: "NAD", Mandatory: true, MaxRepeat: 1},  // Name and address
							{Segment: "DTM", Mandatory: false, MaxRepeat: 1}, // Date/time/period
						},
						Mandatory: false, MaxRepeat: 9,
					},
					{ // Segment group 45
						Group: []SchemaNode{
							{Segment: "EQA", Mandatory: true, MaxRepeat: 1},  // Attached equipment
							{Segment: "EQN", Mandatory: false, MaxRepeat: 1}, // Number of units
						},
						Mandatory: false, MaxRepeat: 99,
					},
				},
				Mandatory: false, MaxRepeat: 999,
			},
		},
		Mandatory: false, MaxRepeat: 99,
	},
}}
