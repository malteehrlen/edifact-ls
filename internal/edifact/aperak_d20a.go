package edifact

// APERAK D.20A message specification: the structural (segment/group
// presence, order, cardinality) branching diagram for the Application Error and Acknowledgement
// message, UN/EDIFACT directory release D.20A. Scope is structural only --
// not element-level content -- per edifact-ls-3uzr's epic non-goal.
// 5 segment groups, max nesting depth 2.
//
// Source: https://service.unece.org/trade/untdid/d20a/trmd/aperak_c.htm
// section 4.3.1 "Segment table". That URL currently returns HTTP 403 from
// Cloudflare when fetched directly; this data was transcribed from the
// Wayback Machine's archived copy instead:
// http://web.archive.org/web/20240415032403/https://service.unece.org/trade/untdid/d20a/trmd/aperak_c.htm
// -- re-check the direct URL first if cross-checking later, in case the
// block has lifted.
//
// Transcription note: generated the same way as iftmcs_d21a.go -- a
// one-off script parsed the source's exact rail-art column positions
// mechanically and verified the result balances (all 5 groups
// open and close consistently) before being transcribed here, rather than
// reading the ASCII nesting by eye.

func init() {
	RegisterSchema(
		MessageID{Type: "APERAK", Version: "D", Release: "20A", Agency: "UN"},
		aperakD20aSchema,
	)
}

var aperakD20aSchema = Schema{Nodes: []SchemaNode{
	{Segment: "BGM", Mandatory: true, MaxRepeat: 1},  // Beginning of message
	{Segment: "DTM", Mandatory: false, MaxRepeat: 9}, // Date/time/period
	{Segment: "FTX", Mandatory: false, MaxRepeat: 9}, // Free text
	{Segment: "CNT", Mandatory: false, MaxRepeat: 9}, // Control total
	{ // Segment group 1
		Group: []SchemaNode{
			{Segment: "DOC", Mandatory: true, MaxRepeat: 1},   // Document/message details
			{Segment: "DTM", Mandatory: false, MaxRepeat: 99}, // Date/time/period
		},
		Mandatory: false, MaxRepeat: 99,
	},
	{ // Segment group 2
		Group: []SchemaNode{
			{Segment: "RFF", Mandatory: true, MaxRepeat: 1},  // Reference
			{Segment: "DTM", Mandatory: false, MaxRepeat: 9}, // Date/time/period
		},
		Mandatory: false, MaxRepeat: 9,
	},
	{ // Segment group 3
		Group: []SchemaNode{
			{Segment: "NAD", Mandatory: true, MaxRepeat: 1},  // Name and address
			{Segment: "CTA", Mandatory: false, MaxRepeat: 9}, // Contact information
			{Segment: "COM", Mandatory: false, MaxRepeat: 9}, // Communication contact
		},
		Mandatory: false, MaxRepeat: 9,
	},
	{ // Segment group 4
		Group: []SchemaNode{
			{Segment: "ERC", Mandatory: true, MaxRepeat: 1},  // Application error information
			{Segment: "FTX", Mandatory: false, MaxRepeat: 1}, // Free text
			{ // Segment group 5
				Group: []SchemaNode{
					{Segment: "RFF", Mandatory: true, MaxRepeat: 1},  // Reference
					{Segment: "FTX", Mandatory: false, MaxRepeat: 9}, // Free text
				},
				Mandatory: false, MaxRepeat: 9,
			},
		},
		Mandatory: false, MaxRepeat: 99999,
	},
}}
