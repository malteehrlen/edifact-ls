package edifact

// SUPRES D.20A message specification: the structural (segment/group
// presence, order, cardinality) branching diagram for this message,
// UN/EDIFACT directory release D.20A. Scope is structural only -- not
// element-level content -- per edifact-ls-3uzr's epic non-goal.
// 22 segment groups, max nesting depth 5.
//
// Source: https://service.unece.org/trade/untdid/d20a/trmd/supres_c.htm
// section 4.3.1 "Segment table". That URL currently returns HTTP 403 from
// Cloudflare when fetched directly; this data was transcribed from the
// Wayback Machine's archived copy instead:
// http://web.archive.org/web/20240421085347/https://service.unece.org/trade/untdid/d20a/trmd/supres_c.htm
// -- re-check the direct URL first if cross-checking later, in case the
// block has lifted.
//
// Transcription note: generated in bulk (edifact-ls-13gu) the same way
// as every other message-type schema -- a script parsed the source's
// exact rail-art column positions mechanically and verified the result
// balances (all 22 groups open and close consistently)
// before being transcribed here, rather than reading the ASCII nesting
// by eye. Not individually spot-checked against the raw source the way
// the first several message types in this project were -- see
// edifact-ls-13gu for which representative samples were spot-checked
// for this batch.

func init() {
	RegisterSchema(
		MessageID{Type: "SUPRES", Version: "D", Release: "20A", Agency: "UN"},
		supresD20aSchema,
		"https://service.unece.org/trade/untdid/d20a/trmd/supres_c.htm",
	)
}

