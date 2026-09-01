package edifact

// SegmentInfo is a short, human-readable description of what a segment
// tag means, independent of any specific message type -- used for hover.
type SegmentInfo struct {
	Name        string // the UNSD segment name, e.g. "Beginning of message"
	Description string // one more sentence on the segment's purpose
}

// segmentDescriptions is a lookup table of documented segment tags.
// Deliberately incomplete: it covers the service segments this project
// already recognizes (see knownServiceSegmentTags in lint.go) plus the
// business segments this project's own fixtures and IFTMCS schema already
// use -- meant to grow incrementally as more segments come up, not to
// enumerate the full UNSD from day one.
//
// Source: the UN/EDIFACT Segment Directory (UNSD). Each segment has its
// own definition page following the pattern
// https://service.unece.org/trade/untdid/d21a/trsd/trsd<tag>.htm (tag
// lowercased) -- confirmed for BGM/CTA/COM/DTM/TSR/CUX/MOA/FTX/CNT/GDS/
// LOC/RFF/NAD/UNH/UNT while sourcing edifact-ls-3uzr's IFTMCS D.21A schema
// (see iftmcs_d21a.go for the same domain's Cloudflare/Wayback caveat).
// UNA/UNB/UNZ/UNG/UNE/UNS are interchange/group envelope segments rather
// than message-body ones, so they don't appear in IFTMCS's own segment
// table; their trsd URLs below follow the identical pattern but weren't
// individually confirmed the same way -- worth checking first if
// cross-checking those six specifically.
var segmentDescriptions = map[string]SegmentInfo{
	"UNA": {"Service string advice", "Specifies the service characters (delimiters) used to compose the interchange, overriding the documented defaults."},
	"UNB": {"Interchange header", "Starts, identifies, and describes an interchange: sender, recipient, date/time, and control reference."},
	"UNZ": {"Interchange trailer", "Ends an interchange, carrying a count of the messages or groups it contains and its control reference."},
	"UNG": {"Group header", "Starts and identifies a functional group of messages of the same type."},
	"UNE": {"Group trailer", "Ends a functional group, carrying a count of the messages it contains and its group reference."},
	"UNH": {"Message header", "Starts and uniquely identifies a message, giving its type, version, release, and controlling agency."},
	"UNT": {"Message trailer", "Ends a message, carrying its segment count and message reference number."},
	"UNS": {"Section control", "Separates the header, detail, and summary sections of a message."},
	"BGM": {"Beginning of message", "Indicates the type and function of a message and transmits its identifying number."},
	"DTM": {"Date/time/period", "Specifies a date, a time, and/or a period."},
	"NAD": {"Name and address", "Specifies a party's name and/or address, and the function it serves (e.g. buyer, consignee)."},
	"LOC": {"Place/location identification", "Identifies a place or location relevant to the segment or message it appears in."},
	"RFF": {"Reference", "Specifies a reference, such as a document or message number, optionally with a date."},
	"MOA": {"Monetary amount", "Specifies a monetary amount."},
	"FTX": {"Free text", "Provides free-form or coded supplementary text."},
	"TSR": {"Transport service requirements", "Specifies transport service, priority, and condition requirements."},
	"CUX": {"Currencies", "Specifies currencies and related information relevant to a message or transaction."},
	"CNT": {"Control total", "Provides a control total, such as a line count or total quantity, for verification."},
	"GDS": {"Nature of cargo", "Indicates, in coded or free-text form, the type of cargo being carried."},
	"CTA": {"Contact information", "Identifies a person or department that communication should be directed to."},
}

// SegmentDescription returns the description for tag, and whether one is
// known. An unknown tag returns a zero SegmentInfo and false.
func SegmentDescription(tag string) (SegmentInfo, bool) {
	info, ok := segmentDescriptions[tag]
	return info, ok
}
