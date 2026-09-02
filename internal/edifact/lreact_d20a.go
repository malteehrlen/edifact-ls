package edifact

// LREACT D.20A message specification: the structural (segment/group
// presence, order, cardinality) branching diagram for this message,
// UN/EDIFACT directory release D.20A. Scope is structural only -- not
// element-level content -- per edifact-ls-3uzr's epic non-goal.
// 9 segment groups, max nesting depth 4.
//
// Source: https://service.unece.org/trade/untdid/d20a/trmd/lreact_c.htm
// section 4.3.1 "Segment table". That URL currently returns HTTP 403 from
// Cloudflare when fetched directly; this data was transcribed from the
// Wayback Machine's archived copy instead:
// http://web.archive.org/web/20231208133635/https://service.unece.org/trade/untdid/d20a/trmd/lreact_c.htm
// -- re-check the direct URL first if cross-checking later, in case the
// block has lifted.
//
// Transcription note: generated in bulk (edifact-ls-13gu) the same way
// as every other message-type schema -- a script parsed the source's
// exact rail-art column positions mechanically and verified the result
// balances (all 9 groups open and close consistently)
// before being transcribed here, rather than reading the ASCII nesting
// by eye. Not individually spot-checked against the raw source the way
// the first several message types in this project were -- see
// edifact-ls-13gu for which representative samples were spot-checked
// for this batch.

func init() {
	RegisterSchema(
		MessageID{Type: "LREACT", Version: "D", Release: "20A", Agency: "UN"},
		lreactD20aSchema,
		"https://service.unece.org/trade/untdid/d20a/trmd/lreact_c.htm",
	)
}

var lreactD20aSchema = Schema{Nodes: []SchemaNode{
	{Segment: "BGM", Mandatory: true, MaxRepeat: 1},  // Beginning of message
	{Segment: "DTM", Mandatory: true, MaxRepeat: 1},  // Date/time/period
	{Segment: "CUX", Mandatory: false, MaxRepeat: 1}, // Currencies
	{Segment: "CTA", Mandatory: false, MaxRepeat: 9}, // Contact information
	{Segment: "FTX", Mandatory: false, MaxRepeat: 1}, // Free text
	{ // Segment group 1
		Group: []SchemaNode{
			{Segment: "GEI", Mandatory: true, MaxRepeat: 1},   // Processing information
			{Segment: "RFF", Mandatory: true, MaxRepeat: 9},   // Reference
			{Segment: "ATT", Mandatory: false, MaxRepeat: 99}, // Attribute
			{Segment: "DTM", Mandatory: false, MaxRepeat: 99}, // Date/time/period
			{Segment: "FTX", Mandatory: false, MaxRepeat: 1},  // Free text
			{ // Segment group 2
				Group: []SchemaNode{
					{Segment: "IMD", Mandatory: true, MaxRepeat: 1},  // Item description
					{Segment: "LOC", Mandatory: true, MaxRepeat: 9},  // Place/location identification
					{Segment: "SEQ", Mandatory: false, MaxRepeat: 1}, // Sequence details
					{Segment: "CUX", Mandatory: false, MaxRepeat: 1}, // Currencies
					{Segment: "PNA", Mandatory: false, MaxRepeat: 9}, // Party identification
				},
				Mandatory: false, MaxRepeat: 1,
			},
			{ // Segment group 3
				Group: []SchemaNode{
					{Segment: "PCD", Mandatory: true, MaxRepeat: 1},  // Percentage details
					{Segment: "PNA", Mandatory: false, MaxRepeat: 1}, // Party identification
					{Segment: "LOC", Mandatory: false, MaxRepeat: 1}, // Place/location identification
				},
				Mandatory: false, MaxRepeat: 1,
			},
			{ // Segment group 4
				Group: []SchemaNode{
					{Segment: "ICD", Mandatory: true, MaxRepeat: 1},  // Insurance cover description
					{Segment: "PCD", Mandatory: true, MaxRepeat: 9},  // Percentage details
					{Segment: "EMP", Mandatory: false, MaxRepeat: 1}, // Employment details
					{Segment: "AGR", Mandatory: false, MaxRepeat: 9}, // Agreement identification
					{Segment: "LOC", Mandatory: false, MaxRepeat: 1}, // Place/location identification
					{ // Segment group 5
						Group: []SchemaNode{
							{Segment: "MOA", Mandatory: true, MaxRepeat: 1}, // Monetary amount
							{Segment: "ARD", Mandatory: true, MaxRepeat: 1}, // Monetary amount function
						},
						Mandatory: false, MaxRepeat: 99,
					},
					{ // Segment group 6
						Group: []SchemaNode{
							{Segment: "RTE", Mandatory: true, MaxRepeat: 1},  // Rate details
							{Segment: "ATT", Mandatory: false, MaxRepeat: 1}, // Attribute
							{Segment: "DTM", Mandatory: false, MaxRepeat: 1}, // Date/time/period
							{ // Segment group 7
								Group: []SchemaNode{
									{Segment: "MOA", Mandatory: true, MaxRepeat: 1}, // Monetary amount
									{Segment: "ARD", Mandatory: true, MaxRepeat: 1}, // Monetary amount function
								},
								Mandatory: false, MaxRepeat: 99,
							},
						},
						Mandatory: false, MaxRepeat: 99,
					},
					{ // Segment group 8
						Group: []SchemaNode{
							{Segment: "PRC", Mandatory: true, MaxRepeat: 1},  // Process identification
							{Segment: "ATT", Mandatory: false, MaxRepeat: 9}, // Attribute
							{Segment: "PCD", Mandatory: false, MaxRepeat: 1}, // Percentage details
							{ // Segment group 9
								Group: []SchemaNode{
									{Segment: "MOA", Mandatory: true, MaxRepeat: 1}, // Monetary amount
									{Segment: "ARD", Mandatory: true, MaxRepeat: 1}, // Monetary amount function
								},
								Mandatory: false, MaxRepeat: 99,
							},
						},
						Mandatory: false, MaxRepeat: 99,
					},
				},
				Mandatory: false, MaxRepeat: 1,
			},
		},
		Mandatory: true, MaxRepeat: 999999,
	},
}}
