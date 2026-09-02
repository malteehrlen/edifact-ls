package edifact

// CUSCAR D.99B message specification: the structural (segment/group
// presence, order, cardinality) branching diagram for the Customs Cargo
// Report message, UN/EDIFACT directory release D.99B. Scope is
// structural only -- not element-level content -- per edifact-ls-3uzr's
// epic non-goal. 17 segment groups, max nesting depth 4.
//
// Release note: requested by the user against a real file declaring
// "CUSCAR:D:96B:UN" -- but D.96B itself isn't archived anywhere, and the
// one archived D.96A snapshot found turned out to be a stub page with no
// real segment table (just placeholder text). D.99B and D.01B were the
// nearest releases with real, complete segment tables; the user chose
// D.99B. This registers for the exact tuple (CUSCAR, D, 99B, UN) --
// per this project's registry design, it will *not* structurally
// validate a message declaring D:96B specifically (different release),
// though it will still produce the "recognized type, different release
// registered" info diagnostic for one rather than silence.
//
// Source: https://service.unece.org/trade/untdid/d99b/trmd/cuscar_c.htm
// section 4.3.1 "Segment table". That URL currently returns HTTP 403 from
// Cloudflare when fetched directly; this data was transcribed from the
// Wayback Machine's archived copy instead:
// http://web.archive.org/web/20220126231724/https://service.unece.org/trade/untdid/d99b/trmd/cuscar_c.htm
// -- re-check the direct URL first if cross-checking later, in case the
// block has lifted.
//
// Transcription note: generated the same way as iftmcs_d21a.go -- a
// one-off script parsed the source's exact rail-art column positions
// mechanically and verified the result balances (all 17 groups open and
// close consistently) before being transcribed here. This particular
// page's table format differs slightly from the D20A/D21A pages the
// script was first written against -- 4-digit position numbers instead
// of 5, and this page actually uses the "change indicator" character
// column (e.g. "0110 + DTM ...") that those pages' tables happened not
// to exercise -- both handled by generalizing the script's parsing
// regexes rather than hand-patching this file's data.

func init() {
	RegisterSchema(
		MessageID{Type: "CUSCAR", Version: "D", Release: "99B", Agency: "UN"},
		cuscarD99bSchema,
		"https://service.unece.org/trade/untdid/d99b/trmd/cuscar_c.htm",
	)
}

