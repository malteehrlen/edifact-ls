package edifact

// RETINS D.20A message specification: the structural (segment/group
// presence, order, cardinality) branching diagram for this message,
// UN/EDIFACT directory release D.20A. Scope is structural only -- not
// element-level content -- per edifact-ls-3uzr's epic non-goal.
// 21 segment groups, max nesting depth 3.
//
// Source: https://service.unece.org/trade/untdid/d20a/trmd/retins_c.htm
// section 4.3.1 "Segment table". That URL currently returns HTTP 403 from
// Cloudflare when fetched directly; this data was transcribed from the
// Wayback Machine's archived copy instead:
// http://web.archive.org/web/20231002205711/https://service.unece.org/trade/untdid/d20a/trmd/retins_c.htm
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
		MessageID{Type: "RETINS", Version: "D", Release: "20A", Agency: "UN"},
		retinsD20aSchema,
		"https://service.unece.org/trade/untdid/d20a/trmd/retins_c.htm",
	)
}

var retinsD20aSchema = Schema{Nodes: []SchemaNode{
	{Segment: "BGM", Mandatory: true, MaxRepeat: 1},  // Beginning of message
	{Segment: "DTM", Mandatory: false, MaxRepeat: 9}, // Date/time/period
	{Segment: "FTX", Mandatory: false, MaxRepeat: 9}, // Free text
	{Segment: "CUX", Mandatory: false, MaxRepeat: 9}, // Currencies
	{ // Segment group 1
		Group: []SchemaNode{
			{Segment: "DOC", Mandatory: true, MaxRepeat: 1},  // Document/message details
			{Segment: "DTM", Mandatory: false, MaxRepeat: 9}, // Date/time/period
		},
		Mandatory: false, MaxRepeat: 9,
	},
	{ // Segment group 2
		Group: []SchemaNode{
			{Segment: "RFF", Mandatory: true, MaxRepeat: 1},  // Reference
			{Segment: "DTM", Mandatory: false, MaxRepeat: 9}, // Date/time/period
		},
		Mandatory: false, MaxRepeat: 9,
	},
	{ // Segment group 3
		Group: []SchemaNode{
			{Segment: "NAD", Mandatory: true, MaxRepeat: 1},  // Name and address
			{Segment: "LOC", Mandatory: false, MaxRepeat: 9}, // Place/location identification
			{ // Segment group 4
				Group: []SchemaNode{
					{Segment: "RFF", Mandatory: true, MaxRepeat: 1},  // Reference
					{Segment: "DTM", Mandatory: false, MaxRepeat: 9}, // Date/time/period
				},
				Mandatory: false, MaxRepeat: 9,
			},
			{ // Segment group 5
				Group: []SchemaNode{
					{Segment: "CTA", Mandatory: true, MaxRepeat: 1},  // Contact information
					{Segment: "COM", Mandatory: false, MaxRepeat: 9}, // Communication contact
				},
				Mandatory: false, MaxRepeat: 9,
			},
		},
		Mandatory: false, MaxRepeat: 9,
	},
	{ // Segment group 6
		Group: []SchemaNode{
			{Segment: "CDI", Mandatory: true, MaxRepeat: 1},  // Physical or logical state
			{Segment: "INP", Mandatory: false, MaxRepeat: 9}, // Parties and instruction
			{Segment: "DTM", Mandatory: false, MaxRepeat: 9}, // Date/time/period
			{Segment: "FTX", Mandatory: false, MaxRepeat: 9}, // Free text
			{ // Segment group 7
				Group: []SchemaNode{
					{Segment: "RFF", Mandatory: true, MaxRepeat: 1},  // Reference
					{Segment: "DTM", Mandatory: false, MaxRepeat: 9}, // Date/time/period
				},
				Mandatory: false, MaxRepeat: 9,
			},
		},
		Mandatory: false, MaxRepeat: 99,
	},
	{ // Segment group 8
		Group: []SchemaNode{
			{Segment: "TOD", Mandatory: true, MaxRepeat: 1},  // Terms of delivery or transport
			{Segment: "LOC", Mandatory: false, MaxRepeat: 9}, // Place/location identification
		},
		Mandatory: false, MaxRepeat: 9,
	},
	{ // Segment group 9
		Group: []SchemaNode{
			{Segment: "TDT", Mandatory: true, MaxRepeat: 1},  // Transport information
			{Segment: "LOC", Mandatory: false, MaxRepeat: 9}, // Place/location identification
			{Segment: "DTM", Mandatory: false, MaxRepeat: 9}, // Date/time/period
		},
		Mandatory: false, MaxRepeat: 9,
	},
	{ // Segment group 10
		Group: []SchemaNode{
			{Segment: "EQD", Mandatory: true, MaxRepeat: 1},  // Equipment details
			{Segment: "MEA", Mandatory: false, MaxRepeat: 9}, // Measurements
			{Segment: "FTX", Mandatory: false, MaxRepeat: 9}, // Free text
		},
		Mandatory: false, MaxRepeat: 9,
	},
	{ // Segment group 11
		Group: []SchemaNode{
			{Segment: "LIN", Mandatory: true, MaxRepeat: 1},  // Line item
			{Segment: "PIA", Mandatory: false, MaxRepeat: 9}, // Additional product id
			{Segment: "IMD", Mandatory: false, MaxRepeat: 9}, // Item description
			{Segment: "MEA", Mandatory: false, MaxRepeat: 9}, // Measurements
			{Segment: "DTM", Mandatory: false, MaxRepeat: 9}, // Date/time/period
			{Segment: "QTY", Mandatory: false, MaxRepeat: 9}, // Quantity
			{Segment: "FTX", Mandatory: false, MaxRepeat: 9}, // Free text
			{Segment: "LOC", Mandatory: false, MaxRepeat: 9}, // Place/location identification
			{Segment: "MOA", Mandatory: false, MaxRepeat: 9}, // Monetary amount
			{ // Segment group 12
				Group: []SchemaNode{
					{Segment: "DOC", Mandatory: true, MaxRepeat: 1},  // Document/message details
					{Segment: "DTM", Mandatory: false, MaxRepeat: 9}, // Date/time/period
				},
				Mandatory: false, MaxRepeat: 9,
			},
			{ // Segment group 13
				Group: []SchemaNode{
					{Segment: "RFF", Mandatory: true, MaxRepeat: 1},  // Reference
					{Segment: "DTM", Mandatory: false, MaxRepeat: 9}, // Date/time/period
				},
				Mandatory: false, MaxRepeat: 9,
			},
			{ // Segment group 14
				Group: []SchemaNode{
					{Segment: "DGS", Mandatory: true, MaxRepeat: 1},  // Dangerous goods
					{Segment: "QTY", Mandatory: false, MaxRepeat: 9}, // Quantity
					{Segment: "FTX", Mandatory: false, MaxRepeat: 9}, // Free text
				},
				Mandatory: false, MaxRepeat: 9,
			},
			{ // Segment group 15
				Group: []SchemaNode{
					{Segment: "HAN", Mandatory: true, MaxRepeat: 1},  // Handling instructions
					{Segment: "FTX", Mandatory: false, MaxRepeat: 9}, // Free text
				},
				Mandatory: false, MaxRepeat: 9,
			},
			{ // Segment group 16
				Group: []SchemaNode{
					{Segment: "PAC", Mandatory: true, MaxRepeat: 1}, // Package
					{ // Segment group 17
						Group: []SchemaNode{
							{Segment: "PCI", Mandatory: true, MaxRepeat: 1},  // Package identification
							{Segment: "DTM", Mandatory: false, MaxRepeat: 9}, // Date/time/period
							{Segment: "GIN", Mandatory: false, MaxRepeat: 9}, // Goods identity number
						},
						Mandatory: false, MaxRepeat: 9,
					},
				},
				Mandatory: false, MaxRepeat: 9,
			},
			{ // Segment group 18
				Group: []SchemaNode{
					{Segment: "CDI", Mandatory: true, MaxRepeat: 1},  // Physical or logical state
					{Segment: "INP", Mandatory: false, MaxRepeat: 9}, // Parties and instruction
					{Segment: "DTM", Mandatory: false, MaxRepeat: 9}, // Date/time/period
					{Segment: "FTX", Mandatory: false, MaxRepeat: 9}, // Free text
					{ // Segment group 19
						Group: []SchemaNode{
							{Segment: "RFF", Mandatory: true, MaxRepeat: 1},  // Reference
							{Segment: "DTM", Mandatory: false, MaxRepeat: 9}, // Date/time/period
						},
						Mandatory: false, MaxRepeat: 9,
					},
				},
				Mandatory: false, MaxRepeat: 99,
			},
			{ // Segment group 20
				Group: []SchemaNode{
					{Segment: "TDT", Mandatory: true, MaxRepeat: 1},  // Transport information
					{Segment: "LOC", Mandatory: false, MaxRepeat: 9}, // Place/location identification
					{Segment: "DTM", Mandatory: false, MaxRepeat: 9}, // Date/time/period
				},
				Mandatory: false, MaxRepeat: 9,
			},
			{ // Segment group 21
				Group: []SchemaNode{
					{Segment: "EQD", Mandatory: true, MaxRepeat: 1},  // Equipment details
					{Segment: "MEA", Mandatory: false, MaxRepeat: 9}, // Measurements
					{Segment: "FTX", Mandatory: false, MaxRepeat: 9}, // Free text
				},
				Mandatory: false, MaxRepeat: 9,
			},
		},
		Mandatory: true, MaxRepeat: 9999,
	},
	{Segment: "CNT", Mandatory: false, MaxRepeat: 9}, // Control total
}}
