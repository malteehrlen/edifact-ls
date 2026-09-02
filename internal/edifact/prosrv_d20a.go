package edifact

// PROSRV D.20A message specification: the structural (segment/group
// presence, order, cardinality) branching diagram for this message,
// UN/EDIFACT directory release D.20A. Scope is structural only -- not
// element-level content -- per edifact-ls-3uzr's epic non-goal.
// 19 segment groups, max nesting depth 5.
//
// Source: https://service.unece.org/trade/untdid/d20a/trmd/prosrv_c.htm
// section 4.3.1 "Segment table". That URL currently returns HTTP 403 from
// Cloudflare when fetched directly; this data was transcribed from the
// Wayback Machine's archived copy instead:
// http://web.archive.org/web/20231207000557/https://service.unece.org/trade/untdid/d20a/trmd/prosrv_c.htm
// -- re-check the direct URL first if cross-checking later, in case the
// block has lifted.
//
// Transcription note: generated in bulk (edifact-ls-13gu) the same way
// as every other message-type schema -- a script parsed the source's
// exact rail-art column positions mechanically and verified the result
// balances (all 19 groups open and close consistently)
// before being transcribed here, rather than reading the ASCII nesting
// by eye. Not individually spot-checked against the raw source the way
// the first several message types in this project were -- see
// edifact-ls-13gu for which representative samples were spot-checked
// for this batch.

func init() {
	RegisterSchema(
		MessageID{Type: "PROSRV", Version: "D", Release: "20A", Agency: "UN"},
		prosrvD20aSchema,
		"https://service.unece.org/trade/untdid/d20a/trmd/prosrv_c.htm",
	)
}

