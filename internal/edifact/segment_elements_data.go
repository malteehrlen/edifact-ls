package edifact

// Real UN/EDIFACT segment element/component structures for BGM, DTM, and
// CTA -- see edifact-ls-9ger for why this is keyed by tag alone rather
// than per message type (a segment's structure is intrinsic to it: each
// source page below lists one structure, then ~150 message types that
// all share it).
//
// Sources (each 403s via Cloudflare when fetched directly; sourced via
// the Wayback Machine instead, same caveat as iftmcs_d21a.go):
//   - BGM: https://service.unece.org/trade/untdid/d21a/trsd/trsdbgm.htm
//   - DTM: https://service.unece.org/trade/untdid/d21a/trsd/trsddtm.htm
//   - CTA: https://service.unece.org/trade/untdid/d21a/trsd/trsdcta.htm
//
// DTM is the interesting one: its only element (C507) is mandatory, and
// within it, component 2005 is mandatory too -- a real "missing
// mandatory" case. BGM and CTA are entirely conditional at every
// element/component -- confirmed deliberately, not assumed, since
// edifact-ls-7uhx already found real UN/EDIFACT data has plenty of
// "nothing here is actually mandatory" segments/groups.
//
// BGM's "Message function code" component and CTA's "Contact function
// code" component are additionally marked CodeList: their actual real
// UN/EDIFACT code lists (1225 and 3139) are registered in
// codelist_1225.go/codelist_3139.go -- see edifact-ls-6xaz for hover's
// coded-value tier and why only these two of BGM/CTA/DTM's several coded
// components are wired up so far (1001, "Document name code", is a
// ~800-entry list deliberately deferred, not attempted here).
func init() {
	RegisterSegmentElementSchema("BGM", SegmentElementSchema{Elements: []ElementSchema{
		{Name: "Document/message name", Mandatory: false, Components: []ComponentSchema{
			{Name: "Document name code", Mandatory: false},
			{Name: "Code list identification code", Mandatory: false},
			{Name: "Code list responsible agency code", Mandatory: false},
			{Name: "Document name", Mandatory: false},
		}},
		{Name: "Document/message identification", Mandatory: false, Components: []ComponentSchema{
			{Name: "Document identifier", Mandatory: false},
			{Name: "Version identifier", Mandatory: false},
			{Name: "Revision identifier", Mandatory: false},
		}},
		{Name: "Message function code", Mandatory: false, Components: []ComponentSchema{
			{Name: "Message function code", Mandatory: false, CodeList: "1225"},
		}},
		{Name: "Response type code", Mandatory: false, Components: []ComponentSchema{
			{Name: "Response type code", Mandatory: false},
		}},
		{Name: "Document status code", Mandatory: false, Components: []ComponentSchema{
			{Name: "Document status code", Mandatory: false},
		}},
		{Name: "Language name code", Mandatory: false, Components: []ComponentSchema{
			{Name: "Language name code", Mandatory: false},
		}},
	}})

	RegisterSegmentElementSchema("DTM", SegmentElementSchema{Elements: []ElementSchema{
		{Name: "Date/time/period", Mandatory: true, Components: []ComponentSchema{
			{Name: "Date or time or period function code qualifier", Mandatory: true},
			{Name: "Date or time or period text", Mandatory: false},
			{Name: "Date or time or period format code", Mandatory: false},
		}},
	}})

	RegisterSegmentElementSchema("CTA", SegmentElementSchema{Elements: []ElementSchema{
		{Name: "Contact function code", Mandatory: false, Components: []ComponentSchema{
			{Name: "Contact function code", Mandatory: false, CodeList: "3139"},
		}},
		{Name: "Contact details", Mandatory: false, Components: []ComponentSchema{
			{Name: "Contact identifier", Mandatory: false},
			{Name: "Contact name", Mandatory: false},
		}},
	}})
}
