package edifact

// HANMOV D.20A message specification: the structural (segment/group
// presence, order, cardinality) branching diagram for this message,
// UN/EDIFACT directory release D.20A. Scope is structural only -- not
// element-level content -- per edifact-ls-3uzr's epic non-goal.
// 23 segment groups, max nesting depth 4.
//
// Source: https://service.unece.org/trade/untdid/d20a/trmd/hanmov_c.htm
// section 4.3.1 "Segment table". That URL currently returns HTTP 403 from
// Cloudflare when fetched directly; this data was transcribed from the
// Wayback Machine's archived copy instead:
// http://web.archive.org/web/20231002212733/https://service.unece.org/trade/untdid/d20a/trmd/hanmov_c.htm
// -- re-check the direct URL first if cross-checking later, in case the
// block has lifted.
//
// Transcription note: generated in bulk (edifact-ls-13gu) the same way
// as every other message-type schema -- a script parsed the source's
// exact rail-art column positions mechanically and verified the result
// balances (all 23 groups open and close consistently)
// before being transcribed here, rather than reading the ASCII nesting
// by eye. Not individually spot-checked against the raw source the way
// the first several message types in this project were -- see
// edifact-ls-13gu for which representative samples were spot-checked
// for this batch.

func init() {
	RegisterSchema(
		MessageID{Type: "HANMOV", Version: "D", Release: "20A", Agency: "UN"},
		hanmovD20aSchema,
		"https://service.unece.org/trade/untdid/d20a/trmd/hanmov_c.htm",
	)
}

