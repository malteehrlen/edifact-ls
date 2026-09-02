package edifact

// OSTENQ D.20A message specification: the structural (segment/group
// presence, order, cardinality) branching diagram for this message,
// UN/EDIFACT directory release D.20A. Scope is structural only -- not
// element-level content -- per edifact-ls-3uzr's epic non-goal.
// 7 segment groups, max nesting depth 3.
//
// Source: https://service.unece.org/trade/untdid/d20a/trmd/ostenq_c.htm
// section 4.3.1 "Segment table". That URL currently returns HTTP 403 from
// Cloudflare when fetched directly; this data was transcribed from the
// Wayback Machine's archived copy instead:
// http://web.archive.org/web/20231202085133/https://service.unece.org/trade/untdid/d20a/trmd/ostenq_c.htm
// -- re-check the direct URL first if cross-checking later, in case the
// block has lifted.
//
// Transcription note: generated in bulk (edifact-ls-13gu) the same way
// as every other message-type schema -- a script parsed the source's
// exact rail-art column positions mechanically and verified the result
// balances (all 7 groups open and close consistently)
// before being transcribed here, rather than reading the ASCII nesting
// by eye. Not individually spot-checked against the raw source the way
// the first several message types in this project were -- see
// edifact-ls-13gu for which representative samples were spot-checked
// for this batch.

func init() {
	RegisterSchema(
		MessageID{Type: "OSTENQ", Version: "D", Release: "20A", Agency: "UN"},
		ostenqD20aSchema,
		"https://service.unece.org/trade/untdid/d20a/trmd/ostenq_c.htm",
	)
}

var ostenqD20aSchema = Schema{Nodes: []SchemaNode{
	{Segment: "BGM", Mandatory: true, MaxRepeat: 1},   // Beginning of message
	{Segment: "DTM", Mandatory: true, MaxRepeat: 5},   // Date/time/period
	{Segment: "IRQ", Mandatory: false, MaxRepeat: 1},  // Information required
	{Segment: "FTX", Mandatory: false, MaxRepeat: 20}, // Free text
	{ // Segment group 1
		Group: []SchemaNode{
			{Segment: "RFF", Mandatory: true, MaxRepeat: 1},  // Reference
			{Segment: "DTM", Mandatory: false, MaxRepeat: 9}, // Date/time/period
		},
		Mandatory: false, MaxRepeat: 9,
	},
	{ // Segment group 2
		Group: []SchemaNode{
			{Segment: "NAD", Mandatory: true, MaxRepeat: 1}, // Name and address
			{ // Segment group 3
				Group: []SchemaNode{
					{Segment: "CTA", Mandatory: true, MaxRepeat: 1},  // Contact information
					{Segment: "COM", Mandatory: false, MaxRepeat: 5}, // Communication contact
				},
				Mandatory: false, MaxRepeat: 5,
			},
		},
		Mandatory: true, MaxRepeat: 25,
	},
	{ // Segment group 4
		Group: []SchemaNode{
			{Segment: "DOC", Mandatory: true, MaxRepeat: 1},  // Document/message details
			{Segment: "DTM", Mandatory: false, MaxRepeat: 5}, // Date/time/period
			{ // Segment group 5
				Group: []SchemaNode{
					{Segment: "LIN", Mandatory: true, MaxRepeat: 1},   // Line item
					{Segment: "PIA", Mandatory: false, MaxRepeat: 25}, // Additional product id
					{Segment: "IMD", Mandatory: false, MaxRepeat: 10}, // Item description
					{Segment: "MEA", Mandatory: false, MaxRepeat: 5},  // Measurements
					{Segment: "STS", Mandatory: false, MaxRepeat: 9},  // Status
					{ // Segment group 6
						Group: []SchemaNode{
							{Segment: "RFF", Mandatory: true, MaxRepeat: 1},  // Reference
							{Segment: "DTM", Mandatory: false, MaxRepeat: 1}, // Date/time/period
						},
						Mandatory: false, MaxRepeat: 99,
					},
					{ // Segment group 7
						Group: []SchemaNode{
							{Segment: "LOC", Mandatory: true, MaxRepeat: 1},  // Place/location identification
							{Segment: "DTM", Mandatory: false, MaxRepeat: 5}, // Date/time/period
							{Segment: "QTY", Mandatory: false, MaxRepeat: 1}, // Quantity
						},
						Mandatory: false, MaxRepeat: 99,
					},
				},
				Mandatory: false, MaxRepeat: 200000,
			},
		},
		Mandatory: false, MaxRepeat: 999,
	},
}}
