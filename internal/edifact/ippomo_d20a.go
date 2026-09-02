package edifact

// IPPOMO D.20A message specification: the structural (segment/group
// presence, order, cardinality) branching diagram for this message,
// UN/EDIFACT directory release D.20A. Scope is structural only -- not
// element-level content -- per edifact-ls-3uzr's epic non-goal.
// 37 segment groups, max nesting depth 5.
//
// Source: https://service.unece.org/trade/untdid/d20a/trmd/ippomo_c.htm
// section 4.3.1 "Segment table". That URL currently returns HTTP 403 from
// Cloudflare when fetched directly; this data was transcribed from the
// Wayback Machine's archived copy instead:
// http://web.archive.org/web/20230608114703/https://service.unece.org/trade/untdid/d20a/trmd/ippomo_c.htm
// -- re-check the direct URL first if cross-checking later, in case the
// block has lifted.
//
// Transcription note: generated in bulk (edifact-ls-13gu) the same way
// as every other message-type schema -- a script parsed the source's
// exact rail-art column positions mechanically and verified the result
// balances (all 37 groups open and close consistently)
// before being transcribed here, rather than reading the ASCII nesting
// by eye. Not individually spot-checked against the raw source the way
// the first several message types in this project were -- see
// edifact-ls-13gu for which representative samples were spot-checked
// for this batch.

func init() {
	RegisterSchema(
		MessageID{Type: "IPPOMO", Version: "D", Release: "20A", Agency: "UN"},
		ippomoD20aSchema,
		"https://service.unece.org/trade/untdid/d20a/trmd/ippomo_c.htm",
	)
}

