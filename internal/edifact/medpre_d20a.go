package edifact

// MEDPRE D.20A message specification: the structural (segment/group
// presence, order, cardinality) branching diagram for this message,
// UN/EDIFACT directory release D.20A. Scope is structural only -- not
// element-level content -- per edifact-ls-3uzr's epic non-goal.
// 18 segment groups, max nesting depth 4.
//
// Source: https://service.unece.org/trade/untdid/d20a/trmd/medpre_c.htm
// section 4.3.1 "Segment table". That URL currently returns HTTP 403 from
// Cloudflare when fetched directly; this data was transcribed from the
// Wayback Machine's archived copy instead:
// http://web.archive.org/web/20230608021210/https://service.unece.org/trade/untdid/d20a/trmd/medpre_c.htm
// -- re-check the direct URL first if cross-checking later, in case the
// block has lifted.
//
// Transcription note: generated in bulk (edifact-ls-13gu) the same way
// as every other message-type schema -- a script parsed the source's
// exact rail-art column positions mechanically and verified the result
// balances (all 18 groups open and close consistently)
// before being transcribed here, rather than reading the ASCII nesting
// by eye. Not individually spot-checked against the raw source the way
// the first several message types in this project were -- see
// edifact-ls-13gu for which representative samples were spot-checked
// for this batch.

func init() {
	RegisterSchema(
		MessageID{Type: "MEDPRE", Version: "D", Release: "20A", Agency: "UN"},
		medpreD20aSchema,
		"https://service.unece.org/trade/untdid/d20a/trmd/medpre_c.htm",
	)
}

