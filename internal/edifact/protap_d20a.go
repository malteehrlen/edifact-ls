package edifact

// PROTAP D.20A message specification: the structural (segment/group
// presence, order, cardinality) branching diagram for this message,
// UN/EDIFACT directory release D.20A. Scope is structural only -- not
// element-level content -- per edifact-ls-3uzr's epic non-goal.
// 25 segment groups, max nesting depth 3.
//
// Source: https://service.unece.org/trade/untdid/d20a/trmd/protap_c.htm
// section 4.3.1 "Segment table". That URL currently returns HTTP 403 from
// Cloudflare when fetched directly; this data was transcribed from the
// Wayback Machine's archived copy instead:
// http://web.archive.org/web/20231202075657/https://service.unece.org/trade/untdid/d20a/trmd/protap_c.htm
// -- re-check the direct URL first if cross-checking later, in case the
// block has lifted.
//
// Transcription note: generated in bulk (edifact-ls-13gu) the same way
// as every other message-type schema -- a script parsed the source's
// exact rail-art column positions mechanically and verified the result
// balances (all 25 groups open and close consistently)
// before being transcribed here, rather than reading the ASCII nesting
// by eye. Not individually spot-checked against the raw source the way
// the first several message types in this project were -- see
// edifact-ls-13gu for which representative samples were spot-checked
// for this batch.

func init() {
	RegisterSchema(
		MessageID{Type: "PROTAP", Version: "D", Release: "20A", Agency: "UN"},
		protapD20aSchema,
		"https://service.unece.org/trade/untdid/d20a/trmd/protap_c.htm",
	)
}

