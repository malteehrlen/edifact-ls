package edifact

// CUSRES D.20A message specification: the structural (segment/group
// presence, order, cardinality) branching diagram for this message,
// UN/EDIFACT directory release D.20A. Scope is structural only -- not
// element-level content -- per edifact-ls-3uzr's epic non-goal.
// 15 segment groups, max nesting depth 3.
//
// Source: https://service.unece.org/trade/untdid/d20a/trmd/cusres_c.htm
// section 4.3.1 "Segment table". That URL currently returns HTTP 403 from
// Cloudflare when fetched directly; this data was transcribed from the
// Wayback Machine's archived copy instead:
// http://web.archive.org/web/20231202080256/https://service.unece.org/trade/untdid/d20a/trmd/cusres_c.htm
// -- re-check the direct URL first if cross-checking later, in case the
// block has lifted.
//
// Transcription note: generated in bulk (edifact-ls-13gu) the same way
// as every other message-type schema -- a script parsed the source's
// exact rail-art column positions mechanically and verified the result
// balances (all 15 groups open and close consistently)
// before being transcribed here, rather than reading the ASCII nesting
// by eye. Not individually spot-checked against the raw source the way
// the first several message types in this project were -- see
// edifact-ls-13gu for which representative samples were spot-checked
// for this batch.

func init() {
	RegisterSchema(
		MessageID{Type: "CUSRES", Version: "D", Release: "20A", Agency: "UN"},
		cusresD20aSchema,
		"https://service.unece.org/trade/untdid/d20a/trmd/cusres_c.htm",
	)
}

