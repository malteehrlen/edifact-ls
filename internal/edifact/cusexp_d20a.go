package edifact

// CUSEXP D.20A message specification: the structural (segment/group
// presence, order, cardinality) branching diagram for this message,
// UN/EDIFACT directory release D.20A. Scope is structural only -- not
// element-level content -- per edifact-ls-3uzr's epic non-goal.
// 17 segment groups, max nesting depth 4.
//
// Source: https://service.unece.org/trade/untdid/d20a/trmd/cusexp_c.htm
// section 4.3.1 "Segment table". That URL currently returns HTTP 403 from
// Cloudflare when fetched directly; this data was transcribed from the
// Wayback Machine's archived copy instead:
// http://web.archive.org/web/20231202083549/https://service.unece.org/trade/untdid/d20a/trmd/cusexp_c.htm
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
		MessageID{Type: "CUSEXP", Version: "D", Release: "20A", Agency: "UN"},
		cusexpD20aSchema,
		"https://service.unece.org/trade/untdid/d20a/trmd/cusexp_c.htm",
	)
}

var cusexpD20aSchema = Schema{Nodes: []SchemaNode{
	{Segment: "BGM", Mandatory: true, MaxRepeat: 1},  // Beginning of message
	{Segment: "DTM", Mandatory: false, MaxRepeat: 5}, // Date/time/period
	{Segment: "LOC", Mandatory: false, MaxRepeat: 5}, // Place/location identification
	{Segment: "CNT", Mandatory: false, MaxRepeat: 9}, // Control total
	{ // Segment group 1
		Group: []SchemaNode{
			{Segment: "NAD", Mandatory: true, MaxRepeat: 1}, // Name and address
			{ // Segment group 2
				Group: []SchemaNode{
					{Segment: "CTA", Mandatory: true, MaxRepeat: 1},  // Contact information
					{Segment: "COM", Mandatory: false, MaxRepeat: 5}, // Communication contact
				},
				Mandatory: false, MaxRepeat: 5,
			},
		},
		Mandatory: true, MaxRepeat: 1,
	},
	{ // Segment group 3
		Group: []SchemaNode{
			{Segment: "TDT", Mandatory: true, MaxRepeat: 1},  // Transport information
			{Segment: "DTM", Mandatory: false, MaxRepeat: 1}, // Date/time/period
			{ // Segment group 4
				Group: []SchemaNode{
					{Segment: "LOC", Mandatory: true, MaxRepeat: 1},  // Place/location identification
					{Segment: "DTM", Mandatory: false, MaxRepeat: 9}, // Date/time/period
				},
				Mandatory: false, MaxRepeat: 9,
			},
		},
		Mandatory: true, MaxRepeat: 1,
	},
	{ // Segment group 5
		Group: []SchemaNode{
			{Segment: "EQD", Mandatory: true, MaxRepeat: 1},  // Equipment details
			{Segment: "SEL", Mandatory: false, MaxRepeat: 9}, // Seal number
		},
		Mandatory: false, MaxRepeat: 99,
	},
	{ // Segment group 6
		Group: []SchemaNode{
			{Segment: "RFF", Mandatory: true, MaxRepeat: 1},  // Reference
			{Segment: "NAD", Mandatory: false, MaxRepeat: 2}, // Name and address
			{Segment: "CNT", Mandatory: false, MaxRepeat: 1}, // Control total
			{ // Segment group 7
				Group: []SchemaNode{
					{Segment: "CNI", Mandatory: true, MaxRepeat: 1},  // Consignment information
					{Segment: "SGP", Mandatory: false, MaxRepeat: 9}, // Split goods placement
					{Segment: "CNT", Mandatory: false, MaxRepeat: 9}, // Control total
					{Segment: "MEA", Mandatory: false, MaxRepeat: 1}, // Measurements
					{Segment: "LOC", Mandatory: false, MaxRepeat: 2}, // Place/location identification
					{Segment: "NAD", Mandatory: false, MaxRepeat: 5}, // Name and address
					{ // Segment group 8
						Group: []SchemaNode{
							{Segment: "GDS", Mandatory: true, MaxRepeat: 1},  // Nature of cargo
							{Segment: "FTX", Mandatory: false, MaxRepeat: 1}, // Free text
						},
						Mandatory: false, MaxRepeat: 1,
					},
					{ // Segment group 9
						Group: []SchemaNode{
							{Segment: "PAC", Mandatory: true, MaxRepeat: 1},  // Package
							{Segment: "PCI", Mandatory: false, MaxRepeat: 1}, // Package identification
						},
						Mandatory: false, MaxRepeat: 999,
					},
					{ // Segment group 10
						Group: []SchemaNode{
							{Segment: "TOD", Mandatory: true, MaxRepeat: 1},  // Terms of delivery or transport
							{Segment: "LOC", Mandatory: false, MaxRepeat: 1}, // Place/location identification
							{Segment: "FTX", Mandatory: false, MaxRepeat: 1}, // Free text
						},
						Mandatory: false, MaxRepeat: 1,
					},
					{ // Segment group 11
						Group: []SchemaNode{
							{Segment: "MOA", Mandatory: true, MaxRepeat: 1}, // Monetary amount
							{ // Segment group 12
								Group: []SchemaNode{
									{Segment: "CUX", Mandatory: true, MaxRepeat: 1},  // Currencies
									{Segment: "DTM", Mandatory: false, MaxRepeat: 1}, // Date/time/period
								},
								Mandatory: false, MaxRepeat: 1,
							},
						},
						Mandatory: false, MaxRepeat: 10,
					},
					{ // Segment group 13
						Group: []SchemaNode{
							{Segment: "TAX", Mandatory: true, MaxRepeat: 1},  // Duty/tax/fee details
							{Segment: "MOA", Mandatory: false, MaxRepeat: 1}, // Monetary amount
							{Segment: "GEI", Mandatory: false, MaxRepeat: 1}, // Processing information
						},
						Mandatory: false, MaxRepeat: 9,
					},
					{ // Segment group 14
						Group: []SchemaNode{
							{Segment: "DOC", Mandatory: true, MaxRepeat: 1},  // Document/message details
							{Segment: "DTM", Mandatory: false, MaxRepeat: 1}, // Date/time/period
							{Segment: "LOC", Mandatory: false, MaxRepeat: 1}, // Place/location identification
						},
						Mandatory: false, MaxRepeat: 9,
					},
					{ // Segment group 15
						Group: []SchemaNode{
							{Segment: "CST", Mandatory: true, MaxRepeat: 1},  // Customs status of goods
							{Segment: "FTX", Mandatory: false, MaxRepeat: 1}, // Free text
							{Segment: "LOC", Mandatory: true, MaxRepeat: 1},  // Place/location identification
							{Segment: "MEA", Mandatory: false, MaxRepeat: 9}, // Measurements
							{ // Segment group 16
								Group: []SchemaNode{
									{Segment: "TAX", Mandatory: true, MaxRepeat: 1},  // Duty/tax/fee details
									{Segment: "MOA", Mandatory: false, MaxRepeat: 1}, // Monetary amount
									{Segment: "GEI", Mandatory: false, MaxRepeat: 1}, // Processing information
								},
								Mandatory: false, MaxRepeat: 9,
							},
						},
						Mandatory: false, MaxRepeat: 99,
					},
				},
				Mandatory: true, MaxRepeat: 9999,
			},
		},
		Mandatory: false, MaxRepeat: 999,
	},
	{ // Segment group 17
		Group: []SchemaNode{
			{Segment: "AUT", Mandatory: true, MaxRepeat: 1},  // Authentication result
			{Segment: "DTM", Mandatory: false, MaxRepeat: 1}, // Date/time/period
		},
		Mandatory: false, MaxRepeat: 1,
	},
}}
