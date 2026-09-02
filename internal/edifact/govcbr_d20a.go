package edifact

// GOVCBR D.20A message specification: the structural (segment/group
// presence, order, cardinality) branching diagram for this message,
// UN/EDIFACT directory release D.20A. Scope is structural only -- not
// element-level content -- per edifact-ls-3uzr's epic non-goal.
// 99 segment groups, max nesting depth 6.
//
// Source: https://service.unece.org/trade/untdid/d20a/trmd/govcbr_c.htm
// section 4.3.1 "Segment table". That URL currently returns HTTP 403 from
// Cloudflare when fetched directly; this data was transcribed from the
// Wayback Machine's archived copy instead:
// http://web.archive.org/web/20230608110620/https://service.unece.org/trade/untdid/d20a/trmd/govcbr_c.htm
// -- re-check the direct URL first if cross-checking later, in case the
// block has lifted.
//
// Transcription note: generated in bulk (edifact-ls-13gu) the same way
// as every other message-type schema -- a script parsed the source's
// exact rail-art column positions mechanically and verified the result
// balances (all 99 groups open and close consistently)
// before being transcribed here, rather than reading the ASCII nesting
// by eye. Not individually spot-checked against the raw source the way
// the first several message types in this project were -- see
// edifact-ls-13gu for which representative samples were spot-checked
// for this batch.

func init() {
	RegisterSchema(
		MessageID{Type: "GOVCBR", Version: "D", Release: "20A", Agency: "UN"},
		govcbrD20aSchema,
		"https://service.unece.org/trade/untdid/d20a/trmd/govcbr_c.htm",
	)
}

