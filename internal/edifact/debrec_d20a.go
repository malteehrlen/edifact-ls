package edifact

// DEBREC D.20A message specification: the structural (segment/group
// presence, order, cardinality) branching diagram for this message,
// UN/EDIFACT directory release D.20A. Scope is structural only -- not
// element-level content -- per edifact-ls-3uzr's epic non-goal.
// 26 segment groups, max nesting depth 5.
//
// Source: https://service.unece.org/trade/untdid/d20a/trmd/debrec_c.htm
// section 4.3.1 "Segment table". That URL currently returns HTTP 403 from
// Cloudflare when fetched directly; this data was transcribed from the
// Wayback Machine's archived copy instead:
// http://web.archive.org/web/20230608112419/https://service.unece.org/trade/untdid/d20a/trmd/debrec_c.htm
// -- re-check the direct URL first if cross-checking later, in case the
// block has lifted.
//
// Transcription note: generated in bulk (edifact-ls-13gu) the same way
// as every other message-type schema -- a script parsed the source's
// exact rail-art column positions mechanically and verified the result
// balances (all 26 groups open and close consistently)
// before being transcribed here, rather than reading the ASCII nesting
// by eye. Not individually spot-checked against the raw source the way
// the first several message types in this project were -- see
// edifact-ls-13gu for which representative samples were spot-checked
// for this batch.

func init() {
	RegisterSchema(
		MessageID{Type: "DEBREC", Version: "D", Release: "20A", Agency: "UN"},
		debrecD20aSchema,
		"https://service.unece.org/trade/untdid/d20a/trmd/debrec_c.htm",
	)
}

