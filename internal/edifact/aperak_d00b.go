package edifact

// APERAK D.00B message specification: the structural (segment/group
// presence, order, cardinality) branching diagram for the Application
// Error and Acknowledgement message, UN/EDIFACT directory release D.00B,
// revision 4 (2000-06-28). Scope is structural only -- not element-level
// content -- per edifact-ls-3uzr's epic non-goal. 5 segment groups, max
// nesting depth 2.
//
// Source: https://service.unece.org/trade/untdid/d00b/trmd/aperak_c.htm
// section 4.3.1 "Segment table". Unlike almost every other source in
// this project (which 403s via Cloudflare and needs the Wayback Machine
// instead), this page was fetched directly by the user in their own
// browser and saved locally -- no archive copy was needed here.
//
// Transcription note: generated the same way as every other message
// type in this project -- a script parsed the source's exact rail-art
// column positions mechanically and verified the result balances (all 5
// groups open and close consistently) before being transcribed here,
// rather than reading the ASCII nesting by eye.
//
// Notable: this tree is structurally identical, tag-for-tag and
// repeat-for-repeat, to aperak_d20a.go's D.20A schema -- APERAK's
// structure hasn't changed across 20 years and several directory
// releases. Registered as its own (Type, Version, Release, Agency) tuple
// anyway, per this project's release-specific matching design: a real
// message self-reporting "APERAK:D:00B:UN" in its UNH still needs an
// exact tuple match, even though the underlying tree happens to be
// identical to a different release's.

func init() {
	RegisterSchema(
		MessageID{Type: "APERAK", Version: "D", Release: "00B", Agency: "UN"},
		aperakD00bSchema,
		"https://service.unece.org/trade/untdid/d00b/trmd/aperak_c.htm",
	)
}

var aperakD00bSchema = Schema{Nodes: []SchemaNode{
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
