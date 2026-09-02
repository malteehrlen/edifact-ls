package edifact

// REGENT D.20A message specification: the structural (segment/group
// presence, order, cardinality) branching diagram for this message,
// UN/EDIFACT directory release D.20A. Scope is structural only -- not
// element-level content -- per edifact-ls-3uzr's epic non-goal.
// 14 segment groups, max nesting depth 5.
//
// Source: https://service.unece.org/trade/untdid/d20a/trmd/regent_c.htm
// section 4.3.1 "Segment table". That URL currently returns HTTP 403 from
// Cloudflare when fetched directly; this data was transcribed from the
// Wayback Machine's archived copy instead:
// http://web.archive.org/web/20230608124459/https://service.unece.org/trade/untdid/d20a/trmd/regent_c.htm
// -- re-check the direct URL first if cross-checking later, in case the
// block has lifted.
//
// Transcription note: generated in bulk (edifact-ls-13gu) the same way
// as every other message-type schema -- a script parsed the source's
// exact rail-art column positions mechanically and verified the result
// balances (all 14 groups open and close consistently)
// before being transcribed here, rather than reading the ASCII nesting
// by eye. Not individually spot-checked against the raw source the way
// the first several message types in this project were -- see
// edifact-ls-13gu for which representative samples were spot-checked
// for this batch.

func init() {
	RegisterSchema(
		MessageID{Type: "REGENT", Version: "D", Release: "20A", Agency: "UN"},
		regentD20aSchema,
		"https://service.unece.org/trade/untdid/d20a/trmd/regent_c.htm",
	)
}