var cusresD20aSchema = Schema{Nodes: []SchemaNode{
	{Segment: "BGM", Mandatory: true, MaxRepeat: 1},    // Beginning of message
	{Segment: "DTM", Mandatory: false, MaxRepeat: 9},   // Date/time/period
	{Segment: "FTX", Mandatory: false, MaxRepeat: 99},  // Free text
	{Segment: "TDT", Mandatory: false, MaxRepeat: 9},   // Transport information
	{Segment: "LOC", Mandatory: false, MaxRepeat: 99},  // Place/location identification
	{Segment: "GEI", Mandatory: false, MaxRepeat: 10},  // Processing information
	{Segment: "EQD", Mandatory: false, MaxRepeat: 999}, // Equipment details
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
		Mandatory: false, MaxRepeat: 99,
	},
	{ // Segment group 3
		Group: []SchemaNode{
			{Segment: "RFF", Mandatory: true, MaxRepeat: 1},  // Reference
			{Segment: "DTM", Mandatory: false, MaxRepeat: 9}, // Date/time/period
			{Segment: "LOC", Mandatory: false, MaxRepeat: 9}, // Place/location identification
		},
		Mandatory: false, MaxRepeat: 999,
	},
	{ // Segment group 4
		Group: []SchemaNode{
			{Segment: "ERP", Mandatory: true, MaxRepeat: 1},   // Error point details
			{Segment: "RFF", Mandatory: false, MaxRepeat: 9},  // Reference
			{Segment: "ERC", Mandatory: false, MaxRepeat: 99}, // Application error information
			{Segment: "FTX", Mandatory: false, MaxRepeat: 99}, // Free text
		},
		Mandatory: false, MaxRepeat: 99999,
	},
	{ // Segment group 5
		Group: []SchemaNode{
			{Segment: "TAX", Mandatory: true, MaxRepeat: 1},   // Duty/tax/fee details
			{Segment: "MOA", Mandatory: false, MaxRepeat: 99}, // Monetary amount
			{Segment: "GEI", Mandatory: false, MaxRepeat: 99}, // Processing information
		},
		Mandatory: false, MaxRepeat: 99,
	},
	{ // Segment group 6
		Group: []SchemaNode{
			{Segment: "DOC", Mandatory: true, MaxRepeat: 1},    // Document/message details
			{Segment: "PAC", Mandatory: false, MaxRepeat: 9},   // Package
			{Segment: "RFF", Mandatory: false, MaxRepeat: 9},   // Reference
			{Segment: "PCI", Mandatory: false, MaxRepeat: 9},   // Package identification
			{Segment: "FTX", Mandatory: false, MaxRepeat: 99},  // Free text
			{Segment: "TDT", Mandatory: false, MaxRepeat: 9},   // Transport information
			{Segment: "LOC", Mandatory: false, MaxRepeat: 9},   // Place/location identification
			{Segment: "DTM", Mandatory: false, MaxRepeat: 9},   // Date/time/period
			{Segment: "GEI", Mandatory: false, MaxRepeat: 99},  // Processing information
			{Segment: "MEA", Mandatory: false, MaxRepeat: 99},  // Measurements
			{Segment: "EQD", Mandatory: false, MaxRepeat: 999}, // Equipment details
			{ // Segment group 7
				Group: []SchemaNode{
					{Segment: "NAD", Mandatory: true, MaxRepeat: 1}, // Name and address
					{ // Segment group 8
						Group: []SchemaNode{
							{Segment: "CTA", Mandatory: true, MaxRepeat: 1},  // Contact information
							{Segment: "COM", Mandatory: false, MaxRepeat: 9}, // Communication contact
						},
						Mandatory: false, MaxRepeat: 9,
					},
				},
				Mandatory: false, MaxRepeat: 9,
			},
			{ // Segment group 9
				Group: []SchemaNode{
					{Segment: "MOA", Mandatory: true, MaxRepeat: 1}, // Monetary amount
					{ // Segment group 10
						Group: []SchemaNode{
							{Segment: "CUX", Mandatory: true, MaxRepeat: 1},  // Currencies
							{Segment: "DTM", Mandatory: false, MaxRepeat: 1}, // Date/time/period
						},
						Mandatory: false, MaxRepeat: 9,
					},
				},
				Mandatory: false, MaxRepeat: 99,
			},
			{ // Segment group 11
				Group: []SchemaNode{
					{Segment: "TAX", Mandatory: true, MaxRepeat: 1},   // Duty/tax/fee details
					{Segment: "MOA", Mandatory: false, MaxRepeat: 99}, // Monetary amount
					{Segment: "GEI", Mandatory: false, MaxRepeat: 99}, // Processing information
				},
				Mandatory: false, MaxRepeat: 99,
			},
			{ // Segment group 12
				Group: []SchemaNode{
					{Segment: "CST", Mandatory: true, MaxRepeat: 1},  // Customs status of goods
					{Segment: "FTX", Mandatory: false, MaxRepeat: 9}, // Free text
					{ // Segment group 13
						Group: []SchemaNode{
							{Segment: "TAX", Mandatory: true, MaxRepeat: 1},   // Duty/tax/fee details
							{Segment: "MOA", Mandatory: false, MaxRepeat: 99}, // Monetary amount
							{Segment: "GEI", Mandatory: false, MaxRepeat: 99}, // Processing information
							{Segment: "MEA", Mandatory: false, MaxRepeat: 9},  // Measurements
							{Segment: "RFF", Mandatory: false, MaxRepeat: 9},  // Reference
						},
						Mandatory: false, MaxRepeat: 999999,
					},
				},
				Mandatory: false, MaxRepeat: 9999,
			},
			{ // Segment group 14
				Group: []SchemaNode{
					{Segment: "ERP", Mandatory: true, MaxRepeat: 1},     // Error point details
					{Segment: "RFF", Mandatory: false, MaxRepeat: 9},    // Reference
					{Segment: "ERC", Mandatory: false, MaxRepeat: 9999}, // Application error information
					{Segment: "FTX", Mandatory: false, MaxRepeat: 9},    // Free text
				},
				Mandatory: false, MaxRepeat: 9999,
			},
		},
		Mandatory: false, MaxRepeat: 99999,
	},
	{Segment: "CNT", Mandatory: false, MaxRepeat: 9}, // Control total
	{ // Segment group 15
		Group: []SchemaNode{
			{Segment: "AUT", Mandatory: true, MaxRepeat: 1},  // Authentication result
			{Segment: "DTM", Mandatory: false, MaxRepeat: 9}, // Date/time/period
		},
		Mandatory: false, MaxRepeat: 9,
	},
}}
