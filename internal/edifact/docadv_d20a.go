package edifact

// DOCADV D.20A message specification: the structural (segment/group
// presence, order, cardinality) branching diagram for this message,
// UN/EDIFACT directory release D.20A. Scope is structural only -- not
// element-level content -- per edifact-ls-3uzr's epic non-goal.
// 17 segment groups, max nesting depth 3.
//
// Source: https://service.unece.org/trade/untdid/d20a/trmd/docadv_c.htm
// section 4.3.1 "Segment table". That URL currently returns HTTP 403 from
// Cloudflare when fetched directly; this data was transcribed from the
// Wayback Machine's archived copy instead:
// http://web.archive.org/web/20230608021331/https://service.unece.org/trade/untdid/d20a/trmd/docadv_c.htm
// -- re-check the direct URL first if cross-checking later, in case the
// block has lifted.
//
// Transcription note: generated in bulk (edifact-ls-13gu) the same way
// as every other message-type schema -- a script parsed the source's
// exact rail-art column positions mechanically and verified the result
// balances (all 17 groups open and close consistently)
// before being transcribed here, rather than reading the ASCII nesting
// by eye. Not individually spot-checked against the raw source the way
// the first several message types in this project were -- see
// edifact-ls-13gu for which representative samples were spot-checked
// for this batch.

func init() {
	RegisterSchema(
		MessageID{Type: "DOCADV", Version: "D", Release: "20A", Agency: "UN"},
		docadvD20aSchema,
		"https://service.unece.org/trade/untdid/d20a/trmd/docadv_c.htm",
	)
}

