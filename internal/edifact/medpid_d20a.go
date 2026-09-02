package edifact

// MEDPID D.20A message specification: the structural (segment/group
// presence, order, cardinality) branching diagram for this message,
// UN/EDIFACT directory release D.20A. Scope is structural only -- not
// element-level content -- per edifact-ls-3uzr's epic non-goal.
// 6 segment groups, max nesting depth 2.
//
// Source: https://service.unece.org/trade/untdid/d20a/trmd/medpid_c.htm
// section 4.3.1 "Segment table". That URL currently returns HTTP 403 from
// Cloudflare when fetched directly; this data was transcribed from the
// Wayback Machine's archived copy instead:
// http://web.archive.org/web/20231207003921/https://service.unece.org/trade/untdid/d20a/trmd/medpid_c.htm
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
		MessageID{Type: "MEDPID", Version: "D", Release: "20A", Agency: "UN"},
		medpidD20aSchema,
		"https://service.unece.org/trade/untdid/d20a/trmd/medpid_c.htm",
	)
}

var medpidD20aSchema = Schema{Nodes: []SchemaNode{
	{Segment: "BGM", Mandatory: true, MaxRepeat: 1},   // Beginning of message
	{Segment: "DTM", Mandatory: false, MaxRepeat: 9},  // Date/time/period
	{Segment: "RFF", Mandatory: false, MaxRepeat: 9},  // Reference
	{Segment: "FTX", Mandatory: false, MaxRepeat: 99}, // Free text
	{ // Segment group 1
		Group: []SchemaNode{
			{Segment: "PNA", Mandatory: true, MaxRepeat: 1},   // Party identification
			{Segment: "ADR", Mandatory: false, MaxRepeat: 9},  // Address
			{Segment: "CTA", Mandatory: false, MaxRepeat: 9},  // Contact information
			{Segment: "COM", Mandatory: false, MaxRepeat: 9},  // Communication contact
			{Segment: "RFF", Mandatory: false, MaxRepeat: 99}, // Reference
			{Segment: "LAN", Mandatory: false, MaxRepeat: 9},  // Language
		},
		Mandatory: true, MaxRepeat: 99,
	},
	{ // Segment group 2
		Group: []SchemaNode{
			{Segment: "GEI", Mandatory: true, MaxRepeat: 1},   // Processing information
			{Segment: "PNA", Mandatory: false, MaxRepeat: 1},  // Party identification
			{Segment: "ADR", Mandatory: false, MaxRepeat: 9},  // Address
			{Segment: "RFF", Mandatory: false, MaxRepeat: 99}, // Reference
			{Segment: "DTM", Mandatory: false, MaxRepeat: 9},  // Date/time/period
			{Segment: "IHC", Mandatory: false, MaxRepeat: 9},  // Person characteristic
			{Segment: "NAT", Mandatory: false, MaxRepeat: 9},  // Nationality
			{Segment: "FTX", Mandatory: false, MaxRepeat: 9},  // Free text
			{Segment: "LAN", Mandatory: false, MaxRepeat: 9},  // Language
			{Segment: "HAN", Mandatory: false, MaxRepeat: 9},  // Handling instructions
			{Segment: "LOC", Mandatory: false, MaxRepeat: 9},  // Place/location identification
			{Segment: "FII", Mandatory: false, MaxRepeat: 9},  // Financial institution information
			{Segment: "CTA", Mandatory: false, MaxRepeat: 9},  // Contact information
			{ // Segment group 3
				Group: []SchemaNode{
					{Segment: "PDI", Mandatory: true, MaxRepeat: 1},  // Person demographic information
					{Segment: "DTM", Mandatory: false, MaxRepeat: 9}, // Date/time/period
				},
				Mandatory: false, MaxRepeat: 9,
			},
			{ // Segment group 4
				Group: []SchemaNode{
					{Segment: "COM", Mandatory: true, MaxRepeat: 1},  // Communication contact
					{Segment: "CTA", Mandatory: false, MaxRepeat: 9}, // Contact information
				},
				Mandatory: false, MaxRepeat: 9,
			},
			{ // Segment group 5
				Group: []SchemaNode{
					{Segment: "REL", Mandatory: true, MaxRepeat: 1},  // Relationship
					{Segment: "PNA", Mandatory: false, MaxRepeat: 1}, // Party identification
					{Segment: "ADR", Mandatory: false, MaxRepeat: 9}, // Address
					{Segment: "RFF", Mandatory: false, MaxRepeat: 9}, // Reference
					{Segment: "DTM", Mandatory: false, MaxRepeat: 9}, // Date/time/period
					{Segment: "PDI", Mandatory: false, MaxRepeat: 1}, // Person demographic information
					{Segment: "IHC", Mandatory: false, MaxRepeat: 9}, // Person characteristic
					{Segment: "NAT", Mandatory: false, MaxRepeat: 9}, // Nationality
					{Segment: "LAN", Mandatory: false, MaxRepeat: 9}, // Language
				},
				Mandatory: false, MaxRepeat: 99,
			},
		},
		Mandatory: true, MaxRepeat: 9999,
	},
	{ // Segment group 6
		Group: []SchemaNode{
			{Segment: "AUT", Mandatory: true, MaxRepeat: 1},  // Authentication result
			{Segment: "DTM", Mandatory: false, MaxRepeat: 9}, // Date/time/period
		},
		Mandatory: false, MaxRepeat: 9,
	},
}}
