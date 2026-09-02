package edifact

// MOVINS D.20A message specification: the structural (segment/group
// presence, order, cardinality) branching diagram for this message,
// UN/EDIFACT directory release D.20A. Scope is structural only -- not
// element-level content -- per edifact-ls-3uzr's epic non-goal.
// 13 segment groups, max nesting depth 5.
//
// Source: https://service.unece.org/trade/untdid/d20a/trmd/movins_c.htm
// section 4.3.1 "Segment table". That URL currently returns HTTP 403 from
// Cloudflare when fetched directly; this data was transcribed from the
// Wayback Machine's archived copy instead:
// http://web.archive.org/web/20231201230645/https://service.unece.org/trade/untdid/d20a/trmd/movins_c.htm
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
		MessageID{Type: "MOVINS", Version: "D", Release: "20A", Agency: "UN"},
		movinsD20aSchema,
		"https://service.unece.org/trade/untdid/d20a/trmd/movins_c.htm",
	)
}

var movinsD20aSchema = Schema{Nodes: []SchemaNode{
	{Segment: "BGM", Mandatory: true, MaxRepeat: 1},  // Beginning of message
	{Segment: "DTM", Mandatory: false, MaxRepeat: 9}, // Date/time/period
	{ // Segment group 1
		Group: []SchemaNode{
			{Segment: "RFF", Mandatory: true, MaxRepeat: 1},  // Reference
			{Segment: "DTM", Mandatory: false, MaxRepeat: 9}, // Date/time/period
		},
		Mandatory: false, MaxRepeat: 9,
	},
	{ // Segment group 2
		Group: []SchemaNode{
			{Segment: "NAD", Mandatory: true, MaxRepeat: 1}, // Name and address
			{ // Segment group 3
				Group: []SchemaNode{
					{Segment: "CTA", Mandatory: true, MaxRepeat: 1},  // Contact information
					{Segment: "COM", Mandatory: false, MaxRepeat: 9}, // Communication contact
				},
				Mandatory: false, MaxRepeat: 9,
			},
		},
		Mandatory: false, MaxRepeat: 9,
	},
	{ // Segment group 4
		Group: []SchemaNode{
			{Segment: "TDT", Mandatory: true, MaxRepeat: 1},  // Transport information
			{Segment: "RFF", Mandatory: false, MaxRepeat: 9}, // Reference
			{Segment: "FTX", Mandatory: false, MaxRepeat: 9}, // Free text
			{ // Segment group 5
				Group: []SchemaNode{
					{Segment: "LOC", Mandatory: true, MaxRepeat: 1},  // Place/location identification
					{Segment: "DTM", Mandatory: false, MaxRepeat: 9}, // Date/time/period
				},
				Mandatory: false, MaxRepeat: 9,
			},
		},
		Mandatory: true, MaxRepeat: 9,
	},
	{ // Segment group 6
		Group: []SchemaNode{
			{Segment: "HAN", Mandatory: true, MaxRepeat: 1}, // Handling instructions
			{ // Segment group 7
				Group: []SchemaNode{
					{Segment: "LOC", Mandatory: true, MaxRepeat: 1},  // Place/location identification
					{Segment: "FTX", Mandatory: false, MaxRepeat: 9}, // Free text
					{Segment: "RFF", Mandatory: false, MaxRepeat: 9}, // Reference
					{ // Segment group 8
						Group: []SchemaNode{
							{Segment: "EQD", Mandatory: true, MaxRepeat: 1},   // Equipment details
							{Segment: "NAD", Mandatory: false, MaxRepeat: 9},  // Name and address
							{Segment: "MEA", Mandatory: false, MaxRepeat: 99}, // Measurements
							{Segment: "HAN", Mandatory: false, MaxRepeat: 99}, // Handling instructions
							{Segment: "DIM", Mandatory: false, MaxRepeat: 9},  // Dimensions
							{Segment: "RFF", Mandatory: false, MaxRepeat: 9},  // Reference
							{Segment: "GDS", Mandatory: false, MaxRepeat: 99}, // Nature of cargo
							{Segment: "FTX", Mandatory: false, MaxRepeat: 9},  // Free text
							{ // Segment group 9
								Group: []SchemaNode{
									{Segment: "LOC", Mandatory: true, MaxRepeat: 1},  // Place/location identification
									{Segment: "TSR", Mandatory: false, MaxRepeat: 1}, // Transport service requirements
									{Segment: "TDT", Mandatory: false, MaxRepeat: 1}, // Transport information
								},
								Mandatory: false, MaxRepeat: 9,
							},
							{ // Segment group 10
								Group: []SchemaNode{
									{Segment: "TMP", Mandatory: true, MaxRepeat: 1},  // Temperature
									{Segment: "RNG", Mandatory: false, MaxRepeat: 1}, // Range details
									{Segment: "DTM", Mandatory: false, MaxRepeat: 1}, // Date/time/period
								},
								Mandatory: false, MaxRepeat: 9,
							},
							{ // Segment group 11
								Group: []SchemaNode{
									{Segment: "EQA", Mandatory: true, MaxRepeat: 1},  // Attached equipment
									{Segment: "NAD", Mandatory: false, MaxRepeat: 1}, // Name and address
								},
								Mandatory: false, MaxRepeat: 9,
							},
							{ // Segment group 12
								Group: []SchemaNode{
									{Segment: "DGS", Mandatory: true, MaxRepeat: 1},  // Dangerous goods
									{Segment: "ATT", Mandatory: false, MaxRepeat: 9}, // Attribute
									{Segment: "MEA", Mandatory: false, MaxRepeat: 9}, // Measurements
									{Segment: "FTX", Mandatory: false, MaxRepeat: 9}, // Free text
									{ // Segment group 13
										Group: []SchemaNode{
											{Segment: "CTA", Mandatory: true, MaxRepeat: 1},  // Contact information
											{Segment: "COM", Mandatory: false, MaxRepeat: 9}, // Communication contact
										},
										Mandatory: false, MaxRepeat: 9,
									},
								},
								Mandatory: false, MaxRepeat: 999,
							},
						},
						Mandatory: false, MaxRepeat: 9,
					},
					{Segment: "CNT", Mandatory: true, MaxRepeat: 1}, // Control total
				},
				Mandatory: false, MaxRepeat: 99999,
			},
		},
		Mandatory: true, MaxRepeat: 9,
	},
}}
