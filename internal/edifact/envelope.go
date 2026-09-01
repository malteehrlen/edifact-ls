package edifact

import "strconv"

// messageSpan tracks one UNH..UNT message while scanning an interchange's
// flat segment list.
type messageSpan struct {
	unh          *Segment
	unt          *Segment
	segmentCount int // inclusive of UNH and, once found, UNT
}

// groupSpan tracks one UNG..UNE functional group while scanning an
// interchange's flat segment list.
type groupSpan struct {
	ung      *Segment
	une      *Segment
	messages []*messageSpan // messages (UNH..UNT spans) found inside this group
}

// ValidateEnvelopes checks interchange-, functional-group-, and
// message-level envelope structure on top of an already-parsed
// Interchange: UNB/UNZ pairing with a matching control count and control
// reference, UNG/UNE pairing (if functional grouping is used) with a
// matching message count and group reference, UNH/UNT pairing for each
// message with a matching segment count and message reference number, and
// UNS's value. It only validates structure Parse itself doesn't know about
// (which segments are semantically special) — it assumes Interchange.Segments
// already reflects whatever was syntactically parsed, valid or not.
func ValidateEnvelopes(ic *Interchange) ErrorList {
	var errs ErrorList
	d := ic.Delimiters
	interchangeStart := Position{Line: 1, Column: 1}

	var unb, unz *Segment
	var allMessages []*messageSpan
	var groups []*groupSpan
	var curGroup *groupSpan
	var curMsg *messageSpan
	sawGroupedMessage := false
	sawUngroupedMessage := false

	for i := range ic.Segments {
		seg := &ic.Segments[i]
		if curMsg != nil {
			curMsg.segmentCount++
		}

		switch seg.Tag {
		case "UNB":
			if unb != nil {
				errs.Add(seg.Pos, SeverityError, "duplicate UNB interchange header segment")
			} else {
				unb = seg
			}
		case "UNZ":
			if unz != nil {
				errs.Add(seg.Pos, SeverityError, "duplicate UNZ interchange trailer segment")
			} else {
				unz = seg
			}
		case "UNG":
			if curGroup != nil {
				errs.Add(seg.Pos, SeverityError, "UNG functional group header found before previous group (ref %q) was closed with UNE", curGroup.ung.Component0(4, d))
			}
			curGroup = &groupSpan{ung: seg}
			groups = append(groups, curGroup)
		case "UNE":
			if curGroup == nil {
				errs.Add(seg.Pos, SeverityError, "UNE functional group trailer found without a preceding UNG")
			} else {
				curGroup.une = seg
				curGroup = nil
			}
		case "UNH":
			if curMsg != nil {
				errs.Add(seg.Pos, SeverityError, "UNH message header found before previous message (ref %q) was closed with UNT", curMsg.unh.Component0(0, d))
			}
			curMsg = &messageSpan{unh: seg, segmentCount: 1}
			allMessages = append(allMessages, curMsg)
			if curGroup != nil {
				curGroup.messages = append(curGroup.messages, curMsg)
				sawGroupedMessage = true
			} else {
				sawUngroupedMessage = true
			}
		case "UNT":
			if curMsg == nil {
				errs.Add(seg.Pos, SeverityError, "UNT message trailer found without a preceding UNH")
			} else {
				curMsg.unt = seg
				curMsg = nil
			}
		case "UNS":
			if val := seg.Component0(0, d); val != "D" && val != "S" {
				errs.Add(seg.Pos, SeverityError, "UNS section control value %q, want \"D\" or \"S\"", val)
			}
		}
	}

	if sawGroupedMessage && sawUngroupedMessage {
		errs.Add(interchangeStart, SeverityError, "interchange mixes messages inside functional groups (UNG/UNE) with messages outside any group, which is not permitted")
	}

	if unb == nil {
		errs.Add(interchangeStart, SeverityError, "missing UNB interchange header segment")
	}
	if unz == nil {
		errs.Add(interchangeStart, SeverityError, "missing UNZ interchange trailer segment")
	}
	if unb != nil && unz != nil {
		wantRef := unb.Component0(4, d)
		// UNZ's control count means the number of functional groups when
		// grouping is used, or the number of messages otherwise.
		var wantCount string
		var countMeaning string
		if len(groups) > 0 {
			wantCount = strconv.Itoa(len(groups))
			countMeaning = "number of functional groups in the interchange"
		} else {
			wantCount = strconv.Itoa(len(allMessages))
			countMeaning = "number of messages in the interchange"
		}

		if gotCount := unz.Component0(0, d); gotCount != wantCount {
			errs.Add(unz.Pos, SeverityError, "UNZ interchange control count is %q, want %q (%s)", gotCount, wantCount, countMeaning)
		}
		if gotRef := unz.Component0(1, d); gotRef != wantRef {
			errs.Add(unz.Pos, SeverityError, "UNZ interchange control reference %q does not match UNB's %q", gotRef, wantRef)
		}
	}

	for _, g := range groups {
		if g.une == nil {
			errs.Add(g.ung.Pos, SeverityError, "UNG functional group (ref %q) is missing its UNE trailer", g.ung.Component0(4, d))
			continue
		}

		wantCount := strconv.Itoa(len(g.messages))
		if gotCount := g.une.Component0(0, d); gotCount != wantCount {
			errs.Add(g.une.Pos, SeverityError, "UNE number of messages is %q, want %q", gotCount, wantCount)
		}

		wantRef := g.ung.Component0(4, d)
		if gotRef := g.une.Component0(1, d); gotRef != wantRef {
			errs.Add(g.une.Pos, SeverityError, "UNE functional group reference %q does not match UNG's %q", gotRef, wantRef)
		}
	}

	for _, m := range allMessages {
		if m.unt == nil {
			errs.Add(m.unh.Pos, SeverityError, "UNH message (ref %q) is missing its UNT trailer", m.unh.Component0(0, d))
			continue
		}

		wantCount := strconv.Itoa(m.segmentCount)
		if gotCount := m.unt.Component0(0, d); gotCount != wantCount {
			errs.Add(m.unt.Pos, SeverityError, "UNT segment count is %q, want %q (segments from UNH to UNT inclusive)", gotCount, wantCount)
		}

		wantRef := m.unh.Component0(0, d)
		if gotRef := m.unt.Component0(1, d); gotRef != wantRef {
			errs.Add(m.unt.Pos, SeverityError, "UNT message reference %q does not match UNH's %q", gotRef, wantRef)
		}
	}

	return errs
}
