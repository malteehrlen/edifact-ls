package edifact

// QALITY D.20A message specification: the structural (segment/group
// presence, order, cardinality) branching diagram for this message,
// UN/EDIFACT directory release D.20A. Scope is structural only -- not
// element-level content -- per edifact-ls-3uzr's epic non-goal.
// 39 segment groups, max nesting depth 5.
//
// Source: https://service.unece.org/trade/untdid/d20a/trmd/qality_c.htm
// section 4.3.1 "Segment table". That URL currently returns HTTP 403 from
// Cloudflare when fetched directly; this data was transcribed from the
// Wayback Machine's archived copy instead:
// http://web.archive.org/web/20231002212316/https://service.unece.org/trade/untdid/d20a/trmd/qality_c.htm
// -- re-check the direct URL first if cross-checking later, in case the
// block has lifted.
//
// Transcription note: generated in bulk (edifact-ls-13gu) the same way
// as every other message-type schema -- a script parsed the source's
// exact rail-art column positions mechanically and verified the result
// balances (all 39 groups open and close consistently)
// before being transcribed here, rather than reading the ASCII nesting
// by eye. Not individually spot-checked against the raw source the way
// the first several message types in this project were -- see
// edifact-ls-13gu for which representative samples were spot-checked
// for this batch.

func init() {
	RegisterSchema(
		MessageID{Type: "QALITY", Version: "D", Release: "20A", Agency: "UN"},
		qalityD20aSchema,
		"https://service.unece.org/trade/untdid/d20a/trmd/qality_c.htm",
	)
}

