package edifact

// RECORD D.20A message specification: the structural (segment/group
// presence, order, cardinality) branching diagram for this message,
// UN/EDIFACT directory release D.20A. Scope is structural only -- not
// element-level content -- per edifact-ls-3uzr's epic non-goal.
// 17 segment groups, max nesting depth 4.
//
// Source: https://service.unece.org/trade/untdid/d20a/trmd/record_c.htm
// section 4.3.1 "Segment table". That URL currently returns HTTP 403 from
// Cloudflare when fetched directly; this data was transcribed from the
// Wayback Machine's archived copy instead:
// http://web.archive.org/web/20231202084057/https://service.unece.org/trade/untdid/d20a/trmd/record_c.htm
// -- re-check the direct URL first if cross-checking later, in case the
// block has lifted.
//
// Transcription note: generated in bulk (edifact-ls-13gu) the same way
// as every other message-type schema -- a script parsed the source's
// exact rail-art column positions mechanically and verified the result
// balances (all 17 groups open and close consistently)
// before being transcribed here, rather than reading the ASCII nesting
// by eye. Not individually spot-checked against the raw source the way
// the first several message types in this project were -- see
// edifact-ls-13gu for which representative samples were spot-checked
// for this batch.

func init() {
	RegisterSchema(
		MessageID{Type: "RECORD", Version: "D", Release: "20A", Agency: "UN"},
		recordD20aSchema,
		"https://service.unece.org/trade/untdid/d20a/trmd/record_c.htm",
	)
}

