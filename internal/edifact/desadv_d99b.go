package edifact

// DESADV D.99B message specification: the structural (segment/group
// presence, order, cardinality) branching diagram for the Despatch Advice
// message, UN/EDIFACT directory release D.99B. Scope is structural only --
// not element-level content -- per edifact-ls-3uzr's epic non-goal.
// 25 segment groups, max nesting depth 4.
//
// Release note: this message type is also registered for D.20A
// (see desadv_d20a.go). D.99B was chosen as a second, older release
// after an extended search for D.96A/D.96B (the originally-requested
// releases) turned up nothing usable: every D.96A trmd page across every
// message type checked is an identical ~1KB placeholder stub with
// exactly one Wayback capture ever; D.95B-D.98B have no archived
// captures at all; and the real D.96B pages the user found directly on
// the live site turned out to be the wrong section (segment-
// clarification narrative, not the branching diagram). D.99B is the
// same release already sourced successfully for CUSCAR
// (see cuscar_d99b.go).
// Source: https://service.unece.org/trade/untdid/d99b/trmd/desadv_c.htm
// section 4.3.1 "Segment table". That URL currently returns HTTP 403 from
// Cloudflare when fetched directly; this data was transcribed from the
// Wayback Machine's archived copy instead:
// http://web.archive.org/web/20220127003011/https://service.unece.org/trade/untdid/d99b/trmd/desadv_c.htm
// -- re-check the direct URL first if cross-checking later, in case the
// block has lifted.
//
// Transcription note: generated the same way as iftmcs_d21a.go -- a
// one-off script parsed the source's exact rail-art column positions
// mechanically and verified the result balances (all 25 groups
// open and close consistently) before being transcribed here, rather than
// reading the ASCII nesting by eye.

func init() {
	RegisterSchema(
		MessageID{Type: "DESADV", Version: "D", Release: "99B", Agency: "UN"},
		desadvD99bSchema,
		"https://service.unece.org/trade/untdid/d99b/trmd/desadv_c.htm",
	)
}