var docadvD20aSchema = Schema{Nodes: []SchemaNode{
	{Segment: "BGM", Mandatory: true, MaxRepeat: 1},   // Beginning of message
	{Segment: "RFF", Mandatory: true, MaxRepeat: 1},   // Reference
	{Segment: "BUS", Mandatory: true, MaxRepeat: 1},   // Business function
	{Segment: "INP", Mandatory: true, MaxRepeat: 10},  // Parties and instruction
	{Segment: "FCA", Mandatory: true, MaxRepeat: 3},   // Financial charges allocation
	{Segment: "DTM", Mandatory: true, MaxRepeat: 3},   // Date/time/period
	{Segment: "FTX", Mandatory: false, MaxRepeat: 20}, // Free text
	{ // Segment group 1
		Group: []SchemaNode{
			{Segment: "FII", Mandatory: true, MaxRepeat: 1},  // Financial institution information
			{Segment: "RFF", Mandatory: false, MaxRepeat: 2}, // Reference
			{Segment: "CTA", Mandatory: false, MaxRepeat: 1}, // Contact information
			{Segment: "COM", Mandatory: false, MaxRepeat: 5}, // Communication contact
		},
		Mandatory: true, MaxRepeat: 9,
	},
	{ // Segment group 2
		Group: []SchemaNode{
			{Segment: "NAD", Mandatory: true, MaxRepeat: 1},  // Name and address
			{Segment: "RFF", Mandatory: false, MaxRepeat: 1}, // Reference
			{Segment: "CTA", Mandatory: false, MaxRepeat: 1}, // Contact information
			{Segment: "COM", Mandatory: false, MaxRepeat: 5}, // Communication contact
		},
		Mandatory: true, MaxRepeat: 9,
	},
	{ // Segment group 3
		Group: []SchemaNode{
			{Segment: "DTM", Mandatory: true, MaxRepeat: 1}, // Date/time/period
			{Segment: "LOC", Mandatory: true, MaxRepeat: 1}, // Place/location identification
		},
		Mandatory: true, MaxRepeat: 1,
	},
	{ // Segment group 4
		Group: []SchemaNode{
			{Segment: "MOA", Mandatory: true, MaxRepeat: 1}, // Monetary amount
			{ // Segment group 5
				Group: []SchemaNode{
					{Segment: "ALC", Mandatory: true, MaxRepeat: 1},  // Allowance or charge
					{Segment: "PCD", Mandatory: false, MaxRepeat: 2}, // Percentage details
				},
				Mandatory: false, MaxRepeat: 1,
			},
		},
		Mandatory: true, MaxRepeat: 5,
	},
	{ // Segment group 6
		Group: []SchemaNode{
			{Segment: "LOC", Mandatory: true, MaxRepeat: 1},  // Place/location identification
			{Segment: "DTM", Mandatory: false, MaxRepeat: 1}, // Date/time/period
		},
		Mandatory: true, MaxRepeat: 3,
	},
	{ // Segment group 7
		Group: []SchemaNode{
			{Segment: "PAI", Mandatory: true, MaxRepeat: 1},  // Payment instructions
			{Segment: "FII", Mandatory: false, MaxRepeat: 1}, // Financial institution information
			{Segment: "LOC", Mandatory: false, MaxRepeat: 1}, // Place/location identification
		},
		Mandatory: true, MaxRepeat: 1,
	},
	{ // Segment group 8
		Group: []SchemaNode{
			{Segment: "PYT", Mandatory: true, MaxRepeat: 1},  // Payment terms
			{Segment: "FII", Mandatory: false, MaxRepeat: 1}, // Financial institution information
			{Segment: "DTM", Mandatory: false, MaxRepeat: 1}, // Date/time/period
			{Segment: "MOA", Mandatory: false, MaxRepeat: 1}, // Monetary amount
			{Segment: "PCD", Mandatory: false, MaxRepeat: 1}, // Percentage details
			{Segment: "FTX", Mandatory: false, MaxRepeat: 1}, // Free text
		},
		Mandatory: true, MaxRepeat: 5,
	},
	{ // Segment group 9
		Group: []SchemaNode{
			{Segment: "TOD", Mandatory: true, MaxRepeat: 1},  // Terms of delivery or transport
			{Segment: "LOC", Mandatory: false, MaxRepeat: 1}, // Place/location identification
		},
		Mandatory: false, MaxRepeat: 1,
	},
	{ // Segment group 10
		Group: []SchemaNode{
			{Segment: "TSR", Mandatory: true, MaxRepeat: 1},  // Transport service requirements
			{Segment: "LOC", Mandatory: false, MaxRepeat: 5}, // Place/location identification
		},
		Mandatory: false, MaxRepeat: 1,
	},
	{ // Segment group 11
		Group: []SchemaNode{
			{Segment: "INP", Mandatory: true, MaxRepeat: 1},  // Parties and instruction
			{Segment: "FTX", Mandatory: false, MaxRepeat: 1}, // Free text
			{Segment: "DTM", Mandatory: false, MaxRepeat: 2}, // Date/time/period
		},
		Mandatory: false, MaxRepeat: 5,
	},
	{ // Segment group 12
		Group: []SchemaNode{
			{Segment: "RFF", Mandatory: true, MaxRepeat: 1},  // Reference
			{Segment: "DTM", Mandatory: false, MaxRepeat: 1}, // Date/time/period
			{Segment: "FTX", Mandatory: false, MaxRepeat: 2}, // Free text
		},
		Mandatory: true, MaxRepeat: 9,
	},
	{ // Segment group 13
		Group: []SchemaNode{
			{Segment: "DOC", Mandatory: true, MaxRepeat: 1},  // Document/message details
			{Segment: "MOA", Mandatory: false, MaxRepeat: 1}, // Monetary amount
			{Segment: "PCD", Mandatory: false, MaxRepeat: 1}, // Percentage details
			{Segment: "LOC", Mandatory: false, MaxRepeat: 1}, // Place/location identification
			{Segment: "FTX", Mandatory: false, MaxRepeat: 1}, // Free text
			{ // Segment group 14
				Group: []SchemaNode{
					{Segment: "ICD", Mandatory: true, MaxRepeat: 1},  // Insurance cover description
					{Segment: "DTM", Mandatory: false, MaxRepeat: 1}, // Date/time/period
					{Segment: "FTX", Mandatory: false, MaxRepeat: 9}, // Free text
				},
				Mandatory: false, MaxRepeat: 20,
			},
			{ // Segment group 15
				Group: []SchemaNode{
					{Segment: "ALI", Mandatory: true, MaxRepeat: 1}, // Additional information
					{ // Segment group 16
						Group: []SchemaNode{
							{Segment: "NAD", Mandatory: true, MaxRepeat: 1},  // Name and address
							{Segment: "CTA", Mandatory: false, MaxRepeat: 1}, // Contact information
							{Segment: "COM", Mandatory: false, MaxRepeat: 5}, // Communication contact
						},
						Mandatory: false, MaxRepeat: 3,
					},
				},
				Mandatory: false, MaxRepeat: 9,
			},
		},
		Mandatory: true, MaxRepeat: 20,
	},
	{ // Segment group 17
		Group: []SchemaNode{
			{Segment: "AUT", Mandatory: true, MaxRepeat: 1},  // Authentication result
			{Segment: "DTM", Mandatory: false, MaxRepeat: 1}, // Date/time/period
		},
		Mandatory: false, MaxRepeat: 1,
	},
}}