var qalityD20aSchema = Schema{Nodes: []SchemaNode{
	{Segment: "BGM", Mandatory: true, MaxRepeat: 1},   // Beginning of message
	{Segment: "DTM", Mandatory: true, MaxRepeat: 10},  // Date/time/period
	{Segment: "IMD", Mandatory: false, MaxRepeat: 10}, // Item description
	{Segment: "MEA", Mandatory: false, MaxRepeat: 10}, // Measurements
	{Segment: "FTX", Mandatory: false, MaxRepeat: 5},  // Free text
	{ // Segment group 1
		Group: []SchemaNode{
			{Segment: "RFF", Mandatory: true, MaxRepeat: 1},  // Reference
			{Segment: "DTM", Mandatory: false, MaxRepeat: 2}, // Date/time/period
		},
		Mandatory: false, MaxRepeat: 10,
	},
	{ // Segment group 2
		Group: []SchemaNode{
			{Segment: "NAD", Mandatory: true, MaxRepeat: 1},  // Name and address
			{Segment: "LOC", Mandatory: false, MaxRepeat: 5}, // Place/location identification
			{ // Segment group 3
				Group: []SchemaNode{
					{Segment: "RFF", Mandatory: true, MaxRepeat: 1},  // Reference
					{Segment: "DTM", Mandatory: false, MaxRepeat: 2}, // Date/time/period
				},
				Mandatory: false, MaxRepeat: 10,
			},
			{ // Segment group 4
				Group: []SchemaNode{
					{Segment: "CTA", Mandatory: true, MaxRepeat: 1},  // Contact information
					{Segment: "COM", Mandatory: false, MaxRepeat: 5}, // Communication contact
				},
				Mandatory: false, MaxRepeat: 5,
			},
		},
		Mandatory: false, MaxRepeat: 10,
	},
	{ // Segment group 5
		Group: []SchemaNode{
			{Segment: "LIN", Mandatory: true, MaxRepeat: 1},   // Line item
			{Segment: "PIA", Mandatory: false, MaxRepeat: 10}, // Additional product id
			{Segment: "IMD", Mandatory: false, MaxRepeat: 10}, // Item description
			{Segment: "MEA", Mandatory: false, MaxRepeat: 10}, // Measurements
			{Segment: "PSD", Mandatory: false, MaxRepeat: 1},  // Physical sample description
			{Segment: "SPS", Mandatory: false, MaxRepeat: 1},  // Sampling parameters for summary statistics
			{Segment: "DTM", Mandatory: false, MaxRepeat: 10}, // Date/time/period
			{Segment: "QTY", Mandatory: false, MaxRepeat: 99}, // Quantity
			{Segment: "FTX", Mandatory: false, MaxRepeat: 5},  // Free text
			{ // Segment group 6
				Group: []SchemaNode{
					{Segment: "RFF", Mandatory: true, MaxRepeat: 1},  // Reference
					{Segment: "DTM", Mandatory: false, MaxRepeat: 2}, // Date/time/period
				},
				Mandatory: false, MaxRepeat: 10,
			},
			{ // Segment group 7
				Group: []SchemaNode{
					{Segment: "NAD", Mandatory: true, MaxRepeat: 1},  // Name and address
					{Segment: "LOC", Mandatory: false, MaxRepeat: 5}, // Place/location identification
					{ // Segment group 8
						Group: []SchemaNode{
							{Segment: "RFF", Mandatory: true, MaxRepeat: 1},  // Reference
							{Segment: "DTM", Mandatory: false, MaxRepeat: 2}, // Date/time/period
						},
						Mandatory: false, MaxRepeat: 10,
					},
					{ // Segment group 9
						Group: []SchemaNode{
							{Segment: "CTA", Mandatory: true, MaxRepeat: 1},  // Contact information
							{Segment: "COM", Mandatory: false, MaxRepeat: 5}, // Communication contact
						},
						Mandatory: false, MaxRepeat: 5,
					},
				},
				Mandatory: false, MaxRepeat: 10,
			},
			{ // Segment group 10
				Group: []SchemaNode{
					{Segment: "TEM", Mandatory: true, MaxRepeat: 1},    // Test method
					{Segment: "MEA", Mandatory: false, MaxRepeat: 100}, // Measurements
					{Segment: "DTM", Mandatory: false, MaxRepeat: 10},  // Date/time/period
					{ // Segment group 11
						Group: []SchemaNode{
							{Segment: "RFF", Mandatory: true, MaxRepeat: 1},  // Reference
							{Segment: "DTM", Mandatory: false, MaxRepeat: 2}, // Date/time/period
						},
						Mandatory: false, MaxRepeat: 10,
					},
				},
				Mandatory: false, MaxRepeat: 100,
			},
			{ // Segment group 12
				Group: []SchemaNode{
					{Segment: "CCI", Mandatory: true, MaxRepeat: 1},   // Characteristic/class id
					{Segment: "PSD", Mandatory: false, MaxRepeat: 10}, // Physical sample description
					{Segment: "SPS", Mandatory: false, MaxRepeat: 10}, // Sampling parameters for summary statistics
					{Segment: "DTM", Mandatory: false, MaxRepeat: 10}, // Date/time/period
					{Segment: "FTX", Mandatory: false, MaxRepeat: 10}, // Free text
					{ // Segment group 13
						Group: []SchemaNode{
							{Segment: "RFF", Mandatory: true, MaxRepeat: 1},  // Reference
							{Segment: "DTM", Mandatory: false, MaxRepeat: 2}, // Date/time/period
						},
						Mandatory: false, MaxRepeat: 10,
					},
					{ // Segment group 14
						Group: []SchemaNode{
							{Segment: "MEA", Mandatory: true, MaxRepeat: 1},   // Measurements
							{Segment: "DTM", Mandatory: false, MaxRepeat: 10}, // Date/time/period
							{ // Segment group 15
								Group: []SchemaNode{
									{Segment: "RFF", Mandatory: true, MaxRepeat: 1},  // Reference
									{Segment: "DTM", Mandatory: false, MaxRepeat: 2}, // Date/time/period
								},
								Mandatory: false, MaxRepeat: 10,
							},
						},
						Mandatory: false, MaxRepeat: 999,
					},
					{ // Segment group 16
						Group: []SchemaNode{
							{Segment: "STA", Mandatory: true, MaxRepeat: 1},   // Statistics
							{Segment: "DTM", Mandatory: false, MaxRepeat: 10}, // Date/time/period
							{ // Segment group 17
								Group: []SchemaNode{
									{Segment: "RFF", Mandatory: true, MaxRepeat: 1},  // Reference
									{Segment: "DTM", Mandatory: false, MaxRepeat: 2}, // Date/time/period
								},
								Mandatory: false, MaxRepeat: 10,
							},
						},
						Mandatory: false, MaxRepeat: 100,
					},
					{ // Segment group 18
						Group: []SchemaNode{
							{Segment: "TEM", Mandatory: true, MaxRepeat: 1},    // Test method
							{Segment: "MEA", Mandatory: false, MaxRepeat: 100}, // Measurements
							{Segment: "DTM", Mandatory: false, MaxRepeat: 10},  // Date/time/period
							{ // Segment group 19
								Group: []SchemaNode{
									{Segment: "RFF", Mandatory: true, MaxRepeat: 1},  // Reference
									{Segment: "DTM", Mandatory: false, MaxRepeat: 2}, // Date/time/period
								},
								Mandatory: false, MaxRepeat: 10,
							},
						},
						Mandatory: false, MaxRepeat: 100,
					},
				},
				Mandatory: false, MaxRepeat: 200,
			},
			{ // Segment group 20
				Group: []SchemaNode{
					{Segment: "GIN", Mandatory: true, MaxRepeat: 1},   // Goods identity number
					{Segment: "DTM", Mandatory: false, MaxRepeat: 10}, // Date/time/period
					{ // Segment group 21
						Group: []SchemaNode{
							{Segment: "RFF", Mandatory: true, MaxRepeat: 1},  // Reference
							{Segment: "DTM", Mandatory: false, MaxRepeat: 2}, // Date/time/period
						},
						Mandatory: false, MaxRepeat: 10,
					},
					{ // Segment group 22
						Group: []SchemaNode{
							{Segment: "CCI", Mandatory: true, MaxRepeat: 1},   // Characteristic/class id
							{Segment: "PSD", Mandatory: false, MaxRepeat: 10}, // Physical sample description
							{Segment: "SPS", Mandatory: false, MaxRepeat: 10}, // Sampling parameters for summary statistics
							{Segment: "DTM", Mandatory: false, MaxRepeat: 10}, // Date/time/period
							{Segment: "FTX", Mandatory: false, MaxRepeat: 10}, // Free text
							{ // Segment group 23
								Group: []SchemaNode{
									{Segment: "RFF", Mandatory: true, MaxRepeat: 1},  // Reference
									{Segment: "DTM", Mandatory: false, MaxRepeat: 2}, // Date/time/period
								},
								Mandatory: false, MaxRepeat: 10,
							},
							{ // Segment group 24
								Group: []SchemaNode{
									{Segment: "MEA", Mandatory: true, MaxRepeat: 1},   // Measurements
									{Segment: "DTM", Mandatory: false, MaxRepeat: 10}, // Date/time/period
									{ // Segment group 25
										Group: []SchemaNode{
											{Segment: "RFF", Mandatory: true, MaxRepeat: 1},  // Reference
											{Segment: "DTM", Mandatory: false, MaxRepeat: 2}, // Date/time/period
										},
										Mandatory: false, MaxRepeat: 10,
									},
								},
								Mandatory: false, MaxRepeat: 999,
							},
							{ // Segment group 26
								Group: []SchemaNode{
									{Segment: "STA", Mandatory: true, MaxRepeat: 1},   // Statistics
									{Segment: "DTM", Mandatory: false, MaxRepeat: 10}, // Date/time/period
									{ // Segment group 27
										Group: []SchemaNode{
											{Segment: "RFF", Mandatory: true, MaxRepeat: 1},  // Reference
											{Segment: "DTM", Mandatory: false, MaxRepeat: 2}, // Date/time/period
										},
										Mandatory: false, MaxRepeat: 10,
									},
								},
								Mandatory: false, MaxRepeat: 100,
							},
							{ // Segment group 28
								Group: []SchemaNode{
									{Segment: "TEM", Mandatory: true, MaxRepeat: 1},    // Test method
									{Segment: "MEA", Mandatory: false, MaxRepeat: 100}, // Measurements
									{Segment: "DTM", Mandatory: false, MaxRepeat: 10},  // Date/time/period
									{ // Segment group 29
										Group: []SchemaNode{
											{Segment: "RFF", Mandatory: true, MaxRepeat: 1},  // Reference
											{Segment: "DTM", Mandatory: false, MaxRepeat: 2}, // Date/time/period
										},
										Mandatory: false, MaxRepeat: 10,
									},
								},
								Mandatory: false, MaxRepeat: 100,
							},
						},
						Mandatory: false, MaxRepeat: 200,
					},
				},
				Mandatory: false, MaxRepeat: 100,
			},
			{ // Segment group 30
				Group: []SchemaNode{
					{Segment: "PRC", Mandatory: true, MaxRepeat: 1}, // Process identification
					{ // Segment group 31
						Group: []SchemaNode{
							{Segment: "NAD", Mandatory: true, MaxRepeat: 1},  // Name and address
							{Segment: "LOC", Mandatory: false, MaxRepeat: 5}, // Place/location identification
						},
						Mandatory: false, MaxRepeat: 10,
					},
					{ // Segment group 32
						Group: []SchemaNode{
							{Segment: "CCI", Mandatory: true, MaxRepeat: 1},   // Characteristic/class id
							{Segment: "PSD", Mandatory: false, MaxRepeat: 10}, // Physical sample description
							{Segment: "SPS", Mandatory: false, MaxRepeat: 10}, // Sampling parameters for summary statistics
							{Segment: "DTM", Mandatory: false, MaxRepeat: 10}, // Date/time/period
							{Segment: "FTX", Mandatory: false, MaxRepeat: 10}, // Free text
							{ // Segment group 33
								Group: []SchemaNode{
									{Segment: "RFF", Mandatory: true, MaxRepeat: 1},  // Reference
									{Segment: "DTM", Mandatory: false, MaxRepeat: 2}, // Date/time/period
								},
								Mandatory: false, MaxRepeat: 10,
							},
							{ // Segment group 34
								Group: []SchemaNode{
									{Segment: "MEA", Mandatory: true, MaxRepeat: 1},   // Measurements
									{Segment: "DTM", Mandatory: false, MaxRepeat: 10}, // Date/time/period
									{ // Segment group 35
										Group: []SchemaNode{
											{Segment: "RFF", Mandatory: true, MaxRepeat: 1},  // Reference
											{Segment: "DTM", Mandatory: false, MaxRepeat: 2}, // Date/time/period
										},
										Mandatory: false, MaxRepeat: 10,
									},
								},
								Mandatory: false, MaxRepeat: 999,
							},
							{ // Segment group 36
								Group: []SchemaNode{
									{Segment: "STA", Mandatory: true, MaxRepeat: 1},   // Statistics
									{Segment: "DTM", Mandatory: false, MaxRepeat: 10}, // Date/time/period
									{ // Segment group 37
										Group: []SchemaNode{
											{Segment: "RFF", Mandatory: true, MaxRepeat: 1},  // Reference
											{Segment: "DTM", Mandatory: false, MaxRepeat: 2}, // Date/time/period
										},
										Mandatory: false, MaxRepeat: 10,
									},
								},
								Mandatory: false, MaxRepeat: 100,
							},
							{ // Segment group 38
								Group: []SchemaNode{
									{Segment: "TEM", Mandatory: true, MaxRepeat: 1},    // Test method
									{Segment: "MEA", Mandatory: false, MaxRepeat: 100}, // Measurements
									{Segment: "DTM", Mandatory: false, MaxRepeat: 10},  // Date/time/period
									{ // Segment group 39
										Group: []SchemaNode{
											{Segment: "RFF", Mandatory: true, MaxRepeat: 1},  // Reference
											{Segment: "DTM", Mandatory: false, MaxRepeat: 2}, // Date/time/period
										},
										Mandatory: false, MaxRepeat: 10,
									},
								},
								Mandatory: false, MaxRepeat: 100,
							},
						},
						Mandatory: false, MaxRepeat: 200,
					},
				},
				Mandatory: false, MaxRepeat: 100,
			},
		},
		Mandatory: false, MaxRepeat: 200,
	},
}}
