package edifact

// PAXLST D.20A message specification: the structural (segment/group
// presence, order, cardinality) branching diagram for this message,
// UN/EDIFACT directory release D.20A. Scope is structural only -- not
// element-level content -- per edifact-ls-3uzr's epic non-goal.
// 7 segment groups, max nesting depth 2.
//
// Source: https://service.unece.org/trade/untdid/d20a/trmd/paxlst_c.htm
// section 4.3.1 "Segment table". That URL currently returns HTTP 403 from
// Cloudflare when fetched directly; this data was transcribed from the
// Wayback Machine's archived copy instead:
// http://web.archive.org/web/20231202081939/https://service.unece.org/trade/untdid/d20a/trmd/paxlst_c.htm
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
		MessageID{Type: "PAXLST", Version: "D", Release: "20A", Agency: "UN"},
		paxlstD20aSchema,
		"https://service.unece.org/trade/untdid/d20a/trmd/paxlst_c.htm",
	)
}

var paxlstD20aSchema = Schema{Nodes: []SchemaNode{
	{Segment: "BGM", Mandatory: true, MaxRepeat: 1},  // Beginning of message
	{Segment: "RFF", Mandatory: false, MaxRepeat: 1}, // Reference
	{Segment: "DTM", Mandatory: false, MaxRepeat: 9}, // Date/time/period
	{Segment: "LOC", Mandatory: false, MaxRepeat: 9}, // Place/location identification
	{ // Segment group 1
		Group: []SchemaNode{
			{Segment: "NAD", Mandatory: true, MaxRepeat: 1},  // Name and address
			{Segment: "CTA", Mandatory: false, MaxRepeat: 1}, // Contact information
			{Segment: "COM", Mandatory: false, MaxRepeat: 9}, // Communication contact
		},
		Mandatory: false, MaxRepeat: 5,
	},
	{ // Segment group 2
		Group: []SchemaNode{
			{Segment: "TDT", Mandatory: true, MaxRepeat: 1},  // Transport information
			{Segment: "DTM", Mandatory: false, MaxRepeat: 1}, // Date/time/period
			{ // Segment group 3
				Group: []SchemaNode{
					{Segment: "LOC", Mandatory: true, MaxRepeat: 1},  // Place/location identification
					{Segment: "DTM", Mandatory: false, MaxRepeat: 5}, // Date/time/period
				},
				Mandatory: false, MaxRepeat: 99,
			},
		},
		Mandatory: true, MaxRepeat: 10,
	},
	{ // Segment group 4
		Group: []SchemaNode{
			{Segment: "NAD", Mandatory: true, MaxRepeat: 1},      // Name and address
			{Segment: "ATT", Mandatory: false, MaxRepeat: 9},     // Attribute
			{Segment: "DTM", Mandatory: false, MaxRepeat: 9},     // Date/time/period
			{Segment: "MEA", Mandatory: false, MaxRepeat: 9},     // Measurements
			{Segment: "GEI", Mandatory: false, MaxRepeat: 5},     // Processing information
			{Segment: "FTX", Mandatory: false, MaxRepeat: 99},    // Free text
			{Segment: "LOC", Mandatory: false, MaxRepeat: 25},    // Place/location identification
			{Segment: "COM", Mandatory: false, MaxRepeat: 1},     // Communication contact
			{Segment: "EMP", Mandatory: false, MaxRepeat: 9},     // Employment details
			{Segment: "NAT", Mandatory: false, MaxRepeat: 9},     // Nationality
			{Segment: "RFF", Mandatory: false, MaxRepeat: 99999}, // Reference
			{ // Segment group 5
				Group: []SchemaNode{
					{Segment: "DOC", Mandatory: true, MaxRepeat: 1},  // Document/message details
					{Segment: "DTM", Mandatory: false, MaxRepeat: 5}, // Date/time/period
					{Segment: "GEI", Mandatory: false, MaxRepeat: 9}, // Processing information
					{Segment: "RFF", Mandatory: false, MaxRepeat: 9}, // Reference
					{Segment: "LOC", Mandatory: false, MaxRepeat: 2}, // Place/location identification
					{Segment: "CPI", Mandatory: false, MaxRepeat: 1}, // Charge payment instructions
					{Segment: "QTY", Mandatory: false, MaxRepeat: 9}, // Quantity
				},
				Mandatory: false, MaxRepeat: 5,
			},
			{ // Segment group 6
				Group: []SchemaNode{
					{Segment: "GID", Mandatory: true, MaxRepeat: 1},  // Goods item details
					{Segment: "FTX", Mandatory: false, MaxRepeat: 9}, // Free text
					{Segment: "QTY", Mandatory: false, MaxRepeat: 9}, // Quantity
				},
				Mandatory: false, MaxRepeat: 99,
			},
			{ // Segment group 7
				Group: []SchemaNode{
					{Segment: "TDT", Mandatory: true, MaxRepeat: 1},  // Transport information
					{Segment: "FTX", Mandatory: false, MaxRepeat: 1}, // Free text
				},
				Mandatory: false, MaxRepeat: 1,
			},
		},
		Mandatory: false, MaxRepeat: 99999,
	},
	{Segment: "CNT", Mandatory: false, MaxRepeat: 1}, // Control total
	{Segment: "AUT", Mandatory: false, MaxRepeat: 1}, // Authentication result
}}
