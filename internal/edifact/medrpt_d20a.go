package edifact

// MEDRPT D.20A message specification: the structural (segment/group
// presence, order, cardinality) branching diagram for this message,
// UN/EDIFACT directory release D.20A. Scope is structural only -- not
// element-level content -- per edifact-ls-3uzr's epic non-goal.
// 19 segment groups, max nesting depth 5.
//
// Source: https://service.unece.org/trade/untdid/d20a/trmd/medrpt_c.htm
// section 4.3.1 "Segment table". That URL currently returns HTTP 403 from
// Cloudflare when fetched directly; this data was transcribed from the
// Wayback Machine's archived copy instead:
// http://web.archive.org/web/20230608030043/https://service.unece.org/trade/untdid/d20a/trmd/medrpt_c.htm
// -- re-check the direct URL first if cross-checking later, in case the
// block has lifted.
//
// Transcription note: generated in bulk (edifact-ls-13gu) the same way
// as every other message-type schema -- a script parsed the source's
// exact rail-art column positions mechanically and verified the result
// balances (all 19 groups open and close consistently)
// before being transcribed here, rather than reading the ASCII nesting
// by eye. Not individually spot-checked against the raw source the way
// the first several message types in this project were -- see
// edifact-ls-13gu for which representative samples were spot-checked
// for this batch.

func init() {
	RegisterSchema(
		MessageID{Type: "MEDRPT", Version: "D", Release: "20A", Agency: "UN"},
		medrptD20aSchema,
		"https://service.unece.org/trade/untdid/d20a/trmd/medrpt_c.htm",
	)
}

