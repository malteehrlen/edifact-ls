package edifact

// STLRPT D.20A message specification: the structural (segment/group
// presence, order, cardinality) branching diagram for this message,
// UN/EDIFACT directory release D.20A. Scope is structural only -- not
// element-level content -- per edifact-ls-3uzr's epic non-goal.
// 12 segment groups, max nesting depth 6.
//
// Source: https://service.unece.org/trade/untdid/d20a/trmd/stlrpt_c.htm
// section 4.3.1 "Segment table". That URL currently returns HTTP 403 from
// Cloudflare when fetched directly; this data was transcribed from the
// Wayback Machine's archived copy instead:
// http://web.archive.org/web/20231002223615/https://service.unece.org/trade/untdid/d20a/trmd/stlrpt_c.htm
// -- re-check the direct URL first if cross-checking later, in case the
// block has lifted.
//
// Transcription note: generated in bulk (edifact-ls-13gu) the same way
// as every other message-type schema -- a script parsed the source's
// exact rail-art column positions mechanically and verified the result
// balances (all 12 groups open and close consistently)
// before being transcribed here, rather than reading the ASCII nesting
// by eye. Not individually spot-checked against the raw source the way
// the first several message types in this project were -- see
// edifact-ls-13gu for which representative samples were spot-checked
// for this batch.

func init() {
	RegisterSchema(
		MessageID{Type: "STLRPT", Version: "D", Release: "20A", Agency: "UN"},
		stlrptD20aSchema,
		"https://service.unece.org/trade/untdid/d20a/trmd/stlrpt_c.htm",
	)
}

