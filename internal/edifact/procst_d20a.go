package edifact

// PROCST D.20A message specification: the structural (segment/group
// presence, order, cardinality) branching diagram for this message,
// UN/EDIFACT directory release D.20A. Scope is structural only -- not
// element-level content -- per edifact-ls-3uzr's epic non-goal.
// 16 segment groups, max nesting depth 4.
//
// Source: https://service.unece.org/trade/untdid/d20a/trmd/procst_c.htm
// section 4.3.1 "Segment table". That URL currently returns HTTP 403 from
// Cloudflare when fetched directly; this data was transcribed from the
// Wayback Machine's archived copy instead:
// http://web.archive.org/web/20231202081422/https://service.unece.org/trade/untdid/d20a/trmd/procst_c.htm
// -- re-check the direct URL first if cross-checking later, in case the
// block has lifted.
//
// Transcription note: generated in bulk (edifact-ls-13gu) the same way
// as every other message-type schema -- a script parsed the source's
// exact rail-art column positions mechanically and verified the result
// balances (all 16 groups open and close consistently)
// before being transcribed here, rather than reading the ASCII nesting
// by eye. Not individually spot-checked against the raw source the way
// the first several message types in this project were -- see
// edifact-ls-13gu for which representative samples were spot-checked
// for this batch.

func init() {
	RegisterSchema(
		MessageID{Type: "PROCST", Version: "D", Release: "20A", Agency: "UN"},
		procstD20aSchema,
		"https://service.unece.org/trade/untdid/d20a/trmd/procst_c.htm",
	)
}