var medrptD20aSchema = Schema{Nodes: []SchemaNode{
	{Segment: "BGM", Mandatory: true, MaxRepeat: 1},  // Beginning of message
	{Segment: "DTM", Mandatory: false, MaxRepeat: 9}, // Date/time/period
	{Segment: "FTX", Mandatory: false, MaxRepeat: 9}, // Free text
	{ // Segment group 1
		Group: []SchemaNode{
			{Segment: "PNA", Mandatory: true, MaxRepeat: 1},  // Party identification
			{Segment: "ADR", Mandatory: false, MaxRepeat: 9}, // Address
			{Segment: "COM", Mandatory: false, MaxRepeat: 9}, // Communication contact
			{Segment: "FTX", Mandatory: false, MaxRepeat: 9}, // Free text
			{Segment: "RFF", Mandatory: false, MaxRepeat: 9}, // Reference
			{Segment: "SEQ", Mandatory: false, MaxRepeat: 1}, // Sequence details
			{Segment: "LAN", Mandatory: false, MaxRepeat: 9}, // Language
			{Segment: "SPR", Mandatory: false, MaxRepeat: 1}, // Organisation classification details
			{Segment: "QUA", Mandatory: false, MaxRepeat: 9}, // Qualification
			{Segment: "EMP", Mandatory: false, MaxRepeat: 9}, // Employment details
		},
		Mandatory: true, MaxRepeat: 99,
	},
	{ // Segment group 2
		Group: []SchemaNode{
			{Segment: "IRQ", Mandatory: true, MaxRepeat: 1},   // Information required
			{Segment: "GEI", Mandatory: true, MaxRepeat: 1},   // Processing information
			{Segment: "RFF", Mandatory: true, MaxRepeat: 9},   // Reference
			{Segment: "DTM", Mandatory: true, MaxRepeat: 1},   // Date/time/period
			{Segment: "STS", Mandatory: false, MaxRepeat: 1},  // Status
			{Segment: "PTY", Mandatory: false, MaxRepeat: 1},  // Priority
			{Segment: "LAN", Mandatory: false, MaxRepeat: 1},  // Language
			{Segment: "FTX", Mandatory: false, MaxRepeat: 99}, // Free text
			{Segment: "TEM", Mandatory: false, MaxRepeat: 1},  // Test method
			{ // Segment group 3
				Group: []SchemaNode{
					{Segment: "RFF", Mandatory: true, MaxRepeat: 1},  // Reference
					{Segment: "DTM", Mandatory: false, MaxRepeat: 1}, // Date/time/period
				},
				Mandatory: false, MaxRepeat: 99,
			},
			{ // Segment group 4
				Group: []SchemaNode{
					{Segment: "FCA", Mandatory: true, MaxRepeat: 1},  // Financial charges allocation
					{Segment: "RFF", Mandatory: true, MaxRepeat: 9},  // Reference
					{Segment: "GEI", Mandatory: false, MaxRepeat: 1}, // Processing information
					{Segment: "DTM", Mandatory: false, MaxRepeat: 9}, // Date/time/period
					{Segment: "PTY", Mandatory: false, MaxRepeat: 1}, // Priority
					{Segment: "CIN", Mandatory: false, MaxRepeat: 9}, // Clinical information
					{Segment: "FTX", Mandatory: false, MaxRepeat: 9}, // Free text
					{Segment: "TEM", Mandatory: false, MaxRepeat: 9}, // Test method
					{ // Segment group 5
						Group: []SchemaNode{
							{Segment: "DOC", Mandatory: true, MaxRepeat: 1},  // Document/message details
							{Segment: "RFF", Mandatory: true, MaxRepeat: 9},  // Reference
							{Segment: "DTM", Mandatory: false, MaxRepeat: 1}, // Date/time/period
						},
						Mandatory: false, MaxRepeat: 99,
					},
				},
				Mandatory: false, MaxRepeat: 99,
			},
			{ // Segment group 6
				Group: []SchemaNode{
					{Segment: "ATT", Mandatory: true, MaxRepeat: 1},  // Attribute
					{Segment: "RFF", Mandatory: false, MaxRepeat: 9}, // Reference
					{Segment: "ADR", Mandatory: false, MaxRepeat: 9}, // Address
					{Segment: "COM", Mandatory: false, MaxRepeat: 9}, // Communication contact
					{Segment: "REL", Mandatory: false, MaxRepeat: 1}, // Relationship
					{Segment: "IMD", Mandatory: false, MaxRepeat: 9}, // Item description
					{Segment: "FTX", Mandatory: false, MaxRepeat: 9}, // Free text
					{Segment: "PNA", Mandatory: false, MaxRepeat: 9}, // Party identification
					{Segment: "DTM", Mandatory: false, MaxRepeat: 9}, // Date/time/period
					{Segment: "PDI", Mandatory: false, MaxRepeat: 1}, // Person demographic information
					{Segment: "NAT", Mandatory: false, MaxRepeat: 9}, // Nationality
					{Segment: "LAN", Mandatory: false, MaxRepeat: 9}, // Language
					{Segment: "HAN", Mandatory: false, MaxRepeat: 9}, // Handling instructions
					{Segment: "CCI", Mandatory: false, MaxRepeat: 9}, // Characteristic/class id
					{ // Segment group 7
						Group: []SchemaNode{
							{Segment: "PAS", Mandatory: true, MaxRepeat: 1},  // Attendance
							{Segment: "LOC", Mandatory: false, MaxRepeat: 9}, // Place/location identification
							{Segment: "DTM", Mandatory: false, MaxRepeat: 9}, // Date/time/period
							{Segment: "RFF", Mandatory: false, MaxRepeat: 1}, // Reference
						},
						Mandatory: false, MaxRepeat: 9,
					},
					{ // Segment group 8
						Group: []SchemaNode{
							{Segment: "CAV", Mandatory: true, MaxRepeat: 1},   // Characteristic value
							{Segment: "CIN", Mandatory: false, MaxRepeat: 9},  // Clinical information
							{Segment: "LAN", Mandatory: false, MaxRepeat: 9},  // Language
							{Segment: "DTM", Mandatory: false, MaxRepeat: 9},  // Date/time/period
							{Segment: "FTX", Mandatory: false, MaxRepeat: 99}, // Free text
							{Segment: "RFF", Mandatory: false, MaxRepeat: 9},  // Reference
							{ // Segment group 9
								Group: []SchemaNode{
									{Segment: "STS", Mandatory: true, MaxRepeat: 1},   // Status
									{Segment: "SEQ", Mandatory: false, MaxRepeat: 1},  // Sequence details
									{Segment: "GEI", Mandatory: false, MaxRepeat: 1},  // Processing information
									{Segment: "RSL", Mandatory: false, MaxRepeat: 1},  // Result
									{Segment: "CCI", Mandatory: true, MaxRepeat: 9},   // Characteristic/class id
									{Segment: "CIN", Mandatory: false, MaxRepeat: 9},  // Clinical information
									{Segment: "DTM", Mandatory: false, MaxRepeat: 99}, // Date/time/period
									{Segment: "FTX", Mandatory: false, MaxRepeat: 99}, // Free text
									{Segment: "RFF", Mandatory: false, MaxRepeat: 99}, // Reference
									{ // Segment group 10
										Group: []SchemaNode{
											{Segment: "RSL", Mandatory: true, MaxRepeat: 1},  // Result
											{Segment: "FTX", Mandatory: false, MaxRepeat: 9}, // Free text
											{Segment: "CCI", Mandatory: false, MaxRepeat: 9}, // Characteristic/class id
										},
										Mandatory: false, MaxRepeat: 99,
									},
									{ // Segment group 11
										Group: []SchemaNode{
											{Segment: "REL", Mandatory: true, MaxRepeat: 1},   // Relationship
											{Segment: "RFF", Mandatory: false, MaxRepeat: 99}, // Reference
										},
										Mandatory: false, MaxRepeat: 99,
									},
								},
								Mandatory: false, MaxRepeat: 99,
							},
							{ // Segment group 12
								Group: []SchemaNode{
									{Segment: "CLI", Mandatory: true, MaxRepeat: 1}, // Clinical intervention
									{ // Segment group 13
										Group: []SchemaNode{
											{Segment: "IMD", Mandatory: true, MaxRepeat: 1},  // Item description
											{Segment: "DSG", Mandatory: false, MaxRepeat: 9}, // Dosage administration
											{Segment: "FTX", Mandatory: false, MaxRepeat: 9}, // Free text
											{Segment: "INP", Mandatory: false, MaxRepeat: 9}, // Parties and instruction
										},
										Mandatory: false, MaxRepeat: 99,
									},
								},
								Mandatory: false, MaxRepeat: 99,
							},
						},
						Mandatory: false, MaxRepeat: 999,
					},
					{ // Segment group 14
						Group: []SchemaNode{
							{Segment: "SEQ", Mandatory: true, MaxRepeat: 1},   // Sequence details
							{Segment: "IMD", Mandatory: true, MaxRepeat: 9},   // Item description
							{Segment: "PRC", Mandatory: false, MaxRepeat: 9},  // Process identification
							{Segment: "RFF", Mandatory: false, MaxRepeat: 9},  // Reference
							{Segment: "QTY", Mandatory: false, MaxRepeat: 1},  // Quantity
							{Segment: "DTM", Mandatory: false, MaxRepeat: 99}, // Date/time/period
							{Segment: "PAC", Mandatory: false, MaxRepeat: 1},  // Package
							{Segment: "FTX", Mandatory: false, MaxRepeat: 9},  // Free text
							{Segment: "TDT", Mandatory: false, MaxRepeat: 9},  // Transport information
							{Segment: "HAN", Mandatory: false, MaxRepeat: 9},  // Handling instructions
							{Segment: "LOC", Mandatory: false, MaxRepeat: 9},  // Place/location identification
							{Segment: "ADR", Mandatory: false, MaxRepeat: 9},  // Address
							{ // Segment group 15
								Group: []SchemaNode{
									{Segment: "CLI", Mandatory: true, MaxRepeat: 1},  // Clinical intervention
									{Segment: "IMD", Mandatory: false, MaxRepeat: 1}, // Item description
									{Segment: "DSG", Mandatory: false, MaxRepeat: 1}, // Dosage administration
									{Segment: "DTM", Mandatory: false, MaxRepeat: 9}, // Date/time/period
									{Segment: "FTX", Mandatory: false, MaxRepeat: 9}, // Free text
									{Segment: "QTY", Mandatory: false, MaxRepeat: 1}, // Quantity
									{Segment: "INP", Mandatory: false, MaxRepeat: 9}, // Parties and instruction
								},
								Mandatory: false, MaxRepeat: 99,
							},
						},
						Mandatory: false, MaxRepeat: 99,
					},
					{ // Segment group 16
						Group: []SchemaNode{
							{Segment: "LIN", Mandatory: true, MaxRepeat: 1},   // Line item
							{Segment: "GEI", Mandatory: true, MaxRepeat: 1},   // Processing information
							{Segment: "RSL", Mandatory: false, MaxRepeat: 1},  // Result
							{Segment: "CCI", Mandatory: true, MaxRepeat: 99},  // Characteristic/class id
							{Segment: "CIN", Mandatory: false, MaxRepeat: 9},  // Clinical information
							{Segment: "SEQ", Mandatory: false, MaxRepeat: 1},  // Sequence details
							{Segment: "STS", Mandatory: false, MaxRepeat: 1},  // Status
							{Segment: "DTM", Mandatory: false, MaxRepeat: 9},  // Date/time/period
							{Segment: "FTX", Mandatory: false, MaxRepeat: 99}, // Free text
							{Segment: "MOA", Mandatory: false, MaxRepeat: 9},  // Monetary amount
							{Segment: "RFF", Mandatory: false, MaxRepeat: 99}, // Reference
							{Segment: "EQD", Mandatory: false, MaxRepeat: 9},  // Equipment details
							{ // Segment group 17
								Group: []SchemaNode{
									{Segment: "REL", Mandatory: true, MaxRepeat: 1},   // Relationship
									{Segment: "RFF", Mandatory: false, MaxRepeat: 99}, // Reference
								},
								Mandatory: false, MaxRepeat: 99,
							},
							{ // Segment group 18
								Group: []SchemaNode{
									{Segment: "RSL", Mandatory: true, MaxRepeat: 1},  // Result
									{Segment: "FTX", Mandatory: false, MaxRepeat: 9}, // Free text
									{Segment: "CCI", Mandatory: false, MaxRepeat: 9}, // Characteristic/class id
								},
								Mandatory: false, MaxRepeat: 99,
							},
						},
						Mandatory: true, MaxRepeat: 99,
					},
				},
				Mandatory: false, MaxRepeat: 99,
			},
		},
		Mandatory: true, MaxRepeat: 99,
	},
	{ // Segment group 19
		Group: []SchemaNode{
			{Segment: "AUT", Mandatory: true, MaxRepeat: 1},  // Authentication result
			{Segment: "DTM", Mandatory: false, MaxRepeat: 9}, // Date/time/period
		},
		Mandatory: false, MaxRepeat: 99,
	},
}}
