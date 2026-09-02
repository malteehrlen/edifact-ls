package edifact

// INSPRE D.20A message specification: the structural (segment/group
// presence, order, cardinality) branching diagram for this message,
// UN/EDIFACT directory release D.20A. Scope is structural only -- not
// element-level content -- per edifact-ls-3uzr's epic non-goal.
// 9 segment groups, max nesting depth 4.
//
// Source: https://service.unece.org/trade/untdid/d20a/trmd/inspre_c.htm
// section 4.3.1 "Segment table". That URL currently returns HTTP 403 from
// Cloudflare when fetched directly; this data was transcribed from the
// Wayback Machine's archived copy instead:
// http://web.archive.org/web/20231208133715/https://service.unece.org/trade/untdid/d20a/trmd/inspre_c.htm
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
		MessageID{Type: "INSPRE", Version: "D", Release: "20A", Agency: "UN"},
		inspreD20aSchema,
		"https://service.unece.org/trade/untdid/d20a/trmd/inspre_c.htm",
	)
}

var inspreD20aSchema = Schema{Nodes: []SchemaNode{
	{Segment: "BGM", Mandatory: false, MaxRepeat: 1}, // Beginning of message
	{Segment: "RFF", Mandatory: false, MaxRepeat: 9}, // Reference
	{Segment: "DTM", Mandatory: true, MaxRepeat: 9},  // Date/time/period
	{ // Segment group 1
		Group: []SchemaNode{
			{Segment: "ATT", Mandatory: true, MaxRepeat: 1},  // Attribute
			{Segment: "PCD", Mandatory: false, MaxRepeat: 1}, // Percentage details
		},
		Mandatory: true, MaxRepeat: 9,
	},
	{Segment: "ICD", Mandatory: false, MaxRepeat: 1}, // Insurance cover description
	{ // Segment group 2
		Group: []SchemaNode{
			{Segment: "NAD", Mandatory: true, MaxRepeat: 1},  // Name and address
			{Segment: "ATT", Mandatory: false, MaxRepeat: 1}, // Attribute
			{Segment: "CTA", Mandatory: false, MaxRepeat: 1}, // Contact information
			{Segment: "COM", Mandatory: false, MaxRepeat: 3}, // Communication contact
			{Segment: "RFF", Mandatory: false, MaxRepeat: 9}, // Reference
		},
		Mandatory: true, MaxRepeat: 9,
	},
	{ // Segment group 3
		Group: []SchemaNode{
			{Segment: "DTM", Mandatory: true, MaxRepeat: 1}, // Date/time/period
			{ // Segment group 4
				Group: []SchemaNode{
					{Segment: "ICD", Mandatory: true, MaxRepeat: 1},  // Insurance cover description
					{Segment: "FTX", Mandatory: false, MaxRepeat: 1}, // Free text
					{Segment: "RFF", Mandatory: false, MaxRepeat: 9}, // Reference
					{ // Segment group 5
						Group: []SchemaNode{
							{Segment: "MOA", Mandatory: true, MaxRepeat: 1},  // Monetary amount
							{Segment: "DTM", Mandatory: false, MaxRepeat: 1}, // Date/time/period
							{ // Segment group 6
								Group: []SchemaNode{
									{Segment: "IND", Mandatory: true, MaxRepeat: 1},  // Index details
									{Segment: "DTM", Mandatory: false, MaxRepeat: 1}, // Date/time/period
								},
								Mandatory: false, MaxRepeat: 2,
							},
							{ // Segment group 7
								Group: []SchemaNode{
									{Segment: "ATT", Mandatory: true, MaxRepeat: 1},  // Attribute
									{Segment: "PCD", Mandatory: false, MaxRepeat: 1}, // Percentage details
								},
								Mandatory: false, MaxRepeat: 2,
							},
							{ // Segment group 8
								Group: []SchemaNode{
									{Segment: "TAX", Mandatory: true, MaxRepeat: 1},  // Duty/tax/fee details
									{Segment: "PCD", Mandatory: false, MaxRepeat: 1}, // Percentage details
								},
								Mandatory: false, MaxRepeat: 1,
							},
						},
						Mandatory: true, MaxRepeat: 99,
					},
				},
				Mandatory: true, MaxRepeat: 99,
			},
			{Segment: "CUX", Mandatory: true, MaxRepeat: 1}, // Currencies
		},
		Mandatory: true, MaxRepeat: 3,
	},
	{Segment: "MOA", Mandatory: true, MaxRepeat: 10}, // Monetary amount
	{Segment: "FTX", Mandatory: false, MaxRepeat: 1}, // Free text
	{ // Segment group 9
		Group: []SchemaNode{
			{Segment: "DTM", Mandatory: true, MaxRepeat: 1}, // Date/time/period
			{Segment: "MOA", Mandatory: true, MaxRepeat: 1}, // Monetary amount
		},
		Mandatory: false, MaxRepeat: 15,
	},
}}
