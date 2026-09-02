package edifact

// ITRRPT D.20A message specification: the structural (segment/group
// presence, order, cardinality) branching diagram for this message,
// UN/EDIFACT directory release D.20A. Scope is structural only -- not
// element-level content -- per edifact-ls-3uzr's epic non-goal.
// 25 segment groups, max nesting depth 5.
//
// Source: https://service.unece.org/trade/untdid/d20a/trmd/itrrpt_c.htm
// section 4.3.1 "Segment table". That URL currently returns HTTP 403 from
// Cloudflare when fetched directly; this data was transcribed from the
// Wayback Machine's archived copy instead:
// http://web.archive.org/web/20231002215239/https://service.unece.org/trade/untdid/d20a/trmd/itrrpt_c.htm
// -- re-check the direct URL first if cross-checking later, in case the
// block has lifted.
//
// Transcription note: generated in bulk (edifact-ls-13gu) the same way
// as every other message-type schema -- a script parsed the source's
// exact rail-art column positions mechanically and verified the result
// balances (all 25 groups open and close consistently)
// before being transcribed here, rather than reading the ASCII nesting
// by eye. Not individually spot-checked against the raw source the way
// the first several message types in this project were -- see
// edifact-ls-13gu for which representative samples were spot-checked
// for this batch.

func init() {
	RegisterSchema(
		MessageID{Type: "ITRRPT", Version: "D", Release: "20A", Agency: "UN"},
		itrrptD20aSchema,
		"https://service.unece.org/trade/untdid/d20a/trmd/itrrpt_c.htm",
	)
}

