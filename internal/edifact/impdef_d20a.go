package edifact

// IMPDEF D.20A message specification: the structural (segment/group
// presence, order, cardinality) branching diagram for this message,
// UN/EDIFACT directory release D.20A. Scope is structural only -- not
// element-level content -- per edifact-ls-3uzr's epic non-goal.
// 13 segment groups, max nesting depth 2.
//
// Source: https://service.unece.org/trade/untdid/d20a/trmd/impdef_c.htm
// section 4.3.1 "Segment table". That URL currently returns HTTP 403 from
// Cloudflare when fetched directly; this data was transcribed from the
// Wayback Machine's archived copy instead:
// http://web.archive.org/web/20240416200647/https://service.unece.org/trade/untdid/d20a/trmd/impdef_c.htm
// -- re-check the direct URL first if cross-checking later, in case the
// block has lifted.
//
// Transcription note: generated in bulk (edifact-ls-13gu) the same way
// as every other message-type schema -- a script parsed the source's
// exact rail-art column positions mechanically and verified the result
// balances (all 13 groups open and close consistently)
// before being transcribed here, rather than reading the ASCII nesting
// by eye. Not individually spot-checked against the raw source the way
// the first several message types in this project were -- see
// edifact-ls-13gu for which representative samples were spot-checked
// for this batch.

func init() {
	RegisterSchema(
		MessageID{Type: "IMPDEF", Version: "D", Release: "20A", Agency: "UN"},
		impdefD20aSchema,
		"https://service.unece.org/trade/untdid/d20a/trmd/impdef_c.htm",
	)
}

var impdefD20aSchema = Schema{Nodes: []SchemaNode{
	{Segment: "BGM", Mandatory: true, MaxRepeat: 1},    // Beginning of message
	{Segment: "MSG", Mandatory: true, MaxRepeat: 1},    // Message type identification
	{Segment: "RCS", Mandatory: false, MaxRepeat: 1},   // Requirements and conditions
	{Segment: "DII", Mandatory: true, MaxRepeat: 1},    // Directory identification
	{Segment: "RFF", Mandatory: false, MaxRepeat: 99},  // Reference
	{Segment: "DTM", Mandatory: false, MaxRepeat: 9},   // Date/time/period
	{Segment: "FTX", Mandatory: false, MaxRepeat: 999}, // Free text
	{ // Segment group 1
		Group: []SchemaNode{
			{Segment: "PNA", Mandatory: true, MaxRepeat: 1},  // Party identification
			{Segment: "ADR", Mandatory: false, MaxRepeat: 1}, // Address
			{ // Segment group 2
				Group: []SchemaNode{
					{Segment: "CTA", Mandatory: true, MaxRepeat: 1},  // Contact information
					{Segment: "COM", Mandatory: false, MaxRepeat: 5}, // Communication contact
				},
				Mandatory: false, MaxRepeat: 9999,
			},
		},
		Mandatory: false, MaxRepeat: 5,
	},
	{ // Segment group 3
		Group: []SchemaNode{
			{Segment: "DFN", Mandatory: true, MaxRepeat: 1},   // Definition function
			{Segment: "FTX", Mandatory: false, MaxRepeat: 99}, // Free text
			{ // Segment group 4
				Group: []SchemaNode{
					{Segment: "GRU", Mandatory: true, MaxRepeat: 1},   // Segment group usage details
					{Segment: "FTX", Mandatory: false, MaxRepeat: 99}, // Free text
				},
				Mandatory: false, MaxRepeat: 1,
			},
			{ // Segment group 5
				Group: []SchemaNode{
					{Segment: "SGU", Mandatory: true, MaxRepeat: 1},   // Segment usage details
					{Segment: "FTX", Mandatory: false, MaxRepeat: 99}, // Free text
				},
				Mandatory: false, MaxRepeat: 999,
			},
			{ // Segment group 6
				Group: []SchemaNode{
					{Segment: "FNT", Mandatory: true, MaxRepeat: 1},   // Footnote
					{Segment: "REL", Mandatory: false, MaxRepeat: 1},  // Relationship
					{Segment: "GIR", Mandatory: false, MaxRepeat: 9},  // Related identification numbers
					{Segment: "FTX", Mandatory: false, MaxRepeat: 99}, // Free text
				},
				Mandatory: false, MaxRepeat: 99,
			},
			{ // Segment group 7
				Group: []SchemaNode{
					{Segment: "RFF", Mandatory: true, MaxRepeat: 1},   // Reference
					{Segment: "FTX", Mandatory: false, MaxRepeat: 99}, // Free text
				},
				Mandatory: false, MaxRepeat: 99,
			},
			{ // Segment group 8
				Group: []SchemaNode{
					{Segment: "ELU", Mandatory: true, MaxRepeat: 1},   // Data element usage details
					{Segment: "ELM", Mandatory: false, MaxRepeat: 1},  // Simple data element details
					{Segment: "EDT", Mandatory: false, MaxRepeat: 1},  // Editing details
					{Segment: "IMD", Mandatory: false, MaxRepeat: 9},  // Item description
					{Segment: "GEI", Mandatory: false, MaxRepeat: 9},  // Processing information
					{Segment: "FTX", Mandatory: false, MaxRepeat: 99}, // Free text
				},
				Mandatory: false, MaxRepeat: 99,
			},
			{ // Segment group 9
				Group: []SchemaNode{
					{Segment: "MEA", Mandatory: true, MaxRepeat: 1},   // Measurements
					{Segment: "FTX", Mandatory: false, MaxRepeat: 99}, // Free text
				},
				Mandatory: false, MaxRepeat: 999,
			},
			{ // Segment group 10
				Group: []SchemaNode{
					{Segment: "ELV", Mandatory: true, MaxRepeat: 1},   // Element value definition
					{Segment: "FTX", Mandatory: false, MaxRepeat: 99}, // Free text
				},
				Mandatory: false, MaxRepeat: 99999,
			},
			{ // Segment group 11
				Group: []SchemaNode{
					{Segment: "CDV", Mandatory: true, MaxRepeat: 1},   // Code value definition
					{Segment: "FTX", Mandatory: false, MaxRepeat: 99}, // Free text
				},
				Mandatory: false, MaxRepeat: 99999,
			},
			{ // Segment group 12
				Group: []SchemaNode{
					{Segment: "DRD", Mandatory: true, MaxRepeat: 1},   // Data representation details
					{Segment: "FTX", Mandatory: false, MaxRepeat: 99}, // Free text
				},
				Mandatory: false, MaxRepeat: 99999,
			},
		},
		Mandatory: false, MaxRepeat: 99999,
	},
	{ // Segment group 13
		Group: []SchemaNode{
			{Segment: "AUT", Mandatory: true, MaxRepeat: 1},  // Authentication result
			{Segment: "DTM", Mandatory: false, MaxRepeat: 1}, // Date/time/period
		},
		Mandatory: false, MaxRepeat: 1,
	},
}}