var recordD20aSchema = Schema{Nodes: []SchemaNode{
	{Segment: "BGM", Mandatory: true, MaxRepeat: 1},  // Beginning of message
	{Segment: "DTM", Mandatory: true, MaxRepeat: 9},  // Date/time/period
	{Segment: "AGR", Mandatory: true, MaxRepeat: 9},  // Agreement identification
	{Segment: "RFF", Mandatory: true, MaxRepeat: 9},  // Reference
	{Segment: "GEI", Mandatory: false, MaxRepeat: 9}, // Processing information
	{Segment: "FTX", Mandatory: false, MaxRepeat: 9}, // Free text
	{ // Segment group 1
		Group: []SchemaNode{
			{Segment: "PNA", Mandatory: true, MaxRepeat: 1},  // Party identification
			{Segment: "BUS", Mandatory: false, MaxRepeat: 1}, // Business function
			{Segment: "RFF", Mandatory: false, MaxRepeat: 9}, // Reference
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
			{Segment: "IDE", Mandatory: true, MaxRepeat: 1},   // Identity
			{Segment: "GEI", Mandatory: false, MaxRepeat: 9},  // Processing information
			{Segment: "ICD", Mandatory: false, MaxRepeat: 1},  // Insurance cover description
			{Segment: "RFF", Mandatory: false, MaxRepeat: 9},  // Reference
			{Segment: "ATT", Mandatory: false, MaxRepeat: 99}, // Attribute
			{Segment: "CUX", Mandatory: false, MaxRepeat: 9},  // Currencies
			{Segment: "FTX", Mandatory: false, MaxRepeat: 9},  // Free text
			{ // Segment group 4
				Group: []SchemaNode{
					{Segment: "PER", Mandatory: true, MaxRepeat: 1}, // Period related details
					{ // Segment group 5
						Group: []SchemaNode{
							{Segment: "DTM", Mandatory: true, MaxRepeat: 1},  // Date/time/period
							{Segment: "GEI", Mandatory: false, MaxRepeat: 1}, // Processing information
						},
						Mandatory: false, MaxRepeat: 2,
					},
				},
				Mandatory: false, MaxRepeat: 9,
			},
			{ // Segment group 6
				Group: []SchemaNode{
					{Segment: "LOC", Mandatory: true, MaxRepeat: 1},  // Place/location identification
					{Segment: "MEA", Mandatory: false, MaxRepeat: 2}, // Measurements
					{Segment: "ADR", Mandatory: false, MaxRepeat: 1}, // Address
				},
				Mandatory: false, MaxRepeat: 99,
			},
			{ // Segment group 7
				Group: []SchemaNode{
					{Segment: "CLA", Mandatory: true, MaxRepeat: 1},  // Clause identification
					{Segment: "APP", Mandatory: false, MaxRepeat: 1}, // Applicability
					{Segment: "MOA", Mandatory: false, MaxRepeat: 9}, // Monetary amount
					{Segment: "PCD", Mandatory: false, MaxRepeat: 9}, // Percentage details
					{Segment: "DTM", Mandatory: false, MaxRepeat: 9}, // Date/time/period
					{Segment: "GEI", Mandatory: false, MaxRepeat: 1}, // Processing information
					{Segment: "FTX", Mandatory: false, MaxRepeat: 1}, // Free text
				},
				Mandatory: false, MaxRepeat: 99,
			},
			{ // Segment group 8
				Group: []SchemaNode{
					{Segment: "PRV", Mandatory: true, MaxRepeat: 1},   // Proviso details
					{Segment: "APP", Mandatory: false, MaxRepeat: 9},  // Applicability
					{Segment: "MOA", Mandatory: false, MaxRepeat: 9},  // Monetary amount
					{Segment: "RTE", Mandatory: false, MaxRepeat: 9},  // Rate details
					{Segment: "DTM", Mandatory: false, MaxRepeat: 9},  // Date/time/period
					{Segment: "QTY", Mandatory: false, MaxRepeat: 9},  // Quantity
					{Segment: "FOR", Mandatory: false, MaxRepeat: 1},  // Formula
					{Segment: "CUX", Mandatory: false, MaxRepeat: 9},  // Currencies
					{Segment: "CLA", Mandatory: false, MaxRepeat: 99}, // Clause identification
					{Segment: "GEI", Mandatory: false, MaxRepeat: 9},  // Processing information
					{Segment: "FTX", Mandatory: false, MaxRepeat: 9},  // Free text
					{ // Segment group 9
						Group: []SchemaNode{
							{Segment: "BAS", Mandatory: true, MaxRepeat: 1},  // Basis
							{Segment: "APP", Mandatory: false, MaxRepeat: 1}, // Applicability
						},
						Mandatory: false, MaxRepeat: 9,
					},
					{ // Segment group 10
						Group: []SchemaNode{
							{Segment: "ATT", Mandatory: true, MaxRepeat: 1},  // Attribute
							{Segment: "PCD", Mandatory: false, MaxRepeat: 1}, // Percentage details
						},
						Mandatory: false, MaxRepeat: 99,
					},
					{ // Segment group 11
						Group: []SchemaNode{
							{Segment: "LOC", Mandatory: true, MaxRepeat: 1},  // Place/location identification
							{Segment: "MEA", Mandatory: false, MaxRepeat: 2}, // Measurements
							{Segment: "ADR", Mandatory: false, MaxRepeat: 1}, // Address
						},
						Mandatory: false, MaxRepeat: 99,
					},
					{ // Segment group 12
						Group: []SchemaNode{
							{Segment: "PYT", Mandatory: true, MaxRepeat: 1},  // Payment terms
							{Segment: "PCD", Mandatory: false, MaxRepeat: 1}, // Percentage details
						},
						Mandatory: false, MaxRepeat: 9,
					},
					{ // Segment group 13
						Group: []SchemaNode{
							{Segment: "RFF", Mandatory: true, MaxRepeat: 1},  // Reference
							{Segment: "MOA", Mandatory: false, MaxRepeat: 1}, // Monetary amount
							{Segment: "PCD", Mandatory: false, MaxRepeat: 9}, // Percentage details
							{Segment: "DTM", Mandatory: false, MaxRepeat: 9}, // Date/time/period
						},
						Mandatory: true, MaxRepeat: 99,
					},
				},
				Mandatory: false, MaxRepeat: 999,
			},
			{ // Segment group 14
				Group: []SchemaNode{
					{Segment: "PNA", Mandatory: true, MaxRepeat: 1},  // Party identification
					{Segment: "ATT", Mandatory: true, MaxRepeat: 9},  // Attribute
					{Segment: "ADR", Mandatory: false, MaxRepeat: 9}, // Address
					{Segment: "CTA", Mandatory: false, MaxRepeat: 1}, // Contact information
					{Segment: "COM", Mandatory: false, MaxRepeat: 9}, // Communication contact
					{Segment: "RFF", Mandatory: false, MaxRepeat: 9}, // Reference
					{ // Segment group 15
						Group: []SchemaNode{
							{Segment: "STS", Mandatory: true, MaxRepeat: 1},  // Status
							{Segment: "RFF", Mandatory: false, MaxRepeat: 9}, // Reference
							{Segment: "DTM", Mandatory: false, MaxRepeat: 9}, // Date/time/period
							{Segment: "MOA", Mandatory: false, MaxRepeat: 9}, // Monetary amount
							{Segment: "PCD", Mandatory: false, MaxRepeat: 9}, // Percentage details
							{Segment: "QTY", Mandatory: false, MaxRepeat: 9}, // Quantity
							{Segment: "FTX", Mandatory: false, MaxRepeat: 9}, // Free text
							{ // Segment group 16
								Group: []SchemaNode{
									{Segment: "RCS", Mandatory: true, MaxRepeat: 1},  // Requirements and conditions
									{Segment: "FTX", Mandatory: false, MaxRepeat: 9}, // Free text
									{Segment: "ATT", Mandatory: false, MaxRepeat: 1}, // Attribute
									{Segment: "RFF", Mandatory: false, MaxRepeat: 1}, // Reference
									{Segment: "PCD", Mandatory: false, MaxRepeat: 1}, // Percentage details
								},
								Mandatory: false, MaxRepeat: 999,
							},
						},
						Mandatory: false, MaxRepeat: 99,
					},
				},
				Mandatory: false, MaxRepeat: 999,
			},
		},
		Mandatory: false, MaxRepeat: 99,
	},
	{ // Segment group 17
		Group: []SchemaNode{
			{Segment: "DOC", Mandatory: true, MaxRepeat: 1},  // Document/message details
			{Segment: "RFF", Mandatory: false, MaxRepeat: 9}, // Reference
			{Segment: "IDE", Mandatory: false, MaxRepeat: 1}, // Identity
			{Segment: "COM", Mandatory: false, MaxRepeat: 1}, // Communication contact
			{Segment: "EFI", Mandatory: false, MaxRepeat: 1}, // External file link identification
			{Segment: "PNA", Mandatory: false, MaxRepeat: 1}, // Party identification
			{Segment: "DTM", Mandatory: false, MaxRepeat: 9}, // Date/time/period
		},
		Mandatory: false, MaxRepeat: 99,
	},
}}
