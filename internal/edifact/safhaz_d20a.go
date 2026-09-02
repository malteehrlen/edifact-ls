package edifact

// SAFHAZ D.20A message specification: the structural (segment/group
// presence, order, cardinality) branching diagram for this message,
// UN/EDIFACT directory release D.20A. Scope is structural only -- not
// element-level content -- per edifact-ls-3uzr's epic non-goal.
// 15 segment groups, max nesting depth 4.
//
// Source: https://service.unece.org/trade/untdid/d20a/trmd/safhaz_c.htm
// section 4.3.1 "Segment table". That URL currently returns HTTP 403 from
// Cloudflare when fetched directly; this data was transcribed from the
// Wayback Machine's archived copy instead:
// http://web.archive.org/web/20231207012537/https://service.unece.org/trade/untdid/d20a/trmd/safhaz_c.htm
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
		MessageID{Type: "SAFHAZ", Version: "D", Release: "20A", Agency: "UN"},
		safhazD20aSchema,
		"https://service.unece.org/trade/untdid/d20a/trmd/safhaz_c.htm",
	)
}

var safhazD20aSchema = Schema{Nodes: []SchemaNode{
	{Segment: "BGM", Mandatory: true, MaxRepeat: 1},   // Beginning of message
	{Segment: "RFF", Mandatory: false, MaxRepeat: 10}, // Reference
	{Segment: "DTM", Mandatory: false, MaxRepeat: 10}, // Date/time/period
	{Segment: "FTX", Mandatory: false, MaxRepeat: 99}, // Free text
	{ // Segment group 1
		Group: []SchemaNode{
			{Segment: "NAD", Mandatory: true, MaxRepeat: 1},   // Name and address
			{Segment: "RFF", Mandatory: false, MaxRepeat: 10}, // Reference
			{ // Segment group 2
				Group: []SchemaNode{
					{Segment: "CTA", Mandatory: true, MaxRepeat: 1},  // Contact information
					{Segment: "COM", Mandatory: false, MaxRepeat: 5}, // Communication contact
				},
				Mandatory: false, MaxRepeat: 10,
			},
		},
		Mandatory: false, MaxRepeat: 10,
	},
	{ // Segment group 3
		Group: []SchemaNode{
			{Segment: "DOC", Mandatory: true, MaxRepeat: 1},    // Document/message details
			{Segment: "IMD", Mandatory: false, MaxRepeat: 999}, // Item description
			{Segment: "PIA", Mandatory: false, MaxRepeat: 10},  // Additional product id
			{Segment: "MEA", Mandatory: false, MaxRepeat: 10},  // Measurements
			{Segment: "RCS", Mandatory: false, MaxRepeat: 10},  // Requirements and conditions
			{Segment: "RFF", Mandatory: false, MaxRepeat: 10},  // Reference
			{Segment: "DTM", Mandatory: false, MaxRepeat: 10},  // Date/time/period
			{Segment: "FTX", Mandatory: false, MaxRepeat: 99},  // Free text
			{ // Segment group 4
				Group: []SchemaNode{
					{Segment: "NAD", Mandatory: true, MaxRepeat: 1},   // Name and address
					{Segment: "RFF", Mandatory: false, MaxRepeat: 10}, // Reference
					{ // Segment group 5
						Group: []SchemaNode{
							{Segment: "CTA", Mandatory: true, MaxRepeat: 1},  // Contact information
							{Segment: "COM", Mandatory: false, MaxRepeat: 5}, // Communication contact
						},
						Mandatory: false, MaxRepeat: 10,
					},
				},
				Mandatory: false, MaxRepeat: 10,
			},
			{ // Segment group 6
				Group: []SchemaNode{
					{Segment: "SFI", Mandatory: true, MaxRepeat: 1},   // Safety information
					{Segment: "PIA", Mandatory: false, MaxRepeat: 10}, // Additional product id
					{Segment: "EQD", Mandatory: false, MaxRepeat: 99}, // Equipment details
					{Segment: "FTX", Mandatory: false, MaxRepeat: 99}, // Free text
					{ // Segment group 7
						Group: []SchemaNode{
							{Segment: "NAD", Mandatory: true, MaxRepeat: 1},   // Name and address
							{Segment: "RFF", Mandatory: false, MaxRepeat: 10}, // Reference
							{ // Segment group 8
								Group: []SchemaNode{
									{Segment: "CTA", Mandatory: true, MaxRepeat: 1},  // Contact information
									{Segment: "COM", Mandatory: false, MaxRepeat: 5}, // Communication contact
								},
								Mandatory: false, MaxRepeat: 10,
							},
						},
						Mandatory: false, MaxRepeat: 10,
					},
					{ // Segment group 9
						Group: []SchemaNode{
							{Segment: "HAN", Mandatory: true, MaxRepeat: 1},   // Handling instructions
							{Segment: "FTX", Mandatory: false, MaxRepeat: 99}, // Free text
						},
						Mandatory: false, MaxRepeat: 10,
					},
					{ // Segment group 10
						Group: []SchemaNode{
							{Segment: "IMD", Mandatory: true, MaxRepeat: 1},   // Item description
							{Segment: "MEA", Mandatory: false, MaxRepeat: 10}, // Measurements
							{Segment: "PCD", Mandatory: false, MaxRepeat: 10}, // Percentage details
							{Segment: "RFF", Mandatory: false, MaxRepeat: 10}, // Reference
							{Segment: "FTX", Mandatory: false, MaxRepeat: 99}, // Free text
						},
						Mandatory: false, MaxRepeat: 999,
					},
					{ // Segment group 11
						Group: []SchemaNode{
							{Segment: "DGS", Mandatory: true, MaxRepeat: 1},   // Dangerous goods
							{Segment: "RFF", Mandatory: false, MaxRepeat: 10}, // Reference
							{Segment: "FTX", Mandatory: false, MaxRepeat: 99}, // Free text
							{ // Segment group 12
								Group: []SchemaNode{
									{Segment: "PAC", Mandatory: true, MaxRepeat: 1},   // Package
									{Segment: "PCI", Mandatory: false, MaxRepeat: 10}, // Package identification
								},
								Mandatory: false, MaxRepeat: 10,
							},
						},
						Mandatory: false, MaxRepeat: 10,
					},
					{ // Segment group 13
						Group: []SchemaNode{
							{Segment: "CCI", Mandatory: true, MaxRepeat: 1},   // Characteristic/class id
							{Segment: "IMD", Mandatory: false, MaxRepeat: 10}, // Item description
							{Segment: "FTX", Mandatory: false, MaxRepeat: 99}, // Free text
							{Segment: "CAV", Mandatory: false, MaxRepeat: 10}, // Characteristic value
							{ // Segment group 14
								Group: []SchemaNode{
									{Segment: "RFF", Mandatory: true, MaxRepeat: 1},   // Reference
									{Segment: "DTM", Mandatory: false, MaxRepeat: 10}, // Date/time/period
								},
								Mandatory: false, MaxRepeat: 10,
							},
							{ // Segment group 15
								Group: []SchemaNode{
									{Segment: "MEA", Mandatory: true, MaxRepeat: 1},   // Measurements
									{Segment: "TEM", Mandatory: false, MaxRepeat: 10}, // Test method
									{Segment: "DTM", Mandatory: false, MaxRepeat: 10}, // Date/time/period
									{Segment: "RFF", Mandatory: false, MaxRepeat: 10}, // Reference
									{Segment: "FTX", Mandatory: false, MaxRepeat: 99}, // Free text
								},
								Mandatory: false, MaxRepeat: 10,
							},
						},
						Mandatory: false, MaxRepeat: 999,
					},
				},
				Mandatory: false, MaxRepeat: 20,
			},
		},
		Mandatory: true, MaxRepeat: 1000,
	},
}}