var cuscarD99bSchema = Schema{Nodes: []SchemaNode{
	{Segment: "BGM", Mandatory: true, MaxRepeat: 1},  // Beginning of message
	{Segment: "DTM", Mandatory: false, MaxRepeat: 9}, // Date/time/period
	{Segment: "LOC", Mandatory: false, MaxRepeat: 9}, // Place/location identification
	{Segment: "QTY", Mandatory: false, MaxRepeat: 9}, // Quantity
	{ // Segment group 1
		Group: []SchemaNode{
			{Segment: "RFF", Mandatory: true, MaxRepeat: 1},  // Reference
			{Segment: "DTM", Mandatory: false, MaxRepeat: 9}, // Date/time/period
		},
		Mandatory: false, MaxRepeat: 99,
	},
	{ // Segment group 2
		Group: []SchemaNode{
			{Segment: "NAD", Mandatory: true, MaxRepeat: 1},  // Name and address
			{Segment: "DTM", Mandatory: false, MaxRepeat: 1}, // Date/time/period
			{ // Segment group 3
				Group: []SchemaNode{
					{Segment: "CTA", Mandatory: true, MaxRepeat: 1},  // Contact information
					{Segment: "COM", Mandatory: false, MaxRepeat: 9}, // Communication contact
				},
				Mandatory: false, MaxRepeat: 9,
			},
		},
		Mandatory: false, MaxRepeat: 99,
	},
	{Segment: "FTX", Mandatory: false, MaxRepeat: 9}, // Free text
	{ // Segment group 4
		Group: []SchemaNode{
			{Segment: "TDT", Mandatory: true, MaxRepeat: 1},   // Details of transport
			{Segment: "LOC", Mandatory: false, MaxRepeat: 99}, // Place/location identification
			{Segment: "DTM", Mandatory: false, MaxRepeat: 99}, // Date/time/period
		},
		Mandatory: false, MaxRepeat: 9,
	},
	{Segment: "GIS", Mandatory: false, MaxRepeat: 9}, // General indicator
	{ // Segment group 5
		Group: []SchemaNode{
			{Segment: "EQD", Mandatory: true, MaxRepeat: 1},  // Equipment details
			{Segment: "TSR", Mandatory: false, MaxRepeat: 9}, // Transport service requirements
			{Segment: "MEA", Mandatory: false, MaxRepeat: 9}, // Measurements
			{Segment: "DIM", Mandatory: false, MaxRepeat: 9}, // Dimensions
			{Segment: "SEL", Mandatory: false, MaxRepeat: 9}, // Seal number
			{Segment: "NAD", Mandatory: false, MaxRepeat: 9}, // Name and address
			{Segment: "GIS", Mandatory: false, MaxRepeat: 9}, // General indicator
			{Segment: "LOC", Mandatory: false, MaxRepeat: 1}, // Place/location identification
			{ // Segment group 6
				Group: []SchemaNode{
					{Segment: "TMP", Mandatory: true, MaxRepeat: 1},  // Temperature
					{Segment: "RNG", Mandatory: false, MaxRepeat: 1}, // Range details
				},
				Mandatory: false, MaxRepeat: 9,
			},
		},
		Mandatory: false, MaxRepeat: 9999,
	},
	{Segment: "CNT", Mandatory: false, MaxRepeat: 9}, // Control total
	{ // Segment group 7
		Group: []SchemaNode{
			{Segment: "CNI", Mandatory: true, MaxRepeat: 1},  // Consignment information
			{Segment: "CNT", Mandatory: false, MaxRepeat: 9}, // Control total
			{ // Segment group 8
				Group: []SchemaNode{
					{Segment: "RFF", Mandatory: true, MaxRepeat: 1},   // Reference
					{Segment: "CNT", Mandatory: false, MaxRepeat: 9},  // Control total
					{Segment: "MOA", Mandatory: false, MaxRepeat: 9},  // Monetary amount
					{Segment: "LOC", Mandatory: false, MaxRepeat: 99}, // Place/location identification
					{Segment: "GIS", Mandatory: false, MaxRepeat: 9},  // General indicator
					{Segment: "CUX", Mandatory: false, MaxRepeat: 9},  // Currencies
					{Segment: "CPI", Mandatory: false, MaxRepeat: 9},  // Charge payment instructions
					{ // Segment group 9
						Group: []SchemaNode{
							{Segment: "TDT", Mandatory: true, MaxRepeat: 1},  // Details of transport
							{Segment: "LOC", Mandatory: false, MaxRepeat: 9}, // Place/location identification
							{Segment: "DTM", Mandatory: false, MaxRepeat: 9}, // Date/time/period
							{Segment: "MEA", Mandatory: false, MaxRepeat: 9}, // Measurements
							{ // Segment group 10
								Group: []SchemaNode{
									{Segment: "RFF", Mandatory: true, MaxRepeat: 1},  // Reference
									{Segment: "TSR", Mandatory: false, MaxRepeat: 9}, // Transport service requirements
								},
								Mandatory: false, MaxRepeat: 9,
							},
						},
						Mandatory: false, MaxRepeat: 9,
					},
					{ // Segment group 11
						Group: []SchemaNode{
							{Segment: "NAD", Mandatory: true, MaxRepeat: 1},  // Name and address
							{Segment: "DTM", Mandatory: false, MaxRepeat: 9}, // Date/time/period
							{Segment: "RFF", Mandatory: false, MaxRepeat: 9}, // Reference
							{ // Segment group 12
								Group: []SchemaNode{
									{Segment: "CTA", Mandatory: true, MaxRepeat: 1},  // Contact information
									{Segment: "COM", Mandatory: false, MaxRepeat: 9}, // Communication contact
								},
								Mandatory: false, MaxRepeat: 9,
							},
						},
						Mandatory: false, MaxRepeat: 99,
					},
					{ // Segment group 13
						Group: []SchemaNode{
							{Segment: "QTY", Mandatory: true, MaxRepeat: 1},  // Quantity
							{Segment: "FTX", Mandatory: false, MaxRepeat: 9}, // Free text
						},
						Mandatory: false, MaxRepeat: 9,
					},
					{ // Segment group 14
						Group: []SchemaNode{
							{Segment: "GID", Mandatory: true, MaxRepeat: 1},     // Goods item details
							{Segment: "PAC", Mandatory: false, MaxRepeat: 9},    // Package
							{Segment: "HAN", Mandatory: false, MaxRepeat: 9},    // Handling instructions
							{Segment: "FTX", Mandatory: false, MaxRepeat: 99},   // Free text
							{Segment: "MEA", Mandatory: false, MaxRepeat: 99},   // Measurements
							{Segment: "MOA", Mandatory: false, MaxRepeat: 9},    // Monetary amount
							{Segment: "SGP", Mandatory: false, MaxRepeat: 9999}, // Split goods placement
							{Segment: "DGS", Mandatory: false, MaxRepeat: 9},    // Dangerous goods
							{Segment: "PCI", Mandatory: false, MaxRepeat: 9},    // Package identification
							{Segment: "CST", Mandatory: false, MaxRepeat: 1},    // Customs status of goods
							{Segment: "LOC", Mandatory: false, MaxRepeat: 9},    // Place/location identification
							{Segment: "TMD", Mandatory: false, MaxRepeat: 9},    // Transport movement details
							{ // Segment group 15
								Group: []SchemaNode{
									{Segment: "GIS", Mandatory: true, MaxRepeat: 1},  // General indicator
									{Segment: "DTM", Mandatory: false, MaxRepeat: 9}, // Date/time/period
									{Segment: "DOC", Mandatory: false, MaxRepeat: 9}, // Document/message details
									{Segment: "PAC", Mandatory: false, MaxRepeat: 9}, // Package
									{Segment: "MEA", Mandatory: false, MaxRepeat: 9}, // Measurements
								},
								Mandatory: false, MaxRepeat: 9,
							},
							{ // Segment group 16
								Group: []SchemaNode{
									{Segment: "QTY", Mandatory: true, MaxRepeat: 1},  // Quantity
									{Segment: "FTX", Mandatory: false, MaxRepeat: 1}, // Free text
								},
								Mandatory: false, MaxRepeat: 9,
							},
						},
						Mandatory: true, MaxRepeat: 9999,
					},
				},
				Mandatory: false, MaxRepeat: 9999,
			},
		},
		Mandatory: false, MaxRepeat: 9999,
	},
	{ // Segment group 17
		Group: []SchemaNode{
			{Segment: "AUT", Mandatory: true, MaxRepeat: 1},  // Authentication result
			{Segment: "DTM", Mandatory: false, MaxRepeat: 9}, // Date/time/period
			{Segment: "RFF", Mandatory: false, MaxRepeat: 9}, // Reference
		},
		Mandatory: false, MaxRepeat: 1,
	},
}}
