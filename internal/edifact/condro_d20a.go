package edifact

// CONDRO D.20A message specification: the structural (segment/group
// presence, order, cardinality) branching diagram for this message,
// UN/EDIFACT directory release D.20A. Scope is structural only -- not
// element-level content -- per edifact-ls-3uzr's epic non-goal.
// 8 segment groups, max nesting depth 2.
//
// Source: https://service.unece.org/trade/untdid/d20a/trmd/condro_c.htm
// section 4.3.1 "Segment table". That URL currently returns HTTP 403 from
// Cloudflare when fetched directly; this data was transcribed from the
// Wayback Machine's archived copy instead:
// http://web.archive.org/web/20240416110932/https://service.unece.org/trade/untdid/d20a/trmd/condro_c.htm
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
		MessageID{Type: "CONDRO", Version: "D", Release: "20A", Agency: "UN"},
		condroD20aSchema,
		"https://service.unece.org/trade/untdid/d20a/trmd/condro_c.htm",
	)
}

var condroD20aSchema = Schema{Nodes: []SchemaNode{
	{Segment: "BGM", Mandatory: true, MaxRepeat: 1},   // Beginning of message
	{Segment: "DTM", Mandatory: true, MaxRepeat: 5},   // Date/time/period
	{Segment: "AUT", Mandatory: false, MaxRepeat: 2},  // Authentication result
	{Segment: "AGR", Mandatory: false, MaxRepeat: 2},  // Agreement identification
	{Segment: "FTX", Mandatory: false, MaxRepeat: 10}, // Free text
	{ // Segment group 1
		Group: []SchemaNode{
			{Segment: "RFF", Mandatory: true, MaxRepeat: 1},  // Reference
			{Segment: "DOC", Mandatory: false, MaxRepeat: 1}, // Document/message details
			{Segment: "DTM", Mandatory: false, MaxRepeat: 1}, // Date/time/period
			{Segment: "FTX", Mandatory: false, MaxRepeat: 5}, // Free text
		},
		Mandatory: true, MaxRepeat: 5,
	},
	{ // Segment group 2
		Group: []SchemaNode{
			{Segment: "NAD", Mandatory: true, MaxRepeat: 1},   // Name and address
			{Segment: "LOC", Mandatory: false, MaxRepeat: 99}, // Place/location identification
			{Segment: "FTX", Mandatory: false, MaxRepeat: 10}, // Free text
			{ // Segment group 3
				Group: []SchemaNode{
					{Segment: "RFF", Mandatory: true, MaxRepeat: 1},   // Reference
					{Segment: "CED", Mandatory: false, MaxRepeat: 10}, // Computer environment details
					{Segment: "DTM", Mandatory: false, MaxRepeat: 1},  // Date/time/period
				},
				Mandatory: false, MaxRepeat: 10,
			},
			{ // Segment group 4
				Group: []SchemaNode{
					{Segment: "CTA", Mandatory: true, MaxRepeat: 1},  // Contact information
					{Segment: "COM", Mandatory: false, MaxRepeat: 5}, // Communication contact
					{Segment: "LOC", Mandatory: false, MaxRepeat: 1}, // Place/location identification
				},
				Mandatory: false, MaxRepeat: 10,
			},
		},
		Mandatory: true, MaxRepeat: 999,
	},
	{ // Segment group 5
		Group: []SchemaNode{
			{Segment: "DOC", Mandatory: true, MaxRepeat: 1},   // Document/message details
			{Segment: "INP", Mandatory: false, MaxRepeat: 99}, // Parties and instruction
			{Segment: "RFF", Mandatory: false, MaxRepeat: 1},  // Reference
			{Segment: "MEA", Mandatory: false, MaxRepeat: 1},  // Measurements
			{ // Segment group 6
				Group: []SchemaNode{
					{Segment: "EFI", Mandatory: true, MaxRepeat: 1},   // External file link identification
					{Segment: "CED", Mandatory: false, MaxRepeat: 10}, // Computer environment details
					{Segment: "DTM", Mandatory: false, MaxRepeat: 5},  // Date/time/period
				},
				Mandatory: false, MaxRepeat: 99,
			},
		},
		Mandatory: false, MaxRepeat: 999,
	},
	{ // Segment group 7
		Group: []SchemaNode{
			{Segment: "BII", Mandatory: true, MaxRepeat: 1},  // Structure identification
			{Segment: "GEI", Mandatory: false, MaxRepeat: 1}, // Processing information
			{Segment: "DTM", Mandatory: false, MaxRepeat: 1}, // Date/time/period
			{Segment: "IMD", Mandatory: false, MaxRepeat: 1}, // Item description
			{Segment: "RFF", Mandatory: false, MaxRepeat: 5}, // Reference
			{Segment: "QTY", Mandatory: false, MaxRepeat: 1}, // Quantity
			{ // Segment group 8
				Group: []SchemaNode{
					{Segment: "LOC", Mandatory: true, MaxRepeat: 1},  // Place/location identification
					{Segment: "DIM", Mandatory: false, MaxRepeat: 5}, // Dimensions
					{Segment: "QTY", Mandatory: false, MaxRepeat: 5}, // Quantity
					{Segment: "MEA", Mandatory: false, MaxRepeat: 1}, // Measurements
				},
				Mandatory: false, MaxRepeat: 1,
			},
		},
		Mandatory: false, MaxRepeat: 100000,
	},
}}
