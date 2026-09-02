package edifact

// SOCADE D.20A message specification: the structural (segment/group
// presence, order, cardinality) branching diagram for this message,
// UN/EDIFACT directory release D.20A. Scope is structural only -- not
// element-level content -- per edifact-ls-3uzr's epic non-goal.
// 12 segment groups, max nesting depth 4.
//
// Source: https://service.unece.org/trade/untdid/d20a/trmd/socade_c.htm
// section 4.3.1 "Segment table". That URL currently returns HTTP 403 from
// Cloudflare when fetched directly; this data was transcribed from the
// Wayback Machine's archived copy instead:
// http://web.archive.org/web/20231206232336/https://service.unece.org/trade/untdid/d20a/trmd/socade_c.htm
// -- re-check the direct URL first if cross-checking later, in case the
// block has lifted.
//
// Transcription note: generated in bulk (edifact-ls-13gu) the same way
// as every other message-type schema -- a script parsed the source's
// exact rail-art column positions mechanically and verified the result
// balances (all 12 groups open and close consistently)
// before being transcribed here, rather than reading the ASCII nesting
// by eye. Not individually spot-checked against the raw source the way
// the first several message types in this project were -- see
// edifact-ls-13gu for which representative samples were spot-checked
// for this batch.

func init() {
	RegisterSchema(
		MessageID{Type: "SOCADE", Version: "D", Release: "20A", Agency: "UN"},
		socadeD20aSchema,
		"https://service.unece.org/trade/untdid/d20a/trmd/socade_c.htm",
	)
}

