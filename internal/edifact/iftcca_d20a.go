package edifact

// IFTCCA D.20A message specification: the structural (segment/group
// presence, order, cardinality) branching diagram for this message,
// UN/EDIFACT directory release D.20A. Scope is structural only -- not
// element-level content -- per edifact-ls-3uzr's epic non-goal.
// 23 segment groups, max nesting depth 4.
//
// Source: https://service.unece.org/trade/untdid/d20a/trmd/iftcca_c.htm
// section 4.3.1 "Segment table". That URL currently returns HTTP 403 from
// Cloudflare when fetched directly; this data was transcribed from the
// Wayback Machine's archived copy instead:
// http://web.archive.org/web/20231202090333/https://service.unece.org/trade/untdid/d20a/trmd/iftcca_c.htm
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
		MessageID{Type: "IFTCCA", Version: "D", Release: "20A", Agency: "UN"},
		iftccaD20aSchema,
		"https://service.unece.org/trade/untdid/d20a/trmd/iftcca_c.htm",
	)
}

var iftccaD20aSchema = Schema{Nodes: []SchemaNode{
	{Segment: "BGM", Mandatory: true, MaxRepeat: 1},   // Beginning of message
	{Segment: "CTA", Mandatory: false, MaxRepeat: 9},  // Contact information
	{Segment: "COM", Mandatory: false, MaxRepeat: 9},  // Communication contact
	{Segment: "DTM", Mandatory: false, MaxRepeat: 9},  // Date/time/period
	{Segment: "TSR", Mandatory: false, MaxRepeat: 9},  // Transport service requirements
	{Segment: "CUX", Mandatory: false, MaxRepeat: 9},  // Currencies
	{Segment: "MOA", Mandatory: false, MaxRepeat: 9},  // Monetary amount
	{Segment: "FTX", Mandatory: false, MaxRepeat: 99}, // Free text
	{ // Segment group 1
		Group: []SchemaNode{
			{Segment: "RFF", Mandatory: true, MaxRepeat: 1},  // Reference
			{Segment: "DTM", Mandatory: false, MaxRepeat: 9}, // Date/time/period
		},
		Mandatory: false, MaxRepeat: 9,
	},
	{Segment: "CNT", Mandatory: false, MaxRepeat: 9}, // Control total
	{ // Segment group 2
		Group: []SchemaNode{
			{Segment: "LOC", Mandatory: true, MaxRepeat: 1},  // Place/location identification
			{Segment: "DTM", Mandatory: false, MaxRepeat: 9}, // Date/time/period
		},
		Mandatory: false, MaxRepeat: 99,
	},
	{ // Segment group 3
		Group: []SchemaNode{
			{Segment: "CPI", Mandatory: true, MaxRepeat: 1},  // Charge payment instructions
			{Segment: "CUX", Mandatory: false, MaxRepeat: 9}, // Currencies
			{Segment: "LOC", Mandatory: false, MaxRepeat: 9}, // Place/location identification
			{Segment: "MOA", Mandatory: false, MaxRepeat: 9}, // Monetary amount
		},
		Mandatory: false, MaxRepeat: 9,
	},
	{ // Segment group 4
		Group: []SchemaNode{
			{Segment: "TDT", Mandatory: true, MaxRepeat: 1},   // Transport information
			{Segment: "DTM", Mandatory: false, MaxRepeat: 9},  // Date/time/period
			{Segment: "TSR", Mandatory: false, MaxRepeat: 9},  // Transport service requirements
			{Segment: "LOC", Mandatory: false, MaxRepeat: 99}, // Place/location identification
			{Segment: "FTX", Mandatory: false, MaxRepeat: 99}, // Free text
			{ // Segment group 5
				Group: []SchemaNode{
					{Segment: "RFF", Mandatory: true, MaxRepeat: 1},  // Reference
					{Segment: "DTM", Mandatory: false, MaxRepeat: 9}, // Date/time/period
				},
				Mandatory: false, MaxRepeat: 9,
			},
		},
		Mandatory: false, MaxRepeat: 99,
	},
	{ // Segment group 6
		Group: []SchemaNode{
			{Segment: "NAD", Mandatory: true, MaxRepeat: 1},  // Name and address
			{Segment: "LOC", Mandatory: false, MaxRepeat: 9}, // Place/location identification
			{ // Segment group 7
				Group: []SchemaNode{
					{Segment: "CTA", Mandatory: true, MaxRepeat: 1},  // Contact information
					{Segment: "COM", Mandatory: false, MaxRepeat: 9}, // Communication contact
				},
				Mandatory: false, MaxRepeat: 9,
			},
			{ // Segment group 8
				Group: []SchemaNode{
					{Segment: "RFF", Mandatory: true, MaxRepeat: 1},  // Reference
					{Segment: "DTM", Mandatory: false, MaxRepeat: 9}, // Date/time/period
				},
				Mandatory: false, MaxRepeat: 9,
			},
		},
		Mandatory: false, MaxRepeat: 99,
	},
	{ // Segment group 9
		Group: []SchemaNode{
			{Segment: "GID", Mandatory: true, MaxRepeat: 1},  // Goods item details
			{Segment: "HAN", Mandatory: false, MaxRepeat: 9}, // Handling instructions
			{Segment: "TMP", Mandatory: false, MaxRepeat: 9}, // Temperature
			{Segment: "RNG", Mandatory: false, MaxRepeat: 9}, // Range details
			{Segment: "LOC", Mandatory: false, MaxRepeat: 9}, // Place/location identification
			{Segment: "FTX", Mandatory: false, MaxRepeat: 9}, // Free text
			{ // Segment group 10
				Group: []SchemaNode{
					{Segment: "GDS", Mandatory: true, MaxRepeat: 1},  // Nature of cargo
					{Segment: "FTX", Mandatory: false, MaxRepeat: 9}, // Free text
				},
				Mandatory: false, MaxRepeat: 9,
			},
			{ // Segment group 11
				Group: []SchemaNode{
					{Segment: "MEA", Mandatory: true, MaxRepeat: 1},  // Measurements
					{Segment: "EQN", Mandatory: false, MaxRepeat: 9}, // Number of units
				},
				Mandatory: false, MaxRepeat: 99,
			},
			{ // Segment group 12
				Group: []SchemaNode{
					{Segment: "DIM", Mandatory: true, MaxRepeat: 1},  // Dimensions
					{Segment: "EQN", Mandatory: false, MaxRepeat: 9}, // Number of units
				},
				Mandatory: false, MaxRepeat: 99,
			},
			{ // Segment group 13
				Group: []SchemaNode{
					{Segment: "RFF", Mandatory: true, MaxRepeat: 1},  // Reference
					{Segment: "EQN", Mandatory: false, MaxRepeat: 9}, // Number of units
				},
				Mandatory: false, MaxRepeat: 9,
			},
			{ // Segment group 14
				Group: []SchemaNode{
					{Segment: "TPL", Mandatory: true, MaxRepeat: 1}, // Transport placement
					{ // Segment group 15
						Group: []SchemaNode{
							{Segment: "MEA", Mandatory: true, MaxRepeat: 1},  // Measurements
							{Segment: "EQN", Mandatory: false, MaxRepeat: 9}, // Number of units
						},
						Mandatory: false, MaxRepeat: 9,
					},
				},
				Mandatory: false, MaxRepeat: 9,
			},
			{ // Segment group 16
				Group: []SchemaNode{
					{Segment: "SGP", Mandatory: true, MaxRepeat: 1}, // Split goods placement
					{ // Segment group 17
						Group: []SchemaNode{
							{Segment: "MEA", Mandatory: true, MaxRepeat: 1},  // Measurements
							{Segment: "EQN", Mandatory: false, MaxRepeat: 9}, // Number of units
						},
						Mandatory: false, MaxRepeat: 9,
					},
				},
				Mandatory: false, MaxRepeat: 999,
			},
			{ // Segment group 18
				Group: []SchemaNode{
					{Segment: "TCC", Mandatory: true, MaxRepeat: 1},  // Charge/rate calculations
					{Segment: "PRI", Mandatory: false, MaxRepeat: 9}, // Price details
					{Segment: "EQN", Mandatory: false, MaxRepeat: 9}, // Number of units
					{Segment: "PCD", Mandatory: false, MaxRepeat: 9}, // Percentage details
					{Segment: "MOA", Mandatory: false, MaxRepeat: 9}, // Monetary amount
					{Segment: "QTY", Mandatory: false, MaxRepeat: 9}, // Quantity
					{Segment: "LOC", Mandatory: false, MaxRepeat: 9}, // Place/location identification
					{Segment: "RFF", Mandatory: false, MaxRepeat: 9}, // Reference
					{Segment: "MEA", Mandatory: false, MaxRepeat: 9}, // Measurements
					{Segment: "CUX", Mandatory: false, MaxRepeat: 9}, // Currencies
					{Segment: "DTM", Mandatory: false, MaxRepeat: 9}, // Date/time/period
					{Segment: "FTX", Mandatory: false, MaxRepeat: 9}, // Free text
				},
				Mandatory: false, MaxRepeat: 99,
			},
			{ // Segment group 19
				Group: []SchemaNode{
					{Segment: "DGS", Mandatory: true, MaxRepeat: 1},   // Dangerous goods
					{Segment: "FTX", Mandatory: false, MaxRepeat: 99}, // Free text
					{ // Segment group 20
						Group: []SchemaNode{
							{Segment: "MEA", Mandatory: true, MaxRepeat: 1},  // Measurements
							{Segment: "EQN", Mandatory: false, MaxRepeat: 1}, // Number of units
						},
						Mandatory: false, MaxRepeat: 9,
					},
					{ // Segment group 21
						Group: []SchemaNode{
							{Segment: "SGP", Mandatory: true, MaxRepeat: 1}, // Split goods placement
							{ // Segment group 22
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
				Mandatory: false, MaxRepeat: 9,
			},
		},
		Mandatory: false, MaxRepeat: 999,
	},
	{ // Segment group 23
		Group: []SchemaNode{
			{Segment: "EQD", Mandatory: true, MaxRepeat: 1},  // Equipment details
			{Segment: "EQN", Mandatory: false, MaxRepeat: 9}, // Number of units
			{Segment: "FTX", Mandatory: false, MaxRepeat: 9}, // Free text
			{Segment: "RFF", Mandatory: false, MaxRepeat: 9}, // Reference
		},
		Mandatory: false, MaxRepeat: 999,
	},
}}
