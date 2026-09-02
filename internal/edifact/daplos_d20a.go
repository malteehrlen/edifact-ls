package edifact

// DAPLOS D.20A message specification: the structural (segment/group
// presence, order, cardinality) branching diagram for this message,
// UN/EDIFACT directory release D.20A. Scope is structural only -- not
// element-level content -- per edifact-ls-3uzr's epic non-goal.
// 21 segment groups, max nesting depth 4.
//
// Source: https://service.unece.org/trade/untdid/d20a/trmd/daplos_c.htm
// section 4.3.1 "Segment table". That URL currently returns HTTP 403 from
// Cloudflare when fetched directly; this data was transcribed from the
// Wayback Machine's archived copy instead:
// http://web.archive.org/web/20230608130237/https://service.unece.org/trade/untdid/d20a/trmd/daplos_c.htm
// -- re-check the direct URL first if cross-checking later, in case the
// block has lifted.
//
// Transcription note: generated in bulk (edifact-ls-13gu) the same way
// as every other message-type schema -- a script parsed the source's
// exact rail-art column positions mechanically and verified the result
// balances (all 21 groups open and close consistently)
// before being transcribed here, rather than reading the ASCII nesting
// by eye. Not individually spot-checked against the raw source the way
// the first several message types in this project were -- see
// edifact-ls-13gu for which representative samples were spot-checked
// for this batch.

func init() {
	RegisterSchema(
		MessageID{Type: "DAPLOS", Version: "D", Release: "20A", Agency: "UN"},
		daplosD20aSchema,
		"https://service.unece.org/trade/untdid/d20a/trmd/daplos_c.htm",
	)
}