var govcbrD20aSchema = Schema{Nodes: []SchemaNode{
	{Segment: "BGM", Mandatory: true, MaxRepeat: 1},      // Beginning of message
	{Segment: "DTM", Mandatory: false, MaxRepeat: 9},     // Date/time/period
	{Segment: "MOA", Mandatory: false, MaxRepeat: 1},     // Monetary amount
	{Segment: "IFD", Mandatory: false, MaxRepeat: 9},     // Information detail
	{Segment: "RFF", Mandatory: false, MaxRepeat: 99999}, // Reference
	{Segment: "CUX", Mandatory: false, MaxRepeat: 9},     // Currencies
	{Segment: "FII", Mandatory: false, MaxRepeat: 1},     // Financial institution information
	{Segment: "GPO", Mandatory: false, MaxRepeat: 9},     // Geographical position
	{Segment: "LAN", Mandatory: false, MaxRepeat: 1},     // Language
	{Segment: "MEA", Mandatory: false, MaxRepeat: 9},     // Measurements
	{ // Segment group 1
		Group: []SchemaNode{
			{Segment: "LOC", Mandatory: true, MaxRepeat: 1},  // Place/location identification
			{Segment: "DTM", Mandatory: false, MaxRepeat: 9}, // Date/time/period
			{Segment: "ADR", Mandatory: false, MaxRepeat: 9}, // Address
			{ // Segment group 2
				Group: []SchemaNode{
					{Segment: "GEI", Mandatory: true, MaxRepeat: 1},   // Processing information
					{Segment: "MOA", Mandatory: false, MaxRepeat: 99}, // Monetary amount
					{Segment: "PYT", Mandatory: false, MaxRepeat: 99}, // Payment terms
				},
				Mandatory: false, MaxRepeat: 99,
			},
		},
		Mandatory: false, MaxRepeat: 99,
	},
	{ // Segment group 3
		Group: []SchemaNode{
			{Segment: "GOR", Mandatory: true, MaxRepeat: 1},  // Governmental requirements
			{Segment: "LOC", Mandatory: false, MaxRepeat: 1}, // Place/location identification
			{Segment: "DTM", Mandatory: false, MaxRepeat: 9}, // Date/time/period
			{Segment: "GEI", Mandatory: false, MaxRepeat: 9}, // Processing information
			{ // Segment group 4
				Group: []SchemaNode{
					{Segment: "CTA", Mandatory: true, MaxRepeat: 1},  // Contact information
					{Segment: "COM", Mandatory: false, MaxRepeat: 9}, // Communication contact
				},
				Mandatory: false, MaxRepeat: 9,
			},
		},
		Mandatory: false, MaxRepeat: 9,
	},
	{ // Segment group 5
		Group: []SchemaNode{
			{Segment: "STS", Mandatory: true, MaxRepeat: 1},   // Status
			{Segment: "DTM", Mandatory: false, MaxRepeat: 99}, // Date/time/period
			{ // Segment group 6
				Group: []SchemaNode{
					{Segment: "ERP", Mandatory: true, MaxRepeat: 1},   // Error point details
					{Segment: "ELU", Mandatory: false, MaxRepeat: 99}, // Data element usage details
					{Segment: "ARR", Mandatory: false, MaxRepeat: 99}, // Array information
				},
				Mandatory: false, MaxRepeat: 99,
			},
		},
		Mandatory: false, MaxRepeat: 9,
	},
	{ // Segment group 7
		Group: []SchemaNode{
			{Segment: "NAD", Mandatory: true, MaxRepeat: 1},  // Name and address
			{Segment: "DTM", Mandatory: false, MaxRepeat: 9}, // Date/time/period
			{Segment: "GEI", Mandatory: false, MaxRepeat: 1}, // Processing information
			{Segment: "STS", Mandatory: false, MaxRepeat: 1}, // Status
			{Segment: "IFD", Mandatory: false, MaxRepeat: 9}, // Information detail
			{Segment: "IDE", Mandatory: false, MaxRepeat: 9}, // Identity
			{ // Segment group 8
				Group: []SchemaNode{
					{Segment: "CTA", Mandatory: true, MaxRepeat: 1},  // Contact information
					{Segment: "COM", Mandatory: false, MaxRepeat: 9}, // Communication contact
				},
				Mandatory: false, MaxRepeat: 9,
			},
		},
		Mandatory: false, MaxRepeat: 99,
	},
	{ // Segment group 9
		Group: []SchemaNode{
			{Segment: "DOC", Mandatory: true, MaxRepeat: 1},   // Document/message details
			{Segment: "RFF", Mandatory: false, MaxRepeat: 9},  // Reference
			{Segment: "DTM", Mandatory: false, MaxRepeat: 9},  // Date/time/period
			{Segment: "MOA", Mandatory: false, MaxRepeat: 9},  // Monetary amount
			{Segment: "ALI", Mandatory: false, MaxRepeat: 1},  // Additional information
			{Segment: "QTY", Mandatory: false, MaxRepeat: 99}, // Quantity
			{Segment: "TSR", Mandatory: false, MaxRepeat: 99}, // Transport service requirements
			{ // Segment group 10
				Group: []SchemaNode{
					{Segment: "NAD", Mandatory: true, MaxRepeat: 1},  // Name and address
					{Segment: "IFD", Mandatory: false, MaxRepeat: 1}, // Information detail
					{Segment: "IDE", Mandatory: false, MaxRepeat: 9}, // Identity
					{ // Segment group 11
						Group: []SchemaNode{
							{Segment: "CTA", Mandatory: true, MaxRepeat: 1},  // Contact information
							{Segment: "COM", Mandatory: false, MaxRepeat: 9}, // Communication contact
						},
						Mandatory: false, MaxRepeat: 9,
					},
				},
				Mandatory: false, MaxRepeat: 99,
			},
			{ // Segment group 12
				Group: []SchemaNode{
					{Segment: "LOC", Mandatory: true, MaxRepeat: 1},   // Place/location identification
					{Segment: "DTM", Mandatory: false, MaxRepeat: 99}, // Date/time/period
				},
				Mandatory: false, MaxRepeat: 99,
			},
		},
		Mandatory: false, MaxRepeat: 9999,
	},
	{ // Segment group 13
		Group: []SchemaNode{
			{Segment: "RCS", Mandatory: true, MaxRepeat: 1},  // Requirements and conditions
			{Segment: "FTX", Mandatory: false, MaxRepeat: 9}, // Free text
			{Segment: "DTM", Mandatory: false, MaxRepeat: 9}, // Date/time/period
			{Segment: "GEI", Mandatory: false, MaxRepeat: 1}, // Processing information
			{Segment: "ALI", Mandatory: false, MaxRepeat: 9}, // Additional information
			{ // Segment group 14
				Group: []SchemaNode{
					{Segment: "ERP", Mandatory: true, MaxRepeat: 1},   // Error point details
					{Segment: "ELU", Mandatory: false, MaxRepeat: 99}, // Data element usage details
					{Segment: "ARR", Mandatory: false, MaxRepeat: 99}, // Array information
				},
				Mandatory: false, MaxRepeat: 999,
			},
		},
		Mandatory: false, MaxRepeat: 99,
	},
	{ // Segment group 15
		Group: []SchemaNode{
			{Segment: "AJT", Mandatory: true, MaxRepeat: 1}, // Adjustment details
			{ // Segment group 16
				Group: []SchemaNode{
					{Segment: "ERP", Mandatory: true, MaxRepeat: 1},   // Error point details
					{Segment: "ELU", Mandatory: false, MaxRepeat: 99}, // Data element usage details
					{Segment: "ARR", Mandatory: false, MaxRepeat: 99}, // Array information
				},
				Mandatory: false, MaxRepeat: 999,
			},
		},
		Mandatory: false, MaxRepeat: 99,
	},
	{ // Segment group 17
		Group: []SchemaNode{
			{Segment: "ERC", Mandatory: true, MaxRepeat: 1}, // Application error information
			{ // Segment group 18
				Group: []SchemaNode{
					{Segment: "ERP", Mandatory: true, MaxRepeat: 1},   // Error point details
					{Segment: "ELU", Mandatory: false, MaxRepeat: 99}, // Data element usage details
					{Segment: "ARR", Mandatory: false, MaxRepeat: 99}, // Array information
				},
				Mandatory: false, MaxRepeat: 999,
			},
		},
		Mandatory: false, MaxRepeat: 99,
	},
	{ // Segment group 19
		Group: []SchemaNode{
			{Segment: "INP", Mandatory: true, MaxRepeat: 1},   // Parties and instruction
			{Segment: "DTM", Mandatory: false, MaxRepeat: 9},  // Date/time/period
			{Segment: "LOC", Mandatory: false, MaxRepeat: 9},  // Place/location identification
			{Segment: "CNT", Mandatory: false, MaxRepeat: 99}, // Control total
			{Segment: "RSL", Mandatory: false, MaxRepeat: 9},  // Result
			{ // Segment group 20
				Group: []SchemaNode{
					{Segment: "NAD", Mandatory: true, MaxRepeat: 1},  // Name and address
					{Segment: "IFD", Mandatory: false, MaxRepeat: 1}, // Information detail
					{Segment: "DTM", Mandatory: false, MaxRepeat: 9}, // Date/time/period
					{Segment: "IDE", Mandatory: false, MaxRepeat: 9}, // Identity
					{ // Segment group 21
						Group: []SchemaNode{
							{Segment: "CTA", Mandatory: true, MaxRepeat: 1},  // Contact information
							{Segment: "COM", Mandatory: false, MaxRepeat: 9}, // Communication contact
						},
						Mandatory: false, MaxRepeat: 9,
					},
				},
				Mandatory: false, MaxRepeat: 1,
			},
		},
		Mandatory: false, MaxRepeat: 1,
	},
	{ // Segment group 22
		Group: []SchemaNode{
			{Segment: "GIR", Mandatory: true, MaxRepeat: 1},   // Related identification numbers
			{Segment: "LOC", Mandatory: false, MaxRepeat: 99}, // Place/location identification
			{Segment: "VLI", Mandatory: false, MaxRepeat: 1},  // Value list identification
			{Segment: "DTM", Mandatory: false, MaxRepeat: 99}, // Date/time/period
			{Segment: "MOA", Mandatory: false, MaxRepeat: 99}, // Monetary amount
			{ // Segment group 23
				Group: []SchemaNode{
					{Segment: "NAD", Mandatory: true, MaxRepeat: 1},  // Name and address
					{Segment: "IFD", Mandatory: false, MaxRepeat: 1}, // Information detail
					{Segment: "DTM", Mandatory: false, MaxRepeat: 9}, // Date/time/period
					{Segment: "IDE", Mandatory: false, MaxRepeat: 9}, // Identity
					{ // Segment group 24
						Group: []SchemaNode{
							{Segment: "CTA", Mandatory: true, MaxRepeat: 1},  // Contact information
							{Segment: "COM", Mandatory: false, MaxRepeat: 9}, // Communication contact
						},
						Mandatory: false, MaxRepeat: 9,
					},
				},
				Mandatory: false, MaxRepeat: 9,
			},
			{ // Segment group 25
				Group: []SchemaNode{
					{Segment: "PRC", Mandatory: true, MaxRepeat: 1}, // Process identification
					{Segment: "SEQ", Mandatory: true, MaxRepeat: 1}, // Sequence details
					{Segment: "DTM", Mandatory: true, MaxRepeat: 9}, // Date/time/period
				},
				Mandatory: false, MaxRepeat: 99,
			},
		},
		Mandatory: false, MaxRepeat: 9,
	},
	{ // Segment group 26
		Group: []SchemaNode{
			{Segment: "TAX", Mandatory: true, MaxRepeat: 1},   // Duty/tax/fee details
			{Segment: "DTM", Mandatory: false, MaxRepeat: 9},  // Date/time/period
			{Segment: "ALI", Mandatory: false, MaxRepeat: 9},  // Additional information
			{Segment: "MOA", Mandatory: false, MaxRepeat: 99}, // Monetary amount
			{Segment: "QTY", Mandatory: false, MaxRepeat: 99}, // Quantity
			{ // Segment group 27
				Group: []SchemaNode{
					{Segment: "PAI", Mandatory: true, MaxRepeat: 1},   // Payment instructions
					{Segment: "RFF", Mandatory: false, MaxRepeat: 99}, // Reference
					{Segment: "MOA", Mandatory: false, MaxRepeat: 99}, // Monetary amount
					{Segment: "DTM", Mandatory: false, MaxRepeat: 99}, // Date/time/period
					{ // Segment group 28
						Group: []SchemaNode{
							{Segment: "GIR", Mandatory: true, MaxRepeat: 1},   // Related identification numbers
							{Segment: "LOC", Mandatory: false, MaxRepeat: 99}, // Place/location identification
							{Segment: "VLI", Mandatory: false, MaxRepeat: 1},  // Value list identification
							{Segment: "DTM", Mandatory: false, MaxRepeat: 99}, // Date/time/period
							{Segment: "MOA", Mandatory: false, MaxRepeat: 99}, // Monetary amount
							{ // Segment group 29
								Group: []SchemaNode{
									{Segment: "NAD", Mandatory: true, MaxRepeat: 1},  // Name and address
									{Segment: "IFD", Mandatory: false, MaxRepeat: 1}, // Information detail
									{Segment: "DTM", Mandatory: false, MaxRepeat: 9}, // Date/time/period
									{Segment: "IDE", Mandatory: false, MaxRepeat: 9}, // Identity
									{ // Segment group 30
										Group: []SchemaNode{
											{Segment: "CTA", Mandatory: true, MaxRepeat: 1},  // Contact information
											{Segment: "COM", Mandatory: false, MaxRepeat: 9}, // Communication contact
										},
										Mandatory: false, MaxRepeat: 9,
									},
								},
								Mandatory: false, MaxRepeat: 1,
							},
							{ // Segment group 31
								Group: []SchemaNode{
									{Segment: "PRC", Mandatory: true, MaxRepeat: 1}, // Process identification
									{Segment: "SEQ", Mandatory: true, MaxRepeat: 1}, // Sequence details
									{Segment: "DTM", Mandatory: true, MaxRepeat: 9}, // Date/time/period
								},
								Mandatory: false, MaxRepeat: 99,
							},
						},
						Mandatory: false, MaxRepeat: 9,
					},
				},
				Mandatory: false, MaxRepeat: 9,
			},
		},
		Mandatory: false, MaxRepeat: 99,
	},
	{ // Segment group 32
		Group: []SchemaNode{
			{Segment: "PAC", Mandatory: true, MaxRepeat: 1},   // Package
			{Segment: "SEQ", Mandatory: true, MaxRepeat: 1},   // Sequence details
			{Segment: "PCI", Mandatory: false, MaxRepeat: 99}, // Package identification
			{Segment: "FTX", Mandatory: false, MaxRepeat: 1},  // Free text
			{Segment: "GEI", Mandatory: false, MaxRepeat: 1},  // Processing information
			{Segment: "DIM", Mandatory: false, MaxRepeat: 1},  // Dimensions
			{Segment: "MEA", Mandatory: false, MaxRepeat: 9},  // Measurements
		},
		Mandatory: false, MaxRepeat: 99999,
	},
	{ // Segment group 33
		Group: []SchemaNode{
			{Segment: "TMP", Mandatory: true, MaxRepeat: 1},  // Temperature
			{Segment: "MEA", Mandatory: false, MaxRepeat: 1}, // Measurements
			{Segment: "DGS", Mandatory: false, MaxRepeat: 1}, // Dangerous goods
		},
		Mandatory: false, MaxRepeat: 9,
	},
	{ // Segment group 34
		Group: []SchemaNode{
			{Segment: "TDT", Mandatory: true, MaxRepeat: 1},   // Transport information
			{Segment: "DTM", Mandatory: false, MaxRepeat: 99}, // Date/time/period
			{Segment: "RFF", Mandatory: false, MaxRepeat: 99}, // Reference
			{Segment: "QTY", Mandatory: false, MaxRepeat: 99}, // Quantity
			{Segment: "MEA", Mandatory: false, MaxRepeat: 99}, // Measurements
			{Segment: "MOA", Mandatory: false, MaxRepeat: 99}, // Monetary amount
			{Segment: "POC", Mandatory: false, MaxRepeat: 99}, // Purpose of conveyance call
			{Segment: "DIM", Mandatory: false, MaxRepeat: 99}, // Dimensions
			{Segment: "GDS", Mandatory: false, MaxRepeat: 99}, // Nature of cargo
			{Segment: "STS", Mandatory: false, MaxRepeat: 1},  // Status
			{Segment: "GPO", Mandatory: false, MaxRepeat: 1},  // Geographical position
			{Segment: "GEI", Mandatory: false, MaxRepeat: 9},  // Processing information
			{Segment: "TMD", Mandatory: false, MaxRepeat: 9},  // Transport movement details
			{Segment: "COM", Mandatory: false, MaxRepeat: 99}, // Communication contact
			{ // Segment group 35
				Group: []SchemaNode{
					{Segment: "SEQ", Mandatory: true, MaxRepeat: 1},     // Sequence details
					{Segment: "EVE", Mandatory: true, MaxRepeat: 1},     // Event
					{Segment: "DTM", Mandatory: false, MaxRepeat: 9},    // Date/time/period
					{Segment: "FTX", Mandatory: true, MaxRepeat: 99999}, // Free text
				},
				Mandatory: false, MaxRepeat: 99,
			},
			{ // Segment group 36
				Group: []SchemaNode{
					{Segment: "LOC", Mandatory: true, MaxRepeat: 1},   // Place/location identification
					{Segment: "SEQ", Mandatory: true, MaxRepeat: 1},   // Sequence details
					{Segment: "DTM", Mandatory: false, MaxRepeat: 99}, // Date/time/period
					{Segment: "STS", Mandatory: false, MaxRepeat: 1},  // Status
					{Segment: "FTX", Mandatory: false, MaxRepeat: 9},  // Free text
					{Segment: "GEI", Mandatory: false, MaxRepeat: 9},  // Processing information
					{ // Segment group 37
						Group: []SchemaNode{
							{Segment: "SEQ", Mandatory: true, MaxRepeat: 1},   // Sequence details
							{Segment: "EVE", Mandatory: true, MaxRepeat: 1},   // Event
							{Segment: "DTM", Mandatory: false, MaxRepeat: 99}, // Date/time/period
							{Segment: "FTX", Mandatory: false, MaxRepeat: 99}, // Free text
							{Segment: "GPO", Mandatory: false, MaxRepeat: 1},  // Geographical position
						},
						Mandatory: false, MaxRepeat: 99,
					},
					{ // Segment group 38
						Group: []SchemaNode{
							{Segment: "ADR", Mandatory: true, MaxRepeat: 1},  // Address
							{Segment: "SEQ", Mandatory: true, MaxRepeat: 1},  // Sequence details
							{Segment: "FTX", Mandatory: false, MaxRepeat: 9}, // Free text
							{Segment: "DTM", Mandatory: false, MaxRepeat: 9}, // Date/time/period
							{Segment: "POC", Mandatory: false, MaxRepeat: 9}, // Purpose of conveyance call
							{Segment: "GEI", Mandatory: false, MaxRepeat: 9}, // Processing information
						},
						Mandatory: false, MaxRepeat: 99,
					},
				},
				Mandatory: false, MaxRepeat: 99,
			},
			{ // Segment group 39
				Group: []SchemaNode{
					{Segment: "NAD", Mandatory: true, MaxRepeat: 1},   // Name and address
					{Segment: "SEQ", Mandatory: true, MaxRepeat: 1},   // Sequence details
					{Segment: "ATT", Mandatory: false, MaxRepeat: 1},  // Attribute
					{Segment: "IFD", Mandatory: false, MaxRepeat: 1},  // Information detail
					{Segment: "NAT", Mandatory: false, MaxRepeat: 1},  // Nationality
					{Segment: "DTM", Mandatory: false, MaxRepeat: 99}, // Date/time/period
					{Segment: "LOC", Mandatory: false, MaxRepeat: 99}, // Place/location identification
					{Segment: "EMP", Mandatory: false, MaxRepeat: 99}, // Employment details
					{Segment: "STS", Mandatory: false, MaxRepeat: 1},  // Status
					{Segment: "GEI", Mandatory: false, MaxRepeat: 9},  // Processing information
					{Segment: "RFF", Mandatory: false, MaxRepeat: 9},  // Reference
					{Segment: "IDE", Mandatory: false, MaxRepeat: 9},  // Identity
					{ // Segment group 40
						Group: []SchemaNode{
							{Segment: "CTA", Mandatory: true, MaxRepeat: 1},   // Contact information
							{Segment: "COM", Mandatory: false, MaxRepeat: 99}, // Communication contact
						},
						Mandatory: false, MaxRepeat: 99,
					},
					{ // Segment group 41
						Group: []SchemaNode{
							{Segment: "ADR", Mandatory: true, MaxRepeat: 1},  // Address
							{Segment: "EMP", Mandatory: false, MaxRepeat: 1}, // Employment details
						},
						Mandatory: false, MaxRepeat: 1,
					},
					{ // Segment group 42
						Group: []SchemaNode{
							{Segment: "DOC", Mandatory: true, MaxRepeat: 1},  // Document/message details
							{Segment: "NAD", Mandatory: true, MaxRepeat: 1},  // Name and address
							{Segment: "DTM", Mandatory: false, MaxRepeat: 9}, // Date/time/period
							{Segment: "LOC", Mandatory: false, MaxRepeat: 9}, // Place/location identification
						},
						Mandatory: false, MaxRepeat: 99,
					},
					{ // Segment group 43
						Group: []SchemaNode{
							{Segment: "SEQ", Mandatory: true, MaxRepeat: 1},  // Sequence details
							{Segment: "FTX", Mandatory: false, MaxRepeat: 1}, // Free text
							{Segment: "QTY", Mandatory: false, MaxRepeat: 1}, // Quantity
						},
						Mandatory: false, MaxRepeat: 99,
					},
				},
				Mandatory: false, MaxRepeat: 99999,
			},
			{ // Segment group 44
				Group: []SchemaNode{
					{Segment: "EQD", Mandatory: true, MaxRepeat: 1},   // Equipment details
					{Segment: "SEQ", Mandatory: true, MaxRepeat: 1},   // Sequence details
					{Segment: "QTY", Mandatory: false, MaxRepeat: 1},  // Quantity
					{Segment: "GEI", Mandatory: false, MaxRepeat: 99}, // Processing information
					{Segment: "RFF", Mandatory: false, MaxRepeat: 99}, // Reference
					{Segment: "RCS", Mandatory: false, MaxRepeat: 1},  // Requirements and conditions
					{Segment: "DTM", Mandatory: false, MaxRepeat: 9},  // Date/time/period
					{ // Segment group 45
						Group: []SchemaNode{
							{Segment: "SEL", Mandatory: true, MaxRepeat: 1}, // Seal number
							{Segment: "SEQ", Mandatory: true, MaxRepeat: 1}, // Sequence details
						},
						Mandatory: false, MaxRepeat: 99,
					},
					{ // Segment group 46
						Group: []SchemaNode{
							{Segment: "NAD", Mandatory: true, MaxRepeat: 1},   // Name and address
							{Segment: "IFD", Mandatory: false, MaxRepeat: 1},  // Information detail
							{Segment: "DTM", Mandatory: false, MaxRepeat: 99}, // Date/time/period
							{Segment: "IDE", Mandatory: false, MaxRepeat: 9},  // Identity
							{ // Segment group 47
								Group: []SchemaNode{
									{Segment: "CTA", Mandatory: true, MaxRepeat: 1},   // Contact information
									{Segment: "COM", Mandatory: false, MaxRepeat: 99}, // Communication contact
								},
								Mandatory: false, MaxRepeat: 99,
							},
						},
						Mandatory: false, MaxRepeat: 9,
					},
					{ // Segment group 48
						Group: []SchemaNode{
							{Segment: "LOC", Mandatory: true, MaxRepeat: 1},   // Place/location identification
							{Segment: "DTM", Mandatory: false, MaxRepeat: 99}, // Date/time/period
						},
						Mandatory: false, MaxRepeat: 1,
					},
					{ // Segment group 49
						Group: []SchemaNode{
							{Segment: "SEQ", Mandatory: true, MaxRepeat: 1},   // Sequence details
							{Segment: "EQD", Mandatory: true, MaxRepeat: 1},   // Equipment details
							{Segment: "RFF", Mandatory: false, MaxRepeat: 99}, // Reference
						},
						Mandatory: false, MaxRepeat: 9999,
					},
				},
				Mandatory: false, MaxRepeat: 99999,
			},
			{ // Segment group 50
				Group: []SchemaNode{
					{Segment: "LIN", Mandatory: true, MaxRepeat: 1},   // Line item
					{Segment: "MEA", Mandatory: false, MaxRepeat: 99}, // Measurements
					{Segment: "LOC", Mandatory: false, MaxRepeat: 99}, // Place/location identification
					{Segment: "DTM", Mandatory: false, MaxRepeat: 9},  // Date/time/period
					{Segment: "GEI", Mandatory: false, MaxRepeat: 9},  // Processing information
					{ // Segment group 51
						Group: []SchemaNode{
							{Segment: "GID", Mandatory: true, MaxRepeat: 1},   // Goods item details
							{Segment: "FTX", Mandatory: false, MaxRepeat: 1},  // Free text
							{Segment: "TCC", Mandatory: false, MaxRepeat: 99}, // Charge/rate calculations
							{Segment: "DGS", Mandatory: false, MaxRepeat: 1},  // Dangerous goods
							{Segment: "MEA", Mandatory: false, MaxRepeat: 99}, // Measurements
							{Segment: "LOC", Mandatory: false, MaxRepeat: 99}, // Place/location identification
							{Segment: "CTA", Mandatory: false, MaxRepeat: 99}, // Contact information
						},
						Mandatory: false, MaxRepeat: 99,
					},
					{ // Segment group 52
						Group: []SchemaNode{
							{Segment: "ADR", Mandatory: true, MaxRepeat: 1},  // Address
							{Segment: "SEQ", Mandatory: true, MaxRepeat: 1},  // Sequence details
							{Segment: "DTM", Mandatory: false, MaxRepeat: 9}, // Date/time/period
						},
						Mandatory: false, MaxRepeat: 99,
					},
				},
				Mandatory: false, MaxRepeat: 9999,
			},
			{ // Segment group 53
				Group: []SchemaNode{
					{Segment: "CCI", Mandatory: true, MaxRepeat: 1},  // Characteristic/class id
					{Segment: "CAV", Mandatory: false, MaxRepeat: 1}, // Characteristic value
				},
				Mandatory: false, MaxRepeat: 9999,
			},
			{ // Segment group 54
				Group: []SchemaNode{
					{Segment: "SFI", Mandatory: true, MaxRepeat: 1},   // Safety information
					{Segment: "FTX", Mandatory: false, MaxRepeat: 99}, // Free text
					{ // Segment group 55
						Group: []SchemaNode{
							{Segment: "DOC", Mandatory: true, MaxRepeat: 1},  // Document/message details
							{Segment: "DTM", Mandatory: false, MaxRepeat: 9}, // Date/time/period
							{Segment: "LOC", Mandatory: false, MaxRepeat: 9}, // Place/location identification
							{Segment: "NAD", Mandatory: false, MaxRepeat: 9}, // Name and address
							{Segment: "GEI", Mandatory: false, MaxRepeat: 9}, // Processing information
						},
						Mandatory: false, MaxRepeat: 99,
					},
				},
				Mandatory: false, MaxRepeat: 99,
			},
		},
		Mandatory: false, MaxRepeat: 9,
	},
	{Segment: "UNS", Mandatory: true, MaxRepeat: 1}, // Section control
	{ // Segment group 56
		Group: []SchemaNode{
			{Segment: "SEQ", Mandatory: true, MaxRepeat: 1},   // Sequence details
			{Segment: "ALI", Mandatory: false, MaxRepeat: 99}, // Additional information
			{Segment: "IFD", Mandatory: false, MaxRepeat: 99}, // Information detail
			{Segment: "DTM", Mandatory: false, MaxRepeat: 99}, // Date/time/period
			{Segment: "RFF", Mandatory: false, MaxRepeat: 99}, // Reference
			{ // Segment group 57
				Group: []SchemaNode{
					{Segment: "NAD", Mandatory: true, MaxRepeat: 1},   // Name and address
					{Segment: "IFD", Mandatory: false, MaxRepeat: 1},  // Information detail
					{Segment: "DTM", Mandatory: false, MaxRepeat: 99}, // Date/time/period
					{Segment: "DMS", Mandatory: false, MaxRepeat: 99}, // Document/message summary
					{Segment: "IDE", Mandatory: false, MaxRepeat: 9},  // Identity
					{ // Segment group 58
						Group: []SchemaNode{
							{Segment: "CTA", Mandatory: true, MaxRepeat: 1},   // Contact information
							{Segment: "COM", Mandatory: false, MaxRepeat: 99}, // Communication contact
						},
						Mandatory: false, MaxRepeat: 99,
					},
				},
				Mandatory: false, MaxRepeat: 99,
			},
			{ // Segment group 59
				Group: []SchemaNode{
					{Segment: "LOC", Mandatory: true, MaxRepeat: 1},   // Place/location identification
					{Segment: "DTM", Mandatory: false, MaxRepeat: 99}, // Date/time/period
				},
				Mandatory: false, MaxRepeat: 99,
			},
			{ // Segment group 60
				Group: []SchemaNode{
					{Segment: "DOC", Mandatory: true, MaxRepeat: 1},   // Document/message details
					{Segment: "SEQ", Mandatory: true, MaxRepeat: 1},   // Sequence details
					{Segment: "RFF", Mandatory: false, MaxRepeat: 99}, // Reference
					{Segment: "DTM", Mandatory: false, MaxRepeat: 99}, // Date/time/period
					{Segment: "MOA", Mandatory: false, MaxRepeat: 99}, // Monetary amount
					{Segment: "ALI", Mandatory: false, MaxRepeat: 1},  // Additional information
					{Segment: "QTY", Mandatory: false, MaxRepeat: 99}, // Quantity
					{Segment: "TSR", Mandatory: false, MaxRepeat: 99}, // Transport service requirements
					{ // Segment group 61
						Group: []SchemaNode{
							{Segment: "NAD", Mandatory: true, MaxRepeat: 1},  // Name and address
							{Segment: "IFD", Mandatory: false, MaxRepeat: 1}, // Information detail
							{Segment: "IDE", Mandatory: false, MaxRepeat: 9}, // Identity
							{ // Segment group 62
								Group: []SchemaNode{
									{Segment: "CTA", Mandatory: true, MaxRepeat: 1},   // Contact information
									{Segment: "COM", Mandatory: false, MaxRepeat: 99}, // Communication contact
								},
								Mandatory: false, MaxRepeat: 99,
							},
						},
						Mandatory: false, MaxRepeat: 9,
					},
					{ // Segment group 63
						Group: []SchemaNode{
							{Segment: "LOC", Mandatory: true, MaxRepeat: 1},   // Place/location identification
							{Segment: "DTM", Mandatory: false, MaxRepeat: 99}, // Date/time/period
						},
						Mandatory: false, MaxRepeat: 99,
					},
				},
				Mandatory: false, MaxRepeat: 999,
			},
			{ // Segment group 64
				Group: []SchemaNode{
					{Segment: "GEI", Mandatory: true, MaxRepeat: 1},   // Processing information
					{Segment: "MOA", Mandatory: false, MaxRepeat: 99}, // Monetary amount
					{Segment: "PYT", Mandatory: false, MaxRepeat: 99}, // Payment terms
					{Segment: "CUX", Mandatory: false, MaxRepeat: 1},  // Currencies
					{ // Segment group 65
						Group: []SchemaNode{
							{Segment: "LOC", Mandatory: true, MaxRepeat: 1},   // Place/location identification
							{Segment: "DTM", Mandatory: false, MaxRepeat: 99}, // Date/time/period
						},
						Mandatory: false, MaxRepeat: 1,
					},
				},
				Mandatory: false, MaxRepeat: 99,
			},
			{ // Segment group 66
				Group: []SchemaNode{
					{Segment: "RCS", Mandatory: true, MaxRepeat: 1},   // Requirements and conditions
					{Segment: "FTX", Mandatory: false, MaxRepeat: 9},  // Free text
					{Segment: "DTM", Mandatory: false, MaxRepeat: 99}, // Date/time/period
					{Segment: "GEI", Mandatory: false, MaxRepeat: 99}, // Processing information
					{Segment: "ALI", Mandatory: false, MaxRepeat: 99}, // Additional information
					{ // Segment group 67
						Group: []SchemaNode{
							{Segment: "ERP", Mandatory: true, MaxRepeat: 1},   // Error point details
							{Segment: "ELU", Mandatory: false, MaxRepeat: 99}, // Data element usage details
							{Segment: "ARR", Mandatory: false, MaxRepeat: 99}, // Array information
						},
						Mandatory: false, MaxRepeat: 9,
					},
				},
				Mandatory: false, MaxRepeat: 99,
			},
			{ // Segment group 68
				Group: []SchemaNode{
					{Segment: "INP", Mandatory: true, MaxRepeat: 1},   // Parties and instruction
					{Segment: "DTM", Mandatory: false, MaxRepeat: 9},  // Date/time/period
					{Segment: "LOC", Mandatory: false, MaxRepeat: 99}, // Place/location identification
					{Segment: "CNT", Mandatory: false, MaxRepeat: 99}, // Control total
					{Segment: "RSL", Mandatory: false, MaxRepeat: 9},  // Result
					{ // Segment group 69
						Group: []SchemaNode{
							{Segment: "NAD", Mandatory: true, MaxRepeat: 1},   // Name and address
							{Segment: "IFD", Mandatory: false, MaxRepeat: 1},  // Information detail
							{Segment: "DTM", Mandatory: false, MaxRepeat: 99}, // Date/time/period
							{Segment: "IDE", Mandatory: false, MaxRepeat: 9},  // Identity
							{ // Segment group 70
								Group: []SchemaNode{
									{Segment: "CTA", Mandatory: true, MaxRepeat: 1},   // Contact information
									{Segment: "COM", Mandatory: false, MaxRepeat: 99}, // Communication contact
								},
								Mandatory: false, MaxRepeat: 99,
							},
						},
						Mandatory: false, MaxRepeat: 1,
					},
				},
				Mandatory: false, MaxRepeat: 1,
			},
			{ // Segment group 71
				Group: []SchemaNode{
					{Segment: "TAX", Mandatory: true, MaxRepeat: 1},   // Duty/tax/fee details
					{Segment: "DTM", Mandatory: false, MaxRepeat: 99}, // Date/time/period
					{Segment: "ALI", Mandatory: false, MaxRepeat: 99}, // Additional information
					{Segment: "MOA", Mandatory: false, MaxRepeat: 99}, // Monetary amount
					{Segment: "QTY", Mandatory: false, MaxRepeat: 99}, // Quantity
					{ // Segment group 72
						Group: []SchemaNode{
							{Segment: "PAI", Mandatory: true, MaxRepeat: 1},   // Payment instructions
							{Segment: "RFF", Mandatory: false, MaxRepeat: 99}, // Reference
							{Segment: "MOA", Mandatory: false, MaxRepeat: 99}, // Monetary amount
							{Segment: "DTM", Mandatory: false, MaxRepeat: 99}, // Date/time/period
							{ // Segment group 73
								Group: []SchemaNode{
									{Segment: "GIR", Mandatory: true, MaxRepeat: 1},   // Related identification numbers
									{Segment: "LOC", Mandatory: false, MaxRepeat: 99}, // Place/location identification
									{Segment: "VLI", Mandatory: false, MaxRepeat: 1},  // Value list identification
									{Segment: "DTM", Mandatory: false, MaxRepeat: 99}, // Date/time/period
									{ // Segment group 74
										Group: []SchemaNode{
											{Segment: "NAD", Mandatory: true, MaxRepeat: 1},   // Name and address
											{Segment: "IFD", Mandatory: false, MaxRepeat: 1},  // Information detail
											{Segment: "DTM", Mandatory: false, MaxRepeat: 99}, // Date/time/period
											{Segment: "IDE", Mandatory: false, MaxRepeat: 9},  // Identity
											{ // Segment group 75
												Group: []SchemaNode{
													{Segment: "CTA", Mandatory: true, MaxRepeat: 1},   // Contact information
													{Segment: "COM", Mandatory: false, MaxRepeat: 99}, // Communication contact
												},
												Mandatory: false, MaxRepeat: 99,
											},
										},
										Mandatory: false, MaxRepeat: 1,
									},
								},
								Mandatory: false, MaxRepeat: 9,
							},
						},
						Mandatory: false, MaxRepeat: 9,
					},
				},
				Mandatory: false, MaxRepeat: 99,
			},
			{ // Segment group 76
				Group: []SchemaNode{
					{Segment: "GOR", Mandatory: true, MaxRepeat: 1},  // Governmental requirements
					{Segment: "LOC", Mandatory: false, MaxRepeat: 1}, // Place/location identification
				},
				Mandatory: false, MaxRepeat: 9,
			},
			{ // Segment group 77
				Group: []SchemaNode{
					{Segment: "MEA", Mandatory: true, MaxRepeat: 1},   // Measurements
					{Segment: "QTY", Mandatory: false, MaxRepeat: 99}, // Quantity
				},
				Mandatory: false, MaxRepeat: 9,
			},
			{ // Segment group 78
				Group: []SchemaNode{
					{Segment: "STS", Mandatory: true, MaxRepeat: 1},   // Status
					{Segment: "DTM", Mandatory: false, MaxRepeat: 99}, // Date/time/period
					{ // Segment group 79
						Group: []SchemaNode{
							{Segment: "ERP", Mandatory: true, MaxRepeat: 1},   // Error point details
							{Segment: "ELU", Mandatory: false, MaxRepeat: 99}, // Data element usage details
						},
						Mandatory: false, MaxRepeat: 99,
					},
				},
				Mandatory: false, MaxRepeat: 9,
			},
			{ // Segment group 80
				Group: []SchemaNode{
					{Segment: "TMP", Mandatory: true, MaxRepeat: 1},   // Temperature
					{Segment: "MEA", Mandatory: false, MaxRepeat: 99}, // Measurements
					{Segment: "DGS", Mandatory: false, MaxRepeat: 99}, // Dangerous goods
				},
				Mandatory: false, MaxRepeat: 9,
			},
			{ // Segment group 81
				Group: []SchemaNode{
					{Segment: "TOD", Mandatory: true, MaxRepeat: 1},   // Terms of delivery or transport
					{Segment: "LOC", Mandatory: false, MaxRepeat: 99}, // Place/location identification
					{Segment: "GEI", Mandatory: false, MaxRepeat: 99}, // Processing information
				},
				Mandatory: false, MaxRepeat: 1,
			},
			{ // Segment group 82
				Group: []SchemaNode{
					{Segment: "CNI", Mandatory: true, MaxRepeat: 1},   // Consignment information
					{Segment: "RFF", Mandatory: false, MaxRepeat: 99}, // Reference
					{Segment: "GEI", Mandatory: false, MaxRepeat: 99}, // Processing information
					{Segment: "STS", Mandatory: false, MaxRepeat: 99}, // Status
					{Segment: "CNT", Mandatory: false, MaxRepeat: 1},  // Control total
					{Segment: "QTY", Mandatory: false, MaxRepeat: 99}, // Quantity
					{Segment: "MOA", Mandatory: false, MaxRepeat: 99}, // Monetary amount
					{Segment: "IFD", Mandatory: false, MaxRepeat: 99}, // Information detail
					{Segment: "TOD", Mandatory: false, MaxRepeat: 9},  // Terms of delivery or transport
					{ // Segment group 83
						Group: []SchemaNode{
							{Segment: "NAD", Mandatory: true, MaxRepeat: 1},   // Name and address
							{Segment: "IFD", Mandatory: false, MaxRepeat: 9},  // Information detail
							{Segment: "DTM", Mandatory: false, MaxRepeat: 99}, // Date/time/period
							{Segment: "IDE", Mandatory: false, MaxRepeat: 9},  // Identity
							{ // Segment group 84
								Group: []SchemaNode{
									{Segment: "CTA", Mandatory: true, MaxRepeat: 1},   // Contact information
									{Segment: "COM", Mandatory: false, MaxRepeat: 99}, // Communication contact
								},
								Mandatory: false, MaxRepeat: 99,
							},
						},
						Mandatory: false, MaxRepeat: 99,
					},
					{ // Segment group 85
						Group: []SchemaNode{
							{Segment: "LOC", Mandatory: true, MaxRepeat: 1},   // Place/location identification
							{Segment: "DTM", Mandatory: false, MaxRepeat: 99}, // Date/time/period
							{Segment: "ARR", Mandatory: false, MaxRepeat: 1},  // Array information
							{Segment: "ADR", Mandatory: false, MaxRepeat: 9},  // Address
							{Segment: "TPL", Mandatory: false, MaxRepeat: 9},  // Transport placement
							{ // Segment group 86
								Group: []SchemaNode{
									{Segment: "GEI", Mandatory: true, MaxRepeat: 1},   // Processing information
									{Segment: "MOA", Mandatory: false, MaxRepeat: 99}, // Monetary amount
									{Segment: "PYT", Mandatory: false, MaxRepeat: 99}, // Payment terms
								},
								Mandatory: false, MaxRepeat: 9,
							},
						},
						Mandatory: false, MaxRepeat: 99,
					},
					{ // Segment group 87
						Group: []SchemaNode{
							{Segment: "DOC", Mandatory: true, MaxRepeat: 1},   // Document/message details
							{Segment: "RFF", Mandatory: false, MaxRepeat: 99}, // Reference
							{Segment: "DTM", Mandatory: false, MaxRepeat: 99}, // Date/time/period
							{Segment: "MOA", Mandatory: false, MaxRepeat: 99}, // Monetary amount
							{Segment: "ALI", Mandatory: false, MaxRepeat: 1},  // Additional information
							{Segment: "QTY", Mandatory: false, MaxRepeat: 99}, // Quantity
							{Segment: "TSR", Mandatory: false, MaxRepeat: 99}, // Transport service requirements
							{ // Segment group 88
								Group: []SchemaNode{
									{Segment: "NAD", Mandatory: true, MaxRepeat: 1},  // Name and address
									{Segment: "IFD", Mandatory: false, MaxRepeat: 1}, // Information detail
									{Segment: "IDE", Mandatory: false, MaxRepeat: 9}, // Identity
									{ // Segment group 89
										Group: []SchemaNode{
											{Segment: "CTA", Mandatory: true, MaxRepeat: 1},   // Contact information
											{Segment: "COM", Mandatory: false, MaxRepeat: 99}, // Communication contact
										},
										Mandatory: false, MaxRepeat: 99,
									},
								},
								Mandatory: false, MaxRepeat: 9,
							},
							{ // Segment group 90
								Group: []SchemaNode{
									{Segment: "LOC", Mandatory: true, MaxRepeat: 1},   // Place/location identification
									{Segment: "DTM", Mandatory: false, MaxRepeat: 99}, // Date/time/period
								},
								Mandatory: false, MaxRepeat: 99,
							},
						},
						Mandatory: false, MaxRepeat: 99,
					},
					{ // Segment group 91
						Group: []SchemaNode{
							{Segment: "RCS", Mandatory: true, MaxRepeat: 1},   // Requirements and conditions
							{Segment: "FTX", Mandatory: false, MaxRepeat: 9},  // Free text
							{Segment: "DTM", Mandatory: false, MaxRepeat: 99}, // Date/time/period
							{Segment: "GEI", Mandatory: false, MaxRepeat: 99}, // Processing information
							{Segment: "ALI", Mandatory: false, MaxRepeat: 99}, // Additional information
							{ // Segment group 92
								Group: []SchemaNode{
									{Segment: "ERP", Mandatory: true, MaxRepeat: 1},   // Error point details
									{Segment: "ELU", Mandatory: false, MaxRepeat: 99}, // Data element usage details
									{Segment: "ARR", Mandatory: false, MaxRepeat: 99}, // Array information
								},
								Mandatory: false, MaxRepeat: 9,
							},
						},
						Mandatory: false, MaxRepeat: 99,
					},
					{ // Segment group 93
						Group: []SchemaNode{
							{Segment: "INP", Mandatory: true, MaxRepeat: 1},   // Parties and instruction
							{Segment: "DTM", Mandatory: false, MaxRepeat: 9},  // Date/time/period
							{Segment: "LOC", Mandatory: false, MaxRepeat: 99}, // Place/location identification
							{Segment: "CNT", Mandatory: false, MaxRepeat: 99}, // Control total
							{ // Segment group 94
								Group: []SchemaNode{
									{Segment: "NAD", Mandatory: true, MaxRepeat: 1},   // Name and address
									{Segment: "IFD", Mandatory: false, MaxRepeat: 1},  // Information detail
									{Segment: "DTM", Mandatory: false, MaxRepeat: 99}, // Date/time/period
									{Segment: "IDE", Mandatory: false, MaxRepeat: 9},  // Identity
									{ // Segment group 95
										Group: []SchemaNode{
											{Segment: "CTA", Mandatory: true, MaxRepeat: 1},   // Contact information
											{Segment: "COM", Mandatory: false, MaxRepeat: 99}, // Communication contact
										},
										Mandatory: false, MaxRepeat: 99,
									},
								},
								Mandatory: false, MaxRepeat: 1,
							},
						},
						Mandatory: false, MaxRepeat: 1,
					},
					{ // Segment group 96
						Group: []SchemaNode{
							{Segment: "PAC", Mandatory: true, MaxRepeat: 1},   // Package
							{Segment: "PCI", Mandatory: false, MaxRepeat: 99}, // Package identification
							{Segment: "FTX", Mandatory: false, MaxRepeat: 99}, // Free text
							{Segment: "GEI", Mandatory: false, MaxRepeat: 99}, // Processing information
							{Segment: "DIM", Mandatory: false, MaxRepeat: 99}, // Dimensions
							{Segment: "MEA", Mandatory: false, MaxRepeat: 99}, // Measurements
						},
						Mandatory: false, MaxRepeat: 99,
					},
					{ // Segment group 97
						Group: []SchemaNode{
							{Segment: "TMP", Mandatory: true, MaxRepeat: 1},   // Temperature
							{Segment: "MEA", Mandatory: false, MaxRepeat: 99}, // Measurements
							{Segment: "DGS", Mandatory: false, MaxRepeat: 99}, // Dangerous goods
						},
						Mandatory: false, MaxRepeat: 99,
					},
					{ // Segment group 98
						Group: []SchemaNode{
							{Segment: "EQD", Mandatory: true, MaxRepeat: 1},   // Equipment details
							{Segment: "SEQ", Mandatory: true, MaxRepeat: 1},   // Sequence details
							{Segment: "QTY", Mandatory: false, MaxRepeat: 1},  // Quantity
							{Segment: "GEI", Mandatory: false, MaxRepeat: 99}, // Processing information
							{ // Segment group 99
								Group: []SchemaNode{
									{Segment: "SEL", Mandatory: true, MaxRepeat: 1}, // Seal number
									{Segment: "SEQ", Mandatory: true, MaxRepeat: 1}, // Sequence details
								},
								Mandatory: false, MaxRepeat: 99,
							},
							{Segment: "----", Mandatory: false, MaxRepeat: 99}, // Segment group 100------------------
							{Segment: "NAD", Mandatory: true, MaxRepeat: 1},    // Name and address
							{Segment: "IFD", Mandatory: false, MaxRepeat: 1},   // Information detail
							{Segment: "DTM", Mandatory: false, MaxRepeat: 99},  // Date/time/period
							{Segment: "IDE", Mandatory: false, MaxRepeat: 9},   // Identity
							{Segment: "----", Mandatory: false, MaxRepeat: 99}, // Segment group 101------------------
							{Segment: "CTA", Mandatory: true, MaxRepeat: 1},    // Contact information
							{Segment: "COM", Mandatory: false, MaxRepeat: 99},  // Communication contact
						},
						Mandatory: false, MaxRepeat: 99,
					},
					{Segment: "----", Mandatory: false, MaxRepeat: 99}, // Segment group 102------------------
					{Segment: "TDT", Mandatory: true, MaxRepeat: 1},    // Transport information
					{Segment: "DTM", Mandatory: false, MaxRepeat: 99},  // Date/time/period
					{Segment: "RFF", Mandatory: false, MaxRepeat: 99},  // Reference
					{Segment: "QTY", Mandatory: false, MaxRepeat: 99},  // Quantity
					{Segment: "MEA", Mandatory: false, MaxRepeat: 99},  // Measurements
					{Segment: "MOA", Mandatory: false, MaxRepeat: 99},  // Monetary amount
					{Segment: "POC", Mandatory: false, MaxRepeat: 99},  // Purpose of conveyance call
					{Segment: "DIM", Mandatory: false, MaxRepeat: 99},  // Dimensions
					{Segment: "FTX", Mandatory: false, MaxRepeat: 99},  // Free text
					{Segment: "GDS", Mandatory: false, MaxRepeat: 99},  // Nature of cargo
					{Segment: "STS", Mandatory: false, MaxRepeat: 1},   // Status
					{Segment: "GPO", Mandatory: false, MaxRepeat: 1},   // Geographical position
					{Segment: "GEI", Mandatory: false, MaxRepeat: 9},   // Processing information
					{Segment: "----", Mandatory: false, MaxRepeat: 99}, // Segment group 103------------------
					{Segment: "LOC", Mandatory: true, MaxRepeat: 1},    // Place/location identification
					{Segment: "SEQ", Mandatory: true, MaxRepeat: 1},    // Sequence details
					{Segment: "DTM", Mandatory: false, MaxRepeat: 99},  // Date/time/period
					{Segment: "----", Mandatory: false, MaxRepeat: 99}, // Segment group 104------------------
					{Segment: "NAD", Mandatory: true, MaxRepeat: 1},    // Name and address
					{Segment: "IFD", Mandatory: false, MaxRepeat: 1},   // Information detail
					{Segment: "NAT", Mandatory: false, MaxRepeat: 1},   // Nationality
					{Segment: "DTM", Mandatory: false, MaxRepeat: 99},  // Date/time/period
					{Segment: "IDE", Mandatory: false, MaxRepeat: 9},   // Identity
					{Segment: "----", Mandatory: false, MaxRepeat: 99}, // Segment group 105------------------
					{Segment: "CTA", Mandatory: true, MaxRepeat: 1},    // Contact information
					{Segment: "COM", Mandatory: false, MaxRepeat: 99},  // Communication contact
					{Segment: "----", Mandatory: false, MaxRepeat: 99}, // Segment group 106------------------
					{Segment: "EQD", Mandatory: true, MaxRepeat: 1},    // Equipment details
					{Segment: "SEQ", Mandatory: true, MaxRepeat: 1},    // Sequence details
					{Segment: "QTY", Mandatory: false, MaxRepeat: 1},   // Quantity
					{Segment: "GEI", Mandatory: false, MaxRepeat: 99},  // Processing information
					{Segment: "----", Mandatory: false, MaxRepeat: 99}, // Segment group 107------------------
					{Segment: "SEL", Mandatory: true, MaxRepeat: 1},    // Seal number
					{Segment: "SEQ", Mandatory: true, MaxRepeat: 1},    // Sequence details
					{Segment: "----", Mandatory: false, MaxRepeat: 9},  // Segment group 108------------------
					{Segment: "NAD", Mandatory: true, MaxRepeat: 1},    // Name and address
					{Segment: "IFD", Mandatory: false, MaxRepeat: 1},   // Information detail
					{Segment: "DTM", Mandatory: false, MaxRepeat: 99},  // Date/time/period
					{Segment: "IDE", Mandatory: false, MaxRepeat: 9},   // Identity
					{Segment: "----", Mandatory: false, MaxRepeat: 99}, // Segment group 109------------------
					{Segment: "CTA", Mandatory: true, MaxRepeat: 1},    // Contact information
					{Segment: "COM", Mandatory: false, MaxRepeat: 99},  // Communication contact
				},
				Mandatory: false, MaxRepeat: 9999,
			},
			{Segment: "----", Mandatory: false, MaxRepeat: 99},   // Segment group 110------------------
			{Segment: "ARD", Mandatory: true, MaxRepeat: 1},      // Monetary amount function
			{Segment: "MOA", Mandatory: false, MaxRepeat: 99},    // Monetary amount
			{Segment: "GEI", Mandatory: false, MaxRepeat: 99},    // Processing information
			{Segment: "PCD", Mandatory: false, MaxRepeat: 99},    // Percentage details
			{Segment: "----", Mandatory: false, MaxRepeat: 9999}, // Segment group 111------------------
			{Segment: "LIN", Mandatory: true, MaxRepeat: 1},      // Line item
			{Segment: "MOA", Mandatory: false, MaxRepeat: 99},    // Monetary amount
			{Segment: "IFD", Mandatory: false, MaxRepeat: 99},    // Information detail
			{Segment: "STS", Mandatory: false, MaxRepeat: 9},     // Status
			{Segment: "ALI", Mandatory: false, MaxRepeat: 99},    // Additional information
			{Segment: "RFF", Mandatory: false, MaxRepeat: 9},     // Reference
			{Segment: "----", Mandatory: false, MaxRepeat: 99},   // Segment group 112------------------
			{Segment: "NAD", Mandatory: true, MaxRepeat: 1},      // Name and address
			{Segment: "IFD", Mandatory: false, MaxRepeat: 9},     // Information detail
			{Segment: "DTM", Mandatory: false, MaxRepeat: 99},    // Date/time/period
			{Segment: "IDE", Mandatory: false, MaxRepeat: 9},     // Identity
			{Segment: "----", Mandatory: false, MaxRepeat: 99},   // Segment group 113------------------
			{Segment: "CTA", Mandatory: true, MaxRepeat: 1},      // Contact information
			{Segment: "COM", Mandatory: false, MaxRepeat: 99},    // Communication contact
			{Segment: "----", Mandatory: false, MaxRepeat: 99},   // Segment group 114------------------
			{Segment: "LOC", Mandatory: true, MaxRepeat: 1},      // Place/location identification
			{Segment: "DTM", Mandatory: false, MaxRepeat: 99},    // Date/time/period
			{Segment: "FTX", Mandatory: false, MaxRepeat: 9},     // Free text
			{Segment: "ADR", Mandatory: false, MaxRepeat: 9},     // Address
			{Segment: "----", Mandatory: false, MaxRepeat: 99},   // Segment group 115------------------
			{Segment: "DOC", Mandatory: true, MaxRepeat: 1},      // Document/message details
			{Segment: "RFF", Mandatory: false, MaxRepeat: 99},    // Reference
			{Segment: "DTM", Mandatory: false, MaxRepeat: 99},    // Date/time/period
			{Segment: "MOA", Mandatory: false, MaxRepeat: 99},    // Monetary amount
			{Segment: "ALI", Mandatory: false, MaxRepeat: 1},     // Additional information
			{Segment: "QTY", Mandatory: false, MaxRepeat: 99},    // Quantity
			{Segment: "TSR", Mandatory: false, MaxRepeat: 99},    // Transport service requirements
			{Segment: "GEI", Mandatory: false, MaxRepeat: 9},     // Processing information
			{Segment: "----", Mandatory: false, MaxRepeat: 99},   // Segment group 116------------------
			{Segment: "NAD", Mandatory: true, MaxRepeat: 1},      // Name and address
			{Segment: "IFD", Mandatory: false, MaxRepeat: 1},     // Information detail
			{Segment: "IDE", Mandatory: false, MaxRepeat: 9},     // Identity
			{Segment: "----", Mandatory: false, MaxRepeat: 99},   // Segment group 117------------------
			{Segment: "CTA", Mandatory: true, MaxRepeat: 1},      // Contact information
			{Segment: "COM", Mandatory: false, MaxRepeat: 99},    // Communication contact
			{Segment: "----", Mandatory: false, MaxRepeat: 99},   // Segment group 118------------------
			{Segment: "LOC", Mandatory: true, MaxRepeat: 1},      // Place/location identification
			{Segment: "DTM", Mandatory: false, MaxRepeat: 99},    // Date/time/period
			{Segment: "----", Mandatory: false, MaxRepeat: 99},   // Segment group 119------------------
			{Segment: "RCS", Mandatory: true, MaxRepeat: 1},      // Requirements and conditions
			{Segment: "FTX", Mandatory: false, MaxRepeat: 9},     // Free text
			{Segment: "DTM", Mandatory: false, MaxRepeat: 99},    // Date/time/period
			{Segment: "GEI", Mandatory: false, MaxRepeat: 99},    // Processing information
			{Segment: "ALI", Mandatory: false, MaxRepeat: 99},    // Additional information
			{Segment: "----", Mandatory: false, MaxRepeat: 99},   // Segment group 120------------------
			{Segment: "ERP", Mandatory: true, MaxRepeat: 1},      // Error point details
			{Segment: "ELU", Mandatory: false, MaxRepeat: 99},    // Data element usage details
			{Segment: "ARR", Mandatory: false, MaxRepeat: 99},    // Array information
			{Segment: "----", Mandatory: false, MaxRepeat: 1},    // Segment group 121------------------
			{Segment: "INP", Mandatory: true, MaxRepeat: 1},      // Parties and instruction
			{Segment: "DTM", Mandatory: false, MaxRepeat: 9},     // Date/time/period
			{Segment: "LOC", Mandatory: false, MaxRepeat: 99},    // Place/location identification
			{Segment: "CNT", Mandatory: false, MaxRepeat: 9},     // Control total
			{Segment: "----", Mandatory: false, MaxRepeat: 1},    // Segment group 122------------------
			{Segment: "NAD", Mandatory: true, MaxRepeat: 1},      // Name and address
			{Segment: "IFD", Mandatory: false, MaxRepeat: 1},     // Information detail
			{Segment: "DTM", Mandatory: false, MaxRepeat: 99},    // Date/time/period
			{Segment: "IDE", Mandatory: false, MaxRepeat: 9},     // Identity
			{Segment: "----", Mandatory: false, MaxRepeat: 99},   // Segment group 123------------------
			{Segment: "CTA", Mandatory: true, MaxRepeat: 1},      // Contact information
			{Segment: "COM", Mandatory: false, MaxRepeat: 99},    // Communication contact
			{Segment: "----", Mandatory: false, MaxRepeat: 99},   // Segment group 124------------------
			{Segment: "MEA", Mandatory: true, MaxRepeat: 1},      // Measurements
			{Segment: "QTY", Mandatory: false, MaxRepeat: 99},    // Quantity
			{Segment: "----", Mandatory: false, MaxRepeat: 999},  // Segment group 125------------------
			{Segment: "PAC", Mandatory: true, MaxRepeat: 1},      // Package
			{Segment: "SEQ", Mandatory: true, MaxRepeat: 1},      // Sequence details
			{Segment: "PCI", Mandatory: false, MaxRepeat: 99},    // Package identification
			{Segment: "FTX", Mandatory: false, MaxRepeat: 99},    // Free text
			{Segment: "GEI", Mandatory: false, MaxRepeat: 99},    // Processing information
			{Segment: "DIM", Mandatory: false, MaxRepeat: 99},    // Dimensions
			{Segment: "MEA", Mandatory: false, MaxRepeat: 99},    // Measurements
			{Segment: "----", Mandatory: false, MaxRepeat: 99},   // Segment group 126------------------
			{Segment: "ARD", Mandatory: true, MaxRepeat: 1},      // Monetary amount function
			{Segment: "MOA", Mandatory: false, MaxRepeat: 99},    // Monetary amount
			{Segment: "GEI", Mandatory: false, MaxRepeat: 99},    // Processing information
			{Segment: "PCD", Mandatory: false, MaxRepeat: 99},    // Percentage details
			{Segment: "----", Mandatory: false, MaxRepeat: 99},   // Segment group 127------------------
			{Segment: "GID", Mandatory: true, MaxRepeat: 1},      // Goods item details
			{Segment: "IMD", Mandatory: false, MaxRepeat: 99},    // Item description
			{Segment: "APP", Mandatory: false, MaxRepeat: 99},    // Applicability
			{Segment: "FTX", Mandatory: false, MaxRepeat: 99},    // Free text
			{Segment: "RFF", Mandatory: false, MaxRepeat: 99},    // Reference
			{Segment: "DTM", Mandatory: false, MaxRepeat: 99},    // Date/time/period
			{Segment: "GIR", Mandatory: false, MaxRepeat: 99},    // Related identification numbers
			{Segment: "GIN", Mandatory: false, MaxRepeat: 99},    // Goods identity number
			{Segment: "GEI", Mandatory: false, MaxRepeat: 99},    // Processing information
			{Segment: "MEA", Mandatory: false, MaxRepeat: 99},    // Measurements
			{Segment: "MOA", Mandatory: false, MaxRepeat: 99},    // Monetary amount
			{Segment: "CDI", Mandatory: false, MaxRepeat: 99},    // Physical or logical state
			{Segment: "PGI", Mandatory: false, MaxRepeat: 99},    // Product group information
			{Segment: "TCC", Mandatory: false, MaxRepeat: 99},    // Charge/rate calculations
			{Segment: "CNT", Mandatory: false, MaxRepeat: 99},    // Control total
			{Segment: "IFD", Mandatory: false, MaxRepeat: 99},    // Information detail
			{Segment: "TDT", Mandatory: false, MaxRepeat: 9},     // Transport information
			{Segment: "----", Mandatory: false, MaxRepeat: 99},   // Segment group 128------------------
			{Segment: "LOC", Mandatory: true, MaxRepeat: 1},      // Place/location identification
			{Segment: "SEQ", Mandatory: true, MaxRepeat: 1},      // Sequence details
			{Segment: "DTM", Mandatory: false, MaxRepeat: 99},    // Date/time/period
			{Segment: "----", Mandatory: false, MaxRepeat: 99},   // Segment group 129------------------
			{Segment: "NAD", Mandatory: true, MaxRepeat: 1},      // Name and address
			{Segment: "IFD", Mandatory: false, MaxRepeat: 9},     // Information detail
			{Segment: "DTM", Mandatory: false, MaxRepeat: 99},    // Date/time/period
			{Segment: "IDE", Mandatory: false, MaxRepeat: 9},     // Identity
			{Segment: "----", Mandatory: false, MaxRepeat: 99},   // Segment group 130------------------
			{Segment: "CTA", Mandatory: true, MaxRepeat: 1},      // Contact information
			{Segment: "COM", Mandatory: false, MaxRepeat: 99},    // Communication contact
			{Segment: "----", Mandatory: false, MaxRepeat: 99},   // Segment group 131------------------
			{Segment: "DOC", Mandatory: true, MaxRepeat: 1},      // Document/message details
			{Segment: "RFF", Mandatory: false, MaxRepeat: 99},    // Reference
			{Segment: "DTM", Mandatory: false, MaxRepeat: 99},    // Date/time/period
			{Segment: "MOA", Mandatory: false, MaxRepeat: 99},    // Monetary amount
			{Segment: "ALI", Mandatory: false, MaxRepeat: 1},     // Additional information
			{Segment: "QTY", Mandatory: false, MaxRepeat: 99},    // Quantity
			{Segment: "TSR", Mandatory: false, MaxRepeat: 99},    // Transport service requirements
			{Segment: "GEI", Mandatory: false, MaxRepeat: 9},     // Processing information
			{Segment: "GIN", Mandatory: false, MaxRepeat: 99},    // Goods identity number
			{Segment: "GIR", Mandatory: false, MaxRepeat: 99},    // Related identification numbers
			{Segment: "IMD", Mandatory: false, MaxRepeat: 99},    // Item description
			{Segment: "----", Mandatory: false, MaxRepeat: 99},   // Segment group 132------------------
			{Segment: "NAD", Mandatory: true, MaxRepeat: 1},      // Name and address
			{Segment: "IFD", Mandatory: false, MaxRepeat: 1},     // Information detail
			{Segment: "IDE", Mandatory: false, MaxRepeat: 9},     // Identity
			{Segment: "----", Mandatory: false, MaxRepeat: 99},   // Segment group 133------------------
			{Segment: "CTA", Mandatory: true, MaxRepeat: 1},      // Contact information
			{Segment: "COM", Mandatory: false, MaxRepeat: 99},    // Communication contact
			{Segment: "----", Mandatory: false, MaxRepeat: 99},   // Segment group 134------------------
			{Segment: "LOC", Mandatory: true, MaxRepeat: 1},      // Place/location identification
			{Segment: "DTM", Mandatory: false, MaxRepeat: 99},    // Date/time/period
			{Segment: "----", Mandatory: false, MaxRepeat: 99},   // Segment group 135------------------
			{Segment: "RCS", Mandatory: true, MaxRepeat: 1},      // Requirements and conditions
			{Segment: "FTX", Mandatory: false, MaxRepeat: 9},     // Free text
			{Segment: "DTM", Mandatory: false, MaxRepeat: 99},    // Date/time/period
			{Segment: "GEI", Mandatory: false, MaxRepeat: 99},    // Processing information
			{Segment: "ALI", Mandatory: false, MaxRepeat: 99},    // Additional information
			{Segment: "----", Mandatory: false, MaxRepeat: 99},   // Segment group 136------------------
			{Segment: "ERP", Mandatory: true, MaxRepeat: 1},      // Error point details
			{Segment: "ELU", Mandatory: false, MaxRepeat: 99},    // Data element usage details
			{Segment: "ARR", Mandatory: false, MaxRepeat: 99},    // Array information
			{Segment: "----", Mandatory: false, MaxRepeat: 999},  // Segment group 137------------------
			{Segment: "PAC", Mandatory: true, MaxRepeat: 1},      // Package
			{Segment: "PCI", Mandatory: false, MaxRepeat: 99},    // Package identification
			{Segment: "FTX", Mandatory: false, MaxRepeat: 99},    // Free text
			{Segment: "GEI", Mandatory: false, MaxRepeat: 99},    // Processing information
			{Segment: "DIM", Mandatory: false, MaxRepeat: 99},    // Dimensions
			{Segment: "MEA", Mandatory: false, MaxRepeat: 99},    // Measurements
			{Segment: "----", Mandatory: false, MaxRepeat: 99},   // Segment group 138------------------
			{Segment: "COD", Mandatory: true, MaxRepeat: 1},      // Component details
			{Segment: "QTY", Mandatory: false, MaxRepeat: 99},    // Quantity
			{Segment: "PCD", Mandatory: false, MaxRepeat: 99},    // Percentage details
			{Segment: "GEI", Mandatory: false, MaxRepeat: 99},    // Processing information
			{Segment: "LOC", Mandatory: false, MaxRepeat: 9},     // Place/location identification
			{Segment: "----", Mandatory: false, MaxRepeat: 1},    // Segment group 139------------------
			{Segment: "INP", Mandatory: true, MaxRepeat: 1},      // Parties and instruction
			{Segment: "DTM", Mandatory: false, MaxRepeat: 9},     // Date/time/period
			{Segment: "LOC", Mandatory: false, MaxRepeat: 99},    // Place/location identification
			{Segment: "CNT", Mandatory: false, MaxRepeat: 99},    // Control total
			{Segment: "----", Mandatory: false, MaxRepeat: 1},    // Segment group 140------------------
			{Segment: "NAD", Mandatory: true, MaxRepeat: 1},      // Name and address
			{Segment: "IFD", Mandatory: false, MaxRepeat: 1},     // Information detail
			{Segment: "DTM", Mandatory: false, MaxRepeat: 99},    // Date/time/period
			{Segment: "IDE", Mandatory: false, MaxRepeat: 9},     // Identity
			{Segment: "----", Mandatory: false, MaxRepeat: 99},   // Segment group 141------------------
			{Segment: "CTA", Mandatory: true, MaxRepeat: 1},      // Contact information
			{Segment: "COM", Mandatory: false, MaxRepeat: 99},    // Communication contact
			{Segment: "----", Mandatory: false, MaxRepeat: 99},   // Segment group 142------------------
			{Segment: "TAX", Mandatory: true, MaxRepeat: 1},      // Duty/tax/fee details
			{Segment: "DTM", Mandatory: false, MaxRepeat: 99},    // Date/time/period
			{Segment: "ALI", Mandatory: false, MaxRepeat: 99},    // Additional information
			{Segment: "MOA", Mandatory: false, MaxRepeat: 99},    // Monetary amount
			{Segment: "QTY", Mandatory: false, MaxRepeat: 99},    // Quantity
			{Segment: "TCC", Mandatory: false, MaxRepeat: 9},     // Charge/rate calculations
			{Segment: "----", Mandatory: false, MaxRepeat: 99},   // Segment group 143------------------
			{Segment: "PAI", Mandatory: true, MaxRepeat: 1},      // Payment instructions
			{Segment: "RFF", Mandatory: false, MaxRepeat: 99},    // Reference
			{Segment: "MOA", Mandatory: false, MaxRepeat: 99},    // Monetary amount
			{Segment: "DTM", Mandatory: false, MaxRepeat: 99},    // Date/time/period
			{Segment: "----", Mandatory: false, MaxRepeat: 99},   // Segment group 144------------------
			{Segment: "GIR", Mandatory: true, MaxRepeat: 1},      // Related identification numbers
			{Segment: "LOC", Mandatory: false, MaxRepeat: 99},    // Place/location identification
			{Segment: "VLI", Mandatory: false, MaxRepeat: 1},     // Value list identification
			{Segment: "DTM", Mandatory: false, MaxRepeat: 99},    // Date/time/period
			{Segment: "MOA", Mandatory: false, MaxRepeat: 99},    // Monetary amount
			{Segment: "----", Mandatory: false, MaxRepeat: 99},   // Segment group 145------------------
			{Segment: "NAD", Mandatory: true, MaxRepeat: 1},      // Name and address
			{Segment: "IFD", Mandatory: false, MaxRepeat: 1},     // Information detail
			{Segment: "DTM", Mandatory: false, MaxRepeat: 99},    // Date/time/period
			{Segment: "IDE", Mandatory: false, MaxRepeat: 9},     // Identity
			{Segment: "----", Mandatory: false, MaxRepeat: 99},   // Segment group 146------------------
			{Segment: "CTA", Mandatory: true, MaxRepeat: 1},      // Contact information
			{Segment: "COM", Mandatory: false, MaxRepeat: 99},    // Communication contact
			{Segment: "----", Mandatory: false, MaxRepeat: 9},    // Segment group 147------------------
			{Segment: "ARR", Mandatory: true, MaxRepeat: 1},      // Array information
			{Segment: "DTM", Mandatory: true, MaxRepeat: 9},      // Date/time/period
			{Segment: "----", Mandatory: false, MaxRepeat: 99},   // Segment group 148------------------
			{Segment: "DLI", Mandatory: true, MaxRepeat: 1},      // Document line identification
			{Segment: "MOA", Mandatory: false, MaxRepeat: 99},    // Monetary amount
			{Segment: "----", Mandatory: false, MaxRepeat: 99},   // Segment group 149------------------
			{Segment: "STS", Mandatory: true, MaxRepeat: 1},      // Status
			{Segment: "DTM", Mandatory: false, MaxRepeat: 99},    // Date/time/period
			{Segment: "----", Mandatory: false, MaxRepeat: 99},   // Segment group 150------------------
			{Segment: "ERP", Mandatory: true, MaxRepeat: 1},      // Error point details
			{Segment: "ELU", Mandatory: false, MaxRepeat: 99},    // Data element usage details
			{Segment: "----", Mandatory: false, MaxRepeat: 99},   // Segment group 151------------------
			{Segment: "TMP", Mandatory: true, MaxRepeat: 1},      // Temperature
			{Segment: "MEA", Mandatory: false, MaxRepeat: 99},    // Measurements
			{Segment: "DGS", Mandatory: false, MaxRepeat: 99},    // Dangerous goods
			{Segment: "----", Mandatory: false, MaxRepeat: 99},   // Segment group 152------------------
			{Segment: "EQD", Mandatory: true, MaxRepeat: 1},      // Equipment details
			{Segment: "SEQ", Mandatory: true, MaxRepeat: 1},      // Sequence details
			{Segment: "----", Mandatory: false, MaxRepeat: 99},   // Segment group 153------------------
			{Segment: "PRC", Mandatory: true, MaxRepeat: 1},      // Process identification
			{Segment: "SEQ", Mandatory: true, MaxRepeat: 1},      // Sequence details
			{Segment: "TMP", Mandatory: false, MaxRepeat: 9},     // Temperature
			{Segment: "DTM", Mandatory: false, MaxRepeat: 9},     // Date/time/period
			{Segment: "----", Mandatory: false, MaxRepeat: 99},   // Segment group 154------------------
			{Segment: "NAD", Mandatory: true, MaxRepeat: 1},      // Name and address
			{Segment: "IFD", Mandatory: false, MaxRepeat: 1},     // Information detail
			{Segment: "IDE", Mandatory: false, MaxRepeat: 9},     // Identity
			{Segment: "----", Mandatory: false, MaxRepeat: 99},   // Segment group 155------------------
			{Segment: "CTA", Mandatory: true, MaxRepeat: 1},      // Contact information
			{Segment: "COM", Mandatory: false, MaxRepeat: 99},    // Communication contact
			{Segment: "----", Mandatory: false, MaxRepeat: 9},    // Segment group 156------------------
			{Segment: "ATT", Mandatory: true, MaxRepeat: 1},      // Attribute
			{Segment: "DGS", Mandatory: true, MaxRepeat: 1},      // Dangerous goods
			{Segment: "FTX", Mandatory: false, MaxRepeat: 9},     // Free text
			{Segment: "----", Mandatory: false, MaxRepeat: 9},    // Segment group 157------------------
			{Segment: "NAD", Mandatory: true, MaxRepeat: 1},      // Name and address
			{Segment: "IFD", Mandatory: false, MaxRepeat: 9},     // Information detail
			{Segment: "IDE", Mandatory: false, MaxRepeat: 9},     // Identity
			{Segment: "----", Mandatory: false, MaxRepeat: 9},    // Segment group 158------------------
			{Segment: "CTA", Mandatory: true, MaxRepeat: 1},      // Contact information
			{Segment: "COM", Mandatory: false, MaxRepeat: 9},     // Communication contact
			{Segment: "----", Mandatory: false, MaxRepeat: 99},   // Segment group 159------------------
			{Segment: "PYT", Mandatory: true, MaxRepeat: 1},      // Payment terms
			{Segment: "GEI", Mandatory: false, MaxRepeat: 9},     // Processing information
			{Segment: "MOA", Mandatory: false, MaxRepeat: 99},    // Monetary amount
			{Segment: "CUX", Mandatory: false, MaxRepeat: 1},     // Currencies
			{Segment: "----", Mandatory: false, MaxRepeat: 1},    // Segment group 160------------------
			{Segment: "LOC", Mandatory: true, MaxRepeat: 1},      // Place/location identification
			{Segment: "DTM", Mandatory: false, MaxRepeat: 99},    // Date/time/period
		},
		Mandatory: false, MaxRepeat: 9999,
	},
	{Segment: "HYN", Mandatory: true, MaxRepeat: 1},      // Hierarchy information
	{Segment: "----", Mandatory: false, MaxRepeat: 9999}, // Segment group 161------------------
	{Segment: "CNI", Mandatory: true, MaxRepeat: 1},      // Consignment information
	{Segment: "RFF", Mandatory: false, MaxRepeat: 99},    // Reference
	{Segment: "GEI", Mandatory: false, MaxRepeat: 99},    // Processing information
	{Segment: "STS", Mandatory: false, MaxRepeat: 99},    // Status
	{Segment: "MEA", Mandatory: false, MaxRepeat: 1},     // Measurements
	{Segment: "QTY", Mandatory: false, MaxRepeat: 99},    // Quantity
	{Segment: "MOA", Mandatory: false, MaxRepeat: 99},    // Monetary amount
	{Segment: "IFD", Mandatory: false, MaxRepeat: 99},    // Information detail
	{Segment: "HAN", Mandatory: false, MaxRepeat: 1},     // Handling instructions
	{Segment: "CNT", Mandatory: false, MaxRepeat: 9},     // Control total
	{Segment: "DTM", Mandatory: false, MaxRepeat: 9},     // Date/time/period
	{Segment: "----", Mandatory: false, MaxRepeat: 99},   // Segment group 162------------------
	{Segment: "NAD", Mandatory: true, MaxRepeat: 1},      // Name and address
	{Segment: "IFD", Mandatory: false, MaxRepeat: 99},    // Information detail
	{Segment: "DTM", Mandatory: false, MaxRepeat: 99},    // Date/time/period
	{Segment: "DMS", Mandatory: false, MaxRepeat: 99},    // Document/message summary
	{Segment: "IDE", Mandatory: false, MaxRepeat: 9},     // Identity
	{Segment: "----", Mandatory: false, MaxRepeat: 99},   // Segment group 163------------------
	{Segment: "CTA", Mandatory: true, MaxRepeat: 1},      // Contact information
	{Segment: "COM", Mandatory: false, MaxRepeat: 99},    // Communication contact
	{Segment: "----", Mandatory: false, MaxRepeat: 99},   // Segment group 164------------------
	{Segment: "LOC", Mandatory: true, MaxRepeat: 1},      // Place/location identification
	{Segment: "SEQ", Mandatory: true, MaxRepeat: 1},      // Sequence details
	{Segment: "DTM", Mandatory: false, MaxRepeat: 99},    // Date/time/period
	{Segment: "ADR", Mandatory: false, MaxRepeat: 9},     // Address
	{Segment: "TPL", Mandatory: false, MaxRepeat: 9},     // Transport placement
	{Segment: "GEI", Mandatory: false, MaxRepeat: 9},     // Processing information
	{Segment: "----", Mandatory: false, MaxRepeat: 99},   // Segment group 165------------------
	{Segment: "DOC", Mandatory: true, MaxRepeat: 1},      // Document/message details
	{Segment: "RFF", Mandatory: false, MaxRepeat: 99},    // Reference
	{Segment: "DTM", Mandatory: false, MaxRepeat: 99},    // Date/time/period
	{Segment: "MOA", Mandatory: false, MaxRepeat: 99},    // Monetary amount
	{Segment: "ALI", Mandatory: false, MaxRepeat: 1},     // Additional information
	{Segment: "QTY", Mandatory: false, MaxRepeat: 99},    // Quantity
	{Segment: "TSR", Mandatory: false, MaxRepeat: 99},    // Transport service requirements
	{Segment: "----", Mandatory: false, MaxRepeat: 99},   // Segment group 166------------------
	{Segment: "NAD", Mandatory: true, MaxRepeat: 1},      // Name and address
	{Segment: "IFD", Mandatory: false, MaxRepeat: 1},     // Information detail
	{Segment: "IDE", Mandatory: false, MaxRepeat: 9},     // Identity
	{Segment: "----", Mandatory: false, MaxRepeat: 99},   // Segment group 167------------------
	{Segment: "CTA", Mandatory: true, MaxRepeat: 1},      // Contact information
	{Segment: "COM", Mandatory: false, MaxRepeat: 99},    // Communication contact
	{Segment: "----", Mandatory: false, MaxRepeat: 99},   // Segment group 168------------------
	{Segment: "LOC", Mandatory: true, MaxRepeat: 1},      // Place/location identification
	{Segment: "DTM", Mandatory: false, MaxRepeat: 99},    // Date/time/period
	{Segment: "----", Mandatory: false, MaxRepeat: 99},   // Segment group 169------------------
	{Segment: "RCS", Mandatory: true, MaxRepeat: 1},      // Requirements and conditions
	{Segment: "FTX", Mandatory: false, MaxRepeat: 9},     // Free text
	{Segment: "DTM", Mandatory: false, MaxRepeat: 99},    // Date/time/period
	{Segment: "GEI", Mandatory: false, MaxRepeat: 99},    // Processing information
	{Segment: "ALI", Mandatory: false, MaxRepeat: 99},    // Additional information
	{Segment: "----", Mandatory: false, MaxRepeat: 99},   // Segment group 170------------------
	{Segment: "ERP", Mandatory: true, MaxRepeat: 1},      // Error point details
	{Segment: "ELU", Mandatory: false, MaxRepeat: 99},    // Data element usage details
	{Segment: "ARR", Mandatory: false, MaxRepeat: 99},    // Array information
	{Segment: "----", Mandatory: false, MaxRepeat: 1},    // Segment group 171------------------
	{Segment: "INP", Mandatory: true, MaxRepeat: 1},      // Parties and instruction
	{Segment: "DTM", Mandatory: false, MaxRepeat: 9},     // Date/time/period
	{Segment: "LOC", Mandatory: false, MaxRepeat: 99},    // Place/location identification
	{Segment: "CNT", Mandatory: false, MaxRepeat: 99},    // Control total
	{Segment: "FTX", Mandatory: false, MaxRepeat: 9},     // Free text
	{Segment: "----", Mandatory: false, MaxRepeat: 1},    // Segment group 172------------------
	{Segment: "NAD", Mandatory: true, MaxRepeat: 1},      // Name and address
	{Segment: "IFD", Mandatory: false, MaxRepeat: 1},     // Information detail
	{Segment: "DTM", Mandatory: false, MaxRepeat: 99},    // Date/time/period
	{Segment: "IDE", Mandatory: false, MaxRepeat: 9},     // Identity
	{Segment: "----", Mandatory: false, MaxRepeat: 99},   // Segment group 173------------------
	{Segment: "CTA", Mandatory: true, MaxRepeat: 1},      // Contact information
	{Segment: "COM", Mandatory: false, MaxRepeat: 99},    // Communication contact
	{Segment: "----", Mandatory: false, MaxRepeat: 99},   // Segment group 174------------------
	{Segment: "TOD", Mandatory: true, MaxRepeat: 1},      // Terms of delivery or transport
	{Segment: "RTE", Mandatory: false, MaxRepeat: 1},     // Rate details
	{Segment: "MOA", Mandatory: false, MaxRepeat: 1},     // Monetary amount
	{Segment: "----", Mandatory: false, MaxRepeat: 9},    // Segment group 175------------------
	{Segment: "PAC", Mandatory: true, MaxRepeat: 1},      // Package
	{Segment: "PCI", Mandatory: false, MaxRepeat: 99},    // Package identification
	{Segment: "FTX", Mandatory: false, MaxRepeat: 99},    // Free text
	{Segment: "GEI", Mandatory: false, MaxRepeat: 99},    // Processing information
	{Segment: "DIM", Mandatory: false, MaxRepeat: 99},    // Dimensions
	{Segment: "MEA", Mandatory: false, MaxRepeat: 99},    // Measurements
	{Segment: "----", Mandatory: false, MaxRepeat: 99},   // Segment group 176------------------
	{Segment: "TMP", Mandatory: true, MaxRepeat: 1},      // Temperature
	{Segment: "MEA", Mandatory: false, MaxRepeat: 99},    // Measurements
	{Segment: "DGS", Mandatory: false, MaxRepeat: 99},    // Dangerous goods
	{Segment: "----", Mandatory: false, MaxRepeat: 99},   // Segment group 177------------------
	{Segment: "RSL", Mandatory: true, MaxRepeat: 1},      // Result
	{Segment: "QTY", Mandatory: false, MaxRepeat: 99},    // Quantity
	{Segment: "DTM", Mandatory: false, MaxRepeat: 99},    // Date/time/period
	{Segment: "----", Mandatory: false, MaxRepeat: 1},    // Segment group 178------------------
	{Segment: "GOR", Mandatory: true, MaxRepeat: 1},      // Governmental requirements
	{Segment: "LOC", Mandatory: false, MaxRepeat: 1},     // Place/location identification
	{Segment: "----", Mandatory: false, MaxRepeat: 9999}, // Segment group 179------------------
	{Segment: "EQD", Mandatory: true, MaxRepeat: 1},      // Equipment details
	{Segment: "SEQ", Mandatory: true, MaxRepeat: 1},      // Sequence details
	{Segment: "QTY", Mandatory: false, MaxRepeat: 1},     // Quantity
	{Segment: "MEA", Mandatory: false, MaxRepeat: 9},     // Measurements
	{Segment: "GEI", Mandatory: false, MaxRepeat: 99},    // Processing information
	{Segment: "TSR", Mandatory: false, MaxRepeat: 99},    // Transport service requirements
	{Segment: "----", Mandatory: false, MaxRepeat: 99},   // Segment group 180------------------
	{Segment: "SEL", Mandatory: true, MaxRepeat: 1},      // Seal number
	{Segment: "SEQ", Mandatory: true, MaxRepeat: 1},      // Sequence details
	{Segment: "----", Mandatory: false, MaxRepeat: 99},   // Segment group 181------------------
	{Segment: "NAD", Mandatory: true, MaxRepeat: 1},      // Name and address
	{Segment: "IFD", Mandatory: false, MaxRepeat: 1},     // Information detail
	{Segment: "DTM", Mandatory: false, MaxRepeat: 99},    // Date/time/period
	{Segment: "IDE", Mandatory: false, MaxRepeat: 9},     // Identity
	{Segment: "----", Mandatory: false, MaxRepeat: 99},   // Segment group 182------------------
	{Segment: "CTA", Mandatory: true, MaxRepeat: 1},      // Contact information
	{Segment: "COM", Mandatory: false, MaxRepeat: 99},    // Communication contact
	{Segment: "----", Mandatory: false, MaxRepeat: 9},    // Segment group 183------------------
	{Segment: "LOC", Mandatory: true, MaxRepeat: 1},      // Place/location identification
	{Segment: "DTM", Mandatory: false, MaxRepeat: 99},    // Date/time/period
	{Segment: "----", Mandatory: false, MaxRepeat: 99},   // Segment group 184------------------
	{Segment: "DOC", Mandatory: true, MaxRepeat: 1},      // Document/message details
	{Segment: "DTM", Mandatory: false, MaxRepeat: 99},    // Date/time/period
	{Segment: "MOA", Mandatory: false, MaxRepeat: 99},    // Monetary amount
	{Segment: "GEI", Mandatory: false, MaxRepeat: 9},     // Processing information
	{Segment: "----", Mandatory: false, MaxRepeat: 99},   // Segment group 185------------------
	{Segment: "NAD", Mandatory: true, MaxRepeat: 1},      // Name and address
	{Segment: "IFD", Mandatory: false, MaxRepeat: 1},     // Information detail
	{Segment: "IDE", Mandatory: false, MaxRepeat: 9},     // Identity
	{Segment: "----", Mandatory: false, MaxRepeat: 99},   // Segment group 186------------------
	{Segment: "CTA", Mandatory: true, MaxRepeat: 1},      // Contact information
	{Segment: "COM", Mandatory: false, MaxRepeat: 99},    // Communication contact
	{Segment: "----", Mandatory: false, MaxRepeat: 1},    // Segment group 187------------------
	{Segment: "LOC", Mandatory: true, MaxRepeat: 1},      // Place/location identification
	{Segment: "DTM", Mandatory: false, MaxRepeat: 99},    // Date/time/period
	{Segment: "----", Mandatory: false, MaxRepeat: 99},   // Segment group 188------------------
	{Segment: "TDT", Mandatory: true, MaxRepeat: 1},      // Transport information
	{Segment: "DTM", Mandatory: false, MaxRepeat: 99},    // Date/time/period
	{Segment: "RFF", Mandatory: false, MaxRepeat: 99},    // Reference
	{Segment: "QTY", Mandatory: false, MaxRepeat: 99},    // Quantity
	{Segment: "MEA", Mandatory: false, MaxRepeat: 99},    // Measurements
	{Segment: "MOA", Mandatory: false, MaxRepeat: 99},    // Monetary amount
	{Segment: "POC", Mandatory: false, MaxRepeat: 99},    // Purpose of conveyance call
	{Segment: "DIM", Mandatory: false, MaxRepeat: 99},    // Dimensions
	{Segment: "FTX", Mandatory: false, MaxRepeat: 99},    // Free text
	{Segment: "GDS", Mandatory: false, MaxRepeat: 99},    // Nature of cargo
	{Segment: "STS", Mandatory: false, MaxRepeat: 1},     // Status
	{Segment: "GPO", Mandatory: false, MaxRepeat: 1},     // Geographical position
	{Segment: "----", Mandatory: false, MaxRepeat: 99},   // Segment group 189------------------
	{Segment: "LOC", Mandatory: true, MaxRepeat: 1},      // Place/location identification
	{Segment: "SEQ", Mandatory: true, MaxRepeat: 1},      // Sequence details
	{Segment: "DTM", Mandatory: false, MaxRepeat: 99},    // Date/time/period
	{Segment: "----", Mandatory: false, MaxRepeat: 99},   // Segment group 190------------------
	{Segment: "NAD", Mandatory: true, MaxRepeat: 1},      // Name and address
	{Segment: "IFD", Mandatory: false, MaxRepeat: 1},     // Information detail
	{Segment: "NAT", Mandatory: false, MaxRepeat: 1},     // Nationality
	{Segment: "DTM", Mandatory: false, MaxRepeat: 99},    // Date/time/period
	{Segment: "LOC", Mandatory: false, MaxRepeat: 99},    // Place/location identification
	{Segment: "EMP", Mandatory: false, MaxRepeat: 99},    // Employment details
	{Segment: "IDE", Mandatory: false, MaxRepeat: 9},     // Identity
	{Segment: "----", Mandatory: false, MaxRepeat: 99},   // Segment group 191------------------
	{Segment: "CTA", Mandatory: true, MaxRepeat: 1},      // Contact information
	{Segment: "COM", Mandatory: false, MaxRepeat: 99},    // Communication contact
	{Segment: "----", Mandatory: false, MaxRepeat: 99},   // Segment group 192------------------
	{Segment: "ADR", Mandatory: true, MaxRepeat: 1},      // Address
	{Segment: "EMP", Mandatory: false, MaxRepeat: 1},     // Employment details
	{Segment: "----", Mandatory: false, MaxRepeat: 99},   // Segment group 193------------------
	{Segment: "DOC", Mandatory: true, MaxRepeat: 1},      // Document/message details
	{Segment: "RFF", Mandatory: false, MaxRepeat: 99},    // Reference
	{Segment: "DTM", Mandatory: false, MaxRepeat: 99},    // Date/time/period
	{Segment: "MOA", Mandatory: false, MaxRepeat: 99},    // Monetary amount
	{Segment: "ALI", Mandatory: false, MaxRepeat: 1},     // Additional information
	{Segment: "QTY", Mandatory: false, MaxRepeat: 99},    // Quantity
	{Segment: "TSR", Mandatory: false, MaxRepeat: 99},    // Transport service requirements
	{Segment: "GEI", Mandatory: false, MaxRepeat: 9},     // Processing information
	{Segment: "----", Mandatory: false, MaxRepeat: 99},   // Segment group 194------------------
	{Segment: "NAD", Mandatory: true, MaxRepeat: 1},      // Name and address
	{Segment: "IFD", Mandatory: false, MaxRepeat: 1},     // Information detail
	{Segment: "IDE", Mandatory: false, MaxRepeat: 9},     // Identity
	{Segment: "----", Mandatory: false, MaxRepeat: 99},   // Segment group 195------------------
	{Segment: "CTA", Mandatory: true, MaxRepeat: 1},      // Contact information
	{Segment: "COM", Mandatory: false, MaxRepeat: 99},    // Communication contact
	{Segment: "----", Mandatory: false, MaxRepeat: 99},   // Segment group 196------------------
	{Segment: "LOC", Mandatory: true, MaxRepeat: 1},      // Place/location identification
	{Segment: "DTM", Mandatory: false, MaxRepeat: 99},    // Date/time/period
	{Segment: "----", Mandatory: false, MaxRepeat: 99},   // Segment group 197------------------
	{Segment: "AJT", Mandatory: true, MaxRepeat: 1},      // Adjustment details
	{Segment: "----", Mandatory: false, MaxRepeat: 99},   // Segment group 198------------------
	{Segment: "ERP", Mandatory: true, MaxRepeat: 1},      // Error point details
	{Segment: "ELU", Mandatory: false, MaxRepeat: 99},    // Data element usage details
	{Segment: "----", Mandatory: false, MaxRepeat: 99},   // Segment group 199------------------
	{Segment: "EQD", Mandatory: true, MaxRepeat: 1},      // Equipment details
	{Segment: "SEQ", Mandatory: true, MaxRepeat: 1},      // Sequence details
	{Segment: "QTY", Mandatory: false, MaxRepeat: 1},     // Quantity
	{Segment: "GEI", Mandatory: false, MaxRepeat: 99},    // Processing information
	{Segment: "----", Mandatory: false, MaxRepeat: 99},   // Segment group 200------------------
	{Segment: "SEL", Mandatory: true, MaxRepeat: 1},      // Seal number
	{Segment: "SEQ", Mandatory: true, MaxRepeat: 1},      // Sequence details
	{Segment: "----", Mandatory: false, MaxRepeat: 99},   // Segment group 201------------------
	{Segment: "NAD", Mandatory: true, MaxRepeat: 1},      // Name and address
	{Segment: "IFD", Mandatory: false, MaxRepeat: 1},     // Information detail
	{Segment: "DTM", Mandatory: false, MaxRepeat: 99},    // Date/time/period
	{Segment: "IDE", Mandatory: false, MaxRepeat: 9},     // Identity
	{Segment: "----", Mandatory: false, MaxRepeat: 99},   // Segment group 202------------------
	{Segment: "CTA", Mandatory: true, MaxRepeat: 1},      // Contact information
	{Segment: "COM", Mandatory: false, MaxRepeat: 99},    // Communication contact
	{Segment: "----", Mandatory: false, MaxRepeat: 99},   // Segment group 203------------------
	{Segment: "LOC", Mandatory: true, MaxRepeat: 1},      // Place/location identification
	{Segment: "DTM", Mandatory: false, MaxRepeat: 99},    // Date/time/period
	{Segment: "----", Mandatory: false, MaxRepeat: 9999}, // Segment group 204------------------
	{Segment: "CCI", Mandatory: true, MaxRepeat: 1},      // Characteristic/class id
	{Segment: "CAV", Mandatory: false, MaxRepeat: 1},     // Characteristic value
	{Segment: "----", Mandatory: false, MaxRepeat: 9999}, // Segment group 205------------------
	{Segment: "SEQ", Mandatory: true, MaxRepeat: 1},      // Sequence details
	{Segment: "GEI", Mandatory: false, MaxRepeat: 99},    // Processing information
	{Segment: "STS", Mandatory: false, MaxRepeat: 99},    // Status
	{Segment: "RFF", Mandatory: false, MaxRepeat: 99},    // Reference
	{Segment: "IFD", Mandatory: false, MaxRepeat: 99},    // Information detail
	{Segment: "HAN", Mandatory: false, MaxRepeat: 99},    // Handling instructions
	{Segment: "TDT", Mandatory: false, MaxRepeat: 9},     // Transport information
	{Segment: "----", Mandatory: false, MaxRepeat: 99},   // Segment group 206------------------
	{Segment: "NAD", Mandatory: true, MaxRepeat: 1},      // Name and address
	{Segment: "IFD", Mandatory: false, MaxRepeat: 1},     // Information detail
	{Segment: "DTM", Mandatory: false, MaxRepeat: 99},    // Date/time/period
	{Segment: "IDE", Mandatory: false, MaxRepeat: 9},     // Identity
	{Segment: "----", Mandatory: false, MaxRepeat: 99},   // Segment group 207------------------
	{Segment: "CTA", Mandatory: true, MaxRepeat: 1},      // Contact information
	{Segment: "COM", Mandatory: false, MaxRepeat: 99},    // Communication contact
	{Segment: "----", Mandatory: false, MaxRepeat: 99},   // Segment group 208------------------
	{Segment: "LOC", Mandatory: true, MaxRepeat: 1},      // Place/location identification
	{Segment: "DTM", Mandatory: false, MaxRepeat: 99},    // Date/time/period
	{Segment: "ADR", Mandatory: false, MaxRepeat: 9},     // Address
	{Segment: "----", Mandatory: false, MaxRepeat: 99},   // Segment group 209------------------
	{Segment: "DOC", Mandatory: true, MaxRepeat: 1},      // Document/message details
	{Segment: "RFF", Mandatory: false, MaxRepeat: 99},    // Reference
	{Segment: "DTM", Mandatory: false, MaxRepeat: 99},    // Date/time/period
	{Segment: "MOA", Mandatory: false, MaxRepeat: 99},    // Monetary amount
	{Segment: "ALI", Mandatory: false, MaxRepeat: 1},     // Additional information
	{Segment: "QTY", Mandatory: false, MaxRepeat: 99},    // Quantity
	{Segment: "TSR", Mandatory: false, MaxRepeat: 99},    // Transport service requirements
	{Segment: "----", Mandatory: false, MaxRepeat: 99},   // Segment group 210------------------
	{Segment: "NAD", Mandatory: true, MaxRepeat: 1},      // Name and address
	{Segment: "IFD", Mandatory: false, MaxRepeat: 1},     // Information detail
	{Segment: "IDE", Mandatory: false, MaxRepeat: 9},     // Identity
	{Segment: "----", Mandatory: false, MaxRepeat: 99},   // Segment group 211------------------
	{Segment: "CTA", Mandatory: true, MaxRepeat: 1},      // Contact information
	{Segment: "COM", Mandatory: false, MaxRepeat: 99},    // Communication contact
	{Segment: "----", Mandatory: false, MaxRepeat: 99},   // Segment group 212------------------
	{Segment: "LOC", Mandatory: true, MaxRepeat: 1},      // Place/location identification
	{Segment: "DTM", Mandatory: false, MaxRepeat: 99},    // Date/time/period
	{Segment: "----", Mandatory: false, MaxRepeat: 99},   // Segment group 213------------------
	{Segment: "AJT", Mandatory: true, MaxRepeat: 1},      // Adjustment details
	{Segment: "----", Mandatory: false, MaxRepeat: 99},   // Segment group 214------------------
	{Segment: "ERP", Mandatory: true, MaxRepeat: 1},      // Error point details
	{Segment: "ELU", Mandatory: false, MaxRepeat: 99},    // Data element usage details
	{Segment: "----", Mandatory: false, MaxRepeat: 99},   // Segment group 215------------------
	{Segment: "RCS", Mandatory: true, MaxRepeat: 1},      // Requirements and conditions
	{Segment: "FTX", Mandatory: false, MaxRepeat: 9},     // Free text
	{Segment: "DTM", Mandatory: false, MaxRepeat: 99},    // Date/time/period
	{Segment: "GEI", Mandatory: false, MaxRepeat: 99},    // Processing information
	{Segment: "ALI", Mandatory: false, MaxRepeat: 99},    // Additional information
	{Segment: "----", Mandatory: false, MaxRepeat: 99},   // Segment group 216------------------
	{Segment: "ERP", Mandatory: true, MaxRepeat: 1},      // Error point details
	{Segment: "ELU", Mandatory: false, MaxRepeat: 99},    // Data element usage details
	{Segment: "ARR", Mandatory: false, MaxRepeat: 99},    // Array information
	{Segment: "----", Mandatory: false, MaxRepeat: 99},   // Segment group 217------------------
	{Segment: "INP", Mandatory: true, MaxRepeat: 1},      // Parties and instruction
	{Segment: "DTM", Mandatory: false, MaxRepeat: 9},     // Date/time/period
	{Segment: "LOC", Mandatory: false, MaxRepeat: 99},    // Place/location identification
	{Segment: "CNT", Mandatory: false, MaxRepeat: 99},    // Control total
	{Segment: "----", Mandatory: false, MaxRepeat: 1},    // Segment group 218------------------
	{Segment: "NAD", Mandatory: true, MaxRepeat: 1},      // Name and address
	{Segment: "IFD", Mandatory: false, MaxRepeat: 1},     // Information detail
	{Segment: "DTM", Mandatory: false, MaxRepeat: 99},    // Date/time/period
	{Segment: "IDE", Mandatory: false, MaxRepeat: 9},     // Identity
	{Segment: "----", Mandatory: false, MaxRepeat: 99},   // Segment group 219------------------
	{Segment: "CTA", Mandatory: true, MaxRepeat: 1},      // Contact information
	{Segment: "COM", Mandatory: false, MaxRepeat: 99},    // Communication contact
	{Segment: "----", Mandatory: false, MaxRepeat: 99},   // Segment group 220------------------
	{Segment: "MEA", Mandatory: true, MaxRepeat: 1},      // Measurements
	{Segment: "QTY", Mandatory: false, MaxRepeat: 99},    // Quantity
	{Segment: "PCD", Mandatory: false, MaxRepeat: 1},     // Percentage details
	{Segment: "----", Mandatory: false, MaxRepeat: 999},  // Segment group 221------------------
	{Segment: "PAC", Mandatory: true, MaxRepeat: 1},      // Package
	{Segment: "SEQ", Mandatory: true, MaxRepeat: 1},      // Sequence details
	{Segment: "PCI", Mandatory: false, MaxRepeat: 99},    // Package identification
	{Segment: "FTX", Mandatory: false, MaxRepeat: 99},    // Free text
	{Segment: "GEI", Mandatory: false, MaxRepeat: 99},    // Processing information
	{Segment: "DIM", Mandatory: false, MaxRepeat: 99},    // Dimensions
	{Segment: "MEA", Mandatory: false, MaxRepeat: 99},    // Measurements
	{Segment: "----", Mandatory: false, MaxRepeat: 99},   // Segment group 222------------------
	{Segment: "NAD", Mandatory: true, MaxRepeat: 1},      // Name and address
	{Segment: "IFD", Mandatory: false, MaxRepeat: 1},     // Information detail
	{Segment: "DTM", Mandatory: false, MaxRepeat: 99},    // Date/time/period
	{Segment: "IDE", Mandatory: false, MaxRepeat: 9},     // Identity
	{Segment: "----", Mandatory: false, MaxRepeat: 99},   // Segment group 223------------------
	{Segment: "CTA", Mandatory: true, MaxRepeat: 1},      // Contact information
	{Segment: "COM", Mandatory: false, MaxRepeat: 99},    // Communication contact
	{Segment: "----", Mandatory: false, MaxRepeat: 99},   // Segment group 224------------------
	{Segment: "TMP", Mandatory: true, MaxRepeat: 1},      // Temperature
	{Segment: "MEA", Mandatory: false, MaxRepeat: 99},    // Measurements
	{Segment: "DGS", Mandatory: false, MaxRepeat: 99},    // Dangerous goods
	{Segment: "----", Mandatory: false, MaxRepeat: 9999}, // Segment group 225------------------
	{Segment: "EQD", Mandatory: true, MaxRepeat: 1},      // Equipment details
	{Segment: "SEQ", Mandatory: true, MaxRepeat: 1},      // Sequence details
	{Segment: "QTY", Mandatory: false, MaxRepeat: 1},     // Quantity
	{Segment: "MEA", Mandatory: false, MaxRepeat: 9},     // Measurements
	{Segment: "GEI", Mandatory: false, MaxRepeat: 99},    // Processing information
	{Segment: "----", Mandatory: false, MaxRepeat: 99},   // Segment group 226------------------
	{Segment: "SEL", Mandatory: true, MaxRepeat: 1},      // Seal number
	{Segment: "SEQ", Mandatory: true, MaxRepeat: 1},      // Sequence details
	{Segment: "----", Mandatory: false, MaxRepeat: 99},   // Segment group 227------------------
	{Segment: "NAD", Mandatory: true, MaxRepeat: 1},      // Name and address
	{Segment: "IFD", Mandatory: false, MaxRepeat: 1},     // Information detail
	{Segment: "DTM", Mandatory: false, MaxRepeat: 99},    // Date/time/period
	{Segment: "IDE", Mandatory: false, MaxRepeat: 9},     // Identity
	{Segment: "----", Mandatory: false, MaxRepeat: 99},   // Segment group 228------------------
	{Segment: "CTA", Mandatory: true, MaxRepeat: 1},      // Contact information
	{Segment: "COM", Mandatory: false, MaxRepeat: 99},    // Communication contact
	{Segment: "----", Mandatory: false, MaxRepeat: 99},   // Segment group 229------------------
	{Segment: "LOC", Mandatory: true, MaxRepeat: 1},      // Place/location identification
	{Segment: "DTM", Mandatory: false, MaxRepeat: 99},    // Date/time/period
	{Segment: "----", Mandatory: false, MaxRepeat: 999},  // Segment group 230------------------
	{Segment: "PAC", Mandatory: true, MaxRepeat: 1},      // Package
	{Segment: "SEQ", Mandatory: true, MaxRepeat: 1},      // Sequence details
	{Segment: "PCI", Mandatory: false, MaxRepeat: 99},    // Package identification
	{Segment: "FTX", Mandatory: false, MaxRepeat: 99},    // Free text
	{Segment: "GEI", Mandatory: false, MaxRepeat: 99},    // Processing information
	{Segment: "DIM", Mandatory: false, MaxRepeat: 99},    // Dimensions
	{Segment: "MEA", Mandatory: false, MaxRepeat: 99},    // Measurements
	{Segment: "----", Mandatory: false, MaxRepeat: 99},   // Segment group 231------------------
	{Segment: "NAD", Mandatory: true, MaxRepeat: 1},      // Name and address
	{Segment: "IFD", Mandatory: false, MaxRepeat: 1},     // Information detail
	{Segment: "DTM", Mandatory: false, MaxRepeat: 99},    // Date/time/period
	{Segment: "IDE", Mandatory: false, MaxRepeat: 9},     // Identity
	{Segment: "----", Mandatory: false, MaxRepeat: 99},   // Segment group 232------------------
	{Segment: "CTA", Mandatory: true, MaxRepeat: 1},      // Contact information
	{Segment: "COM", Mandatory: false, MaxRepeat: 99},    // Communication contact
	{Segment: "----", Mandatory: false, MaxRepeat: 99},   // Segment group 233------------------
	{Segment: "ARR", Mandatory: true, MaxRepeat: 1},      // Array information
	{Segment: "FTX", Mandatory: false, MaxRepeat: 99},    // Free text
	{Segment: "TCC", Mandatory: false, MaxRepeat: 99},    // Charge/rate calculations
	{Segment: "----", Mandatory: false, MaxRepeat: 9},    // Segment group 234------------------
	{Segment: "ATT", Mandatory: true, MaxRepeat: 1},      // Attribute
	{Segment: "DGS", Mandatory: true, MaxRepeat: 1},      // Dangerous goods
	{Segment: "FTX", Mandatory: false, MaxRepeat: 9},     // Free text
	{Segment: "----", Mandatory: false, MaxRepeat: 9},    // Segment group 235------------------
	{Segment: "NAD", Mandatory: true, MaxRepeat: 1},      // Name and address
	{Segment: "IDE", Mandatory: false, MaxRepeat: 9},     // Identity
	{Segment: "----", Mandatory: false, MaxRepeat: 9},    // Segment group 236------------------
	{Segment: "CTA", Mandatory: true, MaxRepeat: 1},      // Contact information
	{Segment: "COM", Mandatory: false, MaxRepeat: 9},     // Communication contact
	{Segment: "----", Mandatory: false, MaxRepeat: 99},   // Segment group 237------------------
	{Segment: "GID", Mandatory: true, MaxRepeat: 1},      // Goods item details
	{Segment: "IMD", Mandatory: false, MaxRepeat: 99},    // Item description
	{Segment: "APP", Mandatory: false, MaxRepeat: 99},    // Applicability
	{Segment: "FTX", Mandatory: false, MaxRepeat: 99},    // Free text
	{Segment: "RFF", Mandatory: false, MaxRepeat: 99},    // Reference
	{Segment: "DTM", Mandatory: false, MaxRepeat: 99},    // Date/time/period
	{Segment: "GIR", Mandatory: false, MaxRepeat: 99},    // Related identification numbers
	{Segment: "GIN", Mandatory: false, MaxRepeat: 99},    // Goods identity number
	{Segment: "GEI", Mandatory: false, MaxRepeat: 99},    // Processing information
	{Segment: "MEA", Mandatory: false, MaxRepeat: 99},    // Measurements
	{Segment: "MOA", Mandatory: false, MaxRepeat: 99},    // Monetary amount
	{Segment: "CDI", Mandatory: false, MaxRepeat: 99},    // Physical or logical state
	{Segment: "PGI", Mandatory: false, MaxRepeat: 99},    // Product group information
	{Segment: "TCC", Mandatory: false, MaxRepeat: 99},    // Charge/rate calculations
	{Segment: "CNT", Mandatory: false, MaxRepeat: 99},    // Control total
	{Segment: "IFD", Mandatory: false, MaxRepeat: 99},    // Information detail
	{Segment: "TDT", Mandatory: false, MaxRepeat: 9},     // Transport information
	{Segment: "----", Mandatory: false, MaxRepeat: 99},   // Segment group 238------------------
	{Segment: "LOC", Mandatory: true, MaxRepeat: 1},      // Place/location identification
	{Segment: "SEQ", Mandatory: true, MaxRepeat: 1},      // Sequence details
	{Segment: "DTM", Mandatory: false, MaxRepeat: 99},    // Date/time/period
	{Segment: "----", Mandatory: false, MaxRepeat: 99},   // Segment group 239------------------
	{Segment: "NAD", Mandatory: true, MaxRepeat: 1},      // Name and address
	{Segment: "IFD", Mandatory: false, MaxRepeat: 9},     // Information detail
	{Segment: "DTM", Mandatory: false, MaxRepeat: 99},    // Date/time/period
	{Segment: "IDE", Mandatory: false, MaxRepeat: 9},     // Identity
	{Segment: "----", Mandatory: false, MaxRepeat: 99},   // Segment group 240------------------
	{Segment: "CTA", Mandatory: true, MaxRepeat: 1},      // Contact information
	{Segment: "COM", Mandatory: false, MaxRepeat: 99},    // Communication contact
	{Segment: "----", Mandatory: false, MaxRepeat: 99},   // Segment group 241------------------
	{Segment: "DOC", Mandatory: true, MaxRepeat: 1},      // Document/message details
	{Segment: "RFF", Mandatory: false, MaxRepeat: 99},    // Reference
	{Segment: "DTM", Mandatory: false, MaxRepeat: 99},    // Date/time/period
	{Segment: "MOA", Mandatory: false, MaxRepeat: 99},    // Monetary amount
	{Segment: "ALI", Mandatory: false, MaxRepeat: 1},     // Additional information
	{Segment: "QTY", Mandatory: false, MaxRepeat: 99},    // Quantity
	{Segment: "TSR", Mandatory: false, MaxRepeat: 99},    // Transport service requirements
	{Segment: "GEI", Mandatory: false, MaxRepeat: 9},     // Processing information
	{Segment: "GIN", Mandatory: false, MaxRepeat: 99},    // Goods identity number
	{Segment: "GIR", Mandatory: false, MaxRepeat: 99},    // Related identification numbers
	{Segment: "IMD", Mandatory: false, MaxRepeat: 99},    // Item description
	{Segment: "----", Mandatory: false, MaxRepeat: 99},   // Segment group 242------------------
	{Segment: "NAD", Mandatory: true, MaxRepeat: 1},      // Name and address
	{Segment: "IFD", Mandatory: false, MaxRepeat: 1},     // Information detail
	{Segment: "IDE", Mandatory: false, MaxRepeat: 9},     // Identity
	{Segment: "----", Mandatory: false, MaxRepeat: 99},   // Segment group 243------------------
	{Segment: "CTA", Mandatory: true, MaxRepeat: 1},      // Contact information
	{Segment: "COM", Mandatory: false, MaxRepeat: 99},    // Communication contact
	{Segment: "----", Mandatory: false, MaxRepeat: 99},   // Segment group 244------------------
	{Segment: "LOC", Mandatory: true, MaxRepeat: 1},      // Place/location identification
	{Segment: "DTM", Mandatory: false, MaxRepeat: 99},    // Date/time/period
	{Segment: "----", Mandatory: false, MaxRepeat: 99},   // Segment group 245------------------
	{Segment: "RCS", Mandatory: true, MaxRepeat: 1},      // Requirements and conditions
	{Segment: "FTX", Mandatory: false, MaxRepeat: 9},     // Free text
	{Segment: "DTM", Mandatory: false, MaxRepeat: 99},    // Date/time/period
	{Segment: "GEI", Mandatory: false, MaxRepeat: 99},    // Processing information
	{Segment: "ALI", Mandatory: false, MaxRepeat: 99},    // Additional information
	{Segment: "----", Mandatory: false, MaxRepeat: 99},   // Segment group 246------------------
	{Segment: "ERP", Mandatory: true, MaxRepeat: 1},      // Error point details
	{Segment: "ELU", Mandatory: false, MaxRepeat: 99},    // Data element usage details
	{Segment: "----", Mandatory: false, MaxRepeat: 999},  // Segment group 247------------------
	{Segment: "PAC", Mandatory: true, MaxRepeat: 1},      // Package
	{Segment: "PCI", Mandatory: false, MaxRepeat: 99},    // Package identification
	{Segment: "FTX", Mandatory: false, MaxRepeat: 99},    // Free text
	{Segment: "GEI", Mandatory: false, MaxRepeat: 99},    // Processing information
	{Segment: "DIM", Mandatory: false, MaxRepeat: 99},    // Dimensions
	{Segment: "MEA", Mandatory: false, MaxRepeat: 99},    // Measurements
	{Segment: "----", Mandatory: false, MaxRepeat: 99},   // Segment group 248------------------
	{Segment: "COD", Mandatory: true, MaxRepeat: 1},      // Component details
	{Segment: "QTY", Mandatory: false, MaxRepeat: 99},    // Quantity
	{Segment: "PCD", Mandatory: false, MaxRepeat: 99},    // Percentage details
	{Segment: "GEI", Mandatory: false, MaxRepeat: 99},    // Processing information
	{Segment: "LOC", Mandatory: false, MaxRepeat: 9},     // Place/location identification
	{Segment: "----", Mandatory: false, MaxRepeat: 1},    // Segment group 249------------------
	{Segment: "INP", Mandatory: true, MaxRepeat: 1},      // Parties and instruction
	{Segment: "DTM", Mandatory: false, MaxRepeat: 9},     // Date/time/period
	{Segment: "LOC", Mandatory: false, MaxRepeat: 99},    // Place/location identification
	{Segment: "----", Mandatory: false, MaxRepeat: 1},    // Segment group 250------------------
	{Segment: "NAD", Mandatory: true, MaxRepeat: 1},      // Name and address
	{Segment: "IFD", Mandatory: false, MaxRepeat: 1},     // Information detail
	{Segment: "DTM", Mandatory: false, MaxRepeat: 99},    // Date/time/period
	{Segment: "IDE", Mandatory: false, MaxRepeat: 9},     // Identity
	{Segment: "----", Mandatory: false, MaxRepeat: 99},   // Segment group 251------------------
	{Segment: "CTA", Mandatory: true, MaxRepeat: 1},      // Contact information
	{Segment: "COM", Mandatory: false, MaxRepeat: 99},    // Communication contact
	{Segment: "----", Mandatory: false, MaxRepeat: 99},   // Segment group 252------------------
	{Segment: "TAX", Mandatory: true, MaxRepeat: 1},      // Duty/tax/fee details
	{Segment: "DTM", Mandatory: false, MaxRepeat: 99},    // Date/time/period
	{Segment: "ALI", Mandatory: false, MaxRepeat: 99},    // Additional information
	{Segment: "MOA", Mandatory: false, MaxRepeat: 99},    // Monetary amount
	{Segment: "QTY", Mandatory: false, MaxRepeat: 99},    // Quantity
	{Segment: "----", Mandatory: false, MaxRepeat: 99},   // Segment group 253------------------
	{Segment: "PAI", Mandatory: true, MaxRepeat: 1},      // Payment instructions
	{Segment: "RFF", Mandatory: false, MaxRepeat: 99},    // Reference
	{Segment: "MOA", Mandatory: false, MaxRepeat: 99},    // Monetary amount
	{Segment: "DTM", Mandatory: false, MaxRepeat: 99},    // Date/time/period
	{Segment: "----", Mandatory: false, MaxRepeat: 99},   // Segment group 254------------------
	{Segment: "GIR", Mandatory: true, MaxRepeat: 1},      // Related identification numbers
	{Segment: "LOC", Mandatory: false, MaxRepeat: 99},    // Place/location identification
	{Segment: "VLI", Mandatory: false, MaxRepeat: 1},     // Value list identification
	{Segment: "DTM", Mandatory: false, MaxRepeat: 99},    // Date/time/period
	{Segment: "----", Mandatory: false, MaxRepeat: 99},   // Segment group 255------------------
	{Segment: "NAD", Mandatory: true, MaxRepeat: 1},      // Name and address
	{Segment: "IFD", Mandatory: false, MaxRepeat: 1},     // Information detail
	{Segment: "DTM", Mandatory: false, MaxRepeat: 99},    // Date/time/period
	{Segment: "IDE", Mandatory: false, MaxRepeat: 9},     // Identity
	{Segment: "----", Mandatory: false, MaxRepeat: 99},   // Segment group 256------------------
	{Segment: "CTA", Mandatory: true, MaxRepeat: 1},      // Contact information
	{Segment: "COM", Mandatory: false, MaxRepeat: 99},    // Communication contact
	{Segment: "----", Mandatory: false, MaxRepeat: 99},   // Segment group 257------------------
	{Segment: "DLI", Mandatory: true, MaxRepeat: 1},      // Document line identification
	{Segment: "MOA", Mandatory: false, MaxRepeat: 99},    // Monetary amount
	{Segment: "----", Mandatory: false, MaxRepeat: 99},   // Segment group 258------------------
	{Segment: "STS", Mandatory: true, MaxRepeat: 1},      // Status
	{Segment: "DTM", Mandatory: false, MaxRepeat: 99},    // Date/time/period
	{Segment: "----", Mandatory: false, MaxRepeat: 99},   // Segment group 259------------------
	{Segment: "ERP", Mandatory: true, MaxRepeat: 1},      // Error point details
	{Segment: "ELU", Mandatory: false, MaxRepeat: 99},    // Data element usage details
	{Segment: "----", Mandatory: false, MaxRepeat: 99},   // Segment group 260------------------
	{Segment: "TMP", Mandatory: true, MaxRepeat: 1},      // Temperature
	{Segment: "MEA", Mandatory: false, MaxRepeat: 99},    // Measurements
	{Segment: "DGS", Mandatory: false, MaxRepeat: 99},    // Dangerous goods
	{Segment: "----", Mandatory: false, MaxRepeat: 99},   // Segment group 261------------------
	{Segment: "EQD", Mandatory: true, MaxRepeat: 1},      // Equipment details
	{Segment: "SEQ", Mandatory: true, MaxRepeat: 1},      // Sequence details
	{Segment: "----", Mandatory: false, MaxRepeat: 99},   // Segment group 262------------------
	{Segment: "PRC", Mandatory: true, MaxRepeat: 1},      // Process identification
	{Segment: "SEQ", Mandatory: true, MaxRepeat: 1},      // Sequence details
	{Segment: "TMP", Mandatory: false, MaxRepeat: 9},     // Temperature
	{Segment: "DTM", Mandatory: false, MaxRepeat: 9},     // Date/time/period
	{Segment: "----", Mandatory: false, MaxRepeat: 99},   // Segment group 263------------------
	{Segment: "NAD", Mandatory: true, MaxRepeat: 1},      // Name and address
	{Segment: "IFD", Mandatory: false, MaxRepeat: 1},     // Information detail
	{Segment: "IDE", Mandatory: false, MaxRepeat: 9},     // Identity
	{Segment: "----", Mandatory: false, MaxRepeat: 99},   // Segment group 264------------------
	{Segment: "CTA", Mandatory: true, MaxRepeat: 1},      // Contact information
	{Segment: "COM", Mandatory: false, MaxRepeat: 99},    // Communication contact
	{Segment: "----", Mandatory: false, MaxRepeat: 9},    // Segment group 265------------------
	{Segment: "ATT", Mandatory: true, MaxRepeat: 1},      // Attribute
	{Segment: "DGS", Mandatory: true, MaxRepeat: 1},      // Dangerous goods
	{Segment: "FTX", Mandatory: false, MaxRepeat: 9},     // Free text
	{Segment: "PAC", Mandatory: false, MaxRepeat: 9},     // Package
	{Segment: "MEA", Mandatory: false, MaxRepeat: 9},     // Measurements
	{Segment: "----", Mandatory: false, MaxRepeat: 9},    // Segment group 266------------------
	{Segment: "NAD", Mandatory: true, MaxRepeat: 1},      // Name and address
	{Segment: "IDE", Mandatory: false, MaxRepeat: 9},     // Identity
	{Segment: "----", Mandatory: false, MaxRepeat: 9},    // Segment group 267------------------
	{Segment: "CTA", Mandatory: true, MaxRepeat: 1},      // Contact information
	{Segment: "COM", Mandatory: false, MaxRepeat: 9},     // Communication contact
	{Segment: "----", Mandatory: false, MaxRepeat: 99},   // Segment group 268------------------
	{Segment: "TOD", Mandatory: true, MaxRepeat: 1},      // Terms of delivery or transport
	{Segment: "RTE", Mandatory: false, MaxRepeat: 1},     // Rate details
	{Segment: "MOA", Mandatory: false, MaxRepeat: 1},     // Monetary amount
	{Segment: "----", Mandatory: false, MaxRepeat: 9999}, // Segment group 269------------------
	{Segment: "LIN", Mandatory: true, MaxRepeat: 1},      // Line item
	{Segment: "MOA", Mandatory: false, MaxRepeat: 99},    // Monetary amount
	{Segment: "IFD", Mandatory: false, MaxRepeat: 99},    // Information detail
	{Segment: "STS", Mandatory: false, MaxRepeat: 99},    // Status
	{Segment: "----", Mandatory: false, MaxRepeat: 99},   // Segment group 270------------------
	{Segment: "NAD", Mandatory: true, MaxRepeat: 1},      // Name and address
	{Segment: "IFD", Mandatory: false, MaxRepeat: 1},     // Information detail
	{Segment: "DTM", Mandatory: false, MaxRepeat: 99},    // Date/time/period
	{Segment: "IDE", Mandatory: false, MaxRepeat: 9},     // Identity
	{Segment: "----", Mandatory: false, MaxRepeat: 99},   // Segment group 271------------------
	{Segment: "CTA", Mandatory: true, MaxRepeat: 1},      // Contact information
	{Segment: "COM", Mandatory: false, MaxRepeat: 99},    // Communication contact
	{Segment: "----", Mandatory: false, MaxRepeat: 99},   // Segment group 272------------------
	{Segment: "LOC", Mandatory: true, MaxRepeat: 1},      // Place/location identification
	{Segment: "DTM", Mandatory: false, MaxRepeat: 99},    // Date/time/period
	{Segment: "----", Mandatory: false, MaxRepeat: 99},   // Segment group 273------------------
	{Segment: "DOC", Mandatory: true, MaxRepeat: 1},      // Document/message details
	{Segment: "RFF", Mandatory: false, MaxRepeat: 99},    // Reference
	{Segment: "DTM", Mandatory: false, MaxRepeat: 99},    // Date/time/period
	{Segment: "MOA", Mandatory: false, MaxRepeat: 99},    // Monetary amount
	{Segment: "ALI", Mandatory: false, MaxRepeat: 1},     // Additional information
	{Segment: "QTY", Mandatory: false, MaxRepeat: 99},    // Quantity
	{Segment: "TSR", Mandatory: false, MaxRepeat: 99},    // Transport service requirements
	{Segment: "----", Mandatory: false, MaxRepeat: 99},   // Segment group 274------------------
	{Segment: "NAD", Mandatory: true, MaxRepeat: 1},      // Name and address
	{Segment: "IFD", Mandatory: false, MaxRepeat: 1},     // Information detail
	{Segment: "IDE", Mandatory: false, MaxRepeat: 9},     // Identity
	{Segment: "----", Mandatory: false, MaxRepeat: 99},   // Segment group 275------------------
	{Segment: "CTA", Mandatory: true, MaxRepeat: 1},      // Contact information
	{Segment: "COM", Mandatory: false, MaxRepeat: 99},    // Communication contact
	{Segment: "----", Mandatory: false, MaxRepeat: 99},   // Segment group 276------------------
	{Segment: "LOC", Mandatory: true, MaxRepeat: 1},      // Place/location identification
	{Segment: "DTM", Mandatory: false, MaxRepeat: 99},    // Date/time/period
	{Segment: "----", Mandatory: false, MaxRepeat: 99},   // Segment group 277------------------
	{Segment: "RCS", Mandatory: true, MaxRepeat: 1},      // Requirements and conditions
	{Segment: "FTX", Mandatory: false, MaxRepeat: 9},     // Free text
	{Segment: "DTM", Mandatory: false, MaxRepeat: 99},    // Date/time/period
	{Segment: "GEI", Mandatory: false, MaxRepeat: 99},    // Processing information
	{Segment: "ALI", Mandatory: false, MaxRepeat: 99},    // Additional information
	{Segment: "----", Mandatory: false, MaxRepeat: 99},   // Segment group 278------------------
	{Segment: "ERP", Mandatory: true, MaxRepeat: 1},      // Error point details
	{Segment: "ELU", Mandatory: false, MaxRepeat: 99},    // Data element usage details
	{Segment: "ARR", Mandatory: false, MaxRepeat: 99},    // Array information
	{Segment: "----", Mandatory: false, MaxRepeat: 1},    // Segment group 279------------------
	{Segment: "INP", Mandatory: true, MaxRepeat: 1},      // Parties and instruction
	{Segment: "DTM", Mandatory: false, MaxRepeat: 9},     // Date/time/period
	{Segment: "LOC", Mandatory: false, MaxRepeat: 99},    // Place/location identification
	{Segment: "----", Mandatory: false, MaxRepeat: 1},    // Segment group 280------------------
	{Segment: "NAD", Mandatory: true, MaxRepeat: 1},      // Name and address
	{Segment: "IFD", Mandatory: false, MaxRepeat: 1},     // Information detail
	{Segment: "DTM", Mandatory: false, MaxRepeat: 99},    // Date/time/period
	{Segment: "IDE", Mandatory: false, MaxRepeat: 9},     // Identity
	{Segment: "----", Mandatory: false, MaxRepeat: 99},   // Segment group 281------------------
	{Segment: "CTA", Mandatory: true, MaxRepeat: 1},      // Contact information
	{Segment: "COM", Mandatory: false, MaxRepeat: 99},    // Communication contact
	{Segment: "----", Mandatory: false, MaxRepeat: 99},   // Segment group 282------------------
	{Segment: "MEA", Mandatory: true, MaxRepeat: 1},      // Measurements
	{Segment: "QTY", Mandatory: false, MaxRepeat: 99},    // Quantity
	{Segment: "----", Mandatory: false, MaxRepeat: 99},   // Segment group 283------------------
	{Segment: "PAC", Mandatory: true, MaxRepeat: 1},      // Package
	{Segment: "SEQ", Mandatory: true, MaxRepeat: 1},      // Sequence details
	{Segment: "PCI", Mandatory: false, MaxRepeat: 99},    // Package identification
	{Segment: "FTX", Mandatory: false, MaxRepeat: 99},    // Free text
	{Segment: "GEI", Mandatory: false, MaxRepeat: 99},    // Processing information
	{Segment: "DIM", Mandatory: false, MaxRepeat: 99},    // Dimensions
	{Segment: "MEA", Mandatory: false, MaxRepeat: 99},    // Measurements
	{Segment: "----", Mandatory: false, MaxRepeat: 99},   // Segment group 284------------------
	{Segment: "ARD", Mandatory: true, MaxRepeat: 1},      // Monetary amount function
	{Segment: "MOA", Mandatory: false, MaxRepeat: 99},    // Monetary amount
	{Segment: "GEI", Mandatory: false, MaxRepeat: 99},    // Processing information
	{Segment: "PCD", Mandatory: false, MaxRepeat: 99},    // Percentage details
	{Segment: "----", Mandatory: false, MaxRepeat: 9999}, // Segment group 285------------------
	{Segment: "ASI", Mandatory: true, MaxRepeat: 1},      // Array structure identification
	{Segment: "CNI", Mandatory: true, MaxRepeat: 1},      // Consignment information
	{Segment: "RFF", Mandatory: false, MaxRepeat: 99},    // Reference
	{Segment: "MOA", Mandatory: false, MaxRepeat: 99},    // Monetary amount
	{Segment: "CNT", Mandatory: false, MaxRepeat: 9},     // Control total
	{Segment: "TOD", Mandatory: false, MaxRepeat: 99},    // Terms of delivery or transport
	{Segment: "GEI", Mandatory: false, MaxRepeat: 99},    // Processing information
	{Segment: "----", Mandatory: false, MaxRepeat: 99},   // Segment group 286------------------
	{Segment: "NAD", Mandatory: true, MaxRepeat: 1},      // Name and address
	{Segment: "IDE", Mandatory: false, MaxRepeat: 9},     // Identity
	{Segment: "----", Mandatory: false, MaxRepeat: 99},   // Segment group 287------------------
	{Segment: "CTA", Mandatory: true, MaxRepeat: 1},      // Contact information
	{Segment: "COM", Mandatory: false, MaxRepeat: 99},    // Communication contact
	{Segment: "----", Mandatory: false, MaxRepeat: 99},   // Segment group 288------------------
	{Segment: "LOC", Mandatory: true, MaxRepeat: 1},      // Place/location identification
	{Segment: "SEQ", Mandatory: true, MaxRepeat: 1},      // Sequence details
	{Segment: "ADR", Mandatory: false, MaxRepeat: 9},     // Address
	{Segment: "----", Mandatory: false, MaxRepeat: 99},   // Segment group 289------------------
	{Segment: "DOC", Mandatory: true, MaxRepeat: 1},      // Document/message details
	{Segment: "RFF", Mandatory: false, MaxRepeat: 99},    // Reference
	{Segment: "DTM", Mandatory: false, MaxRepeat: 99},    // Date/time/period
	{Segment: "----", Mandatory: false, MaxRepeat: 9999}, // Segment group 290------------------
	{Segment: "EQD", Mandatory: true, MaxRepeat: 1},      // Equipment details
	{Segment: "SEQ", Mandatory: true, MaxRepeat: 1},      // Sequence details
	{Segment: "QTY", Mandatory: false, MaxRepeat: 1},     // Quantity
	{Segment: "----", Mandatory: false, MaxRepeat: 99},   // Segment group 291------------------
	{Segment: "SEL", Mandatory: true, MaxRepeat: 1},      // Seal number
	{Segment: "SEQ", Mandatory: true, MaxRepeat: 1},      // Sequence details
	{Segment: "----", Mandatory: false, MaxRepeat: 99},   // Segment group 292------------------
	{Segment: "TDT", Mandatory: true, MaxRepeat: 1},      // Transport information
	{Segment: "----", Mandatory: false, MaxRepeat: 99},   // Segment group 293------------------
	{Segment: "LOC", Mandatory: true, MaxRepeat: 1},      // Place/location identification
	{Segment: "SEQ", Mandatory: true, MaxRepeat: 1},      // Sequence details
	{Segment: "----", Mandatory: false, MaxRepeat: 9999}, // Segment group 294------------------
	{Segment: "SEQ", Mandatory: true, MaxRepeat: 1},      // Sequence details
	{Segment: "RFF", Mandatory: false, MaxRepeat: 99},    // Reference
	{Segment: "TDT", Mandatory: false, MaxRepeat: 9},     // Transport information
	{Segment: "MOA", Mandatory: false, MaxRepeat: 9},     // Monetary amount
	{Segment: "MEA", Mandatory: false, MaxRepeat: 99},    // Measurements
	{Segment: "TOD", Mandatory: false, MaxRepeat: 99},    // Terms of delivery or transport
	{Segment: "LOC", Mandatory: false, MaxRepeat: 99},    // Place/location identification
	{Segment: "----", Mandatory: false, MaxRepeat: 99},   // Segment group 295------------------
	{Segment: "NAD", Mandatory: true, MaxRepeat: 1},      // Name and address
	{Segment: "IFD", Mandatory: false, MaxRepeat: 1},     // Information detail
	{Segment: "IDE", Mandatory: false, MaxRepeat: 9},     // Identity
	{Segment: "----", Mandatory: false, MaxRepeat: 99},   // Segment group 296------------------
	{Segment: "CTA", Mandatory: true, MaxRepeat: 1},      // Contact information
	{Segment: "COM", Mandatory: false, MaxRepeat: 99},    // Communication contact
	{Segment: "----", Mandatory: false, MaxRepeat: 99},   // Segment group 297------------------
	{Segment: "DOC", Mandatory: true, MaxRepeat: 1},      // Document/message details
	{Segment: "RFF", Mandatory: false, MaxRepeat: 99},    // Reference
	{Segment: "DTM", Mandatory: false, MaxRepeat: 99},    // Date/time/period
	{Segment: "----", Mandatory: false, MaxRepeat: 99},   // Segment group 298------------------
	{Segment: "RCS", Mandatory: true, MaxRepeat: 1},      // Requirements and conditions
	{Segment: "FTX", Mandatory: false, MaxRepeat: 9},     // Free text
	{Segment: "----", Mandatory: false, MaxRepeat: 999},  // Segment group 299------------------
	{Segment: "PAC", Mandatory: true, MaxRepeat: 1},      // Package
	{Segment: "SEQ", Mandatory: true, MaxRepeat: 1},      // Sequence details
	{Segment: "PCI", Mandatory: false, MaxRepeat: 99},    // Package identification
	{Segment: "----", Mandatory: false, MaxRepeat: 9999}, // Segment group 300------------------
	{Segment: "EQD", Mandatory: true, MaxRepeat: 1},      // Equipment details
	{Segment: "SEQ", Mandatory: true, MaxRepeat: 1},      // Sequence details
	{Segment: "----", Mandatory: false, MaxRepeat: 99},   // Segment group 301------------------
	{Segment: "GID", Mandatory: true, MaxRepeat: 1},      // Goods item details
	{Segment: "IMD", Mandatory: false, MaxRepeat: 99},    // Item description
	{Segment: "TCC", Mandatory: false, MaxRepeat: 99},    // Charge/rate calculations
	{Segment: "----", Mandatory: false, MaxRepeat: 99},   // Segment group 302------------------
	{Segment: "TMP", Mandatory: true, MaxRepeat: 1},      // Temperature
	{Segment: "DGS", Mandatory: false, MaxRepeat: 99},    // Dangerous goods
	{Segment: "----", Mandatory: false, MaxRepeat: 9999}, // Segment group 303------------------
	{Segment: "DMS", Mandatory: true, MaxRepeat: 1},      // Document/message summary
	{Segment: "SEQ", Mandatory: true, MaxRepeat: 1},      // Sequence details
	{Segment: "----", Mandatory: false, MaxRepeat: 99},   // Segment group 304------------------
	{Segment: "NAD", Mandatory: true, MaxRepeat: 1},      // Name and address
	{Segment: "IDE", Mandatory: false, MaxRepeat: 9},     // Identity
	{Segment: "----", Mandatory: false, MaxRepeat: 99},   // Segment group 305------------------
	{Segment: "CTA", Mandatory: true, MaxRepeat: 1},      // Contact information
	{Segment: "COM", Mandatory: false, MaxRepeat: 99},    // Communication contact
	{Segment: "----", Mandatory: false, MaxRepeat: 9999}, // Segment group 306------------------
	{Segment: "LIN", Mandatory: true, MaxRepeat: 1},      // Line item
	{Segment: "----", Mandatory: false, MaxRepeat: 99},   // Segment group 307------------------
	{Segment: "NAD", Mandatory: true, MaxRepeat: 1},      // Name and address
	{Segment: "IDE", Mandatory: false, MaxRepeat: 9},     // Identity
	{Segment: "----", Mandatory: false, MaxRepeat: 99},   // Segment group 308------------------
	{Segment: "CTA", Mandatory: true, MaxRepeat: 1},      // Contact information
	{Segment: "COM", Mandatory: false, MaxRepeat: 99},    // Communication contact
	{Segment: "----", Mandatory: false, MaxRepeat: 99},   // Segment group 309------------------
	{Segment: "GID", Mandatory: true, MaxRepeat: 1},      // Goods item details
	{Segment: "TCC", Mandatory: false, MaxRepeat: 99},    // Charge/rate calculations
	{Segment: "UNS", Mandatory: true, MaxRepeat: 1},      // Section control
	{Segment: "AUT", Mandatory: false, MaxRepeat: 1},     // Authentication result
	{Segment: "CNT", Mandatory: false, MaxRepeat: 9},     // Control total
}}
