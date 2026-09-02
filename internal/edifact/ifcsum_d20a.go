package edifact

// IFCSUM D.20A message specification: the structural (segment/group
// presence, order, cardinality) branching diagram for this message,
// UN/EDIFACT directory release D.20A. Scope is structural only -- not
// element-level content -- per edifact-ls-3uzr's epic non-goal.
// 80 segment groups, max nesting depth 5.
//
// Source: https://service.unece.org/trade/untdid/d20a/trmd/ifcsum_c.htm
// section 4.3.1 "Segment table". That URL currently returns HTTP 403 from
// Cloudflare when fetched directly; this data was transcribed from the
// Wayback Machine's archived copy instead:
// http://web.archive.org/web/20231002205456/https://service.unece.org/trade/untdid/d20a/trmd/ifcsum_c.htm
// -- re-check the direct URL first if cross-checking later, in case the
// block has lifted.
//
// Transcription note: generated in bulk (edifact-ls-13gu) the same way
// as every other message-type schema -- a script parsed the source's
// exact rail-art column positions mechanically and verified the result
// balances (all 80 groups open and close consistently)
// before being transcribed here, rather than reading the ASCII nesting
// by eye. Not individually spot-checked against the raw source the way
// the first several message types in this project were -- see
// edifact-ls-13gu for which representative samples were spot-checked
// for this batch.

func init() {
	RegisterSchema(
		MessageID{Type: "IFCSUM", Version: "D", Release: "20A", Agency: "UN"},
		ifcsumD20aSchema,
		"https://service.unece.org/trade/untdid/d20a/trmd/ifcsum_c.htm",
	)
}

