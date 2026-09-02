package edifact

// CUSREP D.20A message specification: the structural (segment/group
// presence, order, cardinality) branching diagram for this message,
// UN/EDIFACT directory release D.20A. Scope is structural only -- not
// element-level content -- per edifact-ls-3uzr's epic non-goal.
// 12 segment groups, max nesting depth 2.
//
// Source: https://service.unece.org/trade/untdid/d20a/trmd/cusrep_c.htm
// section 4.3.1 "Segment table". That URL currently returns HTTP 403 from
// Cloudflare when fetched directly; this data was transcribed from the
// Wayback Machine's archived copy instead:
// http://web.archive.org/web/20231202090638/https://service.unece.org/trade/untdid/d20a/trmd/cusrep_c.htm
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
		MessageID{Type: "CUSREP", Version: "D", Release: "20A", Agency: "UN"},
		cusrepD20aSchema,
		"https://service.unece.org/trade/untdid/d20a/trmd/cusrep_c.htm",
	)
}

var cusrepD20aSchema = Schema{Nodes: []SchemaNode{
	{Segment: "BGM", Mandatory: true, MaxRepeat: 1},   // Beginning of message
	{Segment: "DTM", Mandatory: false, MaxRepeat: 9},  // Date/time/period
	{Segment: "QTY", Mandatory: false, MaxRepeat: 9},  // Quantity
	{Segment: "POC", Mandatory: false, MaxRepeat: 99}, // Purpose of conveyance call
	{Segment: "FTX", Mandatory: false, MaxRepeat: 9},  // Free text
	{Segment: "MEA", Mandatory: false, MaxRepeat: 9},  // Measurements
	{Segment: "GEI", Mandatory: false, MaxRepeat: 9},  // Processing information
	{Segment: "GPO", Mandatory: false, MaxRepeat: 1},  // Geographical position
	{Segment: "STS", Mandatory: false, MaxRepeat: 9},  // Status
	{ // Segment group 1
		Group: []SchemaNode{
			{Segment: "RFF", Mandatory: true, MaxRepeat: 1},  // Reference
			{Segment: "DTM", Mandatory: false, MaxRepeat: 1}, // Date/time/period
		},
		Mandatory: false, MaxRepeat: 99,
	},
	{ // Segment group 2
		Group: []SchemaNode{
			{Segment: "GDS", Mandatory: true, MaxRepeat: 1},  // Nature of cargo
			{Segment: "FTX", Mandatory: false, MaxRepeat: 9}, // Free text
		},
		Mandatory: false, MaxRepeat: 9,
	},
	{ // Segment group 3
		Group: []SchemaNode{
			{Segment: "LOC", Mandatory: true, MaxRepeat: 1},  // Place/location identification
			{Segment: "DTM", Mandatory: false, MaxRepeat: 9}, // Date/time/period
		},
		Mandatory: false, MaxRepeat: 99,
	},
	{ // Segment group 4
		Group: []SchemaNode{
			{Segment: "DOC", Mandatory: true, MaxRepeat: 1},  // Document/message details
			{Segment: "RFF", Mandatory: false, MaxRepeat: 1}, // Reference
			{Segment: "DTM", Mandatory: false, MaxRepeat: 1}, // Date/time/period
			{Segment: "LOC", Mandatory: false, MaxRepeat: 1}, // Place/location identification
		},
		Mandatory: false, MaxRepeat: 9,
	},
	{ // Segment group 5
		Group: []SchemaNode{
			{Segment: "TAX", Mandatory: true, MaxRepeat: 1},  // Duty/tax/fee details
			{Segment: "MOA", Mandatory: false, MaxRepeat: 1}, // Monetary amount
			{Segment: "FII", Mandatory: false, MaxRepeat: 1}, // Financial institution information
			{Segment: "LOC", Mandatory: false, MaxRepeat: 1}, // Place/location identification
			{Segment: "RFF", Mandatory: false, MaxRepeat: 1}, // Reference
			{Segment: "DTM", Mandatory: false, MaxRepeat: 9}, // Date/time/period
			{Segment: "GEI", Mandatory: false, MaxRepeat: 1}, // Processing information
		},
		Mandatory: false, MaxRepeat: 9,
	},
	{ // Segment group 6
		Group: []SchemaNode{
			{Segment: "NAD", Mandatory: true, MaxRepeat: 1}, // Name and address
			{ // Segment group 7
				Group: []SchemaNode{
					{Segment: "CTA", Mandatory: true, MaxRepeat: 1},  // Contact information
					{Segment: "COM", Mandatory: false, MaxRepeat: 5}, // Communication contact
				},
				Mandatory: false, MaxRepeat: 9,
			},
			{ // Segment group 8
				Group: []SchemaNode{
					{Segment: "RFF", Mandatory: true, MaxRepeat: 1},  // Reference
					{Segment: "DTM", Mandatory: false, MaxRepeat: 1}, // Date/time/period
				},
				Mandatory: false, MaxRepeat: 9,
			},
		},
		Mandatory: false, MaxRepeat: 9,
	},
	{ // Segment group 9
		Group: []SchemaNode{
			{Segment: "TDT", Mandatory: true, MaxRepeat: 1},  // Transport information
			{Segment: "TPL", Mandatory: false, MaxRepeat: 1}, // Transport placement
			{Segment: "DTM", Mandatory: false, MaxRepeat: 1}, // Date/time/period
			{ // Segment group 10
				Group: []SchemaNode{
					{Segment: "LOC", Mandatory: true, MaxRepeat: 1},   // Place/location identification
					{Segment: "GPO", Mandatory: false, MaxRepeat: 1},  // Geographical position
					{Segment: "DTM", Mandatory: false, MaxRepeat: 9},  // Date/time/period
					{Segment: "QTY", Mandatory: false, MaxRepeat: 9},  // Quantity
					{Segment: "NAD", Mandatory: false, MaxRepeat: 99}, // Name and address
					{Segment: "MEA", Mandatory: false, MaxRepeat: 9},  // Measurements
					{Segment: "POC", Mandatory: false, MaxRepeat: 9},  // Purpose of conveyance call
					{Segment: "STS", Mandatory: false, MaxRepeat: 9},  // Status
					{Segment: "FTX", Mandatory: false, MaxRepeat: 9},  // Free text
				},
				Mandatory: false, MaxRepeat: 99,
			},
		},
		Mandatory: false, MaxRepeat: 99,
	},
	{ // Segment group 11
		Group: []SchemaNode{
			{Segment: "EQD", Mandatory: true, MaxRepeat: 1},  // Equipment details
			{Segment: "EQN", Mandatory: false, MaxRepeat: 1}, // Number of units
		},
		Mandatory: false, MaxRepeat: 999,
	},
	{ // Segment group 12
		Group: []SchemaNode{
			{Segment: "AUT", Mandatory: true, MaxRepeat: 1},  // Authentication result
			{Segment: "DTM", Mandatory: false, MaxRepeat: 1}, // Date/time/period
		},
		Mandatory: false, MaxRepeat: 1,
	},
}}
