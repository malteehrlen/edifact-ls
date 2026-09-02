package edifact

// BERMAN D.20A message specification: the structural (segment/group
// presence, order, cardinality) branching diagram for this message,
// UN/EDIFACT directory release D.20A. Scope is structural only -- not
// element-level content -- per edifact-ls-3uzr's epic non-goal.
// 10 segment groups, max nesting depth 4.
//
// Source: https://service.unece.org/trade/untdid/d20a/trmd/berman_c.htm
// section 4.3.1 "Segment table". That URL currently returns HTTP 403 from
// Cloudflare when fetched directly; this data was transcribed from the
// Wayback Machine's archived copy instead:
// http://web.archive.org/web/20231201234855/https://service.unece.org/trade/untdid/d20a/trmd/berman_c.htm
// -- re-check the direct URL first if cross-checking later, in case the
// block has lifted.
//
// Transcription note: generated in bulk (edifact-ls-13gu) the same way
// as every other message-type schema -- a script parsed the source's
// exact rail-art column positions mechanically and verified the result
// balances (all 10 groups open and close consistently)
// before being transcribed here, rather than reading the ASCII nesting
// by eye. Not individually spot-checked against the raw source the way
// the first several message types in this project were -- see
// edifact-ls-13gu for which representative samples were spot-checked
// for this batch.

func init() {
	RegisterSchema(
		MessageID{Type: "BERMAN", Version: "D", Release: "20A", Agency: "UN"},
		bermanD20aSchema,
		"https://service.unece.org/trade/untdid/d20a/trmd/berman_c.htm",
	)
}

var bermanD20aSchema = Schema{Nodes: []SchemaNode{
	{Segment: "BGM", Mandatory: true, MaxRepeat: 1},  // Beginning of message
	{Segment: "DTM", Mandatory: false, MaxRepeat: 9}, // Date/time/period
	{Segment: "FTX", Mandatory: false, MaxRepeat: 9}, // Free text
	{Segment: "RFF", Mandatory: false, MaxRepeat: 9}, // Reference
	{Segment: "QTY", Mandatory: false, MaxRepeat: 9}, // Quantity
	{ // Segment group 1
		Group: []SchemaNode{
			{Segment: "NAD", Mandatory: true, MaxRepeat: 1}, // Name and address
			{ // Segment group 2
				Group: []SchemaNode{
					{Segment: "CTA", Mandatory: true, MaxRepeat: 1},  // Contact information
					{Segment: "COM", Mandatory: false, MaxRepeat: 9}, // Communication contact
				},
				Mandatory: false, MaxRepeat: 9,
			},
		},
		Mandatory: true, MaxRepeat: 99,
	},
	{ // Segment group 3
		Group: []SchemaNode{
			{Segment: "TDT", Mandatory: true, MaxRepeat: 1},   // Transport information
			{Segment: "RFF", Mandatory: false, MaxRepeat: 9},  // Reference
			{Segment: "DTM", Mandatory: false, MaxRepeat: 9},  // Date/time/period
			{Segment: "MEA", Mandatory: false, MaxRepeat: 99}, // Measurements
			{Segment: "FTX", Mandatory: false, MaxRepeat: 9},  // Free text
			{Segment: "COM", Mandatory: false, MaxRepeat: 9},  // Communication contact
			{ // Segment group 4
				Group: []SchemaNode{
					{Segment: "LOC", Mandatory: true, MaxRepeat: 1},  // Place/location identification
					{Segment: "DTM", Mandatory: false, MaxRepeat: 9}, // Date/time/period
				},
				Mandatory: false, MaxRepeat: 999,
			},
			{ // Segment group 5
				Group: []SchemaNode{
					{Segment: "GOR", Mandatory: true, MaxRepeat: 1},   // Governmental requirements
					{Segment: "RFF", Mandatory: false, MaxRepeat: 9},  // Reference
					{Segment: "NAD", Mandatory: false, MaxRepeat: 99}, // Name and address
					{ // Segment group 6
						Group: []SchemaNode{
							{Segment: "DOC", Mandatory: true, MaxRepeat: 1},  // Document/message details
							{Segment: "DTM", Mandatory: false, MaxRepeat: 9}, // Date/time/period
							{Segment: "LOC", Mandatory: false, MaxRepeat: 9}, // Place/location identification
						},
						Mandatory: false, MaxRepeat: 9,
					},
				},
				Mandatory: false, MaxRepeat: 9,
			},
		},
		Mandatory: true, MaxRepeat: 99,
	},
	{ // Segment group 7
		Group: []SchemaNode{
			{Segment: "TSR", Mandatory: true, MaxRepeat: 1},  // Transport service requirements
			{Segment: "QTY", Mandatory: false, MaxRepeat: 9}, // Quantity
			{Segment: "FTX", Mandatory: false, MaxRepeat: 9}, // Free text
			{ // Segment group 8
				Group: []SchemaNode{
					{Segment: "LOC", Mandatory: true, MaxRepeat: 1},  // Place/location identification
					{Segment: "MEA", Mandatory: false, MaxRepeat: 9}, // Measurements
					{Segment: "DTM", Mandatory: false, MaxRepeat: 9}, // Date/time/period
					{Segment: "QTY", Mandatory: false, MaxRepeat: 9}, // Quantity
					{Segment: "POC", Mandatory: false, MaxRepeat: 9}, // Purpose of conveyance call
					{Segment: "FTX", Mandatory: false, MaxRepeat: 9}, // Free text
					{ // Segment group 9
						Group: []SchemaNode{
							{Segment: "HAN", Mandatory: true, MaxRepeat: 1},  // Handling instructions
							{Segment: "NAD", Mandatory: false, MaxRepeat: 9}, // Name and address
							{ // Segment group 10
								Group: []SchemaNode{
									{Segment: "GDS", Mandatory: true, MaxRepeat: 1},  // Nature of cargo
									{Segment: "FTX", Mandatory: false, MaxRepeat: 9}, // Free text
									{Segment: "MEA", Mandatory: false, MaxRepeat: 9}, // Measurements
									{Segment: "EQN", Mandatory: false, MaxRepeat: 9}, // Number of units
									{Segment: "DGS", Mandatory: false, MaxRepeat: 9}, // Dangerous goods
								},
								Mandatory: false, MaxRepeat: 9,
							},
						},
						Mandatory: false, MaxRepeat: 9,
					},
				},
				Mandatory: false, MaxRepeat: 9,
			},
		},
		Mandatory: false, MaxRepeat: 99,
	},
}}
