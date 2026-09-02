package edifact

// CUSPED D.20A message specification: the structural (segment/group
// presence, order, cardinality) branching diagram for this message,
// UN/EDIFACT directory release D.20A. Scope is structural only -- not
// element-level content -- per edifact-ls-3uzr's epic non-goal.
// 34 segment groups, max nesting depth 5.
//
// Source: https://service.unece.org/trade/untdid/d20a/trmd/cusped_c.htm
// section 4.3.1 "Segment table". That URL currently returns HTTP 403 from
// Cloudflare when fetched directly; this data was transcribed from the
// Wayback Machine's archived copy instead:
// http://web.archive.org/web/20231002214644/https://service.unece.org/trade/untdid/d20a/trmd/cusped_c.htm
// -- re-check the direct URL first if cross-checking later, in case the
// block has lifted.
//
// Transcription note: generated in bulk (edifact-ls-13gu) the same way
// as every other message-type schema -- a script parsed the source's
// exact rail-art column positions mechanically and verified the result
// balances (all 34 groups open and close consistently)
// before being transcribed here, rather than reading the ASCII nesting
// by eye. Not individually spot-checked against the raw source the way
// the first several message types in this project were -- see
// edifact-ls-13gu for which representative samples were spot-checked
// for this batch.

func init() {
	RegisterSchema(
		MessageID{Type: "CUSPED", Version: "D", Release: "20A", Agency: "UN"},
		cuspedD20aSchema,
		"https://service.unece.org/trade/untdid/d20a/trmd/cusped_c.htm",
	)
}

