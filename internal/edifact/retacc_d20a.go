package edifact

// RETACC D.20A message specification: the structural (segment/group
// presence, order, cardinality) branching diagram for this message,
// UN/EDIFACT directory release D.20A. Scope is structural only -- not
// element-level content -- per edifact-ls-3uzr's epic non-goal.
// 6 segment groups, max nesting depth 3.
//
// Source: https://service.unece.org/trade/untdid/d20a/trmd/retacc_c.htm
// section 4.3.1 "Segment table". That URL currently returns HTTP 403 from
// Cloudflare when fetched directly; this data was transcribed from the
// Wayback Machine's archived copy instead:
// http://web.archive.org/web/20231208115345/https://service.unece.org/trade/untdid/d20a/trmd/retacc_c.htm
// -- re-check the direct URL first if cross-checking later, in case the
// block has lifted.
//
// Transcription note: generated in bulk (edifact-ls-13gu) the same way
// as every other message-type schema -- a script parsed the source's
// exact rail-art column positions mechanically and verified the result
// balances (all 6 groups open and close consistently)
// before being transcribed here, rather than reading the ASCII nesting
// by eye. Not individually spot-checked against the raw source the way
// the first several message types in this project were -- see
// edifact-ls-13gu for which representative samples were spot-checked
// for this batch.

func init() {
	RegisterSchema(
		MessageID{Type: "RETACC", Version: "D", Release: "20A", Agency: "UN"},
		retaccD20aSchema,
		"https://service.unece.org/trade/untdid/d20a/trmd/retacc_c.htm",
	)
}

var retaccD20aSchema = Schema{Nodes: []SchemaNode{
	{Segment: "GEI", Mandatory: true, MaxRepeat: 6}, // Processing information
	{ // Segment group 1
		Group: []SchemaNode{
			{Segment: "NAD", Mandatory: true, MaxRepeat: 1},  // Name and address
			{Segment: "CTA", Mandatory: false, MaxRepeat: 1}, // Contact information
			{Segment: "COM", Mandatory: false, MaxRepeat: 5}, // Communication contact
			{Segment: "RFF", Mandatory: false, MaxRepeat: 9}, // Reference
		},
		Mandatory: true, MaxRepeat: 9,
	},
	{Segment: "DTM", Mandatory: true, MaxRepeat: 4},  // Date/time/period
	{Segment: "FTX", Mandatory: false, MaxRepeat: 5}, // Free text
	{ // Segment group 2
		Group: []SchemaNode{
			{Segment: "CUX", Mandatory: true, MaxRepeat: 1},  // Currencies
			{Segment: "RFF", Mandatory: true, MaxRepeat: 5},  // Reference
			{Segment: "FTX", Mandatory: false, MaxRepeat: 1}, // Free text
			{ // Segment group 3
				Group: []SchemaNode{
					{Segment: "MOA", Mandatory: true, MaxRepeat: 1},   // Monetary amount
					{Segment: "DTM", Mandatory: true, MaxRepeat: 10},  // Date/time/period
					{Segment: "PCD", Mandatory: false, MaxRepeat: 2},  // Percentage details
					{Segment: "GEI", Mandatory: false, MaxRepeat: 10}, // Processing information
					{Segment: "LOC", Mandatory: false, MaxRepeat: 9},  // Place/location identification
					{Segment: "FTX", Mandatory: false, MaxRepeat: 3},  // Free text
					{ // Segment group 4
						Group: []SchemaNode{
							{Segment: "RFF", Mandatory: true, MaxRepeat: 1},  // Reference
							{Segment: "DTM", Mandatory: false, MaxRepeat: 1}, // Date/time/period
							{Segment: "COM", Mandatory: false, MaxRepeat: 1}, // Communication contact
							{Segment: "FTX", Mandatory: false, MaxRepeat: 1}, // Free text
						},
						Mandatory: false, MaxRepeat: 9,
					},
				},
				Mandatory: true, MaxRepeat: 999,
			},
		},
		Mandatory: true, MaxRepeat: 99,
	},
	{Segment: "UNS", Mandatory: true, MaxRepeat: 1}, // Section control
	{ // Segment group 5
		Group: []SchemaNode{
			{Segment: "MOA", Mandatory: true, MaxRepeat: 1},  // Monetary amount
			{Segment: "GEI", Mandatory: false, MaxRepeat: 1}, // Processing information
			{ // Segment group 6
				Group: []SchemaNode{
					{Segment: "NAD", Mandatory: true, MaxRepeat: 1},  // Name and address
					{Segment: "PCD", Mandatory: false, MaxRepeat: 1}, // Percentage details
					{Segment: "RFF", Mandatory: false, MaxRepeat: 2}, // Reference
					{Segment: "GEI", Mandatory: false, MaxRepeat: 1}, // Processing information
					{Segment: "MOA", Mandatory: false, MaxRepeat: 1}, // Monetary amount
				},
				Mandatory: false, MaxRepeat: 999,
			},
			{Segment: "PAI", Mandatory: true, MaxRepeat: 1},  // Payment instructions
			{Segment: "RFF", Mandatory: false, MaxRepeat: 9}, // Reference
			{Segment: "DTM", Mandatory: false, MaxRepeat: 3}, // Date/time/period
		},
		Mandatory: false, MaxRepeat: 99,
	},
}}
