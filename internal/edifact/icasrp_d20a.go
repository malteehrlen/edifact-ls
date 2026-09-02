package edifact

// ICASRP D.20A message specification: the structural (segment/group
// presence, order, cardinality) branching diagram for this message,
// UN/EDIFACT directory release D.20A. Scope is structural only -- not
// element-level content -- per edifact-ls-3uzr's epic non-goal.
// 25 segment groups, max nesting depth 5.
//
// Source: https://service.unece.org/trade/untdid/d20a/trmd/icasrp_c.htm
// section 4.3.1 "Segment table". That URL currently returns HTTP 403 from
// Cloudflare when fetched directly; this data was transcribed from the
// Wayback Machine's archived copy instead:
// http://web.archive.org/web/20230608112155/https://service.unece.org/trade/untdid/d20a/trmd/icasrp_c.htm
// -- re-check the direct URL first if cross-checking later, in case the
// block has lifted.
//
// Transcription note: generated in bulk (edifact-ls-13gu) the same way
// as every other message-type schema -- a script parsed the source's
// exact rail-art column positions mechanically and verified the result
// balances (all 25 groups open and close consistently)
// before being transcribed here, rather than reading the ASCII nesting
// by eye. Not individually spot-checked against the raw source the way
// the first several message types in this project were -- see
// edifact-ls-13gu for which representative samples were spot-checked
// for this batch.

func init() {
	RegisterSchema(
		MessageID{Type: "ICASRP", Version: "D", Release: "20A", Agency: "UN"},
		icasrpD20aSchema,
		"https://service.unece.org/trade/untdid/d20a/trmd/icasrp_c.htm",
	)
}

