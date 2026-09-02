package edifact

// CLASET D.20A message specification: the structural (segment/group
// presence, order, cardinality) branching diagram for this message,
// UN/EDIFACT directory release D.20A. Scope is structural only -- not
// element-level content -- per edifact-ls-3uzr's epic non-goal.
// 20 segment groups, max nesting depth 5.
//
// Source: https://service.unece.org/trade/untdid/d20a/trmd/claset_c.htm
// section 4.3.1 "Segment table". That URL currently returns HTTP 403 from
// Cloudflare when fetched directly; this data was transcribed from the
// Wayback Machine's archived copy instead:
// http://web.archive.org/web/20231202081727/https://service.unece.org/trade/untdid/d20a/trmd/claset_c.htm
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
		MessageID{Type: "CLASET", Version: "D", Release: "20A", Agency: "UN"},
		clasetD20aSchema,
		"https://service.unece.org/trade/untdid/d20a/trmd/claset_c.htm",
	)
}

var clasetD20aSchema = Schema{Nodes: []SchemaNode{
	{Segment: "BGM", Mandatory: true, MaxRepeat: 1},  // Beginning of message
	{Segment: "DTM", Mandatory: false, MaxRepeat: 1}, // Date/time/period
	{ // Segment group 1
		Group: []SchemaNode{
			{Segment: "RFF", Mandatory: true, MaxRepeat: 1},  // Reference
			{Segment: "DTM", Mandatory: false, MaxRepeat: 9}, // Date/time/period
		},
		Mandatory: false, MaxRepeat: 9,
	},
	{ // Segment group 2
		Group: []SchemaNode{
			{Segment: "PNA", Mandatory: true, MaxRepeat: 1},  // Party identification
			{Segment: "ADR", Mandatory: false, MaxRepeat: 1}, // Address
			{ // Segment group 3
				Group: []SchemaNode{
					{Segment: "CTA", Mandatory: true, MaxRepeat: 1},  // Contact information
					{Segment: "COM", Mandatory: false, MaxRepeat: 9}, // Communication contact
				},
				Mandatory: false, MaxRepeat: 9,
			},
		},
		Mandatory: false, MaxRepeat: 9,
	},
	{ // Segment group 4
		Group: []SchemaNode{
			{Segment: "VLI", Mandatory: true, MaxRepeat: 1},  // Value list identification
			{Segment: "DTM", Mandatory: false, MaxRepeat: 9}, // Date/time/period
			{Segment: "EQN", Mandatory: false, MaxRepeat: 1}, // Number of units
			{ // Segment group 5
				Group: []SchemaNode{
					{Segment: "PNA", Mandatory: true, MaxRepeat: 1},  // Party identification
					{Segment: "LAN", Mandatory: false, MaxRepeat: 1}, // Language
					{ // Segment group 6
						Group: []SchemaNode{
							{Segment: "CTA", Mandatory: true, MaxRepeat: 1},  // Contact information
							{Segment: "ADR", Mandatory: false, MaxRepeat: 1}, // Address
							{Segment: "COM", Mandatory: false, MaxRepeat: 9}, // Communication contact
						},
						Mandatory: false, MaxRepeat: 9,
					},
				},
				Mandatory: false, MaxRepeat: 9,
			},
			{ // Segment group 7
				Group: []SchemaNode{
					{Segment: "STS", Mandatory: true, MaxRepeat: 1},  // Status
					{Segment: "DTM", Mandatory: false, MaxRepeat: 9}, // Date/time/period
				},
				Mandatory: false, MaxRepeat: 99,
			},
			{ // Segment group 8
				Group: []SchemaNode{
					{Segment: "ATT", Mandatory: true, MaxRepeat: 1},  // Attribute
					{Segment: "DTM", Mandatory: false, MaxRepeat: 9}, // Date/time/period
					{Segment: "ELM", Mandatory: false, MaxRepeat: 1}, // Simple data element details
					{ // Segment group 9
						Group: []SchemaNode{
							{Segment: "CAV", Mandatory: true, MaxRepeat: 1},   // Characteristic value
							{Segment: "FTX", Mandatory: false, MaxRepeat: 99}, // Free text
						},
						Mandatory: false, MaxRepeat: 9,
					},
					{ // Segment group 10
						Group: []SchemaNode{
							{Segment: "STS", Mandatory: true, MaxRepeat: 1},  // Status
							{Segment: "DTM", Mandatory: false, MaxRepeat: 9}, // Date/time/period
						},
						Mandatory: false, MaxRepeat: 99,
					},
				},
				Mandatory: false, MaxRepeat: 99,
			},
			{ // Segment group 11
				Group: []SchemaNode{
					{Segment: "SCD", Mandatory: true, MaxRepeat: 1},  // Structure component definition
					{Segment: "DTM", Mandatory: false, MaxRepeat: 9}, // Date/time/period
					{Segment: "EQN", Mandatory: false, MaxRepeat: 1}, // Number of units
					{ // Segment group 12
						Group: []SchemaNode{
							{Segment: "STS", Mandatory: true, MaxRepeat: 1},  // Status
							{Segment: "DTM", Mandatory: false, MaxRepeat: 9}, // Date/time/period
						},
						Mandatory: false, MaxRepeat: 99,
					},
					{ // Segment group 13
						Group: []SchemaNode{
							{Segment: "ATT", Mandatory: true, MaxRepeat: 1},  // Attribute
							{Segment: "DTM", Mandatory: false, MaxRepeat: 9}, // Date/time/period
							{Segment: "ELM", Mandatory: false, MaxRepeat: 1}, // Simple data element details
							{ // Segment group 14
								Group: []SchemaNode{
									{Segment: "CAV", Mandatory: true, MaxRepeat: 1},   // Characteristic value
									{Segment: "FTX", Mandatory: false, MaxRepeat: 99}, // Free text
								},
								Mandatory: false, MaxRepeat: 99,
							},
							{ // Segment group 15
								Group: []SchemaNode{
									{Segment: "STS", Mandatory: true, MaxRepeat: 1},  // Status
									{Segment: "DTM", Mandatory: false, MaxRepeat: 9}, // Date/time/period
								},
								Mandatory: false, MaxRepeat: 99,
							},
						},
						Mandatory: false, MaxRepeat: 99,
					},
					{ // Segment group 16
						Group: []SchemaNode{
							{Segment: "IDE", Mandatory: true, MaxRepeat: 1},  // Identity
							{Segment: "DTM", Mandatory: false, MaxRepeat: 9}, // Date/time/period
							{Segment: "EQN", Mandatory: false, MaxRepeat: 1}, // Number of units
							{ // Segment group 17
								Group: []SchemaNode{
									{Segment: "STS", Mandatory: true, MaxRepeat: 1},  // Status
									{Segment: "DTM", Mandatory: false, MaxRepeat: 9}, // Date/time/period
								},
								Mandatory: false, MaxRepeat: 99,
							},
							{ // Segment group 18
								Group: []SchemaNode{
									{Segment: "ATT", Mandatory: true, MaxRepeat: 1},  // Attribute
									{Segment: "DTM", Mandatory: false, MaxRepeat: 2}, // Date/time/period
									{Segment: "ELM", Mandatory: false, MaxRepeat: 1}, // Simple data element details
									{ // Segment group 19
										Group: []SchemaNode{
											{Segment: "CAV", Mandatory: true, MaxRepeat: 1},   // Characteristic value
											{Segment: "FTX", Mandatory: false, MaxRepeat: 99}, // Free text
										},
										Mandatory: false, MaxRepeat: 99,
									},
									{ // Segment group 20
										Group: []SchemaNode{
											{Segment: "STS", Mandatory: true, MaxRepeat: 1},  // Status
											{Segment: "DTM", Mandatory: false, MaxRepeat: 9}, // Date/time/period
										},
										Mandatory: false, MaxRepeat: 99,
									},
								},
								Mandatory: false, MaxRepeat: 99,
							},
						},
						Mandatory: false, MaxRepeat: 99,
					},
				},
				Mandatory: false, MaxRepeat: 999999,
			},
		},
		Mandatory: false, MaxRepeat: 99,
	},
}}