var prosrvD20aSchema = Schema{Nodes: []SchemaNode{
	{Segment: "BGM", Mandatory: true, MaxRepeat: 1},  // Beginning of message
	{Segment: "DTM", Mandatory: true, MaxRepeat: 99}, // Date/time/period
	{Segment: "CUX", Mandatory: false, MaxRepeat: 9}, // Currencies
	{ // Segment group 1
		Group: []SchemaNode{
			{Segment: "RFF", Mandatory: true, MaxRepeat: 1},   // Reference
			{Segment: "DTM", Mandatory: false, MaxRepeat: 99}, // Date/time/period
		},
		Mandatory: false, MaxRepeat: 9999,
	},
	{ // Segment group 2
		Group: []SchemaNode{
			{Segment: "NAD", Mandatory: true, MaxRepeat: 1},   // Name and address
			{Segment: "FII", Mandatory: false, MaxRepeat: 99}, // Financial institution information
			{ // Segment group 3
				Group: []SchemaNode{
					{Segment: "RFF", Mandatory: true, MaxRepeat: 1},  // Reference
					{Segment: "DTM", Mandatory: false, MaxRepeat: 9}, // Date/time/period
				},
				Mandatory: false, MaxRepeat: 99,
			},
			{ // Segment group 4
				Group: []SchemaNode{
					{Segment: "CTA", Mandatory: true, MaxRepeat: 1},  // Contact information
					{Segment: "COM", Mandatory: false, MaxRepeat: 9}, // Communication contact
				},
				Mandatory: false, MaxRepeat: 99,
			},
		},
		Mandatory: false, MaxRepeat: 999,
	},
	{ // Segment group 5
		Group: []SchemaNode{
			{Segment: "DOC", Mandatory: true, MaxRepeat: 1},  // Document/message details
			{Segment: "DTM", Mandatory: false, MaxRepeat: 9}, // Date/time/period
			{Segment: "STS", Mandatory: false, MaxRepeat: 9}, // Status
			{Segment: "MOA", Mandatory: false, MaxRepeat: 9}, // Monetary amount
			{Segment: "FTX", Mandatory: false, MaxRepeat: 9}, // Free text
			{ // Segment group 6
				Group: []SchemaNode{
					{Segment: "LIN", Mandatory: true, MaxRepeat: 1},   // Line item
					{Segment: "PIA", Mandatory: false, MaxRepeat: 99}, // Additional product id
					{Segment: "IMD", Mandatory: false, MaxRepeat: 99}, // Item description
					{Segment: "QTY", Mandatory: false, MaxRepeat: 99}, // Quantity
					{Segment: "DTM", Mandatory: false, MaxRepeat: 99}, // Date/time/period
					{Segment: "ALI", Mandatory: false, MaxRepeat: 9},  // Additional information
					{Segment: "GIR", Mandatory: false, MaxRepeat: 99}, // Related identification numbers
					{ // Segment group 7
						Group: []SchemaNode{
							{Segment: "RFF", Mandatory: true, MaxRepeat: 1},  // Reference
							{Segment: "DTM", Mandatory: false, MaxRepeat: 9}, // Date/time/period
						},
						Mandatory: false, MaxRepeat: 999,
					},
					{ // Segment group 8
						Group: []SchemaNode{
							{Segment: "NAD", Mandatory: true, MaxRepeat: 1}, // Name and address
							{ // Segment group 9
								Group: []SchemaNode{
									{Segment: "RFF", Mandatory: true, MaxRepeat: 1},  // Reference
									{Segment: "DTM", Mandatory: false, MaxRepeat: 9}, // Date/time/period
								},
								Mandatory: false, MaxRepeat: 99,
							},
							{ // Segment group 10
								Group: []SchemaNode{
									{Segment: "CTA", Mandatory: true, MaxRepeat: 1},  // Contact information
									{Segment: "COM", Mandatory: false, MaxRepeat: 9}, // Communication contact
								},
								Mandatory: false, MaxRepeat: 9,
							},
						},
						Mandatory: false, MaxRepeat: 999,
					},
					{ // Segment group 11
						Group: []SchemaNode{
							{Segment: "CCI", Mandatory: true, MaxRepeat: 1},   // Characteristic/class id
							{Segment: "CAV", Mandatory: false, MaxRepeat: 99}, // Characteristic value
							{Segment: "FTX", Mandatory: false, MaxRepeat: 99}, // Free text
						},
						Mandatory: false, MaxRepeat: 9999,
					},
					{ // Segment group 12
						Group: []SchemaNode{
							{Segment: "ALC", Mandatory: true, MaxRepeat: 1},  // Allowance or charge
							{Segment: "QTY", Mandatory: false, MaxRepeat: 9}, // Quantity
							{Segment: "MOA", Mandatory: false, MaxRepeat: 9}, // Monetary amount
							{Segment: "RTE", Mandatory: false, MaxRepeat: 9}, // Rate details
							{ // Segment group 13
								Group: []SchemaNode{
									{Segment: "TAX", Mandatory: true, MaxRepeat: 1},  // Duty/tax/fee details
									{Segment: "MOA", Mandatory: false, MaxRepeat: 9}, // Monetary amount
								},
								Mandatory: false, MaxRepeat: 9,
							},
						},
						Mandatory: false, MaxRepeat: 999,
					},
					{ // Segment group 14
						Group: []SchemaNode{
							{Segment: "HYN", Mandatory: true, MaxRepeat: 1},  // Hierarchy information
							{Segment: "PIA", Mandatory: false, MaxRepeat: 9}, // Additional product id
							{Segment: "IMD", Mandatory: false, MaxRepeat: 9}, // Item description
							{Segment: "QTY", Mandatory: false, MaxRepeat: 9}, // Quantity
							{Segment: "ALI", Mandatory: false, MaxRepeat: 9}, // Additional information
							{Segment: "PRI", Mandatory: false, MaxRepeat: 9}, // Price details
							{ // Segment group 15
								Group: []SchemaNode{
									{Segment: "RFF", Mandatory: true, MaxRepeat: 1},  // Reference
									{Segment: "DTM", Mandatory: false, MaxRepeat: 9}, // Date/time/period
								},
								Mandatory: false, MaxRepeat: 99,
							},
							{ // Segment group 16
								Group: []SchemaNode{
									{Segment: "ALC", Mandatory: true, MaxRepeat: 1},  // Allowance or charge
									{Segment: "QTY", Mandatory: false, MaxRepeat: 9}, // Quantity
									{Segment: "MOA", Mandatory: false, MaxRepeat: 9}, // Monetary amount
									{Segment: "RTE", Mandatory: false, MaxRepeat: 9}, // Rate details
									{ // Segment group 17
										Group: []SchemaNode{
											{Segment: "TAX", Mandatory: true, MaxRepeat: 1},  // Duty/tax/fee details
											{Segment: "MOA", Mandatory: false, MaxRepeat: 9}, // Monetary amount
										},
										Mandatory: false, MaxRepeat: 9,
									},
								},
								Mandatory: false, MaxRepeat: 999,
							},
							{ // Segment group 18
								Group: []SchemaNode{
									{Segment: "CCI", Mandatory: true, MaxRepeat: 1},   // Characteristic/class id
									{Segment: "CAV", Mandatory: false, MaxRepeat: 99}, // Characteristic value
									{Segment: "FTX", Mandatory: false, MaxRepeat: 99}, // Free text
								},
								Mandatory: false, MaxRepeat: 9999,
							},
						},
						Mandatory: false, MaxRepeat: 999,
					},
				},
				Mandatory: false, MaxRepeat: 99999,
			},
		},
		Mandatory: false, MaxRepeat: 99999,
	},
	{Segment: "UNS", Mandatory: true, MaxRepeat: 1},   // Section control
	{Segment: "CNT", Mandatory: false, MaxRepeat: 99}, // Control total
	{Segment: "MOA", Mandatory: false, MaxRepeat: 99}, // Monetary amount
	{ // Segment group 19
		Group: []SchemaNode{
			{Segment: "TAX", Mandatory: true, MaxRepeat: 1},  // Duty/tax/fee details
			{Segment: "MOA", Mandatory: false, MaxRepeat: 9}, // Monetary amount
		},
		Mandatory: false, MaxRepeat: 99,
	},
}}
