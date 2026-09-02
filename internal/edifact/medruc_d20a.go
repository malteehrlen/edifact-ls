package edifact

// MEDRUC D.20A message specification: the structural (segment/group
// presence, order, cardinality) branching diagram for this message,
// UN/EDIFACT directory release D.20A. Scope is structural only -- not
// element-level content -- per edifact-ls-3uzr's epic non-goal.
// 26 segment groups, max nesting depth 4.
//
// Source: https://service.unece.org/trade/untdid/d20a/trmd/medruc_c.htm
// section 4.3.1 "Segment table". That URL currently returns HTTP 403 from
// Cloudflare when fetched directly; this data was transcribed from the
// Wayback Machine's archived copy instead:
// http://web.archive.org/web/20230608022219/https://service.unece.org/trade/untdid/d20a/trmd/medruc_c.htm
// -- re-check the direct URL first if cross-checking later, in case the
// block has lifted.
//
// Transcription note: generated in bulk (edifact-ls-13gu) the same way
// as every other message-type schema -- a script parsed the source's
// exact rail-art column positions mechanically and verified the result
// balances (all 26 groups open and close consistently)
// before being transcribed here, rather than reading the ASCII nesting
// by eye. Not individually spot-checked against the raw source the way
// the first several message types in this project were -- see
// edifact-ls-13gu for which representative samples were spot-checked
// for this batch.

func init() {
	RegisterSchema(
		MessageID{Type: "MEDRUC", Version: "D", Release: "20A", Agency: "UN"},
		medrucD20aSchema,
		"https://service.unece.org/trade/untdid/d20a/trmd/medruc_c.htm",
	)
}

