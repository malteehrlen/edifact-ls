package edifact

// IFTSTA D.20A message specification: the structural (segment/group
// presence, order, cardinality) branching diagram for the International Multimodal Status Report
// message, UN/EDIFACT directory release D.20A. Scope is structural only --
// not element-level content -- per edifact-ls-3uzr's epic non-goal.
// 29 segment groups, max nesting depth 5.
//
// Source: https://service.unece.org/trade/untdid/d20a/trmd/iftsta_c.htm
// section 4.3.1 "Segment table". That URL currently returns HTTP 403 from
// Cloudflare when fetched directly; this data was transcribed from the
// Wayback Machine's archived copy instead:
// http://web.archive.org/web/20250710164050/https://service.unece.org/trade/untdid/d20a/trmd/iftsta_c.htm
// -- re-check the direct URL first if cross-checking later, in case the
// block has lifted.
//
// Transcription note: generated the same way as iftmcs_d21a.go -- a
// one-off script parsed the source's exact rail-art column positions
// mechanically and verified the result balances (all 29 groups
// open and close consistently) before being transcribed here, rather than
// reading the ASCII nesting by eye.

func init() {
	RegisterSchema(
		MessageID{Type: "IFTSTA", Version: "D", Release: "20A", Agency: "UN"},
		iftstaD20aSchema,
		"https://service.unece.org/trade/untdid/d20a/trmd/iftsta_c.htm",
	)
}

