package edifact

// IFTSAI D.20A message specification: the structural (segment/group
// presence, order, cardinality) branching diagram for this message,
// UN/EDIFACT directory release D.20A. Scope is structural only -- not
// element-level content -- per edifact-ls-3uzr's epic non-goal.
// 17 segment groups, max nesting depth 4.
//
// Source: https://service.unece.org/trade/untdid/d20a/trmd/iftsai_c.htm
// section 4.3.1 "Segment table". That URL currently returns HTTP 403 from
// Cloudflare when fetched directly; this data was transcribed from the
// Wayback Machine's archived copy instead:
// http://web.archive.org/web/20231201231420/https://service.unece.org/trade/untdid/d20a/trmd/iftsai_c.htm
// -- re-check the direct URL first if cross-checking later, in case the
// block has lifted.
//
// Transcription note: generated in bulk (edifact-ls-13gu) the same way
// as every other message-type schema -- a script parsed the source's
// exact rail-art column positions mechanically and verified the result
// balances (all 17 groups open and close consistently)
// before being transcribed here, rather than reading the ASCII nesting
// by eye. Not individually spot-checked against the raw source the way
// the first several message types in this project were -- see
// edifact-ls-13gu for which representative samples were spot-checked
// for this batch.

func init() {
	RegisterSchema(
		MessageID{Type: "IFTSAI", Version: "D", Release: "20A", Agency: "UN"},
		iftsaiD20aSchema,
		"https://service.unece.org/trade/untdid/d20a/trmd/iftsai_c.htm",
	)
}

