package edifact

// UTILMD D.20A message specification: the structural (segment/group
// presence, order, cardinality) branching diagram for this message,
// UN/EDIFACT directory release D.20A. Scope is structural only -- not
// element-level content -- per edifact-ls-3uzr's epic non-goal.
// 13 segment groups, max nesting depth 3.
//
// Source: https://service.unece.org/trade/untdid/d20a/trmd/utilmd_c.htm
// section 4.3.1 "Segment table". That URL currently returns HTTP 403 from
// Cloudflare when fetched directly; this data was transcribed from the
// Wayback Machine's archived copy instead:
// http://web.archive.org/web/20231202080514/https://service.unece.org/trade/untdid/d20a/trmd/utilmd_c.htm
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
		MessageID{Type: "UTILMD", Version: "D", Release: "20A", Agency: "UN"},
		utilmdD20aSchema,
		"https://service.unece.org/trade/untdid/d20a/trmd/utilmd_c.htm",
	)
}

var utilmdD20aSchema = Schema{Nodes: []SchemaNode{
	{Segment: "BGM", Mandatory: true, MaxRepeat: 1},  // Beginning of message
	{Segment: "DTM", Mandatory: true, MaxRepeat: 9},  // Date/time/period
	{Segment: "MKS", Mandatory: false, MaxRepeat: 9}, // Market/sales channel information
	{Segment: "FTX", Mandatory: false, MaxRepeat: 9}, // Free text
	{Segment: "TSR", Mandatory: false, MaxRepeat: 9}, // Transport service requirements
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
			{Segment: "FII", Mandatory: false, MaxRepeat: 1}, // Financial institution information
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
			{Segment: "IDE", Mandatory: true, MaxRepeat: 1},   // Identity
			{Segment: "LIN", Mandatory: false, MaxRepeat: 1},  // Line item
			{Segment: "PIA", Mandatory: false, MaxRepeat: 9},  // Additional product id
			{Segment: "IMD", Mandatory: false, MaxRepeat: 9},  // Item description
			{Segment: "DTM", Mandatory: false, MaxRepeat: 99}, // Date/time/period
			{Segment: "PRC", Mandatory: false, MaxRepeat: 9},  // Process identification
			{Segment: "STS", Mandatory: false, MaxRepeat: 9},  // Status
			{Segment: "TAX", Mandatory: false, MaxRepeat: 9},  // Duty/tax/fee details
			{Segment: "PTY", Mandatory: false, MaxRepeat: 9},  // Priority
			{Segment: "FTX", Mandatory: false, MaxRepeat: 9},  // Free text
			{Segment: "AGR", Mandatory: false, MaxRepeat: 9},  // Agreement identification
			{Segment: "INP", Mandatory: false, MaxRepeat: 9},  // Parties and instruction
			{Segment: "TSR", Mandatory: false, MaxRepeat: 9},  // Transport service requirements
			{ // Segment group 5
				Group: []SchemaNode{
					{Segment: "LOC", Mandatory: true, MaxRepeat: 1},  // Place/location identification
					{Segment: "HYN", Mandatory: false, MaxRepeat: 9}, // Hierarchy information
				},
				Mandatory: false, MaxRepeat: 999999,
			},
			{ // Segment group 6
				Group: []SchemaNode{
					{Segment: "RFF", Mandatory: true, MaxRepeat: 1},  // Reference
					{Segment: "DTM", Mandatory: false, MaxRepeat: 9}, // Date/time/period
				},
				Mandatory: false, MaxRepeat: 99,
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
					{Segment: "RFF", Mandatory: false, MaxRepeat: 9}, // Reference
					{Segment: "PIA", Mandatory: false, MaxRepeat: 9}, // Additional product id
					{ // Segment group 9
						Group: []SchemaNode{
							{Segment: "QTY", Mandatory: true, MaxRepeat: 1},  // Quantity
							{Segment: "DTM", Mandatory: false, MaxRepeat: 9}, // Date/time/period
							{Segment: "STS", Mandatory: false, MaxRepeat: 9}, // Status
							{Segment: "LIN", Mandatory: false, MaxRepeat: 9}, // Line item
						},
						Mandatory: false, MaxRepeat: 99,
					},
					{ // Segment group 10
						Group: []SchemaNode{
							{Segment: "CCI", Mandatory: true, MaxRepeat: 1},   // Characteristic/class id
							{Segment: "CAV", Mandatory: false, MaxRepeat: 99}, // Characteristic value
						},
						Mandatory: false, MaxRepeat: 99,
					},
				},
				Mandatory: false, MaxRepeat: 99999,
			},
			{ // Segment group 11
				Group: []SchemaNode{
					{Segment: "MOA", Mandatory: true, MaxRepeat: 1},  // Monetary amount
					{Segment: "RFF", Mandatory: false, MaxRepeat: 9}, // Reference
					{Segment: "DTM", Mandatory: false, MaxRepeat: 9}, // Date/time/period
				},
				Mandatory: false, MaxRepeat: 99,
			},
			{ // Segment group 12
				Group: []SchemaNode{
					{Segment: "NAD", Mandatory: true, MaxRepeat: 1},  // Name and address
					{Segment: "RFF", Mandatory: false, MaxRepeat: 9}, // Reference
					{Segment: "DTM", Mandatory: false, MaxRepeat: 9}, // Date/time/period
					{Segment: "FII", Mandatory: false, MaxRepeat: 1}, // Financial institution information
					{Segment: "LAN", Mandatory: false, MaxRepeat: 9}, // Language
					{ // Segment group 13
						Group: []SchemaNode{
							{Segment: "CTA", Mandatory: true, MaxRepeat: 1},  // Contact information
							{Segment: "COM", Mandatory: false, MaxRepeat: 9}, // Communication contact
						},
						Mandatory: false, MaxRepeat: 9,
					},
				},
				Mandatory: false, MaxRepeat: 99,
			},
		},
		Mandatory: false, MaxRepeat: 99999,
	},
	{Segment: "CNT", Mandatory: false, MaxRepeat: 9}, // Control total
}}
