package edifact

// RECLAM D.20A message specification: the structural (segment/group
// presence, order, cardinality) branching diagram for this message,
// UN/EDIFACT directory release D.20A. Scope is structural only -- not
// element-level content -- per edifact-ls-3uzr's epic non-goal.
// 10 segment groups, max nesting depth 3.
//
// Source: https://service.unece.org/trade/untdid/d20a/trmd/reclam_c.htm
// section 4.3.1 "Segment table". That URL currently returns HTTP 403 from
// Cloudflare when fetched directly; this data was transcribed from the
// Wayback Machine's archived copy instead:
// http://web.archive.org/web/20231208124908/https://service.unece.org/trade/untdid/d20a/trmd/reclam_c.htm
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
		MessageID{Type: "RECLAM", Version: "D", Release: "20A", Agency: "UN"},
		reclamD20aSchema,
		"https://service.unece.org/trade/untdid/d20a/trmd/reclam_c.htm",
	)
}

var reclamD20aSchema = Schema{Nodes: []SchemaNode{
	{ // Segment group 1
		Group: []SchemaNode{
			{Segment: "GEI", Mandatory: true, MaxRepeat: 1},  // Processing information
			{Segment: "FTX", Mandatory: false, MaxRepeat: 1}, // Free text
		},
		Mandatory: true, MaxRepeat: 99,
	},
	{ // Segment group 2
		Group: []SchemaNode{
			{Segment: "NAD", Mandatory: true, MaxRepeat: 1},  // Name and address
			{Segment: "PCD", Mandatory: false, MaxRepeat: 1}, // Percentage details
			{Segment: "CTA", Mandatory: false, MaxRepeat: 1}, // Contact information
			{Segment: "COM", Mandatory: false, MaxRepeat: 5}, // Communication contact
			{Segment: "RFF", Mandatory: false, MaxRepeat: 5}, // Reference
			{Segment: "GEI", Mandatory: false, MaxRepeat: 2}, // Processing information
			{Segment: "FTX", Mandatory: false, MaxRepeat: 2}, // Free text
			{Segment: "DTM", Mandatory: false, MaxRepeat: 9}, // Date/time/period
		},
		Mandatory: true, MaxRepeat: 999,
	},
	{Segment: "UNS", Mandatory: true, MaxRepeat: 1},    // Section control
	{Segment: "DTM", Mandatory: true, MaxRepeat: 8},    // Date/time/period
	{Segment: "FTX", Mandatory: false, MaxRepeat: 999}, // Free text
	{ // Segment group 3
		Group: []SchemaNode{
			{Segment: "RFF", Mandatory: true, MaxRepeat: 1},  // Reference
			{Segment: "DTM", Mandatory: false, MaxRepeat: 1}, // Date/time/period
			{Segment: "COM", Mandatory: false, MaxRepeat: 1}, // Communication contact
			{Segment: "FTX", Mandatory: false, MaxRepeat: 1}, // Free text
		},
		Mandatory: false, MaxRepeat: 9,
	},
	{Segment: "LOC", Mandatory: false, MaxRepeat: 1}, // Place/location identification
	{Segment: "MOA", Mandatory: false, MaxRepeat: 1}, // Monetary amount
	{Segment: "PCD", Mandatory: false, MaxRepeat: 1}, // Percentage details
	{ // Segment group 4
		Group: []SchemaNode{
			{Segment: "GEI", Mandatory: true, MaxRepeat: 1},  // Processing information
			{Segment: "RFF", Mandatory: false, MaxRepeat: 8}, // Reference
			{Segment: "LOC", Mandatory: false, MaxRepeat: 9}, // Place/location identification
			{Segment: "DTM", Mandatory: false, MaxRepeat: 1}, // Date/time/period
			{Segment: "FTX", Mandatory: false, MaxRepeat: 3}, // Free text
			{Segment: "PCD", Mandatory: false, MaxRepeat: 2}, // Percentage details
			{ // Segment group 5
				Group: []SchemaNode{
					{Segment: "CUX", Mandatory: true, MaxRepeat: 1},  // Currencies
					{Segment: "GEI", Mandatory: false, MaxRepeat: 1}, // Processing information
					{Segment: "FTX", Mandatory: false, MaxRepeat: 1}, // Free text
					{ // Segment group 6
						Group: []SchemaNode{
							{Segment: "MOA", Mandatory: true, MaxRepeat: 1},  // Monetary amount
							{Segment: "GEI", Mandatory: true, MaxRepeat: 3},  // Processing information
							{Segment: "CUX", Mandatory: false, MaxRepeat: 1}, // Currencies
							{Segment: "FTX", Mandatory: false, MaxRepeat: 2}, // Free text
						},
						Mandatory: true, MaxRepeat: 99,
					},
					{Segment: "PCD", Mandatory: true, MaxRepeat: 1}, // Percentage details
				},
				Mandatory: false, MaxRepeat: 3,
			},
			{ // Segment group 7
				Group: []SchemaNode{
					{Segment: "NAD", Mandatory: true, MaxRepeat: 1},  // Name and address
					{Segment: "PCD", Mandatory: false, MaxRepeat: 1}, // Percentage details
				},
				Mandatory: false, MaxRepeat: 99,
			},
		},
		Mandatory: false, MaxRepeat: 99,
	},
	{Segment: "UNS", Mandatory: true, MaxRepeat: 1}, // Section control
	{ // Segment group 8
		Group: []SchemaNode{
			{Segment: "CUX", Mandatory: true, MaxRepeat: 1}, // Currencies
			{ // Segment group 9
				Group: []SchemaNode{
					{Segment: "MOA", Mandatory: true, MaxRepeat: 1},   // Monetary amount
					{Segment: "PCD", Mandatory: false, MaxRepeat: 2},  // Percentage details
					{Segment: "CUX", Mandatory: false, MaxRepeat: 1},  // Currencies
					{Segment: "FTX", Mandatory: false, MaxRepeat: 10}, // Free text
					{ // Segment group 10
						Group: []SchemaNode{
							{Segment: "RFF", Mandatory: true, MaxRepeat: 1},  // Reference
							{Segment: "DTM", Mandatory: false, MaxRepeat: 1}, // Date/time/period
							{Segment: "COM", Mandatory: false, MaxRepeat: 1}, // Communication contact
							{Segment: "FTX", Mandatory: false, MaxRepeat: 1}, // Free text
						},
						Mandatory: false, MaxRepeat: 99,
					},
					{Segment: "GEI", Mandatory: true, MaxRepeat: 99}, // Processing information
				},
				Mandatory: true, MaxRepeat: 999,
			},
		},
		Mandatory: false, MaxRepeat: 3,
	},
}}
