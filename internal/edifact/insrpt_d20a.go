package edifact

// INSRPT D.20A message specification: the structural (segment/group
// presence, order, cardinality) branching diagram for this message,
// UN/EDIFACT directory release D.20A. Scope is structural only -- not
// element-level content -- per edifact-ls-3uzr's epic non-goal.
// 9 segment groups, max nesting depth 4.
//
// Source: https://service.unece.org/trade/untdid/d20a/trmd/insrpt_c.htm
// section 4.3.1 "Segment table". That URL currently returns HTTP 403 from
// Cloudflare when fetched directly; this data was transcribed from the
// Wayback Machine's archived copy instead:
// http://web.archive.org/web/20230924164058/https://service.unece.org/trade/untdid/d20a/trmd/insrpt_c.htm
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
		MessageID{Type: "INSRPT", Version: "D", Release: "20A", Agency: "UN"},
		insrptD20aSchema,
		"https://service.unece.org/trade/untdid/d20a/trmd/insrpt_c.htm",
	)
}

var insrptD20aSchema = Schema{Nodes: []SchemaNode{
	{Segment: "BGM", Mandatory: true, MaxRepeat: 1},  // Beginning of message
	{Segment: "DTM", Mandatory: true, MaxRepeat: 99}, // Date/time/period
	{ // Segment group 1
		Group: []SchemaNode{
			{Segment: "RFF", Mandatory: true, MaxRepeat: 1},   // Reference
			{Segment: "DTM", Mandatory: false, MaxRepeat: 99}, // Date/time/period
		},
		Mandatory: false, MaxRepeat: 99,
	},
	{ // Segment group 2
		Group: []SchemaNode{
			{Segment: "NAD", Mandatory: true, MaxRepeat: 1},   // Name and address
			{Segment: "LOC", Mandatory: false, MaxRepeat: 99}, // Place/location identification
		},
		Mandatory: true, MaxRepeat: 99,
	},
	{ // Segment group 3
		Group: []SchemaNode{
			{Segment: "DOC", Mandatory: true, MaxRepeat: 1}, // Document/message details
			{ // Segment group 4
				Group: []SchemaNode{
					{Segment: "RFF", Mandatory: true, MaxRepeat: 1},   // Reference
					{Segment: "DTM", Mandatory: false, MaxRepeat: 99}, // Date/time/period
				},
				Mandatory: false, MaxRepeat: 99,
			},
			{ // Segment group 5
				Group: []SchemaNode{
					{Segment: "NAD", Mandatory: true, MaxRepeat: 1}, // Name and address
					{ // Segment group 6
						Group: []SchemaNode{
							{Segment: "CTA", Mandatory: true, MaxRepeat: 1},   // Contact information
							{Segment: "COM", Mandatory: false, MaxRepeat: 99}, // Communication contact
						},
						Mandatory: false, MaxRepeat: 99,
					},
				},
				Mandatory: false, MaxRepeat: 99,
			},
			{ // Segment group 7
				Group: []SchemaNode{
					{Segment: "LIN", Mandatory: true, MaxRepeat: 1},    // Line item
					{Segment: "PIA", Mandatory: false, MaxRepeat: 99},  // Additional product id
					{Segment: "IMD", Mandatory: false, MaxRepeat: 99},  // Item description
					{Segment: "DTM", Mandatory: false, MaxRepeat: 99},  // Date/time/period
					{Segment: "STS", Mandatory: false, MaxRepeat: 99},  // Status
					{Segment: "MOA", Mandatory: false, MaxRepeat: 99},  // Monetary amount
					{Segment: "QTY", Mandatory: false, MaxRepeat: 99},  // Quantity
					{Segment: "MEA", Mandatory: false, MaxRepeat: 99},  // Measurements
					{Segment: "GIN", Mandatory: false, MaxRepeat: 999}, // Goods identity number
					{Segment: "RFF", Mandatory: false, MaxRepeat: 99},  // Reference
					{Segment: "FTX", Mandatory: false, MaxRepeat: 9},   // Free text
					{ // Segment group 8
						Group: []SchemaNode{
							{Segment: "NAD", Mandatory: true, MaxRepeat: 1},   // Name and address
							{Segment: "GEI", Mandatory: false, MaxRepeat: 99}, // Processing information
							{Segment: "LOC", Mandatory: false, MaxRepeat: 99}, // Place/location identification
							{Segment: "QTY", Mandatory: false, MaxRepeat: 99}, // Quantity
							{Segment: "QVR", Mandatory: false, MaxRepeat: 99}, // Quantity variances
							{Segment: "RFF", Mandatory: false, MaxRepeat: 99}, // Reference
							{Segment: "FTX", Mandatory: false, MaxRepeat: 99}, // Free text
							{ // Segment group 9
								Group: []SchemaNode{
									{Segment: "GIN", Mandatory: true, MaxRepeat: 1},   // Goods identity number
									{Segment: "RFF", Mandatory: false, MaxRepeat: 99}, // Reference
									{Segment: "STS", Mandatory: false, MaxRepeat: 99}, // Status
								},
								Mandatory: false, MaxRepeat: 999,
							},
						},
						Mandatory: false, MaxRepeat: 999,
					},
				},
				Mandatory: false, MaxRepeat: 999,
			},
		},
		Mandatory: true, MaxRepeat: 99,
	},
}}
