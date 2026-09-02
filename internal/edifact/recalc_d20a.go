package edifact

// RECALC D.20A message specification: the structural (segment/group
// presence, order, cardinality) branching diagram for this message,
// UN/EDIFACT directory release D.20A. Scope is structural only -- not
// element-level content -- per edifact-ls-3uzr's epic non-goal.
// 6 segment groups, max nesting depth 4.
//
// Source: https://service.unece.org/trade/untdid/d20a/trmd/recalc_c.htm
// section 4.3.1 "Segment table". That URL currently returns HTTP 403 from
// Cloudflare when fetched directly; this data was transcribed from the
// Wayback Machine's archived copy instead:
// http://web.archive.org/web/20231201142139/https://service.unece.org/trade/untdid/d20a/trmd/recalc_c.htm
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
		MessageID{Type: "RECALC", Version: "D", Release: "20A", Agency: "UN"},
		recalcD20aSchema,
		"https://service.unece.org/trade/untdid/d20a/trmd/recalc_c.htm",
	)
}

var recalcD20aSchema = Schema{Nodes: []SchemaNode{
	{Segment: "GEI", Mandatory: true, MaxRepeat: 5}, // Processing information
	{ // Segment group 1
		Group: []SchemaNode{
			{Segment: "NAD", Mandatory: true, MaxRepeat: 1},  // Name and address
			{Segment: "CTA", Mandatory: false, MaxRepeat: 1}, // Contact information
			{Segment: "COM", Mandatory: false, MaxRepeat: 5}, // Communication contact
			{Segment: "RFF", Mandatory: false, MaxRepeat: 9}, // Reference
		},
		Mandatory: true, MaxRepeat: 9,
	},
	{Segment: "DTM", Mandatory: true, MaxRepeat: 9},  // Date/time/period
	{Segment: "FTX", Mandatory: false, MaxRepeat: 3}, // Free text
	{ // Segment group 2
		Group: []SchemaNode{
			{Segment: "REL", Mandatory: true, MaxRepeat: 1},  // Relationship
			{Segment: "RFF", Mandatory: false, MaxRepeat: 9}, // Reference
			{Segment: "GEI", Mandatory: false, MaxRepeat: 5}, // Processing information
			{Segment: "CUX", Mandatory: true, MaxRepeat: 1},  // Currencies
			{Segment: "DTM", Mandatory: false, MaxRepeat: 9}, // Date/time/period
			{Segment: "QTY", Mandatory: false, MaxRepeat: 9}, // Quantity
			{Segment: "LOC", Mandatory: false, MaxRepeat: 1}, // Place/location identification
			{Segment: "PCD", Mandatory: false, MaxRepeat: 9}, // Percentage details
			{Segment: "FTX", Mandatory: false, MaxRepeat: 3}, // Free text
			{ // Segment group 3
				Group: []SchemaNode{
					{Segment: "ARD", Mandatory: true, MaxRepeat: 1}, // Monetary amount function
					{ // Segment group 4
						Group: []SchemaNode{
							{Segment: "MOA", Mandatory: true, MaxRepeat: 1},  // Monetary amount
							{Segment: "GEI", Mandatory: false, MaxRepeat: 9}, // Processing information
							{Segment: "PCD", Mandatory: false, MaxRepeat: 9}, // Percentage details
							{Segment: "FTX", Mandatory: false, MaxRepeat: 2}, // Free text
							{Segment: "DTM", Mandatory: false, MaxRepeat: 9}, // Date/time/period
							{ // Segment group 5
								Group: []SchemaNode{
									{Segment: "RFF", Mandatory: true, MaxRepeat: 1},  // Reference
									{Segment: "DTM", Mandatory: false, MaxRepeat: 2}, // Date/time/period
									{Segment: "COM", Mandatory: false, MaxRepeat: 1}, // Communication contact
									{Segment: "FTX", Mandatory: false, MaxRepeat: 1}, // Free text
								},
								Mandatory: false, MaxRepeat: 9,
							},
							{ // Segment group 6
								Group: []SchemaNode{
									{Segment: "CUX", Mandatory: true, MaxRepeat: 1},  // Currencies
									{Segment: "MOA", Mandatory: false, MaxRepeat: 1}, // Monetary amount
									{Segment: "AJT", Mandatory: true, MaxRepeat: 1},  // Adjustment details
								},
								Mandatory: false, MaxRepeat: 1,
							},
						},
						Mandatory: true, MaxRepeat: 99,
					},
				},
				Mandatory: true, MaxRepeat: 9,
			},
		},
		Mandatory: true, MaxRepeat: 99,
	},
}}