var daplosD20aSchema = Schema{Nodes: []SchemaNode{
	{Segment: "BGM", Mandatory: true, MaxRepeat: 1},  // Beginning of message
	{Segment: "DTM", Mandatory: true, MaxRepeat: 9},  // Date/time/period
	{Segment: "IMD", Mandatory: false, MaxRepeat: 1}, // Item description
	{ // Segment group 1
		Group: []SchemaNode{
			{Segment: "NAD", Mandatory: true, MaxRepeat: 1},  // Name and address
			{Segment: "RFF", Mandatory: false, MaxRepeat: 9}, // Reference
		},
		Mandatory: false, MaxRepeat: 9,
	},
	{ // Segment group 2
		Group: []SchemaNode{
			{Segment: "LIN", Mandatory: true, MaxRepeat: 1},  // Line item
			{Segment: "DTM", Mandatory: true, MaxRepeat: 9},  // Date/time/period
			{Segment: "PIA", Mandatory: true, MaxRepeat: 9},  // Additional product id
			{Segment: "QTY", Mandatory: false, MaxRepeat: 1}, // Quantity
			{ // Segment group 3
				Group: []SchemaNode{
					{Segment: "CCI", Mandatory: true, MaxRepeat: 1},  // Characteristic/class id
					{Segment: "CAV", Mandatory: false, MaxRepeat: 1}, // Characteristic value
				},
				Mandatory: false, MaxRepeat: 9,
			},
			{ // Segment group 4
				Group: []SchemaNode{
					{Segment: "RFF", Mandatory: true, MaxRepeat: 1}, // Reference
					{ // Segment group 5
						Group: []SchemaNode{
							{Segment: "MEA", Mandatory: true, MaxRepeat: 1},     // Measurements
							{Segment: "GPO", Mandatory: false, MaxRepeat: 9999}, // Geographical position
						},
						Mandatory: false, MaxRepeat: 9,
					},
				},
				Mandatory: false, MaxRepeat: 99,
			},
			{ // Segment group 6
				Group: []SchemaNode{
					{Segment: "IMD", Mandatory: true, MaxRepeat: 1},  // Item description
					{Segment: "QTY", Mandatory: false, MaxRepeat: 1}, // Quantity
					{Segment: "NAD", Mandatory: false, MaxRepeat: 1}, // Name and address
					{ // Segment group 7
						Group: []SchemaNode{
							{Segment: "RFF", Mandatory: true, MaxRepeat: 1},  // Reference
							{Segment: "DTM", Mandatory: false, MaxRepeat: 1}, // Date/time/period
						},
						Mandatory: false, MaxRepeat: 9,
					},
				},
				Mandatory: false, MaxRepeat: 99,
			},
			{ // Segment group 8
				Group: []SchemaNode{
					{Segment: "HYN", Mandatory: true, MaxRepeat: 1},  // Hierarchy information
					{Segment: "PIA", Mandatory: false, MaxRepeat: 9}, // Additional product id
					{ // Segment group 9
						Group: []SchemaNode{
							{Segment: "CCI", Mandatory: true, MaxRepeat: 1},  // Characteristic/class id
							{Segment: "CAV", Mandatory: false, MaxRepeat: 1}, // Characteristic value
						},
						Mandatory: false, MaxRepeat: 9,
					},
					{ // Segment group 10
						Group: []SchemaNode{
							{Segment: "IMD", Mandatory: true, MaxRepeat: 1},  // Item description
							{Segment: "DTM", Mandatory: false, MaxRepeat: 1}, // Date/time/period
							{Segment: "QTY", Mandatory: false, MaxRepeat: 1}, // Quantity
							{Segment: "NAD", Mandatory: false, MaxRepeat: 1}, // Name and address
						},
						Mandatory: false, MaxRepeat: 9,
					},
				},
				Mandatory: false, MaxRepeat: 9,
			},
			{ // Segment group 11
				Group: []SchemaNode{
					{Segment: "DOC", Mandatory: true, MaxRepeat: 1},  // Document/message details
					{Segment: "NAD", Mandatory: false, MaxRepeat: 1}, // Name and address
					{Segment: "DTM", Mandatory: false, MaxRepeat: 9}, // Date/time/period
				},
				Mandatory: false, MaxRepeat: 99,
			},
			{ // Segment group 12
				Group: []SchemaNode{
					{Segment: "EVE", Mandatory: true, MaxRepeat: 1},  // Event
					{Segment: "DTM", Mandatory: false, MaxRepeat: 9}, // Date/time/period
					{Segment: "IMD", Mandatory: false, MaxRepeat: 9}, // Item description
					{Segment: "NAD", Mandatory: false, MaxRepeat: 1}, // Name and address
					{Segment: "FTX", Mandatory: false, MaxRepeat: 9}, // Free text
					{ // Segment group 13
						Group: []SchemaNode{
							{Segment: "CCI", Mandatory: true, MaxRepeat: 1},   // Characteristic/class id
							{Segment: "CAV", Mandatory: false, MaxRepeat: 99}, // Characteristic value
						},
						Mandatory: false, MaxRepeat: 1,
					},
					{ // Segment group 14
						Group: []SchemaNode{
							{Segment: "MEA", Mandatory: true, MaxRepeat: 1},     // Measurements
							{Segment: "GPO", Mandatory: false, MaxRepeat: 9999}, // Geographical position
						},
						Mandatory: false, MaxRepeat: 9,
					},
					{ // Segment group 15
						Group: []SchemaNode{
							{Segment: "IFD", Mandatory: true, MaxRepeat: 1},  // Information detail
							{Segment: "PIA", Mandatory: false, MaxRepeat: 9}, // Additional product id
							{Segment: "QTY", Mandatory: false, MaxRepeat: 9}, // Quantity
							{Segment: "NAD", Mandatory: false, MaxRepeat: 1}, // Name and address
							{ // Segment group 16
								Group: []SchemaNode{
									{Segment: "COD", Mandatory: true, MaxRepeat: 1},  // Component details
									{Segment: "MEA", Mandatory: false, MaxRepeat: 1}, // Measurements
								},
								Mandatory: false, MaxRepeat: 99,
							},
							{ // Segment group 17
								Group: []SchemaNode{
									{Segment: "CCI", Mandatory: true, MaxRepeat: 1},  // Characteristic/class id
									{Segment: "CAV", Mandatory: false, MaxRepeat: 1}, // Characteristic value
								},
								Mandatory: false, MaxRepeat: 9,
							},
							{ // Segment group 18
								Group: []SchemaNode{
									{Segment: "RFF", Mandatory: true, MaxRepeat: 1},  // Reference
									{Segment: "QTY", Mandatory: false, MaxRepeat: 9}, // Quantity
									{Segment: "MEA", Mandatory: false, MaxRepeat: 1}, // Measurements
								},
								Mandatory: false, MaxRepeat: 999,
							},
							{ // Segment group 19
								Group: []SchemaNode{
									{Segment: "DOC", Mandatory: true, MaxRepeat: 1},  // Document/message details
									{Segment: "NAD", Mandatory: false, MaxRepeat: 1}, // Name and address
									{Segment: "DTM", Mandatory: false, MaxRepeat: 9}, // Date/time/period
								},
								Mandatory: false, MaxRepeat: 99,
							},
						},
						Mandatory: false, MaxRepeat: 99,
					},
					{ // Segment group 20
						Group: []SchemaNode{
							{Segment: "HYN", Mandatory: true, MaxRepeat: 1},  // Hierarchy information
							{Segment: "NAD", Mandatory: false, MaxRepeat: 1}, // Name and address
							{Segment: "FTX", Mandatory: false, MaxRepeat: 9}, // Free text
							{ // Segment group 21
								Group: []SchemaNode{
									{Segment: "RFF", Mandatory: true, MaxRepeat: 1},  // Reference
									{Segment: "DTM", Mandatory: false, MaxRepeat: 1}, // Date/time/period
								},
								Mandatory: false, MaxRepeat: 9,
							},
						},
						Mandatory: false, MaxRepeat: 1,
					},
				},
				Mandatory: false, MaxRepeat: 999,
			},
		},
		Mandatory: false, MaxRepeat: 9999,
	},
}}
