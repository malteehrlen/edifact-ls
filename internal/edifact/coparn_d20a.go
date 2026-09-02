package edifact

// COPARN D.20A message specification: the structural (segment/group
// presence, order, cardinality) branching diagram for this message,
// UN/EDIFACT directory release D.20A. Scope is structural only -- not
// element-level content -- per edifact-ls-3uzr's epic non-goal.
// 20 segment groups, max nesting depth 3.
//
// Source: https://service.unece.org/trade/untdid/d20a/trmd/coparn_c.htm
// section 4.3.1 "Segment table". That URL currently returns HTTP 403 from
// Cloudflare when fetched directly; this data was transcribed from the
// Wayback Machine's archived copy instead:
// http://web.archive.org/web/20231201223122/https://service.unece.org/trade/untdid/d20a/trmd/coparn_c.htm
// -- re-check the direct URL first if cross-checking later, in case the
// block has lifted.
//
// Transcription note: generated in bulk (edifact-ls-13gu) the same way
// as every other message-type schema -- a script parsed the source's
// exact rail-art column positions mechanically and verified the result
// balances (all 20 groups open and close consistently)
// before being transcribed here, rather than reading the ASCII nesting
// by eye. Not individually spot-checked against the raw source the way
// the first several message types in this project were -- see
// edifact-ls-13gu for which representative samples were spot-checked
// for this batch.

func init() {
	RegisterSchema(
		MessageID{Type: "COPARN", Version: "D", Release: "20A", Agency: "UN"},
		coparnD20aSchema,
		"https://service.unece.org/trade/untdid/d20a/trmd/coparn_c.htm",
	)
}

