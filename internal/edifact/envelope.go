package edifact

import "strconv"

// messageSpan tracks one UNH..UNT message while scanning an interchange's
// flat segment list.
type messageSpan struct {
	unh          *Segment
	unt          *Segment
	segmentCount int // inclusive of UNH and, once found, UNT
}

// ValidateEnvelopes checks interchange- and message-level envelope
// structure on top of an already-parsed Interchange: UNB/UNZ pairing with a
// matching control count and control reference, and, for each message,
// UNH/UNT pairing with a matching segment count and message reference
// number. It only validates structure Parse itself doesn't know about
// (which segments are semantically special) — it assumes Interchange.Segments
// already reflects whatever was syntactically parsed, valid or not.
func ValidateEnvelopes(ic *Interchange) ErrorList {
	var errs ErrorList
	d := ic.Delimiters
	interchangeStart := Position{Line: 1, Column: 1}

	var unb, unz *Segment
	var messages []*messageSpan
	var cur *messageSpan

	for i := range ic.Segments {
		seg := &ic.Segments[i]
		if cur != nil {
			cur.segmentCount++
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
		case "UNH":
			if cur != nil {
				errs.Add(seg.Pos, SeverityError, "UNH message header found before previous message (ref %q) was closed with UNT", cur.unh.Component0(0, d))
			}
			cur = &messageSpan{unh: seg, segmentCount: 1}
			messages = append(messages, cur)
		case "UNT":
			if cur == nil {
				errs.Add(seg.Pos, SeverityError, "UNT message trailer found without a preceding UNH")
			} else {
				cur.unt = seg
				cur = nil
			}
		}
	}

	if unb == nil {
		errs.Add(interchangeStart, SeverityError, "missing UNB interchange header segment")
	}
	if unz == nil {
		errs.Add(interchangeStart, SeverityError, "missing UNZ interchange trailer segment")
	}
	if unb != nil && unz != nil {
		wantRef := unb.Component0(4, d)
		wantCount := strconv.Itoa(len(messages))

		if gotCount := unz.Component0(0, d); gotCount != wantCount {
			errs.Add(unz.Pos, SeverityError, "UNZ interchange control count is %q, want %q (number of messages in the interchange)", gotCount, wantCount)
		}
		if gotRef := unz.Component0(1, d); gotRef != wantRef {
			errs.Add(unz.Pos, SeverityError, "UNZ interchange control reference %q does not match UNB's %q", gotRef, wantRef)
		}
	}

	for _, m := range messages {
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
