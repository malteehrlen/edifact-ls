package edifact

// BMISRM D.20A message specification: the structural (segment/group
// presence, order, cardinality) branching diagram for this message,
// UN/EDIFACT directory release D.20A. Scope is structural only -- not
// element-level content -- per edifact-ls-3uzr's epic non-goal.
// 9 segment groups, max nesting depth 4.
//
// Source: https://service.unece.org/trade/untdid/d20a/trmd/bmisrm_c.htm
// section 4.3.1 "Segment table". That URL currently returns HTTP 403 from
// Cloudflare when fetched directly; this data was transcribed from the
// Wayback Machine's archived copy instead:
// http://web.archive.org/web/20231202081110/https://service.unece.org/trade/untdid/d20a/trmd/bmisrm_c.htm
// -- re-check the direct URL first if cross-checking later, in case the
// block has lifted.
//
// Transcription note: generated in bulk (edifact-ls-13gu) the same way
// as every other message-type schema -- a script parsed the source's
// exact rail-art column positions mechanically and verified the result
// balances (all 9 groups open and close consistently)
// before being transcribed here, rather than reading the ASCII nesting
// by eye. Not individually spot-checked against the raw source the way
// the first several message types in this project were -- see
// edifact-ls-13gu for which representative samples were spot-checked
// for this batch.

func init() {
	RegisterSchema(
		MessageID{Type: "BMISRM", Version: "D", Release: "20A", Agency: "UN"},
		bmisrmD20aSchema,
		"https://service.unece.org/trade/untdid/d20a/trmd/bmisrm_c.htm",
	)
}

var bmisrmD20aSchema = Schema{Nodes: []SchemaNode{
	{Segment: "BGM", Mandatory: true, MaxRepeat: 1},   // Beginning of message
	{Segment: "DTM", Mandatory: true, MaxRepeat: 1},   // Date/time/period
	{Segment: "RFF", Mandatory: true, MaxRepeat: 9},   // Reference
	{Segment: "LOC", Mandatory: true, MaxRepeat: 9},   // Place/location identification
	{Segment: "FTX", Mandatory: false, MaxRepeat: 99}, // Free text
	{ // Segment group 1
		Group: []SchemaNode{
			{Segment: "PNA", Mandatory: true, MaxRepeat: 1},  // Party identification
			{Segment: "ADR", Mandatory: false, MaxRepeat: 1}, // Address
			{Segment: "RFF", Mandatory: false, MaxRepeat: 1}, // Reference
			{ // Segment group 2
				Group: []SchemaNode{
					{Segment: "CTA", Mandatory: true, MaxRepeat: 1},  // Contact information
					{Segment: "COM", Mandatory: false, MaxRepeat: 9}, // Communication contact
				},
				Mandatory: false, MaxRepeat: 9,
			},
		},
		Mandatory: true, MaxRepeat: 9,
	},
	{ // Segment group 3
		Group: []SchemaNode{
			{Segment: "DTM", Mandatory: true, MaxRepeat: 1},  // Date/time/period
			{Segment: "STS", Mandatory: false, MaxRepeat: 1}, // Status
			{Segment: "FTX", Mandatory: false, MaxRepeat: 9}, // Free text
		},
		Mandatory: false, MaxRepeat: 999,
	},
	{ // Segment group 4
		Group: []SchemaNode{
			{Segment: "LIN", Mandatory: true, MaxRepeat: 1},   // Line item
			{Segment: "QTY", Mandatory: false, MaxRepeat: 1},  // Quantity
			{Segment: "FTX", Mandatory: false, MaxRepeat: 99}, // Free text
			{ // Segment group 5
				Group: []SchemaNode{
					{Segment: "PNA", Mandatory: true, MaxRepeat: 1},  // Party identification
					{Segment: "ADR", Mandatory: false, MaxRepeat: 1}, // Address
					{Segment: "RFF", Mandatory: false, MaxRepeat: 1}, // Reference
					{ // Segment group 6
						Group: []SchemaNode{
							{Segment: "CTA", Mandatory: true, MaxRepeat: 1},  // Contact information
							{Segment: "COM", Mandatory: false, MaxRepeat: 9}, // Communication contact
						},
						Mandatory: false, MaxRepeat: 9,
					},
				},
				Mandatory: false, MaxRepeat: 9,
			},
			{ // Segment group 7
				Group: []SchemaNode{
					{Segment: "TDT", Mandatory: true, MaxRepeat: 1},   // Transport information
					{Segment: "EQD", Mandatory: false, MaxRepeat: 1},  // Equipment details
					{Segment: "RFF", Mandatory: false, MaxRepeat: 9},  // Reference
					{Segment: "QTY", Mandatory: false, MaxRepeat: 1},  // Quantity
					{Segment: "MEA", Mandatory: false, MaxRepeat: 99}, // Measurements
					{Segment: "PCD", Mandatory: false, MaxRepeat: 99}, // Percentage details
					{ // Segment group 8
						Group: []SchemaNode{
							{Segment: "PSD", Mandatory: true, MaxRepeat: 1},   // Physical sample description
							{Segment: "LOC", Mandatory: false, MaxRepeat: 99}, // Place/location identification
							{ // Segment group 9
								Group: []SchemaNode{
									{Segment: "TEM", Mandatory: true, MaxRepeat: 1},   // Test method
									{Segment: "MEA", Mandatory: false, MaxRepeat: 99}, // Measurements
									{Segment: "PCD", Mandatory: false, MaxRepeat: 99}, // Percentage details
									{Segment: "DTM", Mandatory: false, MaxRepeat: 99}, // Date/time/period
									{Segment: "FTX", Mandatory: false, MaxRepeat: 99}, // Free text
								},
								Mandatory: false, MaxRepeat: 999,
							},
						},
						Mandatory: false, MaxRepeat: 99,
					},
				},
				Mandatory: false, MaxRepeat: 999,
			},
		},
		Mandatory: true, MaxRepeat: 999,
	},
}}
