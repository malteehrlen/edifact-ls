package edifact

// UTILTS D.20A message specification: the structural (segment/group
// presence, order, cardinality) branching diagram for this message,
// UN/EDIFACT directory release D.20A. Scope is structural only -- not
// element-level content -- per edifact-ls-3uzr's epic non-goal.
// 13 segment groups, max nesting depth 4.
//
// Source: https://service.unece.org/trade/untdid/d20a/trmd/utilts_c.htm
// section 4.3.1 "Segment table". That URL currently returns HTTP 403 from
// Cloudflare when fetched directly; this data was transcribed from the
// Wayback Machine's archived copy instead:
// http://web.archive.org/web/20231207001101/https://service.unece.org/trade/untdid/d20a/trmd/utilts_c.htm
// -- re-check the direct URL first if cross-checking later, in case the
// block has lifted.
//
// Transcription note: generated in bulk (edifact-ls-13gu) the same way
// as every other message-type schema -- a script parsed the source's
// exact rail-art column positions mechanically and verified the result
// balances (all 13 groups open and close consistently)
// before being transcribed here, rather than reading the ASCII nesting
// by eye. Not individually spot-checked against the raw source the way
// the first several message types in this project were -- see
// edifact-ls-13gu for which representative samples were spot-checked
// for this batch.

func init() {
	RegisterSchema(
		MessageID{Type: "UTILTS", Version: "D", Release: "20A", Agency: "UN"},
		utiltsD20aSchema,
		"https://service.unece.org/trade/untdid/d20a/trmd/utilts_c.htm",
	)
}

var utiltsD20aSchema = Schema{Nodes: []SchemaNode{
	{Segment: "BGM", Mandatory: true, MaxRepeat: 1},  // Beginning of message
	{Segment: "DTM", Mandatory: true, MaxRepeat: 9},  // Date/time/period
	{Segment: "MKS", Mandatory: false, MaxRepeat: 9}, // Market/sales channel information
	{Segment: "PRC", Mandatory: false, MaxRepeat: 9}, // Process identification
	{ // Segment group 1
		Group: []SchemaNode{
			{Segment: "RFF", Mandatory: true, MaxRepeat: 1},  // Reference
			{Segment: "DTM", Mandatory: false, MaxRepeat: 9}, // Date/time/period
		},
		Mandatory: false, MaxRepeat: 9,
	},
	{ // Segment group 2
		Group: []SchemaNode{
			{Segment: "NAD", Mandatory: true, MaxRepeat: 1},  // Name and address
			{Segment: "RFF", Mandatory: false, MaxRepeat: 1}, // Reference
			{Segment: "ATT", Mandatory: false, MaxRepeat: 9}, // Attribute
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
			{Segment: "CUX", Mandatory: true, MaxRepeat: 1},  // Currencies
			{Segment: "DTM", Mandatory: false, MaxRepeat: 9}, // Date/time/period
			{Segment: "STS", Mandatory: false, MaxRepeat: 9}, // Status
		},
		Mandatory: false, MaxRepeat: 99,
	},
	{ // Segment group 5
		Group: []SchemaNode{
			{Segment: "IDE", Mandatory: true, MaxRepeat: 1},  // Identity
			{Segment: "LOC", Mandatory: false, MaxRepeat: 9}, // Place/location identification
			{Segment: "NAD", Mandatory: false, MaxRepeat: 9}, // Name and address
			{Segment: "ALI", Mandatory: false, MaxRepeat: 9}, // Additional information
			{Segment: "LIN", Mandatory: false, MaxRepeat: 9}, // Line item
			{Segment: "PIA", Mandatory: false, MaxRepeat: 9}, // Additional product id
			{Segment: "IMD", Mandatory: false, MaxRepeat: 9}, // Item description
			{Segment: "DTM", Mandatory: false, MaxRepeat: 9}, // Date/time/period
			{Segment: "PRC", Mandatory: false, MaxRepeat: 9}, // Process identification
			{Segment: "STS", Mandatory: false, MaxRepeat: 9}, // Status
			{Segment: "AGR", Mandatory: false, MaxRepeat: 9}, // Agreement identification
			{Segment: "MEA", Mandatory: false, MaxRepeat: 9}, // Measurements
			{Segment: "FTX", Mandatory: false, MaxRepeat: 9}, // Free text
			{ // Segment group 6
				Group: []SchemaNode{
					{Segment: "RFF", Mandatory: true, MaxRepeat: 1},  // Reference
					{Segment: "DTM", Mandatory: false, MaxRepeat: 9}, // Date/time/period
				},
				Mandatory: false, MaxRepeat: 99999,
			},
			{ // Segment group 7
				Group: []SchemaNode{
					{Segment: "CCI", Mandatory: true, MaxRepeat: 1},   // Characteristic/class id
					{Segment: "CAV", Mandatory: false, MaxRepeat: 99}, // Characteristic value
				},
				Mandatory: false, MaxRepeat: 99,
			},
			{ // Segment group 8
				Group: []SchemaNode{
					{Segment: "SEQ", Mandatory: true, MaxRepeat: 1},  // Sequence details
					{Segment: "DTM", Mandatory: false, MaxRepeat: 9}, // Date/time/period
					{Segment: "RFF", Mandatory: false, MaxRepeat: 9}, // Reference
					{Segment: "MOA", Mandatory: false, MaxRepeat: 9}, // Monetary amount
					{Segment: "PCD", Mandatory: false, MaxRepeat: 9}, // Percentage details
					{Segment: "GPO", Mandatory: false, MaxRepeat: 9}, // Geographical position
					{ // Segment group 9
						Group: []SchemaNode{
							{Segment: "CCI", Mandatory: true, MaxRepeat: 1},   // Characteristic/class id
							{Segment: "CAV", Mandatory: false, MaxRepeat: 99}, // Characteristic value
						},
						Mandatory: false, MaxRepeat: 99,
					},
					{ // Segment group 10
						Group: []SchemaNode{
							{Segment: "PRI", Mandatory: true, MaxRepeat: 1},  // Price details
							{Segment: "CUX", Mandatory: false, MaxRepeat: 9}, // Currencies
						},
						Mandatory: false, MaxRepeat: 9,
					},
					{ // Segment group 11
						Group: []SchemaNode{
							{Segment: "QTY", Mandatory: true, MaxRepeat: 1},  // Quantity
							{Segment: "DTM", Mandatory: false, MaxRepeat: 9}, // Date/time/period
							{Segment: "STS", Mandatory: false, MaxRepeat: 9}, // Status
							{ // Segment group 12
								Group: []SchemaNode{
									{Segment: "CCI", Mandatory: true, MaxRepeat: 1},   // Characteristic/class id
									{Segment: "CAV", Mandatory: false, MaxRepeat: 99}, // Characteristic value
								},
								Mandatory: false, MaxRepeat: 99,
							},
							{ // Segment group 13
								Group: []SchemaNode{
									{Segment: "PRI", Mandatory: true, MaxRepeat: 1},  // Price details
									{Segment: "CUX", Mandatory: false, MaxRepeat: 9}, // Currencies
								},
								Mandatory: false, MaxRepeat: 9,
							},
						},
						Mandatory: false, MaxRepeat: 99999,
					},
				},
				Mandatory: false, MaxRepeat: 99999,
			},
		},
		Mandatory: false, MaxRepeat: 99999,
	},
	{Segment: "CNT", Mandatory: false, MaxRepeat: 9}, // Control total
}}