var ifcsumD20aSchema = Schema{Nodes: []SchemaNode{
	{Segment: "BGM", Mandatory: true, MaxRepeat: 1},   // Beginning of message
	{Segment: "DTM", Mandatory: false, MaxRepeat: 9},  // Date/time/period
	{Segment: "MOA", Mandatory: false, MaxRepeat: 99}, // Monetary amount
	{Segment: "FTX", Mandatory: false, MaxRepeat: 99}, // Free text
	{Segment: "CNT", Mandatory: false, MaxRepeat: 9},  // Control total
	{Segment: "PCD", Mandatory: false, MaxRepeat: 1},  // Percentage details
	{Segment: "GDS", Mandatory: false, MaxRepeat: 9},  // Nature of cargo
	{ // Segment group 1
		Group: []SchemaNode{
			{Segment: "RFF", Mandatory: true, MaxRepeat: 1},  // Reference
			{Segment: "DTM", Mandatory: false, MaxRepeat: 9}, // Date/time/period
		},
		Mandatory: false, MaxRepeat: 9,
	},
	{ // Segment group 2
		Group: []SchemaNode{
			{Segment: "GOR", Mandatory: true, MaxRepeat: 1},  // Governmental requirements
			{Segment: "DTM", Mandatory: false, MaxRepeat: 9}, // Date/time/period
			{Segment: "LOC", Mandatory: false, MaxRepeat: 9}, // Place/location identification
			{Segment: "SEL", Mandatory: false, MaxRepeat: 9}, // Seal number
			{Segment: "FTX", Mandatory: false, MaxRepeat: 9}, // Free text
			{ // Segment group 3
				Group: []SchemaNode{
					{Segment: "DOC", Mandatory: true, MaxRepeat: 1},  // Document/message details
					{Segment: "DTM", Mandatory: false, MaxRepeat: 1}, // Date/time/period
				},
				Mandatory: false, MaxRepeat: 9,
			},
		},
		Mandatory: false, MaxRepeat: 9,
	},
	{ // Segment group 4
		Group: []SchemaNode{
			{Segment: "NAD", Mandatory: true, MaxRepeat: 1}, // Name and address
			{ // Segment group 5
				Group: []SchemaNode{
					{Segment: "CTA", Mandatory: true, MaxRepeat: 1},  // Contact information
					{Segment: "COM", Mandatory: false, MaxRepeat: 9}, // Communication contact
				},
				Mandatory: false, MaxRepeat: 9,
			},
			{ // Segment group 6
				Group: []SchemaNode{
					{Segment: "RFF", Mandatory: true, MaxRepeat: 1},  // Reference
					{Segment: "DTM", Mandatory: false, MaxRepeat: 9}, // Date/time/period
				},
				Mandatory: false, MaxRepeat: 9,
			},
		},
		Mandatory: false, MaxRepeat: 9,
	},
	{ // Segment group 7
		Group: []SchemaNode{
			{Segment: "TCC", Mandatory: true, MaxRepeat: 1},  // Charge/rate calculations
			{Segment: "CUX", Mandatory: false, MaxRepeat: 1}, // Currencies
			{Segment: "PRI", Mandatory: false, MaxRepeat: 1}, // Price details
			{Segment: "EQN", Mandatory: false, MaxRepeat: 1}, // Number of units
			{Segment: "PCD", Mandatory: false, MaxRepeat: 1}, // Percentage details
			{Segment: "MOA", Mandatory: false, MaxRepeat: 9}, // Monetary amount
			{Segment: "QTY", Mandatory: false, MaxRepeat: 9}, // Quantity
			{Segment: "LOC", Mandatory: false, MaxRepeat: 9}, // Place/location identification
		},
		Mandatory: false, MaxRepeat: 9,
	},
	{ // Segment group 8
		Group: []SchemaNode{
			{Segment: "ICD", Mandatory: true, MaxRepeat: 1},  // Insurance cover description
			{Segment: "DTM", Mandatory: false, MaxRepeat: 1}, // Date/time/period
			{Segment: "FTX", Mandatory: false, MaxRepeat: 9}, // Free text
		},
		Mandatory: false, MaxRepeat: 99,
	},
	{ // Segment group 9
		Group: []SchemaNode{
			{Segment: "TDT", Mandatory: true, MaxRepeat: 1},  // Transport information
			{Segment: "DTM", Mandatory: false, MaxRepeat: 9}, // Date/time/period
			{ // Segment group 10
				Group: []SchemaNode{
					{Segment: "TSR", Mandatory: true, MaxRepeat: 1},  // Transport service requirements
					{Segment: "SCC", Mandatory: false, MaxRepeat: 9}, // Scheduling conditions
				},
				Mandatory: false, MaxRepeat: 9,
			},
			{ // Segment group 11
				Group: []SchemaNode{
					{Segment: "LOC", Mandatory: true, MaxRepeat: 1},  // Place/location identification
					{Segment: "DTM", Mandatory: false, MaxRepeat: 9}, // Date/time/period
				},
				Mandatory: false, MaxRepeat: 9,
			},
			{Segment: "SEL", Mandatory: false, MaxRepeat: 9}, // Seal number
			{Segment: "FTX", Mandatory: false, MaxRepeat: 9}, // Free text
			{ // Segment group 12
				Group: []SchemaNode{
					{Segment: "MEA", Mandatory: true, MaxRepeat: 1},  // Measurements
					{Segment: "EQN", Mandatory: false, MaxRepeat: 1}, // Number of units
				},
				Mandatory: false, MaxRepeat: 9,
			},
			{ // Segment group 13
				Group: []SchemaNode{
					{Segment: "DIM", Mandatory: true, MaxRepeat: 1},  // Dimensions
					{Segment: "EQN", Mandatory: false, MaxRepeat: 1}, // Number of units
				},
				Mandatory: false, MaxRepeat: 9,
			},
			{ // Segment group 14
				Group: []SchemaNode{
					{Segment: "CTA", Mandatory: true, MaxRepeat: 1},  // Contact information
					{Segment: "COM", Mandatory: false, MaxRepeat: 9}, // Communication contact
				},
				Mandatory: false, MaxRepeat: 9,
			},
			{ // Segment group 15
				Group: []SchemaNode{
					{Segment: "TCC", Mandatory: true, MaxRepeat: 1},  // Charge/rate calculations
					{Segment: "MOA", Mandatory: false, MaxRepeat: 9}, // Monetary amount
					{Segment: "PCD", Mandatory: false, MaxRepeat: 1}, // Percentage details
				},
				Mandatory: false, MaxRepeat: 99,
			},
			{ // Segment group 16
				Group: []SchemaNode{
					{Segment: "ICD", Mandatory: true, MaxRepeat: 1},  // Insurance cover description
					{Segment: "DTM", Mandatory: false, MaxRepeat: 1}, // Date/time/period
					{Segment: "FTX", Mandatory: false, MaxRepeat: 9}, // Free text
				},
				Mandatory: false, MaxRepeat: 99,
			},
			{ // Segment group 17
				Group: []SchemaNode{
					{Segment: "RFF", Mandatory: true, MaxRepeat: 1},  // Reference
					{Segment: "DTM", Mandatory: false, MaxRepeat: 9}, // Date/time/period
				},
				Mandatory: false, MaxRepeat: 9,
			},
			{ // Segment group 18
				Group: []SchemaNode{
					{Segment: "NAD", Mandatory: true, MaxRepeat: 1},  // Name and address
					{Segment: "LOC", Mandatory: false, MaxRepeat: 9}, // Place/location identification
					{ // Segment group 19
						Group: []SchemaNode{
							{Segment: "CTA", Mandatory: true, MaxRepeat: 1},  // Contact information
							{Segment: "COM", Mandatory: false, MaxRepeat: 9}, // Communication contact
						},
						Mandatory: false, MaxRepeat: 9,
					},
					{ // Segment group 20
						Group: []SchemaNode{
							{Segment: "DOC", Mandatory: true, MaxRepeat: 1},  // Document/message details
							{Segment: "DTM", Mandatory: false, MaxRepeat: 1}, // Date/time/period
						},
						Mandatory: false, MaxRepeat: 9,
					},
					{ // Segment group 21
						Group: []SchemaNode{
							{Segment: "TCC", Mandatory: true, MaxRepeat: 1},  // Charge/rate calculations
							{Segment: "CUX", Mandatory: false, MaxRepeat: 1}, // Currencies
							{Segment: "PRI", Mandatory: false, MaxRepeat: 1}, // Price details
							{Segment: "EQN", Mandatory: false, MaxRepeat: 1}, // Number of units
							{Segment: "PCD", Mandatory: false, MaxRepeat: 1}, // Percentage details
							{Segment: "MOA", Mandatory: false, MaxRepeat: 9}, // Monetary amount
							{Segment: "QTY", Mandatory: false, MaxRepeat: 9}, // Quantity
						},
						Mandatory: false, MaxRepeat: 9,
					},
					{ // Segment group 22
						Group: []SchemaNode{
							{Segment: "RFF", Mandatory: true, MaxRepeat: 1},  // Reference
							{Segment: "DTM", Mandatory: false, MaxRepeat: 9}, // Date/time/period
						},
						Mandatory: false, MaxRepeat: 9,
					},
				},
				Mandatory: false, MaxRepeat: 9,
			},
		},
		Mandatory: false, MaxRepeat: 9,
	},
	{ // Segment group 23
		Group: []SchemaNode{
			{Segment: "EQD", Mandatory: true, MaxRepeat: 1},  // Equipment details
			{Segment: "EQN", Mandatory: false, MaxRepeat: 1}, // Number of units
			{Segment: "TPL", Mandatory: false, MaxRepeat: 1}, // Transport placement
			{Segment: "TMD", Mandatory: false, MaxRepeat: 1}, // Transport movement details
			{Segment: "MEA", Mandatory: false, MaxRepeat: 9}, // Measurements
			{Segment: "DIM", Mandatory: false, MaxRepeat: 9}, // Dimensions
			{Segment: "SEL", Mandatory: false, MaxRepeat: 9}, // Seal number
			{Segment: "NAD", Mandatory: false, MaxRepeat: 9}, // Name and address
			{Segment: "LOC", Mandatory: false, MaxRepeat: 9}, // Place/location identification
			{Segment: "HAN", Mandatory: false, MaxRepeat: 1}, // Handling instructions
			{Segment: "TMP", Mandatory: false, MaxRepeat: 1}, // Temperature
			{Segment: "FTX", Mandatory: false, MaxRepeat: 9}, // Free text
			{Segment: "RFF", Mandatory: false, MaxRepeat: 9}, // Reference
			{Segment: "PCD", Mandatory: false, MaxRepeat: 9}, // Percentage details
			{ // Segment group 24
				Group: []SchemaNode{
					{Segment: "EQA", Mandatory: true, MaxRepeat: 1},  // Attached equipment
					{Segment: "EQN", Mandatory: false, MaxRepeat: 1}, // Number of units
				},
				Mandatory: false, MaxRepeat: 99,
			},
			{ // Segment group 25
				Group: []SchemaNode{
					{Segment: "DGS", Mandatory: true, MaxRepeat: 1},  // Dangerous goods
					{Segment: "FTX", Mandatory: false, MaxRepeat: 9}, // Free text
					{ // Segment group 26
						Group: []SchemaNode{
							{Segment: "CTA", Mandatory: true, MaxRepeat: 1},  // Contact information
							{Segment: "COM", Mandatory: false, MaxRepeat: 9}, // Communication contact
						},
						Mandatory: false, MaxRepeat: 9,
					},
				},
				Mandatory: false, MaxRepeat: 99,
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
		},
		Mandatory: false, MaxRepeat: 9999,
	},
	{ // Segment group 28
		Group: []SchemaNode{
			{Segment: "CNI", Mandatory: true, MaxRepeat: 1}, // Consignment information
			{ // Segment group 29
				Group: []SchemaNode{
					{Segment: "SGP", Mandatory: true, MaxRepeat: 1}, // Split goods placement
					{ // Segment group 30
						Group: []SchemaNode{
							{Segment: "MEA", Mandatory: true, MaxRepeat: 1},  // Measurements
							{Segment: "EQN", Mandatory: false, MaxRepeat: 1}, // Number of units
						},
						Mandatory: false, MaxRepeat: 9,
					},
				},
				Mandatory: false, MaxRepeat: 9,
			},
			{ // Segment group 31
				Group: []SchemaNode{
					{Segment: "TPL", Mandatory: true, MaxRepeat: 1}, // Transport placement
					{ // Segment group 32
						Group: []SchemaNode{
							{Segment: "MEA", Mandatory: true, MaxRepeat: 1},  // Measurements
							{Segment: "EQN", Mandatory: false, MaxRepeat: 1}, // Number of units
						},
						Mandatory: false, MaxRepeat: 9,
					},
				},
				Mandatory: false, MaxRepeat: 9,
			},
			{Segment: "CTA", Mandatory: false, MaxRepeat: 1},  // Contact information
			{Segment: "COM", Mandatory: false, MaxRepeat: 9},  // Communication contact
			{Segment: "DTM", Mandatory: false, MaxRepeat: 9},  // Date/time/period
			{Segment: "CNT", Mandatory: false, MaxRepeat: 9},  // Control total
			{Segment: "TSR", Mandatory: false, MaxRepeat: 9},  // Transport service requirements
			{Segment: "CUX", Mandatory: false, MaxRepeat: 9},  // Currencies
			{Segment: "PCD", Mandatory: false, MaxRepeat: 9},  // Percentage details
			{Segment: "MOA", Mandatory: false, MaxRepeat: 99}, // Monetary amount
			{Segment: "FTX", Mandatory: false, MaxRepeat: 99}, // Free text
			{Segment: "GDS", Mandatory: false, MaxRepeat: 9},  // Nature of cargo
			{ // Segment group 33
				Group: []SchemaNode{
					{Segment: "LOC", Mandatory: true, MaxRepeat: 1},  // Place/location identification
					{Segment: "DTM", Mandatory: false, MaxRepeat: 9}, // Date/time/period
				},
				Mandatory: false, MaxRepeat: 99,
			},
			{ // Segment group 34
				Group: []SchemaNode{
					{Segment: "TOD", Mandatory: true, MaxRepeat: 1},  // Terms of delivery or transport
					{Segment: "LOC", Mandatory: false, MaxRepeat: 9}, // Place/location identification
				},
				Mandatory: false, MaxRepeat: 2,
			},
			{ // Segment group 35
				Group: []SchemaNode{
					{Segment: "RFF", Mandatory: true, MaxRepeat: 1},  // Reference
					{Segment: "DTM", Mandatory: false, MaxRepeat: 9}, // Date/time/period
				},
				Mandatory: false, MaxRepeat: 999,
			},
			{ // Segment group 36
				Group: []SchemaNode{
					{Segment: "GOR", Mandatory: true, MaxRepeat: 1},  // Governmental requirements
					{Segment: "DTM", Mandatory: false, MaxRepeat: 9}, // Date/time/period
					{Segment: "LOC", Mandatory: false, MaxRepeat: 9}, // Place/location identification
					{Segment: "SEL", Mandatory: false, MaxRepeat: 9}, // Seal number
					{Segment: "FTX", Mandatory: false, MaxRepeat: 9}, // Free text
					{ // Segment group 37
						Group: []SchemaNode{
							{Segment: "DOC", Mandatory: true, MaxRepeat: 1},  // Document/message details
							{Segment: "DTM", Mandatory: false, MaxRepeat: 1}, // Date/time/period
						},
						Mandatory: false, MaxRepeat: 9,
					},
				},
				Mandatory: false, MaxRepeat: 9,
			},
			{ // Segment group 38
				Group: []SchemaNode{
					{Segment: "CPI", Mandatory: true, MaxRepeat: 1},   // Charge payment instructions
					{Segment: "RFF", Mandatory: false, MaxRepeat: 99}, // Reference
					{Segment: "CUX", Mandatory: false, MaxRepeat: 1},  // Currencies
					{Segment: "LOC", Mandatory: false, MaxRepeat: 9},  // Place/location identification
					{Segment: "MOA", Mandatory: false, MaxRepeat: 9},  // Monetary amount
				},
				Mandatory: false, MaxRepeat: 9,
			},
			{ // Segment group 39
				Group: []SchemaNode{
					{Segment: "TCC", Mandatory: true, MaxRepeat: 1},  // Charge/rate calculations
					{Segment: "LOC", Mandatory: false, MaxRepeat: 1}, // Place/location identification
					{Segment: "FTX", Mandatory: false, MaxRepeat: 1}, // Free text
					{Segment: "CUX", Mandatory: false, MaxRepeat: 1}, // Currencies
					{Segment: "PRI", Mandatory: false, MaxRepeat: 1}, // Price details
					{Segment: "EQN", Mandatory: false, MaxRepeat: 1}, // Number of units
					{Segment: "PCD", Mandatory: false, MaxRepeat: 1}, // Percentage details
					{Segment: "MOA", Mandatory: false, MaxRepeat: 9}, // Monetary amount
					{Segment: "QTY", Mandatory: false, MaxRepeat: 9}, // Quantity
				},
				Mandatory: false, MaxRepeat: 99,
			},
			{ // Segment group 40
				Group: []SchemaNode{
					{Segment: "ICD", Mandatory: true, MaxRepeat: 1},  // Insurance cover description
					{Segment: "DTM", Mandatory: false, MaxRepeat: 1}, // Date/time/period
					{Segment: "FTX", Mandatory: false, MaxRepeat: 9}, // Free text
				},
				Mandatory: false, MaxRepeat: 99,
			},
			{ // Segment group 41
				Group: []SchemaNode{
					{Segment: "TDT", Mandatory: true, MaxRepeat: 1},  // Transport information
					{Segment: "DTM", Mandatory: false, MaxRepeat: 9}, // Date/time/period
					{ // Segment group 42
						Group: []SchemaNode{
							{Segment: "TSR", Mandatory: true, MaxRepeat: 1},  // Transport service requirements
							{Segment: "SCC", Mandatory: false, MaxRepeat: 9}, // Scheduling conditions
						},
						Mandatory: false, MaxRepeat: 9,
					},
					{ // Segment group 43
						Group: []SchemaNode{
							{Segment: "LOC", Mandatory: true, MaxRepeat: 1},  // Place/location identification
							{Segment: "DTM", Mandatory: false, MaxRepeat: 9}, // Date/time/period
						},
						Mandatory: false, MaxRepeat: 99,
					},
					{ // Segment group 44
						Group: []SchemaNode{
							{Segment: "RFF", Mandatory: true, MaxRepeat: 1},  // Reference
							{Segment: "DTM", Mandatory: false, MaxRepeat: 1}, // Date/time/period
						},
						Mandatory: false, MaxRepeat: 9,
					},
					{ // Segment group 45
						Group: []SchemaNode{
							{Segment: "TCC", Mandatory: true, MaxRepeat: 1},  // Charge/rate calculations
							{Segment: "MOA", Mandatory: false, MaxRepeat: 9}, // Monetary amount
							{Segment: "PCD", Mandatory: false, MaxRepeat: 1}, // Percentage details
						},
						Mandatory: false, MaxRepeat: 99,
					},
					{ // Segment group 46
						Group: []SchemaNode{
							{Segment: "ICD", Mandatory: true, MaxRepeat: 1},  // Insurance cover description
							{Segment: "DTM", Mandatory: false, MaxRepeat: 1}, // Date/time/period
							{Segment: "FTX", Mandatory: false, MaxRepeat: 9}, // Free text
						},
						Mandatory: false, MaxRepeat: 99,
					},
				},
				Mandatory: false, MaxRepeat: 99,
			},
			{ // Segment group 47
				Group: []SchemaNode{
					{Segment: "NAD", Mandatory: true, MaxRepeat: 1},  // Name and address
					{Segment: "LOC", Mandatory: false, MaxRepeat: 9}, // Place/location identification
					{Segment: "MOA", Mandatory: false, MaxRepeat: 9}, // Monetary amount
					{ // Segment group 48
						Group: []SchemaNode{
							{Segment: "CTA", Mandatory: true, MaxRepeat: 1},  // Contact information
							{Segment: "COM", Mandatory: false, MaxRepeat: 9}, // Communication contact
						},
						Mandatory: false, MaxRepeat: 9,
					},
					{ // Segment group 49
						Group: []SchemaNode{
							{Segment: "DOC", Mandatory: true, MaxRepeat: 1},  // Document/message details
							{Segment: "DTM", Mandatory: false, MaxRepeat: 1}, // Date/time/period
						},
						Mandatory: false, MaxRepeat: 9,
					},
					{ // Segment group 50
						Group: []SchemaNode{
							{Segment: "TCC", Mandatory: true, MaxRepeat: 1},  // Charge/rate calculations
							{Segment: "CUX", Mandatory: false, MaxRepeat: 1}, // Currencies
							{Segment: "PRI", Mandatory: false, MaxRepeat: 1}, // Price details
							{Segment: "EQN", Mandatory: false, MaxRepeat: 1}, // Number of units
							{Segment: "PCD", Mandatory: false, MaxRepeat: 1}, // Percentage details
							{Segment: "MOA", Mandatory: false, MaxRepeat: 9}, // Monetary amount
							{Segment: "QTY", Mandatory: false, MaxRepeat: 9}, // Quantity
						},
						Mandatory: false, MaxRepeat: 99,
					},
					{ // Segment group 51
						Group: []SchemaNode{
							{Segment: "RFF", Mandatory: true, MaxRepeat: 1},  // Reference
							{Segment: "DTM", Mandatory: false, MaxRepeat: 9}, // Date/time/period
						},
						Mandatory: false, MaxRepeat: 9,
					},
					{ // Segment group 52
						Group: []SchemaNode{
							{Segment: "CPI", Mandatory: true, MaxRepeat: 1},   // Charge payment instructions
							{Segment: "RFF", Mandatory: false, MaxRepeat: 99}, // Reference
							{Segment: "CUX", Mandatory: false, MaxRepeat: 1},  // Currencies
							{Segment: "LOC", Mandatory: false, MaxRepeat: 9},  // Place/location identification
							{Segment: "MOA", Mandatory: false, MaxRepeat: 9},  // Monetary amount
						},
						Mandatory: false, MaxRepeat: 9,
					},
					{ // Segment group 53
						Group: []SchemaNode{
							{Segment: "TSR", Mandatory: true, MaxRepeat: 1},  // Transport service requirements
							{Segment: "RFF", Mandatory: false, MaxRepeat: 1}, // Reference
							{Segment: "LOC", Mandatory: false, MaxRepeat: 1}, // Place/location identification
							{Segment: "TPL", Mandatory: false, MaxRepeat: 1}, // Transport placement
							{Segment: "FTX", Mandatory: false, MaxRepeat: 9}, // Free text
						},
						Mandatory: false, MaxRepeat: 99,
					},
				},
				Mandatory: false, MaxRepeat: 99,
			},
			{ // Segment group 54
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
			{ // Segment group 55
				Group: []SchemaNode{
					{Segment: "GID", Mandatory: true, MaxRepeat: 1},   // Goods item details
					{Segment: "HAN", Mandatory: false, MaxRepeat: 1},  // Handling instructions
					{Segment: "TMP", Mandatory: false, MaxRepeat: 9},  // Temperature
					{Segment: "RNG", Mandatory: false, MaxRepeat: 9},  // Range details
					{Segment: "TMD", Mandatory: false, MaxRepeat: 1},  // Transport movement details
					{Segment: "LOC", Mandatory: false, MaxRepeat: 9},  // Place/location identification
					{Segment: "MOA", Mandatory: false, MaxRepeat: 9},  // Monetary amount
					{Segment: "PIA", Mandatory: false, MaxRepeat: 9},  // Additional product id
					{Segment: "GIN", Mandatory: false, MaxRepeat: 99}, // Goods identity number
					{Segment: "FTX", Mandatory: false, MaxRepeat: 9},  // Free text
					{ // Segment group 56
						Group: []SchemaNode{
							{Segment: "NAD", Mandatory: true, MaxRepeat: 1},  // Name and address
							{Segment: "DTM", Mandatory: false, MaxRepeat: 1}, // Date/time/period
						},
						Mandatory: false, MaxRepeat: 9,
					},
					{Segment: "GDS", Mandatory: false, MaxRepeat: 9}, // Nature of cargo
					{ // Segment group 57
						Group: []SchemaNode{
							{Segment: "MEA", Mandatory: true, MaxRepeat: 1},  // Measurements
							{Segment: "EQN", Mandatory: false, MaxRepeat: 1}, // Number of units
						},
						Mandatory: false, MaxRepeat: 99,
					},
					{ // Segment group 58
						Group: []SchemaNode{
							{Segment: "DIM", Mandatory: true, MaxRepeat: 1},  // Dimensions
							{Segment: "EQN", Mandatory: false, MaxRepeat: 1}, // Number of units
						},
						Mandatory: false, MaxRepeat: 99,
					},
					{ // Segment group 59
						Group: []SchemaNode{
							{Segment: "RFF", Mandatory: true, MaxRepeat: 1},  // Reference
							{Segment: "DTM", Mandatory: false, MaxRepeat: 9}, // Date/time/period
						},
						Mandatory: false, MaxRepeat: 9,
					},
					{ // Segment group 60
						Group: []SchemaNode{
							{Segment: "PCI", Mandatory: true, MaxRepeat: 1},     // Package identification
							{Segment: "RFF", Mandatory: false, MaxRepeat: 1},    // Reference
							{Segment: "DTM", Mandatory: false, MaxRepeat: 1},    // Date/time/period
							{Segment: "GIN", Mandatory: false, MaxRepeat: 9999}, // Goods identity number
							{Segment: "MEA", Mandatory: false, MaxRepeat: 9},    // Measurements
							{Segment: "DIM", Mandatory: false, MaxRepeat: 1},    // Dimensions
							{Segment: "FTX", Mandatory: false, MaxRepeat: 9},    // Free text
						},
						Mandatory: false, MaxRepeat: 999,
					},
					{ // Segment group 61
						Group: []SchemaNode{
							{Segment: "DOC", Mandatory: true, MaxRepeat: 1},  // Document/message details
							{Segment: "DTM", Mandatory: false, MaxRepeat: 9}, // Date/time/period
						},
						Mandatory: false, MaxRepeat: 9,
					},
					{ // Segment group 62
						Group: []SchemaNode{
							{Segment: "GOR", Mandatory: true, MaxRepeat: 1},  // Governmental requirements
							{Segment: "DTM", Mandatory: false, MaxRepeat: 9}, // Date/time/period
							{Segment: "LOC", Mandatory: false, MaxRepeat: 9}, // Place/location identification
							{Segment: "SEL", Mandatory: false, MaxRepeat: 9}, // Seal number
							{Segment: "FTX", Mandatory: false, MaxRepeat: 9}, // Free text
							{ // Segment group 63
								Group: []SchemaNode{
									{Segment: "DOC", Mandatory: true, MaxRepeat: 1},  // Document/message details
									{Segment: "DTM", Mandatory: false, MaxRepeat: 1}, // Date/time/period
								},
								Mandatory: false, MaxRepeat: 9,
							},
						},
						Mandatory: false, MaxRepeat: 9,
					},
					{ // Segment group 64
						Group: []SchemaNode{
							{Segment: "TPL", Mandatory: true, MaxRepeat: 1}, // Transport placement
							{ // Segment group 65
								Group: []SchemaNode{
									{Segment: "MEA", Mandatory: true, MaxRepeat: 1},  // Measurements
									{Segment: "EQN", Mandatory: false, MaxRepeat: 1}, // Number of units
								},
								Mandatory: false, MaxRepeat: 9,
							},
						},
						Mandatory: false, MaxRepeat: 9,
					},
					{ // Segment group 66
						Group: []SchemaNode{
							{Segment: "SGP", Mandatory: true, MaxRepeat: 1},  // Split goods placement
							{Segment: "SEQ", Mandatory: false, MaxRepeat: 1}, // Sequence details
							{ // Segment group 67
								Group: []SchemaNode{
									{Segment: "MEA", Mandatory: true, MaxRepeat: 1},  // Measurements
									{Segment: "EQN", Mandatory: false, MaxRepeat: 1}, // Number of units
								},
								Mandatory: false, MaxRepeat: 9,
							},
						},
						Mandatory: false, MaxRepeat: 999,
					},
					{ // Segment group 68
						Group: []SchemaNode{
							{Segment: "TCC", Mandatory: true, MaxRepeat: 1},  // Charge/rate calculations
							{Segment: "CUX", Mandatory: false, MaxRepeat: 1}, // Currencies
							{Segment: "PRI", Mandatory: false, MaxRepeat: 1}, // Price details
							{Segment: "EQN", Mandatory: false, MaxRepeat: 1}, // Number of units
							{Segment: "PCD", Mandatory: false, MaxRepeat: 1}, // Percentage details
							{Segment: "MOA", Mandatory: false, MaxRepeat: 9}, // Monetary amount
							{Segment: "QTY", Mandatory: false, MaxRepeat: 9}, // Quantity
							{Segment: "LOC", Mandatory: false, MaxRepeat: 9}, // Place/location identification
						},
						Mandatory: false, MaxRepeat: 99,
					},
					{ // Segment group 69
						Group: []SchemaNode{
							{Segment: "ICD", Mandatory: true, MaxRepeat: 1},  // Insurance cover description
							{Segment: "DTM", Mandatory: false, MaxRepeat: 1}, // Date/time/period
							{Segment: "FTX", Mandatory: false, MaxRepeat: 9}, // Free text
						},
						Mandatory: false, MaxRepeat: 99,
					},
					{ // Segment group 70
						Group: []SchemaNode{
							{Segment: "DGS", Mandatory: true, MaxRepeat: 1},   // Dangerous goods
							{Segment: "FTX", Mandatory: false, MaxRepeat: 99}, // Free text
							{ // Segment group 71
								Group: []SchemaNode{
									{Segment: "CTA", Mandatory: true, MaxRepeat: 1},  // Contact information
									{Segment: "COM", Mandatory: false, MaxRepeat: 9}, // Communication contact
								},
								Mandatory: false, MaxRepeat: 9,
							},
							{ // Segment group 72
								Group: []SchemaNode{
									{Segment: "MEA", Mandatory: true, MaxRepeat: 1},  // Measurements
									{Segment: "EQN", Mandatory: false, MaxRepeat: 1}, // Number of units
								},
								Mandatory: false, MaxRepeat: 9,
							},
							{ // Segment group 73
								Group: []SchemaNode{
									{Segment: "SGP", Mandatory: true, MaxRepeat: 1}, // Split goods placement
									{ // Segment group 74
										Group: []SchemaNode{
											{Segment: "MEA", Mandatory: true, MaxRepeat: 1},  // Measurements
											{Segment: "EQN", Mandatory: false, MaxRepeat: 1}, // Number of units
										},
										Mandatory: false, MaxRepeat: 9,
									},
								},
								Mandatory: false, MaxRepeat: 999,
							},
						},
						Mandatory: false, MaxRepeat: 99,
					},
				},
				Mandatory: false, MaxRepeat: 99999,
			},
			{ // Segment group 75
				Group: []SchemaNode{
					{Segment: "EQD", Mandatory: true, MaxRepeat: 1},   // Equipment details
					{Segment: "EQN", Mandatory: false, MaxRepeat: 1},  // Number of units
					{Segment: "TMD", Mandatory: false, MaxRepeat: 1},  // Transport movement details
					{Segment: "MEA", Mandatory: false, MaxRepeat: 9},  // Measurements
					{Segment: "DIM", Mandatory: false, MaxRepeat: 9},  // Dimensions
					{Segment: "SEL", Mandatory: false, MaxRepeat: 99}, // Seal number
					{Segment: "TPL", Mandatory: false, MaxRepeat: 9},  // Transport placement
					{Segment: "HAN", Mandatory: false, MaxRepeat: 1},  // Handling instructions
					{Segment: "TMP", Mandatory: false, MaxRepeat: 1},  // Temperature
					{Segment: "RNG", Mandatory: false, MaxRepeat: 1},  // Range details
					{Segment: "FTX", Mandatory: false, MaxRepeat: 9},  // Free text
					{Segment: "PCD", Mandatory: false, MaxRepeat: 9},  // Percentage details
					{ // Segment group 76
						Group: []SchemaNode{
							{Segment: "TCC", Mandatory: true, MaxRepeat: 1},  // Charge/rate calculations
							{Segment: "CUX", Mandatory: false, MaxRepeat: 1}, // Currencies
							{Segment: "PRI", Mandatory: false, MaxRepeat: 1}, // Price details
							{Segment: "EQN", Mandatory: false, MaxRepeat: 1}, // Number of units
							{Segment: "PCD", Mandatory: false, MaxRepeat: 1}, // Percentage details
							{Segment: "MOA", Mandatory: false, MaxRepeat: 9}, // Monetary amount
							{Segment: "QTY", Mandatory: false, MaxRepeat: 9}, // Quantity
						},
						Mandatory: false, MaxRepeat: 99,
					},
					{ // Segment group 77
						Group: []SchemaNode{
							{Segment: "NAD", Mandatory: true, MaxRepeat: 1},  // Name and address
							{Segment: "DTM", Mandatory: false, MaxRepeat: 1}, // Date/time/period
						},
						Mandatory: false, MaxRepeat: 9,
					},
					{ // Segment group 78
						Group: []SchemaNode{
							{Segment: "EQA", Mandatory: true, MaxRepeat: 1},  // Attached equipment
							{Segment: "EQN", Mandatory: false, MaxRepeat: 1}, // Number of units
						},
						Mandatory: false, MaxRepeat: 99,
					},
					{ // Segment group 79
						Group: []SchemaNode{
							{Segment: "DGS", Mandatory: true, MaxRepeat: 1},  // Dangerous goods
							{Segment: "FTX", Mandatory: false, MaxRepeat: 9}, // Free text
							{ // Segment group 80
								Group: []SchemaNode{
									{Segment: "CTA", Mandatory: true, MaxRepeat: 1},  // Contact information
									{Segment: "COM", Mandatory: false, MaxRepeat: 9}, // Communication contact
								},
								Mandatory: false, MaxRepeat: 9,
							},
						},
						Mandatory: false, MaxRepeat: 99,
					},
				},
				Mandatory: false, MaxRepeat: 999,
			},
		},
		Mandatory: false, MaxRepeat: 9999,
	},
}}