var medpreD20aSchema = Schema{Nodes: []SchemaNode{
	{Segment: "BGM", Mandatory: true, MaxRepeat: 1},  // Beginning of message
	{Segment: "DTM", Mandatory: false, MaxRepeat: 9}, // Date/time/period
	{Segment: "FTX", Mandatory: false, MaxRepeat: 9}, // Free text
	{ // Segment group 1
		Group: []SchemaNode{
			{Segment: "SEQ", Mandatory: true, MaxRepeat: 1},  // Sequence details
			{Segment: "PNA", Mandatory: false, MaxRepeat: 9}, // Party identification
			{Segment: "IDE", Mandatory: false, MaxRepeat: 9}, // Identity
			{Segment: "SPR", Mandatory: false, MaxRepeat: 9}, // Organisation classification details
			{Segment: "QUA", Mandatory: false, MaxRepeat: 9}, // Qualification
			{Segment: "EMP", Mandatory: false, MaxRepeat: 1}, // Employment details
			{Segment: "ADR", Mandatory: false, MaxRepeat: 1}, // Address
			{Segment: "COM", Mandatory: false, MaxRepeat: 9}, // Communication contact
			{Segment: "RFF", Mandatory: false, MaxRepeat: 1}, // Reference
			{ // Segment group 2
				Group: []SchemaNode{
					{Segment: "DOC", Mandatory: true, MaxRepeat: 1},  // Document/message details
					{Segment: "COM", Mandatory: false, MaxRepeat: 1}, // Communication contact
				},
				Mandatory: false, MaxRepeat: 9,
			},
		},
		Mandatory: true, MaxRepeat: 99,
	},
	{ // Segment group 3
		Group: []SchemaNode{
			{Segment: "ATT", Mandatory: true, MaxRepeat: 1},  // Attribute
			{Segment: "PNA", Mandatory: false, MaxRepeat: 9}, // Party identification
			{Segment: "PDI", Mandatory: false, MaxRepeat: 1}, // Person demographic information
			{Segment: "IDE", Mandatory: false, MaxRepeat: 9}, // Identity
			{Segment: "DTM", Mandatory: false, MaxRepeat: 2}, // Date/time/period
			{Segment: "NAT", Mandatory: false, MaxRepeat: 1}, // Nationality
			{Segment: "AGR", Mandatory: false, MaxRepeat: 9}, // Agreement identification
			{Segment: "CCI", Mandatory: false, MaxRepeat: 2}, // Characteristic/class id
			{Segment: "STS", Mandatory: false, MaxRepeat: 1}, // Status
			{Segment: "FTX", Mandatory: false, MaxRepeat: 9}, // Free text
			{Segment: "LAN", Mandatory: false, MaxRepeat: 1}, // Language
			{Segment: "CAV", Mandatory: false, MaxRepeat: 1}, // Characteristic value
			{ // Segment group 4
				Group: []SchemaNode{
					{Segment: "HAN", Mandatory: true, MaxRepeat: 1},  // Handling instructions
					{Segment: "FTX", Mandatory: false, MaxRepeat: 9}, // Free text
					{Segment: "DTM", Mandatory: false, MaxRepeat: 2}, // Date/time/period
				},
				Mandatory: false, MaxRepeat: 9,
			},
			{ // Segment group 5
				Group: []SchemaNode{
					{Segment: "ADR", Mandatory: true, MaxRepeat: 1},  // Address
					{Segment: "COM", Mandatory: false, MaxRepeat: 9}, // Communication contact
				},
				Mandatory: false, MaxRepeat: 9,
			},
			{ // Segment group 6
				Group: []SchemaNode{
					{Segment: "REL", Mandatory: true, MaxRepeat: 1},  // Relationship
					{Segment: "PNA", Mandatory: true, MaxRepeat: 9},  // Party identification
					{Segment: "PDI", Mandatory: false, MaxRepeat: 1}, // Person demographic information
					{Segment: "ADR", Mandatory: false, MaxRepeat: 1}, // Address
					{Segment: "COM", Mandatory: false, MaxRepeat: 9}, // Communication contact
					{Segment: "LAN", Mandatory: false, MaxRepeat: 1}, // Language
					{Segment: "CAV", Mandatory: false, MaxRepeat: 1}, // Characteristic value
				},
				Mandatory: false, MaxRepeat: 99,
			},
			{ // Segment group 7
				Group: []SchemaNode{
					{Segment: "LIN", Mandatory: true, MaxRepeat: 1},   // Line item
					{Segment: "CIN", Mandatory: false, MaxRepeat: 9},  // Clinical information
					{Segment: "DTM", Mandatory: false, MaxRepeat: 9},  // Date/time/period
					{Segment: "PNA", Mandatory: false, MaxRepeat: 9},  // Party identification
					{Segment: "LAN", Mandatory: false, MaxRepeat: 1},  // Language
					{Segment: "FTX", Mandatory: false, MaxRepeat: 99}, // Free text
					{Segment: "RSL", Mandatory: false, MaxRepeat: 1},  // Result
					{Segment: "CLI", Mandatory: false, MaxRepeat: 9},  // Clinical intervention
					{Segment: "CCI", Mandatory: false, MaxRepeat: 9},  // Characteristic/class id
					{ // Segment group 8
						Group: []SchemaNode{
							{Segment: "IMD", Mandatory: true, MaxRepeat: 1},  // Item description
							{Segment: "STS", Mandatory: true, MaxRepeat: 1},  // Status
							{Segment: "DSG", Mandatory: false, MaxRepeat: 9}, // Dosage administration
							{Segment: "INP", Mandatory: false, MaxRepeat: 9}, // Parties and instruction
							{Segment: "FTX", Mandatory: false, MaxRepeat: 9}, // Free text
							{Segment: "SCC", Mandatory: false, MaxRepeat: 9}, // Scheduling conditions
							{Segment: "QTY", Mandatory: false, MaxRepeat: 9}, // Quantity
						},
						Mandatory: false, MaxRepeat: 1,
					},
				},
				Mandatory: false, MaxRepeat: 999,
			},
		},
		Mandatory: false, MaxRepeat: 1,
	},
	{ // Segment group 9
		Group: []SchemaNode{
			{Segment: "GEI", Mandatory: true, MaxRepeat: 1},   // Processing information
			{Segment: "IDE", Mandatory: true, MaxRepeat: 9},   // Identity
			{Segment: "DTM", Mandatory: true, MaxRepeat: 9},   // Date/time/period
			{Segment: "PTY", Mandatory: false, MaxRepeat: 1},  // Priority
			{Segment: "AGR", Mandatory: false, MaxRepeat: 1},  // Agreement identification
			{Segment: "LAN", Mandatory: false, MaxRepeat: 1},  // Language
			{Segment: "STS", Mandatory: false, MaxRepeat: 1},  // Status
			{Segment: "FTX", Mandatory: false, MaxRepeat: 99}, // Free text
			{ // Segment group 10
				Group: []SchemaNode{
					{Segment: "DOC", Mandatory: true, MaxRepeat: 1},  // Document/message details
					{Segment: "IDE", Mandatory: false, MaxRepeat: 1}, // Identity
					{Segment: "DTM", Mandatory: false, MaxRepeat: 1}, // Date/time/period
				},
				Mandatory: false, MaxRepeat: 9,
			},
			{ // Segment group 11
				Group: []SchemaNode{
					{Segment: "TOD", Mandatory: true, MaxRepeat: 1},  // Terms of delivery or transport
					{Segment: "TDT", Mandatory: false, MaxRepeat: 1}, // Transport information
					{Segment: "FTX", Mandatory: false, MaxRepeat: 9}, // Free text
					{Segment: "DTM", Mandatory: false, MaxRepeat: 9}, // Date/time/period
					{Segment: "ADR", Mandatory: false, MaxRepeat: 1}, // Address
					{Segment: "COM", Mandatory: false, MaxRepeat: 9}, // Communication contact
					{Segment: "PNA", Mandatory: false, MaxRepeat: 9}, // Party identification
					{Segment: "PTY", Mandatory: false, MaxRepeat: 1}, // Priority
					{Segment: "PAC", Mandatory: false, MaxRepeat: 1}, // Package
				},
				Mandatory: false, MaxRepeat: 1,
			},
			{ // Segment group 12
				Group: []SchemaNode{
					{Segment: "FCA", Mandatory: true, MaxRepeat: 1},  // Financial charges allocation
					{Segment: "PNA", Mandatory: false, MaxRepeat: 9}, // Party identification
					{Segment: "IDE", Mandatory: false, MaxRepeat: 9}, // Identity
					{Segment: "RCS", Mandatory: false, MaxRepeat: 9}, // Requirements and conditions
					{ // Segment group 13
						Group: []SchemaNode{
							{Segment: "ICD", Mandatory: true, MaxRepeat: 1},  // Insurance cover description
							{Segment: "ALC", Mandatory: false, MaxRepeat: 9}, // Allowance or charge
							{Segment: "PCD", Mandatory: false, MaxRepeat: 1}, // Percentage details
							{Segment: "MOA", Mandatory: false, MaxRepeat: 9}, // Monetary amount
							{Segment: "IDE", Mandatory: false, MaxRepeat: 9}, // Identity
							{Segment: "DTM", Mandatory: false, MaxRepeat: 9}, // Date/time/period
						},
						Mandatory: false, MaxRepeat: 9,
					},
				},
				Mandatory: false, MaxRepeat: 9,
			},
			{ // Segment group 14
				Group: []SchemaNode{
					{Segment: "LIN", Mandatory: true, MaxRepeat: 1},   // Line item
					{Segment: "DSG", Mandatory: false, MaxRepeat: 9},  // Dosage administration
					{Segment: "IMD", Mandatory: true, MaxRepeat: 1},   // Item description
					{Segment: "QTY", Mandatory: false, MaxRepeat: 9},  // Quantity
					{Segment: "PGI", Mandatory: false, MaxRepeat: 9},  // Product group information
					{Segment: "PNA", Mandatory: false, MaxRepeat: 9},  // Party identification
					{Segment: "PAC", Mandatory: false, MaxRepeat: 1},  // Package
					{Segment: "IDE", Mandatory: false, MaxRepeat: 9},  // Identity
					{Segment: "DLM", Mandatory: false, MaxRepeat: 9},  // Delivery limitations
					{Segment: "EQN", Mandatory: false, MaxRepeat: 1},  // Number of units
					{Segment: "PRC", Mandatory: false, MaxRepeat: 9},  // Process identification
					{Segment: "DTM", Mandatory: false, MaxRepeat: 9},  // Date/time/period
					{Segment: "STS", Mandatory: false, MaxRepeat: 1},  // Status
					{Segment: "CIN", Mandatory: false, MaxRepeat: 99}, // Clinical information
					{Segment: "FTX", Mandatory: false, MaxRepeat: 99}, // Free text
					{ // Segment group 15
						Group: []SchemaNode{
							{Segment: "CAV", Mandatory: true, MaxRepeat: 1},   // Characteristic value
							{Segment: "IMD", Mandatory: true, MaxRepeat: 1},   // Item description
							{Segment: "QTY", Mandatory: false, MaxRepeat: 9},  // Quantity
							{Segment: "FTX", Mandatory: false, MaxRepeat: 99}, // Free text
						},
						Mandatory: false, MaxRepeat: 99,
					},
					{ // Segment group 16
						Group: []SchemaNode{
							{Segment: "SEQ", Mandatory: true, MaxRepeat: 1},  // Sequence details
							{Segment: "DSG", Mandatory: false, MaxRepeat: 9}, // Dosage administration
							{Segment: "QTY", Mandatory: false, MaxRepeat: 9}, // Quantity
							{Segment: "INP", Mandatory: false, MaxRepeat: 9}, // Parties and instruction
							{Segment: "DTM", Mandatory: false, MaxRepeat: 9}, // Date/time/period
							{Segment: "FTX", Mandatory: false, MaxRepeat: 9}, // Free text
							{Segment: "SCC", Mandatory: false, MaxRepeat: 9}, // Scheduling conditions
							{Segment: "CIN", Mandatory: false, MaxRepeat: 9}, // Clinical information
							{Segment: "PCI", Mandatory: false, MaxRepeat: 9}, // Package identification
							{Segment: "LAN", Mandatory: false, MaxRepeat: 1}, // Language
							{Segment: "EQA", Mandatory: false, MaxRepeat: 9}, // Attached equipment
						},
						Mandatory: false, MaxRepeat: 99,
					},
					{ // Segment group 17
						Group: []SchemaNode{
							{Segment: "FCA", Mandatory: true, MaxRepeat: 1},  // Financial charges allocation
							{Segment: "PNA", Mandatory: false, MaxRepeat: 9}, // Party identification
							{Segment: "IDE", Mandatory: false, MaxRepeat: 9}, // Identity
							{Segment: "RCS", Mandatory: false, MaxRepeat: 9}, // Requirements and conditions
							{ // Segment group 18
								Group: []SchemaNode{
									{Segment: "ICD", Mandatory: true, MaxRepeat: 1},  // Insurance cover description
									{Segment: "ALC", Mandatory: false, MaxRepeat: 9}, // Allowance or charge
									{Segment: "PCD", Mandatory: false, MaxRepeat: 1}, // Percentage details
									{Segment: "MOA", Mandatory: false, MaxRepeat: 9}, // Monetary amount
									{Segment: "IDE", Mandatory: false, MaxRepeat: 9}, // Identity
									{Segment: "DTM", Mandatory: false, MaxRepeat: 9}, // Date/time/period
								},
								Mandatory: false, MaxRepeat: 9,
							},
						},
						Mandatory: false, MaxRepeat: 9,
					},
				},
				Mandatory: false, MaxRepeat: 99,
			},
		},
		Mandatory: true, MaxRepeat: 1,
	},
	{Segment: "CNT", Mandatory: false, MaxRepeat: 1}, // Control total
}}
