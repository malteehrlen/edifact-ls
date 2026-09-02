package edifact

// PRIHIS D.20A message specification: the structural (segment/group
// presence, order, cardinality) branching diagram for this message,
// UN/EDIFACT directory release D.20A. Scope is structural only -- not
// element-level content -- per edifact-ls-3uzr's epic non-goal.
// 17 segment groups, max nesting depth 6.
//
// Source: https://service.unece.org/trade/untdid/d20a/trmd/prihis_c.htm
// section 4.3.1 "Segment table". That URL currently returns HTTP 403 from
// Cloudflare when fetched directly; this data was transcribed from the
// Wayback Machine's archived copy instead:
// http://web.archive.org/web/20231207010652/https://service.unece.org/trade/untdid/d20a/trmd/prihis_c.htm
// -- re-check the direct URL first if cross-checking later, in case the
// block has lifted.
//
// Transcription note: generated in bulk (edifact-ls-13gu) the same way
// as every other message-type schema -- a script parsed the source's
// exact rail-art column positions mechanically and verified the result
// balances (all 17 groups open and close consistently)
// before being transcribed here, rather than reading the ASCII nesting
// by eye. Not individually spot-checked against the raw source the way
// the first several message types in this project were -- see
// edifact-ls-13gu for which representative samples were spot-checked
// for this batch.

func init() {
	RegisterSchema(
		MessageID{Type: "PRIHIS", Version: "D", Release: "20A", Agency: "UN"},
		prihisD20aSchema,
		"https://service.unece.org/trade/untdid/d20a/trmd/prihis_c.htm",
	)
}

