package edifact

// BANSTA D.20A message specification: the structural (segment/group
// presence, order, cardinality) branching diagram for this message,
// UN/EDIFACT directory release D.20A. Scope is structural only -- not
// element-level content -- per edifact-ls-3uzr's epic non-goal.
// 9 segment groups, max nesting depth 3.
//
// Source: https://service.unece.org/trade/untdid/d20a/trmd/bansta_c.htm
// section 4.3.1 "Segment table". That URL currently returns HTTP 403 from
// Cloudflare when fetched directly; this data was transcribed from the
// Wayback Machine's archived copy instead:
// http://web.archive.org/web/20231208120815/https://service.unece.org/trade/untdid/d20a/trmd/bansta_c.htm
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
		MessageID{Type: "BANSTA", Version: "D", Release: "20A", Agency: "UN"},
		banstaD20aSchema,
		"https://service.unece.org/trade/untdid/d20a/trmd/bansta_c.htm",
	)
}

var banstaD20aSchema = Schema{Nodes: []SchemaNode{
	{Segment: "BGM", Mandatory: true, MaxRepeat: 1},  // Beginning of message
	{Segment: "DTM", Mandatory: true, MaxRepeat: 1},  // Date/time/period
	{Segment: "BUS", Mandatory: false, MaxRepeat: 1}, // Business function
	{ // Segment group 1
		Group: []SchemaNode{
			{Segment: "RFF", Mandatory: true, MaxRepeat: 1},  // Reference
			{Segment: "DTM", Mandatory: false, MaxRepeat: 1}, // Date/time/period
		},
		Mandatory: false, MaxRepeat: 2,
	},
	{ // Segment group 2
		Group: []SchemaNode{
			{Segment: "FII", Mandatory: true, MaxRepeat: 1},  // Financial institution information
			{Segment: "CTA", Mandatory: false, MaxRepeat: 1}, // Contact information
			{Segment: "COM", Mandatory: false, MaxRepeat: 5}, // Communication contact
		},
		Mandatory: false, MaxRepeat: 5,
	},
	{ // Segment group 3
		Group: []SchemaNode{
			{Segment: "NAD", Mandatory: true, MaxRepeat: 1},  // Name and address
			{Segment: "CTA", Mandatory: false, MaxRepeat: 1}, // Contact information
			{Segment: "COM", Mandatory: false, MaxRepeat: 5}, // Communication contact
		},
		Mandatory: false, MaxRepeat: 3,
	},
	{ // Segment group 4
		Group: []SchemaNode{
			{Segment: "LIN", Mandatory: true, MaxRepeat: 1}, // Line item
			{ // Segment group 5
				Group: []SchemaNode{
					{Segment: "RFF", Mandatory: true, MaxRepeat: 1},  // Reference
					{Segment: "DTM", Mandatory: false, MaxRepeat: 1}, // Date/time/period
				},
				Mandatory: false, MaxRepeat: 5,
			},
			{ // Segment group 6
				Group: []SchemaNode{
					{Segment: "SEQ", Mandatory: true, MaxRepeat: 1},  // Sequence details
					{Segment: "GEI", Mandatory: true, MaxRepeat: 1},  // Processing information
					{Segment: "DTM", Mandatory: false, MaxRepeat: 2}, // Date/time/period
					{Segment: "MOA", Mandatory: false, MaxRepeat: 1}, // Monetary amount
					{Segment: "CUX", Mandatory: false, MaxRepeat: 1}, // Currencies
					{Segment: "PCD", Mandatory: false, MaxRepeat: 1}, // Percentage details
					{Segment: "FTX", Mandatory: false, MaxRepeat: 1}, // Free text
					{Segment: "DOC", Mandatory: false, MaxRepeat: 5}, // Document/message details
					{ // Segment group 7
						Group: []SchemaNode{
							{Segment: "FII", Mandatory: true, MaxRepeat: 1},  // Financial institution information
							{Segment: "CTA", Mandatory: false, MaxRepeat: 1}, // Contact information
							{Segment: "COM", Mandatory: false, MaxRepeat: 5}, // Communication contact
						},
						Mandatory: false, MaxRepeat: 1,
					},
					{ // Segment group 8
						Group: []SchemaNode{
							{Segment: "NAD", Mandatory: true, MaxRepeat: 1},  // Name and address
							{Segment: "CTA", Mandatory: false, MaxRepeat: 1}, // Contact information
							{Segment: "COM", Mandatory: false, MaxRepeat: 5}, // Communication contact
						},
						Mandatory: false, MaxRepeat: 1,
					},
				},
				Mandatory: false, MaxRepeat: 99,
			},
		},
		Mandatory: true, MaxRepeat: 999,
	},
	{Segment: "CNT", Mandatory: false, MaxRepeat: 5}, // Control total
	{ // Segment group 9
		Group: []SchemaNode{
			{Segment: "AUT", Mandatory: true, MaxRepeat: 1},  // Authentication result
			{Segment: "DTM", Mandatory: false, MaxRepeat: 1}, // Date/time/period
		},
		Mandatory: false, MaxRepeat: 5,
	},
}}