var debrecD20aSchema = Schema{Nodes: []SchemaNode{
	{Segment: "BGM", Mandatory: true, MaxRepeat: 1},  // Beginning of message
	{Segment: "DTM", Mandatory: false, MaxRepeat: 5}, // Date/time/period
	{ // Segment group 1
		Group: []SchemaNode{
			{Segment: "PNA", Mandatory: true, MaxRepeat: 1},  // Party identification
			{Segment: "ADR", Mandatory: false, MaxRepeat: 1}, // Address
			{Segment: "RFF", Mandatory: false, MaxRepeat: 5}, // Reference
			{Segment: "EMP", Mandatory: false, MaxRepeat: 1}, // Employment details
			{Segment: "COM", Mandatory: false, MaxRepeat: 9}, // Communication contact
			{ // Segment group 2
				Group: []SchemaNode{
					{Segment: "CTA", Mandatory: true, MaxRepeat: 1},  // Contact information
					{Segment: "DTM", Mandatory: false, MaxRepeat: 1}, // Date/time/period
					{ // Segment group 3
						Group: []SchemaNode{
							{Segment: "COM", Mandatory: true, MaxRepeat: 1},  // Communication contact
							{Segment: "DTM", Mandatory: false, MaxRepeat: 9}, // Date/time/period
						},
						Mandatory: false, MaxRepeat: 5,
					},
				},
				Mandatory: false, MaxRepeat: 9,
			},
			{ // Segment group 4
				Group: []SchemaNode{
					{Segment: "AUT", Mandatory: true, MaxRepeat: 1},  // Authentication result
					{Segment: "DTM", Mandatory: false, MaxRepeat: 1}, // Date/time/period
				},
				Mandatory: false, MaxRepeat: 1,
			},
		},
		Mandatory: false, MaxRepeat: 99,
	},
	{ // Segment group 5
		Group: []SchemaNode{
			{Segment: "GIR", Mandatory: true, MaxRepeat: 1}, // Related identification numbers
			{ // Segment group 6
				Group: []SchemaNode{
					{Segment: "ATT", Mandatory: true, MaxRepeat: 1},  // Attribute
					{Segment: "DTM", Mandatory: false, MaxRepeat: 3}, // Date/time/period
				},
				Mandatory: false, MaxRepeat: 1,
			},
			{ // Segment group 7
				Group: []SchemaNode{
					{Segment: "EVE", Mandatory: true, MaxRepeat: 1},   // Event
					{Segment: "DTM", Mandatory: false, MaxRepeat: 9},  // Date/time/period
					{Segment: "IRQ", Mandatory: false, MaxRepeat: 99}, // Information required
					{Segment: "FTX", Mandatory: false, MaxRepeat: 99}, // Free text
					{ // Segment group 8
						Group: []SchemaNode{
							{Segment: "ATT", Mandatory: true, MaxRepeat: 1},  // Attribute
							{Segment: "DTM", Mandatory: false, MaxRepeat: 3}, // Date/time/period
						},
						Mandatory: false, MaxRepeat: 9,
					},
					{ // Segment group 9
						Group: []SchemaNode{
							{Segment: "PNA", Mandatory: true, MaxRepeat: 1},  // Party identification
							{Segment: "RFF", Mandatory: false, MaxRepeat: 9}, // Reference
							{Segment: "ATT", Mandatory: false, MaxRepeat: 9}, // Attribute
							{Segment: "EMP", Mandatory: false, MaxRepeat: 1}, // Employment details
							{Segment: "PDI", Mandatory: false, MaxRepeat: 1}, // Person demographic information
							{Segment: "FII", Mandatory: false, MaxRepeat: 9}, // Financial institution information
							{Segment: "DTM", Mandatory: false, MaxRepeat: 9}, // Date/time/period
							{Segment: "MOA", Mandatory: false, MaxRepeat: 5}, // Monetary amount
							{Segment: "QTY", Mandatory: false, MaxRepeat: 9}, // Quantity
							{Segment: "PCD", Mandatory: false, MaxRepeat: 1}, // Percentage details
							{Segment: "COM", Mandatory: false, MaxRepeat: 1}, // Communication contact
							{ // Segment group 10
								Group: []SchemaNode{
									{Segment: "CTA", Mandatory: true, MaxRepeat: 1},  // Contact information
									{Segment: "DTM", Mandatory: false, MaxRepeat: 1}, // Date/time/period
									{Segment: "FTX", Mandatory: false, MaxRepeat: 9}, // Free text
									{ // Segment group 11
										Group: []SchemaNode{
											{Segment: "COM", Mandatory: true, MaxRepeat: 1},  // Communication contact
											{Segment: "DTM", Mandatory: false, MaxRepeat: 9}, // Date/time/period
										},
										Mandatory: false, MaxRepeat: 5,
									},
								},
								Mandatory: false, MaxRepeat: 9,
							},
							{ // Segment group 12
								Group: []SchemaNode{
									{Segment: "ADR", Mandatory: true, MaxRepeat: 1},  // Address
									{Segment: "DTM", Mandatory: false, MaxRepeat: 1}, // Date/time/period
									{ // Segment group 13
										Group: []SchemaNode{
											{Segment: "COM", Mandatory: true, MaxRepeat: 1},  // Communication contact
											{Segment: "DTM", Mandatory: false, MaxRepeat: 9}, // Date/time/period
										},
										Mandatory: false, MaxRepeat: 1,
									},
								},
								Mandatory: false, MaxRepeat: 5,
							},
							{ // Segment group 14
								Group: []SchemaNode{
									{Segment: "REL", Mandatory: true, MaxRepeat: 1},  // Relationship
									{Segment: "RFF", Mandatory: false, MaxRepeat: 1}, // Reference
								},
								Mandatory: false, MaxRepeat: 1,
							},
						},
						Mandatory: false, MaxRepeat: 99,
					},
					{ // Segment group 15
						Group: []SchemaNode{
							{Segment: "DOC", Mandatory: true, MaxRepeat: 1},   // Document/message details
							{Segment: "DTM", Mandatory: false, MaxRepeat: 9},  // Date/time/period
							{Segment: "FTX", Mandatory: false, MaxRepeat: 99}, // Free text
							{Segment: "ATT", Mandatory: false, MaxRepeat: 9},  // Attribute
							{Segment: "MOA", Mandatory: false, MaxRepeat: 9},  // Monetary amount
							{ // Segment group 16
								Group: []SchemaNode{
									{Segment: "PNA", Mandatory: true, MaxRepeat: 1},  // Party identification
									{Segment: "RFF", Mandatory: false, MaxRepeat: 9}, // Reference
									{Segment: "ADR", Mandatory: false, MaxRepeat: 5}, // Address
								},
								Mandatory: false, MaxRepeat: 9,
							},
						},
						Mandatory: false, MaxRepeat: 99,
					},
					{ // Segment group 17
						Group: []SchemaNode{
							{Segment: "IMD", Mandatory: true, MaxRepeat: 1},  // Item description
							{Segment: "STS", Mandatory: false, MaxRepeat: 1}, // Status
							{Segment: "FTX", Mandatory: false, MaxRepeat: 9}, // Free text
							{Segment: "ATT", Mandatory: false, MaxRepeat: 1}, // Attribute
							{ // Segment group 18
								Group: []SchemaNode{
									{Segment: "RFF", Mandatory: true, MaxRepeat: 1},   // Reference
									{Segment: "PIA", Mandatory: false, MaxRepeat: 1},  // Additional product id
									{Segment: "LOC", Mandatory: false, MaxRepeat: 2},  // Place/location identification
									{Segment: "DTM", Mandatory: false, MaxRepeat: 9},  // Date/time/period
									{Segment: "MEA", Mandatory: false, MaxRepeat: 9},  // Measurements
									{Segment: "ATT", Mandatory: false, MaxRepeat: 99}, // Attribute
									{Segment: "FTX", Mandatory: false, MaxRepeat: 1},  // Free text
								},
								Mandatory: false, MaxRepeat: 1,
							},
							{ // Segment group 19
								Group: []SchemaNode{
									{Segment: "ADR", Mandatory: true, MaxRepeat: 1}, // Address
									{ // Segment group 20
										Group: []SchemaNode{
											{Segment: "RFF", Mandatory: true, MaxRepeat: 1},  // Reference
											{Segment: "FTX", Mandatory: false, MaxRepeat: 9}, // Free text
											{Segment: "MEA", Mandatory: false, MaxRepeat: 1}, // Measurements
										},
										Mandatory: false, MaxRepeat: 9,
									},
								},
								Mandatory: false, MaxRepeat: 1,
							},
							{ // Segment group 21
								Group: []SchemaNode{
									{Segment: "DOC", Mandatory: true, MaxRepeat: 1},  // Document/message details
									{Segment: "DTM", Mandatory: false, MaxRepeat: 9}, // Date/time/period
									{Segment: "FTX", Mandatory: false, MaxRepeat: 9}, // Free text
								},
								Mandatory: false, MaxRepeat: 1,
							},
						},
						Mandatory: false, MaxRepeat: 999,
					},
					{ // Segment group 22
						Group: []SchemaNode{
							{Segment: "GEI", Mandatory: true, MaxRepeat: 1},  // Processing information
							{Segment: "DTM", Mandatory: false, MaxRepeat: 3}, // Date/time/period
							{Segment: "FTX", Mandatory: false, MaxRepeat: 9}, // Free text
							{ // Segment group 23
								Group: []SchemaNode{
									{Segment: "DOC", Mandatory: true, MaxRepeat: 1},  // Document/message details
									{Segment: "DTM", Mandatory: false, MaxRepeat: 9}, // Date/time/period
									{Segment: "FTX", Mandatory: false, MaxRepeat: 9}, // Free text
								},
								Mandatory: false, MaxRepeat: 1,
							},
							{ // Segment group 24
								Group: []SchemaNode{
									{Segment: "ARD", Mandatory: true, MaxRepeat: 1},  // Monetary amount function
									{Segment: "MOA", Mandatory: false, MaxRepeat: 2}, // Monetary amount
									{Segment: "DTM", Mandatory: false, MaxRepeat: 5}, // Date/time/period
									{Segment: "TAX", Mandatory: false, MaxRepeat: 9}, // Duty/tax/fee details
									{Segment: "FTX", Mandatory: false, MaxRepeat: 9}, // Free text
									{Segment: "ATT", Mandatory: false, MaxRepeat: 9}, // Attribute
									{ // Segment group 25
										Group: []SchemaNode{
											{Segment: "PCD", Mandatory: true, MaxRepeat: 1},  // Percentage details
											{Segment: "DTM", Mandatory: false, MaxRepeat: 1}, // Date/time/period
										},
										Mandatory: false, MaxRepeat: 9,
									},
									{ // Segment group 26
										Group: []SchemaNode{
											{Segment: "PAI", Mandatory: true, MaxRepeat: 1},  // Payment instructions
											{Segment: "RFF", Mandatory: false, MaxRepeat: 9}, // Reference
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
				Mandatory: false, MaxRepeat: 999,
			},
		},
		Mandatory: false, MaxRepeat: 9999,
	},
	{Segment: "CNT", Mandatory: false, MaxRepeat: 99}, // Control total
}}
