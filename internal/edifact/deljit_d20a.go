package edifact

// DELJIT D.20A message specification: the structural (segment/group
// presence, order, cardinality) branching diagram for this message,
// UN/EDIFACT directory release D.20A. Scope is structural only -- not
// element-level content -- per edifact-ls-3uzr's epic non-goal.
// 16 segment groups, max nesting depth 4.
//
// Source: https://service.unece.org/trade/untdid/d20a/trmd/deljit_c.htm
// section 4.3.1 "Segment table". That URL currently returns HTTP 403 from
// Cloudflare when fetched directly; this data was transcribed from the
// Wayback Machine's archived copy instead:
// http://web.archive.org/web/20231002220820/https://service.unece.org/trade/untdid/d20a/trmd/deljit_c.htm
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
		MessageID{Type: "DELJIT", Version: "D", Release: "20A", Agency: "UN"},
		deljitD20aSchema,
		"https://service.unece.org/trade/untdid/d20a/trmd/deljit_c.htm",
	)
}

var deljitD20aSchema = Schema{Nodes: []SchemaNode{
	{Segment: "BGM", Mandatory: true, MaxRepeat: 1},  // Beginning of message
	{Segment: "DTM", Mandatory: true, MaxRepeat: 10}, // Date/time/period
	{Segment: "FTX", Mandatory: false, MaxRepeat: 5}, // Free text
	{ // Segment group 1
		Group: []SchemaNode{
			{Segment: "RFF", Mandatory: true, MaxRepeat: 1},  // Reference
			{Segment: "DTM", Mandatory: false, MaxRepeat: 1}, // Date/time/period
		},
		Mandatory: false, MaxRepeat: 10,
	},
	{ // Segment group 2
		Group: []SchemaNode{
			{Segment: "NAD", Mandatory: true, MaxRepeat: 1},   // Name and address
			{Segment: "LOC", Mandatory: false, MaxRepeat: 10}, // Place/location identification
			{Segment: "FTX", Mandatory: false, MaxRepeat: 5},  // Free text
			{ // Segment group 3
				Group: []SchemaNode{
					{Segment: "RFF", Mandatory: true, MaxRepeat: 1},  // Reference
					{Segment: "DTM", Mandatory: false, MaxRepeat: 9}, // Date/time/period
				},
				Mandatory: false, MaxRepeat: 99,
			},
			{ // Segment group 4
				Group: []SchemaNode{
					{Segment: "CTA", Mandatory: true, MaxRepeat: 1},  // Contact information
					{Segment: "COM", Mandatory: false, MaxRepeat: 5}, // Communication contact
				},
				Mandatory: false, MaxRepeat: 5,
			},
		},
		Mandatory: false, MaxRepeat: 20,
	},
	{ // Segment group 5
		Group: []SchemaNode{
			{Segment: "SEQ", Mandatory: true, MaxRepeat: 1},   // Sequence details
			{Segment: "DTM", Mandatory: false, MaxRepeat: 5},  // Date/time/period
			{Segment: "GIR", Mandatory: false, MaxRepeat: 99}, // Related identification numbers
			{Segment: "LOC", Mandatory: false, MaxRepeat: 5},  // Place/location identification
			{ // Segment group 6
				Group: []SchemaNode{
					{Segment: "RFF", Mandatory: true, MaxRepeat: 1},  // Reference
					{Segment: "DTM", Mandatory: false, MaxRepeat: 9}, // Date/time/period
				},
				Mandatory: false, MaxRepeat: 9,
			},
			{ // Segment group 7
				Group: []SchemaNode{
					{Segment: "PAC", Mandatory: true, MaxRepeat: 1}, // Package
					{ // Segment group 8
						Group: []SchemaNode{
							{Segment: "PCI", Mandatory: true, MaxRepeat: 1},   // Package identification
							{Segment: "GIN", Mandatory: false, MaxRepeat: 10}, // Goods identity number
						},
						Mandatory: false, MaxRepeat: 999,
					},
				},
				Mandatory: false, MaxRepeat: 5,
			},
			{ // Segment group 9
				Group: []SchemaNode{
					{Segment: "LIN", Mandatory: true, MaxRepeat: 1},   // Line item
					{Segment: "PIA", Mandatory: false, MaxRepeat: 10}, // Additional product id
					{Segment: "IMD", Mandatory: false, MaxRepeat: 99}, // Item description
					{Segment: "ALI", Mandatory: false, MaxRepeat: 5},  // Additional information
					{Segment: "GIR", Mandatory: false, MaxRepeat: 5},  // Related identification numbers
					{Segment: "FTX", Mandatory: false, MaxRepeat: 5},  // Free text
					{Segment: "PAC", Mandatory: false, MaxRepeat: 99}, // Package
					{Segment: "DTM", Mandatory: false, MaxRepeat: 9},  // Date/time/period
					{ // Segment group 10
						Group: []SchemaNode{
							{Segment: "RFF", Mandatory: true, MaxRepeat: 1},  // Reference
							{Segment: "DTM", Mandatory: false, MaxRepeat: 1}, // Date/time/period
						},
						Mandatory: false, MaxRepeat: 9,
					},
					{ // Segment group 11
						Group: []SchemaNode{
							{Segment: "TDT", Mandatory: true, MaxRepeat: 1},  // Transport information
							{Segment: "TMD", Mandatory: false, MaxRepeat: 1}, // Transport movement details
						},
						Mandatory: false, MaxRepeat: 9,
					},
					{ // Segment group 12
						Group: []SchemaNode{
							{Segment: "LOC", Mandatory: true, MaxRepeat: 1}, // Place/location identification
							{ // Segment group 13
								Group: []SchemaNode{
									{Segment: "CTA", Mandatory: true, MaxRepeat: 1},  // Contact information
									{Segment: "COM", Mandatory: false, MaxRepeat: 5}, // Communication contact
								},
								Mandatory: false, MaxRepeat: 5,
							},
						},
						Mandatory: false, MaxRepeat: 99,
					},
					{ // Segment group 14
						Group: []SchemaNode{
							{Segment: "QTY", Mandatory: true, MaxRepeat: 1},  // Quantity
							{Segment: "SCC", Mandatory: false, MaxRepeat: 1}, // Scheduling conditions
							{Segment: "DTM", Mandatory: false, MaxRepeat: 9}, // Date/time/period
							{ // Segment group 15
								Group: []SchemaNode{
									{Segment: "RFF", Mandatory: true, MaxRepeat: 1},  // Reference
									{Segment: "DTM", Mandatory: false, MaxRepeat: 1}, // Date/time/period
								},
								Mandatory: false, MaxRepeat: 99,
							},
						},
						Mandatory: false, MaxRepeat: 999,
					},
					{ // Segment group 16
						Group: []SchemaNode{
							{Segment: "PRI", Mandatory: true, MaxRepeat: 1},  // Price details
							{Segment: "CUX", Mandatory: false, MaxRepeat: 1}, // Currencies
							{Segment: "DTM", Mandatory: false, MaxRepeat: 9}, // Date/time/period
						},
						Mandatory: false, MaxRepeat: 9,
					},
				},
				Mandatory: false, MaxRepeat: 9999,
			},
		},
		Mandatory: true, MaxRepeat: 9999,
	},
}}
