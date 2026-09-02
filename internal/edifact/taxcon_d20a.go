package edifact

// TAXCON D.20A message specification: the structural (segment/group
// presence, order, cardinality) branching diagram for this message,
// UN/EDIFACT directory release D.20A. Scope is structural only -- not
// element-level content -- per edifact-ls-3uzr's epic non-goal.
// 5 segment groups, max nesting depth 3.
//
// Source: https://service.unece.org/trade/untdid/d20a/trmd/taxcon_c.htm
// section 4.3.1 "Segment table". That URL currently returns HTTP 403 from
// Cloudflare when fetched directly; this data was transcribed from the
// Wayback Machine's archived copy instead:
// http://web.archive.org/web/20251116003832/https://service.unece.org/trade/untdid/d20a/trmd/taxcon_c.htm
// -- re-check the direct URL first if cross-checking later, in case the
// block has lifted.
//
// Transcription note: generated in bulk (edifact-ls-13gu) the same way
// as every other message-type schema -- a script parsed the source's
// exact rail-art column positions mechanically and verified the result
// balances (all 5 groups open and close consistently)
// before being transcribed here, rather than reading the ASCII nesting
// by eye. Not individually spot-checked against the raw source the way
// the first several message types in this project were -- see
// edifact-ls-13gu for which representative samples were spot-checked
// for this batch.

func init() {
	RegisterSchema(
		MessageID{Type: "TAXCON", Version: "D", Release: "20A", Agency: "UN"},
		taxconD20aSchema,
		"https://service.unece.org/trade/untdid/d20a/trmd/taxcon_c.htm",
	)
}

var taxconD20aSchema = Schema{Nodes: []SchemaNode{
	{Segment: "BGM", Mandatory: true, MaxRepeat: 1}, // Beginning of message
	{Segment: "DTM", Mandatory: true, MaxRepeat: 5}, // Date/time/period
	{ // Segment group 1
		Group: []SchemaNode{
			{Segment: "RFF", Mandatory: true, MaxRepeat: 1},  // Reference
			{Segment: "DTM", Mandatory: false, MaxRepeat: 5}, // Date/time/period
		},
		Mandatory: false, MaxRepeat: 5,
	},
	{ // Segment group 2
		Group: []SchemaNode{
			{Segment: "NAD", Mandatory: true, MaxRepeat: 1},  // Name and address
			{Segment: "RFF", Mandatory: false, MaxRepeat: 1}, // Reference
		},
		Mandatory: true, MaxRepeat: 5,
	},
	{Segment: "UNS", Mandatory: true, MaxRepeat: 1}, // Section control
	{ // Segment group 3
		Group: []SchemaNode{
			{Segment: "RFF", Mandatory: true, MaxRepeat: 1},  // Reference
			{Segment: "DTM", Mandatory: false, MaxRepeat: 5}, // Date/time/period
			{ // Segment group 4
				Group: []SchemaNode{
					{Segment: "DOC", Mandatory: true, MaxRepeat: 1},  // Document/message details
					{Segment: "MOA", Mandatory: false, MaxRepeat: 9}, // Monetary amount
					{Segment: "CNT", Mandatory: true, MaxRepeat: 1},  // Control total
					{ // Segment group 5
						Group: []SchemaNode{
							{Segment: "TAX", Mandatory: true, MaxRepeat: 1}, // Duty/tax/fee details
							{Segment: "MOA", Mandatory: true, MaxRepeat: 3}, // Monetary amount
						},
						Mandatory: true, MaxRepeat: 9,
					},
				},
				Mandatory: true, MaxRepeat: 9,
			},
		},
		Mandatory: true, MaxRepeat: 9999,
	},
}}
