package edifact

// DMSTAT D.20A message specification: the structural (segment/group
// presence, order, cardinality) branching diagram for this message,
// UN/EDIFACT directory release D.20A. Scope is structural only -- not
// element-level content -- per edifact-ls-3uzr's epic non-goal.
// 4 segment groups, max nesting depth 2.
//
// Source: https://service.unece.org/trade/untdid/d20a/trmd/dmstat_c.htm
// section 4.3.1 "Segment table". That URL currently returns HTTP 403 from
// Cloudflare when fetched directly; this data was transcribed from the
// Wayback Machine's archived copy instead:
// http://web.archive.org/web/20231202092650/https://service.unece.org/trade/untdid/d20a/trmd/dmstat_c.htm
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
		MessageID{Type: "DMSTAT", Version: "D", Release: "20A", Agency: "UN"},
		dmstatD20aSchema,
		"https://service.unece.org/trade/untdid/d20a/trmd/dmstat_c.htm",
	)
}

var dmstatD20aSchema = Schema{Nodes: []SchemaNode{
	{Segment: "BGM", Mandatory: true, MaxRepeat: 1}, // Beginning of message
	{Segment: "DTM", Mandatory: true, MaxRepeat: 1}, // Date/time/period
	{ // Segment group 1
		Group: []SchemaNode{
			{Segment: "PNA", Mandatory: true, MaxRepeat: 1},  // Party identification
			{Segment: "COM", Mandatory: false, MaxRepeat: 9}, // Communication contact
		},
		Mandatory: true, MaxRepeat: 9,
	},
	{ // Segment group 2
		Group: []SchemaNode{
			{Segment: "HYN", Mandatory: true, MaxRepeat: 1},  // Hierarchy information
			{Segment: "RFF", Mandatory: false, MaxRepeat: 5}, // Reference
			{ // Segment group 3
				Group: []SchemaNode{
					{Segment: "STS", Mandatory: true, MaxRepeat: 1},    // Status
					{Segment: "PNA", Mandatory: false, MaxRepeat: 1},   // Party identification
					{Segment: "DTM", Mandatory: false, MaxRepeat: 1},   // Date/time/period
					{Segment: "RFF", Mandatory: false, MaxRepeat: 1},   // Reference
					{Segment: "FTX", Mandatory: false, MaxRepeat: 999}, // Free text
				},
				Mandatory: false, MaxRepeat: 9,
			},
			{ // Segment group 4
				Group: []SchemaNode{
					{Segment: "IRQ", Mandatory: true, MaxRepeat: 1},  // Information required
					{Segment: "RFF", Mandatory: false, MaxRepeat: 1}, // Reference
					{Segment: "PNA", Mandatory: false, MaxRepeat: 2}, // Party identification
					{Segment: "DTM", Mandatory: false, MaxRepeat: 1}, // Date/time/period
				},
				Mandatory: false, MaxRepeat: 9999,
			},
		},
		Mandatory: true, MaxRepeat: 999999,
	},
}}