var protapD20aSchema = Schema{Nodes: []SchemaNode{
	{Segment: "BGM", Mandatory: true, MaxRepeat: 1},  // Beginning of message
	{Segment: "GEI", Mandatory: true, MaxRepeat: 1},  // Processing information
	{Segment: "DTM", Mandatory: true, MaxRepeat: 99}, // Date/time/period
	{ // Segment group 1
		Group: []SchemaNode{
			{Segment: "RFF", Mandatory: true, MaxRepeat: 1},   // Reference
			{Segment: "IMD", Mandatory: false, MaxRepeat: 9},  // Item description
			{Segment: "BII", Mandatory: false, MaxRepeat: 99}, // Structure identification
			{ // Segment group 2
				Group: []SchemaNode{
					{Segment: "GIR", Mandatory: true, MaxRepeat: 1},  // Related identification numbers
					{Segment: "REL", Mandatory: false, MaxRepeat: 1}, // Relationship
				},
				Mandatory: false, MaxRepeat: 99,
			},
		},
		Mandatory: false, MaxRepeat: 99,
	},
	{ // Segment group 3
		Group: []SchemaNode{
			{Segment: "PNA", Mandatory: true, MaxRepeat: 1},  // Party identification
			{Segment: "ADR", Mandatory: false, MaxRepeat: 1}, // Address
			{Segment: "BII", Mandatory: false, MaxRepeat: 1}, // Structure identification
			{ // Segment group 4
				Group: []SchemaNode{
					{Segment: "CTA", Mandatory: true, MaxRepeat: 1},  // Contact information
					{Segment: "COM", Mandatory: false, MaxRepeat: 9}, // Communication contact
				},
				Mandatory: false, MaxRepeat: 9,
			},
			{ // Segment group 5
				Group: []SchemaNode{
					{Segment: "CED", Mandatory: true, MaxRepeat: 1},  // Computer environment details
					{Segment: "DTM", Mandatory: false, MaxRepeat: 9}, // Date/time/period
				},
				Mandatory: false, MaxRepeat: 9,
			},
		},
		Mandatory: false, MaxRepeat: 99,
	},
	{ // Segment group 6
		Group: []SchemaNode{
			{Segment: "EFI", Mandatory: true, MaxRepeat: 1},  // External file link identification
			{Segment: "CED", Mandatory: false, MaxRepeat: 1}, // Computer environment details
			{Segment: "DTM", Mandatory: false, MaxRepeat: 1}, // Date/time/period
		},
		Mandatory: false, MaxRepeat: 99,
	},
	{ // Segment group 7
		Group: []SchemaNode{
			{Segment: "IND", Mandatory: true, MaxRepeat: 1},  // Index details
			{Segment: "BII", Mandatory: false, MaxRepeat: 1}, // Structure identification
			{Segment: "IMD", Mandatory: false, MaxRepeat: 1}, // Item description
			{Segment: "RCS", Mandatory: false, MaxRepeat: 1}, // Requirements and conditions
		},
		Mandatory: false, MaxRepeat: 999,
	},
	{ // Segment group 8
		Group: []SchemaNode{
			{Segment: "LIN", Mandatory: true, MaxRepeat: 1},   // Line item
			{Segment: "BII", Mandatory: false, MaxRepeat: 9},  // Structure identification
			{Segment: "RFF", Mandatory: false, MaxRepeat: 9},  // Reference
			{Segment: "IMD", Mandatory: false, MaxRepeat: 9},  // Item description
			{Segment: "DTM", Mandatory: false, MaxRepeat: 99}, // Date/time/period
			{ // Segment group 9
				Group: []SchemaNode{
					{Segment: "QTY", Mandatory: true, MaxRepeat: 1},   // Quantity
					{Segment: "DTM", Mandatory: false, MaxRepeat: 99}, // Date/time/period
				},
				Mandatory: false, MaxRepeat: 99,
			},
			{ // Segment group 10
				Group: []SchemaNode{
					{Segment: "PRI", Mandatory: true, MaxRepeat: 1},   // Price details
					{Segment: "DTM", Mandatory: false, MaxRepeat: 99}, // Date/time/period
				},
				Mandatory: false, MaxRepeat: 99,
			},
			{ // Segment group 11
				Group: []SchemaNode{
					{Segment: "PCD", Mandatory: true, MaxRepeat: 1},   // Percentage details
					{Segment: "DTM", Mandatory: false, MaxRepeat: 99}, // Date/time/period
				},
				Mandatory: false, MaxRepeat: 99,
			},
			{ // Segment group 12
				Group: []SchemaNode{
					{Segment: "RTE", Mandatory: true, MaxRepeat: 1},   // Rate details
					{Segment: "DTM", Mandatory: false, MaxRepeat: 99}, // Date/time/period
				},
				Mandatory: false, MaxRepeat: 99,
			},
			{ // Segment group 13
				Group: []SchemaNode{
					{Segment: "MOA", Mandatory: true, MaxRepeat: 1},   // Monetary amount
					{Segment: "DTM", Mandatory: false, MaxRepeat: 99}, // Date/time/period
				},
				Mandatory: false, MaxRepeat: 99,
			},
		},
		Mandatory: false, MaxRepeat: 9999,
	},
	{Segment: "UNS", Mandatory: true, MaxRepeat: 1}, // Section control
	{ // Segment group 14
		Group: []SchemaNode{
			{Segment: "RFF", Mandatory: true, MaxRepeat: 1}, // Reference
			{ // Segment group 15
				Group: []SchemaNode{
					{Segment: "SCC", Mandatory: true, MaxRepeat: 1},    // Scheduling conditions
					{Segment: "QTY", Mandatory: false, MaxRepeat: 99},  // Quantity
					{Segment: "DTM", Mandatory: false, MaxRepeat: 999}, // Date/time/period
				},
				Mandatory: false, MaxRepeat: 999,
			},
		},
		Mandatory: false, MaxRepeat: 99,
	},
	{ // Segment group 16
		Group: []SchemaNode{
			{Segment: "LIN", Mandatory: true, MaxRepeat: 1},   // Line item
			{Segment: "STS", Mandatory: false, MaxRepeat: 1},  // Status
			{Segment: "BII", Mandatory: false, MaxRepeat: 9},  // Structure identification
			{Segment: "RFF", Mandatory: false, MaxRepeat: 9},  // Reference
			{Segment: "IMD", Mandatory: false, MaxRepeat: 9},  // Item description
			{Segment: "PNA", Mandatory: false, MaxRepeat: 9},  // Party identification
			{Segment: "CCI", Mandatory: false, MaxRepeat: 9},  // Characteristic/class id
			{Segment: "DTM", Mandatory: false, MaxRepeat: 99}, // Date/time/period
			{Segment: "FTX", Mandatory: false, MaxRepeat: 99}, // Free text
			{ // Segment group 17
				Group: []SchemaNode{
					{Segment: "CTA", Mandatory: true, MaxRepeat: 1},  // Contact information
					{Segment: "COM", Mandatory: false, MaxRepeat: 9}, // Communication contact
				},
				Mandatory: false, MaxRepeat: 9,
			},
			{ // Segment group 18
				Group: []SchemaNode{
					{Segment: "QTY", Mandatory: true, MaxRepeat: 1},   // Quantity
					{Segment: "DTM", Mandatory: false, MaxRepeat: 99}, // Date/time/period
				},
				Mandatory: false, MaxRepeat: 99,
			},
			{ // Segment group 19
				Group: []SchemaNode{
					{Segment: "PCD", Mandatory: true, MaxRepeat: 1},   // Percentage details
					{Segment: "DTM", Mandatory: false, MaxRepeat: 99}, // Date/time/period
				},
				Mandatory: false, MaxRepeat: 99,
			},
			{ // Segment group 20
				Group: []SchemaNode{
					{Segment: "MOA", Mandatory: true, MaxRepeat: 1},   // Monetary amount
					{Segment: "DTM", Mandatory: false, MaxRepeat: 99}, // Date/time/period
				},
				Mandatory: false, MaxRepeat: 99,
			},
			{ // Segment group 21
				Group: []SchemaNode{
					{Segment: "GIR", Mandatory: true, MaxRepeat: 1},   // Related identification numbers
					{Segment: "REL", Mandatory: false, MaxRepeat: 1},  // Relationship
					{Segment: "RCS", Mandatory: false, MaxRepeat: 1},  // Requirements and conditions
					{Segment: "BII", Mandatory: false, MaxRepeat: 9},  // Structure identification
					{Segment: "RFF", Mandatory: false, MaxRepeat: 9},  // Reference
					{Segment: "IMD", Mandatory: false, MaxRepeat: 9},  // Item description
					{Segment: "DTM", Mandatory: false, MaxRepeat: 99}, // Date/time/period
					{ // Segment group 22
						Group: []SchemaNode{
							{Segment: "QTY", Mandatory: true, MaxRepeat: 1},   // Quantity
							{Segment: "DTM", Mandatory: false, MaxRepeat: 99}, // Date/time/period
						},
						Mandatory: false, MaxRepeat: 99,
					},
					{ // Segment group 23
						Group: []SchemaNode{
							{Segment: "PCD", Mandatory: true, MaxRepeat: 1},   // Percentage details
							{Segment: "DTM", Mandatory: false, MaxRepeat: 99}, // Date/time/period
						},
						Mandatory: false, MaxRepeat: 99,
					},
					{ // Segment group 24
						Group: []SchemaNode{
							{Segment: "MOA", Mandatory: true, MaxRepeat: 1},   // Monetary amount
							{Segment: "DTM", Mandatory: false, MaxRepeat: 99}, // Date/time/period
						},
						Mandatory: false, MaxRepeat: 99,
					},
					{ // Segment group 25
						Group: []SchemaNode{
							{Segment: "RTE", Mandatory: true, MaxRepeat: 1},   // Rate details
							{Segment: "DTM", Mandatory: false, MaxRepeat: 99}, // Date/time/period
						},
						Mandatory: false, MaxRepeat: 99,
					},
				},
				Mandatory: false, MaxRepeat: 9999,
			},
		},
		Mandatory: false, MaxRepeat: 9999,
	},
	{Segment: "CNT", Mandatory: false, MaxRepeat: 99}, // Control total
}}
