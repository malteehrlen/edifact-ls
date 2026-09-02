package edifact

// DGRECA D.20A message specification: the structural (segment/group
// presence, order, cardinality) branching diagram for this message,
// UN/EDIFACT directory release D.20A. Scope is structural only -- not
// element-level content -- per edifact-ls-3uzr's epic non-goal.
// 5 segment groups, max nesting depth 2.
//
// Source: https://service.unece.org/trade/untdid/d20a/trmd/dgreca_c.htm
// section 4.3.1 "Segment table". That URL currently returns HTTP 403 from
// Cloudflare when fetched directly; this data was transcribed from the
// Wayback Machine's archived copy instead:
// http://web.archive.org/web/20240416115234/https://service.unece.org/trade/untdid/d20a/trmd/dgreca_c.htm
// -- re-check the direct URL first if cross-checking later, in case the
// block has lifted.
//
// Transcription note: generated in bulk (edifact-ls-13gu) the same way
// as every other message-type schema -- a script parsed the source's
// exact rail-art column positions mechanically and verified the result
// balances (all 5 groups open and close consistently)
// before being transcribed here, rather than reading the ASCII nesting
// by eye. Not individually spot-checked against the raw source the way
// the first several message types in this project were -- see
// edifact-ls-13gu for which representative samples were spot-checked
// for this batch.

func init() {
	RegisterSchema(
		MessageID{Type: "DGRECA", Version: "D", Release: "20A", Agency: "UN"},
		dgrecaD20aSchema,
		"https://service.unece.org/trade/untdid/d20a/trmd/dgreca_c.htm",
	)
}

var dgrecaD20aSchema = Schema{Nodes: []SchemaNode{
	{Segment: "BGM", Mandatory: true, MaxRepeat: 1},  // Beginning of message
	{Segment: "DTM", Mandatory: false, MaxRepeat: 9}, // Date/time/period
	{ // Segment group 1
		Group: []SchemaNode{
			{Segment: "RFF", Mandatory: true, MaxRepeat: 1},  // Reference
			{Segment: "DTM", Mandatory: false, MaxRepeat: 9}, // Date/time/period
		},
		Mandatory: false, MaxRepeat: 9,
	},
	{ // Segment group 2
		Group: []SchemaNode{
			{Segment: "TDT", Mandatory: true, MaxRepeat: 1},  // Transport information
			{Segment: "DTM", Mandatory: true, MaxRepeat: 2},  // Date/time/period
			{Segment: "LOC", Mandatory: true, MaxRepeat: 2},  // Place/location identification
			{Segment: "RFF", Mandatory: false, MaxRepeat: 2}, // Reference
			{Segment: "MEA", Mandatory: false, MaxRepeat: 1}, // Measurements
		},
		Mandatory: true, MaxRepeat: 1,
	},
	{ // Segment group 3
		Group: []SchemaNode{
			{Segment: "NAD", Mandatory: true, MaxRepeat: 1}, // Name and address
			{ // Segment group 4
				Group: []SchemaNode{
					{Segment: "CTA", Mandatory: true, MaxRepeat: 1},  // Contact information
					{Segment: "COM", Mandatory: false, MaxRepeat: 3}, // Communication contact
				},
				Mandatory: false, MaxRepeat: 1,
			},
		},
		Mandatory: true, MaxRepeat: 1,
	},
	{ // Segment group 5
		Group: []SchemaNode{
			{Segment: "DGS", Mandatory: true, MaxRepeat: 1},   // Dangerous goods
			{Segment: "FTX", Mandatory: true, MaxRepeat: 9},   // Free text
			{Segment: "MEA", Mandatory: true, MaxRepeat: 1},   // Measurements
			{Segment: "LOC", Mandatory: false, MaxRepeat: 99}, // Place/location identification
			{Segment: "SGP", Mandatory: false, MaxRepeat: 99}, // Split goods placement
		},
		Mandatory: true, MaxRepeat: 999,
	},
}}
