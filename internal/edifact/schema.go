package edifact

import "fmt"

// Schema describes one UN/EDIFACT message type's expected body structure --
// the ordered branching diagram a UNSM publishes in its "segment table"
// (position, tag, mandatory/conditional, max repeat, nesting), independent
// of which message type it actually belongs to. UNH and UNT themselves are
// not part of a Schema; ValidateSchema checks only the segments between
// them.
type Schema struct {
	Nodes []SchemaNode
}

// SchemaNode is one entry in a branching diagram: either a segment slot
// (Segment set, Group nil) or a nested segment group (Group set, Segment
// ""). Exactly one of the two is set. A group's own leading tag -- used to
// recognize where an occurrence of the group starts -- is taken from its
// first child, per the UN/EDIFACT convention that a group always opens
// with a mandatory segment.
type SchemaNode struct {
	Segment string
	Group   []SchemaNode

	Mandatory bool
	MaxRepeat int // must be >= 1
}

// ValidateSchema checks a message's body segments (the segments strictly
// between its UNH and UNT) against schema, reporting missing mandatory
// segments/groups, exceeded repeat counts, and segments left over once the
// schema is exhausted. atEnd is the position attributed to a "missing
// mandatory" diagnostic when the segment stream runs out before the
// schema does -- callers wiring this to a real message should pass the
// UNT segment's position.
func ValidateSchema(schema Schema, segments []Segment, atEnd Position) ErrorList {
	var errs ErrorList
	// insideGroup=false: at the true top level there's no enclosing
	// repeat that could legitimately reclaim an over-cap match, so
	// overflow is checked everywhere, including the first node.
	pos := matchSequence(schema.Nodes, segments, 0, atEnd, false, &errs)
	for pos < len(segments) {
		seg := segments[pos]
		errs.Add(seg.Pos, SeverityWarning, "segment tag %q is not expected here per the message specification", seg.Tag)
		pos++
	}
	return errs
}

// matchSequence walks children in order starting at segments[pos],
// consuming as many matching occurrences of each as its MaxRepeat allows,
// and returns the position just past everything it consumed.
//
// insideGroup is true when children is a group's own child list (as
// opposed to the schema's top-level nodes). In that case the first child
// is exempt from local "exceeded repeat" detection: its leading tag is,
// by construction, also the enclosing group's own leading tag (see
// leadingTag), so a match beyond the child's own cap is ambiguous -- it
// might just as well be the start of a fresh occurrence of the whole
// group. That ambiguity is resolved one level up, where the enclosing
// call checks the *group's* MaxRepeat instead; double-checking it here
// too would misreport a legitimate new group occurrence as this leaf
// overflowing.
func matchSequence(children []SchemaNode, segments []Segment, pos int, atEnd Position, insideGroup bool, errs *ErrorList) int {
	for i, child := range children {
		repeats := 0
		for repeats < child.MaxRepeat && matchesLeading(child, segments, pos) {
			pos = matchOnce(child, segments, pos, atEnd, errs)
			repeats++
		}

		if repeats == 0 && child.Mandatory {
			errs.Add(posOrEnd(segments, pos, atEnd), SeverityError, "missing mandatory %s", describe(child))
		}

		checkOverflow := !(insideGroup && i == 0)
		if checkOverflow && repeats == child.MaxRepeat && matchesLeading(child, segments, pos) {
			errs.Add(segments[pos].Pos, SeverityError, "%s repeats more than the maximum of %d time(s) allowed", describe(child), child.MaxRepeat)
			// Consume the excess occurrences too, so they aren't
			// misinterpreted as belonging to whatever schema node
			// comes next.
			for matchesLeading(child, segments, pos) {
				pos = matchOnce(child, segments, pos, atEnd, errs)
			}
		}
	}
	return pos
}

// matchOnce consumes exactly one occurrence of node, which the caller has
// already established matches at segments[pos] via matchesLeading.
func matchOnce(node SchemaNode, segments []Segment, pos int, atEnd Position, errs *ErrorList) int {
	if node.Group == nil {
		return pos + 1
	}
	return matchSequence(node.Group, segments, pos, atEnd, true, errs)
}

// matchesLeading reports whether segments[pos], if present, could start an
// occurrence of node.
func matchesLeading(node SchemaNode, segments []Segment, pos int) bool {
	if pos >= len(segments) {
		return false
	}
	return segments[pos].Tag == leadingTag(node)
}

// leadingTag is the segment tag that identifies the start of an occurrence
// of node: node's own tag if it's a segment slot, or its first child's
// leading tag (recursively) if it's a group.
func leadingTag(node SchemaNode) string {
	if node.Group == nil {
		return node.Segment
	}
	if len(node.Group) == 0 {
		return ""
	}
	return leadingTag(node.Group[0])
}

func describe(node SchemaNode) string {
	if node.Group == nil {
		return fmt.Sprintf("segment %q", node.Segment)
	}
	return fmt.Sprintf("segment group starting with %q", leadingTag(node))
}

func posOrEnd(segments []Segment, pos int, atEnd Position) Position {
	if pos < len(segments) {
		return segments[pos].Pos
	}
	return atEnd
}
