package edifact

// IFTSTQ D.20A message specification: the structural (segment/group
// presence, order, cardinality) branching diagram for this message,
// UN/EDIFACT directory release D.20A. Scope is structural only -- not
// element-level content -- per edifact-ls-3uzr's epic non-goal.
// 8 segment groups, max nesting depth 3.
//
// Source: https://service.unece.org/trade/untdid/d20a/trmd/iftstq_c.htm
// section 4.3.1 "Segment table". That URL currently returns HTTP 403 from
// Cloudflare when fetched directly; this data was transcribed from the
// Wayback Machine's archived copy instead:
// http://web.archive.org/web/20240416113215/https://service.unece.org/trade/untdid/d20a/trmd/iftstq_c.htm
// -- re-check the direct URL first if cross-checking later, in case the
// block has lifted.
//
// Transcription note: generated in bulk (edifact-ls-13gu) the same way
// as every other message-type schema -- a script parsed the source's
// exact rail-art column positions mechanically and verified the result
// balances (all 8 groups open and close consistently)
// before being transcribed here, rather than reading the ASCII nesting
// by eye. Not individually spot-checked against the raw source the way
// the first several message types in this project were -- see
// edifact-ls-13gu for which representative samples were spot-checked
// for this batch.

func init() {
	RegisterSchema(
		MessageID{Type: "IFTSTQ", Version: "D", Release: "20A", Agency: "UN"},
		iftstqD20aSchema,
		"https://service.unece.org/trade/untdid/d20a/trmd/iftstq_c.htm",
	)
}

var iftstqD20aSchema = Schema{Nodes: []SchemaNode{
	{Segment: "BGM", Mandatory: true, MaxRepeat: 1},   // Beginning of message
	{Segment: "DTM", Mandatory: false, MaxRepeat: 9},  // Date/time/period
	{Segment: "LOC", Mandatory: false, MaxRepeat: 99}, // Place/location identification
	{ // Segment group 1
		Group: []SchemaNode{
			{Segment: "RFF", Mandatory: true, MaxRepeat: 1},  // Reference
			{Segment: "DTM", Mandatory: false, MaxRepeat: 9}, // Date/time/period
		},
		Mandatory: false, MaxRepeat: 999,
	},
	{Segment: "TDT", Mandatory: false, MaxRepeat: 99},  // Transport information
	{Segment: "EQD", Mandatory: false, MaxRepeat: 999}, // Equipment details
	{ // Segment group 2
		Group: []SchemaNode{
			{Segment: "NAD", Mandatory: true, MaxRepeat: 1}, // Name and address
			{ // Segment group 3
				Group: []SchemaNode{
					{Segment: "CTA", Mandatory: true, MaxRepeat: 1},  // Contact information
					{Segment: "COM", Mandatory: false, MaxRepeat: 9}, // Communication contact
				},
				Mandatory: false, MaxRepeat: 9,
			},
		},
		Mandatory: false, MaxRepeat: 9,
	},
	{ // Segment group 4
		Group: []SchemaNode{
			{Segment: "CNI", Mandatory: true, MaxRepeat: 1},   // Consignment information
			{Segment: "DTM", Mandatory: false, MaxRepeat: 9},  // Date/time/period
			{Segment: "LOC", Mandatory: false, MaxRepeat: 99}, // Place/location identification
			{Segment: "RFF", Mandatory: false, MaxRepeat: 99}, // Reference
			{Segment: "TDT", Mandatory: false, MaxRepeat: 99}, // Transport information
			{ // Segment group 5
				Group: []SchemaNode{
					{Segment: "EQD", Mandatory: true, MaxRepeat: 1},  // Equipment details
					{Segment: "TPL", Mandatory: false, MaxRepeat: 9}, // Transport placement
				},
				Mandatory: false, MaxRepeat: 99,
			},
			{ // Segment group 6
				Group: []SchemaNode{
					{Segment: "NAD", Mandatory: true, MaxRepeat: 1}, // Name and address
					{ // Segment group 7
						Group: []SchemaNode{
							{Segment: "GID", Mandatory: true, MaxRepeat: 1},   // Goods item details
							{Segment: "FTX", Mandatory: false, MaxRepeat: 1},  // Free text
							{Segment: "SGP", Mandatory: false, MaxRepeat: 99}, // Split goods placement
						},
						Mandatory: false, MaxRepeat: 99,
					},
					{ // Segment group 8
						Group: []SchemaNode{
							{Segment: "CTA", Mandatory: true, MaxRepeat: 1},  // Contact information
							{Segment: "COM", Mandatory: false, MaxRepeat: 9}, // Communication contact
						},
						Mandatory: false, MaxRepeat: 9,
					},
				},
				Mandatory: false, MaxRepeat: 9,
			},
		},
		Mandatory: false, MaxRepeat: 999,
	},
}}
