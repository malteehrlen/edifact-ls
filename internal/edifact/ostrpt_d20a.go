package edifact

// OSTRPT D.20A message specification: the structural (segment/group
// presence, order, cardinality) branching diagram for this message,
// UN/EDIFACT directory release D.20A. Scope is structural only -- not
// element-level content -- per edifact-ls-3uzr's epic non-goal.
// 18 segment groups, max nesting depth 5.
//
// Source: https://service.unece.org/trade/untdid/d20a/trmd/ostrpt_c.htm
// section 4.3.1 "Segment table". That URL currently returns HTTP 403 from
// Cloudflare when fetched directly; this data was transcribed from the
// Wayback Machine's archived copy instead:
// http://web.archive.org/web/20230924170318/https://service.unece.org/trade/untdid/d20a/trmd/ostrpt_c.htm
// -- re-check the direct URL first if cross-checking later, in case the
// block has lifted.
//
// Transcription note: generated in bulk (edifact-ls-13gu) the same way
// as every other message-type schema -- a script parsed the source's
// exact rail-art column positions mechanically and verified the result
// balances (all 18 groups open and close consistently)
// before being transcribed here, rather than reading the ASCII nesting
// by eye. Not individually spot-checked against the raw source the way
// the first several message types in this project were -- see
// edifact-ls-13gu for which representative samples were spot-checked
// for this batch.

func init() {
	RegisterSchema(
		MessageID{Type: "OSTRPT", Version: "D", Release: "20A", Agency: "UN"},
		ostrptD20aSchema,
		"https://service.unece.org/trade/untdid/d20a/trmd/ostrpt_c.htm",
	)
}

