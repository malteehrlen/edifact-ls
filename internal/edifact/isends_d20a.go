package edifact

// ISENDS D.20A message specification: the structural (segment/group
// presence, order, cardinality) branching diagram for this message,
// UN/EDIFACT directory release D.20A. Scope is structural only -- not
// element-level content -- per edifact-ls-3uzr's epic non-goal.
// 8 segment groups, max nesting depth 3.
//
// Source: https://service.unece.org/trade/untdid/d20a/trmd/isends_c.htm
// section 4.3.1 "Segment table". That URL currently returns HTTP 403 from
// Cloudflare when fetched directly; this data was transcribed from the
// Wayback Machine's archived copy instead:
// http://web.archive.org/web/20240417090506/https://service.unece.org/trade/untdid/d20a/trmd/isends_c.htm
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
		MessageID{Type: "ISENDS", Version: "D", Release: "20A", Agency: "UN"},
		isendsD20aSchema,
		"https://service.unece.org/trade/untdid/d20a/trmd/isends_c.htm",
	)
}

var isendsD20aSchema = Schema{Nodes: []SchemaNode{
	{Segment: "BGM", Mandatory: true, MaxRepeat: 1},  // Beginning of message
	{Segment: "ATT", Mandatory: true, MaxRepeat: 1},  // Attribute
	{Segment: "DTM", Mandatory: false, MaxRepeat: 9}, // Date/time/period
	{ // Segment group 1
		Group: []SchemaNode{
			{Segment: "PNA", Mandatory: true, MaxRepeat: 1},  // Party identification
			{Segment: "ADR", Mandatory: false, MaxRepeat: 9}, // Address
			{ // Segment group 2
				Group: []SchemaNode{
					{Segment: "CTA", Mandatory: true, MaxRepeat: 1},  // Contact information
					{Segment: "COM", Mandatory: false, MaxRepeat: 9}, // Communication contact
					{Segment: "GEI", Mandatory: false, MaxRepeat: 9}, // Processing information
					{Segment: "SEQ", Mandatory: false, MaxRepeat: 1}, // Sequence details
					{Segment: "ATT", Mandatory: false, MaxRepeat: 9}, // Attribute
				},
				Mandatory: false, MaxRepeat: 9,
			},
		},
		Mandatory: true, MaxRepeat: 9,
	},
	{ // Segment group 3
		Group: []SchemaNode{
			{Segment: "RFF", Mandatory: true, MaxRepeat: 1},  // Reference
			{Segment: "DTM", Mandatory: false, MaxRepeat: 9}, // Date/time/period
		},
		Mandatory: false, MaxRepeat: 9,
	},
	{ // Segment group 4
		Group: []SchemaNode{
			{Segment: "ICD", Mandatory: true, MaxRepeat: 1},  // Insurance cover description
			{Segment: "RFF", Mandatory: false, MaxRepeat: 9}, // Reference
			{Segment: "DTM", Mandatory: false, MaxRepeat: 9}, // Date/time/period
			{Segment: "MKS", Mandatory: false, MaxRepeat: 9}, // Market/sales channel information
			{ // Segment group 5
				Group: []SchemaNode{
					{Segment: "PRC", Mandatory: true, MaxRepeat: 1},  // Process identification
					{Segment: "GEI", Mandatory: false, MaxRepeat: 1}, // Processing information
					{Segment: "DTM", Mandatory: false, MaxRepeat: 9}, // Date/time/period
					{ // Segment group 6
						Group: []SchemaNode{
							{Segment: "ATT", Mandatory: true, MaxRepeat: 1},  // Attribute
							{Segment: "MOA", Mandatory: false, MaxRepeat: 1}, // Monetary amount
							{Segment: "PCD", Mandatory: false, MaxRepeat: 1}, // Percentage details
							{Segment: "RCS", Mandatory: false, MaxRepeat: 9}, // Requirements and conditions
						},
						Mandatory: false, MaxRepeat: 9,
					},
					{ // Segment group 7
						Group: []SchemaNode{
							{Segment: "DOC", Mandatory: true, MaxRepeat: 1},  // Document/message details
							{Segment: "QRS", Mandatory: false, MaxRepeat: 1}, // Query and response
						},
						Mandatory: false, MaxRepeat: 9,
					},
				},
				Mandatory: false, MaxRepeat: 9,
			},
		},
		Mandatory: false, MaxRepeat: 99,
	},
	{ // Segment group 8
		Group: []SchemaNode{
			{Segment: "RNG", Mandatory: true, MaxRepeat: 1},  // Range details
			{Segment: "GEI", Mandatory: false, MaxRepeat: 9}, // Processing information
			{Segment: "MOA", Mandatory: false, MaxRepeat: 9}, // Monetary amount
			{Segment: "RFF", Mandatory: false, MaxRepeat: 9}, // Reference
			{Segment: "DTM", Mandatory: false, MaxRepeat: 9}, // Date/time/period
			{Segment: "MKS", Mandatory: false, MaxRepeat: 9}, // Market/sales channel information
			{Segment: "ICD", Mandatory: false, MaxRepeat: 9}, // Insurance cover description
		},
		Mandatory: false, MaxRepeat: 9,
	},
}}
