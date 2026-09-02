package edifact

// JUPREQ D.20A message specification: the structural (segment/group
// presence, order, cardinality) branching diagram for this message,
// UN/EDIFACT directory release D.20A. Scope is structural only -- not
// element-level content -- per edifact-ls-3uzr's epic non-goal.
// 16 segment groups, max nesting depth 4.
//
// Source: https://service.unece.org/trade/untdid/d20a/trmd/jupreq_c.htm
// section 4.3.1 "Segment table". That URL currently returns HTTP 403 from
// Cloudflare when fetched directly; this data was transcribed from the
// Wayback Machine's archived copy instead:
// http://web.archive.org/web/20231208133323/https://service.unece.org/trade/untdid/d20a/trmd/jupreq_c.htm
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
		MessageID{Type: "JUPREQ", Version: "D", Release: "20A", Agency: "UN"},
		jupreqD20aSchema,
		"https://service.unece.org/trade/untdid/d20a/trmd/jupreq_c.htm",
	)
}

var jupreqD20aSchema = Schema{Nodes: []SchemaNode{
	{Segment: "BGM", Mandatory: true, MaxRepeat: 1},   // Beginning of message
	{Segment: "DTM", Mandatory: true, MaxRepeat: 99},  // Date/time/period
	{Segment: "MOA", Mandatory: false, MaxRepeat: 99}, // Monetary amount
	{Segment: "GEI", Mandatory: false, MaxRepeat: 99}, // Processing information
	{Segment: "FTX", Mandatory: false, MaxRepeat: 99}, // Free text
	{ // Segment group 1
		Group: []SchemaNode{
			{Segment: "RFF", Mandatory: true, MaxRepeat: 1},   // Reference
			{Segment: "DTM", Mandatory: false, MaxRepeat: 99}, // Date/time/period
		},
		Mandatory: false, MaxRepeat: 99,
	},
	{ // Segment group 2
		Group: []SchemaNode{
			{Segment: "NAD", Mandatory: true, MaxRepeat: 1},   // Name and address
			{Segment: "CTA", Mandatory: false, MaxRepeat: 99}, // Contact information
			{Segment: "DTM", Mandatory: false, MaxRepeat: 99}, // Date/time/period
		},
		Mandatory: true, MaxRepeat: 99,
	},
	{ // Segment group 3
		Group: []SchemaNode{
			{Segment: "DOC", Mandatory: true, MaxRepeat: 1},   // Document/message details
			{Segment: "GEI", Mandatory: false, MaxRepeat: 99}, // Processing information
			{Segment: "RFF", Mandatory: false, MaxRepeat: 99}, // Reference
			{Segment: "DTM", Mandatory: false, MaxRepeat: 99}, // Date/time/period
			{Segment: "MOA", Mandatory: false, MaxRepeat: 99}, // Monetary amount
			{Segment: "FTX", Mandatory: false, MaxRepeat: 99}, // Free text
			{ // Segment group 4
				Group: []SchemaNode{
					{Segment: "NAD", Mandatory: true, MaxRepeat: 1},   // Name and address
					{Segment: "FII", Mandatory: false, MaxRepeat: 99}, // Financial institution information
					{Segment: "MOA", Mandatory: false, MaxRepeat: 99}, // Monetary amount
					{ // Segment group 5
						Group: []SchemaNode{
							{Segment: "UGH", Mandatory: true, MaxRepeat: 1}, // Anti-collision segment group header
							{ // Segment group 6
								Group: []SchemaNode{
									{Segment: "REL", Mandatory: true, MaxRepeat: 1},   // Relationship
									{Segment: "NAD", Mandatory: false, MaxRepeat: 99}, // Name and address
									{Segment: "FII", Mandatory: false, MaxRepeat: 99}, // Financial institution information
								},
								Mandatory: true, MaxRepeat: 99,
							},
							{Segment: "UGT", Mandatory: true, MaxRepeat: 1}, // Anti-collision segment group trailer
						},
						Mandatory: true, MaxRepeat: 1,
					},
				},
				Mandatory: true, MaxRepeat: 99,
			},
			{ // Segment group 7
				Group: []SchemaNode{
					{Segment: "DMS", Mandatory: true, MaxRepeat: 1},   // Document/message summary
					{Segment: "RFF", Mandatory: false, MaxRepeat: 99}, // Reference
					{Segment: "NAD", Mandatory: false, MaxRepeat: 99}, // Name and address
				},
				Mandatory: true, MaxRepeat: 99,
			},
			{ // Segment group 8
				Group: []SchemaNode{
					{Segment: "TAX", Mandatory: true, MaxRepeat: 1},   // Duty/tax/fee details
					{Segment: "MOA", Mandatory: false, MaxRepeat: 99}, // Monetary amount
				},
				Mandatory: true, MaxRepeat: 99,
			},
			{ // Segment group 9
				Group: []SchemaNode{
					{Segment: "FOR", Mandatory: true, MaxRepeat: 1},   // Formula
					{Segment: "GEI", Mandatory: false, MaxRepeat: 99}, // Processing information
					{Segment: "DTM", Mandatory: false, MaxRepeat: 99}, // Date/time/period
					{Segment: "RTE", Mandatory: false, MaxRepeat: 99}, // Rate details
					{Segment: "MOA", Mandatory: false, MaxRepeat: 99}, // Monetary amount
					{Segment: "TAX", Mandatory: false, MaxRepeat: 99}, // Duty/tax/fee details
					{ // Segment group 10
						Group: []SchemaNode{
							{Segment: "IND", Mandatory: true, MaxRepeat: 1},   // Index details
							{Segment: "DTM", Mandatory: false, MaxRepeat: 99}, // Date/time/period
							{Segment: "RTE", Mandatory: false, MaxRepeat: 99}, // Rate details
						},
						Mandatory: false, MaxRepeat: 99,
					},
				},
				Mandatory: false, MaxRepeat: 99,
			},
			{ // Segment group 11
				Group: []SchemaNode{
					{Segment: "LIN", Mandatory: true, MaxRepeat: 1},   // Line item
					{Segment: "PIA", Mandatory: false, MaxRepeat: 99}, // Additional product id
					{Segment: "QTY", Mandatory: false, MaxRepeat: 99}, // Quantity
					{Segment: "DTM", Mandatory: false, MaxRepeat: 99}, // Date/time/period
					{Segment: "TAX", Mandatory: false, MaxRepeat: 99}, // Duty/tax/fee details
					{Segment: "RTE", Mandatory: false, MaxRepeat: 99}, // Rate details
					{Segment: "DMS", Mandatory: false, MaxRepeat: 99}, // Document/message summary
					{Segment: "RFF", Mandatory: false, MaxRepeat: 99}, // Reference
					{Segment: "GEI", Mandatory: false, MaxRepeat: 99}, // Processing information
					{ // Segment group 12
						Group: []SchemaNode{
							{Segment: "MOA", Mandatory: true, MaxRepeat: 1},   // Monetary amount
							{Segment: "CUX", Mandatory: false, MaxRepeat: 99}, // Currencies
							{Segment: "PAI", Mandatory: false, MaxRepeat: 99}, // Payment instructions
						},
						Mandatory: false, MaxRepeat: 99,
					},
					{ // Segment group 13
						Group: []SchemaNode{
							{Segment: "NAD", Mandatory: true, MaxRepeat: 1},   // Name and address
							{Segment: "FII", Mandatory: false, MaxRepeat: 99}, // Financial institution information
						},
						Mandatory: false, MaxRepeat: 99,
					},
					{ // Segment group 14
						Group: []SchemaNode{
							{Segment: "FOR", Mandatory: true, MaxRepeat: 1},   // Formula
							{Segment: "GEI", Mandatory: false, MaxRepeat: 99}, // Processing information
							{Segment: "DTM", Mandatory: false, MaxRepeat: 99}, // Date/time/period
							{Segment: "MOA", Mandatory: false, MaxRepeat: 99}, // Monetary amount
						},
						Mandatory: false, MaxRepeat: 99,
					},
					{ // Segment group 15
						Group: []SchemaNode{
							{Segment: "ARD", Mandatory: true, MaxRepeat: 1},   // Monetary amount function
							{Segment: "MOA", Mandatory: false, MaxRepeat: 99}, // Monetary amount
							{Segment: "DTM", Mandatory: false, MaxRepeat: 99}, // Date/time/period
							{Segment: "TAX", Mandatory: false, MaxRepeat: 99}, // Duty/tax/fee details
							{ // Segment group 16
								Group: []SchemaNode{
									{Segment: "SEQ", Mandatory: true, MaxRepeat: 1},   // Sequence details
									{Segment: "GEI", Mandatory: false, MaxRepeat: 99}, // Processing information
									{Segment: "DTM", Mandatory: false, MaxRepeat: 99}, // Date/time/period
									{Segment: "MOA", Mandatory: false, MaxRepeat: 99}, // Monetary amount
									{Segment: "RFF", Mandatory: false, MaxRepeat: 99}, // Reference
								},
								Mandatory: false, MaxRepeat: 99,
							},
						},
						Mandatory: true, MaxRepeat: 99,
					},
				},
				Mandatory: false, MaxRepeat: 999,
			},
		},
		Mandatory: false, MaxRepeat: 999,
	},
}}
