package edifact

// DESTIM D.20A message specification: the structural (segment/group
// presence, order, cardinality) branching diagram for this message,
// UN/EDIFACT directory release D.20A. Scope is structural only -- not
// element-level content -- per edifact-ls-3uzr's epic non-goal.
// 12 segment groups, max nesting depth 3.
//
// Source: https://service.unece.org/trade/untdid/d20a/trmd/destim_c.htm
// section 4.3.1 "Segment table". That URL currently returns HTTP 403 from
// Cloudflare when fetched directly; this data was transcribed from the
// Wayback Machine's archived copy instead:
// http://web.archive.org/web/20231208123706/https://service.unece.org/trade/untdid/d20a/trmd/destim_c.htm
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
		MessageID{Type: "DESTIM", Version: "D", Release: "20A", Agency: "UN"},
		destimD20aSchema,
		"https://service.unece.org/trade/untdid/d20a/trmd/destim_c.htm",
	)
}

var destimD20aSchema = Schema{Nodes: []SchemaNode{
	{Segment: "BGM", Mandatory: true, MaxRepeat: 1},  // Beginning of message
	{Segment: "DTM", Mandatory: true, MaxRepeat: 9},  // Date/time/period
	{Segment: "GEI", Mandatory: false, MaxRepeat: 1}, // Processing information
	{Segment: "CUX", Mandatory: false, MaxRepeat: 9}, // Currencies
	{Segment: "RFF", Mandatory: false, MaxRepeat: 9}, // Reference
	{Segment: "FTX", Mandatory: false, MaxRepeat: 9}, // Free text
	{ // Segment group 1
		Group: []SchemaNode{
			{Segment: "EQD", Mandatory: true, MaxRepeat: 1},  // Equipment details
			{Segment: "DIM", Mandatory: false, MaxRepeat: 1}, // Dimensions
			{Segment: "IMD", Mandatory: false, MaxRepeat: 9}, // Item description
		},
		Mandatory: true, MaxRepeat: 1,
	},
	{ // Segment group 2
		Group: []SchemaNode{
			{Segment: "NAD", Mandatory: true, MaxRepeat: 1},  // Name and address
			{Segment: "MOA", Mandatory: false, MaxRepeat: 9}, // Monetary amount
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
			{Segment: "TDT", Mandatory: true, MaxRepeat: 1},  // Transport information
			{Segment: "DTM", Mandatory: false, MaxRepeat: 9}, // Date/time/period
			{ // Segment group 5
				Group: []SchemaNode{
					{Segment: "LOC", Mandatory: true, MaxRepeat: 1},  // Place/location identification
					{Segment: "DTM", Mandatory: false, MaxRepeat: 9}, // Date/time/period
				},
				Mandatory: false, MaxRepeat: 9,
			},
		},
		Mandatory: false, MaxRepeat: 1,
	},
	{Segment: "UNS", Mandatory: true, MaxRepeat: 1}, // Section control
	{ // Segment group 6
		Group: []SchemaNode{
			{Segment: "DTM", Mandatory: true, MaxRepeat: 1},  // Date/time/period
			{Segment: "LOC", Mandatory: false, MaxRepeat: 1}, // Place/location identification
			{Segment: "NAD", Mandatory: false, MaxRepeat: 9}, // Name and address
		},
		Mandatory: false, MaxRepeat: 9,
	},
	{ // Segment group 7
		Group: []SchemaNode{
			{Segment: "RFF", Mandatory: true, MaxRepeat: 1},  // Reference
			{Segment: "DTM", Mandatory: false, MaxRepeat: 1}, // Date/time/period
			{Segment: "NAD", Mandatory: false, MaxRepeat: 9}, // Name and address
			{Segment: "MOA", Mandatory: false, MaxRepeat: 1}, // Monetary amount
		},
		Mandatory: false, MaxRepeat: 9,
	},
	{ // Segment group 8
		Group: []SchemaNode{
			{Segment: "LIN", Mandatory: true, MaxRepeat: 1},  // Line item
			{Segment: "DIM", Mandatory: false, MaxRepeat: 1}, // Dimensions
			{Segment: "QTY", Mandatory: false, MaxRepeat: 1}, // Quantity
			{Segment: "FTX", Mandatory: false, MaxRepeat: 9}, // Free text
			{ // Segment group 9
				Group: []SchemaNode{
					{Segment: "DAM", Mandatory: true, MaxRepeat: 1},  // Damage
					{Segment: "COD", Mandatory: false, MaxRepeat: 1}, // Component details
				},
				Mandatory: false, MaxRepeat: 1,
			},
			{ // Segment group 10
				Group: []SchemaNode{
					{Segment: "RTE", Mandatory: true, MaxRepeat: 1}, // Rate details
					{Segment: "QTY", Mandatory: true, MaxRepeat: 1}, // Quantity
				},
				Mandatory: false, MaxRepeat: 9,
			},
			{ // Segment group 11
				Group: []SchemaNode{
					{Segment: "NAD", Mandatory: true, MaxRepeat: 1}, // Name and address
					{ // Segment group 12
						Group: []SchemaNode{
							{Segment: "MOA", Mandatory: true, MaxRepeat: 1},  // Monetary amount
							{Segment: "TAX", Mandatory: false, MaxRepeat: 1}, // Duty/tax/fee details
							{Segment: "MEA", Mandatory: false, MaxRepeat: 1}, // Measurements
						},
						Mandatory: true, MaxRepeat: 9,
					},
				},
				Mandatory: false, MaxRepeat: 9,
			},
		},
		Mandatory: false, MaxRepeat: 999,
	},
	{Segment: "CNT", Mandatory: false, MaxRepeat: 9}, // Control total
}}