var prihisD20aSchema = Schema{Nodes: []SchemaNode{
	{Segment: "BGM", Mandatory: true, MaxRepeat: 1},  // Beginning of message
	{Segment: "DTM", Mandatory: true, MaxRepeat: 1},  // Date/time/period
	{Segment: "RFF", Mandatory: false, MaxRepeat: 9}, // Reference
	{Segment: "CUX", Mandatory: false, MaxRepeat: 9}, // Currencies
	{ // Segment group 1
		Group: []SchemaNode{
			{Segment: "NAD", Mandatory: true, MaxRepeat: 1},  // Name and address
			{Segment: "LOC", Mandatory: false, MaxRepeat: 9}, // Place/location identification
			{ // Segment group 2
				Group: []SchemaNode{
					{Segment: "RFF", Mandatory: true, MaxRepeat: 1},  // Reference
					{Segment: "DTM", Mandatory: false, MaxRepeat: 9}, // Date/time/period
				},
				Mandatory: false, MaxRepeat: 9,
			},
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
			{Segment: "LIN", Mandatory: true, MaxRepeat: 1},  // Line item
			{Segment: "PIA", Mandatory: false, MaxRepeat: 9}, // Additional product id
			{Segment: "RFF", Mandatory: false, MaxRepeat: 9}, // Reference
			{Segment: "IMD", Mandatory: false, MaxRepeat: 9}, // Item description
			{Segment: "MEA", Mandatory: false, MaxRepeat: 9}, // Measurements
			{Segment: "PCD", Mandatory: false, MaxRepeat: 9}, // Percentage details
			{Segment: "ALI", Mandatory: false, MaxRepeat: 9}, // Additional information
			{Segment: "GIR", Mandatory: false, MaxRepeat: 9}, // Related identification numbers
			{ // Segment group 5
				Group: []SchemaNode{
					{Segment: "NAD", Mandatory: true, MaxRepeat: 1},  // Name and address
					{Segment: "LOC", Mandatory: false, MaxRepeat: 9}, // Place/location identification
					{Segment: "STS", Mandatory: false, MaxRepeat: 9}, // Status
					{Segment: "PIA", Mandatory: false, MaxRepeat: 9}, // Additional product id
					{ // Segment group 6
						Group: []SchemaNode{
							{Segment: "CTA", Mandatory: true, MaxRepeat: 1},  // Contact information
							{Segment: "COM", Mandatory: false, MaxRepeat: 9}, // Communication contact
						},
						Mandatory: false, MaxRepeat: 9,
					},
					{ // Segment group 7
						Group: []SchemaNode{
							{Segment: "RFF", Mandatory: true, MaxRepeat: 1},  // Reference
							{Segment: "DTM", Mandatory: false, MaxRepeat: 9}, // Date/time/period
							{Segment: "AGR", Mandatory: false, MaxRepeat: 9}, // Agreement identification
							{ // Segment group 8
								Group: []SchemaNode{
									{Segment: "QTY", Mandatory: true, MaxRepeat: 1},  // Quantity
									{Segment: "DTM", Mandatory: false, MaxRepeat: 9}, // Date/time/period
									{Segment: "TAX", Mandatory: false, MaxRepeat: 9}, // Duty/tax/fee details
									{Segment: "FTX", Mandatory: false, MaxRepeat: 9}, // Free text
									{ // Segment group 9
										Group: []SchemaNode{
											{Segment: "PYT", Mandatory: true, MaxRepeat: 1},  // Payment terms
											{Segment: "DTM", Mandatory: false, MaxRepeat: 9}, // Date/time/period
											{Segment: "PCD", Mandatory: false, MaxRepeat: 9}, // Percentage details
											{Segment: "MOA", Mandatory: false, MaxRepeat: 9}, // Monetary amount
										},
										Mandatory: false, MaxRepeat: 9,
									},
									{ // Segment group 10
										Group: []SchemaNode{
											{Segment: "TOD", Mandatory: true, MaxRepeat: 1},  // Terms of delivery or transport
											{Segment: "LOC", Mandatory: false, MaxRepeat: 9}, // Place/location identification
										},
										Mandatory: false, MaxRepeat: 9,
									},
									{ // Segment group 11
										Group: []SchemaNode{
											{Segment: "ALC", Mandatory: true, MaxRepeat: 1},  // Allowance or charge
											{Segment: "ALI", Mandatory: false, MaxRepeat: 9}, // Additional information
											{Segment: "DTM", Mandatory: false, MaxRepeat: 9}, // Date/time/period
											{ // Segment group 12
												Group: []SchemaNode{
													{Segment: "QTY", Mandatory: true, MaxRepeat: 1},  // Quantity
													{Segment: "RNG", Mandatory: false, MaxRepeat: 1}, // Range details
												},
												Mandatory: false, MaxRepeat: 9,
											},
											{ // Segment group 13
												Group: []SchemaNode{
													{Segment: "PCD", Mandatory: true, MaxRepeat: 1},  // Percentage details
													{Segment: "RNG", Mandatory: false, MaxRepeat: 1}, // Range details
												},
												Mandatory: false, MaxRepeat: 9,
											},
											{ // Segment group 14
												Group: []SchemaNode{
													{Segment: "MOA", Mandatory: true, MaxRepeat: 1},  // Monetary amount
													{Segment: "RNG", Mandatory: false, MaxRepeat: 1}, // Range details
												},
												Mandatory: false, MaxRepeat: 9,
											},
											{ // Segment group 15
												Group: []SchemaNode{
													{Segment: "RTE", Mandatory: true, MaxRepeat: 1},  // Rate details
													{Segment: "RNG", Mandatory: false, MaxRepeat: 1}, // Range details
												},
												Mandatory: false, MaxRepeat: 9,
											},
											{ // Segment group 16
												Group: []SchemaNode{
													{Segment: "TAX", Mandatory: true, MaxRepeat: 1},  // Duty/tax/fee details
													{Segment: "MOA", Mandatory: false, MaxRepeat: 1}, // Monetary amount
												},
												Mandatory: false, MaxRepeat: 9,
											},
										},
										Mandatory: false, MaxRepeat: 99,
									},
									{ // Segment group 17
										Group: []SchemaNode{
											{Segment: "PRI", Mandatory: true, MaxRepeat: 1},  // Price details
											{Segment: "DTM", Mandatory: false, MaxRepeat: 9}, // Date/time/period
											{Segment: "CUX", Mandatory: false, MaxRepeat: 9}, // Currencies
											{Segment: "RNG", Mandatory: false, MaxRepeat: 9}, // Range details
										},
										Mandatory: true, MaxRepeat: 9,
									},
								},
								Mandatory: false, MaxRepeat: 99,
							},
						},
						Mandatory: false, MaxRepeat: 9,
					},
				},
				Mandatory: false, MaxRepeat: 999,
			},
		},
		Mandatory: false, MaxRepeat: 99999,
	},
}}
