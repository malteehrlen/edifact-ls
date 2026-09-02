package edifact

// COMDIS D.20A message specification: the structural (segment/group
// presence, order, cardinality) branching diagram for this message,
// UN/EDIFACT directory release D.20A. Scope is structural only -- not
// element-level content -- per edifact-ls-3uzr's epic non-goal.
// 6 segment groups, max nesting depth 3.
//
// Source: https://service.unece.org/trade/untdid/d20a/trmd/comdis_c.htm
// section 4.3.1 "Segment table". That URL currently returns HTTP 403 from
// Cloudflare when fetched directly; this data was transcribed from the
// Wayback Machine's archived copy instead:
// http://web.archive.org/web/20231201123859/https://service.unece.org/trade/untdid/d20a/trmd/comdis_c.htm
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
		MessageID{Type: "COMDIS", Version: "D", Release: "20A", Agency: "UN"},
		comdisD20aSchema,
		"https://service.unece.org/trade/untdid/d20a/trmd/comdis_c.htm",
	)
}

var comdisD20aSchema = Schema{Nodes: []SchemaNode{
	{Segment: "BGM", Mandatory: true, MaxRepeat: 1},  // Beginning of message
	{Segment: "RFF", Mandatory: false, MaxRepeat: 9}, // Reference
	{Segment: "DTM", Mandatory: false, MaxRepeat: 9}, // Date/time/period
	{Segment: "CUX", Mandatory: false, MaxRepeat: 9}, // Currencies
	{ // Segment group 1
		Group: []SchemaNode{
			{Segment: "NAD", Mandatory: true, MaxRepeat: 1},  // Name and address
			{Segment: "CTA", Mandatory: false, MaxRepeat: 1}, // Contact information
			{Segment: "COM", Mandatory: false, MaxRepeat: 5}, // Communication contact
		},
		Mandatory: false, MaxRepeat: 99,
	},
	{ // Segment group 2
		Group: []SchemaNode{
			{Segment: "DOC", Mandatory: true, MaxRepeat: 1},  // Document/message details
			{Segment: "DTM", Mandatory: false, MaxRepeat: 9}, // Date/time/period
			{Segment: "MOA", Mandatory: false, MaxRepeat: 2}, // Monetary amount
			{ // Segment group 3
				Group: []SchemaNode{
					{Segment: "AJT", Mandatory: true, MaxRepeat: 1},  // Adjustment details
					{Segment: "FTX", Mandatory: false, MaxRepeat: 9}, // Free text
				},
				Mandatory: false, MaxRepeat: 9,
			},
			{ // Segment group 4
				Group: []SchemaNode{
					{Segment: "INP", Mandatory: true, MaxRepeat: 1},  // Parties and instruction
					{Segment: "FTX", Mandatory: false, MaxRepeat: 9}, // Free text
				},
				Mandatory: false, MaxRepeat: 9,
			},
			{ // Segment group 5
				Group: []SchemaNode{
					{Segment: "DLI", Mandatory: true, MaxRepeat: 1},  // Document line identification
					{Segment: "MOA", Mandatory: false, MaxRepeat: 2}, // Monetary amount
					{ // Segment group 6
						Group: []SchemaNode{
							{Segment: "AJT", Mandatory: true, MaxRepeat: 1},  // Adjustment details
							{Segment: "FTX", Mandatory: false, MaxRepeat: 9}, // Free text
						},
						Mandatory: false, MaxRepeat: 9,
					},
				},
				Mandatory: false, MaxRepeat: 9999,
			},
		},
		Mandatory: false, MaxRepeat: 9999,
	},
}}
