package edifact

// RECADV D.20A message specification: the structural (segment/group
// presence, order, cardinality) branching diagram for this message,
// UN/EDIFACT directory release D.20A. Scope is structural only -- not
// element-level content -- per edifact-ls-3uzr's epic non-goal.
// 38 segment groups, max nesting depth 5.
//
// Source: https://service.unece.org/trade/untdid/d20a/trmd/recadv_c.htm
// section 4.3.1 "Segment table". That URL currently returns HTTP 403 from
// Cloudflare when fetched directly; this data was transcribed from the
// Wayback Machine's archived copy instead:
// http://web.archive.org/web/20230924163522/https://service.unece.org/trade/untdid/d20a/trmd/recadv_c.htm
// -- re-check the direct URL first if cross-checking later, in case the
// block has lifted.
//
// Transcription note: generated in bulk (edifact-ls-13gu) the same way
// as every other message-type schema -- a script parsed the source's
// exact rail-art column positions mechanically and verified the result
// balances (all 38 groups open and close consistently)
// before being transcribed here, rather than reading the ASCII nesting
// by eye. Not individually spot-checked against the raw source the way
// the first several message types in this project were -- see
// edifact-ls-13gu for which representative samples were spot-checked
// for this batch.

func init() {
	RegisterSchema(
		MessageID{Type: "RECADV", Version: "D", Release: "20A", Agency: "UN"},
		recadvD20aSchema,
		"https://service.unece.org/trade/untdid/d20a/trmd/recadv_c.htm",
	)
}