var ippomoD20aSchema = Schema{Nodes: []SchemaNode{
	{Segment: "BGM", Mandatory: true, MaxRepeat: 1},   // Beginning of message
	{Segment: "GEI", Mandatory: false, MaxRepeat: 1},  // Processing information
	{Segment: "ATT", Mandatory: false, MaxRepeat: 9},  // Attribute
	{Segment: "CTA", Mandatory: false, MaxRepeat: 1},  // Contact information
	{Segment: "RFF", Mandatory: false, MaxRepeat: 99}, // Reference
	{Segment: "QTY", Mandatory: false, MaxRepeat: 1},  // Quantity
	{Segment: "SEQ", Mandatory: false, MaxRepeat: 1},  // Sequence details
	{Segment: "DTM", Mandatory: false, MaxRepeat: 9},  // Date/time/period
	{ // Segment group 1
		Group: []SchemaNode{
			{Segment: "PNA", Mandatory: true, MaxRepeat: 1},  // Party identification
			{Segment: "QRS", Mandatory: false, MaxRepeat: 9}, // Query and response
			{ // Segment group 2
				Group: []SchemaNode{
					{Segment: "RFF", Mandatory: true, MaxRepeat: 1},  // Reference
					{Segment: "DTM", Mandatory: false, MaxRepeat: 9}, // Date/time/period
				},
				Mandatory: false, MaxRepeat: 9,
			},
		},
		Mandatory: false, MaxRepeat: 9,
	},
	{ // Segment group 3
		Group: []SchemaNode{
			{Segment: "ICD", Mandatory: true, MaxRepeat: 1},   // Insurance cover description
			{Segment: "RFF", Mandatory: false, MaxRepeat: 9},  // Reference
			{Segment: "DTM", Mandatory: false, MaxRepeat: 99}, // Date/time/period
			{Segment: "ATT", Mandatory: false, MaxRepeat: 9},  // Attribute
			{Segment: "PCD", Mandatory: false, MaxRepeat: 9},  // Percentage details
			{Segment: "MOA", Mandatory: false, MaxRepeat: 9},  // Monetary amount
			{Segment: "QRS", Mandatory: false, MaxRepeat: 99}, // Query and response
			{Segment: "QTY", Mandatory: false, MaxRepeat: 9},  // Quantity
			{Segment: "FTX", Mandatory: false, MaxRepeat: 99}, // Free text
			{ // Segment group 4
				Group: []SchemaNode{
					{Segment: "DOC", Mandatory: true, MaxRepeat: 1},  // Document/message details
					{Segment: "ATT", Mandatory: false, MaxRepeat: 9}, // Attribute
					{Segment: "FTX", Mandatory: false, MaxRepeat: 1}, // Free text
					{ // Segment group 5
						Group: []SchemaNode{
							{Segment: "PCC", Mandatory: true, MaxRepeat: 1},   // Premium calculation component details
							{Segment: "GEI", Mandatory: false, MaxRepeat: 1},  // Processing information
							{Segment: "MOA", Mandatory: false, MaxRepeat: 9},  // Monetary amount
							{Segment: "PCD", Mandatory: false, MaxRepeat: 9},  // Percentage details
							{Segment: "QTY", Mandatory: false, MaxRepeat: 9},  // Quantity
							{Segment: "DTM", Mandatory: false, MaxRepeat: 9},  // Date/time/period
							{Segment: "FTX", Mandatory: false, MaxRepeat: 99}, // Free text
							{ // Segment group 6
								Group: []SchemaNode{
									{Segment: "ATT", Mandatory: true, MaxRepeat: 1},  // Attribute
									{Segment: "PCD", Mandatory: false, MaxRepeat: 1}, // Percentage details
									{Segment: "MOA", Mandatory: false, MaxRepeat: 9}, // Monetary amount
									{Segment: "FTX", Mandatory: false, MaxRepeat: 9}, // Free text
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
		Mandatory: true, MaxRepeat: 9,
	},
	{ // Segment group 7
		Group: []SchemaNode{
			{Segment: "PNA", Mandatory: true, MaxRepeat: 1},   // Party identification
			{Segment: "GEI", Mandatory: false, MaxRepeat: 1},  // Processing information
			{Segment: "ATT", Mandatory: false, MaxRepeat: 9},  // Attribute
			{Segment: "PCD", Mandatory: false, MaxRepeat: 1},  // Percentage details
			{Segment: "RFF", Mandatory: false, MaxRepeat: 9},  // Reference
			{Segment: "EMP", Mandatory: false, MaxRepeat: 99}, // Employment details
			{Segment: "DTM", Mandatory: false, MaxRepeat: 9},  // Date/time/period
			{Segment: "MEM", Mandatory: false, MaxRepeat: 9},  // Membership details
			{Segment: "QTY", Mandatory: false, MaxRepeat: 9},  // Quantity
			{Segment: "REL", Mandatory: false, MaxRepeat: 1},  // Relationship
			{Segment: "LOC", Mandatory: false, MaxRepeat: 1},  // Place/location identification
			{Segment: "QRS", Mandatory: false, MaxRepeat: 99}, // Query and response
			{Segment: "FTX", Mandatory: false, MaxRepeat: 99}, // Free text
			{ // Segment group 8
				Group: []SchemaNode{
					{Segment: "ADR", Mandatory: true, MaxRepeat: 1},  // Address
					{Segment: "DTM", Mandatory: false, MaxRepeat: 1}, // Date/time/period
				},
				Mandatory: false, MaxRepeat: 9,
			},
			{ // Segment group 9
				Group: []SchemaNode{
					{Segment: "CTA", Mandatory: true, MaxRepeat: 1},  // Contact information
					{Segment: "COM", Mandatory: false, MaxRepeat: 9}, // Communication contact
				},
				Mandatory: false, MaxRepeat: 9,
			},
			{ // Segment group 10
				Group: []SchemaNode{
					{Segment: "EVE", Mandatory: true, MaxRepeat: 1},   // Event
					{Segment: "GEI", Mandatory: false, MaxRepeat: 1},  // Processing information
					{Segment: "QRS", Mandatory: false, MaxRepeat: 9},  // Query and response
					{Segment: "RFF", Mandatory: false, MaxRepeat: 9},  // Reference
					{Segment: "ATT", Mandatory: false, MaxRepeat: 9},  // Attribute
					{Segment: "ICD", Mandatory: false, MaxRepeat: 9},  // Insurance cover description
					{Segment: "MOA", Mandatory: false, MaxRepeat: 9},  // Monetary amount
					{Segment: "DTM", Mandatory: false, MaxRepeat: 9},  // Date/time/period
					{Segment: "QTY", Mandatory: false, MaxRepeat: 9},  // Quantity
					{Segment: "PCD", Mandatory: false, MaxRepeat: 9},  // Percentage details
					{Segment: "FTX", Mandatory: false, MaxRepeat: 99}, // Free text
					{ // Segment group 11
						Group: []SchemaNode{
							{Segment: "PNA", Mandatory: true, MaxRepeat: 1},  // Party identification
							{Segment: "ADR", Mandatory: false, MaxRepeat: 1}, // Address
							{Segment: "RFF", Mandatory: false, MaxRepeat: 9}, // Reference
							{Segment: "CTA", Mandatory: false, MaxRepeat: 1}, // Contact information
							{Segment: "COM", Mandatory: false, MaxRepeat: 9}, // Communication contact
						},
						Mandatory: false, MaxRepeat: 9,
					},
				},
				Mandatory: false, MaxRepeat: 99,
			},
			{ // Segment group 12
				Group: []SchemaNode{
					{Segment: "DOC", Mandatory: true, MaxRepeat: 1},  // Document/message details
					{Segment: "DTM", Mandatory: false, MaxRepeat: 9}, // Date/time/period
					{Segment: "ATT", Mandatory: false, MaxRepeat: 9}, // Attribute
					{Segment: "LOC", Mandatory: false, MaxRepeat: 1}, // Place/location identification
					{Segment: "FTX", Mandatory: false, MaxRepeat: 1}, // Free text
				},
				Mandatory: true, MaxRepeat: 9,
			},
		},
		Mandatory: true, MaxRepeat: 999,
	},
	{ // Segment group 13
		Group: []SchemaNode{
			{Segment: "PYT", Mandatory: true, MaxRepeat: 1},  // Payment terms
			{Segment: "DTM", Mandatory: false, MaxRepeat: 9}, // Date/time/period
			{Segment: "ATT", Mandatory: false, MaxRepeat: 9}, // Attribute
			{Segment: "FII", Mandatory: false, MaxRepeat: 1}, // Financial institution information
			{Segment: "RFF", Mandatory: false, MaxRepeat: 9}, // Reference
			{Segment: "PCD", Mandatory: false, MaxRepeat: 9}, // Percentage details
			{ // Segment group 14
				Group: []SchemaNode{
					{Segment: "PNA", Mandatory: true, MaxRepeat: 1},  // Party identification
					{Segment: "ADR", Mandatory: false, MaxRepeat: 1}, // Address
				},
				Mandatory: false, MaxRepeat: 9,
			},
			{ // Segment group 15
				Group: []SchemaNode{
					{Segment: "MOA", Mandatory: true, MaxRepeat: 1},   // Monetary amount
					{Segment: "DTM", Mandatory: false, MaxRepeat: 99}, // Date/time/period
					{Segment: "ATT", Mandatory: false, MaxRepeat: 1},  // Attribute
					{ // Segment group 16
						Group: []SchemaNode{
							{Segment: "TAX", Mandatory: true, MaxRepeat: 1},  // Duty/tax/fee details
							{Segment: "MOA", Mandatory: false, MaxRepeat: 9}, // Monetary amount
							{Segment: "LOC", Mandatory: false, MaxRepeat: 1}, // Place/location identification
							{Segment: "PCD", Mandatory: true, MaxRepeat: 1},  // Percentage details
						},
						Mandatory: false, MaxRepeat: 99,
					},
				},
				Mandatory: false, MaxRepeat: 99,
			},
		},
		Mandatory: false, MaxRepeat: 9,
	},
	{ // Segment group 17
		Group: []SchemaNode{
			{Segment: "ROD", Mandatory: true, MaxRepeat: 1},   // Risk object type
			{Segment: "GEI", Mandatory: false, MaxRepeat: 1},  // Processing information
			{Segment: "RFF", Mandatory: false, MaxRepeat: 9},  // Reference
			{Segment: "DTM", Mandatory: false, MaxRepeat: 9},  // Date/time/period
			{Segment: "ATT", Mandatory: false, MaxRepeat: 99}, // Attribute
			{Segment: "QTY", Mandatory: false, MaxRepeat: 99}, // Quantity
			{Segment: "MOA", Mandatory: false, MaxRepeat: 9},  // Monetary amount
			{Segment: "MEA", Mandatory: false, MaxRepeat: 9},  // Measurements
			{Segment: "ADR", Mandatory: false, MaxRepeat: 99}, // Address
			{Segment: "PCD", Mandatory: false, MaxRepeat: 9},  // Percentage details
			{Segment: "QRS", Mandatory: false, MaxRepeat: 9},  // Query and response
			{Segment: "FTX", Mandatory: false, MaxRepeat: 99}, // Free text
			{ // Segment group 18
				Group: []SchemaNode{
					{Segment: "PNA", Mandatory: true, MaxRepeat: 1},  // Party identification
					{Segment: "GEI", Mandatory: false, MaxRepeat: 1}, // Processing information
					{Segment: "ADR", Mandatory: false, MaxRepeat: 1}, // Address
					{Segment: "FTX", Mandatory: false, MaxRepeat: 9}, // Free text
				},
				Mandatory: false, MaxRepeat: 9,
			},
			{ // Segment group 19
				Group: []SchemaNode{
					{Segment: "EQD", Mandatory: true, MaxRepeat: 1},  // Equipment details
					{Segment: "GEI", Mandatory: false, MaxRepeat: 1}, // Processing information
					{Segment: "ATT", Mandatory: false, MaxRepeat: 9}, // Attribute
					{Segment: "DTM", Mandatory: false, MaxRepeat: 1}, // Date/time/period
					{Segment: "PNA", Mandatory: false, MaxRepeat: 9}, // Party identification
					{Segment: "QRS", Mandatory: false, MaxRepeat: 9}, // Query and response
					{Segment: "RFF", Mandatory: false, MaxRepeat: 9}, // Reference
					{Segment: "FTX", Mandatory: false, MaxRepeat: 9}, // Free text
				},
				Mandatory: false, MaxRepeat: 9,
			},
			{ // Segment group 20
				Group: []SchemaNode{
					{Segment: "DOC", Mandatory: true, MaxRepeat: 1},   // Document/message details
					{Segment: "ATT", Mandatory: false, MaxRepeat: 1},  // Attribute
					{Segment: "DTM", Mandatory: false, MaxRepeat: 99}, // Date/time/period
					{Segment: "PCD", Mandatory: false, MaxRepeat: 1},  // Percentage details
					{Segment: "FTX", Mandatory: false, MaxRepeat: 9},  // Free text
				},
				Mandatory: false, MaxRepeat: 99,
			},
			{ // Segment group 21
				Group: []SchemaNode{
					{Segment: "EVE", Mandatory: true, MaxRepeat: 1},   // Event
					{Segment: "GEI", Mandatory: false, MaxRepeat: 1},  // Processing information
					{Segment: "QRS", Mandatory: false, MaxRepeat: 9},  // Query and response
					{Segment: "RFF", Mandatory: false, MaxRepeat: 9},  // Reference
					{Segment: "ATT", Mandatory: false, MaxRepeat: 9},  // Attribute
					{Segment: "ICD", Mandatory: false, MaxRepeat: 1},  // Insurance cover description
					{Segment: "MOA", Mandatory: true, MaxRepeat: 9},   // Monetary amount
					{Segment: "DTM", Mandatory: false, MaxRepeat: 9},  // Date/time/period
					{Segment: "QTY", Mandatory: false, MaxRepeat: 9},  // Quantity
					{Segment: "PCD", Mandatory: false, MaxRepeat: 9},  // Percentage details
					{Segment: "FTX", Mandatory: false, MaxRepeat: 99}, // Free text
					{ // Segment group 22
						Group: []SchemaNode{
							{Segment: "PNA", Mandatory: true, MaxRepeat: 1},  // Party identification
							{Segment: "ADR", Mandatory: false, MaxRepeat: 1}, // Address
							{Segment: "RFF", Mandatory: false, MaxRepeat: 9}, // Reference
							{Segment: "CTA", Mandatory: false, MaxRepeat: 1}, // Contact information
							{Segment: "COM", Mandatory: false, MaxRepeat: 9}, // Communication contact
						},
						Mandatory: false, MaxRepeat: 9,
					},
				},
				Mandatory: false, MaxRepeat: 99,
			},
			{ // Segment group 23
				Group: []SchemaNode{
					{Segment: "ICD", Mandatory: true, MaxRepeat: 1},  // Insurance cover description
					{Segment: "GEI", Mandatory: false, MaxRepeat: 1}, // Processing information
					{Segment: "DTM", Mandatory: false, MaxRepeat: 9}, // Date/time/period
					{Segment: "RFF", Mandatory: false, MaxRepeat: 9}, // Reference
					{Segment: "ATT", Mandatory: false, MaxRepeat: 9}, // Attribute
					{Segment: "QRS", Mandatory: false, MaxRepeat: 9}, // Query and response
					{Segment: "QTY", Mandatory: false, MaxRepeat: 9}, // Quantity
					{Segment: "FTX", Mandatory: false, MaxRepeat: 9}, // Free text
					{ // Segment group 24
						Group: []SchemaNode{
							{Segment: "MOA", Mandatory: true, MaxRepeat: 1},  // Monetary amount
							{Segment: "PCD", Mandatory: false, MaxRepeat: 1}, // Percentage details
							{Segment: "LOC", Mandatory: false, MaxRepeat: 1}, // Place/location identification
						},
						Mandatory: false, MaxRepeat: 99,
					},
					{ // Segment group 25
						Group: []SchemaNode{
							{Segment: "ROD", Mandatory: true, MaxRepeat: 1},  // Risk object type
							{Segment: "GEI", Mandatory: false, MaxRepeat: 1}, // Processing information
							{Segment: "MOA", Mandatory: false, MaxRepeat: 9}, // Monetary amount
						},
						Mandatory: false, MaxRepeat: 9,
					},
					{ // Segment group 26
						Group: []SchemaNode{
							{Segment: "COD", Mandatory: true, MaxRepeat: 1},  // Component details
							{Segment: "GEI", Mandatory: false, MaxRepeat: 1}, // Processing information
							{ // Segment group 27
								Group: []SchemaNode{
									{Segment: "PCC", Mandatory: true, MaxRepeat: 1},  // Premium calculation component details
									{Segment: "GEI", Mandatory: false, MaxRepeat: 1}, // Processing information
									{Segment: "MOA", Mandatory: false, MaxRepeat: 9}, // Monetary amount
									{Segment: "PCD", Mandatory: false, MaxRepeat: 9}, // Percentage details
									{Segment: "QRS", Mandatory: false, MaxRepeat: 9}, // Query and response
									{Segment: "QTY", Mandatory: false, MaxRepeat: 9}, // Quantity
									{Segment: "DTM", Mandatory: false, MaxRepeat: 9}, // Date/time/period
									{Segment: "FTX", Mandatory: false, MaxRepeat: 9}, // Free text
									{ // Segment group 28
										Group: []SchemaNode{
											{Segment: "PNA", Mandatory: true, MaxRepeat: 1},  // Party identification
											{Segment: "ICD", Mandatory: false, MaxRepeat: 9}, // Insurance cover description
										},
										Mandatory: false, MaxRepeat: 99,
									},
									{ // Segment group 29
										Group: []SchemaNode{
											{Segment: "ATT", Mandatory: true, MaxRepeat: 1},  // Attribute
											{Segment: "PCD", Mandatory: false, MaxRepeat: 1}, // Percentage details
											{Segment: "MOA", Mandatory: false, MaxRepeat: 9}, // Monetary amount
											{Segment: "FTX", Mandatory: false, MaxRepeat: 9}, // Free text
										},
										Mandatory: true, MaxRepeat: 9,
									},
								},
								Mandatory: false, MaxRepeat: 99,
							},
						},
						Mandatory: true, MaxRepeat: 1,
					},
					{ // Segment group 30
						Group: []SchemaNode{
							{Segment: "PNA", Mandatory: true, MaxRepeat: 1},  // Party identification
							{Segment: "GEI", Mandatory: false, MaxRepeat: 1}, // Processing information
							{Segment: "RFF", Mandatory: false, MaxRepeat: 9}, // Reference
							{Segment: "DTM", Mandatory: false, MaxRepeat: 9}, // Date/time/period
							{Segment: "MOA", Mandatory: false, MaxRepeat: 9}, // Monetary amount
							{Segment: "PCD", Mandatory: false, MaxRepeat: 9}, // Percentage details
							{Segment: "QTY", Mandatory: false, MaxRepeat: 9}, // Quantity
							{Segment: "QRS", Mandatory: false, MaxRepeat: 9}, // Query and response
							{Segment: "REL", Mandatory: false, MaxRepeat: 1}, // Relationship
						},
						Mandatory: false, MaxRepeat: 999,
					},
					{ // Segment group 31
						Group: []SchemaNode{
							{Segment: "EVE", Mandatory: true, MaxRepeat: 1},  // Event
							{Segment: "GEI", Mandatory: false, MaxRepeat: 1}, // Processing information
							{Segment: "ATT", Mandatory: false, MaxRepeat: 9}, // Attribute
							{Segment: "DTM", Mandatory: false, MaxRepeat: 9}, // Date/time/period
							{Segment: "MOA", Mandatory: false, MaxRepeat: 9}, // Monetary amount
							{Segment: "PCD", Mandatory: false, MaxRepeat: 9}, // Percentage details
							{Segment: "QRS", Mandatory: false, MaxRepeat: 9}, // Query and response
							{Segment: "QTY", Mandatory: false, MaxRepeat: 9}, // Quantity
							{Segment: "RFF", Mandatory: false, MaxRepeat: 9}, // Reference
							{Segment: "PCC", Mandatory: false, MaxRepeat: 9}, // Premium calculation component details
							{Segment: "LOC", Mandatory: false, MaxRepeat: 9}, // Place/location identification
							{Segment: "FTX", Mandatory: false, MaxRepeat: 9}, // Free text
							{ // Segment group 32
								Group: []SchemaNode{
									{Segment: "PNA", Mandatory: true, MaxRepeat: 1},  // Party identification
									{Segment: "ADR", Mandatory: false, MaxRepeat: 1}, // Address
									{Segment: "RFF", Mandatory: false, MaxRepeat: 9}, // Reference
									{Segment: "CTA", Mandatory: false, MaxRepeat: 1}, // Contact information
									{Segment: "COM", Mandatory: false, MaxRepeat: 9}, // Communication contact
								},
								Mandatory: false, MaxRepeat: 9,
							},
						},
						Mandatory: false, MaxRepeat: 9,
					},
				},
				Mandatory: true, MaxRepeat: 99,
			},
		},
		Mandatory: false, MaxRepeat: 9999,
	},
	{Segment: "UNS", Mandatory: true, MaxRepeat: 1}, // Section control
	{ // Segment group 33
		Group: []SchemaNode{
			{Segment: "DTM", Mandatory: true, MaxRepeat: 1},  // Date/time/period
			{Segment: "RFF", Mandatory: false, MaxRepeat: 9}, // Reference
			{Segment: "QRS", Mandatory: false, MaxRepeat: 9}, // Query and response
			{Segment: "FTX", Mandatory: false, MaxRepeat: 9}, // Free text
			{ // Segment group 34
				Group: []SchemaNode{
					{Segment: "PNA", Mandatory: true, MaxRepeat: 1},  // Party identification
					{Segment: "ADR", Mandatory: false, MaxRepeat: 1}, // Address
					{Segment: "CTA", Mandatory: false, MaxRepeat: 1}, // Contact information
					{Segment: "COM", Mandatory: false, MaxRepeat: 9}, // Communication contact
				},
				Mandatory: false, MaxRepeat: 9,
			},
			{ // Segment group 35
				Group: []SchemaNode{
					{Segment: "GEI", Mandatory: true, MaxRepeat: 1},  // Processing information
					{Segment: "DTM", Mandatory: true, MaxRepeat: 9},  // Date/time/period
					{Segment: "QTY", Mandatory: true, MaxRepeat: 9},  // Quantity
					{Segment: "MOA", Mandatory: false, MaxRepeat: 9}, // Monetary amount
					{ // Segment group 36
						Group: []SchemaNode{
							{Segment: "ICD", Mandatory: true, MaxRepeat: 1},  // Insurance cover description
							{Segment: "ATT", Mandatory: false, MaxRepeat: 9}, // Attribute
							{Segment: "MOA", Mandatory: true, MaxRepeat: 9},  // Monetary amount
							{Segment: "QTY", Mandatory: false, MaxRepeat: 9}, // Quantity
							{Segment: "FTX", Mandatory: false, MaxRepeat: 1}, // Free text
						},
						Mandatory: true, MaxRepeat: 99,
					},
				},
				Mandatory: false, MaxRepeat: 99,
			},
			{ // Segment group 37
				Group: []SchemaNode{
					{Segment: "EVE", Mandatory: true, MaxRepeat: 1},  // Event
					{Segment: "GEI", Mandatory: false, MaxRepeat: 1}, // Processing information
					{Segment: "ATT", Mandatory: false, MaxRepeat: 9}, // Attribute
					{Segment: "DTM", Mandatory: false, MaxRepeat: 9}, // Date/time/period
					{Segment: "ICD", Mandatory: true, MaxRepeat: 1},  // Insurance cover description
					{Segment: "MOA", Mandatory: false, MaxRepeat: 9}, // Monetary amount
					{Segment: "RFF", Mandatory: false, MaxRepeat: 9}, // Reference
					{Segment: "QRS", Mandatory: false, MaxRepeat: 9}, // Query and response
					{Segment: "FTX", Mandatory: false, MaxRepeat: 9}, // Free text
				},
				Mandatory: false, MaxRepeat: 99,
			},
		},
		Mandatory: false, MaxRepeat: 99,
	},
}}