var cuspedD20aSchema = Schema{Nodes: []SchemaNode{
	{Segment: "BGM", Mandatory: true, MaxRepeat: 1},   // Beginning of message
	{Segment: "CST", Mandatory: false, MaxRepeat: 1},  // Customs status of goods
	{Segment: "LOC", Mandatory: false, MaxRepeat: 99}, // Place/location identification
	{Segment: "DTM", Mandatory: false, MaxRepeat: 99}, // Date/time/period
	{Segment: "GEI", Mandatory: false, MaxRepeat: 99}, // Processing information
	{Segment: "FII", Mandatory: false, MaxRepeat: 1},  // Financial institution information
	{ // Segment group 1
		Group: []SchemaNode{
			{Segment: "RFF", Mandatory: true, MaxRepeat: 1},  // Reference
			{Segment: "DTM", Mandatory: false, MaxRepeat: 9}, // Date/time/period
		},
		Mandatory: false, MaxRepeat: 99,
	},
	{ // Segment group 2
		Group: []SchemaNode{
			{Segment: "NAD", Mandatory: true, MaxRepeat: 1},  // Name and address
			{Segment: "RFF", Mandatory: false, MaxRepeat: 9}, // Reference
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
			{Segment: "MOA", Mandatory: true, MaxRepeat: 1}, // Monetary amount
			{ // Segment group 5
				Group: []SchemaNode{
					{Segment: "CUX", Mandatory: true, MaxRepeat: 1},  // Currencies
					{Segment: "DTM", Mandatory: false, MaxRepeat: 1}, // Date/time/period
				},
				Mandatory: false, MaxRepeat: 9,
			},
		},
		Mandatory: false, MaxRepeat: 99,
	},
	{ // Segment group 6
		Group: []SchemaNode{
			{Segment: "TAX", Mandatory: true, MaxRepeat: 1},  // Duty/tax/fee details
			{Segment: "MOA", Mandatory: false, MaxRepeat: 9}, // Monetary amount
			{Segment: "GEI", Mandatory: false, MaxRepeat: 9}, // Processing information
		},
		Mandatory: false, MaxRepeat: 99,
	},
	{Segment: "CNT", Mandatory: false, MaxRepeat: 9}, // Control total
	{ // Segment group 7
		Group: []SchemaNode{
			{Segment: "DMS", Mandatory: true, MaxRepeat: 1},    // Document/message summary
			{Segment: "CNT", Mandatory: false, MaxRepeat: 9},   // Control total
			{Segment: "RFF", Mandatory: false, MaxRepeat: 9},   // Reference
			{Segment: "DTM", Mandatory: false, MaxRepeat: 9},   // Date/time/period
			{Segment: "GEI", Mandatory: false, MaxRepeat: 99},  // Processing information
			{Segment: "LOC", Mandatory: false, MaxRepeat: 99},  // Place/location identification
			{Segment: "EQD", Mandatory: false, MaxRepeat: 999}, // Equipment details
			{Segment: "PAC", Mandatory: false, MaxRepeat: 9},   // Package
			{Segment: "TDT", Mandatory: false, MaxRepeat: 9},   // Transport information
			{ // Segment group 8
				Group: []SchemaNode{
					{Segment: "NAD", Mandatory: true, MaxRepeat: 1},  // Name and address
					{Segment: "RFF", Mandatory: false, MaxRepeat: 9}, // Reference
					{ // Segment group 9
						Group: []SchemaNode{
							{Segment: "CTA", Mandatory: true, MaxRepeat: 1},  // Contact information
							{Segment: "COM", Mandatory: false, MaxRepeat: 9}, // Communication contact
						},
						Mandatory: false, MaxRepeat: 9,
					},
				},
				Mandatory: false, MaxRepeat: 9,
			},
			{ // Segment group 10
				Group: []SchemaNode{
					{Segment: "DOC", Mandatory: true, MaxRepeat: 1},  // Document/message details
					{Segment: "DTM", Mandatory: false, MaxRepeat: 9}, // Date/time/period
					{Segment: "LOC", Mandatory: false, MaxRepeat: 1}, // Place/location identification
					{Segment: "NAD", Mandatory: false, MaxRepeat: 1}, // Name and address
				},
				Mandatory: false, MaxRepeat: 99,
			},
			{ // Segment group 11
				Group: []SchemaNode{
					{Segment: "TOD", Mandatory: true, MaxRepeat: 1},  // Terms of delivery or transport
					{Segment: "LOC", Mandatory: false, MaxRepeat: 9}, // Place/location identification
					{Segment: "FTX", Mandatory: false, MaxRepeat: 9}, // Free text
				},
				Mandatory: false, MaxRepeat: 9,
			},
			{ // Segment group 12
				Group: []SchemaNode{
					{Segment: "MOA", Mandatory: true, MaxRepeat: 1}, // Monetary amount
					{ // Segment group 13
						Group: []SchemaNode{
							{Segment: "CUX", Mandatory: true, MaxRepeat: 1},  // Currencies
							{Segment: "DTM", Mandatory: false, MaxRepeat: 1}, // Date/time/period
						},
						Mandatory: false, MaxRepeat: 9,
					},
				},
				Mandatory: false, MaxRepeat: 99,
			},
			{ // Segment group 14
				Group: []SchemaNode{
					{Segment: "TAX", Mandatory: true, MaxRepeat: 1},  // Duty/tax/fee details
					{Segment: "MOA", Mandatory: false, MaxRepeat: 9}, // Monetary amount
					{Segment: "GEI", Mandatory: false, MaxRepeat: 9}, // Processing information
				},
				Mandatory: false, MaxRepeat: 99,
			},
			{ // Segment group 15
				Group: []SchemaNode{
					{Segment: "CST", Mandatory: true, MaxRepeat: 1},   // Customs status of goods
					{Segment: "FTX", Mandatory: false, MaxRepeat: 9},  // Free text
					{Segment: "LOC", Mandatory: false, MaxRepeat: 99}, // Place/location identification
					{Segment: "DTM", Mandatory: false, MaxRepeat: 9},  // Date/time/period
					{Segment: "MEA", Mandatory: false, MaxRepeat: 99}, // Measurements
					{Segment: "NAD", Mandatory: false, MaxRepeat: 9},  // Name and address
					{Segment: "TDT", Mandatory: false, MaxRepeat: 9},  // Transport information
					{ // Segment group 16
						Group: []SchemaNode{
							{Segment: "PAC", Mandatory: true, MaxRepeat: 1}, // Package
							{ // Segment group 17
								Group: []SchemaNode{
									{Segment: "PCI", Mandatory: true, MaxRepeat: 1},  // Package identification
									{Segment: "FTX", Mandatory: false, MaxRepeat: 1}, // Free text
								},
								Mandatory: false, MaxRepeat: 999,
							},
						},
						Mandatory: false, MaxRepeat: 99,
					},
					{ // Segment group 18
						Group: []SchemaNode{
							{Segment: "MOA", Mandatory: true, MaxRepeat: 1}, // Monetary amount
							{ // Segment group 19
								Group: []SchemaNode{
									{Segment: "CUX", Mandatory: true, MaxRepeat: 1},  // Currencies
									{Segment: "DTM", Mandatory: false, MaxRepeat: 1}, // Date/time/period
								},
								Mandatory: false, MaxRepeat: 1,
							},
						},
						Mandatory: false, MaxRepeat: 99,
					},
					{ // Segment group 20
						Group: []SchemaNode{
							{Segment: "DOC", Mandatory: true, MaxRepeat: 1},   // Document/message details
							{Segment: "DTM", Mandatory: false, MaxRepeat: 9},  // Date/time/period
							{Segment: "LOC", Mandatory: false, MaxRepeat: 1},  // Place/location identification
							{Segment: "NAD", Mandatory: false, MaxRepeat: 1},  // Name and address
							{Segment: "GIN", Mandatory: false, MaxRepeat: 99}, // Goods identity number
							{Segment: "MOA", Mandatory: false, MaxRepeat: 1},  // Monetary amount
							{ // Segment group 21
								Group: []SchemaNode{
									{Segment: "IMD", Mandatory: true, MaxRepeat: 1},  // Item description
									{Segment: "FTX", Mandatory: false, MaxRepeat: 9}, // Free text
								},
								Mandatory: false, MaxRepeat: 1,
							},
						},
						Mandatory: false, MaxRepeat: 999,
					},
					{ // Segment group 22
						Group: []SchemaNode{
							{Segment: "RFF", Mandatory: true, MaxRepeat: 1},  // Reference
							{Segment: "DTM", Mandatory: false, MaxRepeat: 1}, // Date/time/period
						},
						Mandatory: false, MaxRepeat: 99,
					},
					{ // Segment group 23
						Group: []SchemaNode{
							{Segment: "TOD", Mandatory: true, MaxRepeat: 1},  // Terms of delivery or transport
							{Segment: "LOC", Mandatory: false, MaxRepeat: 9}, // Place/location identification
							{Segment: "FTX", Mandatory: false, MaxRepeat: 1}, // Free text
						},
						Mandatory: false, MaxRepeat: 1,
					},
					{ // Segment group 24
						Group: []SchemaNode{
							{Segment: "GDS", Mandatory: true, MaxRepeat: 1},  // Nature of cargo
							{Segment: "FTX", Mandatory: false, MaxRepeat: 1}, // Free text
						},
						Mandatory: false, MaxRepeat: 1,
					},
					{ // Segment group 25
						Group: []SchemaNode{
							{Segment: "GEI", Mandatory: true, MaxRepeat: 1},  // Processing information
							{Segment: "PCD", Mandatory: false, MaxRepeat: 1}, // Percentage details
							{Segment: "DTM", Mandatory: false, MaxRepeat: 1}, // Date/time/period
							{Segment: "RFF", Mandatory: false, MaxRepeat: 1}, // Reference
						},
						Mandatory: false, MaxRepeat: 99,
					},
					{ // Segment group 26
						Group: []SchemaNode{
							{Segment: "TAX", Mandatory: true, MaxRepeat: 1},  // Duty/tax/fee details
							{Segment: "MOA", Mandatory: false, MaxRepeat: 9}, // Monetary amount
							{Segment: "GEI", Mandatory: false, MaxRepeat: 1}, // Processing information
						},
						Mandatory: false, MaxRepeat: 99,
					},
					{ // Segment group 27
						Group: []SchemaNode{
							{Segment: "QVR", Mandatory: true, MaxRepeat: 1},  // Quantity variances
							{Segment: "QTY", Mandatory: false, MaxRepeat: 1}, // Quantity
							{ // Segment group 28
								Group: []SchemaNode{
									{Segment: "RFF", Mandatory: true, MaxRepeat: 1},  // Reference
									{Segment: "DTM", Mandatory: false, MaxRepeat: 1}, // Date/time/period
								},
								Mandatory: false, MaxRepeat: 1,
							},
						},
						Mandatory: false, MaxRepeat: 1,
					},
					{ // Segment group 29
						Group: []SchemaNode{
							{Segment: "GIR", Mandatory: true, MaxRepeat: 1},  // Related identification numbers
							{Segment: "GEI", Mandatory: false, MaxRepeat: 9}, // Processing information
							{Segment: "NAD", Mandatory: false, MaxRepeat: 1}, // Name and address
							{Segment: "MEA", Mandatory: false, MaxRepeat: 9}, // Measurements
							{ // Segment group 30
								Group: []SchemaNode{
									{Segment: "MOA", Mandatory: true, MaxRepeat: 1}, // Monetary amount
									{ // Segment group 31
										Group: []SchemaNode{
											{Segment: "CUX", Mandatory: true, MaxRepeat: 1},  // Currencies
											{Segment: "DTM", Mandatory: false, MaxRepeat: 1}, // Date/time/period
										},
										Mandatory: false, MaxRepeat: 1,
									},
								},
								Mandatory: false, MaxRepeat: 99,
							},
							{ // Segment group 32
								Group: []SchemaNode{
									{Segment: "TAX", Mandatory: true, MaxRepeat: 1},  // Duty/tax/fee details
									{Segment: "MOA", Mandatory: false, MaxRepeat: 9}, // Monetary amount
									{Segment: "GEI", Mandatory: false, MaxRepeat: 1}, // Processing information
								},
								Mandatory: false, MaxRepeat: 99,
							},
							{ // Segment group 33
								Group: []SchemaNode{
									{Segment: "DOC", Mandatory: true, MaxRepeat: 1},  // Document/message details
									{Segment: "DTM", Mandatory: false, MaxRepeat: 9}, // Date/time/period
									{Segment: "LOC", Mandatory: false, MaxRepeat: 1}, // Place/location identification
									{Segment: "NAD", Mandatory: false, MaxRepeat: 1}, // Name and address
								},
								Mandatory: false, MaxRepeat: 99,
							},
						},
						Mandatory: false, MaxRepeat: 99,
					},
				},
				Mandatory: false, MaxRepeat: 9999,
			},
		},
		Mandatory: false, MaxRepeat: 9999,
	},
	{ // Segment group 34
		Group: []SchemaNode{
			{Segment: "AUT", Mandatory: true, MaxRepeat: 1},  // Authentication result
			{Segment: "DTM", Mandatory: false, MaxRepeat: 9}, // Date/time/period
		},
		Mandatory: false, MaxRepeat: 1,
	},
}}
