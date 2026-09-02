package edifact

// JINFDE D.20A message specification: the structural (segment/group
// presence, order, cardinality) branching diagram for this message,
// UN/EDIFACT directory release D.20A. Scope is structural only -- not
// element-level content -- per edifact-ls-3uzr's epic non-goal.
// 9 segment groups, max nesting depth 3.
//
// Source: https://service.unece.org/trade/untdid/d20a/trmd/jinfde_c.htm
// section 4.3.1 "Segment table". That URL currently returns HTTP 403 from
// Cloudflare when fetched directly; this data was transcribed from the
// Wayback Machine's archived copy instead:
// http://web.archive.org/web/20231207010150/https://service.unece.org/trade/untdid/d20a/trmd/jinfde_c.htm
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
		MessageID{Type: "JINFDE", Version: "D", Release: "20A", Agency: "UN"},
		jinfdeD20aSchema,
		"https://service.unece.org/trade/untdid/d20a/trmd/jinfde_c.htm",
	)
}

var jinfdeD20aSchema = Schema{Nodes: []SchemaNode{
	{Segment: "BGM", Mandatory: true, MaxRepeat: 1}, // Beginning of message
	{Segment: "DTM", Mandatory: true, MaxRepeat: 2}, // Date/time/period
	{ // Segment group 1
		Group: []SchemaNode{
			{Segment: "PNA", Mandatory: true, MaxRepeat: 1},  // Party identification
			{Segment: "ADR", Mandatory: false, MaxRepeat: 1}, // Address
			{Segment: "LOC", Mandatory: false, MaxRepeat: 5}, // Place/location identification
			{Segment: "DTM", Mandatory: false, MaxRepeat: 5}, // Date/time/period
			{ // Segment group 2
				Group: []SchemaNode{
					{Segment: "CTA", Mandatory: true, MaxRepeat: 1},  // Contact information
					{Segment: "COM", Mandatory: false, MaxRepeat: 5}, // Communication contact
				},
				Mandatory: false, MaxRepeat: 5,
			},
		},
		Mandatory: true, MaxRepeat: 9,
	},
	{ // Segment group 3
		Group: []SchemaNode{
			{Segment: "RFF", Mandatory: true, MaxRepeat: 1},  // Reference
			{Segment: "DTM", Mandatory: false, MaxRepeat: 5}, // Date/time/period
		},
		Mandatory: false, MaxRepeat: 5,
	},
	{ // Segment group 4
		Group: []SchemaNode{
			{Segment: "GEI", Mandatory: true, MaxRepeat: 1},  // Processing information
			{Segment: "FTX", Mandatory: false, MaxRepeat: 1}, // Free text
		},
		Mandatory: false, MaxRepeat: 5,
	},
	{Segment: "UNS", Mandatory: true, MaxRepeat: 1}, // Section control
	{ // Segment group 5
		Group: []SchemaNode{
			{Segment: "RFF", Mandatory: true, MaxRepeat: 1}, // Reference
			{ // Segment group 6
				Group: []SchemaNode{
					{Segment: "PNA", Mandatory: true, MaxRepeat: 1},  // Party identification
					{Segment: "ADR", Mandatory: false, MaxRepeat: 1}, // Address
					{Segment: "LOC", Mandatory: false, MaxRepeat: 2}, // Place/location identification
					{Segment: "DTM", Mandatory: false, MaxRepeat: 5}, // Date/time/period
					{Segment: "NAT", Mandatory: false, MaxRepeat: 9}, // Nationality
					{Segment: "PDI", Mandatory: false, MaxRepeat: 1}, // Person demographic information
					{Segment: "DOC", Mandatory: false, MaxRepeat: 9}, // Document/message details
					{ // Segment group 7
						Group: []SchemaNode{
							{Segment: "CTA", Mandatory: true, MaxRepeat: 1},  // Contact information
							{Segment: "COM", Mandatory: false, MaxRepeat: 5}, // Communication contact
						},
						Mandatory: false, MaxRepeat: 5,
					},
				},
				Mandatory: true, MaxRepeat: 999,
			},
			{ // Segment group 8
				Group: []SchemaNode{
					{Segment: "RFF", Mandatory: true, MaxRepeat: 1},  // Reference
					{Segment: "DTM", Mandatory: false, MaxRepeat: 5}, // Date/time/period
				},
				Mandatory: false, MaxRepeat: 5,
			},
			{Segment: "FTX", Mandatory: true, MaxRepeat: 30}, // Free text
			{ // Segment group 9
				Group: []SchemaNode{
					{Segment: "GEI", Mandatory: true, MaxRepeat: 1},  // Processing information
					{Segment: "FTX", Mandatory: false, MaxRepeat: 1}, // Free text
				},
				Mandatory: false, MaxRepeat: 30,
			},
		},
		Mandatory: true, MaxRepeat: 999,
	},
}}
