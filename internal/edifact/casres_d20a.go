package edifact

// CASRES D.20A message specification: the structural (segment/group
// presence, order, cardinality) branching diagram for this message,
// UN/EDIFACT directory release D.20A. Scope is structural only -- not
// element-level content -- per edifact-ls-3uzr's epic non-goal.
// 2 segment groups, max nesting depth 1.
//
// Source: https://service.unece.org/trade/untdid/d20a/trmd/casres_c.htm
// section 4.3.1 "Segment table". That URL currently returns HTTP 403 from
// Cloudflare when fetched directly; this data was transcribed from the
// Wayback Machine's archived copy instead:
// http://web.archive.org/web/20240421091426/https://service.unece.org/trade/untdid/d20a/trmd/casres_c.htm
// -- re-check the direct URL first if cross-checking later, in case the
// block has lifted.
//
// Transcription note: generated in bulk (edifact-ls-13gu) the same way
// as every other message-type schema -- a script parsed the source's
// exact rail-art column positions mechanically and verified the result
// balances (all 2 groups open and close consistently)
// before being transcribed here, rather than reading the ASCII nesting
// by eye. Not individually spot-checked against the raw source the way
// the first several message types in this project were -- see
// edifact-ls-13gu for which representative samples were spot-checked
// for this batch.

func init() {
	RegisterSchema(
		MessageID{Type: "CASRES", Version: "D", Release: "20A", Agency: "UN"},
		casresD20aSchema,
		"https://service.unece.org/trade/untdid/d20a/trmd/casres_c.htm",
	)
}

var casresD20aSchema = Schema{Nodes: []SchemaNode{
	{Segment: "BGM", Mandatory: true, MaxRepeat: 1},   // Beginning of message
	{Segment: "DTM", Mandatory: true, MaxRepeat: 9},   // Date/time/period
	{Segment: "ERC", Mandatory: false, MaxRepeat: 9},  // Application error information
	{Segment: "LOC", Mandatory: false, MaxRepeat: 9},  // Place/location identification
	{Segment: "RFF", Mandatory: false, MaxRepeat: 9},  // Reference
	{Segment: "MOA", Mandatory: false, MaxRepeat: 99}, // Monetary amount
	{Segment: "FTX", Mandatory: false, MaxRepeat: 9},  // Free text
	{ // Segment group 1
		Group: []SchemaNode{
			{Segment: "NAD", Mandatory: true, MaxRepeat: 1},  // Name and address
			{Segment: "SEQ", Mandatory: false, MaxRepeat: 1}, // Sequence details
			{Segment: "FTX", Mandatory: false, MaxRepeat: 9}, // Free text
		},
		Mandatory: false, MaxRepeat: 99,
	},
	{ // Segment group 2
		Group: []SchemaNode{
			{Segment: "DOC", Mandatory: true, MaxRepeat: 1},   // Document/message details
			{Segment: "GIR", Mandatory: false, MaxRepeat: 99}, // Related identification numbers
			{Segment: "PYT", Mandatory: false, MaxRepeat: 1},  // Payment terms
			{Segment: "DTM", Mandatory: false, MaxRepeat: 9},  // Date/time/period
			{Segment: "MOA", Mandatory: false, MaxRepeat: 99}, // Monetary amount
			{Segment: "FTX", Mandatory: false, MaxRepeat: 9},  // Free text
		},
		Mandatory: false, MaxRepeat: 99,
	},
}}
