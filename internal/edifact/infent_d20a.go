package edifact

// INFENT D.20A message specification: the structural (segment/group
// presence, order, cardinality) branching diagram for this message,
// UN/EDIFACT directory release D.20A. Scope is structural only -- not
// element-level content -- per edifact-ls-3uzr's epic non-goal.
// 8 segment groups, max nesting depth 3.
//
// Source: https://service.unece.org/trade/untdid/d20a/trmd/infent_c.htm
// section 4.3.1 "Segment table". That URL currently returns HTTP 403 from
// Cloudflare when fetched directly; this data was transcribed from the
// Wayback Machine's archived copy instead:
// http://web.archive.org/web/20231202084937/https://service.unece.org/trade/untdid/d20a/trmd/infent_c.htm
// -- re-check the direct URL first if cross-checking later, in case the
// block has lifted.
//
// Transcription note: generated in bulk (edifact-ls-13gu) the same way
// as every other message-type schema -- a script parsed the source's
// exact rail-art column positions mechanically and verified the result
// balances (all 8 groups open and close consistently)
// before being transcribed here, rather than reading the ASCII nesting
// by eye. Not individually spot-checked against the raw source the way
// the first several message types in this project were -- see
// edifact-ls-13gu for which representative samples were spot-checked
// for this batch.

func init() {
	RegisterSchema(
		MessageID{Type: "INFENT", Version: "D", Release: "20A", Agency: "UN"},
		infentD20aSchema,
		"https://service.unece.org/trade/untdid/d20a/trmd/infent_c.htm",
	)
}

var infentD20aSchema = Schema{Nodes: []SchemaNode{
	{Segment: "BGM", Mandatory: true, MaxRepeat: 1},   // Beginning of message
	{Segment: "DTM", Mandatory: true, MaxRepeat: 99},  // Date/time/period
	{Segment: "RFF", Mandatory: false, MaxRepeat: 99}, // Reference
	{Segment: "CUX", Mandatory: false, MaxRepeat: 99}, // Currencies
	{Segment: "FTX", Mandatory: false, MaxRepeat: 1},  // Free text
	{ // Segment group 1
		Group: []SchemaNode{
			{Segment: "NAD", Mandatory: true, MaxRepeat: 1},  // Name and address
			{Segment: "RFF", Mandatory: false, MaxRepeat: 9}, // Reference
			{Segment: "FII", Mandatory: false, MaxRepeat: 1}, // Financial institution information
			{ // Segment group 2
				Group: []SchemaNode{
					{Segment: "CTA", Mandatory: true, MaxRepeat: 1},  // Contact information
					{Segment: "COM", Mandatory: false, MaxRepeat: 9}, // Communication contact
				},
				Mandatory: false, MaxRepeat: 99,
			},
		},
		Mandatory: true, MaxRepeat: 99,
	},
	{ // Segment group 3
		Group: []SchemaNode{
			{Segment: "CCI", Mandatory: true, MaxRepeat: 1},  // Characteristic/class id
			{Segment: "CAV", Mandatory: false, MaxRepeat: 1}, // Characteristic value
		},
		Mandatory: false, MaxRepeat: 99,
	},
	{ // Segment group 4
		Group: []SchemaNode{
			{Segment: "SEQ", Mandatory: true, MaxRepeat: 1},   // Sequence details
			{Segment: "IND", Mandatory: true, MaxRepeat: 1},   // Index details
			{Segment: "RFF", Mandatory: false, MaxRepeat: 9},  // Reference
			{Segment: "MOA", Mandatory: false, MaxRepeat: 99}, // Monetary amount
			{Segment: "FTX", Mandatory: false, MaxRepeat: 99}, // Free text
			{Segment: "DTM", Mandatory: false, MaxRepeat: 9},  // Date/time/period
			{Segment: "QTY", Mandatory: false, MaxRepeat: 9},  // Quantity
			{Segment: "CUX", Mandatory: false, MaxRepeat: 9},  // Currencies
			{Segment: "PCD", Mandatory: false, MaxRepeat: 9},  // Percentage details
			{Segment: "FII", Mandatory: false, MaxRepeat: 9},  // Financial institution information
			{Segment: "CPT", Mandatory: false, MaxRepeat: 9},  // Account identification
			{ // Segment group 5
				Group: []SchemaNode{
					{Segment: "NAD", Mandatory: true, MaxRepeat: 1}, // Name and address
					{ // Segment group 6
						Group: []SchemaNode{
							{Segment: "CTA", Mandatory: true, MaxRepeat: 1},  // Contact information
							{Segment: "COM", Mandatory: false, MaxRepeat: 9}, // Communication contact
						},
						Mandatory: false, MaxRepeat: 99,
					},
				},
				Mandatory: false, MaxRepeat: 9,
			},
			{ // Segment group 7
				Group: []SchemaNode{
					{Segment: "CCI", Mandatory: true, MaxRepeat: 1},  // Characteristic/class id
					{Segment: "CAV", Mandatory: false, MaxRepeat: 1}, // Characteristic value
				},
				Mandatory: false, MaxRepeat: 99,
			},
		},
		Mandatory: true, MaxRepeat: 99999,
	},
	{ // Segment group 8
		Group: []SchemaNode{
			{Segment: "AUT", Mandatory: true, MaxRepeat: 1},  // Authentication result
			{Segment: "DTM", Mandatory: false, MaxRepeat: 1}, // Date/time/period
		},
		Mandatory: false, MaxRepeat: 1,
	},
}}
