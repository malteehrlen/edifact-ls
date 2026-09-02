package edifact

// SANCRT D.20A message specification: the structural (segment/group
// presence, order, cardinality) branching diagram for this message,
// UN/EDIFACT directory release D.20A. Scope is structural only -- not
// element-level content -- per edifact-ls-3uzr's epic non-goal.
// 21 segment groups, max nesting depth 4.
//
// Source: https://service.unece.org/trade/untdid/d20a/trmd/sancrt_c.htm
// section 4.3.1 "Segment table". That URL currently returns HTTP 403 from
// Cloudflare when fetched directly; this data was transcribed from the
// Wayback Machine's archived copy instead:
// http://web.archive.org/web/20231002210905/https://service.unece.org/trade/untdid/d20a/trmd/sancrt_c.htm
// -- re-check the direct URL first if cross-checking later, in case the
// block has lifted.
//
// Transcription note: generated in bulk (edifact-ls-13gu) the same way
// as every other message-type schema -- a script parsed the source's
// exact rail-art column positions mechanically and verified the result
// balances (all 21 groups open and close consistently)
// before being transcribed here, rather than reading the ASCII nesting
// by eye. Not individually spot-checked against the raw source the way
// the first several message types in this project were -- see
// edifact-ls-13gu for which representative samples were spot-checked
// for this batch.

func init() {
	RegisterSchema(
		MessageID{Type: "SANCRT", Version: "D", Release: "20A", Agency: "UN"},
		sancrtD20aSchema,
		"https://service.unece.org/trade/untdid/d20a/trmd/sancrt_c.htm",
	)
}

