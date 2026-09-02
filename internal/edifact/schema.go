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
	numbering := numberGroups(schema.Nodes, new(int))
	ctx := &matchContext{errs: &errs}
	// insideGroup=false: at the true top level there's no enclosing
	// repeat that could legitimately reclaim an over-cap match, so
	// overflow is checked everywhere, including the first node.
	pos := matchSequence(schema.Nodes, numbering, segments, 0, atEnd, false, nil, ctx)
	for pos < len(segments) {
		seg := segments[pos]
		errs.Add(seg.Pos, SeverityWarning, "segment tag %q is not expected here per the message specification", seg.Tag)
		pos++
	}
	return errs
}

// GroupPathAt returns the sequence of segment-group numbers (outermost
// first) that segments[targetIndex] falls within per schema -- the same
// numbering UN/EDIFACT branching diagrams themselves use, each group
// numbered sequentially in a single structural pass over the tree,
// independent of how many times it repeats at runtime (see numberGroups).
// Returns nil if targetIndex is out of range or the segment there isn't
// inside any group (including a top-level segment, or one the schema
// doesn't actually match). Used by hover to show a segment occurrence's
// message-specific context -- see edifact-ls-pcm0.
func GroupPathAt(schema Schema, segments []Segment, targetIndex int) []int {
	if targetIndex < 0 || targetIndex >= len(segments) {
		return nil
	}
	numbering := numberGroups(schema.Nodes, new(int))
	var discard ErrorList
	var result []int
	ctx := &matchContext{errs: &discard, visit: func(segIndex int, groupPath []int) {
		if segIndex == targetIndex && len(groupPath) > 0 {
			result = append([]int{}, groupPath...)
		}
	}}
	matchSequence(schema.Nodes, numbering, segments, 0, Position{}, false, nil, ctx)
	return result
}

// groupNumbering mirrors a []SchemaNode's shape index-for-index, holding
// the segment-group number assigned to each Group node (unused/zero for a
// segment leaf) -- see numberGroups.
type groupNumbering struct {
	Number   int
	Children []groupNumbering
}

// numberGroups assigns each group in nodes (recursively) the sequential
// number UN/EDIFACT branching diagrams themselves use, matching the
// "Segment group N" numbering already visible as comments in this
// project's generated schema-data source (e.g. iftmcs_d21a.go): a single
// structural preorder pass over the tree, counting *tree positions*, not
// runtime matches. This has to be computed once, up front, and then only
// ever looked up (never recomputed) during real segment matching --
// recomputing it inside a repeating group's own match loop would
// reassign a fresh, wrong number to every nested group each time the
// enclosing group repeats, instead of the one fixed number that group
// actually has in the spec.
func numberGroups(nodes []SchemaNode, counter *int) []groupNumbering {
	result := make([]groupNumbering, len(nodes))
	for i, n := range nodes {
		if n.Group == nil {
			continue
		}
		*counter++
		// Captured into a local before the recursive call below: Go
		// doesn't guarantee a struct literal's field expressions
		// evaluate left-to-right the way function-call arguments do, so
		// inlining *counter directly into the Number field risked
		// reading it *after* the recursive call had already advanced it
		// further for any nested groups -- which is exactly what
		// happened here until this was pulled out explicitly.
		num := *counter
		result[i] = groupNumbering{Number: num, Children: numberGroups(n.Group, counter)}
	}
	return result
}

// matchContext carries state threaded through matchSequence/matchOnce
// that's shared across a whole matching walk but doesn't vary per node:
// where to report diagnostics, and an optional callback letting a caller
// (GroupPathAt) observe every real segment occurrence matched along with
// its enclosing group-number path, without a second, separately
// maintained matching implementation that could drift from this one.
type matchContext struct {
	errs  *ErrorList
	visit func(segIndex int, groupPath []int)
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
//
// The same kind of ambiguity also arises between *siblings* later in
// this same children list: two different schema nodes that happen to
// share a leading tag (e.g. a message with two independent UNS
// occurrences marking the header/detail and detail/summary boundaries,
// or -- found while bulk-sourcing edifact-ls-13gu -- two differently-
// purposed segment groups that both happen to lead with the same
// segment, such as several party-detail groups each starting with NAD).
// These aren't necessarily immediately adjacent -- other, conditional
// siblings the current message instance doesn't happen to use can sit
// between them (e.g. UNS ... [optional group] ... UNS) -- so a child at
// its own cap checks every remaining sibling, not just the next one,
// for a shared leading tag before concluding a further match is
// genuinely its own overflow rather than the legitimate start of a
// later sibling's occurrence.
func matchSequence(children []SchemaNode, numbering []groupNumbering, segments []Segment, pos int, atEnd Position, insideGroup bool, groupPath []int, ctx *matchContext) int {
	for i, child := range children {
		// childPath is only meaningful (non-nil) when child is itself a
		// group -- it's groupPath extended with this tree position's own
		// number, precomputed in numbering rather than counted live here,
		// so it stays the same across every repeat occurrence of child.
		var childPath []int
		if child.Group != nil {
			childPath = append(append([]int{}, groupPath...), numbering[i].Number)
		}

		repeats := 0
		for repeats < child.MaxRepeat && matchesLeading(child, segments, pos) {
			pos = matchOnce(child, numbering[i].Children, segments, pos, atEnd, groupPath, childPath, ctx)
			repeats++
		}

		if repeats == 0 && child.Mandatory {
			ctx.errs.Add(posOrEnd(segments, pos, atEnd), SeverityError, "missing mandatory %s", describe(child))
		}

		checkOverflow := !(insideGroup && i == 0) && !laterSiblingSharesLeadingTag(children, i)
		if checkOverflow && repeats == child.MaxRepeat && matchesLeading(child, segments, pos) {
			ctx.errs.Add(segments[pos].Pos, SeverityError, "%s repeats more than the maximum of %d time(s) allowed", describe(child), child.MaxRepeat)
			// Consume the excess occurrences too, so they aren't
			// misinterpreted as belonging to whatever schema node
			// comes next.
			for matchesLeading(child, segments, pos) {
				pos = matchOnce(child, numbering[i].Children, segments, pos, atEnd, groupPath, childPath, ctx)
			}
		}
	}
	return pos
}

// laterSiblingSharesLeadingTag reports whether any of children[i+1:] has
// the same leading tag as children[i]. Scanning every remaining sibling
// (not just the immediate next one) is what lets this see past
// conditional siblings that a given message instance doesn't happen to
// use -- e.g. UNS ... [an unused optional group] ... UNS still resolves
// correctly, since the second UNS is still found even though it isn't
// textually adjacent to the first.
func laterSiblingSharesLeadingTag(children []SchemaNode, i int) bool {
	tag := leadingTag(children[i])
	for _, sibling := range children[i+1:] {
		if leadingTag(sibling) == tag {
			return true
		}
	}
	return false
}

// matchOnce consumes exactly one occurrence of node, which the caller has
// already established matches at segments[pos] via matchesLeading.
// groupPath is the path leading to node's own children list (used for a
// leaf's visit callback, since a leaf doesn't add to the path itself);
// childPath is groupPath plus node's own number, used only when
// recursing into node.Group.
func matchOnce(node SchemaNode, numbering []groupNumbering, segments []Segment, pos int, atEnd Position, groupPath, childPath []int, ctx *matchContext) int {
	if node.Group == nil {
		if ctx.visit != nil {
			ctx.visit(pos, groupPath)
		}
		return pos + 1
	}
	return matchSequence(node.Group, numbering, segments, pos, atEnd, true, childPath, ctx)
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