var stlrptD20aSchema = Schema{Nodes: []SchemaNode{
	{Segment: "BGM", Mandatory: true, MaxRepeat: 1}, // Beginning of message
	{Segment: "DTM", Mandatory: true, MaxRepeat: 2}, // Date/time/period
	{Segment: "NAD", Mandatory: true, MaxRepeat: 1}, // Name and address
	{ // Segment group 1
		Group: []SchemaNode{
			{Segment: "CUX", Mandatory: true, MaxRepeat: 1}, // Currencies
			{Segment: "EQN", Mandatory: true, MaxRepeat: 1}, // Number of units
			{Segment: "MOA", Mandatory: true, MaxRepeat: 6}, // Monetary amount
		},
		Mandatory: false, MaxRepeat: 999,
	},
	{ // Segment group 2
		Group: []SchemaNode{
			{Segment: "DTM", Mandatory: true, MaxRepeat: 1},  // Date/time/period
			{Segment: "GEI", Mandatory: false, MaxRepeat: 1}, // Processing information
			{ // Segment group 3
				Group: []SchemaNode{
					{Segment: "CUX", Mandatory: true, MaxRepeat: 1},  // Currencies
					{Segment: "EQN", Mandatory: true, MaxRepeat: 1},  // Number of units
					{Segment: "MOA", Mandatory: true, MaxRepeat: 6},  // Monetary amount
					{Segment: "QVR", Mandatory: false, MaxRepeat: 1}, // Quantity variances
				},
				Mandatory: false, MaxRepeat: 999,
			},
			{ // Segment group 4
				Group: []SchemaNode{
					{Segment: "RFF", Mandatory: true, MaxRepeat: 1},  // Reference
					{Segment: "LOC", Mandatory: false, MaxRepeat: 1}, // Place/location identification
					{Segment: "DTM", Mandatory: false, MaxRepeat: 1}, // Date/time/period
					{Segment: "FTX", Mandatory: false, MaxRepeat: 9}, // Free text
					{ // Segment group 5
						Group: []SchemaNode{
							{Segment: "CUX", Mandatory: true, MaxRepeat: 1},  // Currencies
							{Segment: "MOA", Mandatory: true, MaxRepeat: 9},  // Monetary amount
							{Segment: "QVR", Mandatory: false, MaxRepeat: 1}, // Quantity variances
						},
						Mandatory: false, MaxRepeat: 999,
					},
					{ // Segment group 6
						Group: []SchemaNode{
							{Segment: "BUS", Mandatory: true, MaxRepeat: 1}, // Business function
							{ // Segment group 7
								Group: []SchemaNode{
									{Segment: "CUX", Mandatory: true, MaxRepeat: 1},  // Currencies
									{Segment: "MOA", Mandatory: true, MaxRepeat: 9},  // Monetary amount
									{Segment: "QVR", Mandatory: false, MaxRepeat: 1}, // Quantity variances
								},
								Mandatory: false, MaxRepeat: 999,
							},
							{ // Segment group 8
								Group: []SchemaNode{
									{Segment: "LIN", Mandatory: true, MaxRepeat: 1},    // Line item
									{Segment: "DTM", Mandatory: false, MaxRepeat: 1},   // Date/time/period
									{Segment: "RFF", Mandatory: false, MaxRepeat: 9},   // Reference
									{Segment: "NAD", Mandatory: false, MaxRepeat: 9},   // Name and address
									{Segment: "LOC", Mandatory: false, MaxRepeat: 1},   // Place/location identification
									{Segment: "CUX", Mandatory: false, MaxRepeat: 1},   // Currencies
									{Segment: "GEI", Mandatory: false, MaxRepeat: 2},   // Processing information
									{Segment: "IMD", Mandatory: false, MaxRepeat: 9},   // Item description
									{Segment: "MOA", Mandatory: false, MaxRepeat: 9},   // Monetary amount
									{Segment: "ALC", Mandatory: false, MaxRepeat: 9},   // Allowance or charge
									{Segment: "FTX", Mandatory: false, MaxRepeat: 99},  // Free text
									{Segment: "TAX", Mandatory: false, MaxRepeat: 999}, // Duty/tax/fee details
									{ // Segment group 9
										Group: []SchemaNode{
											{Segment: "DOC", Mandatory: true, MaxRepeat: 1},  // Document/message details
											{Segment: "DTM", Mandatory: true, MaxRepeat: 1},  // Date/time/period
											{Segment: "FTX", Mandatory: false, MaxRepeat: 1}, // Free text
											{Segment: "PIA", Mandatory: false, MaxRepeat: 1}, // Additional product id
											{Segment: "SEQ", Mandatory: false, MaxRepeat: 1}, // Sequence details
										},
										Mandatory: false, MaxRepeat: 99,
									},
									{ // Segment group 10
										Group: []SchemaNode{
											{Segment: "PYT", Mandatory: true, MaxRepeat: 1},  // Payment terms
											{Segment: "MOA", Mandatory: false, MaxRepeat: 9}, // Monetary amount
											{Segment: "RFF", Mandatory: false, MaxRepeat: 9}, // Reference
											{Segment: "DTM", Mandatory: false, MaxRepeat: 2}, // Date/time/period
											{Segment: "AGR", Mandatory: false, MaxRepeat: 1}, // Agreement identification
											{Segment: "GEI", Mandatory: false, MaxRepeat: 2}, // Processing information
											{Segment: "NAD", Mandatory: false, MaxRepeat: 1}, // Name and address
											{Segment: "FTX", Mandatory: false, MaxRepeat: 9}, // Free text
											{Segment: "IMD", Mandatory: false, MaxRepeat: 2}, // Item description
										},
										Mandatory: false, MaxRepeat: 9999,
									},
									{ // Segment group 11
										Group: []SchemaNode{
											{Segment: "GIR", Mandatory: true, MaxRepeat: 1},  // Related identification numbers
											{Segment: "IMD", Mandatory: false, MaxRepeat: 2}, // Item description
											{Segment: "GIN", Mandatory: false, MaxRepeat: 2}, // Goods identity number
											{ // Segment group 12
												Group: []SchemaNode{
													{Segment: "SEQ", Mandatory: true, MaxRepeat: 1},   // Sequence details
													{Segment: "FTX", Mandatory: false, MaxRepeat: 99}, // Free text
													{Segment: "ALC", Mandatory: false, MaxRepeat: 1},  // Allowance or charge
													{Segment: "GIN", Mandatory: false, MaxRepeat: 1},  // Goods identity number
													{Segment: "RFF", Mandatory: false, MaxRepeat: 2},  // Reference
													{Segment: "IMD", Mandatory: false, MaxRepeat: 9},  // Item description
													{Segment: "NAD", Mandatory: false, MaxRepeat: 1},  // Name and address
													{Segment: "TDT", Mandatory: false, MaxRepeat: 1},  // Transport information
													{Segment: "DTM", Mandatory: false, MaxRepeat: 9},  // Date/time/period
													{Segment: "LOC", Mandatory: false, MaxRepeat: 9},  // Place/location identification
												},
												Mandatory: false, MaxRepeat: 99,
											},
										},
										Mandatory: false, MaxRepeat: 999,
									},
								},
								Mandatory: true, MaxRepeat: 99999,
							},
						},
						Mandatory: true, MaxRepeat: 99,
					},
					{Segment: "CNT", Mandatory: true, MaxRepeat: 1}, // Control total
				},
				Mandatory: true, MaxRepeat: 999999,
			},
		},
		Mandatory: true, MaxRepeat: 9,
	},
}}
