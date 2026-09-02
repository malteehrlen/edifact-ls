package edifact

// DIRDEF D.20A message specification: the structural (segment/group
// presence, order, cardinality) branching diagram for this message,
// UN/EDIFACT directory release D.20A. Scope is structural only -- not
// element-level content -- per edifact-ls-3uzr's epic non-goal.
// 13 segment groups, max nesting depth 3.
//
// Source: https://service.unece.org/trade/untdid/d20a/trmd/dirdef_c.htm
// section 4.3.1 "Segment table". That URL currently returns HTTP 403 from
// Cloudflare when fetched directly; this data was transcribed from the
// Wayback Machine's archived copy instead:
// http://web.archive.org/web/20240416212139/https://service.unece.org/trade/untdid/d20a/trmd/dirdef_c.htm
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
		MessageID{Type: "DIRDEF", Version: "D", Release: "20A", Agency: "UN"},
		dirdefD20aSchema,
		"https://service.unece.org/trade/untdid/d20a/trmd/dirdef_c.htm",
	)
}

var dirdefD20aSchema = Schema{Nodes: []SchemaNode{
	{Segment: "BGM", Mandatory: true, MaxRepeat: 1},  // Beginning of message
	{Segment: "DII", Mandatory: true, MaxRepeat: 1},  // Directory identification
	{Segment: "DTM", Mandatory: false, MaxRepeat: 9}, // Date/time/period
	{Segment: "FTX", Mandatory: false, MaxRepeat: 9}, // Free text
	{ // Segment group 1
		Group: []SchemaNode{
			{Segment: "NAD", Mandatory: true, MaxRepeat: 1}, // Name and address
			{ // Segment group 2
				Group: []SchemaNode{
					{Segment: "CTA", Mandatory: true, MaxRepeat: 1},  // Contact information
					{Segment: "COM", Mandatory: false, MaxRepeat: 9}, // Communication contact
				},
				Mandatory: false, MaxRepeat: 9,
			},
		},
		Mandatory: false, MaxRepeat: 9,
	},
	{ // Segment group 3
		Group: []SchemaNode{
			{Segment: "MSG", Mandatory: true, MaxRepeat: 1},    // Message type identification
			{Segment: "ATT", Mandatory: false, MaxRepeat: 99},  // Attribute
			{Segment: "FTX", Mandatory: false, MaxRepeat: 999}, // Free text
			{Segment: "DTM", Mandatory: true, MaxRepeat: 1},    // Date/time/period
			{ // Segment group 4
				Group: []SchemaNode{
					{Segment: "SGU", Mandatory: true, MaxRepeat: 1},   // Segment usage details
					{Segment: "ATT", Mandatory: false, MaxRepeat: 9},  // Attribute
					{Segment: "FTX", Mandatory: false, MaxRepeat: 99}, // Free text
					{ // Segment group 5
						Group: []SchemaNode{
							{Segment: "GRU", Mandatory: true, MaxRepeat: 1},   // Segment group usage details
							{Segment: "ATT", Mandatory: false, MaxRepeat: 9},  // Attribute
							{Segment: "FTX", Mandatory: false, MaxRepeat: 99}, // Free text
						},
						Mandatory: false, MaxRepeat: 1,
					},
				},
				Mandatory: false, MaxRepeat: 999,
			},
			{ // Segment group 6
				Group: []SchemaNode{
					{Segment: "FNT", Mandatory: true, MaxRepeat: 1},  // Footnote
					{Segment: "REL", Mandatory: false, MaxRepeat: 1}, // Relationship
					{Segment: "GIR", Mandatory: false, MaxRepeat: 9}, // Related identification numbers
					{Segment: "ATT", Mandatory: false, MaxRepeat: 9}, // Attribute
					{Segment: "FTX", Mandatory: false, MaxRepeat: 9}, // Free text
				},
				Mandatory: false, MaxRepeat: 99,
			},
		},
		Mandatory: false, MaxRepeat: 9999,
	},
	{ // Segment group 7
		Group: []SchemaNode{
			{Segment: "SEG", Mandatory: true, MaxRepeat: 1},   // Segment identification
			{Segment: "ATT", Mandatory: false, MaxRepeat: 9},  // Attribute
			{Segment: "FTX", Mandatory: false, MaxRepeat: 9},  // Free text
			{Segment: "ELU", Mandatory: false, MaxRepeat: 99}, // Data element usage details
			{ // Segment group 8
				Group: []SchemaNode{
					{Segment: "FNT", Mandatory: true, MaxRepeat: 1},  // Footnote
					{Segment: "REL", Mandatory: false, MaxRepeat: 1}, // Relationship
					{Segment: "GIR", Mandatory: false, MaxRepeat: 9}, // Related identification numbers
					{Segment: "ATT", Mandatory: false, MaxRepeat: 9}, // Attribute
					{Segment: "FTX", Mandatory: false, MaxRepeat: 9}, // Free text
				},
				Mandatory: false, MaxRepeat: 99,
			},
		},
		Mandatory: false, MaxRepeat: 9999,
	},
	{ // Segment group 9
		Group: []SchemaNode{
			{Segment: "CMP", Mandatory: true, MaxRepeat: 1},   // Composite data element identification
			{Segment: "ATT", Mandatory: false, MaxRepeat: 9},  // Attribute
			{Segment: "FTX", Mandatory: false, MaxRepeat: 9},  // Free text
			{Segment: "ELU", Mandatory: false, MaxRepeat: 99}, // Data element usage details
			{ // Segment group 10
				Group: []SchemaNode{
					{Segment: "FNT", Mandatory: true, MaxRepeat: 1},  // Footnote
					{Segment: "REL", Mandatory: false, MaxRepeat: 1}, // Relationship
					{Segment: "GIR", Mandatory: false, MaxRepeat: 9}, // Related identification numbers
					{Segment: "ATT", Mandatory: false, MaxRepeat: 9}, // Attribute
					{Segment: "FTX", Mandatory: false, MaxRepeat: 9}, // Free text
				},
				Mandatory: false, MaxRepeat: 99,
			},
		},
		Mandatory: false, MaxRepeat: 9999,
	},
	{ // Segment group 11
		Group: []SchemaNode{
			{Segment: "ELM", Mandatory: true, MaxRepeat: 1},  // Simple data element details
			{Segment: "ATT", Mandatory: false, MaxRepeat: 9}, // Attribute
			{Segment: "FTX", Mandatory: false, MaxRepeat: 9}, // Free text
		},
		Mandatory: false, MaxRepeat: 9999,
	},
	{ // Segment group 12
		Group: []SchemaNode{
			{Segment: "CDS", Mandatory: true, MaxRepeat: 1},  // Code set identification
			{Segment: "FTX", Mandatory: false, MaxRepeat: 9}, // Free text
			{ // Segment group 13
				Group: []SchemaNode{
					{Segment: "CDV", Mandatory: true, MaxRepeat: 1},  // Code value definition
					{Segment: "ATT", Mandatory: false, MaxRepeat: 9}, // Attribute
					{Segment: "FTX", Mandatory: false, MaxRepeat: 9}, // Free text
				},
				Mandatory: false, MaxRepeat: 9999,
			},
		},
		Mandatory: false, MaxRepeat: 9999,
	},
}}