var procstD20aSchema = Schema{Nodes: []SchemaNode{
	{Segment: "BGM", Mandatory: true, MaxRepeat: 1},   // Beginning of message
	{Segment: "PRI", Mandatory: false, MaxRepeat: 1},  // Price details
	{Segment: "CUX", Mandatory: false, MaxRepeat: 1},  // Currencies
	{Segment: "ATT", Mandatory: false, MaxRepeat: 99}, // Attribute
	{Segment: "MOA", Mandatory: false, MaxRepeat: 99}, // Monetary amount
	{Segment: "PCD", Mandatory: false, MaxRepeat: 99}, // Percentage details
	{Segment: "DTM", Mandatory: false, MaxRepeat: 99}, // Date/time/period
	{ // Segment group 1
		Group: []SchemaNode{
			{Segment: "RFF", Mandatory: true, MaxRepeat: 1},  // Reference
			{Segment: "IMD", Mandatory: false, MaxRepeat: 9}, // Item description
		},
		Mandatory: false, MaxRepeat: 99,
	},
	{ // Segment group 2
		Group: []SchemaNode{
			{Segment: "QTY", Mandatory: true, MaxRepeat: 1},  // Quantity
			{Segment: "IMD", Mandatory: false, MaxRepeat: 1}, // Item description
			{Segment: "CCI", Mandatory: false, MaxRepeat: 1}, // Characteristic/class id
			{Segment: "ATT", Mandatory: false, MaxRepeat: 9}, // Attribute
		},
		Mandatory: false, MaxRepeat: 99,
	},
	{ // Segment group 3
		Group: []SchemaNode{
			{Segment: "PNA", Mandatory: true, MaxRepeat: 1},  // Party identification
			{Segment: "ADR", Mandatory: false, MaxRepeat: 1}, // Address
			{ // Segment group 4
				Group: []SchemaNode{
					{Segment: "CTA", Mandatory: true, MaxRepeat: 1},  // Contact information
					{Segment: "COM", Mandatory: false, MaxRepeat: 9}, // Communication contact
					{Segment: "DTM", Mandatory: false, MaxRepeat: 9}, // Date/time/period
				},
				Mandatory: false, MaxRepeat: 9,
			},
		},
		Mandatory: false, MaxRepeat: 99,
	},
	{ // Segment group 5
		Group: []SchemaNode{
			{Segment: "EFI", Mandatory: true, MaxRepeat: 1},  // External file link identification
			{Segment: "CED", Mandatory: false, MaxRepeat: 1}, // Computer environment details
			{Segment: "DTM", Mandatory: false, MaxRepeat: 1}, // Date/time/period
		},
		Mandatory: false, MaxRepeat: 99,
	},
	{ // Segment group 6
		Group: []SchemaNode{
			{Segment: "LIN", Mandatory: true, MaxRepeat: 1},  // Line item
			{Segment: "ATT", Mandatory: false, MaxRepeat: 9}, // Attribute
			{Segment: "DTM", Mandatory: false, MaxRepeat: 9}, // Date/time/period
			{ // Segment group 7
				Group: []SchemaNode{
					{Segment: "RFF", Mandatory: true, MaxRepeat: 1},  // Reference
					{Segment: "QTY", Mandatory: false, MaxRepeat: 1}, // Quantity
					{Segment: "DTM", Mandatory: false, MaxRepeat: 1}, // Date/time/period
					{Segment: "IMD", Mandatory: false, MaxRepeat: 1}, // Item description
				},
				Mandatory: false, MaxRepeat: 9,
			},
			{ // Segment group 8
				Group: []SchemaNode{
					{Segment: "IDE", Mandatory: true, MaxRepeat: 1},   // Identity
					{Segment: "REL", Mandatory: false, MaxRepeat: 1},  // Relationship
					{Segment: "IMD", Mandatory: false, MaxRepeat: 1},  // Item description
					{Segment: "RCS", Mandatory: false, MaxRepeat: 9},  // Requirements and conditions
					{Segment: "CCI", Mandatory: false, MaxRepeat: 9},  // Characteristic/class id
					{Segment: "RFF", Mandatory: false, MaxRepeat: 9},  // Reference
					{Segment: "DTM", Mandatory: false, MaxRepeat: 9},  // Date/time/period
					{Segment: "STA", Mandatory: false, MaxRepeat: 99}, // Statistics
					{Segment: "FTX", Mandatory: false, MaxRepeat: 99}, // Free text
					{ // Segment group 9
						Group: []SchemaNode{
							{Segment: "CTA", Mandatory: true, MaxRepeat: 1},  // Contact information
							{Segment: "COM", Mandatory: false, MaxRepeat: 9}, // Communication contact
						},
						Mandatory: false, MaxRepeat: 9,
					},
					{ // Segment group 10
						Group: []SchemaNode{
							{Segment: "MOA", Mandatory: true, MaxRepeat: 1},  // Monetary amount
							{Segment: "DTM", Mandatory: false, MaxRepeat: 9}, // Date/time/period
						},
						Mandatory: false, MaxRepeat: 99,
					},
					{ // Segment group 11
						Group: []SchemaNode{
							{Segment: "QTY", Mandatory: true, MaxRepeat: 1},  // Quantity
							{Segment: "DTM", Mandatory: false, MaxRepeat: 9}, // Date/time/period
						},
						Mandatory: false, MaxRepeat: 99,
					},
					{ // Segment group 12
						Group: []SchemaNode{
							{Segment: "PCD", Mandatory: true, MaxRepeat: 1},  // Percentage details
							{Segment: "DTM", Mandatory: false, MaxRepeat: 9}, // Date/time/period
						},
						Mandatory: false, MaxRepeat: 99,
					},
					{ // Segment group 13
						Group: []SchemaNode{
							{Segment: "CAV", Mandatory: true, MaxRepeat: 1},   // Characteristic value
							{Segment: "PRC", Mandatory: false, MaxRepeat: 1},  // Process identification
							{Segment: "IMD", Mandatory: false, MaxRepeat: 1},  // Item description
							{Segment: "RFF", Mandatory: false, MaxRepeat: 9},  // Reference
							{Segment: "DTM", Mandatory: false, MaxRepeat: 9},  // Date/time/period
							{Segment: "STA", Mandatory: false, MaxRepeat: 99}, // Statistics
							{ // Segment group 14
								Group: []SchemaNode{
									{Segment: "MOA", Mandatory: true, MaxRepeat: 1},  // Monetary amount
									{Segment: "DTM", Mandatory: false, MaxRepeat: 9}, // Date/time/period
								},
								Mandatory: false, MaxRepeat: 99,
							},
							{ // Segment group 15
								Group: []SchemaNode{
									{Segment: "QTY", Mandatory: true, MaxRepeat: 1},  // Quantity
									{Segment: "DTM", Mandatory: false, MaxRepeat: 9}, // Date/time/period
								},
								Mandatory: false, MaxRepeat: 99,
							},
							{ // Segment group 16
								Group: []SchemaNode{
									{Segment: "PCD", Mandatory: true, MaxRepeat: 1},  // Percentage details
									{Segment: "DTM", Mandatory: false, MaxRepeat: 9}, // Date/time/period
								},
								Mandatory: false, MaxRepeat: 99,
							},
						},
						Mandatory: false, MaxRepeat: 999,
					},
				},
				Mandatory: false, MaxRepeat: 9999,
			},
		},
		Mandatory: true, MaxRepeat: 99,
	},
	{Segment: "CNT", Mandatory: false, MaxRepeat: 9}, // Control total
}}
