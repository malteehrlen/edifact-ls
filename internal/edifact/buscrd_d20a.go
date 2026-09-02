package edifact

// BUSCRD D.20A message specification: the structural (segment/group
// presence, order, cardinality) branching diagram for this message,
// UN/EDIFACT directory release D.20A. Scope is structural only -- not
// element-level content -- per edifact-ls-3uzr's epic non-goal.
// 22 segment groups, max nesting depth 4.
//
// Source: https://service.unece.org/trade/untdid/d20a/trmd/buscrd_c.htm
// section 4.3.1 "Segment table". That URL currently returns HTTP 403 from
// Cloudflare when fetched directly; this data was transcribed from the
// Wayback Machine's archived copy instead:
// http://web.archive.org/web/20230608124149/https://service.unece.org/trade/untdid/d20a/trmd/buscrd_c.htm
// -- re-check the direct URL first if cross-checking later, in case the
// block has lifted.
//
// Transcription note: generated in bulk (edifact-ls-13gu) the same way
// as every other message-type schema -- a script parsed the source's
// exact rail-art column positions mechanically and verified the result
// balances (all 22 groups open and close consistently)
// before being transcribed here, rather than reading the ASCII nesting
// by eye. Not individually spot-checked against the raw source the way
// the first several message types in this project were -- see
// edifact-ls-13gu for which representative samples were spot-checked
// for this batch.

func init() {
	RegisterSchema(
		MessageID{Type: "BUSCRD", Version: "D", Release: "20A", Agency: "UN"},
		buscrdD20aSchema,
		"https://service.unece.org/trade/untdid/d20a/trmd/buscrd_c.htm",
	)
}

