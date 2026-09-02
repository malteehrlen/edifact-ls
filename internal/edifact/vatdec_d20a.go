package edifact

// VATDEC D.20A message specification: the structural (segment/group
// presence, order, cardinality) branching diagram for this message,
// UN/EDIFACT directory release D.20A. Scope is structural only -- not
// element-level content -- per edifact-ls-3uzr's epic non-goal.
// 12 segment groups, max nesting depth 5.
//
// Source: https://service.unece.org/trade/untdid/d20a/trmd/vatdec_c.htm
// section 4.3.1 "Segment table". That URL currently returns HTTP 403 from
// Cloudflare when fetched directly; this data was transcribed from the
// Wayback Machine's archived copy instead:
// http://web.archive.org/web/20240421084705/https://service.unece.org/trade/untdid/d20a/trmd/vatdec_c.htm
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
		MessageID{Type: "VATDEC", Version: "D", Release: "20A", Agency: "UN"},
		vatdecD20aSchema,
		"https://service.unece.org/trade/untdid/d20a/trmd/vatdec_c.htm",
	)
}

var vatdecD20aSchema = Schema{Nodes: []SchemaNode{
	{Segment: "BGM", Mandatory: true, MaxRepeat: 1},  // Beginning of message
	{Segment: "DTM", Mandatory: false, MaxRepeat: 1}, // Date/time/period
	{Segment: "LOC", Mandatory: false, MaxRepeat: 1}, // Place/location identification
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
			{ // Segment group 3
				Group: []SchemaNode{
					{Segment: "RFF", Mandatory: true, MaxRepeat: 1},  // Reference
					{Segment: "FTX", Mandatory: false, MaxRepeat: 9}, // Free text
				},
				Mandatory: false, MaxRepeat: 9,
			},
			{Segment: "FII", Mandatory: false, MaxRepeat: 9}, // Financial institution information
			{Segment: "PAI", Mandatory: false, MaxRepeat: 9}, // Payment instructions
			{Segment: "GEI", Mandatory: false, MaxRepeat: 9}, // Processing information
			{Segment: "LOC", Mandatory: false, MaxRepeat: 9}, // Place/location identification
			{ // Segment group 4
				Group: []SchemaNode{
					{Segment: "DTM", Mandatory: true, MaxRepeat: 1},  // Date/time/period
					{Segment: "GEI", Mandatory: false, MaxRepeat: 9}, // Processing information
					{Segment: "MEA", Mandatory: false, MaxRepeat: 9}, // Measurements
					{ // Segment group 5
						Group: []SchemaNode{
							{Segment: "DMS", Mandatory: true, MaxRepeat: 1},  // Document/message summary
							{Segment: "MOA", Mandatory: false, MaxRepeat: 1}, // Monetary amount
							{Segment: "MEA", Mandatory: false, MaxRepeat: 1}, // Measurements
							{Segment: "PCD", Mandatory: false, MaxRepeat: 1}, // Percentage details
							{Segment: "GEI", Mandatory: false, MaxRepeat: 1}, // Processing information
						},
						Mandatory: false, MaxRepeat: 9,
					},
					{ // Segment group 6
						Group: []SchemaNode{
							{Segment: "LIN", Mandatory: true, MaxRepeat: 1}, // Line item
							{ // Segment group 7
								Group: []SchemaNode{
									{Segment: "NAD", Mandatory: true, MaxRepeat: 1}, // Name and address
									{ // Segment group 8
										Group: []SchemaNode{
											{Segment: "CTA", Mandatory: true, MaxRepeat: 1},  // Contact information
											{Segment: "COM", Mandatory: false, MaxRepeat: 9}, // Communication contact
										},
										Mandatory: false, MaxRepeat: 9999,
									},
									{Segment: "RFF", Mandatory: false, MaxRepeat: 9}, // Reference
									{Segment: "FII", Mandatory: false, MaxRepeat: 1}, // Financial institution information
								},
								Mandatory: false, MaxRepeat: 9999,
							},
							{Segment: "GEI", Mandatory: false, MaxRepeat: 9}, // Processing information
							{ // Segment group 9
								Group: []SchemaNode{
									{Segment: "MOA", Mandatory: true, MaxRepeat: 1},  // Monetary amount
									{Segment: "ARD", Mandatory: false, MaxRepeat: 1}, // Monetary amount function
									{Segment: "PCD", Mandatory: false, MaxRepeat: 1}, // Percentage details
									{Segment: "GEI", Mandatory: false, MaxRepeat: 9}, // Processing information
									{ // Segment group 10
										Group: []SchemaNode{
											{Segment: "DMS", Mandatory: true, MaxRepeat: 1},  // Document/message summary
											{Segment: "MOA", Mandatory: false, MaxRepeat: 1}, // Monetary amount
											{Segment: "MEA", Mandatory: false, MaxRepeat: 1}, // Measurements
											{Segment: "PCD", Mandatory: false, MaxRepeat: 1}, // Percentage details
											{Segment: "GEI", Mandatory: true, MaxRepeat: 1},  // Processing information
										},
										Mandatory: false, MaxRepeat: 9,
									},
								},
								Mandatory: true, MaxRepeat: 9999,
							},
						},
						Mandatory: true, MaxRepeat: 9999,
					},
					{ // Segment group 11
						Group: []SchemaNode{
							{Segment: "CNT", Mandatory: true, MaxRepeat: 1}, // Control total
							{Segment: "MOA", Mandatory: true, MaxRepeat: 1}, // Monetary amount
						},
						Mandatory: false, MaxRepeat: 1,
					},
				},
				Mandatory: false, MaxRepeat: 99,
			},
		},
		Mandatory: true, MaxRepeat: 9999,
	},
	{ // Segment group 12
		Group: []SchemaNode{
			{Segment: "AUT", Mandatory: true, MaxRepeat: 1},  // Authentication result
			{Segment: "DTM", Mandatory: false, MaxRepeat: 1}, // Date/time/period
			{Segment: "RFF", Mandatory: false, MaxRepeat: 1}, // Reference
		},
		Mandatory: false, MaxRepeat: 1,
	},
}}
