package edifact

// PRODEX D.20A message specification: the structural (segment/group
// presence, order, cardinality) branching diagram for this message,
// UN/EDIFACT directory release D.20A. Scope is structural only -- not
// element-level content -- per edifact-ls-3uzr's epic non-goal.
// 4 segment groups, max nesting depth 3.
//
// Source: https://service.unece.org/trade/untdid/d20a/trmd/prodex_c.htm
// section 4.3.1 "Segment table". That URL currently returns HTTP 403 from
// Cloudflare when fetched directly; this data was transcribed from the
// Wayback Machine's archived copy instead:
// http://web.archive.org/web/20240416120021/https://service.unece.org/trade/untdid/d20a/trmd/prodex_c.htm
// -- re-check the direct URL first if cross-checking later, in case the
// block has lifted.
//
// Transcription note: generated in bulk (edifact-ls-13gu) the same way
// as every other message-type schema -- a script parsed the source's
// exact rail-art column positions mechanically and verified the result
// balances (all 4 groups open and close consistently)
// before being transcribed here, rather than reading the ASCII nesting
// by eye. Not individually spot-checked against the raw source the way
// the first several message types in this project were -- see
// edifact-ls-13gu for which representative samples were spot-checked
// for this batch.

func init() {
	RegisterSchema(
		MessageID{Type: "PRODEX", Version: "D", Release: "20A", Agency: "UN"},
		prodexD20aSchema,
		"https://service.unece.org/trade/untdid/d20a/trmd/prodex_c.htm",
	)
}

var prodexD20aSchema = Schema{Nodes: []SchemaNode{
	{Segment: "BGM", Mandatory: true, MaxRepeat: 1}, // Beginning of message
	{Segment: "DTM", Mandatory: true, MaxRepeat: 2}, // Date/time/period
	{Segment: "MEA", Mandatory: true, MaxRepeat: 1}, // Measurements
	{Segment: "NAD", Mandatory: true, MaxRepeat: 2}, // Name and address
	{ // Segment group 1
		Group: []SchemaNode{
			{Segment: "RFF", Mandatory: true, MaxRepeat: 1},  // Reference
			{Segment: "DTM", Mandatory: false, MaxRepeat: 1}, // Date/time/period
		},
		Mandatory: true, MaxRepeat: 5,
	},
	{ // Segment group 2
		Group: []SchemaNode{
			{Segment: "IMD", Mandatory: true, MaxRepeat: 1},   // Item description
			{Segment: "QTY", Mandatory: false, MaxRepeat: 10}, // Quantity
			{ // Segment group 3
				Group: []SchemaNode{
					{Segment: "LIN", Mandatory: true, MaxRepeat: 1},  // Line item
					{Segment: "GEI", Mandatory: false, MaxRepeat: 2}, // Processing information
					{Segment: "LOC", Mandatory: false, MaxRepeat: 2}, // Place/location identification
					{Segment: "DTM", Mandatory: false, MaxRepeat: 1}, // Date/time/period
					{Segment: "MEA", Mandatory: false, MaxRepeat: 5}, // Measurements
					{Segment: "QTY", Mandatory: false, MaxRepeat: 5}, // Quantity
					{Segment: "TDT", Mandatory: false, MaxRepeat: 5}, // Transport information
					{ // Segment group 4
						Group: []SchemaNode{
							{Segment: "RFF", Mandatory: true, MaxRepeat: 1},  // Reference
							{Segment: "DTM", Mandatory: false, MaxRepeat: 1}, // Date/time/period
						},
						Mandatory: false, MaxRepeat: 5,
					},
				},
				Mandatory: false, MaxRepeat: 9999,
			},
		},
		Mandatory: true, MaxRepeat: 99,
	},
}}
