package edifact

// IPPOAD D.20A message specification: the structural (segment/group
// presence, order, cardinality) branching diagram for this message,
// UN/EDIFACT directory release D.20A. Scope is structural only -- not
// element-level content -- per edifact-ls-3uzr's epic non-goal.
// 20 segment groups, max nesting depth 3.
//
// Source: https://service.unece.org/trade/untdid/d20a/trmd/ippoad_c.htm
// section 4.3.1 "Segment table". That URL currently returns HTTP 403 from
// Cloudflare when fetched directly; this data was transcribed from the
// Wayback Machine's archived copy instead:
// http://web.archive.org/web/20230608114744/https://service.unece.org/trade/untdid/d20a/trmd/ippoad_c.htm
// -- re-check the direct URL first if cross-checking later, in case the
// block has lifted.
//
// Transcription note: generated in bulk (edifact-ls-13gu) the same way
// as every other message-type schema -- a script parsed the source's
// exact rail-art column positions mechanically and verified the result
// balances (all 20 groups open and close consistently)
// before being transcribed here, rather than reading the ASCII nesting
// by eye. Not individually spot-checked against the raw source the way
// the first several message types in this project were -- see
// edifact-ls-13gu for which representative samples were spot-checked
// for this batch.

func init() {
	RegisterSchema(
		MessageID{Type: "IPPOAD", Version: "D", Release: "20A", Agency: "UN"},
		ippoadD20aSchema,
		"https://service.unece.org/trade/untdid/d20a/trmd/ippoad_c.htm",
	)
}

