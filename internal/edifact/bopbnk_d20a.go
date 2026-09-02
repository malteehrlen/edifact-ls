package edifact

// BOPBNK D.20A message specification: the structural (segment/group
// presence, order, cardinality) branching diagram for this message,
// UN/EDIFACT directory release D.20A. Scope is structural only -- not
// element-level content -- per edifact-ls-3uzr's epic non-goal.
// 7 segment groups, max nesting depth 4.
//
// Source: https://service.unece.org/trade/untdid/d20a/trmd/bopbnk_c.htm
// section 4.3.1 "Segment table". That URL currently returns HTTP 403 from
// Cloudflare when fetched directly; this data was transcribed from the
// Wayback Machine's archived copy instead:
// http://web.archive.org/web/20231201140917/https://service.unece.org/trade/untdid/d20a/trmd/bopbnk_c.htm
// -- re-check the direct URL first if cross-checking later, in case the
// block has lifted.
//
// Transcription note: generated in bulk (edifact-ls-13gu) the same way
// as every other message-type schema -- a script parsed the source's
// exact rail-art column positions mechanically and verified the result
// balances (all 7 groups open and close consistently)
// before being transcribed here, rather than reading the ASCII nesting
// by eye. Not individually spot-checked against the raw source the way
// the first several message types in this project were -- see
// edifact-ls-13gu for which representative samples were spot-checked
// for this batch.

func init() {
	RegisterSchema(
		MessageID{Type: "BOPBNK", Version: "D", Release: "20A", Agency: "UN"},
		bopbnkD20aSchema,
		"https://service.unece.org/trade/untdid/d20a/trmd/bopbnk_c.htm",
	)
}

var bopbnkD20aSchema = Schema{Nodes: []SchemaNode{
	{Segment: "BGM", Mandatory: true, MaxRepeat: 1}, // Beginning of message
	{Segment: "DTM", Mandatory: true, MaxRepeat: 9}, // Date/time/period
	{ // Segment group 1
		Group: []SchemaNode{
			{Segment: "RFF", Mandatory: true, MaxRepeat: 1},  // Reference
			{Segment: "DTM", Mandatory: false, MaxRepeat: 1}, // Date/time/period
		},
		Mandatory: false, MaxRepeat: 9,
	},
	{ // Segment group 2
		Group: []SchemaNode{
			{Segment: "NAD", Mandatory: true, MaxRepeat: 1},   // Name and address
			{Segment: "CTA", Mandatory: false, MaxRepeat: 1},  // Contact information
			{Segment: "COM", Mandatory: false, MaxRepeat: 9},  // Communication contact
			{Segment: "FTX", Mandatory: false, MaxRepeat: 99}, // Free text
		},
		Mandatory: true, MaxRepeat: 9,
	},
	{ // Segment group 3
		Group: []SchemaNode{
			{Segment: "RFF", Mandatory: true, MaxRepeat: 1},  // Reference
			{Segment: "CUX", Mandatory: false, MaxRepeat: 1}, // Currencies
			{Segment: "MOA", Mandatory: false, MaxRepeat: 9}, // Monetary amount
			{Segment: "LOC", Mandatory: false, MaxRepeat: 1}, // Place/location identification
			{ // Segment group 4
				Group: []SchemaNode{
					{Segment: "RCS", Mandatory: true, MaxRepeat: 1},  // Requirements and conditions
					{Segment: "FTX", Mandatory: false, MaxRepeat: 1}, // Free text
					{ // Segment group 5
						Group: []SchemaNode{
							{Segment: "MOA", Mandatory: true, MaxRepeat: 1},  // Monetary amount
							{Segment: "ATT", Mandatory: false, MaxRepeat: 1}, // Attribute
							{Segment: "NAD", Mandatory: false, MaxRepeat: 1}, // Name and address
							{ // Segment group 6
								Group: []SchemaNode{
									{Segment: "GIR", Mandatory: true, MaxRepeat: 1},  // Related identification numbers
									{Segment: "QTY", Mandatory: false, MaxRepeat: 1}, // Quantity
									{Segment: "PRI", Mandatory: false, MaxRepeat: 1}, // Price details
								},
								Mandatory: false, MaxRepeat: 1,
							},
							{ // Segment group 7
								Group: []SchemaNode{
									{Segment: "RFF", Mandatory: true, MaxRepeat: 1},  // Reference
									{Segment: "DTM", Mandatory: false, MaxRepeat: 1}, // Date/time/period
								},
								Mandatory: false, MaxRepeat: 1,
							},
							{Segment: "LOC", Mandatory: true, MaxRepeat: 9}, // Place/location identification
						},
						Mandatory: true, MaxRepeat: 9999,
					},
				},
				Mandatory: true, MaxRepeat: 999,
			},
		},
		Mandatory: true, MaxRepeat: 999,
	},
	{Segment: "CNT", Mandatory: false, MaxRepeat: 9}, // Control total
}}