var medrucD20aSchema = Schema{Nodes: []SchemaNode{
	{Segment: "BGM", Mandatory: true, MaxRepeat: 1},  // Beginning of message
	{Segment: "DTM", Mandatory: false, MaxRepeat: 9}, // Date/time/period
	{Segment: "RFF", Mandatory: false, MaxRepeat: 9}, // Reference
	{Segment: "FTX", Mandatory: false, MaxRepeat: 9}, // Free text
	{Segment: "CUX", Mandatory: false, MaxRepeat: 1}, // Currencies
	{Segment: "CNT", Mandatory: false, MaxRepeat: 9}, // Control total
	{ // Segment group 1
		Group: []SchemaNode{
			{Segment: "MOA", Mandatory: true, MaxRepeat: 1},  // Monetary amount
			{Segment: "TAX", Mandatory: false, MaxRepeat: 1}, // Duty/tax/fee details
		},
		Mandatory: false, MaxRepeat: 9,
	},
	{ // Segment group 2
		Group: []SchemaNode{
			{Segment: "PNA", Mandatory: true, MaxRepeat: 1},   // Party identification
			{Segment: "SEQ", Mandatory: false, MaxRepeat: 1},  // Sequence details
			{Segment: "ADR", Mandatory: false, MaxRepeat: 9},  // Address
			{Segment: "CTA", Mandatory: false, MaxRepeat: 9},  // Contact information
			{Segment: "COM", Mandatory: false, MaxRepeat: 9},  // Communication contact
			{Segment: "RFF", Mandatory: false, MaxRepeat: 99}, // Reference
			{Segment: "REL", Mandatory: false, MaxRepeat: 9},  // Relationship
			{Segment: "NAT", Mandatory: false, MaxRepeat: 9},  // Nationality
			{Segment: "LAN", Mandatory: false, MaxRepeat: 9},  // Language
			{Segment: "SPR", Mandatory: false, MaxRepeat: 1},  // Organisation classification details
			{Segment: "EMP", Mandatory: false, MaxRepeat: 9},  // Employment details
			{Segment: "QUA", Mandatory: false, MaxRepeat: 9},  // Qualification
			{Segment: "DTM", Mandatory: false, MaxRepeat: 99}, // Date/time/period
			{Segment: "STS", Mandatory: false, MaxRepeat: 1},  // Status
			{Segment: "RCS", Mandatory: false, MaxRepeat: 99}, // Requirements and conditions
			{Segment: "PDI", Mandatory: false, MaxRepeat: 1},  // Person demographic information
			{Segment: "FII", Mandatory: false, MaxRepeat: 9},  // Financial institution information
		},
		Mandatory: false, MaxRepeat: 999,
	},
	{ // Segment group 3
		Group: []SchemaNode{
			{Segment: "AGR", Mandatory: true, MaxRepeat: 1},   // Agreement identification
			{Segment: "SEQ", Mandatory: false, MaxRepeat: 1},  // Sequence details
			{Segment: "RFF", Mandatory: false, MaxRepeat: 99}, // Reference
			{Segment: "DTM", Mandatory: false, MaxRepeat: 99}, // Date/time/period
			{ // Segment group 4
				Group: []SchemaNode{
					{Segment: "ICD", Mandatory: true, MaxRepeat: 1},  // Insurance cover description
					{Segment: "FTX", Mandatory: false, MaxRepeat: 9}, // Free text
				},
				Mandatory: false, MaxRepeat: 99,
			},
		},
		Mandatory: false, MaxRepeat: 999,
	},
	{ // Segment group 5
		Group: []SchemaNode{
			{Segment: "DOC", Mandatory: true, MaxRepeat: 1},  // Document/message details
			{Segment: "DTM", Mandatory: false, MaxRepeat: 9}, // Date/time/period
			{ // Segment group 6
				Group: []SchemaNode{
					{Segment: "PNA", Mandatory: true, MaxRepeat: 1},  // Party identification
					{Segment: "ADR", Mandatory: false, MaxRepeat: 9}, // Address
					{Segment: "STS", Mandatory: false, MaxRepeat: 1}, // Status
					{Segment: "RCS", Mandatory: false, MaxRepeat: 9}, // Requirements and conditions
				},
				Mandatory: false, MaxRepeat: 9,
			},
		},
		Mandatory: false, MaxRepeat: 99,
	},
	{ // Segment group 7
		Group: []SchemaNode{
			{Segment: "FCA", Mandatory: true, MaxRepeat: 1},   // Financial charges allocation
			{Segment: "GEI", Mandatory: false, MaxRepeat: 9},  // Processing information
			{Segment: "RFF", Mandatory: false, MaxRepeat: 99}, // Reference
			{Segment: "DTM", Mandatory: false, MaxRepeat: 99}, // Date/time/period
			{Segment: "FTX", Mandatory: false, MaxRepeat: 99}, // Free text
			{Segment: "AGR", Mandatory: false, MaxRepeat: 99}, // Agreement identification
			{Segment: "DOC", Mandatory: false, MaxRepeat: 9},  // Document/message details
			{ // Segment group 8
				Group: []SchemaNode{
					{Segment: "MOA", Mandatory: true, MaxRepeat: 1},  // Monetary amount
					{Segment: "TAX", Mandatory: false, MaxRepeat: 1}, // Duty/tax/fee details
				},
				Mandatory: false, MaxRepeat: 99,
			},
			{ // Segment group 9
				Group: []SchemaNode{
					{Segment: "RCS", Mandatory: true, MaxRepeat: 1},  // Requirements and conditions
					{Segment: "FTX", Mandatory: false, MaxRepeat: 9}, // Free text
				},
				Mandatory: false, MaxRepeat: 99,
			},
			{ // Segment group 10
				Group: []SchemaNode{
					{Segment: "PNA", Mandatory: true, MaxRepeat: 1},   // Party identification
					{Segment: "RFF", Mandatory: false, MaxRepeat: 99}, // Reference
					{Segment: "ADR", Mandatory: false, MaxRepeat: 9},  // Address
					{Segment: "COM", Mandatory: false, MaxRepeat: 9},  // Communication contact
					{Segment: "DTM", Mandatory: false, MaxRepeat: 99}, // Date/time/period
					{Segment: "PDI", Mandatory: false, MaxRepeat: 1},  // Person demographic information
					{Segment: "RSL", Mandatory: false, MaxRepeat: 9},  // Result
					{Segment: "NAT", Mandatory: false, MaxRepeat: 9},  // Nationality
					{Segment: "LAN", Mandatory: false, MaxRepeat: 9},  // Language
					{Segment: "HAN", Mandatory: false, MaxRepeat: 9},  // Handling instructions
					{Segment: "REL", Mandatory: false, MaxRepeat: 1},  // Relationship
				},
				Mandatory: false, MaxRepeat: 999,
			},
			{ // Segment group 11
				Group: []SchemaNode{
					{Segment: "PRC", Mandatory: true, MaxRepeat: 1},   // Process identification
					{Segment: "GEI", Mandatory: false, MaxRepeat: 1},  // Processing information
					{Segment: "STS", Mandatory: false, MaxRepeat: 9},  // Status
					{Segment: "IMD", Mandatory: false, MaxRepeat: 9},  // Item description
					{Segment: "CIN", Mandatory: false, MaxRepeat: 99}, // Clinical information
					{Segment: "PNA", Mandatory: false, MaxRepeat: 99}, // Party identification
					{Segment: "DTM", Mandatory: false, MaxRepeat: 99}, // Date/time/period
					{Segment: "RFF", Mandatory: false, MaxRepeat: 99}, // Reference
					{Segment: "FTX", Mandatory: false, MaxRepeat: 99}, // Free text
					{Segment: "QTY", Mandatory: false, MaxRepeat: 99}, // Quantity
					{Segment: "PCD", Mandatory: false, MaxRepeat: 99}, // Percentage details
					{ // Segment group 12
						Group: []SchemaNode{
							{Segment: "PAS", Mandatory: true, MaxRepeat: 1},  // Attendance
							{Segment: "LOC", Mandatory: false, MaxRepeat: 9}, // Place/location identification
							{Segment: "DTM", Mandatory: false, MaxRepeat: 9}, // Date/time/period
							{Segment: "RFF", Mandatory: false, MaxRepeat: 9}, // Reference
						},
						Mandatory: false, MaxRepeat: 99,
					},
					{ // Segment group 13
						Group: []SchemaNode{
							{Segment: "MOA", Mandatory: true, MaxRepeat: 1},  // Monetary amount
							{Segment: "TAX", Mandatory: false, MaxRepeat: 1}, // Duty/tax/fee details
						},
						Mandatory: false, MaxRepeat: 99,
					},
					{ // Segment group 14
						Group: []SchemaNode{
							{Segment: "RCS", Mandatory: true, MaxRepeat: 1},  // Requirements and conditions
							{Segment: "FTX", Mandatory: false, MaxRepeat: 9}, // Free text
						},
						Mandatory: false, MaxRepeat: 99,
					},
					{ // Segment group 15
						Group: []SchemaNode{
							{Segment: "CLI", Mandatory: true, MaxRepeat: 1},   // Clinical intervention
							{Segment: "GEI", Mandatory: false, MaxRepeat: 1},  // Processing information
							{Segment: "IMD", Mandatory: false, MaxRepeat: 9},  // Item description
							{Segment: "CIN", Mandatory: false, MaxRepeat: 99}, // Clinical information
							{Segment: "PNA", Mandatory: false, MaxRepeat: 99}, // Party identification
							{Segment: "DTM", Mandatory: false, MaxRepeat: 99}, // Date/time/period
							{Segment: "RFF", Mandatory: false, MaxRepeat: 99}, // Reference
							{Segment: "FTX", Mandatory: false, MaxRepeat: 99}, // Free text
							{Segment: "QTY", Mandatory: false, MaxRepeat: 99}, // Quantity
							{Segment: "PCD", Mandatory: false, MaxRepeat: 99}, // Percentage details
							{Segment: "PAC", Mandatory: false, MaxRepeat: 9},  // Package
							{ // Segment group 16
								Group: []SchemaNode{
									{Segment: "MOA", Mandatory: true, MaxRepeat: 1},  // Monetary amount
									{Segment: "TAX", Mandatory: false, MaxRepeat: 1}, // Duty/tax/fee details
								},
								Mandatory: false, MaxRepeat: 99,
							},
							{ // Segment group 17
								Group: []SchemaNode{
									{Segment: "RCS", Mandatory: true, MaxRepeat: 1},  // Requirements and conditions
									{Segment: "FTX", Mandatory: false, MaxRepeat: 9}, // Free text
								},
								Mandatory: false, MaxRepeat: 99,
							},
							{ // Segment group 18
								Group: []SchemaNode{
									{Segment: "EQD", Mandatory: true, MaxRepeat: 1},  // Equipment details
									{Segment: "SEQ", Mandatory: false, MaxRepeat: 1}, // Sequence details
								},
								Mandatory: false, MaxRepeat: 99,
							},
						},
						Mandatory: false, MaxRepeat: 999,
					},
				},
				Mandatory: false, MaxRepeat: 999,
			},
		},
		Mandatory: false, MaxRepeat: 999,
	},
	{ // Segment group 19
		Group: []SchemaNode{
			{Segment: "TMD", Mandatory: true, MaxRepeat: 1},   // Transport movement details
			{Segment: "SEQ", Mandatory: false, MaxRepeat: 1},  // Sequence details
			{Segment: "GEI", Mandatory: false, MaxRepeat: 1},  // Processing information
			{Segment: "FTX", Mandatory: false, MaxRepeat: 99}, // Free text
			{ // Segment group 20
				Group: []SchemaNode{
					{Segment: "TDT", Mandatory: true, MaxRepeat: 1},   // Transport information
					{Segment: "GEI", Mandatory: false, MaxRepeat: 1},  // Processing information
					{Segment: "IMD", Mandatory: false, MaxRepeat: 9},  // Item description
					{Segment: "PNA", Mandatory: false, MaxRepeat: 99}, // Party identification
					{Segment: "RFF", Mandatory: false, MaxRepeat: 99}, // Reference
					{Segment: "FTX", Mandatory: false, MaxRepeat: 99}, // Free text
					{Segment: "CIN", Mandatory: false, MaxRepeat: 99}, // Clinical information
					{Segment: "QTY", Mandatory: false, MaxRepeat: 99}, // Quantity
					{ // Segment group 21
						Group: []SchemaNode{
							{Segment: "ADR", Mandatory: true, MaxRepeat: 1},  // Address
							{Segment: "DTM", Mandatory: false, MaxRepeat: 9}, // Date/time/period
						},
						Mandatory: false, MaxRepeat: 99,
					},
					{ // Segment group 22
						Group: []SchemaNode{
							{Segment: "TSR", Mandatory: true, MaxRepeat: 1},  // Transport service requirements
							{Segment: "FTX", Mandatory: false, MaxRepeat: 9}, // Free text
						},
						Mandatory: false, MaxRepeat: 99,
					},
					{ // Segment group 23
						Group: []SchemaNode{
							{Segment: "TCC", Mandatory: true, MaxRepeat: 1},   // Charge/rate calculations
							{Segment: "GEI", Mandatory: false, MaxRepeat: 1},  // Processing information
							{Segment: "IMD", Mandatory: false, MaxRepeat: 9},  // Item description
							{Segment: "DTM", Mandatory: false, MaxRepeat: 99}, // Date/time/period
							{Segment: "RFF", Mandatory: false, MaxRepeat: 99}, // Reference
							{Segment: "FTX", Mandatory: false, MaxRepeat: 99}, // Free text
							{Segment: "QTY", Mandatory: false, MaxRepeat: 99}, // Quantity
							{Segment: "PCD", Mandatory: false, MaxRepeat: 99}, // Percentage details
							{ // Segment group 24
								Group: []SchemaNode{
									{Segment: "MOA", Mandatory: true, MaxRepeat: 1},  // Monetary amount
									{Segment: "TAX", Mandatory: false, MaxRepeat: 1}, // Duty/tax/fee details
								},
								Mandatory: false, MaxRepeat: 99,
							},
							{ // Segment group 25
								Group: []SchemaNode{
									{Segment: "TSR", Mandatory: true, MaxRepeat: 1},  // Transport service requirements
									{Segment: "FTX", Mandatory: false, MaxRepeat: 9}, // Free text
								},
								Mandatory: false, MaxRepeat: 99,
							},
						},
						Mandatory: false, MaxRepeat: 999,
					},
				},
				Mandatory: false, MaxRepeat: 999,
			},
		},
		Mandatory: false, MaxRepeat: 999,
	},
	{ // Segment group 26
		Group: []SchemaNode{
			{Segment: "AUT", Mandatory: true, MaxRepeat: 1},  // Authentication result
			{Segment: "DTM", Mandatory: false, MaxRepeat: 9}, // Date/time/period
		},
		Mandatory: false, MaxRepeat: 9,
	},
}}