var regentD20aSchema = Schema{Nodes: []SchemaNode{
	{Segment: "BGM", Mandatory: true, MaxRepeat: 1},  // Beginning of message
	{Segment: "CUX", Mandatory: false, MaxRepeat: 1}, // Currencies
	{ // Segment group 1
		Group: []SchemaNode{
			{Segment: "PNA", Mandatory: true, MaxRepeat: 1},  // Party identification
			{Segment: "RFF", Mandatory: false, MaxRepeat: 9}, // Reference
			{Segment: "ATT", Mandatory: false, MaxRepeat: 9}, // Attribute
			{Segment: "ADR", Mandatory: false, MaxRepeat: 9}, // Address
			{Segment: "COM", Mandatory: false, MaxRepeat: 9}, // Communication contact
		},
		Mandatory: true, MaxRepeat: 99,
	},
	{ // Segment group 2
		Group: []SchemaNode{
			{Segment: "DOC", Mandatory: true, MaxRepeat: 1},   // Document/message details
			{Segment: "RFF", Mandatory: false, MaxRepeat: 9},  // Reference
			{Segment: "DTM", Mandatory: false, MaxRepeat: 9},  // Date/time/period
			{Segment: "ATT", Mandatory: false, MaxRepeat: 9},  // Attribute
			{Segment: "ADR", Mandatory: false, MaxRepeat: 9},  // Address
			{Segment: "LOC", Mandatory: false, MaxRepeat: 9},  // Place/location identification
			{Segment: "COM", Mandatory: false, MaxRepeat: 9},  // Communication contact
			{Segment: "GEI", Mandatory: false, MaxRepeat: 1},  // Processing information
			{Segment: "FTX", Mandatory: false, MaxRepeat: 99}, // Free text
		},
		Mandatory: false, MaxRepeat: 99,
	},
	{ // Segment group 3
		Group: []SchemaNode{
			{Segment: "MOA", Mandatory: true, MaxRepeat: 1},  // Monetary amount
			{Segment: "PAI", Mandatory: false, MaxRepeat: 1}, // Payment instructions
			{Segment: "FII", Mandatory: false, MaxRepeat: 1}, // Financial institution information
			{Segment: "DTM", Mandatory: false, MaxRepeat: 9}, // Date/time/period
		},
		Mandatory: false, MaxRepeat: 9,
	},
	{ // Segment group 4
		Group: []SchemaNode{
			{Segment: "EVE", Mandatory: true, MaxRepeat: 1},  // Event
			{Segment: "ATT", Mandatory: false, MaxRepeat: 9}, // Attribute
			{Segment: "DTM", Mandatory: false, MaxRepeat: 9}, // Date/time/period
			{Segment: "FTX", Mandatory: false, MaxRepeat: 9}, // Free text
		},
		Mandatory: true, MaxRepeat: 99,
	},
	{ // Segment group 5
		Group: []SchemaNode{
			{Segment: "DSI", Mandatory: true, MaxRepeat: 1},  // Data set identification
			{Segment: "PNA", Mandatory: false, MaxRepeat: 9}, // Party identification
			{Segment: "REL", Mandatory: false, MaxRepeat: 1}, // Relationship
			{Segment: "RFF", Mandatory: false, MaxRepeat: 9}, // Reference
			{ // Segment group 6
				Group: []SchemaNode{
					{Segment: "EVE", Mandatory: true, MaxRepeat: 1},   // Event
					{Segment: "ATT", Mandatory: false, MaxRepeat: 9},  // Attribute
					{Segment: "DTM", Mandatory: false, MaxRepeat: 9},  // Date/time/period
					{Segment: "FTX", Mandatory: false, MaxRepeat: 99}, // Free text
					{ // Segment group 7
						Group: []SchemaNode{
							{Segment: "REL", Mandatory: true, MaxRepeat: 1},    // Relationship
							{Segment: "PNA", Mandatory: false, MaxRepeat: 9},   // Party identification
							{Segment: "RFF", Mandatory: false, MaxRepeat: 999}, // Reference
							{Segment: "NAT", Mandatory: false, MaxRepeat: 2},   // Nationality
							{Segment: "PDI", Mandatory: false, MaxRepeat: 1},   // Person demographic information
							{Segment: "DTM", Mandatory: false, MaxRepeat: 9},   // Date/time/period
							{Segment: "ADR", Mandatory: false, MaxRepeat: 9},   // Address
							{Segment: "LOC", Mandatory: false, MaxRepeat: 9},   // Place/location identification
							{Segment: "COM", Mandatory: false, MaxRepeat: 9},   // Communication contact
							{Segment: "QTY", Mandatory: false, MaxRepeat: 99},  // Quantity
							{Segment: "FTX", Mandatory: false, MaxRepeat: 999}, // Free text
							{ // Segment group 8
								Group: []SchemaNode{
									{Segment: "ATT", Mandatory: true, MaxRepeat: 1},   // Attribute
									{Segment: "GEI", Mandatory: false, MaxRepeat: 1},  // Processing information
									{Segment: "DTM", Mandatory: false, MaxRepeat: 9},  // Date/time/period
									{Segment: "MEA", Mandatory: false, MaxRepeat: 1},  // Measurements
									{Segment: "FTX", Mandatory: false, MaxRepeat: 99}, // Free text
								},
								Mandatory: false, MaxRepeat: 99,
							},
							{ // Segment group 9
								Group: []SchemaNode{
									{Segment: "MOA", Mandatory: true, MaxRepeat: 1},   // Monetary amount
									{Segment: "ATT", Mandatory: false, MaxRepeat: 9},  // Attribute
									{Segment: "DTM", Mandatory: false, MaxRepeat: 9},  // Date/time/period
									{Segment: "QTY", Mandatory: false, MaxRepeat: 9},  // Quantity
									{Segment: "FTX", Mandatory: false, MaxRepeat: 99}, // Free text
								},
								Mandatory: false, MaxRepeat: 99,
							},
							{ // Segment group 10
								Group: []SchemaNode{
									{Segment: "DOC", Mandatory: true, MaxRepeat: 1},  // Document/message details
									{Segment: "ATT", Mandatory: false, MaxRepeat: 9}, // Attribute
									{Segment: "LOC", Mandatory: false, MaxRepeat: 9}, // Place/location identification
									{Segment: "DTM", Mandatory: false, MaxRepeat: 9}, // Date/time/period
								},
								Mandatory: false, MaxRepeat: 99,
							},
							{ // Segment group 11
								Group: []SchemaNode{
									{Segment: "TAX", Mandatory: true, MaxRepeat: 1},  // Duty/tax/fee details
									{Segment: "ATT", Mandatory: false, MaxRepeat: 9}, // Attribute
									{Segment: "LOC", Mandatory: false, MaxRepeat: 9}, // Place/location identification
									{Segment: "FTX", Mandatory: false, MaxRepeat: 9}, // Free text
								},
								Mandatory: false, MaxRepeat: 9,
							},
							{ // Segment group 12
								Group: []SchemaNode{
									{Segment: "EMP", Mandatory: true, MaxRepeat: 1},  // Employment details
									{Segment: "ATT", Mandatory: false, MaxRepeat: 9}, // Attribute
									{Segment: "LOC", Mandatory: false, MaxRepeat: 9}, // Place/location identification
									{Segment: "DTM", Mandatory: false, MaxRepeat: 9}, // Date/time/period
								},
								Mandatory: false, MaxRepeat: 9,
							},
							{ // Segment group 13
								Group: []SchemaNode{
									{Segment: "ICD", Mandatory: true, MaxRepeat: 1}, // Insurance cover description
									{ // Segment group 14
										Group: []SchemaNode{
											{Segment: "PNA", Mandatory: true, MaxRepeat: 1},  // Party identification
											{Segment: "ADR", Mandatory: false, MaxRepeat: 9}, // Address
											{Segment: "LOC", Mandatory: false, MaxRepeat: 9}, // Place/location identification
										},
										Mandatory: false, MaxRepeat: 9,
									},
								},
								Mandatory: false, MaxRepeat: 9,
							},
						},
						Mandatory: true, MaxRepeat: 9999,
					},
				},
				Mandatory: true, MaxRepeat: 999,
			},
		},
		Mandatory: true, MaxRepeat: 99999,
	},
}}