var recadvD20aSchema = Schema{Nodes: []SchemaNode{
	{Segment: "BGM", Mandatory: true, MaxRepeat: 1},   // Beginning of message
	{Segment: "DTM", Mandatory: true, MaxRepeat: 10},  // Date/time/period
	{Segment: "ALI", Mandatory: false, MaxRepeat: 5},  // Additional information
	{Segment: "CUX", Mandatory: false, MaxRepeat: 9},  // Currencies
	{Segment: "FTX", Mandatory: false, MaxRepeat: 99}, // Free text
	{ // Segment group 1
		Group: []SchemaNode{
			{Segment: "RFF", Mandatory: true, MaxRepeat: 1},   // Reference
			{Segment: "DTM", Mandatory: false, MaxRepeat: 9},  // Date/time/period
			{Segment: "GEI", Mandatory: false, MaxRepeat: 99}, // Processing information
			{Segment: "MOA", Mandatory: false, MaxRepeat: 99}, // Monetary amount
			{Segment: "ALC", Mandatory: false, MaxRepeat: 1},  // Allowance or charge
		},
		Mandatory: false, MaxRepeat: 10,
	},
	{ // Segment group 2
		Group: []SchemaNode{
			{Segment: "DOC", Mandatory: true, MaxRepeat: 1}, // Document/message details
			{ // Segment group 3
				Group: []SchemaNode{
					{Segment: "CDI", Mandatory: true, MaxRepeat: 1},  // Physical or logical state
					{Segment: "INP", Mandatory: false, MaxRepeat: 5}, // Parties and instruction
				},
				Mandatory: false, MaxRepeat: 10,
			},
		},
		Mandatory: false, MaxRepeat: 10,
	},
	{ // Segment group 4
		Group: []SchemaNode{
			{Segment: "NAD", Mandatory: true, MaxRepeat: 1},   // Name and address
			{Segment: "LOC", Mandatory: false, MaxRepeat: 10}, // Place/location identification
			{ // Segment group 5
				Group: []SchemaNode{
					{Segment: "RFF", Mandatory: true, MaxRepeat: 1},  // Reference
					{Segment: "DTM", Mandatory: false, MaxRepeat: 9}, // Date/time/period
				},
				Mandatory: false, MaxRepeat: 10,
			},
			{ // Segment group 6
				Group: []SchemaNode{
					{Segment: "CTA", Mandatory: true, MaxRepeat: 1},  // Contact information
					{Segment: "COM", Mandatory: false, MaxRepeat: 5}, // Communication contact
				},
				Mandatory: false, MaxRepeat: 10,
			},
		},
		Mandatory: true, MaxRepeat: 99,
	},
	{ // Segment group 7
		Group: []SchemaNode{
			{Segment: "TOD", Mandatory: true, MaxRepeat: 1}, // Terms of delivery or transport
			{ // Segment group 8
				Group: []SchemaNode{
					{Segment: "CDI", Mandatory: true, MaxRepeat: 1},  // Physical or logical state
					{Segment: "INP", Mandatory: false, MaxRepeat: 5}, // Parties and instruction
				},
				Mandatory: false, MaxRepeat: 10,
			},
			{ // Segment group 9
				Group: []SchemaNode{
					{Segment: "LOC", Mandatory: true, MaxRepeat: 1},   // Place/location identification
					{Segment: "CDI", Mandatory: false, MaxRepeat: 10}, // Physical or logical state
				},
				Mandatory: false, MaxRepeat: 10,
			},
		},
		Mandatory: false, MaxRepeat: 10,
	},
	{ // Segment group 10
		Group: []SchemaNode{
			{Segment: "TDT", Mandatory: true, MaxRepeat: 1},   // Transport information
			{Segment: "DTM", Mandatory: false, MaxRepeat: 10}, // Date/time/period
			{Segment: "CDI", Mandatory: false, MaxRepeat: 20}, // Physical or logical state
		},
		Mandatory: false, MaxRepeat: 10,
	},
	{ // Segment group 11
		Group: []SchemaNode{
			{Segment: "EQD", Mandatory: true, MaxRepeat: 1}, // Equipment details
			{ // Segment group 12
				Group: []SchemaNode{
					{Segment: "CDI", Mandatory: true, MaxRepeat: 1},  // Physical or logical state
					{Segment: "INP", Mandatory: false, MaxRepeat: 5}, // Parties and instruction
				},
				Mandatory: false, MaxRepeat: 10,
			},
			{ // Segment group 13
				Group: []SchemaNode{
					{Segment: "SEL", Mandatory: true, MaxRepeat: 1},  // Seal number
					{Segment: "CDI", Mandatory: true, MaxRepeat: 10}, // Physical or logical state
				},
				Mandatory: false, MaxRepeat: 25,
			},
			{ // Segment group 14
				Group: []SchemaNode{
					{Segment: "EQA", Mandatory: true, MaxRepeat: 1}, // Attached equipment
					{ // Segment group 15
						Group: []SchemaNode{
							{Segment: "CDI", Mandatory: true, MaxRepeat: 1},  // Physical or logical state
							{Segment: "INP", Mandatory: false, MaxRepeat: 5}, // Parties and instruction
						},
						Mandatory: false, MaxRepeat: 10,
					},
				},
				Mandatory: false, MaxRepeat: 10,
			},
		},
		Mandatory: false, MaxRepeat: 9999,
	},
	{ // Segment group 16
		Group: []SchemaNode{
			{Segment: "CPS", Mandatory: true, MaxRepeat: 1}, // Consignment packing sequence
			{ // Segment group 17
				Group: []SchemaNode{
					{Segment: "RFF", Mandatory: true, MaxRepeat: 1},  // Reference
					{Segment: "DTM", Mandatory: false, MaxRepeat: 9}, // Date/time/period
				},
				Mandatory: false, MaxRepeat: 9,
			},
			{ // Segment group 18
				Group: []SchemaNode{
					{Segment: "TDT", Mandatory: true, MaxRepeat: 1},  // Transport information
					{Segment: "DTM", Mandatory: false, MaxRepeat: 9}, // Date/time/period
				},
				Mandatory: false, MaxRepeat: 9,
			},
			{ // Segment group 19
				Group: []SchemaNode{
					{Segment: "NAD", Mandatory: true, MaxRepeat: 1},  // Name and address
					{Segment: "LOC", Mandatory: false, MaxRepeat: 9}, // Place/location identification
				},
				Mandatory: false, MaxRepeat: 9,
			},
			{ // Segment group 20
				Group: []SchemaNode{
					{Segment: "PAC", Mandatory: true, MaxRepeat: 1},  // Package
					{Segment: "QVR", Mandatory: false, MaxRepeat: 1}, // Quantity variances
					{ // Segment group 21
						Group: []SchemaNode{
							{Segment: "PCI", Mandatory: true, MaxRepeat: 1},  // Package identification
							{Segment: "RFF", Mandatory: false, MaxRepeat: 9}, // Reference
							{ // Segment group 22
								Group: []SchemaNode{
									{Segment: "CDI", Mandatory: true, MaxRepeat: 1},  // Physical or logical state
									{Segment: "INP", Mandatory: false, MaxRepeat: 5}, // Parties and instruction
								},
								Mandatory: false, MaxRepeat: 10,
							},
							{ // Segment group 23
								Group: []SchemaNode{
									{Segment: "GIN", Mandatory: true, MaxRepeat: 1}, // Goods identity number
									{ // Segment group 24
										Group: []SchemaNode{
											{Segment: "CDI", Mandatory: true, MaxRepeat: 1},  // Physical or logical state
											{Segment: "INP", Mandatory: false, MaxRepeat: 5}, // Parties and instruction
										},
										Mandatory: false, MaxRepeat: 10,
									},
								},
								Mandatory: false, MaxRepeat: 999,
							},
						},
						Mandatory: false, MaxRepeat: 999,
					},
				},
				Mandatory: false, MaxRepeat: 9999,
			},
			{ // Segment group 25
				Group: []SchemaNode{
					{Segment: "LIN", Mandatory: true, MaxRepeat: 1},   // Line item
					{Segment: "PIA", Mandatory: false, MaxRepeat: 10}, // Additional product id
					{Segment: "IMD", Mandatory: false, MaxRepeat: 99}, // Item description
					{Segment: "QTY", Mandatory: false, MaxRepeat: 10}, // Quantity
					{Segment: "QVR", Mandatory: false, MaxRepeat: 10}, // Quantity variances
					{Segment: "DTM", Mandatory: false, MaxRepeat: 9},  // Date/time/period
					{Segment: "PRI", Mandatory: false, MaxRepeat: 1},  // Price details
					{Segment: "PCD", Mandatory: false, MaxRepeat: 9},  // Percentage details
					{Segment: "FTX", Mandatory: false, MaxRepeat: 99}, // Free text
					{ // Segment group 26
						Group: []SchemaNode{
							{Segment: "NAD", Mandatory: true, MaxRepeat: 1},  // Name and address
							{Segment: "LOC", Mandatory: false, MaxRepeat: 9}, // Place/location identification
							{ // Segment group 27
								Group: []SchemaNode{
									{Segment: "CTA", Mandatory: true, MaxRepeat: 1},  // Contact information
									{Segment: "COM", Mandatory: false, MaxRepeat: 9}, // Communication contact
								},
								Mandatory: false, MaxRepeat: 99,
							},
						},
						Mandatory: false, MaxRepeat: 99,
					},
					{ // Segment group 28
						Group: []SchemaNode{
							{Segment: "CDI", Mandatory: true, MaxRepeat: 1},  // Physical or logical state
							{Segment: "INP", Mandatory: false, MaxRepeat: 5}, // Parties and instruction
						},
						Mandatory: false, MaxRepeat: 10,
					},
					{ // Segment group 29
						Group: []SchemaNode{
							{Segment: "DOC", Mandatory: true, MaxRepeat: 1}, // Document/message details
							{ // Segment group 30
								Group: []SchemaNode{
									{Segment: "CDI", Mandatory: true, MaxRepeat: 1},  // Physical or logical state
									{Segment: "INP", Mandatory: false, MaxRepeat: 5}, // Parties and instruction
								},
								Mandatory: false, MaxRepeat: 10,
							},
						},
						Mandatory: false, MaxRepeat: 10,
					},
					{ // Segment group 31
						Group: []SchemaNode{
							{Segment: "GIN", Mandatory: true, MaxRepeat: 1}, // Goods identity number
							{ // Segment group 32
								Group: []SchemaNode{
									{Segment: "CDI", Mandatory: true, MaxRepeat: 1},  // Physical or logical state
									{Segment: "INP", Mandatory: false, MaxRepeat: 5}, // Parties and instruction
								},
								Mandatory: false, MaxRepeat: 10,
							},
						},
						Mandatory: false, MaxRepeat: 99,
					},
					{ // Segment group 33
						Group: []SchemaNode{
							{Segment: "RFF", Mandatory: true, MaxRepeat: 1},   // Reference
							{Segment: "DTM", Mandatory: false, MaxRepeat: 9},  // Date/time/period
							{Segment: "GEI", Mandatory: false, MaxRepeat: 99}, // Processing information
							{Segment: "MOA", Mandatory: false, MaxRepeat: 99}, // Monetary amount
							{Segment: "ALC", Mandatory: false, MaxRepeat: 1},  // Allowance or charge
						},
						Mandatory: false, MaxRepeat: 10,
					},
					{ // Segment group 34
						Group: []SchemaNode{
							{Segment: "PCI", Mandatory: true, MaxRepeat: 1},  // Package identification
							{Segment: "QTY", Mandatory: false, MaxRepeat: 9}, // Quantity
							{Segment: "QVR", Mandatory: false, MaxRepeat: 1}, // Quantity variances
							{ // Segment group 35
								Group: []SchemaNode{
									{Segment: "CDI", Mandatory: true, MaxRepeat: 1},  // Physical or logical state
									{Segment: "INP", Mandatory: false, MaxRepeat: 5}, // Parties and instruction
								},
								Mandatory: false, MaxRepeat: 10,
							},
							{ // Segment group 36
								Group: []SchemaNode{
									{Segment: "GIN", Mandatory: true, MaxRepeat: 1}, // Goods identity number
									{ // Segment group 37
										Group: []SchemaNode{
											{Segment: "CDI", Mandatory: true, MaxRepeat: 1},  // Physical or logical state
											{Segment: "INP", Mandatory: false, MaxRepeat: 5}, // Parties and instruction
										},
										Mandatory: false, MaxRepeat: 10,
									},
								},
								Mandatory: false, MaxRepeat: 10,
							},
							{ // Segment group 38
								Group: []SchemaNode{
									{Segment: "RFF", Mandatory: true, MaxRepeat: 1},  // Reference
									{Segment: "DTM", Mandatory: false, MaxRepeat: 9}, // Date/time/period
								},
								Mandatory: false, MaxRepeat: 9,
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
	{Segment: "CNT", Mandatory: false, MaxRepeat: 1}, // Control total
}}
