package edifact

// CONRPW D.20A message specification: the structural (segment/group
// presence, order, cardinality) branching diagram for this message,
// UN/EDIFACT directory release D.20A. Scope is structural only -- not
// element-level content -- per edifact-ls-3uzr's epic non-goal.
// 2 segment groups, max nesting depth 2.
//
// Source: https://service.unece.org/trade/untdid/d20a/trmd/conrpw_c.htm
// section 4.3.1 "Segment table". That URL currently returns HTTP 403 from
// Cloudflare when fetched directly; this data was transcribed from the
// Wayback Machine's archived copy instead:
// http://web.archive.org/web/20240416105123/https://service.unece.org/trade/untdid/d20a/trmd/conrpw_c.htm
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
		MessageID{Type: "CONRPW", Version: "D", Release: "20A", Agency: "UN"},
		conrpwD20aSchema,
		"https://service.unece.org/trade/untdid/d20a/trmd/conrpw_c.htm",
	)
}

var conrpwD20aSchema = Schema{Nodes: []SchemaNode{
	{Segment: "BGM", Mandatory: true, MaxRepeat: 1}, // Beginning of message
	{Segment: "RFF", Mandatory: true, MaxRepeat: 9}, // Reference
	{Segment: "DTM", Mandatory: true, MaxRepeat: 9}, // Date/time/period
	{ // Segment group 1
		Group: []SchemaNode{
			{Segment: "NAD", Mandatory: true, MaxRepeat: 1}, // Name and address
			{ // Segment group 2
				Group: []SchemaNode{
					{Segment: "CTA", Mandatory: true, MaxRepeat: 1},  // Contact information
					{Segment: "COM", Mandatory: false, MaxRepeat: 3}, // Communication contact
				},
				Mandatory: false, MaxRepeat: 2,
			},
		},
		Mandatory: true, MaxRepeat: 3,
	},
	{Segment: "LOC", Mandatory: true, MaxRepeat: 99}, // Place/location identification
	{Segment: "FTX", Mandatory: true, MaxRepeat: 99}, // Free text
	{Segment: "DOC", Mandatory: false, MaxRepeat: 9}, // Document/message details
	{Segment: "CNT", Mandatory: false, MaxRepeat: 5}, // Control total
	{Segment: "AUT", Mandatory: false, MaxRepeat: 1}, // Authentication result
}}
