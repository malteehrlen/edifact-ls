package edifact

// SUPMAN D.20A message specification: the structural (segment/group
// presence, order, cardinality) branching diagram for this message,
// UN/EDIFACT directory release D.20A. Scope is structural only -- not
// element-level content -- per edifact-ls-3uzr's epic non-goal.
// 10 segment groups, max nesting depth 3.
//
// Source: https://service.unece.org/trade/untdid/d20a/trmd/supman_c.htm
// section 4.3.1 "Segment table". That URL currently returns HTTP 403 from
// Cloudflare when fetched directly; this data was transcribed from the
// Wayback Machine's archived copy instead:
// http://web.archive.org/web/20231201135149/https://service.unece.org/trade/untdid/d20a/trmd/supman_c.htm
// -- re-check the direct URL first if cross-checking later, in case the
// block has lifted.
//
// Transcription note: generated in bulk (edifact-ls-13gu) the same way
// as every other message-type schema -- a script parsed the source's
// exact rail-art column positions mechanically and verified the result
// balances (all 10 groups open and close consistently)
// before being transcribed here, rather than reading the ASCII nesting
// by eye. Not individually spot-checked against the raw source the way
// the first several message types in this project were -- see
// edifact-ls-13gu for which representative samples were spot-checked
// for this batch.

func init() {
	RegisterSchema(
		MessageID{Type: "SUPMAN", Version: "D", Release: "20A", Agency: "UN"},
		supmanD20aSchema,
		"https://service.unece.org/trade/untdid/d20a/trmd/supman_c.htm",
	)
}

var supmanD20aSchema = Schema{Nodes: []SchemaNode{
	{Segment: "BGM", Mandatory: true, MaxRepeat: 1},  // Beginning of message
	{Segment: "RFF", Mandatory: true, MaxRepeat: 6},  // Reference
	{Segment: "CUX", Mandatory: false, MaxRepeat: 1}, // Currencies
	{Segment: "DTM", Mandatory: false, MaxRepeat: 9}, // Date/time/period
	{Segment: "FTX", Mandatory: false, MaxRepeat: 5}, // Free text
	{ // Segment group 1
		Group: []SchemaNode{
			{Segment: "NAD", Mandatory: true, MaxRepeat: 1}, // Name and address
			{ // Segment group 2
				Group: []SchemaNode{
					{Segment: "CTA", Mandatory: true, MaxRepeat: 1},  // Contact information
					{Segment: "COM", Mandatory: false, MaxRepeat: 1}, // Communication contact
				},
				Mandatory: false, MaxRepeat: 5,
			},
		},
		Mandatory: false, MaxRepeat: 6,
	},
	{Segment: "UNS", Mandatory: true, MaxRepeat: 1}, // Section control
	{ // Segment group 3
		Group: []SchemaNode{
			{Segment: "NAD", Mandatory: true, MaxRepeat: 1},  // Name and address
			{Segment: "DTM", Mandatory: true, MaxRepeat: 15}, // Date/time/period
			{Segment: "ATT", Mandatory: false, MaxRepeat: 9}, // Attribute
			{Segment: "RFF", Mandatory: false, MaxRepeat: 9}, // Reference
			{ // Segment group 4
				Group: []SchemaNode{
					{Segment: "REL", Mandatory: true, MaxRepeat: 1},  // Relationship
					{Segment: "NAD", Mandatory: true, MaxRepeat: 1},  // Name and address
					{Segment: "PCD", Mandatory: false, MaxRepeat: 1}, // Percentage details
				},
				Mandatory: false, MaxRepeat: 99,
			},
			{ // Segment group 5
				Group: []SchemaNode{
					{Segment: "EMP", Mandatory: true, MaxRepeat: 1},  // Employment details
					{Segment: "PCD", Mandatory: false, MaxRepeat: 1}, // Percentage details
					{Segment: "CUX", Mandatory: false, MaxRepeat: 1}, // Currencies
					{Segment: "NAD", Mandatory: false, MaxRepeat: 9}, // Name and address
					{ // Segment group 6
						Group: []SchemaNode{
							{Segment: "MOA", Mandatory: true, MaxRepeat: 1},  // Monetary amount
							{Segment: "PYT", Mandatory: false, MaxRepeat: 1}, // Payment terms
							{Segment: "DTM", Mandatory: false, MaxRepeat: 1}, // Date/time/period
						},
						Mandatory: false, MaxRepeat: 9,
					},
				},
				Mandatory: false, MaxRepeat: 9,
			},
			{ // Segment group 7
				Group: []SchemaNode{
					{Segment: "GEI", Mandatory: true, MaxRepeat: 1},  // Processing information
					{Segment: "DTM", Mandatory: false, MaxRepeat: 1}, // Date/time/period
				},
				Mandatory: true, MaxRepeat: 20,
			},
			{ // Segment group 8
				Group: []SchemaNode{
					{Segment: "MEM", Mandatory: true, MaxRepeat: 1}, // Membership details
					{ // Segment group 9
						Group: []SchemaNode{
							{Segment: "ATT", Mandatory: true, MaxRepeat: 1},  // Attribute
							{Segment: "PCD", Mandatory: false, MaxRepeat: 1}, // Percentage details
						},
						Mandatory: false, MaxRepeat: 9,
					},
					{ // Segment group 10
						Group: []SchemaNode{
							{Segment: "COT", Mandatory: true, MaxRepeat: 1},  // Contribution details
							{Segment: "MOA", Mandatory: false, MaxRepeat: 1}, // Monetary amount
							{Segment: "PCD", Mandatory: false, MaxRepeat: 1}, // Percentage details
							{Segment: "DTM", Mandatory: false, MaxRepeat: 3}, // Date/time/period
							{Segment: "PYT", Mandatory: false, MaxRepeat: 1}, // Payment terms
							{Segment: "FTX", Mandatory: false, MaxRepeat: 3}, // Free text
						},
						Mandatory: false, MaxRepeat: 99,
					},
				},
				Mandatory: false, MaxRepeat: 9,
			},
		},
		Mandatory: true, MaxRepeat: 999999,
	},
	{Segment: "CNT", Mandatory: false, MaxRepeat: 9}, // Control total
	{Segment: "AUT", Mandatory: false, MaxRepeat: 1}, // Authentication result
}}
