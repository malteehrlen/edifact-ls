package edifact

// GESMES D.20A message specification: the structural (segment/group
// presence, order, cardinality) branching diagram for this message,
// UN/EDIFACT directory release D.20A. Scope is structural only -- not
// element-level content -- per edifact-ls-3uzr's epic non-goal.
// 25 segment groups, max nesting depth 6.
//
// Source: https://service.unece.org/trade/untdid/d20a/trmd/gesmes_c.htm
// section 4.3.1 "Segment table". That URL currently returns HTTP 403 from
// Cloudflare when fetched directly; this data was transcribed from the
// Wayback Machine's archived copy instead:
// http://web.archive.org/web/20231202094224/https://service.unece.org/trade/untdid/d20a/trmd/gesmes_c.htm
// -- re-check the direct URL first if cross-checking later, in case the
// block has lifted.
//
// Transcription note: generated in bulk (edifact-ls-13gu) the same way
// as every other message-type schema -- a script parsed the source's
// exact rail-art column positions mechanically and verified the result
// balances (all 25 groups open and close consistently)
// before being transcribed here, rather than reading the ASCII nesting
// by eye. Not individually spot-checked against the raw source the way
// the first several message types in this project were -- see
// edifact-ls-13gu for which representative samples were spot-checked
// for this batch.

func init() {
	RegisterSchema(
		MessageID{Type: "GESMES", Version: "D", Release: "20A", Agency: "UN"},
		gesmesD20aSchema,
		"https://service.unece.org/trade/untdid/d20a/trmd/gesmes_c.htm",
	)
}