var sancrtD20aSchema = Schema{Nodes: []SchemaNode{
	{Segment: "BGM", Mandatory: true, MaxRepeat: 1},   // Beginning of message
	{Segment: "DTM", Mandatory: false, MaxRepeat: 99}, // Date/time/period
	{Segment: "STS", Mandatory: false, MaxRepeat: 99}, // Status
	{Segment: "LOC", Mandatory: false, MaxRepeat: 99}, // Place/location identification
	{Segment: "RFF", Mandatory: false, MaxRepeat: 99}, // Reference
	{Segment: "FTX", Mandatory: false, MaxRepeat: 99}, // Free text
	{Segment: "MEA", Mandatory: false, MaxRepeat: 99}, // Measurements
	{Segment: "MOA", Mandatory: false, MaxRepeat: 99}, // Monetary amount
	{Segment: "GEI", Mandatory: false, MaxRepeat: 99}, // Processing information
	{Segment: "CST", Mandatory: false, MaxRepeat: 1},  // Customs status of goods
	{ // Segment group 1
		Group: []SchemaNode{
			{Segment: "DOC", Mandatory: true, MaxRepeat: 1},  // Document/message details
			{Segment: "DTM", Mandatory: false, MaxRepeat: 9}, // Date/time/period
			{Segment: "LOC", Mandatory: false, MaxRepeat: 9}, // Place/location identification
		},
		Mandatory: false, MaxRepeat: 99,
	},
	{ // Segment group 2
		Group: []SchemaNode{
			{Segment: "PNA", Mandatory: true, MaxRepeat: 1},  // Party identification
			{Segment: "ADR", Mandatory: false, MaxRepeat: 9}, // Address
			{ // Segment group 3
				Group: []SchemaNode{
					{Segment: "CTA", Mandatory: true, MaxRepeat: 1},  // Contact information
					{Segment: "COM", Mandatory: false, MaxRepeat: 9}, // Communication contact
				},
				Mandatory: false, MaxRepeat: 9,
			},
		},
		Mandatory: false, MaxRepeat: 99,
	},
	{ // Segment group 4
		Group: []SchemaNode{
			{Segment: "TDT", Mandatory: true, MaxRepeat: 1},  // Transport information
			{Segment: "DTM", Mandatory: false, MaxRepeat: 9}, // Date/time/period
			{Segment: "LOC", Mandatory: false, MaxRepeat: 9}, // Place/location identification
			{Segment: "RFF", Mandatory: false, MaxRepeat: 9}, // Reference
		},
		Mandatory: false, MaxRepeat: 99,
	},
	{ // Segment group 5
		Group: []SchemaNode{
			{Segment: "PAC", Mandatory: true, MaxRepeat: 1},  // Package
			{Segment: "PCI", Mandatory: false, MaxRepeat: 9}, // Package identification
			{Segment: "MEA", Mandatory: false, MaxRepeat: 9}, // Measurements
		},
		Mandatory: false, MaxRepeat: 99,
	},
	{ // Segment group 6
		Group: []SchemaNode{
			{Segment: "EQD", Mandatory: true, MaxRepeat: 1},  // Equipment details
			{Segment: "TMP", Mandatory: false, MaxRepeat: 9}, // Temperature
			{ // Segment group 7
				Group: []SchemaNode{
					{Segment: "SEL", Mandatory: true, MaxRepeat: 1},  // Seal number
					{Segment: "DTM", Mandatory: false, MaxRepeat: 9}, // Date/time/period
					{Segment: "LOC", Mandatory: false, MaxRepeat: 9}, // Place/location identification
				},
				Mandatory: false, MaxRepeat: 99,
			},
		},
		Mandatory: false, MaxRepeat: 99,
	},
	{ // Segment group 8
		Group: []SchemaNode{
			{Segment: "PRC", Mandatory: true, MaxRepeat: 1},   // Process identification
			{Segment: "IMD", Mandatory: false, MaxRepeat: 9},  // Item description
			{Segment: "MEA", Mandatory: false, MaxRepeat: 9},  // Measurements
			{Segment: "DTM", Mandatory: false, MaxRepeat: 9},  // Date/time/period
			{Segment: "LOC", Mandatory: false, MaxRepeat: 9},  // Place/location identification
			{Segment: "DOC", Mandatory: false, MaxRepeat: 9},  // Document/message details
			{Segment: "RFF", Mandatory: false, MaxRepeat: 9},  // Reference
			{Segment: "TMP", Mandatory: false, MaxRepeat: 9},  // Temperature
			{Segment: "GEI", Mandatory: false, MaxRepeat: 99}, // Processing information
			{Segment: "FTX", Mandatory: false, MaxRepeat: 99}, // Free text
			{ // Segment group 9
				Group: []SchemaNode{
					{Segment: "PNA", Mandatory: true, MaxRepeat: 1},  // Party identification
					{Segment: "ADR", Mandatory: false, MaxRepeat: 9}, // Address
					{ // Segment group 10
						Group: []SchemaNode{
							{Segment: "CTA", Mandatory: true, MaxRepeat: 1},  // Contact information
							{Segment: "COM", Mandatory: false, MaxRepeat: 9}, // Communication contact
						},
						Mandatory: false, MaxRepeat: 9,
					},
				},
				Mandatory: false, MaxRepeat: 99,
			},
		},
		Mandatory: false, MaxRepeat: 99,
	},
	{ // Segment group 11
		Group: []SchemaNode{
			{Segment: "LIN", Mandatory: true, MaxRepeat: 1},     // Line item
			{Segment: "CST", Mandatory: false, MaxRepeat: 9},    // Customs status of goods
			{Segment: "MEA", Mandatory: false, MaxRepeat: 9},    // Measurements
			{Segment: "PIA", Mandatory: false, MaxRepeat: 9},    // Additional product id
			{Segment: "IMD", Mandatory: false, MaxRepeat: 9},    // Item description
			{Segment: "GIN", Mandatory: false, MaxRepeat: 9999}, // Goods identity number
			{Segment: "RFF", Mandatory: false, MaxRepeat: 9},    // Reference
			{Segment: "ATT", Mandatory: false, MaxRepeat: 9},    // Attribute
			{Segment: "DTM", Mandatory: false, MaxRepeat: 9},    // Date/time/period
			{Segment: "LOC", Mandatory: false, MaxRepeat: 9},    // Place/location identification
			{Segment: "FTX", Mandatory: false, MaxRepeat: 9},    // Free text
			{Segment: "QTY", Mandatory: false, MaxRepeat: 9},    // Quantity
			{Segment: "MOA", Mandatory: false, MaxRepeat: 9},    // Monetary amount
			{ // Segment group 12
				Group: []SchemaNode{
					{Segment: "DOC", Mandatory: true, MaxRepeat: 1},  // Document/message details
					{Segment: "DTM", Mandatory: false, MaxRepeat: 9}, // Date/time/period
					{Segment: "LOC", Mandatory: false, MaxRepeat: 9}, // Place/location identification
				},
				Mandatory: false, MaxRepeat: 99,
			},
			{ // Segment group 13
				Group: []SchemaNode{
					{Segment: "PNA", Mandatory: true, MaxRepeat: 1},  // Party identification
					{Segment: "ADR", Mandatory: false, MaxRepeat: 9}, // Address
					{ // Segment group 14
						Group: []SchemaNode{
							{Segment: "CTA", Mandatory: true, MaxRepeat: 1},  // Contact information
							{Segment: "COM", Mandatory: false, MaxRepeat: 9}, // Communication contact
						},
						Mandatory: false, MaxRepeat: 9,
					},
				},
				Mandatory: false, MaxRepeat: 99,
			},
			{ // Segment group 15
				Group: []SchemaNode{
					{Segment: "PAC", Mandatory: true, MaxRepeat: 1},  // Package
					{Segment: "PCI", Mandatory: false, MaxRepeat: 9}, // Package identification
					{Segment: "MEA", Mandatory: false, MaxRepeat: 9}, // Measurements
				},
				Mandatory: false, MaxRepeat: 99,
			},
			{ // Segment group 16
				Group: []SchemaNode{
					{Segment: "EQD", Mandatory: true, MaxRepeat: 1},  // Equipment details
					{Segment: "TMP", Mandatory: false, MaxRepeat: 9}, // Temperature
					{ // Segment group 17
						Group: []SchemaNode{
							{Segment: "SEL", Mandatory: true, MaxRepeat: 1},  // Seal number
							{Segment: "DTM", Mandatory: false, MaxRepeat: 9}, // Date/time/period
							{Segment: "LOC", Mandatory: false, MaxRepeat: 9}, // Place/location identification
						},
						Mandatory: false, MaxRepeat: 99,
					},
				},
				Mandatory: false, MaxRepeat: 99,
			},
			{ // Segment group 18
				Group: []SchemaNode{
					{Segment: "PRC", Mandatory: true, MaxRepeat: 1},   // Process identification
					{Segment: "IMD", Mandatory: false, MaxRepeat: 9},  // Item description
					{Segment: "MEA", Mandatory: false, MaxRepeat: 9},  // Measurements
					{Segment: "DTM", Mandatory: false, MaxRepeat: 9},  // Date/time/period
					{Segment: "LOC", Mandatory: false, MaxRepeat: 9},  // Place/location identification
					{Segment: "DOC", Mandatory: false, MaxRepeat: 9},  // Document/message details
					{Segment: "RFF", Mandatory: false, MaxRepeat: 9},  // Reference
					{Segment: "TMP", Mandatory: false, MaxRepeat: 9},  // Temperature
					{Segment: "GEI", Mandatory: false, MaxRepeat: 99}, // Processing information
					{Segment: "FTX", Mandatory: false, MaxRepeat: 99}, // Free text
					{ // Segment group 19
						Group: []SchemaNode{
							{Segment: "PNA", Mandatory: true, MaxRepeat: 1},  // Party identification
							{Segment: "ADR", Mandatory: false, MaxRepeat: 9}, // Address
							{ // Segment group 20
								Group: []SchemaNode{
									{Segment: "CTA", Mandatory: true, MaxRepeat: 1},  // Contact information
									{Segment: "COM", Mandatory: false, MaxRepeat: 9}, // Communication contact
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
		Mandatory: false, MaxRepeat: 9999,
	},
	{Segment: "CNT", Mandatory: false, MaxRepeat: 9}, // Control total
	{ // Segment group 21
		Group: []SchemaNode{
			{Segment: "AUT", Mandatory: true, MaxRepeat: 1},  // Authentication result
			{Segment: "DTM", Mandatory: false, MaxRepeat: 9}, // Date/time/period
		},
		Mandatory: false, MaxRepeat: 99,
	},
}}