var buscrdD20aSchema = Schema{Nodes: []SchemaNode{
	{Segment: "BGM", Mandatory: true, MaxRepeat: 1},   // Beginning of message
	{Segment: "DTM", Mandatory: true, MaxRepeat: 9},   // Date/time/period
	{Segment: "RFF", Mandatory: false, MaxRepeat: 99}, // Reference
	{ // Segment group 1
		Group: []SchemaNode{
			{Segment: "PNA", Mandatory: true, MaxRepeat: 1},  // Party identification
			{Segment: "ADR", Mandatory: false, MaxRepeat: 1}, // Address
			{ // Segment group 2
				Group: []SchemaNode{
					{Segment: "CTA", Mandatory: true, MaxRepeat: 1},  // Contact information
					{Segment: "COM", Mandatory: false, MaxRepeat: 9}, // Communication contact
				},
				Mandatory: false, MaxRepeat: 9,
			},
		},
		Mandatory: true, MaxRepeat: 9,
	},
	{ // Segment group 3
		Group: []SchemaNode{
			{Segment: "LIN", Mandatory: true, MaxRepeat: 1},   // Line item
			{Segment: "DTM", Mandatory: false, MaxRepeat: 99}, // Date/time/period
			{Segment: "PNA", Mandatory: false, MaxRepeat: 9},  // Party identification
			{Segment: "CUX", Mandatory: false, MaxRepeat: 1},  // Currencies
			{Segment: "LAN", Mandatory: false, MaxRepeat: 1},  // Language
			{Segment: "IFD", Mandatory: false, MaxRepeat: 99}, // Information detail
			{Segment: "FTX", Mandatory: false, MaxRepeat: 99}, // Free text
			{ // Segment group 4
				Group: []SchemaNode{
					{Segment: "HYN", Mandatory: true, MaxRepeat: 1},   // Hierarchy information
					{Segment: "IRQ", Mandatory: false, MaxRepeat: 99}, // Information required
					{Segment: "IFD", Mandatory: false, MaxRepeat: 99}, // Information detail
					{Segment: "PCD", Mandatory: false, MaxRepeat: 99}, // Percentage details
					{Segment: "DTM", Mandatory: false, MaxRepeat: 99}, // Date/time/period
					{Segment: "RFF", Mandatory: false, MaxRepeat: 99}, // Reference
					{Segment: "MEA", Mandatory: false, MaxRepeat: 99}, // Measurements
					{Segment: "CUX", Mandatory: false, MaxRepeat: 1},  // Currencies
					{Segment: "FTX", Mandatory: false, MaxRepeat: 99}, // Free text
					{ // Segment group 5
						Group: []SchemaNode{
							{Segment: "MOA", Mandatory: true, MaxRepeat: 1},  // Monetary amount
							{Segment: "ARD", Mandatory: false, MaxRepeat: 9}, // Monetary amount function
							{Segment: "PCD", Mandatory: false, MaxRepeat: 9}, // Percentage details
							{Segment: "FTX", Mandatory: false, MaxRepeat: 1}, // Free text
						},
						Mandatory: false, MaxRepeat: 999,
					},
					{ // Segment group 6
						Group: []SchemaNode{
							{Segment: "ATT", Mandatory: true, MaxRepeat: 1},  // Attribute
							{Segment: "MEA", Mandatory: false, MaxRepeat: 9}, // Measurements
							{Segment: "PCD", Mandatory: false, MaxRepeat: 1}, // Percentage details
						},
						Mandatory: false, MaxRepeat: 999,
					},
					{ // Segment group 7
						Group: []SchemaNode{
							{Segment: "RSL", Mandatory: true, MaxRepeat: 1},   // Result
							{Segment: "CAV", Mandatory: false, MaxRepeat: 99}, // Characteristic value
							{Segment: "FTX", Mandatory: false, MaxRepeat: 99}, // Free text
						},
						Mandatory: false, MaxRepeat: 999,
					},
					{ // Segment group 8
						Group: []SchemaNode{
							{Segment: "RCS", Mandatory: true, MaxRepeat: 1},   // Requirements and conditions
							{Segment: "IFD", Mandatory: false, MaxRepeat: 99}, // Information detail
							{Segment: "MEA", Mandatory: false, MaxRepeat: 9},  // Measurements
							{Segment: "ATT", Mandatory: false, MaxRepeat: 99}, // Attribute
							{ // Segment group 9
								Group: []SchemaNode{
									{Segment: "MOA", Mandatory: true, MaxRepeat: 1},  // Monetary amount
									{Segment: "ARD", Mandatory: false, MaxRepeat: 9}, // Monetary amount function
									{Segment: "FTX", Mandatory: false, MaxRepeat: 1}, // Free text
								},
								Mandatory: false, MaxRepeat: 999,
							},
						},
						Mandatory: false, MaxRepeat: 99999,
					},
					{ // Segment group 10
						Group: []SchemaNode{
							{Segment: "QTY", Mandatory: true, MaxRepeat: 1},  // Quantity
							{Segment: "FTX", Mandatory: false, MaxRepeat: 1}, // Free text
						},
						Mandatory: false, MaxRepeat: 999,
					},
					{ // Segment group 11
						Group: []SchemaNode{
							{Segment: "BUS", Mandatory: true, MaxRepeat: 1},  // Business function
							{Segment: "PYT", Mandatory: false, MaxRepeat: 9}, // Payment terms
							{Segment: "DTM", Mandatory: false, MaxRepeat: 9}, // Date/time/period
							{Segment: "RNG", Mandatory: false, MaxRepeat: 9}, // Range details
							{Segment: "PCD", Mandatory: false, MaxRepeat: 1}, // Percentage details
						},
						Mandatory: false, MaxRepeat: 99,
					},
					{ // Segment group 12
						Group: []SchemaNode{
							{Segment: "GEI", Mandatory: true, MaxRepeat: 1},   // Processing information
							{Segment: "IMD", Mandatory: false, MaxRepeat: 99}, // Item description
							{Segment: "CAV", Mandatory: false, MaxRepeat: 99}, // Characteristic value
							{Segment: "PCD", Mandatory: false, MaxRepeat: 1},  // Percentage details
							{Segment: "EQN", Mandatory: false, MaxRepeat: 1},  // Number of units
							{Segment: "RFF", Mandatory: false, MaxRepeat: 99}, // Reference
							{ // Segment group 13
								Group: []SchemaNode{
									{Segment: "MOA", Mandatory: true, MaxRepeat: 1},  // Monetary amount
									{Segment: "ARD", Mandatory: false, MaxRepeat: 9}, // Monetary amount function
									{Segment: "FTX", Mandatory: false, MaxRepeat: 1}, // Free text
								},
								Mandatory: false, MaxRepeat: 9,
							},
						},
						Mandatory: false, MaxRepeat: 999,
					},
					{ // Segment group 14
						Group: []SchemaNode{
							{Segment: "REL", Mandatory: true, MaxRepeat: 1},   // Relationship
							{Segment: "ATT", Mandatory: false, MaxRepeat: 99}, // Attribute
						},
						Mandatory: false, MaxRepeat: 999,
					},
					{ // Segment group 15
						Group: []SchemaNode{
							{Segment: "STS", Mandatory: true, MaxRepeat: 1},   // Status
							{Segment: "ATT", Mandatory: false, MaxRepeat: 99}, // Attribute
						},
						Mandatory: false, MaxRepeat: 999,
					},
					{ // Segment group 16
						Group: []SchemaNode{
							{Segment: "PNA", Mandatory: true, MaxRepeat: 1},   // Party identification
							{Segment: "REL", Mandatory: false, MaxRepeat: 9},  // Relationship
							{Segment: "ADR", Mandatory: false, MaxRepeat: 99}, // Address
							{Segment: "FII", Mandatory: false, MaxRepeat: 9},  // Financial institution information
							{Segment: "ATT", Mandatory: false, MaxRepeat: 99}, // Attribute
							{Segment: "PDI", Mandatory: false, MaxRepeat: 1},  // Person demographic information
							{Segment: "EMP", Mandatory: false, MaxRepeat: 9},  // Employment details
							{Segment: "NAT", Mandatory: false, MaxRepeat: 9},  // Nationality
							{Segment: "RFF", Mandatory: false, MaxRepeat: 9},  // Reference
							{Segment: "LAN", Mandatory: false, MaxRepeat: 9},  // Language
							{Segment: "DTM", Mandatory: false, MaxRepeat: 99}, // Date/time/period
							{Segment: "IFD", Mandatory: false, MaxRepeat: 99}, // Information detail
							{Segment: "PCD", Mandatory: false, MaxRepeat: 9},  // Percentage details
							{Segment: "COM", Mandatory: false, MaxRepeat: 9},  // Communication contact
							{Segment: "QTY", Mandatory: false, MaxRepeat: 99}, // Quantity
							{Segment: "FTX", Mandatory: false, MaxRepeat: 99}, // Free text
							{ // Segment group 17
								Group: []SchemaNode{
									{Segment: "MOA", Mandatory: true, MaxRepeat: 1},  // Monetary amount
									{Segment: "ARD", Mandatory: false, MaxRepeat: 9}, // Monetary amount function
									{Segment: "FTX", Mandatory: false, MaxRepeat: 1}, // Free text
								},
								Mandatory: false, MaxRepeat: 999,
							},
						},
						Mandatory: false, MaxRepeat: 999,
					},
					{ // Segment group 18
						Group: []SchemaNode{
							{Segment: "EVE", Mandatory: true, MaxRepeat: 1},   // Event
							{Segment: "DTM", Mandatory: false, MaxRepeat: 99}, // Date/time/period
							{Segment: "IFD", Mandatory: false, MaxRepeat: 99}, // Information detail
							{Segment: "FTX", Mandatory: false, MaxRepeat: 99}, // Free text
						},
						Mandatory: false, MaxRepeat: 999,
					},
					{ // Segment group 19
						Group: []SchemaNode{
							{Segment: "LOC", Mandatory: true, MaxRepeat: 1},   // Place/location identification
							{Segment: "REL", Mandatory: false, MaxRepeat: 9},  // Relationship
							{Segment: "ADR", Mandatory: false, MaxRepeat: 9},  // Address
							{Segment: "DTM", Mandatory: false, MaxRepeat: 99}, // Date/time/period
							{Segment: "COM", Mandatory: false, MaxRepeat: 99}, // Communication contact
							{Segment: "IFD", Mandatory: false, MaxRepeat: 99}, // Information detail
							{Segment: "FTX", Mandatory: false, MaxRepeat: 99}, // Free text
							{ // Segment group 20
								Group: []SchemaNode{
									{Segment: "PNA", Mandatory: true, MaxRepeat: 1},   // Party identification
									{Segment: "REL", Mandatory: false, MaxRepeat: 99}, // Relationship
								},
								Mandatory: false, MaxRepeat: 999,
							},
							{ // Segment group 21
								Group: []SchemaNode{
									{Segment: "MOA", Mandatory: true, MaxRepeat: 1},  // Monetary amount
									{Segment: "ARD", Mandatory: false, MaxRepeat: 9}, // Monetary amount function
									{Segment: "FTX", Mandatory: false, MaxRepeat: 1}, // Free text
								},
								Mandatory: false, MaxRepeat: 99,
							},
						},
						Mandatory: false, MaxRepeat: 999,
					},
					{ // Segment group 22
						Group: []SchemaNode{
							{Segment: "APR", Mandatory: true, MaxRepeat: 1},   // Additional price information
							{Segment: "LOC", Mandatory: false, MaxRepeat: 99}, // Place/location identification
							{Segment: "PCD", Mandatory: false, MaxRepeat: 1},  // Percentage details
							{Segment: "FTX", Mandatory: false, MaxRepeat: 99}, // Free text
						},
						Mandatory: false, MaxRepeat: 999,
					},
				},
				Mandatory: false, MaxRepeat: 999999,
			},
		},
		Mandatory: false, MaxRepeat: 999999,
	},
}}