var supresD20aSchema = Schema{Nodes: []SchemaNode{
	{Segment: "BGM", Mandatory: true, MaxRepeat: 1},  // Beginning of message
	{Segment: "LAN", Mandatory: false, MaxRepeat: 1}, // Language
	{Segment: "PCD", Mandatory: false, MaxRepeat: 1}, // Percentage details
	{Segment: "DTM", Mandatory: false, MaxRepeat: 5}, // Date/time/period
	{Segment: "FTX", Mandatory: false, MaxRepeat: 5}, // Free text
	{Segment: "LOC", Mandatory: false, MaxRepeat: 2}, // Place/location identification
	{Segment: "MOA", Mandatory: false, MaxRepeat: 5}, // Monetary amount
	{Segment: "FII", Mandatory: false, MaxRepeat: 5}, // Financial institution information
	{ // Segment group 1
		Group: []SchemaNode{
			{Segment: "RFF", Mandatory: true, MaxRepeat: 1},  // Reference
			{Segment: "DTM", Mandatory: false, MaxRepeat: 1}, // Date/time/period
		},
		Mandatory: false, MaxRepeat: 5,
	},
	{ // Segment group 2
		Group: []SchemaNode{
			{Segment: "NAD", Mandatory: true, MaxRepeat: 1},  // Name and address
			{Segment: "RFF", Mandatory: false, MaxRepeat: 5}, // Reference
			{Segment: "COM", Mandatory: false, MaxRepeat: 5}, // Communication contact
			{Segment: "LAN", Mandatory: false, MaxRepeat: 1}, // Language
			{ // Segment group 3
				Group: []SchemaNode{
					{Segment: "CTA", Mandatory: true, MaxRepeat: 1},  // Contact information
					{Segment: "COM", Mandatory: false, MaxRepeat: 5}, // Communication contact
				},
				Mandatory: false, MaxRepeat: 5,
			},
		},
		Mandatory: false, MaxRepeat: 5,
	},
	{ // Segment group 4
		Group: []SchemaNode{
			{Segment: "PAI", Mandatory: true, MaxRepeat: 1},   // Payment instructions
			{Segment: "RFF", Mandatory: false, MaxRepeat: 10}, // Reference
		},
		Mandatory: false, MaxRepeat: 1,
	},
	{ // Segment group 5
		Group: []SchemaNode{
			{Segment: "PYT", Mandatory: true, MaxRepeat: 1},   // Payment terms
			{Segment: "DTM", Mandatory: false, MaxRepeat: 10}, // Date/time/period
		},
		Mandatory: false, MaxRepeat: 1,
	},
	{ // Segment group 6
		Group: []SchemaNode{
			{Segment: "SEQ", Mandatory: true, MaxRepeat: 1}, // Sequence details
			{ // Segment group 7
				Group: []SchemaNode{
					{Segment: "IMD", Mandatory: true, MaxRepeat: 1},   // Item description
					{Segment: "QTY", Mandatory: false, MaxRepeat: 1},  // Quantity
					{Segment: "LAN", Mandatory: false, MaxRepeat: 1},  // Language
					{Segment: "DTM", Mandatory: false, MaxRepeat: 20}, // Date/time/period
					{Segment: "RFF", Mandatory: false, MaxRepeat: 99}, // Reference
					{Segment: "PCD", Mandatory: false, MaxRepeat: 1},  // Percentage details
					{Segment: "DIM", Mandatory: false, MaxRepeat: 2},  // Dimensions
					{Segment: "FTX", Mandatory: false, MaxRepeat: 5},  // Free text
					{ // Segment group 8
						Group: []SchemaNode{
							{Segment: "NAD", Mandatory: true, MaxRepeat: 1},  // Name and address
							{Segment: "COM", Mandatory: false, MaxRepeat: 5}, // Communication contact
							{Segment: "LAN", Mandatory: false, MaxRepeat: 1}, // Language
							{ // Segment group 9
								Group: []SchemaNode{
									{Segment: "CTA", Mandatory: true, MaxRepeat: 1},  // Contact information
									{Segment: "COM", Mandatory: false, MaxRepeat: 5}, // Communication contact
								},
								Mandatory: false, MaxRepeat: 1,
							},
						},
						Mandatory: false, MaxRepeat: 1,
					},
					{ // Segment group 10
						Group: []SchemaNode{
							{Segment: "LOC", Mandatory: true, MaxRepeat: 1},   // Place/location identification
							{Segment: "DTM", Mandatory: false, MaxRepeat: 10}, // Date/time/period
							{ // Segment group 11
								Group: []SchemaNode{
									{Segment: "NAD", Mandatory: true, MaxRepeat: 1},   // Name and address
									{Segment: "DTM", Mandatory: false, MaxRepeat: 10}, // Date/time/period
									{ // Segment group 12
										Group: []SchemaNode{
											{Segment: "CTA", Mandatory: true, MaxRepeat: 1},  // Contact information
											{Segment: "COM", Mandatory: false, MaxRepeat: 5}, // Communication contact
										},
										Mandatory: false, MaxRepeat: 5,
									},
								},
								Mandatory: false, MaxRepeat: 1,
							},
						},
						Mandatory: false, MaxRepeat: 10,
					},
					{ // Segment group 13
						Group: []SchemaNode{
							{Segment: "MOA", Mandatory: true, MaxRepeat: 1},  // Monetary amount
							{Segment: "QTY", Mandatory: false, MaxRepeat: 1}, // Quantity
						},
						Mandatory: false, MaxRepeat: 99,
					},
					{ // Segment group 14
						Group: []SchemaNode{
							{Segment: "PAI", Mandatory: true, MaxRepeat: 1},  // Payment instructions
							{Segment: "RFF", Mandatory: false, MaxRepeat: 1}, // Reference
						},
						Mandatory: false, MaxRepeat: 1,
					},
					{ // Segment group 15
						Group: []SchemaNode{
							{Segment: "PYT", Mandatory: true, MaxRepeat: 1},  // Payment terms
							{Segment: "DTM", Mandatory: false, MaxRepeat: 1}, // Date/time/period
						},
						Mandatory: false, MaxRepeat: 1,
					},
					{ // Segment group 16
						Group: []SchemaNode{
							{Segment: "RCS", Mandatory: true, MaxRepeat: 1},  // Requirements and conditions
							{Segment: "FII", Mandatory: false, MaxRepeat: 5}, // Financial institution information
							{Segment: "MOA", Mandatory: false, MaxRepeat: 2}, // Monetary amount
							{Segment: "RFF", Mandatory: false, MaxRepeat: 1}, // Reference
							{Segment: "FTX", Mandatory: false, MaxRepeat: 1}, // Free text
							{Segment: "DOC", Mandatory: false, MaxRepeat: 1}, // Document/message details
						},
						Mandatory: false, MaxRepeat: 99,
					},
				},
				Mandatory: false, MaxRepeat: 99,
			},
			{ // Segment group 17
				Group: []SchemaNode{
					{Segment: "MEM", Mandatory: true, MaxRepeat: 1}, // Membership details
					{ // Segment group 18
						Group: []SchemaNode{
							{Segment: "NAD", Mandatory: true, MaxRepeat: 1},   // Name and address
							{Segment: "ATT", Mandatory: false, MaxRepeat: 99}, // Attribute
							{Segment: "RFF", Mandatory: false, MaxRepeat: 99}, // Reference
							{Segment: "LAN", Mandatory: false, MaxRepeat: 1},  // Language
							{Segment: "FTX", Mandatory: false, MaxRepeat: 5},  // Free text
							{Segment: "COM", Mandatory: false, MaxRepeat: 5},  // Communication contact
							{ // Segment group 19
								Group: []SchemaNode{
									{Segment: "CTA", Mandatory: true, MaxRepeat: 1},  // Contact information
									{Segment: "COM", Mandatory: false, MaxRepeat: 5}, // Communication contact
								},
								Mandatory: false, MaxRepeat: 1,
							},
							{ // Segment group 20
								Group: []SchemaNode{
									{Segment: "DOC", Mandatory: true, MaxRepeat: 1},   // Document/message details
									{Segment: "LOC", Mandatory: false, MaxRepeat: 1},  // Place/location identification
									{Segment: "DTM", Mandatory: false, MaxRepeat: 10}, // Date/time/period
								},
								Mandatory: false, MaxRepeat: 5,
							},
							{ // Segment group 21
								Group: []SchemaNode{
									{Segment: "PAI", Mandatory: true, MaxRepeat: 1},  // Payment instructions
									{Segment: "PYT", Mandatory: false, MaxRepeat: 1}, // Payment terms
									{Segment: "FII", Mandatory: false, MaxRepeat: 1}, // Financial institution information
									{Segment: "MOA", Mandatory: false, MaxRepeat: 1}, // Monetary amount
									{Segment: "RFF", Mandatory: false, MaxRepeat: 9}, // Reference
								},
								Mandatory: false, MaxRepeat: 10,
							},
							{ // Segment group 22
								Group: []SchemaNode{
									{Segment: "RCS", Mandatory: true, MaxRepeat: 1},  // Requirements and conditions
									{Segment: "MOA", Mandatory: false, MaxRepeat: 2}, // Monetary amount
									{Segment: "RFF", Mandatory: false, MaxRepeat: 2}, // Reference
									{Segment: "FTX", Mandatory: false, MaxRepeat: 1}, // Free text
								},
								Mandatory: false, MaxRepeat: 10,
							},
						},
						Mandatory: false, MaxRepeat: 99,
					},
				},
				Mandatory: false, MaxRepeat: 99,
			},
		},
		Mandatory: false, MaxRepeat: 99,
	},
}}