var iftsaiD20aSchema = Schema{Nodes: []SchemaNode{
	{Segment: "BGM", Mandatory: true, MaxRepeat: 1},   // Beginning of message
	{Segment: "DTM", Mandatory: false, MaxRepeat: 9},  // Date/time/period
	{Segment: "FTX", Mandatory: false, MaxRepeat: 99}, // Free text
	{Segment: "GEI", Mandatory: false, MaxRepeat: 1},  // Processing information
	{ // Segment group 1
		Group: []SchemaNode{
			{Segment: "RFF", Mandatory: true, MaxRepeat: 1},  // Reference
			{Segment: "DTM", Mandatory: false, MaxRepeat: 9}, // Date/time/period
		},
		Mandatory: false, MaxRepeat: 9,
	},
	{ // Segment group 2
		Group: []SchemaNode{
			{Segment: "LOC", Mandatory: true, MaxRepeat: 1},  // Place/location identification
			{Segment: "DTM", Mandatory: false, MaxRepeat: 9}, // Date/time/period
		},
		Mandatory: false, MaxRepeat: 9,
	},
	{ // Segment group 3
		Group: []SchemaNode{
			{Segment: "EQD", Mandatory: true, MaxRepeat: 1},  // Equipment details
			{Segment: "EQN", Mandatory: false, MaxRepeat: 9}, // Number of units
			{Segment: "MEA", Mandatory: false, MaxRepeat: 9}, // Measurements
			{Segment: "DIM", Mandatory: false, MaxRepeat: 9}, // Dimensions
			{Segment: "FTX", Mandatory: false, MaxRepeat: 9}, // Free text
			{Segment: "RFF", Mandatory: false, MaxRepeat: 9}, // Reference
			{Segment: "TPL", Mandatory: false, MaxRepeat: 1}, // Transport placement
			{Segment: "DTM", Mandatory: false, MaxRepeat: 9}, // Date/time/period
		},
		Mandatory: false, MaxRepeat: 999,
	},
	{ // Segment group 4
		Group: []SchemaNode{
			{Segment: "TDT", Mandatory: true, MaxRepeat: 1},   // Transport information
			{Segment: "DTM", Mandatory: false, MaxRepeat: 9},  // Date/time/period
			{Segment: "TSR", Mandatory: false, MaxRepeat: 9},  // Transport service requirements
			{Segment: "RFF", Mandatory: false, MaxRepeat: 9},  // Reference
			{Segment: "FTX", Mandatory: false, MaxRepeat: 9},  // Free text
			{Segment: "EQD", Mandatory: false, MaxRepeat: 99}, // Equipment details
			{Segment: "QTY", Mandatory: false, MaxRepeat: 9},  // Quantity
			{Segment: "MEA", Mandatory: false, MaxRepeat: 9},  // Measurements
			{ // Segment group 5
				Group: []SchemaNode{
					{Segment: "LOC", Mandatory: true, MaxRepeat: 1},  // Place/location identification
					{Segment: "RFF", Mandatory: false, MaxRepeat: 9}, // Reference
					{Segment: "FTX", Mandatory: false, MaxRepeat: 9}, // Free text
					{ // Segment group 6
						Group: []SchemaNode{
							{Segment: "DTM", Mandatory: true, MaxRepeat: 1},  // Date/time/period
							{Segment: "FTX", Mandatory: false, MaxRepeat: 9}, // Free text
						},
						Mandatory: false, MaxRepeat: 9,
					},
					{ // Segment group 7
						Group: []SchemaNode{
							{Segment: "EQD", Mandatory: true, MaxRepeat: 1},  // Equipment details
							{Segment: "EQN", Mandatory: false, MaxRepeat: 1}, // Number of units
						},
						Mandatory: false, MaxRepeat: 9,
					},
					{ // Segment group 8
						Group: []SchemaNode{
							{Segment: "TDT", Mandatory: true, MaxRepeat: 1}, // Transport information
							{ // Segment group 9
								Group: []SchemaNode{
									{Segment: "LOC", Mandatory: true, MaxRepeat: 1},    // Place/location identification
									{Segment: "DTM", Mandatory: false, MaxRepeat: 999}, // Date/time/period
								},
								Mandatory: false, MaxRepeat: 999,
							},
							{ // Segment group 10
								Group: []SchemaNode{
									{Segment: "EQD", Mandatory: true, MaxRepeat: 1},  // Equipment details
									{Segment: "EQN", Mandatory: false, MaxRepeat: 1}, // Number of units
								},
								Mandatory: false, MaxRepeat: 9,
							},
						},
						Mandatory: false, MaxRepeat: 999,
					},
					{Segment: "CNT", Mandatory: true, MaxRepeat: 1}, // Control total
				},
				Mandatory: false, MaxRepeat: 999,
			},
		},
		Mandatory: false, MaxRepeat: 999,
	},
	{ // Segment group 11
		Group: []SchemaNode{
			{Segment: "NAD", Mandatory: true, MaxRepeat: 1},  // Name and address
			{Segment: "LOC", Mandatory: false, MaxRepeat: 9}, // Place/location identification
			{ // Segment group 12
				Group: []SchemaNode{
					{Segment: "CTA", Mandatory: true, MaxRepeat: 1},  // Contact information
					{Segment: "COM", Mandatory: false, MaxRepeat: 9}, // Communication contact
				},
				Mandatory: false, MaxRepeat: 9,
			},
		},
		Mandatory: false, MaxRepeat: 9,
	},
	{ // Segment group 13
		Group: []SchemaNode{
			{Segment: "GID", Mandatory: true, MaxRepeat: 1},  // Goods item details
			{Segment: "HAN", Mandatory: false, MaxRepeat: 9}, // Handling instructions
			{Segment: "FTX", Mandatory: false, MaxRepeat: 9}, // Free text
			{ // Segment group 14
				Group: []SchemaNode{
					{Segment: "GDS", Mandatory: true, MaxRepeat: 1},  // Nature of cargo
					{Segment: "FTX", Mandatory: false, MaxRepeat: 9}, // Free text
				},
				Mandatory: false, MaxRepeat: 9,
			},
			{ // Segment group 15
				Group: []SchemaNode{
					{Segment: "MEA", Mandatory: true, MaxRepeat: 1},  // Measurements
					{Segment: "EQN", Mandatory: false, MaxRepeat: 9}, // Number of units
				},
				Mandatory: false, MaxRepeat: 9,
			},
			{ // Segment group 16
				Group: []SchemaNode{
					{Segment: "DIM", Mandatory: true, MaxRepeat: 1},  // Dimensions
					{Segment: "EQN", Mandatory: false, MaxRepeat: 9}, // Number of units
				},
				Mandatory: false, MaxRepeat: 9,
			},
			{ // Segment group 17
				Group: []SchemaNode{
					{Segment: "DGS", Mandatory: true, MaxRepeat: 1},  // Dangerous goods
					{Segment: "FTX", Mandatory: false, MaxRepeat: 9}, // Free text
					{Segment: "DTM", Mandatory: false, MaxRepeat: 9}, // Date/time/period
				},
				Mandatory: false, MaxRepeat: 9,
			},
		},
		Mandatory: false, MaxRepeat: 9,
	},
}}