var coparnD20aSchema = Schema{Nodes: []SchemaNode{
	{Segment: "BGM", Mandatory: true, MaxRepeat: 1},   // Beginning of message
	{Segment: "TMD", Mandatory: false, MaxRepeat: 1},  // Transport movement details
	{Segment: "DTM", Mandatory: false, MaxRepeat: 9},  // Date/time/period
	{Segment: "TSR", Mandatory: false, MaxRepeat: 9},  // Transport service requirements
	{Segment: "FTX", Mandatory: false, MaxRepeat: 99}, // Free text
	{Segment: "LOC", Mandatory: false, MaxRepeat: 9},  // Place/location identification
	{ // Segment group 1
		Group: []SchemaNode{
			{Segment: "RFF", Mandatory: true, MaxRepeat: 1},  // Reference
			{Segment: "DTM", Mandatory: false, MaxRepeat: 9}, // Date/time/period
		},
		Mandatory: false, MaxRepeat: 9,
	},
	{ // Segment group 2
		Group: []SchemaNode{
			{Segment: "TDT", Mandatory: true, MaxRepeat: 1},  // Transport information
			{Segment: "DTM", Mandatory: false, MaxRepeat: 9}, // Date/time/period
			{Segment: "RFF", Mandatory: false, MaxRepeat: 9}, // Reference
			{ // Segment group 3
				Group: []SchemaNode{
					{Segment: "LOC", Mandatory: true, MaxRepeat: 1},  // Place/location identification
					{Segment: "DTM", Mandatory: false, MaxRepeat: 9}, // Date/time/period
				},
				Mandatory: false, MaxRepeat: 9,
			},
		},
		Mandatory: false, MaxRepeat: 9,
	},
	{ // Segment group 4
		Group: []SchemaNode{
			{Segment: "NAD", Mandatory: true, MaxRepeat: 1}, // Name and address
			{ // Segment group 5
				Group: []SchemaNode{
					{Segment: "CTA", Mandatory: true, MaxRepeat: 1},  // Contact information
					{Segment: "COM", Mandatory: false, MaxRepeat: 9}, // Communication contact
				},
				Mandatory: false, MaxRepeat: 9,
			},
			{Segment: "RFF", Mandatory: false, MaxRepeat: 9}, // Reference
			{Segment: "DTM", Mandatory: false, MaxRepeat: 9}, // Date/time/period
		},
		Mandatory: true, MaxRepeat: 9,
	},
	{ // Segment group 6
		Group: []SchemaNode{
			{Segment: "GID", Mandatory: true, MaxRepeat: 1},  // Goods item details
			{Segment: "HAN", Mandatory: false, MaxRepeat: 9}, // Handling instructions
			{Segment: "FTX", Mandatory: false, MaxRepeat: 9}, // Free text
			{Segment: "RFF", Mandatory: false, MaxRepeat: 9}, // Reference
			{Segment: "PIA", Mandatory: false, MaxRepeat: 9}, // Additional product id
			{ // Segment group 7
				Group: []SchemaNode{
					{Segment: "NAD", Mandatory: true, MaxRepeat: 1},  // Name and address
					{Segment: "DTM", Mandatory: false, MaxRepeat: 9}, // Date/time/period
					{Segment: "RFF", Mandatory: false, MaxRepeat: 9}, // Reference
				},
				Mandatory: false, MaxRepeat: 99,
			},
			{Segment: "MEA", Mandatory: false, MaxRepeat: 9}, // Measurements
			{Segment: "DIM", Mandatory: false, MaxRepeat: 9}, // Dimensions
			{ // Segment group 8
				Group: []SchemaNode{
					{Segment: "DOC", Mandatory: true, MaxRepeat: 1},  // Document/message details
					{Segment: "DTM", Mandatory: false, MaxRepeat: 9}, // Date/time/period
					{Segment: "LOC", Mandatory: false, MaxRepeat: 9}, // Place/location identification
				},
				Mandatory: false, MaxRepeat: 9,
			},
			{ // Segment group 9
				Group: []SchemaNode{
					{Segment: "SGP", Mandatory: true, MaxRepeat: 1},  // Split goods placement
					{Segment: "MEA", Mandatory: false, MaxRepeat: 9}, // Measurements
				},
				Mandatory: false, MaxRepeat: 999,
			},
			{ // Segment group 10
				Group: []SchemaNode{
					{Segment: "DGS", Mandatory: true, MaxRepeat: 1},  // Dangerous goods
					{Segment: "ATT", Mandatory: false, MaxRepeat: 9}, // Attribute
					{Segment: "FTX", Mandatory: false, MaxRepeat: 9}, // Free text
					{Segment: "MEA", Mandatory: false, MaxRepeat: 9}, // Measurements
					{ // Segment group 11
						Group: []SchemaNode{
							{Segment: "CTA", Mandatory: true, MaxRepeat: 1},  // Contact information
							{Segment: "COM", Mandatory: false, MaxRepeat: 9}, // Communication contact
						},
						Mandatory: false, MaxRepeat: 9,
					},
				},
				Mandatory: false, MaxRepeat: 999,
			},
			{ // Segment group 12
				Group: []SchemaNode{
					{Segment: "TMP", Mandatory: true, MaxRepeat: 1},  // Temperature
					{Segment: "RNG", Mandatory: false, MaxRepeat: 1}, // Range details
				},
				Mandatory: false, MaxRepeat: 9,
			},
		},
		Mandatory: false, MaxRepeat: 99999,
	},
	{ // Segment group 13
		Group: []SchemaNode{
			{Segment: "EQD", Mandatory: true, MaxRepeat: 1},  // Equipment details
			{Segment: "RFF", Mandatory: false, MaxRepeat: 9}, // Reference
			{Segment: "EQN", Mandatory: false, MaxRepeat: 1}, // Number of units
			{Segment: "TMD", Mandatory: false, MaxRepeat: 9}, // Transport movement details
			{Segment: "DTM", Mandatory: false, MaxRepeat: 9}, // Date/time/period
			{Segment: "TSR", Mandatory: false, MaxRepeat: 9}, // Transport service requirements
			{Segment: "LOC", Mandatory: false, MaxRepeat: 9}, // Place/location identification
			{Segment: "MEA", Mandatory: false, MaxRepeat: 9}, // Measurements
			{Segment: "DIM", Mandatory: false, MaxRepeat: 9}, // Dimensions
			{ // Segment group 14
				Group: []SchemaNode{
					{Segment: "TMP", Mandatory: true, MaxRepeat: 1},  // Temperature
					{Segment: "RNG", Mandatory: false, MaxRepeat: 1}, // Range details
				},
				Mandatory: false, MaxRepeat: 9,
			},
			{Segment: "SEL", Mandatory: false, MaxRepeat: 9}, // Seal number
			{Segment: "FTX", Mandatory: false, MaxRepeat: 9}, // Free text
			{Segment: "PCD", Mandatory: false, MaxRepeat: 9}, // Percentage details
			{ // Segment group 15
				Group: []SchemaNode{
					{Segment: "DGS", Mandatory: true, MaxRepeat: 1},  // Dangerous goods
					{Segment: "ATT", Mandatory: false, MaxRepeat: 9}, // Attribute
					{Segment: "FTX", Mandatory: false, MaxRepeat: 9}, // Free text
					{Segment: "MEA", Mandatory: false, MaxRepeat: 9}, // Measurements
					{ // Segment group 16
						Group: []SchemaNode{
							{Segment: "CTA", Mandatory: true, MaxRepeat: 1},  // Contact information
							{Segment: "COM", Mandatory: false, MaxRepeat: 9}, // Communication contact
						},
						Mandatory: false, MaxRepeat: 9,
					},
				},
				Mandatory: false, MaxRepeat: 999,
			},
			{Segment: "MOA", Mandatory: false, MaxRepeat: 9}, // Monetary amount
			{Segment: "GOR", Mandatory: false, MaxRepeat: 9}, // Governmental requirements
			{Segment: "EQA", Mandatory: false, MaxRepeat: 1}, // Attached equipment
			{Segment: "COD", Mandatory: false, MaxRepeat: 1}, // Component details
			{Segment: "HAN", Mandatory: false, MaxRepeat: 9}, // Handling instructions
			{ // Segment group 17
				Group: []SchemaNode{
					{Segment: "DAM", Mandatory: true, MaxRepeat: 1},  // Damage
					{Segment: "COD", Mandatory: false, MaxRepeat: 1}, // Component details
				},
				Mandatory: false, MaxRepeat: 9,
			},
			{ // Segment group 18
				Group: []SchemaNode{
					{Segment: "TDT", Mandatory: true, MaxRepeat: 1},  // Transport information
					{Segment: "DTM", Mandatory: false, MaxRepeat: 9}, // Date/time/period
					{Segment: "RFF", Mandatory: false, MaxRepeat: 9}, // Reference
					{ // Segment group 19
						Group: []SchemaNode{
							{Segment: "LOC", Mandatory: true, MaxRepeat: 1},  // Place/location identification
							{Segment: "DTM", Mandatory: false, MaxRepeat: 9}, // Date/time/period
						},
						Mandatory: false, MaxRepeat: 9,
					},
				},
				Mandatory: false, MaxRepeat: 9,
			},
			{ // Segment group 20
				Group: []SchemaNode{
					{Segment: "NAD", Mandatory: true, MaxRepeat: 1},  // Name and address
					{Segment: "DTM", Mandatory: false, MaxRepeat: 1}, // Date/time/period
					{Segment: "CTA", Mandatory: false, MaxRepeat: 1}, // Contact information
					{Segment: "COM", Mandatory: false, MaxRepeat: 1}, // Communication contact
					{Segment: "RFF", Mandatory: false, MaxRepeat: 9}, // Reference
				},
				Mandatory: false, MaxRepeat: 9,
			},
		},
		Mandatory: true, MaxRepeat: 99999,
	},
	{Segment: "CNT", Mandatory: false, MaxRepeat: 1}, // Control total
}}