var itrrptD20aSchema = Schema{Nodes: []SchemaNode{
	{Segment: "BGM", Mandatory: true, MaxRepeat: 1},   // Beginning of message
	{Segment: "DTM", Mandatory: false, MaxRepeat: 10}, // Date/time/period
	{Segment: "MEA", Mandatory: false, MaxRepeat: 5},  // Measurements
	{Segment: "CUX", Mandatory: false, MaxRepeat: 9},  // Currencies
	{ // Segment group 1
		Group: []SchemaNode{
			{Segment: "RFF", Mandatory: true, MaxRepeat: 1},  // Reference
			{Segment: "DTM", Mandatory: false, MaxRepeat: 5}, // Date/time/period
		},
		Mandatory: false, MaxRepeat: 10,
	},
	{ // Segment group 2
		Group: []SchemaNode{
			{Segment: "NAD", Mandatory: true, MaxRepeat: 1},   // Name and address
			{Segment: "LOC", Mandatory: false, MaxRepeat: 10}, // Place/location identification
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
				Mandatory: false, MaxRepeat: 10,
			},
		},
		Mandatory: false, MaxRepeat: 10,
	},
	{ // Segment group 5
		Group: []SchemaNode{
			{Segment: "TDT", Mandatory: true, MaxRepeat: 1},  // Transport information
			{Segment: "PCD", Mandatory: false, MaxRepeat: 5}, // Percentage details
			{ // Segment group 6
				Group: []SchemaNode{
					{Segment: "LOC", Mandatory: true, MaxRepeat: 1},   // Place/location identification
					{Segment: "DTM", Mandatory: false, MaxRepeat: 10}, // Date/time/period
				},
				Mandatory: false, MaxRepeat: 10,
			},
		},
		Mandatory: false, MaxRepeat: 10,
	},
	{ // Segment group 7
		Group: []SchemaNode{
			{Segment: "EQD", Mandatory: true, MaxRepeat: 1},   // Equipment details
			{Segment: "MEA", Mandatory: false, MaxRepeat: 5},  // Measurements
			{Segment: "SEL", Mandatory: false, MaxRepeat: 99}, // Seal number
			{Segment: "EQA", Mandatory: false, MaxRepeat: 5},  // Attached equipment
			{Segment: "PCD", Mandatory: false, MaxRepeat: 5},  // Percentage details
			{ // Segment group 8
				Group: []SchemaNode{
					{Segment: "HAN", Mandatory: true, MaxRepeat: 1},   // Handling instructions
					{Segment: "FTX", Mandatory: false, MaxRepeat: 10}, // Free text
				},
				Mandatory: false, MaxRepeat: 10,
			},
		},
		Mandatory: false, MaxRepeat: 10,
	},
	{ // Segment group 9
		Group: []SchemaNode{
			{Segment: "GEI", Mandatory: true, MaxRepeat: 1},   // Processing information
			{Segment: "CDI", Mandatory: false, MaxRepeat: 1},  // Physical or logical state
			{Segment: "DTM", Mandatory: false, MaxRepeat: 10}, // Date/time/period
			{Segment: "ALI", Mandatory: false, MaxRepeat: 5},  // Additional information
			{Segment: "MEA", Mandatory: false, MaxRepeat: 5},  // Measurements
			{Segment: "MOA", Mandatory: false, MaxRepeat: 5},  // Monetary amount
			{ // Segment group 10
				Group: []SchemaNode{
					{Segment: "RFF", Mandatory: true, MaxRepeat: 1},  // Reference
					{Segment: "DTM", Mandatory: false, MaxRepeat: 1}, // Date/time/period
				},
				Mandatory: true, MaxRepeat: 10,
			},
			{ // Segment group 11
				Group: []SchemaNode{
					{Segment: "NAD", Mandatory: true, MaxRepeat: 1},   // Name and address
					{Segment: "LOC", Mandatory: false, MaxRepeat: 10}, // Place/location identification
					{ // Segment group 12
						Group: []SchemaNode{
							{Segment: "RFF", Mandatory: true, MaxRepeat: 1},  // Reference
							{Segment: "DTM", Mandatory: false, MaxRepeat: 1}, // Date/time/period
						},
						Mandatory: false, MaxRepeat: 10,
					},
					{ // Segment group 13
						Group: []SchemaNode{
							{Segment: "CTA", Mandatory: true, MaxRepeat: 1},  // Contact information
							{Segment: "COM", Mandatory: false, MaxRepeat: 5}, // Communication contact
						},
						Mandatory: false, MaxRepeat: 10,
					},
				},
				Mandatory: false, MaxRepeat: 10,
			},
			{ // Segment group 14
				Group: []SchemaNode{
					{Segment: "TOD", Mandatory: true, MaxRepeat: 1},  // Terms of delivery or transport
					{Segment: "LOC", Mandatory: false, MaxRepeat: 5}, // Place/location identification
					{Segment: "FTX", Mandatory: false, MaxRepeat: 5}, // Free text
				},
				Mandatory: false, MaxRepeat: 10,
			},
			{ // Segment group 15
				Group: []SchemaNode{
					{Segment: "CPS", Mandatory: true, MaxRepeat: 1},  // Consignment packing sequence
					{Segment: "FTX", Mandatory: false, MaxRepeat: 5}, // Free text
					{ // Segment group 16
						Group: []SchemaNode{
							{Segment: "PAC", Mandatory: true, MaxRepeat: 1},   // Package
							{Segment: "MEA", Mandatory: false, MaxRepeat: 10}, // Measurements
							{Segment: "QTY", Mandatory: false, MaxRepeat: 10}, // Quantity
							{ // Segment group 17
								Group: []SchemaNode{
									{Segment: "HAN", Mandatory: true, MaxRepeat: 1},   // Handling instructions
									{Segment: "FTX", Mandatory: false, MaxRepeat: 10}, // Free text
								},
								Mandatory: false, MaxRepeat: 10,
							},
							{ // Segment group 18
								Group: []SchemaNode{
									{Segment: "PCI", Mandatory: true, MaxRepeat: 1},  // Package identification
									{Segment: "RFF", Mandatory: false, MaxRepeat: 1}, // Reference
									{Segment: "CDI", Mandatory: false, MaxRepeat: 1}, // Physical or logical state
									{Segment: "DTM", Mandatory: false, MaxRepeat: 5}, // Date/time/period
									{ // Segment group 19
										Group: []SchemaNode{
											{Segment: "GIN", Mandatory: true, MaxRepeat: 1},    // Goods identity number
											{Segment: "CDI", Mandatory: false, MaxRepeat: 1},   // Physical or logical state
											{Segment: "DLM", Mandatory: false, MaxRepeat: 999}, // Delivery limitations
										},
										Mandatory: false, MaxRepeat: 99,
									},
								},
								Mandatory: false, MaxRepeat: 9999,
							},
						},
						Mandatory: false, MaxRepeat: 9999,
					},
					{ // Segment group 20
						Group: []SchemaNode{
							{Segment: "LIN", Mandatory: true, MaxRepeat: 1},    // Line item
							{Segment: "PIA", Mandatory: false, MaxRepeat: 10},  // Additional product id
							{Segment: "IMD", Mandatory: false, MaxRepeat: 99},  // Item description
							{Segment: "MEA", Mandatory: false, MaxRepeat: 10},  // Measurements
							{Segment: "CDI", Mandatory: false, MaxRepeat: 1},   // Physical or logical state
							{Segment: "QTY", Mandatory: false, MaxRepeat: 10},  // Quantity
							{Segment: "ALI", Mandatory: false, MaxRepeat: 10},  // Additional information
							{Segment: "GIN", Mandatory: false, MaxRepeat: 999}, // Goods identity number
							{Segment: "GIR", Mandatory: false, MaxRepeat: 999}, // Related identification numbers
							{Segment: "DTM", Mandatory: false, MaxRepeat: 5},   // Date/time/period
							{Segment: "MOA", Mandatory: false, MaxRepeat: 1},   // Monetary amount
							{Segment: "FTX", Mandatory: false, MaxRepeat: 5},   // Free text
							{ // Segment group 21
								Group: []SchemaNode{
									{Segment: "RFF", Mandatory: true, MaxRepeat: 1},  // Reference
									{Segment: "DTM", Mandatory: false, MaxRepeat: 1}, // Date/time/period
								},
								Mandatory: false, MaxRepeat: 10,
							},
							{ // Segment group 22
								Group: []SchemaNode{
									{Segment: "DGS", Mandatory: true, MaxRepeat: 1},  // Dangerous goods
									{Segment: "QTY", Mandatory: false, MaxRepeat: 1}, // Quantity
									{Segment: "FTX", Mandatory: false, MaxRepeat: 5}, // Free text
								},
								Mandatory: false, MaxRepeat: 10,
							},
							{ // Segment group 23
								Group: []SchemaNode{
									{Segment: "PCI", Mandatory: true, MaxRepeat: 1},   // Package identification
									{Segment: "CDI", Mandatory: false, MaxRepeat: 1},  // Physical or logical state
									{Segment: "DTM", Mandatory: false, MaxRepeat: 5},  // Date/time/period
									{Segment: "QTY", Mandatory: false, MaxRepeat: 1},  // Quantity
									{Segment: "MEA", Mandatory: false, MaxRepeat: 10}, // Measurements
									{ // Segment group 24
										Group: []SchemaNode{
											{Segment: "GIN", Mandatory: true, MaxRepeat: 1},    // Goods identity number
											{Segment: "CDI", Mandatory: false, MaxRepeat: 1},   // Physical or logical state
											{Segment: "DLM", Mandatory: false, MaxRepeat: 999}, // Delivery limitations
										},
										Mandatory: false, MaxRepeat: 10,
									},
									{ // Segment group 25
										Group: []SchemaNode{
											{Segment: "HAN", Mandatory: true, MaxRepeat: 1},     // Handling instructions
											{Segment: "FTX", Mandatory: false, MaxRepeat: 5},    // Free text
											{Segment: "GIN", Mandatory: false, MaxRepeat: 9999}, // Goods identity number
										},
										Mandatory: false, MaxRepeat: 10,
									},
								},
								Mandatory: false, MaxRepeat: 9999,
							},
						},
						Mandatory: false, MaxRepeat: 1,
					},
				},
				Mandatory: false, MaxRepeat: 9999,
			},
		},
		Mandatory: true, MaxRepeat: 99,
	},
	{Segment: "CNT", Mandatory: false, MaxRepeat: 1}, // Control total
}}