var icasrpD20aSchema = Schema{Nodes: []SchemaNode{
	{Segment: "BGM", Mandatory: true, MaxRepeat: 1},  // Beginning of message
	{Segment: "ATT", Mandatory: false, MaxRepeat: 9}, // Attribute
	{Segment: "CUX", Mandatory: false, MaxRepeat: 1}, // Currencies
	{Segment: "DTM", Mandatory: false, MaxRepeat: 9}, // Date/time/period
	{Segment: "QRS", Mandatory: false, MaxRepeat: 9}, // Query and response
	{Segment: "FTX", Mandatory: false, MaxRepeat: 9}, // Free text
	{ // Segment group 1
		Group: []SchemaNode{
			{Segment: "PNA", Mandatory: true, MaxRepeat: 1}, // Party identification
			{ // Segment group 2
				Group: []SchemaNode{
					{Segment: "RFF", Mandatory: true, MaxRepeat: 1},  // Reference
					{Segment: "DTM", Mandatory: false, MaxRepeat: 9}, // Date/time/period
				},
				Mandatory: false, MaxRepeat: 9,
			},
		},
		Mandatory: true, MaxRepeat: 9,
	},
	{ // Segment group 3
		Group: []SchemaNode{
			{Segment: "EVE", Mandatory: true, MaxRepeat: 1},  // Event
			{Segment: "IDE", Mandatory: false, MaxRepeat: 9}, // Identity
			{Segment: "ATT", Mandatory: false, MaxRepeat: 9}, // Attribute
			{Segment: "DTM", Mandatory: false, MaxRepeat: 9}, // Date/time/period
			{Segment: "RFF", Mandatory: false, MaxRepeat: 9}, // Reference
			{Segment: "PCD", Mandatory: false, MaxRepeat: 9}, // Percentage details
			{Segment: "QRS", Mandatory: false, MaxRepeat: 9}, // Query and response
			{Segment: "QTY", Mandatory: false, MaxRepeat: 9}, // Quantity
			{Segment: "FTX", Mandatory: false, MaxRepeat: 9}, // Free text
			{ // Segment group 4
				Group: []SchemaNode{
					{Segment: "TCC", Mandatory: true, MaxRepeat: 1},  // Charge/rate calculations
					{Segment: "IDE", Mandatory: false, MaxRepeat: 9}, // Identity
					{Segment: "ATT", Mandatory: false, MaxRepeat: 9}, // Attribute
					{Segment: "MOA", Mandatory: false, MaxRepeat: 9}, // Monetary amount
					{Segment: "PCD", Mandatory: false, MaxRepeat: 9}, // Percentage details
					{Segment: "FTX", Mandatory: false, MaxRepeat: 9}, // Free text
					{ // Segment group 5
						Group: []SchemaNode{
							{Segment: "COT", Mandatory: true, MaxRepeat: 1},  // Contribution details
							{Segment: "IDE", Mandatory: false, MaxRepeat: 9}, // Identity
							{Segment: "ATT", Mandatory: false, MaxRepeat: 9}, // Attribute
							{Segment: "MOA", Mandatory: false, MaxRepeat: 9}, // Monetary amount
							{Segment: "PCD", Mandatory: false, MaxRepeat: 9}, // Percentage details
							{Segment: "RFF", Mandatory: false, MaxRepeat: 9}, // Reference
							{Segment: "FTX", Mandatory: false, MaxRepeat: 9}, // Free text
						},
						Mandatory: false, MaxRepeat: 99,
					},
				},
				Mandatory: false, MaxRepeat: 9,
			},
			{ // Segment group 6
				Group: []SchemaNode{
					{Segment: "COD", Mandatory: true, MaxRepeat: 1},  // Component details
					{Segment: "IDE", Mandatory: false, MaxRepeat: 9}, // Identity
					{Segment: "ADR", Mandatory: false, MaxRepeat: 1}, // Address
					{Segment: "ATT", Mandatory: false, MaxRepeat: 9}, // Attribute
					{Segment: "COM", Mandatory: false, MaxRepeat: 9}, // Communication contact
					{Segment: "DTM", Mandatory: false, MaxRepeat: 9}, // Date/time/period
					{Segment: "MOA", Mandatory: false, MaxRepeat: 9}, // Monetary amount
					{Segment: "PCD", Mandatory: false, MaxRepeat: 9}, // Percentage details
					{Segment: "QRS", Mandatory: false, MaxRepeat: 9}, // Query and response
					{Segment: "QTY", Mandatory: false, MaxRepeat: 9}, // Quantity
					{Segment: "RFF", Mandatory: false, MaxRepeat: 9}, // Reference
					{Segment: "FTX", Mandatory: false, MaxRepeat: 9}, // Free text
					{ // Segment group 7
						Group: []SchemaNode{
							{Segment: "PCC", Mandatory: true, MaxRepeat: 1},  // Premium calculation component details
							{Segment: "ATT", Mandatory: false, MaxRepeat: 9}, // Attribute
							{Segment: "DTM", Mandatory: false, MaxRepeat: 9}, // Date/time/period
							{Segment: "MOA", Mandatory: false, MaxRepeat: 9}, // Monetary amount
							{Segment: "PCD", Mandatory: false, MaxRepeat: 9}, // Percentage details
							{Segment: "QRS", Mandatory: false, MaxRepeat: 9}, // Query and response
							{Segment: "QTY", Mandatory: false, MaxRepeat: 9}, // Quantity
							{Segment: "FTX", Mandatory: false, MaxRepeat: 9}, // Free text
							{ // Segment group 8
								Group: []SchemaNode{
									{Segment: "TCC", Mandatory: true, MaxRepeat: 1},  // Charge/rate calculations
									{Segment: "IDE", Mandatory: false, MaxRepeat: 9}, // Identity
									{Segment: "ATT", Mandatory: false, MaxRepeat: 9}, // Attribute
									{Segment: "MOA", Mandatory: false, MaxRepeat: 9}, // Monetary amount
									{Segment: "PCD", Mandatory: false, MaxRepeat: 9}, // Percentage details
									{Segment: "FTX", Mandatory: false, MaxRepeat: 9}, // Free text
									{ // Segment group 9
										Group: []SchemaNode{
											{Segment: "COT", Mandatory: true, MaxRepeat: 1},  // Contribution details
											{Segment: "IDE", Mandatory: false, MaxRepeat: 9}, // Identity
											{Segment: "ATT", Mandatory: false, MaxRepeat: 9}, // Attribute
											{Segment: "MOA", Mandatory: false, MaxRepeat: 9}, // Monetary amount
											{Segment: "PCD", Mandatory: false, MaxRepeat: 9}, // Percentage details
											{Segment: "RFF", Mandatory: false, MaxRepeat: 9}, // Reference
											{Segment: "FTX", Mandatory: false, MaxRepeat: 9}, // Free text
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
				Mandatory: false, MaxRepeat: 99,
			},
		},
		Mandatory: true, MaxRepeat: 99,
	},
	{ // Segment group 10
		Group: []SchemaNode{
			{Segment: "ICD", Mandatory: true, MaxRepeat: 1},  // Insurance cover description
			{Segment: "IDE", Mandatory: false, MaxRepeat: 9}, // Identity
			{Segment: "ATT", Mandatory: false, MaxRepeat: 9}, // Attribute
			{Segment: "DTM", Mandatory: false, MaxRepeat: 9}, // Date/time/period
			{Segment: "MOA", Mandatory: false, MaxRepeat: 9}, // Monetary amount
			{Segment: "RFF", Mandatory: false, MaxRepeat: 9}, // Reference
			{Segment: "QRS", Mandatory: false, MaxRepeat: 9}, // Query and response
			{Segment: "FTX", Mandatory: false, MaxRepeat: 9}, // Free text
		},
		Mandatory: false, MaxRepeat: 999,
	},
	{ // Segment group 11
		Group: []SchemaNode{
			{Segment: "DOC", Mandatory: true, MaxRepeat: 1},  // Document/message details
			{Segment: "IDE", Mandatory: false, MaxRepeat: 9}, // Identity
			{Segment: "ATT", Mandatory: false, MaxRepeat: 9}, // Attribute
			{Segment: "ADR", Mandatory: false, MaxRepeat: 9}, // Address
			{Segment: "EFI", Mandatory: false, MaxRepeat: 1}, // External file link identification
			{Segment: "RFF", Mandatory: false, MaxRepeat: 9}, // Reference
			{Segment: "DTM", Mandatory: false, MaxRepeat: 9}, // Date/time/period
			{Segment: "MOA", Mandatory: false, MaxRepeat: 9}, // Monetary amount
			{Segment: "QRS", Mandatory: false, MaxRepeat: 9}, // Query and response
			{Segment: "QTY", Mandatory: false, MaxRepeat: 9}, // Quantity
			{Segment: "FTX", Mandatory: false, MaxRepeat: 9}, // Free text
		},
		Mandatory: false, MaxRepeat: 99,
	},
	{ // Segment group 12
		Group: []SchemaNode{
			{Segment: "PNA", Mandatory: true, MaxRepeat: 1},   // Party identification
			{Segment: "IDE", Mandatory: false, MaxRepeat: 9},  // Identity
			{Segment: "ATT", Mandatory: false, MaxRepeat: 99}, // Attribute
			{Segment: "DTM", Mandatory: false, MaxRepeat: 9},  // Date/time/period
			{Segment: "MEA", Mandatory: false, MaxRepeat: 9},  // Measurements
			{Segment: "RFF", Mandatory: false, MaxRepeat: 9},  // Reference
			{Segment: "QRS", Mandatory: false, MaxRepeat: 9},  // Query and response
			{Segment: "PCD", Mandatory: false, MaxRepeat: 9},  // Percentage details
			{Segment: "FTX", Mandatory: false, MaxRepeat: 9},  // Free text
			{ // Segment group 13
				Group: []SchemaNode{
					{Segment: "COD", Mandatory: true, MaxRepeat: 1},  // Component details
					{Segment: "IDE", Mandatory: false, MaxRepeat: 9}, // Identity
					{Segment: "ADR", Mandatory: false, MaxRepeat: 1}, // Address
					{Segment: "FII", Mandatory: false, MaxRepeat: 1}, // Financial institution information
					{Segment: "ATT", Mandatory: false, MaxRepeat: 9}, // Attribute
					{Segment: "COM", Mandatory: false, MaxRepeat: 9}, // Communication contact
					{Segment: "DTM", Mandatory: false, MaxRepeat: 9}, // Date/time/period
					{Segment: "EMP", Mandatory: false, MaxRepeat: 9}, // Employment details
					{Segment: "MOA", Mandatory: false, MaxRepeat: 9}, // Monetary amount
					{Segment: "QRS", Mandatory: false, MaxRepeat: 9}, // Query and response
					{Segment: "QTY", Mandatory: false, MaxRepeat: 9}, // Quantity
					{Segment: "RFF", Mandatory: false, MaxRepeat: 9}, // Reference
					{Segment: "FTX", Mandatory: false, MaxRepeat: 9}, // Free text
				},
				Mandatory: false, MaxRepeat: 99,
			},
		},
		Mandatory: true, MaxRepeat: 99,
	},
	{ // Segment group 14
		Group: []SchemaNode{
			{Segment: "ROD", Mandatory: true, MaxRepeat: 1},   // Risk object type
			{Segment: "IDE", Mandatory: false, MaxRepeat: 9},  // Identity
			{Segment: "RFF", Mandatory: false, MaxRepeat: 9},  // Reference
			{Segment: "MOA", Mandatory: false, MaxRepeat: 99}, // Monetary amount
			{Segment: "PCD", Mandatory: false, MaxRepeat: 9},  // Percentage details
			{Segment: "TAX", Mandatory: false, MaxRepeat: 1},  // Duty/tax/fee details
			{Segment: "DTM", Mandatory: false, MaxRepeat: 99}, // Date/time/period
			{Segment: "LOC", Mandatory: false, MaxRepeat: 9},  // Place/location identification
			{Segment: "ATT", Mandatory: false, MaxRepeat: 99}, // Attribute
			{Segment: "MEA", Mandatory: false, MaxRepeat: 9},  // Measurements
			{Segment: "QTY", Mandatory: false, MaxRepeat: 9},  // Quantity
			{Segment: "QRS", Mandatory: false, MaxRepeat: 99}, // Query and response
			{Segment: "RTE", Mandatory: false, MaxRepeat: 9},  // Rate details
			{Segment: "FTX", Mandatory: false, MaxRepeat: 9},  // Free text
			{ // Segment group 15
				Group: []SchemaNode{
					{Segment: "TCC", Mandatory: true, MaxRepeat: 1},  // Charge/rate calculations
					{Segment: "IDE", Mandatory: false, MaxRepeat: 9}, // Identity
					{Segment: "ATT", Mandatory: false, MaxRepeat: 9}, // Attribute
					{Segment: "MOA", Mandatory: false, MaxRepeat: 9}, // Monetary amount
					{Segment: "PCD", Mandatory: false, MaxRepeat: 9}, // Percentage details
					{Segment: "FTX", Mandatory: false, MaxRepeat: 9}, // Free text
					{ // Segment group 16
						Group: []SchemaNode{
							{Segment: "COT", Mandatory: true, MaxRepeat: 1},  // Contribution details
							{Segment: "IDE", Mandatory: false, MaxRepeat: 9}, // Identity
							{Segment: "ATT", Mandatory: false, MaxRepeat: 9}, // Attribute
							{Segment: "MOA", Mandatory: false, MaxRepeat: 9}, // Monetary amount
							{Segment: "PCD", Mandatory: false, MaxRepeat: 9}, // Percentage details
							{Segment: "RFF", Mandatory: false, MaxRepeat: 9}, // Reference
							{Segment: "FTX", Mandatory: false, MaxRepeat: 9}, // Free text
						},
						Mandatory: false, MaxRepeat: 99,
					},
				},
				Mandatory: false, MaxRepeat: 99,
			},
			{ // Segment group 17
				Group: []SchemaNode{
					{Segment: "COD", Mandatory: true, MaxRepeat: 1},  // Component details
					{Segment: "IDE", Mandatory: false, MaxRepeat: 9}, // Identity
					{Segment: "ADR", Mandatory: false, MaxRepeat: 1}, // Address
					{Segment: "ATT", Mandatory: false, MaxRepeat: 9}, // Attribute
					{Segment: "COM", Mandatory: false, MaxRepeat: 9}, // Communication contact
					{Segment: "DAM", Mandatory: false, MaxRepeat: 9}, // Damage
					{Segment: "DTM", Mandatory: false, MaxRepeat: 9}, // Date/time/period
					{Segment: "MEA", Mandatory: false, MaxRepeat: 9}, // Measurements
					{Segment: "MOA", Mandatory: false, MaxRepeat: 9}, // Monetary amount
					{Segment: "PCD", Mandatory: false, MaxRepeat: 9}, // Percentage details
					{Segment: "QRS", Mandatory: false, MaxRepeat: 9}, // Query and response
					{Segment: "QTY", Mandatory: false, MaxRepeat: 9}, // Quantity
					{Segment: "RFF", Mandatory: false, MaxRepeat: 9}, // Reference
					{Segment: "FTX", Mandatory: false, MaxRepeat: 9}, // Free text
					{ // Segment group 18
						Group: []SchemaNode{
							{Segment: "TCC", Mandatory: true, MaxRepeat: 1},  // Charge/rate calculations
							{Segment: "IDE", Mandatory: false, MaxRepeat: 9}, // Identity
							{Segment: "ATT", Mandatory: false, MaxRepeat: 9}, // Attribute
							{Segment: "MOA", Mandatory: false, MaxRepeat: 9}, // Monetary amount
							{Segment: "PCD", Mandatory: false, MaxRepeat: 9}, // Percentage details
							{Segment: "FTX", Mandatory: false, MaxRepeat: 9}, // Free text
							{ // Segment group 19
								Group: []SchemaNode{
									{Segment: "COT", Mandatory: true, MaxRepeat: 1},  // Contribution details
									{Segment: "IDE", Mandatory: false, MaxRepeat: 9}, // Identity
									{Segment: "ATT", Mandatory: false, MaxRepeat: 9}, // Attribute
									{Segment: "MOA", Mandatory: false, MaxRepeat: 9}, // Monetary amount
									{Segment: "PCD", Mandatory: false, MaxRepeat: 5}, // Percentage details
									{Segment: "RFF", Mandatory: false, MaxRepeat: 9}, // Reference
									{Segment: "FTX", Mandatory: false, MaxRepeat: 9}, // Free text
								},
								Mandatory: false, MaxRepeat: 99,
							},
						},
						Mandatory: false, MaxRepeat: 99,
					},
					{ // Segment group 20
						Group: []SchemaNode{
							{Segment: "IMD", Mandatory: true, MaxRepeat: 1},  // Item description
							{Segment: "IDE", Mandatory: false, MaxRepeat: 9}, // Identity
							{Segment: "ATT", Mandatory: false, MaxRepeat: 9}, // Attribute
							{Segment: "DTM", Mandatory: false, MaxRepeat: 9}, // Date/time/period
							{Segment: "MEA", Mandatory: false, MaxRepeat: 9}, // Measurements
							{Segment: "PCD", Mandatory: false, MaxRepeat: 9}, // Percentage details
							{Segment: "QRS", Mandatory: false, MaxRepeat: 9}, // Query and response
							{Segment: "QTY", Mandatory: false, MaxRepeat: 9}, // Quantity
							{Segment: "RFF", Mandatory: false, MaxRepeat: 9}, // Reference
							{Segment: "FTX", Mandatory: false, MaxRepeat: 9}, // Free text
							{ // Segment group 21
								Group: []SchemaNode{
									{Segment: "TCC", Mandatory: true, MaxRepeat: 1},  // Charge/rate calculations
									{Segment: "IDE", Mandatory: false, MaxRepeat: 9}, // Identity
									{Segment: "ATT", Mandatory: false, MaxRepeat: 9}, // Attribute
									{Segment: "MOA", Mandatory: false, MaxRepeat: 9}, // Monetary amount
									{Segment: "PCD", Mandatory: false, MaxRepeat: 9}, // Percentage details
									{Segment: "FTX", Mandatory: false, MaxRepeat: 9}, // Free text
									{ // Segment group 22
										Group: []SchemaNode{
											{Segment: "COT", Mandatory: true, MaxRepeat: 1},  // Contribution details
											{Segment: "IDE", Mandatory: false, MaxRepeat: 9}, // Identity
											{Segment: "ATT", Mandatory: false, MaxRepeat: 9}, // Attribute
											{Segment: "MOA", Mandatory: false, MaxRepeat: 9}, // Monetary amount
											{Segment: "PCD", Mandatory: false, MaxRepeat: 9}, // Percentage details
											{Segment: "RFF", Mandatory: false, MaxRepeat: 9}, // Reference
											{Segment: "FTX", Mandatory: false, MaxRepeat: 9}, // Free text
										},
										Mandatory: false, MaxRepeat: 99,
									},
								},
								Mandatory: false, MaxRepeat: 9,
							},
						},
						Mandatory: false, MaxRepeat: 99,
					},
				},
				Mandatory: false, MaxRepeat: 999,
			},
		},
		Mandatory: true, MaxRepeat: 99,
	},
	{ // Segment group 23
		Group: []SchemaNode{
			{Segment: "PRC", Mandatory: true, MaxRepeat: 1},   // Process identification
			{Segment: "IDE", Mandatory: false, MaxRepeat: 9},  // Identity
			{Segment: "ATT", Mandatory: false, MaxRepeat: 9},  // Attribute
			{Segment: "DTM", Mandatory: false, MaxRepeat: 99}, // Date/time/period
			{Segment: "MOA", Mandatory: false, MaxRepeat: 9},  // Monetary amount
			{Segment: "QRS", Mandatory: false, MaxRepeat: 99}, // Query and response
			{Segment: "RFF", Mandatory: false, MaxRepeat: 9},  // Reference
			{Segment: "FTX", Mandatory: false, MaxRepeat: 9},  // Free text
		},
		Mandatory: false, MaxRepeat: 99,
	},
	{ // Segment group 24
		Group: []SchemaNode{
			{Segment: "SEQ", Mandatory: true, MaxRepeat: 1},  // Sequence details
			{Segment: "IDE", Mandatory: false, MaxRepeat: 9}, // Identity
			{Segment: "ATT", Mandatory: false, MaxRepeat: 9}, // Attribute
			{Segment: "DTM", Mandatory: false, MaxRepeat: 9}, // Date/time/period
			{Segment: "RFF", Mandatory: false, MaxRepeat: 9}, // Reference
			{ // Segment group 25
				Group: []SchemaNode{
					{Segment: "COD", Mandatory: true, MaxRepeat: 1},  // Component details
					{Segment: "IDE", Mandatory: false, MaxRepeat: 9}, // Identity
					{Segment: "ATT", Mandatory: false, MaxRepeat: 9}, // Attribute
					{Segment: "DTM", Mandatory: false, MaxRepeat: 9}, // Date/time/period
					{Segment: "MOA", Mandatory: false, MaxRepeat: 9}, // Monetary amount
					{Segment: "QRS", Mandatory: false, MaxRepeat: 9}, // Query and response
					{Segment: "RFF", Mandatory: false, MaxRepeat: 9}, // Reference
					{Segment: "FTX", Mandatory: false, MaxRepeat: 9}, // Free text
				},
				Mandatory: false, MaxRepeat: 99,
			},
		},
		Mandatory: false, MaxRepeat: 99,
	},
}}