var ippoadD20aSchema = Schema{Nodes: []SchemaNode{
	{Segment: "BGM", Mandatory: true, MaxRepeat: 1},  // Beginning of message
	{Segment: "ATT", Mandatory: false, MaxRepeat: 9}, // Attribute
	{Segment: "RFF", Mandatory: false, MaxRepeat: 9}, // Reference
	{Segment: "DTM", Mandatory: false, MaxRepeat: 9}, // Date/time/period
	{Segment: "QTY", Mandatory: false, MaxRepeat: 9}, // Quantity
	{Segment: "SEQ", Mandatory: false, MaxRepeat: 1}, // Sequence details
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
		Mandatory: false, MaxRepeat: 9,
	},
	{ // Segment group 3
		Group: []SchemaNode{
			{Segment: "ICD", Mandatory: true, MaxRepeat: 1},   // Insurance cover description
			{Segment: "IDE", Mandatory: true, MaxRepeat: 9},   // Identity
			{Segment: "GEI", Mandatory: false, MaxRepeat: 1},  // Processing information
			{Segment: "ATT", Mandatory: false, MaxRepeat: 99}, // Attribute
			{Segment: "CUX", Mandatory: false, MaxRepeat: 1},  // Currencies
			{Segment: "DTM", Mandatory: false, MaxRepeat: 99}, // Date/time/period
			{Segment: "MOA", Mandatory: false, MaxRepeat: 9},  // Monetary amount
			{Segment: "PCD", Mandatory: false, MaxRepeat: 9},  // Percentage details
			{Segment: "QTY", Mandatory: false, MaxRepeat: 9},  // Quantity
			{Segment: "QRS", Mandatory: false, MaxRepeat: 99}, // Query and response
			{Segment: "RFF", Mandatory: false, MaxRepeat: 9},  // Reference
			{Segment: "FTX", Mandatory: false, MaxRepeat: 99}, // Free text
			{ // Segment group 4
				Group: []SchemaNode{
					{Segment: "COD", Mandatory: true, MaxRepeat: 1},   // Component details
					{Segment: "IDE", Mandatory: false, MaxRepeat: 99}, // Identity
					{Segment: "GEI", Mandatory: false, MaxRepeat: 1},  // Processing information
					{Segment: "ATT", Mandatory: false, MaxRepeat: 99}, // Attribute
					{Segment: "COM", Mandatory: false, MaxRepeat: 9},  // Communication contact
					{Segment: "DTM", Mandatory: false, MaxRepeat: 99}, // Date/time/period
					{Segment: "EFI", Mandatory: false, MaxRepeat: 1},  // External file link identification
					{Segment: "MOA", Mandatory: false, MaxRepeat: 9},  // Monetary amount
					{Segment: "PCD", Mandatory: false, MaxRepeat: 9},  // Percentage details
					{Segment: "QRS", Mandatory: false, MaxRepeat: 99}, // Query and response
					{Segment: "QTY", Mandatory: false, MaxRepeat: 9},  // Quantity
					{Segment: "RFF", Mandatory: false, MaxRepeat: 9},  // Reference
					{Segment: "FTX", Mandatory: false, MaxRepeat: 9},  // Free text
					{ // Segment group 5
						Group: []SchemaNode{
							{Segment: "PCC", Mandatory: true, MaxRepeat: 1},   // Premium calculation component details
							{Segment: "IDE", Mandatory: false, MaxRepeat: 9},  // Identity
							{Segment: "ADR", Mandatory: false, MaxRepeat: 99}, // Address
							{Segment: "ATT", Mandatory: false, MaxRepeat: 9},  // Attribute
							{Segment: "DTM", Mandatory: false, MaxRepeat: 9},  // Date/time/period
							{Segment: "MOA", Mandatory: false, MaxRepeat: 9},  // Monetary amount
							{Segment: "PCD", Mandatory: false, MaxRepeat: 9},  // Percentage details
							{Segment: "QRS", Mandatory: false, MaxRepeat: 9},  // Query and response
							{Segment: "QTY", Mandatory: false, MaxRepeat: 9},  // Quantity
							{Segment: "FTX", Mandatory: false, MaxRepeat: 9},  // Free text
						},
						Mandatory: false, MaxRepeat: 99,
					},
				},
				Mandatory: false, MaxRepeat: 999,
			},
		},
		Mandatory: true, MaxRepeat: 999,
	},
	{ // Segment group 6
		Group: []SchemaNode{
			{Segment: "PNA", Mandatory: true, MaxRepeat: 1},   // Party identification
			{Segment: "IDE", Mandatory: false, MaxRepeat: 99}, // Identity
			{Segment: "GEI", Mandatory: false, MaxRepeat: 1},  // Processing information
			{Segment: "ATT", Mandatory: false, MaxRepeat: 99}, // Attribute
			{Segment: "COM", Mandatory: false, MaxRepeat: 9},  // Communication contact
			{Segment: "DTM", Mandatory: false, MaxRepeat: 9},  // Date/time/period
			{Segment: "LOC", Mandatory: false, MaxRepeat: 9},  // Place/location identification
			{Segment: "MOA", Mandatory: false, MaxRepeat: 9},  // Monetary amount
			{Segment: "PCD", Mandatory: false, MaxRepeat: 9},  // Percentage details
			{Segment: "QRS", Mandatory: false, MaxRepeat: 99}, // Query and response
			{Segment: "QTY", Mandatory: false, MaxRepeat: 9},  // Quantity
			{Segment: "RFF", Mandatory: false, MaxRepeat: 9},  // Reference
			{Segment: "FTX", Mandatory: false, MaxRepeat: 99}, // Free text
			{ // Segment group 7
				Group: []SchemaNode{
					{Segment: "CTA", Mandatory: true, MaxRepeat: 1},  // Contact information
					{Segment: "COM", Mandatory: false, MaxRepeat: 9}, // Communication contact
				},
				Mandatory: false, MaxRepeat: 9,
			},
			{ // Segment group 8
				Group: []SchemaNode{
					{Segment: "COD", Mandatory: true, MaxRepeat: 1},   // Component details
					{Segment: "IDE", Mandatory: false, MaxRepeat: 9},  // Identity
					{Segment: "GEI", Mandatory: false, MaxRepeat: 1},  // Processing information
					{Segment: "ADR", Mandatory: false, MaxRepeat: 9},  // Address
					{Segment: "ATT", Mandatory: false, MaxRepeat: 9},  // Attribute
					{Segment: "COM", Mandatory: false, MaxRepeat: 9},  // Communication contact
					{Segment: "DTM", Mandatory: false, MaxRepeat: 9},  // Date/time/period
					{Segment: "EFI", Mandatory: false, MaxRepeat: 1},  // External file link identification
					{Segment: "EMP", Mandatory: false, MaxRepeat: 1},  // Employment details
					{Segment: "FII", Mandatory: false, MaxRepeat: 1},  // Financial institution information
					{Segment: "IMD", Mandatory: false, MaxRepeat: 9},  // Item description
					{Segment: "MOA", Mandatory: false, MaxRepeat: 9},  // Monetary amount
					{Segment: "PCD", Mandatory: false, MaxRepeat: 9},  // Percentage details
					{Segment: "QRS", Mandatory: false, MaxRepeat: 99}, // Query and response
					{Segment: "QTY", Mandatory: false, MaxRepeat: 9},  // Quantity
					{Segment: "RFF", Mandatory: false, MaxRepeat: 9},  // Reference
					{Segment: "FTX", Mandatory: false, MaxRepeat: 9},  // Free text
					{ // Segment group 9
						Group: []SchemaNode{
							{Segment: "PCC", Mandatory: true, MaxRepeat: 1},   // Premium calculation component details
							{Segment: "IDE", Mandatory: false, MaxRepeat: 9},  // Identity
							{Segment: "ADR", Mandatory: false, MaxRepeat: 99}, // Address
							{Segment: "ATT", Mandatory: false, MaxRepeat: 9},  // Attribute
							{Segment: "DTM", Mandatory: false, MaxRepeat: 9},  // Date/time/period
							{Segment: "MOA", Mandatory: false, MaxRepeat: 9},  // Monetary amount
							{Segment: "PCD", Mandatory: false, MaxRepeat: 9},  // Percentage details
							{Segment: "QRS", Mandatory: false, MaxRepeat: 9},  // Query and response
							{Segment: "QTY", Mandatory: false, MaxRepeat: 9},  // Quantity
							{Segment: "FTX", Mandatory: false, MaxRepeat: 9},  // Free text
						},
						Mandatory: false, MaxRepeat: 99,
					},
				},
				Mandatory: false, MaxRepeat: 999,
			},
		},
		Mandatory: false, MaxRepeat: 9999,
	},
	{ // Segment group 10
		Group: []SchemaNode{
			{Segment: "ROD", Mandatory: true, MaxRepeat: 1},   // Risk object type
			{Segment: "IDE", Mandatory: false, MaxRepeat: 99}, // Identity
			{Segment: "GEI", Mandatory: false, MaxRepeat: 1},  // Processing information
			{Segment: "ATT", Mandatory: false, MaxRepeat: 99}, // Attribute
			{Segment: "COM", Mandatory: false, MaxRepeat: 9},  // Communication contact
			{Segment: "DTM", Mandatory: false, MaxRepeat: 99}, // Date/time/period
			{Segment: "MEA", Mandatory: false, MaxRepeat: 99}, // Measurements
			{Segment: "MOA", Mandatory: false, MaxRepeat: 9},  // Monetary amount
			{Segment: "PCD", Mandatory: false, MaxRepeat: 9},  // Percentage details
			{Segment: "QRS", Mandatory: false, MaxRepeat: 99}, // Query and response
			{Segment: "QTY", Mandatory: false, MaxRepeat: 99}, // Quantity
			{Segment: "RFF", Mandatory: false, MaxRepeat: 9},  // Reference
			{Segment: "FTX", Mandatory: false, MaxRepeat: 99}, // Free text
			{ // Segment group 11
				Group: []SchemaNode{
					{Segment: "COD", Mandatory: true, MaxRepeat: 1},   // Component details
					{Segment: "IDE", Mandatory: false, MaxRepeat: 9},  // Identity
					{Segment: "GEI", Mandatory: false, MaxRepeat: 1},  // Processing information
					{Segment: "ADR", Mandatory: false, MaxRepeat: 9},  // Address
					{Segment: "ATT", Mandatory: false, MaxRepeat: 99}, // Attribute
					{Segment: "COM", Mandatory: false, MaxRepeat: 9},  // Communication contact
					{Segment: "DTM", Mandatory: false, MaxRepeat: 9},  // Date/time/period
					{Segment: "EFI", Mandatory: false, MaxRepeat: 1},  // External file link identification
					{Segment: "EMP", Mandatory: false, MaxRepeat: 9},  // Employment details
					{Segment: "IMD", Mandatory: false, MaxRepeat: 9},  // Item description
					{Segment: "MOA", Mandatory: false, MaxRepeat: 9},  // Monetary amount
					{Segment: "MEA", Mandatory: false, MaxRepeat: 9},  // Measurements
					{Segment: "PCD", Mandatory: false, MaxRepeat: 9},  // Percentage details
					{Segment: "QRS", Mandatory: false, MaxRepeat: 99}, // Query and response
					{Segment: "QTY", Mandatory: false, MaxRepeat: 9},  // Quantity
					{Segment: "RFF", Mandatory: false, MaxRepeat: 9},  // Reference
					{Segment: "FTX", Mandatory: false, MaxRepeat: 99}, // Free text
					{ // Segment group 12
						Group: []SchemaNode{
							{Segment: "PCC", Mandatory: true, MaxRepeat: 1},   // Premium calculation component details
							{Segment: "IDE", Mandatory: false, MaxRepeat: 9},  // Identity
							{Segment: "ADR", Mandatory: false, MaxRepeat: 99}, // Address
							{Segment: "ATT", Mandatory: false, MaxRepeat: 9},  // Attribute
							{Segment: "DTM", Mandatory: false, MaxRepeat: 9},  // Date/time/period
							{Segment: "MOA", Mandatory: false, MaxRepeat: 9},  // Monetary amount
							{Segment: "PCD", Mandatory: false, MaxRepeat: 9},  // Percentage details
							{Segment: "QRS", Mandatory: false, MaxRepeat: 9},  // Query and response
							{Segment: "QTY", Mandatory: false, MaxRepeat: 9},  // Quantity
							{Segment: "FTX", Mandatory: false, MaxRepeat: 9},  // Free text
						},
						Mandatory: false, MaxRepeat: 99,
					},
				},
				Mandatory: false, MaxRepeat: 999,
			},
		},
		Mandatory: false, MaxRepeat: 9999,
	},
	{ // Segment group 13
		Group: []SchemaNode{
			{Segment: "EVE", Mandatory: true, MaxRepeat: 1},   // Event
			{Segment: "IDE", Mandatory: false, MaxRepeat: 99}, // Identity
			{Segment: "GEI", Mandatory: false, MaxRepeat: 1},  // Processing information
			{Segment: "ATT", Mandatory: false, MaxRepeat: 99}, // Attribute
			{Segment: "DTM", Mandatory: false, MaxRepeat: 9},  // Date/time/period
			{Segment: "MEA", Mandatory: false, MaxRepeat: 9},  // Measurements
			{Segment: "MOA", Mandatory: false, MaxRepeat: 9},  // Monetary amount
			{Segment: "PCD", Mandatory: false, MaxRepeat: 9},  // Percentage details
			{Segment: "QRS", Mandatory: false, MaxRepeat: 99}, // Query and response
			{Segment: "QTY", Mandatory: false, MaxRepeat: 99}, // Quantity
			{Segment: "RFF", Mandatory: false, MaxRepeat: 9},  // Reference
			{Segment: "FTX", Mandatory: false, MaxRepeat: 99}, // Free text
			{ // Segment group 14
				Group: []SchemaNode{
					{Segment: "COD", Mandatory: true, MaxRepeat: 1},   // Component details
					{Segment: "IDE", Mandatory: false, MaxRepeat: 99}, // Identity
					{Segment: "GEI", Mandatory: false, MaxRepeat: 1},  // Processing information
					{Segment: "ADR", Mandatory: false, MaxRepeat: 9},  // Address
					{Segment: "ATT", Mandatory: false, MaxRepeat: 99}, // Attribute
					{Segment: "COM", Mandatory: false, MaxRepeat: 9},  // Communication contact
					{Segment: "DTM", Mandatory: false, MaxRepeat: 9},  // Date/time/period
					{Segment: "EFI", Mandatory: false, MaxRepeat: 1},  // External file link identification
					{Segment: "EMP", Mandatory: false, MaxRepeat: 9},  // Employment details
					{Segment: "MOA", Mandatory: false, MaxRepeat: 9},  // Monetary amount
					{Segment: "PCD", Mandatory: false, MaxRepeat: 9},  // Percentage details
					{Segment: "QRS", Mandatory: false, MaxRepeat: 99}, // Query and response
					{Segment: "QTY", Mandatory: false, MaxRepeat: 9},  // Quantity
					{Segment: "RFF", Mandatory: false, MaxRepeat: 9},  // Reference
					{Segment: "FTX", Mandatory: false, MaxRepeat: 99}, // Free text
					{ // Segment group 15
						Group: []SchemaNode{
							{Segment: "PCC", Mandatory: true, MaxRepeat: 1},   // Premium calculation component details
							{Segment: "IDE", Mandatory: false, MaxRepeat: 9},  // Identity
							{Segment: "ADR", Mandatory: false, MaxRepeat: 99}, // Address
							{Segment: "ATT", Mandatory: false, MaxRepeat: 9},  // Attribute
							{Segment: "DTM", Mandatory: false, MaxRepeat: 9},  // Date/time/period
							{Segment: "MOA", Mandatory: false, MaxRepeat: 9},  // Monetary amount
							{Segment: "PCD", Mandatory: false, MaxRepeat: 9},  // Percentage details
							{Segment: "QRS", Mandatory: false, MaxRepeat: 9},  // Query and response
							{Segment: "QTY", Mandatory: false, MaxRepeat: 9},  // Quantity
							{Segment: "FTX", Mandatory: false, MaxRepeat: 9},  // Free text
						},
						Mandatory: false, MaxRepeat: 99,
					},
				},
				Mandatory: false, MaxRepeat: 999,
			},
		},
		Mandatory: false, MaxRepeat: 9999,
	},
	{ // Segment group 16
		Group: []SchemaNode{
			{Segment: "SEQ", Mandatory: true, MaxRepeat: 1},  // Sequence details
			{Segment: "IDE", Mandatory: false, MaxRepeat: 9}, // Identity
			{Segment: "GEI", Mandatory: false, MaxRepeat: 9}, // Processing information
			{Segment: "ATT", Mandatory: false, MaxRepeat: 9}, // Attribute
			{Segment: "DTM", Mandatory: false, MaxRepeat: 9}, // Date/time/period
			{Segment: "MOA", Mandatory: false, MaxRepeat: 9}, // Monetary amount
			{Segment: "PCD", Mandatory: false, MaxRepeat: 9}, // Percentage details
			{Segment: "PNA", Mandatory: false, MaxRepeat: 9}, // Party identification
			{Segment: "QRS", Mandatory: false, MaxRepeat: 9}, // Query and response
			{Segment: "QTY", Mandatory: false, MaxRepeat: 9}, // Quantity
			{Segment: "RFF", Mandatory: false, MaxRepeat: 9}, // Reference
			{Segment: "FTX", Mandatory: false, MaxRepeat: 9}, // Free text
			{ // Segment group 17
				Group: []SchemaNode{
					{Segment: "COD", Mandatory: true, MaxRepeat: 1},   // Component details
					{Segment: "IDE", Mandatory: false, MaxRepeat: 99}, // Identity
					{Segment: "GEI", Mandatory: false, MaxRepeat: 9},  // Processing information
					{Segment: "ATT", Mandatory: false, MaxRepeat: 9},  // Attribute
					{Segment: "DTM", Mandatory: false, MaxRepeat: 9},  // Date/time/period
					{Segment: "MOA", Mandatory: false, MaxRepeat: 9},  // Monetary amount
					{Segment: "PCD", Mandatory: false, MaxRepeat: 9},  // Percentage details
					{Segment: "PNA", Mandatory: false, MaxRepeat: 9},  // Party identification
					{Segment: "QRS", Mandatory: false, MaxRepeat: 9},  // Query and response
					{Segment: "QTY", Mandatory: false, MaxRepeat: 9},  // Quantity
					{Segment: "RFF", Mandatory: false, MaxRepeat: 9},  // Reference
					{Segment: "FTX", Mandatory: false, MaxRepeat: 9},  // Free text
					{ // Segment group 18
						Group: []SchemaNode{
							{Segment: "IMD", Mandatory: true, MaxRepeat: 1},  // Item description
							{Segment: "IDE", Mandatory: false, MaxRepeat: 9}, // Identity
							{Segment: "DTM", Mandatory: false, MaxRepeat: 9}, // Date/time/period
							{Segment: "PNA", Mandatory: false, MaxRepeat: 9}, // Party identification
							{Segment: "QRS", Mandatory: false, MaxRepeat: 9}, // Query and response
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
	{ // Segment group 19
		Group: []SchemaNode{
			{Segment: "PYT", Mandatory: true, MaxRepeat: 1},  // Payment terms
			{Segment: "IDE", Mandatory: false, MaxRepeat: 9}, // Identity
			{Segment: "GEI", Mandatory: false, MaxRepeat: 1}, // Processing information
			{Segment: "ATT", Mandatory: false, MaxRepeat: 9}, // Attribute
			{Segment: "DTM", Mandatory: false, MaxRepeat: 9}, // Date/time/period
			{Segment: "MOA", Mandatory: false, MaxRepeat: 9}, // Monetary amount
			{Segment: "PCD", Mandatory: false, MaxRepeat: 9}, // Percentage details
			{Segment: "QRS", Mandatory: false, MaxRepeat: 9}, // Query and response
			{Segment: "QTY", Mandatory: false, MaxRepeat: 9}, // Quantity
			{Segment: "RFF", Mandatory: false, MaxRepeat: 9}, // Reference
			{Segment: "FTX", Mandatory: false, MaxRepeat: 9}, // Free text
			{ // Segment group 20
				Group: []SchemaNode{
					{Segment: "PCC", Mandatory: true, MaxRepeat: 1},  // Premium calculation component details
					{Segment: "IDE", Mandatory: false, MaxRepeat: 9}, // Identity
					{Segment: "ADR", Mandatory: false, MaxRepeat: 1}, // Address
					{Segment: "ATT", Mandatory: false, MaxRepeat: 9}, // Attribute
					{Segment: "DTM", Mandatory: false, MaxRepeat: 9}, // Date/time/period
					{Segment: "MOA", Mandatory: true, MaxRepeat: 9},  // Monetary amount
					{Segment: "PCD", Mandatory: false, MaxRepeat: 9}, // Percentage details
					{Segment: "QRS", Mandatory: false, MaxRepeat: 9}, // Query and response
					{Segment: "QTY", Mandatory: false, MaxRepeat: 9}, // Quantity
					{Segment: "FTX", Mandatory: false, MaxRepeat: 9}, // Free text
				},
				Mandatory: false, MaxRepeat: 99,
			},
		},
		Mandatory: false, MaxRepeat: 9,
	},
}}