var hanmovD20aSchema = Schema{Nodes: []SchemaNode{
	{Segment: "BGM", Mandatory: true, MaxRepeat: 1},  // Beginning of message
	{Segment: "DTM", Mandatory: false, MaxRepeat: 9}, // Date/time/period
	{Segment: "HAN", Mandatory: false, MaxRepeat: 9}, // Handling instructions
	{Segment: "FTX", Mandatory: false, MaxRepeat: 9}, // Free text
	{Segment: "CNT", Mandatory: false, MaxRepeat: 9}, // Control total
	{Segment: "CUX", Mandatory: false, MaxRepeat: 9}, // Currencies
	{ // Segment group 1
		Group: []SchemaNode{
			{Segment: "TCC", Mandatory: true, MaxRepeat: 1},  // Charge/rate calculations
			{Segment: "LOC", Mandatory: false, MaxRepeat: 9}, // Place/location identification
			{Segment: "FTX", Mandatory: false, MaxRepeat: 9}, // Free text
			{Segment: "CUX", Mandatory: false, MaxRepeat: 9}, // Currencies
		},
		Mandatory: false, MaxRepeat: 9,
	},
	{ // Segment group 2
		Group: []SchemaNode{
			{Segment: "RFF", Mandatory: true, MaxRepeat: 1},  // Reference
			{Segment: "DTM", Mandatory: false, MaxRepeat: 9}, // Date/time/period
			{ // Segment group 3
				Group: []SchemaNode{
					{Segment: "STS", Mandatory: true, MaxRepeat: 1},  // Status
					{Segment: "DTM", Mandatory: false, MaxRepeat: 9}, // Date/time/period
					{Segment: "FTX", Mandatory: false, MaxRepeat: 9}, // Free text
					{Segment: "HAN", Mandatory: false, MaxRepeat: 9}, // Handling instructions
				},
				Mandatory: false, MaxRepeat: 1,
			},
			{Segment: "MOA", Mandatory: false, MaxRepeat: 9}, // Monetary amount
		},
		Mandatory: false, MaxRepeat: 9,
	},
	{ // Segment group 4
		Group: []SchemaNode{
			{Segment: "NAD", Mandatory: true, MaxRepeat: 1},  // Name and address
			{Segment: "LOC", Mandatory: false, MaxRepeat: 9}, // Place/location identification
			{Segment: "DOC", Mandatory: false, MaxRepeat: 9}, // Document/message details
			{ // Segment group 5
				Group: []SchemaNode{
					{Segment: "RFF", Mandatory: true, MaxRepeat: 1},  // Reference
					{Segment: "DTM", Mandatory: false, MaxRepeat: 9}, // Date/time/period
				},
				Mandatory: false, MaxRepeat: 9,
			},
			{ // Segment group 6
				Group: []SchemaNode{
					{Segment: "CTA", Mandatory: true, MaxRepeat: 1},  // Contact information
					{Segment: "COM", Mandatory: false, MaxRepeat: 9}, // Communication contact
				},
				Mandatory: false, MaxRepeat: 1,
			},
		},
		Mandatory: false, MaxRepeat: 9,
	},
	{ // Segment group 7
		Group: []SchemaNode{
			{Segment: "TDT", Mandatory: true, MaxRepeat: 1}, // Transport information
			{ // Segment group 8
				Group: []SchemaNode{
					{Segment: "STS", Mandatory: true, MaxRepeat: 1},  // Status
					{Segment: "LOC", Mandatory: false, MaxRepeat: 9}, // Place/location identification
					{Segment: "DTM", Mandatory: false, MaxRepeat: 9}, // Date/time/period
					{ // Segment group 9
						Group: []SchemaNode{
							{Segment: "RFF", Mandatory: true, MaxRepeat: 1},  // Reference
							{Segment: "DTM", Mandatory: false, MaxRepeat: 9}, // Date/time/period
						},
						Mandatory: false, MaxRepeat: 1,
					},
					{ // Segment group 10
						Group: []SchemaNode{
							{Segment: "DOC", Mandatory: true, MaxRepeat: 1},  // Document/message details
							{Segment: "DTM", Mandatory: false, MaxRepeat: 9}, // Date/time/period
						},
						Mandatory: false, MaxRepeat: 1,
					},
				},
				Mandatory: false, MaxRepeat: 1,
			},
		},
		Mandatory: false, MaxRepeat: 9,
	},
	{ // Segment group 11
		Group: []SchemaNode{
			{Segment: "TOD", Mandatory: true, MaxRepeat: 1},  // Terms of delivery or transport
			{Segment: "LOC", Mandatory: false, MaxRepeat: 9}, // Place/location identification
		},
		Mandatory: false, MaxRepeat: 9,
	},
	{ // Segment group 12
		Group: []SchemaNode{
			{Segment: "LIN", Mandatory: true, MaxRepeat: 1}, // Line item
			{ // Segment group 13
				Group: []SchemaNode{
					{Segment: "STS", Mandatory: true, MaxRepeat: 1},   // Status
					{Segment: "PIA", Mandatory: false, MaxRepeat: 9},  // Additional product id
					{Segment: "IMD", Mandatory: false, MaxRepeat: 99}, // Item description
					{Segment: "HAN", Mandatory: false, MaxRepeat: 9},  // Handling instructions
					{Segment: "TCC", Mandatory: false, MaxRepeat: 9},  // Charge/rate calculations
					{ // Segment group 14
						Group: []SchemaNode{
							{Segment: "NAD", Mandatory: true, MaxRepeat: 1}, // Name and address
							{ // Segment group 15
								Group: []SchemaNode{
									{Segment: "CTA", Mandatory: true, MaxRepeat: 1},  // Contact information
									{Segment: "COM", Mandatory: false, MaxRepeat: 9}, // Communication contact
								},
								Mandatory: false, MaxRepeat: 9,
							},
						},
						Mandatory: false, MaxRepeat: 9,
					},
					{Segment: "MEA", Mandatory: false, MaxRepeat: 9},  // Measurements
					{Segment: "QTY", Mandatory: false, MaxRepeat: 9},  // Quantity
					{Segment: "GIN", Mandatory: false, MaxRepeat: 99}, // Goods identity number
					{Segment: "GIR", Mandatory: false, MaxRepeat: 99}, // Related identification numbers
					{ // Segment group 16
						Group: []SchemaNode{
							{Segment: "DTM", Mandatory: true, MaxRepeat: 1},  // Date/time/period
							{Segment: "QTY", Mandatory: false, MaxRepeat: 9}, // Quantity
						},
						Mandatory: false, MaxRepeat: 1,
					},
					{Segment: "FTX", Mandatory: false, MaxRepeat: 9}, // Free text
					{ // Segment group 17
						Group: []SchemaNode{
							{Segment: "RFF", Mandatory: true, MaxRepeat: 1},  // Reference
							{Segment: "DTM", Mandatory: false, MaxRepeat: 9}, // Date/time/period
						},
						Mandatory: false, MaxRepeat: 1,
					},
					{Segment: "LOC", Mandatory: false, MaxRepeat: 9}, // Place/location identification
					{ // Segment group 18
						Group: []SchemaNode{
							{Segment: "PAC", Mandatory: true, MaxRepeat: 1},  // Package
							{Segment: "MEA", Mandatory: false, MaxRepeat: 9}, // Measurements
							{Segment: "QTY", Mandatory: false, MaxRepeat: 9}, // Quantity
							{ // Segment group 19
								Group: []SchemaNode{
									{Segment: "PCI", Mandatory: true, MaxRepeat: 1},  // Package identification
									{Segment: "GIN", Mandatory: false, MaxRepeat: 9}, // Goods identity number
								},
								Mandatory: false, MaxRepeat: 9,
							},
						},
						Mandatory: false, MaxRepeat: 1,
					},
				},
				Mandatory: false, MaxRepeat: 1,
			},
		},
		Mandatory: false, MaxRepeat: 99,
	},
	{ // Segment group 20
		Group: []SchemaNode{
			{Segment: "GID", Mandatory: true, MaxRepeat: 1},  // Goods item details
			{Segment: "HAN", Mandatory: false, MaxRepeat: 9}, // Handling instructions
			{Segment: "CST", Mandatory: false, MaxRepeat: 9}, // Customs status of goods
			{Segment: "TMP", Mandatory: false, MaxRepeat: 9}, // Temperature
			{Segment: "RNG", Mandatory: false, MaxRepeat: 9}, // Range details
			{Segment: "LOC", Mandatory: false, MaxRepeat: 9}, // Place/location identification
			{Segment: "MOA", Mandatory: false, MaxRepeat: 9}, // Monetary amount
			{Segment: "PIA", Mandatory: false, MaxRepeat: 9}, // Additional product id
			{Segment: "FTX", Mandatory: false, MaxRepeat: 9}, // Free text
			{Segment: "NAD", Mandatory: false, MaxRepeat: 9}, // Name and address
			{Segment: "MEA", Mandatory: false, MaxRepeat: 9}, // Measurements
			{Segment: "DIM", Mandatory: false, MaxRepeat: 9}, // Dimensions
			{Segment: "RFF", Mandatory: false, MaxRepeat: 9}, // Reference
			{Segment: "PCI", Mandatory: false, MaxRepeat: 9}, // Package identification
			{ // Segment group 21
				Group: []SchemaNode{
					{Segment: "DOC", Mandatory: true, MaxRepeat: 1},  // Document/message details
					{Segment: "DTM", Mandatory: false, MaxRepeat: 9}, // Date/time/period
				},
				Mandatory: false, MaxRepeat: 9,
			},
			{Segment: "SGP", Mandatory: false, MaxRepeat: 9}, // Split goods placement
			{Segment: "TCC", Mandatory: false, MaxRepeat: 9}, // Charge/rate calculations
			{ // Segment group 22
				Group: []SchemaNode{
					{Segment: "DGS", Mandatory: true, MaxRepeat: 1},  // Dangerous goods
					{Segment: "FTX", Mandatory: false, MaxRepeat: 9}, // Free text
					{Segment: "MEA", Mandatory: false, MaxRepeat: 9}, // Measurements
				},
				Mandatory: false, MaxRepeat: 9,
			},
		},
		Mandatory: false, MaxRepeat: 9,
	},
	{ // Segment group 23
		Group: []SchemaNode{
			{Segment: "EQD", Mandatory: true, MaxRepeat: 1},   // Equipment details
			{Segment: "EQN", Mandatory: false, MaxRepeat: 9},  // Number of units
			{Segment: "SEL", Mandatory: false, MaxRepeat: 9},  // Seal number
			{Segment: "EQA", Mandatory: false, MaxRepeat: 99}, // Attached equipment
			{Segment: "RFF", Mandatory: false, MaxRepeat: 9},  // Reference
		},
		Mandatory: false, MaxRepeat: 9,
	},
}}
