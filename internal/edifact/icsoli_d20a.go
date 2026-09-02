package edifact

// ICSOLI D.20A message specification: the structural (segment/group
// presence, order, cardinality) branching diagram for this message,
// UN/EDIFACT directory release D.20A. Scope is structural only -- not
// element-level content -- per edifact-ls-3uzr's epic non-goal.
// 10 segment groups, max nesting depth 2.
//
// Source: https://service.unece.org/trade/untdid/d20a/trmd/icsoli_c.htm
// section 4.3.1 "Segment table". That URL currently returns HTTP 403 from
// Cloudflare when fetched directly; this data was transcribed from the
// Wayback Machine's archived copy instead:
// http://web.archive.org/web/20230608115320/https://service.unece.org/trade/untdid/d20a/trmd/icsoli_c.htm
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
		MessageID{Type: "ICSOLI", Version: "D", Release: "20A", Agency: "UN"},
		icsoliD20aSchema,
		"https://service.unece.org/trade/untdid/d20a/trmd/icsoli_c.htm",
	)
}

var icsoliD20aSchema = Schema{Nodes: []SchemaNode{
	{Segment: "BGM", Mandatory: true, MaxRepeat: 1},  // Beginning of message
	{Segment: "RFF", Mandatory: false, MaxRepeat: 9}, // Reference
	{Segment: "SEQ", Mandatory: false, MaxRepeat: 1}, // Sequence details
	{Segment: "DTM", Mandatory: false, MaxRepeat: 9}, // Date/time/period
	{Segment: "EFI", Mandatory: false, MaxRepeat: 9}, // External file link identification
	{ // Segment group 1
		Group: []SchemaNode{
			{Segment: "ICD", Mandatory: true, MaxRepeat: 1},  // Insurance cover description
			{Segment: "MOA", Mandatory: false, MaxRepeat: 9}, // Monetary amount
			{Segment: "ATT", Mandatory: false, MaxRepeat: 9}, // Attribute
			{Segment: "RFF", Mandatory: false, MaxRepeat: 9}, // Reference
			{Segment: "DTM", Mandatory: false, MaxRepeat: 9}, // Date/time/period
			{Segment: "QRS", Mandatory: false, MaxRepeat: 9}, // Query and response
			{Segment: "FTX", Mandatory: false, MaxRepeat: 9}, // Free text
		},
		Mandatory: true, MaxRepeat: 1,
	},
	{ // Segment group 2
		Group: []SchemaNode{
			{Segment: "PNA", Mandatory: true, MaxRepeat: 1},   // Party identification
			{Segment: "RFF", Mandatory: false, MaxRepeat: 99}, // Reference
			{Segment: "EMP", Mandatory: false, MaxRepeat: 9},  // Employment details
			{Segment: "DTM", Mandatory: false, MaxRepeat: 9},  // Date/time/period
			{Segment: "PCD", Mandatory: false, MaxRepeat: 9},  // Percentage details
			{Segment: "QRS", Mandatory: false, MaxRepeat: 9},  // Query and response
			{Segment: "FTX", Mandatory: false, MaxRepeat: 1},  // Free text
			{Segment: "COM", Mandatory: false, MaxRepeat: 9},  // Communication contact
			{Segment: "ATT", Mandatory: false, MaxRepeat: 9},  // Attribute
			{ // Segment group 3
				Group: []SchemaNode{
					{Segment: "ADR", Mandatory: true, MaxRepeat: 1},  // Address
					{Segment: "COM", Mandatory: false, MaxRepeat: 9}, // Communication contact
				},
				Mandatory: false, MaxRepeat: 9,
			},
			{ // Segment group 4
				Group: []SchemaNode{
					{Segment: "ROD", Mandatory: true, MaxRepeat: 1},  // Risk object type
					{Segment: "MOA", Mandatory: true, MaxRepeat: 1},  // Monetary amount
					{Segment: "FTX", Mandatory: false, MaxRepeat: 1}, // Free text
				},
				Mandatory: false, MaxRepeat: 9,
			},
			{ // Segment group 5
				Group: []SchemaNode{
					{Segment: "COD", Mandatory: true, MaxRepeat: 1},  // Component details
					{Segment: "ATT", Mandatory: false, MaxRepeat: 9}, // Attribute
					{Segment: "IDE", Mandatory: false, MaxRepeat: 1}, // Identity
					{Segment: "RFF", Mandatory: false, MaxRepeat: 9}, // Reference
					{Segment: "DTM", Mandatory: false, MaxRepeat: 9}, // Date/time/period
					{Segment: "MOA", Mandatory: false, MaxRepeat: 9}, // Monetary amount
					{Segment: "EQN", Mandatory: false, MaxRepeat: 9}, // Number of units
					{Segment: "QRS", Mandatory: false, MaxRepeat: 9}, // Query and response
				},
				Mandatory: false, MaxRepeat: 99,
			},
		},
		Mandatory: true, MaxRepeat: 9,
	},
	{Segment: "UNS", Mandatory: true, MaxRepeat: 1}, // Section control
	{ // Segment group 6
		Group: []SchemaNode{
			{Segment: "EVE", Mandatory: true, MaxRepeat: 1},  // Event
			{Segment: "ADR", Mandatory: false, MaxRepeat: 9}, // Address
			{Segment: "RFF", Mandatory: false, MaxRepeat: 9}, // Reference
			{Segment: "DTM", Mandatory: false, MaxRepeat: 9}, // Date/time/period
			{Segment: "QRS", Mandatory: false, MaxRepeat: 9}, // Query and response
			{Segment: "FTX", Mandatory: false, MaxRepeat: 9}, // Free text
			{ // Segment group 7
				Group: []SchemaNode{
					{Segment: "ATT", Mandatory: true, MaxRepeat: 1},  // Attribute
					{Segment: "MOA", Mandatory: false, MaxRepeat: 1}, // Monetary amount
				},
				Mandatory: false, MaxRepeat: 9,
			},
		},
		Mandatory: false, MaxRepeat: 9,
	},
	{ // Segment group 8
		Group: []SchemaNode{
			{Segment: "PYT", Mandatory: true, MaxRepeat: 1},  // Payment terms
			{Segment: "MOA", Mandatory: true, MaxRepeat: 9},  // Monetary amount
			{Segment: "PNA", Mandatory: false, MaxRepeat: 1}, // Party identification
			{Segment: "FII", Mandatory: false, MaxRepeat: 1}, // Financial institution information
			{Segment: "ATT", Mandatory: false, MaxRepeat: 9}, // Attribute
			{Segment: "RFF", Mandatory: false, MaxRepeat: 9}, // Reference
		},
		Mandatory: false, MaxRepeat: 99,
	},
	{ // Segment group 9
		Group: []SchemaNode{
			{Segment: "COD", Mandatory: true, MaxRepeat: 1},  // Component details
			{Segment: "ATT", Mandatory: false, MaxRepeat: 9}, // Attribute
			{Segment: "RFF", Mandatory: false, MaxRepeat: 9}, // Reference
			{Segment: "ADR", Mandatory: false, MaxRepeat: 9}, // Address
			{Segment: "DTM", Mandatory: false, MaxRepeat: 9}, // Date/time/period
			{Segment: "QRS", Mandatory: false, MaxRepeat: 9}, // Query and response
			{Segment: "FTX", Mandatory: false, MaxRepeat: 9}, // Free text
			{ // Segment group 10
				Group: []SchemaNode{
					{Segment: "IMD", Mandatory: true, MaxRepeat: 1},  // Item description
					{Segment: "ATT", Mandatory: false, MaxRepeat: 9}, // Attribute
					{Segment: "RFF", Mandatory: false, MaxRepeat: 9}, // Reference
					{Segment: "DTM", Mandatory: false, MaxRepeat: 9}, // Date/time/period
					{Segment: "MOA", Mandatory: false, MaxRepeat: 9}, // Monetary amount
					{Segment: "FTX", Mandatory: false, MaxRepeat: 9}, // Free text
				},
				Mandatory: false, MaxRepeat: 99,
			},
		},
		Mandatory: false, MaxRepeat: 99,
	},
}}
