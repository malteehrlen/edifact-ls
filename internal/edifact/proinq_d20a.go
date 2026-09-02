package edifact

// PROINQ D.20A message specification: the structural (segment/group
// presence, order, cardinality) branching diagram for this message,
// UN/EDIFACT directory release D.20A. Scope is structural only -- not
// element-level content -- per edifact-ls-3uzr's epic non-goal.
// 7 segment groups, max nesting depth 3.
//
// Source: https://service.unece.org/trade/untdid/d20a/trmd/proinq_c.htm
// section 4.3.1 "Segment table". That URL currently returns HTTP 403 from
// Cloudflare when fetched directly; this data was transcribed from the
// Wayback Machine's archived copy instead:
// http://web.archive.org/web/20231202084858/https://service.unece.org/trade/untdid/d20a/trmd/proinq_c.htm
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
		MessageID{Type: "PROINQ", Version: "D", Release: "20A", Agency: "UN"},
		proinqD20aSchema,
		"https://service.unece.org/trade/untdid/d20a/trmd/proinq_c.htm",
	)
}

var proinqD20aSchema = Schema{Nodes: []SchemaNode{
	{Segment: "BGM", Mandatory: true, MaxRepeat: 1},  // Beginning of message
	{Segment: "DTM", Mandatory: true, MaxRepeat: 9},  // Date/time/period
	{Segment: "CUX", Mandatory: false, MaxRepeat: 1}, // Currencies
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
		Mandatory: true, MaxRepeat: 99,
	},
	{ // Segment group 3
		Group: []SchemaNode{
			{Segment: "RFF", Mandatory: true, MaxRepeat: 1},  // Reference
			{Segment: "DTM", Mandatory: false, MaxRepeat: 9}, // Date/time/period
		},
		Mandatory: false, MaxRepeat: 9,
	},
	{ // Segment group 4
		Group: []SchemaNode{
			{Segment: "LIN", Mandatory: true, MaxRepeat: 1}, // Line item
			{ // Segment group 5
				Group: []SchemaNode{
					{Segment: "CCI", Mandatory: true, MaxRepeat: 1},   // Characteristic/class id
					{Segment: "CAV", Mandatory: false, MaxRepeat: 10}, // Characteristic value
					{Segment: "MEA", Mandatory: false, MaxRepeat: 10}, // Measurements
				},
				Mandatory: false, MaxRepeat: 999,
			},
			{ // Segment group 6
				Group: []SchemaNode{
					{Segment: "IRQ", Mandatory: true, MaxRepeat: 1},   // Information required
					{Segment: "PIA", Mandatory: false, MaxRepeat: 10}, // Additional product id
					{Segment: "IMD", Mandatory: false, MaxRepeat: 10}, // Item description
					{Segment: "MEA", Mandatory: false, MaxRepeat: 10}, // Measurements
					{Segment: "NAD", Mandatory: false, MaxRepeat: 10}, // Name and address
					{Segment: "PGI", Mandatory: false, MaxRepeat: 9},  // Product group information
					{Segment: "DTM", Mandatory: false, MaxRepeat: 5},  // Date/time/period
					{ // Segment group 7
						Group: []SchemaNode{
							{Segment: "PRI", Mandatory: true, MaxRepeat: 1},  // Price details
							{Segment: "RNG", Mandatory: false, MaxRepeat: 1}, // Range details
						},
						Mandatory: false, MaxRepeat: 10,
					},
				},
				Mandatory: false, MaxRepeat: 9999,
			},
		},
		Mandatory: true, MaxRepeat: 9999,
	},
}}