var iftstaD20aSchema = Schema{Nodes: []SchemaNode{
	{Segment: "BGM", Mandatory: true, MaxRepeat: 1},  // Beginning of message
	{Segment: "DTM", Mandatory: false, MaxRepeat: 9}, // Date/time/period
	{Segment: "TSR", Mandatory: false, MaxRepeat: 1}, // Transport service requirements
	{ // Segment group 1
		Group: []SchemaNode{
			{Segment: "NAD", Mandatory: true, MaxRepeat: 1}, // Name and address
			{ // Segment group 2
				Group: []SchemaNode{
					{Segment: "CTA", Mandatory: true, MaxRepeat: 1},  // Contact information
					{Segment: "COM", Mandatory: false, MaxRepeat: 9}, // Communication contact
				},
				Mandatory: false, MaxRepeat: 9,
			},
		},
		Mandatory: false, MaxRepeat: 9,
	},
	{ // Segment group 3
		Group: []SchemaNode{
			{Segment: "RFF", Mandatory: true, MaxRepeat: 1},  // Reference
			{Segment: "DTM", Mandatory: false, MaxRepeat: 1}, // Date/time/period
		},
		Mandatory: false, MaxRepeat: 9,
	},
	{Segment: "LOC", Mandatory: false, MaxRepeat: 9}, // Place/location identification
	{Segment: "FTX", Mandatory: false, MaxRepeat: 9}, // Free text
	{Segment: "CNT", Mandatory: false, MaxRepeat: 9}, // Control total
	{ // Segment group 4
		Group: []SchemaNode{
			{Segment: "EQD", Mandatory: true, MaxRepeat: 1},    // Equipment details
			{Segment: "NAD", Mandatory: false, MaxRepeat: 9},   // Name and address
			{Segment: "DTM", Mandatory: false, MaxRepeat: 9},   // Date/time/period
			{Segment: "FTX", Mandatory: false, MaxRepeat: 9},   // Free text
			{Segment: "MEA", Mandatory: false, MaxRepeat: 9},   // Measurements
			{Segment: "DIM", Mandatory: false, MaxRepeat: 9},   // Dimensions
			{Segment: "SEL", Mandatory: false, MaxRepeat: 9},   // Seal number
			{Segment: "RFF", Mandatory: false, MaxRepeat: 999}, // Reference
			{Segment: "TPL", Mandatory: false, MaxRepeat: 9},   // Transport placement
			{Segment: "TMD", Mandatory: false, MaxRepeat: 1},   // Transport movement details
			{ // Segment group 5
				Group: []SchemaNode{
					{Segment: "DGS", Mandatory: true, MaxRepeat: 1},  // Dangerous goods
					{Segment: "FTX", Mandatory: false, MaxRepeat: 9}, // Free text
				},
				Mandatory: false, MaxRepeat: 9,
			},
			{ // Segment group 6
				Group: []SchemaNode{
					{Segment: "LOC", Mandatory: true, MaxRepeat: 1},  // Place/location identification
					{Segment: "DTM", Mandatory: false, MaxRepeat: 9}, // Date/time/period
				},
				Mandatory: false, MaxRepeat: 9,
			},
			{ // Segment group 7
				Group: []SchemaNode{
					{Segment: "STS", Mandatory: true, MaxRepeat: 1},  // Status
					{Segment: "LOC", Mandatory: false, MaxRepeat: 1}, // Place/location identification
					{Segment: "DTM", Mandatory: false, MaxRepeat: 1}, // Date/time/period
				},
				Mandatory: false, MaxRepeat: 99,
			},
			{ // Segment group 8
				Group: []SchemaNode{
					{Segment: "EQA", Mandatory: true, MaxRepeat: 1},  // Attached equipment
					{Segment: "SEL", Mandatory: false, MaxRepeat: 9}, // Seal number
					{ // Segment group 9
						Group: []SchemaNode{
							{Segment: "LOC", Mandatory: true, MaxRepeat: 1},  // Place/location identification
							{Segment: "DTM", Mandatory: false, MaxRepeat: 9}, // Date/time/period
						},
						Mandatory: false, MaxRepeat: 9,
					},
				},
				Mandatory: false, MaxRepeat: 99,
			},
			{ // Segment group 10
				Group: []SchemaNode{
					{Segment: "TDT", Mandatory: true, MaxRepeat: 1},  // Transport information
					{Segment: "DTM", Mandatory: false, MaxRepeat: 9}, // Date/time/period
					{Segment: "RFF", Mandatory: false, MaxRepeat: 9}, // Reference
					{ // Segment group 11
						Group: []SchemaNode{
							{Segment: "LOC", Mandatory: true, MaxRepeat: 1},  // Place/location identification
							{Segment: "DTM", Mandatory: false, MaxRepeat: 9}, // Date/time/period
						},
						Mandatory: false, MaxRepeat: 9,
					},
				},
				Mandatory: false, MaxRepeat: 99,
			},
			{ // Segment group 12
				Group: []SchemaNode{
					{Segment: "TMP", Mandatory: true, MaxRepeat: 1},  // Temperature
					{Segment: "RNG", Mandatory: false, MaxRepeat: 1}, // Range details
				},
				Mandatory: false, MaxRepeat: 1,
			},
			{ // Segment group 13
				Group: []SchemaNode{
					{Segment: "DAM", Mandatory: true, MaxRepeat: 1},  // Damage
					{Segment: "COD", Mandatory: false, MaxRepeat: 9}, // Component details
				},
				Mandatory: false, MaxRepeat: 9,
			},
		},
		Mandatory: false, MaxRepeat: 99999,
	},
	{ // Segment group 14
		Group: []SchemaNode{
			{Segment: "CNI", Mandatory: true, MaxRepeat: 1},  // Consignment information
			{Segment: "LOC", Mandatory: false, MaxRepeat: 9}, // Place/location identification
			{Segment: "CNT", Mandatory: false, MaxRepeat: 9}, // Control total
			{ // Segment group 15
				Group: []SchemaNode{
					{Segment: "STS", Mandatory: true, MaxRepeat: 1},    // Status
					{Segment: "RFF", Mandatory: false, MaxRepeat: 999}, // Reference
					{Segment: "DTM", Mandatory: false, MaxRepeat: 9},   // Date/time/period
					{Segment: "DOC", Mandatory: false, MaxRepeat: 1},   // Document/message details
					{Segment: "FTX", Mandatory: false, MaxRepeat: 9},   // Free text
					{Segment: "MOA", Mandatory: false, MaxRepeat: 9},   // Monetary amount
					{ // Segment group 16
						Group: []SchemaNode{
							{Segment: "EFI", Mandatory: true, MaxRepeat: 1},   // External file link identification
							{Segment: "CED", Mandatory: false, MaxRepeat: 99}, // Computer environment details
							{Segment: "COM", Mandatory: false, MaxRepeat: 9},  // Communication contact
							{Segment: "RFF", Mandatory: false, MaxRepeat: 9},  // Reference
							{Segment: "DTM", Mandatory: false, MaxRepeat: 9},  // Date/time/period
							{Segment: "QTY", Mandatory: false, MaxRepeat: 9},  // Quantity
						},
						Mandatory: false, MaxRepeat: 99,
					},
					{ // Segment group 17
						Group: []SchemaNode{
							{Segment: "NAD", Mandatory: true, MaxRepeat: 1}, // Name and address
							{ // Segment group 18
								Group: []SchemaNode{
									{Segment: "CTA", Mandatory: true, MaxRepeat: 1},  // Contact information
									{Segment: "COM", Mandatory: false, MaxRepeat: 9}, // Communication contact
								},
								Mandatory: false, MaxRepeat: 9,
							},
						},
						Mandatory: false, MaxRepeat: 9,
					},
					{Segment: "LOC", Mandatory: false, MaxRepeat: 1},  // Place/location identification
					{Segment: "PCI", Mandatory: false, MaxRepeat: 99}, // Package identification
					{ // Segment group 19
						Group: []SchemaNode{
							{Segment: "TDT", Mandatory: true, MaxRepeat: 1},  // Transport information
							{Segment: "DTM", Mandatory: false, MaxRepeat: 9}, // Date/time/period
							{Segment: "RFF", Mandatory: false, MaxRepeat: 9}, // Reference
							{ // Segment group 20
								Group: []SchemaNode{
									{Segment: "LOC", Mandatory: true, MaxRepeat: 1},  // Place/location identification
									{Segment: "DTM", Mandatory: false, MaxRepeat: 9}, // Date/time/period
								},
								Mandatory: false, MaxRepeat: 9,
							},
						},
						Mandatory: false, MaxRepeat: 99,
					},
					{ // Segment group 21
						Group: []SchemaNode{
							{Segment: "EQD", Mandatory: true, MaxRepeat: 1},  // Equipment details
							{Segment: "MEA", Mandatory: false, MaxRepeat: 9}, // Measurements
							{Segment: "DIM", Mandatory: false, MaxRepeat: 9}, // Dimensions
							{Segment: "SEL", Mandatory: false, MaxRepeat: 9}, // Seal number
							{Segment: "RFF", Mandatory: false, MaxRepeat: 9}, // Reference
							{Segment: "TPL", Mandatory: false, MaxRepeat: 9}, // Transport placement
							{Segment: "TMD", Mandatory: false, MaxRepeat: 1}, // Transport movement details
							{ // Segment group 22
								Group: []SchemaNode{
									{Segment: "LOC", Mandatory: true, MaxRepeat: 1},  // Place/location identification
									{Segment: "DTM", Mandatory: false, MaxRepeat: 9}, // Date/time/period
								},
								Mandatory: false, MaxRepeat: 9,
							},
							{ // Segment group 23
								Group: []SchemaNode{
									{Segment: "EQA", Mandatory: true, MaxRepeat: 1},  // Attached equipment
									{Segment: "SEL", Mandatory: false, MaxRepeat: 9}, // Seal number
									{ // Segment group 24
										Group: []SchemaNode{
											{Segment: "LOC", Mandatory: true, MaxRepeat: 1},  // Place/location identification
											{Segment: "DTM", Mandatory: false, MaxRepeat: 9}, // Date/time/period
										},
										Mandatory: false, MaxRepeat: 9,
									},
								},
								Mandatory: false, MaxRepeat: 99,
							},
						},
						Mandatory: false, MaxRepeat: 99,
					},
					{ // Segment group 25
						Group: []SchemaNode{
							{Segment: "GID", Mandatory: true, MaxRepeat: 1},   // Goods item details
							{Segment: "HAN", Mandatory: false, MaxRepeat: 9},  // Handling instructions
							{Segment: "SGP", Mandatory: false, MaxRepeat: 99}, // Split goods placement
							{Segment: "DGS", Mandatory: false, MaxRepeat: 9},  // Dangerous goods
							{Segment: "FTX", Mandatory: false, MaxRepeat: 9},  // Free text
							{Segment: "GDS", Mandatory: false, MaxRepeat: 9},  // Nature of cargo
							{Segment: "LOC", Mandatory: false, MaxRepeat: 9},  // Place/location identification
							{Segment: "QTY", Mandatory: false, MaxRepeat: 9},  // Quantity
							{ // Segment group 26
								Group: []SchemaNode{
									{Segment: "MEA", Mandatory: true, MaxRepeat: 1},  // Measurements
									{Segment: "EQN", Mandatory: false, MaxRepeat: 1}, // Number of units
								},
								Mandatory: false, MaxRepeat: 99,
							},
							{ // Segment group 27
								Group: []SchemaNode{
									{Segment: "DIM", Mandatory: true, MaxRepeat: 1},  // Dimensions
									{Segment: "EQN", Mandatory: false, MaxRepeat: 1}, // Number of units
								},
								Mandatory: false, MaxRepeat: 99,
							},
							{ // Segment group 28
								Group: []SchemaNode{
									{Segment: "RFF", Mandatory: true, MaxRepeat: 1},  // Reference
									{Segment: "DTM", Mandatory: false, MaxRepeat: 9}, // Date/time/period
								},
								Mandatory: false, MaxRepeat: 9,
							},
							{ // Segment group 29
								Group: []SchemaNode{
									{Segment: "PCI", Mandatory: true, MaxRepeat: 1},     // Package identification
									{Segment: "GIN", Mandatory: false, MaxRepeat: 9999}, // Goods identity number
								},
								Mandatory: false, MaxRepeat: 99,
							},
						},
						Mandatory: false, MaxRepeat: 9999,
					},
				},
				Mandatory: true, MaxRepeat: 99,
			},
		},
		Mandatory: false, MaxRepeat: 99999,
	},
}}
