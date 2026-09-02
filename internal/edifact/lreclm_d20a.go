package edifact

// LRECLM D.20A message specification: the structural (segment/group
// presence, order, cardinality) branching diagram for this message,
// UN/EDIFACT directory release D.20A. Scope is structural only -- not
// element-level content -- per edifact-ls-3uzr's epic non-goal.
// 9 segment groups, max nesting depth 5.
//
// Source: https://service.unece.org/trade/untdid/d20a/trmd/lreclm_c.htm
// section 4.3.1 "Segment table". That URL currently returns HTTP 403 from
// Cloudflare when fetched directly; this data was transcribed from the
// Wayback Machine's archived copy instead:
// http://web.archive.org/web/20231208132606/https://service.unece.org/trade/untdid/d20a/trmd/lreclm_c.htm
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
		MessageID{Type: "LRECLM", Version: "D", Release: "20A", Agency: "UN"},
		lreclmD20aSchema,
		"https://service.unece.org/trade/untdid/d20a/trmd/lreclm_c.htm",
	)
}

var lreclmD20aSchema = Schema{Nodes: []SchemaNode{
	{Segment: "BGM", Mandatory: true, MaxRepeat: 1},  // Beginning of message
	{Segment: "DTM", Mandatory: true, MaxRepeat: 1},  // Date/time/period
	{Segment: "CUX", Mandatory: false, MaxRepeat: 1}, // Currencies
	{Segment: "CTA", Mandatory: false, MaxRepeat: 9}, // Contact information
	{Segment: "FTX", Mandatory: false, MaxRepeat: 1}, // Free text
	{ // Segment group 1
		Group: []SchemaNode{
			{Segment: "IMD", Mandatory: true, MaxRepeat: 1},  // Item description
			{Segment: "ATT", Mandatory: true, MaxRepeat: 99}, // Attribute
			{Segment: "DTM", Mandatory: false, MaxRepeat: 9}, // Date/time/period
			{Segment: "LOC", Mandatory: false, MaxRepeat: 1}, // Place/location identification
			{Segment: "PCD", Mandatory: false, MaxRepeat: 9}, // Percentage details
			{Segment: "RFF", Mandatory: false, MaxRepeat: 9}, // Reference
			{Segment: "CUX", Mandatory: false, MaxRepeat: 1}, // Currencies
			{Segment: "FTX", Mandatory: false, MaxRepeat: 1}, // Free text
			{ // Segment group 2
				Group: []SchemaNode{
					{Segment: "MOA", Mandatory: true, MaxRepeat: 1}, // Monetary amount
					{Segment: "ARD", Mandatory: true, MaxRepeat: 1}, // Monetary amount function
				},
				Mandatory: false, MaxRepeat: 9,
			},
			{ // Segment group 3
				Group: []SchemaNode{
					{Segment: "PNA", Mandatory: true, MaxRepeat: 1},  // Party identification
					{Segment: "DTM", Mandatory: true, MaxRepeat: 9},  // Date/time/period
					{Segment: "LOC", Mandatory: false, MaxRepeat: 1}, // Place/location identification
					{Segment: "EMP", Mandatory: false, MaxRepeat: 1}, // Employment details
					{ // Segment group 4
						Group: []SchemaNode{
							{Segment: "MOA", Mandatory: true, MaxRepeat: 1}, // Monetary amount
							{Segment: "ARD", Mandatory: true, MaxRepeat: 1}, // Monetary amount function
						},
						Mandatory: false, MaxRepeat: 9,
					},
				},
				Mandatory: true, MaxRepeat: 1,
			},
			{ // Segment group 5
				Group: []SchemaNode{
					{Segment: "GEI", Mandatory: true, MaxRepeat: 1},   // Processing information
					{Segment: "RFF", Mandatory: true, MaxRepeat: 9},   // Reference
					{Segment: "PNA", Mandatory: false, MaxRepeat: 1},  // Party identification
					{Segment: "DTM", Mandatory: false, MaxRepeat: 99}, // Date/time/period
					{ // Segment group 6
						Group: []SchemaNode{
							{Segment: "MOA", Mandatory: true, MaxRepeat: 1}, // Monetary amount
							{Segment: "ARD", Mandatory: true, MaxRepeat: 1}, // Monetary amount function
						},
						Mandatory: false, MaxRepeat: 9,
					},
					{ // Segment group 7
						Group: []SchemaNode{
							{Segment: "IDE", Mandatory: true, MaxRepeat: 1}, // Identity
							{Segment: "PNA", Mandatory: true, MaxRepeat: 1}, // Party identification
							{Segment: "DTM", Mandatory: true, MaxRepeat: 1}, // Date/time/period
							{ // Segment group 8
								Group: []SchemaNode{
									{Segment: "ICD", Mandatory: true, MaxRepeat: 1},   // Insurance cover description
									{Segment: "DTM", Mandatory: false, MaxRepeat: 99}, // Date/time/period
									{Segment: "RFF", Mandatory: false, MaxRepeat: 9},  // Reference
									{Segment: "ATT", Mandatory: false, MaxRepeat: 9},  // Attribute
									{Segment: "PCD", Mandatory: false, MaxRepeat: 1},  // Percentage details
									{Segment: "AGR", Mandatory: false, MaxRepeat: 1},  // Agreement identification
									{ // Segment group 9
										Group: []SchemaNode{
											{Segment: "MOA", Mandatory: true, MaxRepeat: 1}, // Monetary amount
											{Segment: "ARD", Mandatory: true, MaxRepeat: 1}, // Monetary amount function
										},
										Mandatory: false, MaxRepeat: 9,
									},
								},
								Mandatory: false, MaxRepeat: 99,
							},
						},
						Mandatory: false, MaxRepeat: 99,
					},
				},
				Mandatory: true, MaxRepeat: 99,
			},
		},
		Mandatory: true, MaxRepeat: 999999,
	},
}}