var ostrptD20aSchema = Schema{Nodes: []SchemaNode{
	{Segment: "BGM", Mandatory: true, MaxRepeat: 1},   // Beginning of message
	{Segment: "DTM", Mandatory: true, MaxRepeat: 5},   // Date/time/period
	{Segment: "RFF", Mandatory: false, MaxRepeat: 1},  // Reference
	{Segment: "IRQ", Mandatory: false, MaxRepeat: 1},  // Information required
	{Segment: "FTX", Mandatory: false, MaxRepeat: 99}, // Free text
	{ // Segment group 1
		Group: []SchemaNode{
			{Segment: "NAD", Mandatory: true, MaxRepeat: 1},    // Name and address
			{Segment: "LOC", Mandatory: false, MaxRepeat: 999}, // Place/location identification
			{ // Segment group 2
				Group: []SchemaNode{
					{Segment: "CTA", Mandatory: true, MaxRepeat: 1},  // Contact information
					{Segment: "COM", Mandatory: false, MaxRepeat: 5}, // Communication contact
				},
				Mandatory: false, MaxRepeat: 5,
			},
		},
		Mandatory: true, MaxRepeat: 99,
	},
	{ // Segment group 3
		Group: []SchemaNode{
			{Segment: "DOC", Mandatory: true, MaxRepeat: 1},   // Document/message details
			{Segment: "DTM", Mandatory: false, MaxRepeat: 5},  // Date/time/period
			{Segment: "CNT", Mandatory: false, MaxRepeat: 9},  // Control total
			{Segment: "RFF", Mandatory: false, MaxRepeat: 99}, // Reference
			{Segment: "TOD", Mandatory: false, MaxRepeat: 99}, // Terms of delivery or transport
			{ // Segment group 4
				Group: []SchemaNode{
					{Segment: "STS", Mandatory: true, MaxRepeat: 1},  // Status
					{Segment: "DTM", Mandatory: false, MaxRepeat: 5}, // Date/time/period
					{Segment: "PCD", Mandatory: false, MaxRepeat: 1}, // Percentage details
					{Segment: "QTY", Mandatory: false, MaxRepeat: 1}, // Quantity
					{ // Segment group 5
						Group: []SchemaNode{
							{Segment: "NAD", Mandatory: true, MaxRepeat: 1},  // Name and address
							{Segment: "LOC", Mandatory: false, MaxRepeat: 9}, // Place/location identification
						},
						Mandatory: false, MaxRepeat: 99,
					},
				},
				Mandatory: false, MaxRepeat: 99,
			},
			{ // Segment group 6
				Group: []SchemaNode{
					{Segment: "LIN", Mandatory: true, MaxRepeat: 1},   // Line item
					{Segment: "PIA", Mandatory: false, MaxRepeat: 99}, // Additional product id
					{Segment: "IMD", Mandatory: false, MaxRepeat: 99}, // Item description
					{Segment: "MEA", Mandatory: false, MaxRepeat: 5},  // Measurements
					{Segment: "GIN", Mandatory: false, MaxRepeat: 99}, // Goods identity number
					{Segment: "RCS", Mandatory: false, MaxRepeat: 5},  // Requirements and conditions
					{Segment: "FTX", Mandatory: false, MaxRepeat: 99}, // Free text
					{ // Segment group 7
						Group: []SchemaNode{
							{Segment: "RFF", Mandatory: true, MaxRepeat: 1},  // Reference
							{Segment: "DTM", Mandatory: false, MaxRepeat: 1}, // Date/time/period
						},
						Mandatory: false, MaxRepeat: 999,
					},
					{ // Segment group 8
						Group: []SchemaNode{
							{Segment: "STS", Mandatory: true, MaxRepeat: 1},  // Status
							{Segment: "DTM", Mandatory: false, MaxRepeat: 5}, // Date/time/period
							{Segment: "PCD", Mandatory: false, MaxRepeat: 1}, // Percentage details
							{Segment: "QTY", Mandatory: false, MaxRepeat: 1}, // Quantity
							{ // Segment group 9
								Group: []SchemaNode{
									{Segment: "NAD", Mandatory: true, MaxRepeat: 1},  // Name and address
									{Segment: "RFF", Mandatory: false, MaxRepeat: 5}, // Reference
									{Segment: "LOC", Mandatory: false, MaxRepeat: 9}, // Place/location identification
									{ // Segment group 10
										Group: []SchemaNode{
											{Segment: "CTA", Mandatory: true, MaxRepeat: 1},  // Contact information
											{Segment: "COM", Mandatory: false, MaxRepeat: 5}, // Communication contact
										},
										Mandatory: false, MaxRepeat: 5,
									},
								},
								Mandatory: false, MaxRepeat: 99,
							},
							{ // Segment group 11
								Group: []SchemaNode{
									{Segment: "TDT", Mandatory: true, MaxRepeat: 1},  // Transport information
									{Segment: "LOC", Mandatory: false, MaxRepeat: 1}, // Place/location identification
								},
								Mandatory: false, MaxRepeat: 5,
							},
							{ // Segment group 12
								Group: []SchemaNode{
									{Segment: "TOD", Mandatory: true, MaxRepeat: 1},  // Terms of delivery or transport
									{Segment: "LOC", Mandatory: false, MaxRepeat: 1}, // Place/location identification
								},
								Mandatory: false, MaxRepeat: 5,
							},
							{ // Segment group 13
								Group: []SchemaNode{
									{Segment: "EQD", Mandatory: true, MaxRepeat: 1},  // Equipment details
									{Segment: "HAN", Mandatory: false, MaxRepeat: 5}, // Handling instructions
								},
								Mandatory: false, MaxRepeat: 5,
							},
							{ // Segment group 14
								Group: []SchemaNode{
									{Segment: "PAC", Mandatory: true, MaxRepeat: 1},  // Package
									{Segment: "QTY", Mandatory: false, MaxRepeat: 1}, // Quantity
									{ // Segment group 15
										Group: []SchemaNode{
											{Segment: "PCI", Mandatory: true, MaxRepeat: 1},   // Package identification
											{Segment: "RFF", Mandatory: false, MaxRepeat: 99}, // Reference
											{Segment: "GIN", Mandatory: false, MaxRepeat: 99}, // Goods identity number
										},
										Mandatory: false, MaxRepeat: 10,
									},
								},
								Mandatory: false, MaxRepeat: 1,
							},
						},
						Mandatory: false, MaxRepeat: 99,
					},
					{ // Segment group 16
						Group: []SchemaNode{
							{Segment: "SCC", Mandatory: true, MaxRepeat: 1}, // Scheduling conditions
							{ // Segment group 17
								Group: []SchemaNode{
									{Segment: "LOC", Mandatory: true, MaxRepeat: 1},  // Place/location identification
									{Segment: "DTM", Mandatory: false, MaxRepeat: 5}, // Date/time/period
									{Segment: "QTY", Mandatory: false, MaxRepeat: 1}, // Quantity
									{ // Segment group 18
										Group: []SchemaNode{
											{Segment: "STS", Mandatory: true, MaxRepeat: 1},  // Status
											{Segment: "DTM", Mandatory: false, MaxRepeat: 5}, // Date/time/period
											{Segment: "PCD", Mandatory: false, MaxRepeat: 1}, // Percentage details
											{Segment: "QTY", Mandatory: false, MaxRepeat: 1}, // Quantity
											{Segment: "NAD", Mandatory: false, MaxRepeat: 1}, // Name and address
										},
										Mandatory: false, MaxRepeat: 99,
									},
								},
								Mandatory: false, MaxRepeat: 99,
							},
						},
						Mandatory: false, MaxRepeat: 99,
					},
				},
				Mandatory: false, MaxRepeat: 999999,
			},
		},
		Mandatory: true, MaxRepeat: 999,
	},
}}