var desadvD99bSchema = Schema{Nodes: []SchemaNode{
	{Segment: "BGM", Mandatory: true, MaxRepeat: 1},   // Beginning of message
	{Segment: "DTM", Mandatory: false, MaxRepeat: 10}, // Date/time/period
	{Segment: "ALI", Mandatory: false, MaxRepeat: 5},  // Additional information
	{Segment: "MEA", Mandatory: false, MaxRepeat: 5},  // Measurements
	{Segment: "MOA", Mandatory: false, MaxRepeat: 5},  // Monetary amount
	{Segment: "CUX", Mandatory: false, MaxRepeat: 9},  // Currencies
	{ // Segment group 1
		Group: []SchemaNode{
			{Segment: "RFF", Mandatory: true, MaxRepeat: 1},  // Reference
			{Segment: "DTM", Mandatory: false, MaxRepeat: 1}, // Date/time/period
		},
		Mandatory: false, MaxRepeat: 10,
	},
	{ // Segment group 2
		Group: []SchemaNode{
			{Segment: "NAD", Mandatory: true, MaxRepeat: 1},   // Name and address
			{Segment: "LOC", Mandatory: false, MaxRepeat: 10}, // Place/location identification
			{ // Segment group 3
				Group: []SchemaNode{
					{Segment: "RFF", Mandatory: true, MaxRepeat: 1},  // Reference
					{Segment: "DTM", Mandatory: false, MaxRepeat: 1}, // Date/time/period
				},
				Mandatory: false, MaxRepeat: 10,
			},
			{ // Segment group 4
				Group: []SchemaNode{
					{Segment: "CTA", Mandatory: true, MaxRepeat: 1},  // Contact information
					{Segment: "COM", Mandatory: false, MaxRepeat: 5}, // Communication contact
				},
				Mandatory: false, MaxRepeat: 10,
			},
		},
		Mandatory: false, MaxRepeat: 99,
	},
	{ // Segment group 5
		Group: []SchemaNode{
			{Segment: "TOD", Mandatory: true, MaxRepeat: 1},  // Terms of delivery or transport
			{Segment: "LOC", Mandatory: false, MaxRepeat: 5}, // Place/location identification
			{Segment: "FTX", Mandatory: false, MaxRepeat: 5}, // Free text
		},
		Mandatory: false, MaxRepeat: 10,
	},
	{ // Segment group 6
		Group: []SchemaNode{
			{Segment: "TDT", Mandatory: true, MaxRepeat: 1},  // Details of transport
			{Segment: "PCD", Mandatory: false, MaxRepeat: 6}, // Percentage details
			{Segment: "TMD", Mandatory: false, MaxRepeat: 1}, // Transport movement details
			{ // Segment group 7
				Group: []SchemaNode{
					{Segment: "LOC", Mandatory: true, MaxRepeat: 1},   // Place/location identification
					{Segment: "DTM", Mandatory: false, MaxRepeat: 10}, // Date/time/period
				},
				Mandatory: false, MaxRepeat: 10,
			},
		},
		Mandatory: false, MaxRepeat: 10,
	},
	{ // Segment group 8
		Group: []SchemaNode{
			{Segment: "EQD", Mandatory: true, MaxRepeat: 1},   // Equipment details
			{Segment: "MEA", Mandatory: false, MaxRepeat: 5},  // Measurements
			{Segment: "SEL", Mandatory: false, MaxRepeat: 25}, // Seal number
			{Segment: "EQA", Mandatory: false, MaxRepeat: 5},  // Attached equipment
			{ // Segment group 9
				Group: []SchemaNode{
					{Segment: "HAN", Mandatory: true, MaxRepeat: 1},   // Handling instructions
					{Segment: "FTX", Mandatory: false, MaxRepeat: 10}, // Free text
				},
				Mandatory: false, MaxRepeat: 10,
			},
		},
		Mandatory: false, MaxRepeat: 10,
	},
	{ // Segment group 10
		Group: []SchemaNode{
			{Segment: "CPS", Mandatory: true, MaxRepeat: 1},  // Consignment packing sequence
			{Segment: "FTX", Mandatory: false, MaxRepeat: 5}, // Free text
			{Segment: "QVR", Mandatory: false, MaxRepeat: 9}, // Quantity variances
			{ // Segment group 11
				Group: []SchemaNode{
					{Segment: "PAC", Mandatory: true, MaxRepeat: 1},   // Package
					{Segment: "MEA", Mandatory: false, MaxRepeat: 10}, // Measurements
					{Segment: "QTY", Mandatory: false, MaxRepeat: 10}, // Quantity
					{ // Segment group 12
						Group: []SchemaNode{
							{Segment: "HAN", Mandatory: true, MaxRepeat: 1},   // Handling instructions
							{Segment: "FTX", Mandatory: false, MaxRepeat: 10}, // Free text
						},
						Mandatory: false, MaxRepeat: 10,
					},
					{ // Segment group 13
						Group: []SchemaNode{
							{Segment: "PCI", Mandatory: true, MaxRepeat: 1},  // Package identification
							{Segment: "RFF", Mandatory: false, MaxRepeat: 1}, // Reference
							{Segment: "DTM", Mandatory: false, MaxRepeat: 5}, // Date/time/period
							{ // Segment group 14
								Group: []SchemaNode{
									{Segment: "GIR", Mandatory: true, MaxRepeat: 1},  // Related identification numbers
									{Segment: "DTM", Mandatory: false, MaxRepeat: 5}, // Date/time/period
								},
								Mandatory: false, MaxRepeat: 99,
							},
							{ // Segment group 15
								Group: []SchemaNode{
									{Segment: "GIN", Mandatory: true, MaxRepeat: 1},   // Goods identity number
									{Segment: "DLM", Mandatory: false, MaxRepeat: 10}, // Delivery limitations
								},
								Mandatory: false, MaxRepeat: 99,
							},
							{ // Segment group 16
								Group: []SchemaNode{
									{Segment: "COD", Mandatory: true, MaxRepeat: 1},  // Component details
									{Segment: "MEA", Mandatory: false, MaxRepeat: 9}, // Measurements
									{Segment: "QTY", Mandatory: false, MaxRepeat: 9}, // Quantity
									{Segment: "PCD", Mandatory: false, MaxRepeat: 9}, // Percentage details
								},
								Mandatory: false, MaxRepeat: 99,
							},
						},
						Mandatory: false, MaxRepeat: 1000,
					},
				},
				Mandatory: false, MaxRepeat: 9999,
			},
			{ // Segment group 17
				Group: []SchemaNode{
					{Segment: "LIN", Mandatory: true, MaxRepeat: 1},    // Line item
					{Segment: "PIA", Mandatory: false, MaxRepeat: 10},  // Additional product id
					{Segment: "IMD", Mandatory: false, MaxRepeat: 25},  // Item description
					{Segment: "MEA", Mandatory: false, MaxRepeat: 10},  // Measurements
					{Segment: "QTY", Mandatory: false, MaxRepeat: 10},  // Quantity
					{Segment: "ALI", Mandatory: false, MaxRepeat: 10},  // Additional information
					{Segment: "GIN", Mandatory: false, MaxRepeat: 100}, // Goods identity number
					{Segment: "GIR", Mandatory: false, MaxRepeat: 100}, // Related identification numbers
					{Segment: "DLM", Mandatory: false, MaxRepeat: 100}, // Delivery limitations
					{Segment: "DTM", Mandatory: false, MaxRepeat: 5},   // Date/time/period
					{Segment: "NAD", Mandatory: false, MaxRepeat: 99},  // Name and address
					{Segment: "TDT", Mandatory: false, MaxRepeat: 1},   // Details of transport
					{Segment: "TMD", Mandatory: false, MaxRepeat: 1},   // Transport movement details
					{Segment: "HAN", Mandatory: false, MaxRepeat: 20},  // Handling instructions
					{Segment: "FTX", Mandatory: false, MaxRepeat: 99},  // Free text
					{Segment: "MOA", Mandatory: false, MaxRepeat: 5},   // Monetary amount
					{ // Segment group 18
						Group: []SchemaNode{
							{Segment: "RFF", Mandatory: true, MaxRepeat: 1},  // Reference
							{Segment: "NAD", Mandatory: false, MaxRepeat: 1}, // Name and address
							{Segment: "CTA", Mandatory: false, MaxRepeat: 1}, // Contact information
							{Segment: "DTM", Mandatory: false, MaxRepeat: 1}, // Date/time/period
						},
						Mandatory: false, MaxRepeat: 99,
					},
					{ // Segment group 19
						Group: []SchemaNode{
							{Segment: "DGS", Mandatory: true, MaxRepeat: 1},  // Dangerous goods
							{Segment: "QTY", Mandatory: false, MaxRepeat: 1}, // Quantity
							{Segment: "FTX", Mandatory: false, MaxRepeat: 5}, // Free text
						},
						Mandatory: false, MaxRepeat: 9999,
					},
					{ // Segment group 20
						Group: []SchemaNode{
							{Segment: "LOC", Mandatory: true, MaxRepeat: 1},   // Place/location identification
							{Segment: "NAD", Mandatory: false, MaxRepeat: 1},  // Name and address
							{Segment: "DTM", Mandatory: false, MaxRepeat: 1},  // Date/time/period
							{Segment: "QTY", Mandatory: false, MaxRepeat: 10}, // Quantity
						},
						Mandatory: false, MaxRepeat: 100,
					},
					{ // Segment group 21
						Group: []SchemaNode{
							{Segment: "SGP", Mandatory: true, MaxRepeat: 1},   // Split goods placement
							{Segment: "QTY", Mandatory: false, MaxRepeat: 10}, // Quantity
						},
						Mandatory: false, MaxRepeat: 1000,
					},
					{ // Segment group 22
						Group: []SchemaNode{
							{Segment: "PCI", Mandatory: true, MaxRepeat: 1},   // Package identification
							{Segment: "DTM", Mandatory: false, MaxRepeat: 5},  // Date/time/period
							{Segment: "MEA", Mandatory: false, MaxRepeat: 10}, // Measurements
							{Segment: "QTY", Mandatory: false, MaxRepeat: 1},  // Quantity
							{ // Segment group 23
								Group: []SchemaNode{
									{Segment: "GIN", Mandatory: true, MaxRepeat: 1},    // Goods identity number
									{Segment: "DLM", Mandatory: false, MaxRepeat: 100}, // Delivery limitations
								},
								Mandatory: false, MaxRepeat: 10,
							},
							{ // Segment group 24
								Group: []SchemaNode{
									{Segment: "HAN", Mandatory: true, MaxRepeat: 1},     // Handling instructions
									{Segment: "FTX", Mandatory: false, MaxRepeat: 5},    // Free text
									{Segment: "GIN", Mandatory: false, MaxRepeat: 1000}, // Goods identity number
								},
								Mandatory: false, MaxRepeat: 10,
							},
						},
						Mandatory: false, MaxRepeat: 9999,
					},
					{ // Segment group 25
						Group: []SchemaNode{
							{Segment: "QVR", Mandatory: true, MaxRepeat: 1},  // Quantity variances
							{Segment: "DTM", Mandatory: false, MaxRepeat: 5}, // Date/time/period
						},
						Mandatory: false, MaxRepeat: 10,
					},
				},
				Mandatory: false, MaxRepeat: 9999,
			},
		},
		Mandatory: false, MaxRepeat: 9999,
	},
	{Segment: "CNT", Mandatory: false, MaxRepeat: 5}, // Control total
}}