var socadeD20aSchema = Schema{Nodes: []SchemaNode{
	{Segment: "BGM", Mandatory: true, MaxRepeat: 1},  // Beginning of message
	{Segment: "DTM", Mandatory: true, MaxRepeat: 5},  // Date/time/period
	{Segment: "RFF", Mandatory: false, MaxRepeat: 2}, // Reference
	{ // Segment group 1
		Group: []SchemaNode{
			{Segment: "PNA", Mandatory: true, MaxRepeat: 1},  // Party identification
			{Segment: "RFF", Mandatory: false, MaxRepeat: 9}, // Reference
			{Segment: "CTA", Mandatory: false, MaxRepeat: 1}, // Contact information
			{Segment: "COM", Mandatory: false, MaxRepeat: 5}, // Communication contact
			{ // Segment group 2
				Group: []SchemaNode{
					{Segment: "ADR", Mandatory: true, MaxRepeat: 1}, // Address
					{Segment: "FTX", Mandatory: true, MaxRepeat: 9}, // Free text
				},
				Mandatory: false, MaxRepeat: 5,
			},
		},
		Mandatory: true, MaxRepeat: 9,
	},
	{Segment: "UNS", Mandatory: true, MaxRepeat: 1}, // Section control
	{ // Segment group 3
		Group: []SchemaNode{
			{Segment: "PNA", Mandatory: true, MaxRepeat: 1},   // Party identification
			{Segment: "FTX", Mandatory: false, MaxRepeat: 9},  // Free text
			{Segment: "PDI", Mandatory: false, MaxRepeat: 1},  // Person demographic information
			{Segment: "REL", Mandatory: false, MaxRepeat: 1},  // Relationship
			{Segment: "DTM", Mandatory: false, MaxRepeat: 9},  // Date/time/period
			{Segment: "NAT", Mandatory: false, MaxRepeat: 1},  // Nationality
			{Segment: "LOC", Mandatory: false, MaxRepeat: 1},  // Place/location identification
			{Segment: "FII", Mandatory: false, MaxRepeat: 1},  // Financial institution information
			{Segment: "ATT", Mandatory: false, MaxRepeat: 9},  // Attribute
			{Segment: "EMP", Mandatory: false, MaxRepeat: 1},  // Employment details
			{Segment: "MOA", Mandatory: false, MaxRepeat: 99}, // Monetary amount
			{Segment: "QTY", Mandatory: false, MaxRepeat: 99}, // Quantity
			{Segment: "CTA", Mandatory: false, MaxRepeat: 1},  // Contact information
			{Segment: "COM", Mandatory: false, MaxRepeat: 5},  // Communication contact
			{ // Segment group 4
				Group: []SchemaNode{
					{Segment: "RFF", Mandatory: true, MaxRepeat: 1},  // Reference
					{Segment: "DTM", Mandatory: false, MaxRepeat: 1}, // Date/time/period
				},
				Mandatory: false, MaxRepeat: 99,
			},
			{ // Segment group 5
				Group: []SchemaNode{
					{Segment: "ADR", Mandatory: true, MaxRepeat: 1},  // Address
					{Segment: "FTX", Mandatory: false, MaxRepeat: 1}, // Free text
				},
				Mandatory: false, MaxRepeat: 9,
			},
			{ // Segment group 6
				Group: []SchemaNode{
					{Segment: "PRC", Mandatory: true, MaxRepeat: 1},   // Process identification
					{Segment: "FTX", Mandatory: false, MaxRepeat: 99}, // Free text
					{Segment: "ATT", Mandatory: false, MaxRepeat: 9},  // Attribute
					{Segment: "DTM", Mandatory: false, MaxRepeat: 99}, // Date/time/period
					{Segment: "EMP", Mandatory: false, MaxRepeat: 9},  // Employment details
					{Segment: "ADR", Mandatory: false, MaxRepeat: 9},  // Address
					{Segment: "DOC", Mandatory: false, MaxRepeat: 9},  // Document/message details
					{Segment: "IND", Mandatory: false, MaxRepeat: 9},  // Index details
					{Segment: "STS", Mandatory: false, MaxRepeat: 9},  // Status
					{Segment: "LOC", Mandatory: false, MaxRepeat: 9},  // Place/location identification
					{ // Segment group 7
						Group: []SchemaNode{
							{Segment: "QTY", Mandatory: true, MaxRepeat: 1},  // Quantity
							{Segment: "DTM", Mandatory: false, MaxRepeat: 1}, // Date/time/period
						},
						Mandatory: false, MaxRepeat: 99,
					},
					{ // Segment group 8
						Group: []SchemaNode{
							{Segment: "RFF", Mandatory: true, MaxRepeat: 1},  // Reference
							{Segment: "DTM", Mandatory: false, MaxRepeat: 1}, // Date/time/period
							{Segment: "GEI", Mandatory: false, MaxRepeat: 9}, // Processing information
						},
						Mandatory: false, MaxRepeat: 99,
					},
					{ // Segment group 9
						Group: []SchemaNode{
							{Segment: "MOA", Mandatory: true, MaxRepeat: 1},  // Monetary amount
							{Segment: "RFF", Mandatory: false, MaxRepeat: 9}, // Reference
							{Segment: "DTM", Mandatory: false, MaxRepeat: 9}, // Date/time/period
						},
						Mandatory: false, MaxRepeat: 99,
					},
					{ // Segment group 10
						Group: []SchemaNode{
							{Segment: "PCD", Mandatory: true, MaxRepeat: 1}, // Percentage details
							{ // Segment group 11
								Group: []SchemaNode{
									{Segment: "MOA", Mandatory: true, MaxRepeat: 1},  // Monetary amount
									{Segment: "DTM", Mandatory: false, MaxRepeat: 9}, // Date/time/period
								},
								Mandatory: false, MaxRepeat: 9,
							},
						},
						Mandatory: false, MaxRepeat: 9,
					},
				},
				Mandatory: false, MaxRepeat: 99,
			},
		},
		Mandatory: true, MaxRepeat: 99999,
	},
	{ // Segment group 12
		Group: []SchemaNode{
			{Segment: "AUT", Mandatory: true, MaxRepeat: 1},  // Authentication result
			{Segment: "DTM", Mandatory: false, MaxRepeat: 1}, // Date/time/period
		},
		Mandatory: false, MaxRepeat: 1,
	},
}}