var gesmesD20aSchema = Schema{Nodes: []SchemaNode{
	{Segment: "BGM", Mandatory: false, MaxRepeat: 1}, // Beginning of message
	{Segment: "DTM", Mandatory: false, MaxRepeat: 9}, // Date/time/period
	{ // Segment group 1
		Group: []SchemaNode{
			{Segment: "FNT", Mandatory: true, MaxRepeat: 1},     // Footnote
			{Segment: "FTX", Mandatory: false, MaxRepeat: 9999}, // Free text
		},
		Mandatory: false, MaxRepeat: 9999,
	},
	{ // Segment group 2
		Group: []SchemaNode{
			{Segment: "NAD", Mandatory: true, MaxRepeat: 1},  // Name and address
			{Segment: "IDE", Mandatory: false, MaxRepeat: 1}, // Identity
			{ // Segment group 3
				Group: []SchemaNode{
					{Segment: "CTA", Mandatory: true, MaxRepeat: 1},  // Contact information
					{Segment: "COM", Mandatory: false, MaxRepeat: 5}, // Communication contact
				},
				Mandatory: false, MaxRepeat: 5,
			},
		},
		Mandatory: false, MaxRepeat: 9,
	},
	{ // Segment group 4
		Group: []SchemaNode{
			{Segment: "VLI", Mandatory: true, MaxRepeat: 1},     // Value list identification
			{Segment: "DTM", Mandatory: false, MaxRepeat: 9999}, // Date/time/period
			{Segment: "FTX", Mandatory: false, MaxRepeat: 9999}, // Free text
			{ // Segment group 5
				Group: []SchemaNode{
					{Segment: "IDE", Mandatory: true, MaxRepeat: 1},  // Identity
					{Segment: "DTM", Mandatory: false, MaxRepeat: 9}, // Date/time/period
				},
				Mandatory: false, MaxRepeat: 9,
			},
			{ // Segment group 6
				Group: []SchemaNode{
					{Segment: "GIR", Mandatory: true, MaxRepeat: 1},  // Related identification numbers
					{Segment: "DTM", Mandatory: false, MaxRepeat: 2}, // Date/time/period
				},
				Mandatory: false, MaxRepeat: 5,
			},
			{ // Segment group 7
				Group: []SchemaNode{
					{Segment: "CDV", Mandatory: true, MaxRepeat: 1},    // Code value definition
					{Segment: "FTX", Mandatory: false, MaxRepeat: 999}, // Free text
					{ // Segment group 8
						Group: []SchemaNode{
							{Segment: "IDE", Mandatory: true, MaxRepeat: 1},  // Identity
							{Segment: "DTM", Mandatory: false, MaxRepeat: 9}, // Date/time/period
						},
						Mandatory: false, MaxRepeat: 9,
					},
				},
				Mandatory: false, MaxRepeat: 9999,
			},
		},
		Mandatory: false, MaxRepeat: 99,
	},
	{ // Segment group 9
		Group: []SchemaNode{
			{Segment: "STC", Mandatory: true, MaxRepeat: 1},  // Statistical concept
			{Segment: "FTX", Mandatory: false, MaxRepeat: 5}, // Free text
			{Segment: "IDE", Mandatory: false, MaxRepeat: 5}, // Identity
		},
		Mandatory: false, MaxRepeat: 999,
	},
	{ // Segment group 10
		Group: []SchemaNode{
			{Segment: "ASI", Mandatory: true, MaxRepeat: 1},  // Array structure identification
			{Segment: "GEI", Mandatory: false, MaxRepeat: 5}, // Processing information
			{Segment: "FTX", Mandatory: false, MaxRepeat: 9}, // Free text
			{Segment: "IDE", Mandatory: false, MaxRepeat: 5}, // Identity
			{ // Segment group 11
				Group: []SchemaNode{
					{Segment: "SCD", Mandatory: true, MaxRepeat: 1},   // Structure component definition
					{Segment: "ATT", Mandatory: false, MaxRepeat: 99}, // Attribute
					{ // Segment group 12
						Group: []SchemaNode{
							{Segment: "IDE", Mandatory: true, MaxRepeat: 1},   // Identity
							{Segment: "ATT", Mandatory: false, MaxRepeat: 99}, // Attribute
							{Segment: "DTM", Mandatory: false, MaxRepeat: 2},  // Date/time/period
						},
						Mandatory: false, MaxRepeat: 9999,
					},
				},
				Mandatory: false, MaxRepeat: 999,
			},
		},
		Mandatory: false, MaxRepeat: 9999,
	},
	{ // Segment group 13
		Group: []SchemaNode{
			{Segment: "DSI", Mandatory: true, MaxRepeat: 1},  // Data set identification
			{Segment: "STS", Mandatory: false, MaxRepeat: 9}, // Status
			{Segment: "DTM", Mandatory: false, MaxRepeat: 5}, // Date/time/period
			{Segment: "FTX", Mandatory: false, MaxRepeat: 5}, // Free text
			{Segment: "GIR", Mandatory: false, MaxRepeat: 2}, // Related identification numbers
			{ // Segment group 14
				Group: []SchemaNode{
					{Segment: "IDE", Mandatory: true, MaxRepeat: 1},  // Identity
					{Segment: "GEI", Mandatory: false, MaxRepeat: 5}, // Processing information
					{Segment: "DTM", Mandatory: false, MaxRepeat: 5}, // Date/time/period
					{ // Segment group 15
						Group: []SchemaNode{
							{Segment: "CDV", Mandatory: true, MaxRepeat: 1},   // Code value definition
							{Segment: "FTX", Mandatory: false, MaxRepeat: 99}, // Free text
						},
						Mandatory: false, MaxRepeat: 999,
					},
					{ // Segment group 16
						Group: []SchemaNode{
							{Segment: "SCD", Mandatory: true, MaxRepeat: 1},     // Structure component definition
							{Segment: "ATT", Mandatory: false, MaxRepeat: 9},    // Attribute
							{Segment: "DTM", Mandatory: false, MaxRepeat: 9999}, // Date/time/period
							{ // Segment group 17
								Group: []SchemaNode{
									{Segment: "CDV", Mandatory: true, MaxRepeat: 1},   // Code value definition
									{Segment: "FTX", Mandatory: false, MaxRepeat: 99}, // Free text
								},
								Mandatory: false, MaxRepeat: 9999,
							},
						},
						Mandatory: false, MaxRepeat: 999,
					},
				},
				Mandatory: false, MaxRepeat: 99,
			},
			{Segment: "ARR", Mandatory: false, MaxRepeat: 999999}, // Array information
			{ // Segment group 18
				Group: []SchemaNode{
					{Segment: "NAD", Mandatory: true, MaxRepeat: 1},  // Name and address
					{Segment: "IDE", Mandatory: false, MaxRepeat: 1}, // Identity
					{ // Segment group 19
						Group: []SchemaNode{
							{Segment: "CTA", Mandatory: true, MaxRepeat: 1},  // Contact information
							{Segment: "COM", Mandatory: false, MaxRepeat: 5}, // Communication contact
						},
						Mandatory: false, MaxRepeat: 5,
					},
				},
				Mandatory: false, MaxRepeat: 5,
			},
			{ // Segment group 20
				Group: []SchemaNode{
					{Segment: "FNS", Mandatory: true, MaxRepeat: 1}, // Footnote set
					{ // Segment group 21
						Group: []SchemaNode{
							{Segment: "REL", Mandatory: true, MaxRepeat: 1}, // Relationship
							{ // Segment group 22
								Group: []SchemaNode{
									{Segment: "ARR", Mandatory: true, MaxRepeat: 1},     // Array information
									{Segment: "FTX", Mandatory: false, MaxRepeat: 9999}, // Free text
									{ // Segment group 23
										Group: []SchemaNode{
											{Segment: "IDE", Mandatory: true, MaxRepeat: 1},     // Identity
											{Segment: "FTX", Mandatory: false, MaxRepeat: 9999}, // Free text
											{ // Segment group 24
												Group: []SchemaNode{
													{Segment: "CDV", Mandatory: true, MaxRepeat: 1},     // Code value definition
													{Segment: "FTX", Mandatory: false, MaxRepeat: 9999}, // Free text
												},
												Mandatory: false, MaxRepeat: 9999,
											},
										},
										Mandatory: false, MaxRepeat: 9999,
									},
								},
								Mandatory: false, MaxRepeat: 9999,
							},
						},
						Mandatory: false, MaxRepeat: 9999,
					},
				},
				Mandatory: false, MaxRepeat: 999,
			},
			{ // Segment group 25
				Group: []SchemaNode{
					{Segment: "FNT", Mandatory: true, MaxRepeat: 1},     // Footnote
					{Segment: "FTX", Mandatory: false, MaxRepeat: 9999}, // Free text
				},
				Mandatory: false, MaxRepeat: 9999,
			},
		},
		Mandatory: false, MaxRepeat: 999,
	},
}}
