package edifact

import (
	"fmt"
	"testing"
)

// These tests validate every currently-registered structural schema
// mechanically, derived directly from each schema's own tree, rather than
// hand-writing bespoke conformant/violation fixtures per message type --
// impractical once the registry holds dozens of message types across
// several releases. They complement, not replace, the per-schema
// transcription checks: extract_msgtype.py's own balance verification
// (every group's rails must close consistently) plus manual spot-checks
// against the real source assure the schema *data* is faithful to
// UNECE's table; these generic tests instead assure the *engine*
// correctly accepts what a schema claims is minimally required and
// correctly rejects what's missing or over-repeated -- for every
// registered schema, automatically, including ones registered after
// these tests were written.

// minimalMandatoryTags walks nodes and returns the tag sequence of the
// smallest message that could possibly satisfy them: one occurrence of
// each unconditionally-mandatory node, recursing into a mandatory
// group's own children the same way (a mandatory group's own optional
// children are, by definition, not required for the group itself to be
// present).
func minimalMandatoryTags(nodes []SchemaNode) []string {
	var tags []string
	for _, n := range nodes {
		if !n.Mandatory {
			continue
		}
		if n.Group == nil {
			tags = append(tags, n.Segment)
		} else {
			tags = append(tags, minimalMandatoryTags(n.Group)...)
		}
	}
	return tags
}

func segsFromTags(tags []string) []Segment {
	segs := make([]Segment, len(tags))
	for i, tag := range tags {
		segs[i] = seg(tag, i+1)
	}
	return segs
}

func hasErrorSeverity(errs ErrorList) bool {
	for _, e := range errs {
		if e.Severity == SeverityError {
			return true
		}
	}
	return false
}

// TestAllRegisteredSchemasAcceptMinimalConformantMessage confirms that
// for every registered schema, the message built from exactly its own
// unconditionally-mandatory tags validates clean -- i.e. the schema
// doesn't (say) require something as mandatory that the matcher then
// refuses to accept, or vice versa.
func TestAllRegisteredSchemasAcceptMinimalConformantMessage(t *testing.T) {
	for _, info := range ListRegisteredSchemas() {
		id := info.ID
		t.Run(fmt.Sprintf("%s_%s_%s_%s", id.Type, id.Version, id.Release, id.Agency), func(t *testing.T) {
			rs := schemaRegistry[id]
			tags := minimalMandatoryTags(rs.Schema.Nodes)
			errs := ValidateSchema(rs.Schema, segsFromTags(tags), testEnd)
			if len(errs) != 0 {
				t.Errorf("minimal message %v unexpectedly produced diagnostics: %v", tags, errs)
			}
		})
	}
}

// TestAllRegisteredSchemasFlagMissingMandatoryTag confirms that removing
// the last tag from a schema's own minimal-conformant sequence is
// flagged -- proving every tag in that derived sequence is actually
// enforced, not silently optional in practice.
func TestAllRegisteredSchemasFlagMissingMandatoryTag(t *testing.T) {
	for _, info := range ListRegisteredSchemas() {
		id := info.ID
		t.Run(fmt.Sprintf("%s_%s_%s_%s", id.Type, id.Version, id.Release, id.Agency), func(t *testing.T) {
			rs := schemaRegistry[id]
			tags := minimalMandatoryTags(rs.Schema.Nodes)
			if len(tags) == 0 {
				t.Skipf("%v has no unconditionally mandatory top-level tags to remove", id)
			}
			truncated := tags[:len(tags)-1]
			errs := ValidateSchema(rs.Schema, segsFromTags(truncated), testEnd)
			if !hasErrorSeverity(errs) {
				t.Errorf("dropping the last mandatory tag %q from %v produced no error-severity diagnostic (got %v)", tags[len(tags)-1], tags, errs)
			}
		})
	}
}

// TestAllRegisteredSchemasFlagExceededTopLevelRepeat covers the schemas
// whose very first top-level node has a small enough MaxRepeat to
// synthesize an over-cap message from that node alone (most real
// messages' leading segment, e.g. BGM, is capped at 1) -- confirming the
// repeat-overflow diagnostic actually fires, not just the missing-
// mandatory one.
func TestAllRegisteredSchemasFlagExceededTopLevelRepeat(t *testing.T) {
	const maxSynthesizableRepeat = 5

	for _, info := range ListRegisteredSchemas() {
		id := info.ID
		rs := schemaRegistry[id]
		if len(rs.Schema.Nodes) == 0 {
			continue
		}
		first := rs.Schema.Nodes[0]
		if first.MaxRepeat > maxSynthesizableRepeat {
			continue
		}
		t.Run(fmt.Sprintf("%s_%s_%s_%s", id.Type, id.Version, id.Release, id.Agency), func(t *testing.T) {
			tag := leadingTag(first)
			var tags []string
			for i := 0; i < first.MaxRepeat+1; i++ {
				tags = append(tags, tag)
			}
			errs := ValidateSchema(rs.Schema, segsFromTags(tags), testEnd)
			want := fmt.Sprintf("maximum of %d", first.MaxRepeat)
			if !containsMessage(errs, want) {
				t.Errorf("repeating %q %d times (cap %d) did not report %q (got %v)", tag, first.MaxRepeat+1, first.MaxRepeat, want, errs)
			}
		})
	}
}

// TestAllRegisteredSchemasGroupPathAtStaysInRange exercises GroupPathAt
// (edifact-ls-pcm0's hover group-context lookup) against every
// registered schema's own minimal-conformant message, at every index --
// a broad smoke test against real, deeply-nested trees (unlike the
// small hand-built ones GroupPathAt's own dedicated tests use), asserting
// every reported group number is a real, in-range group of that schema.
func TestAllRegisteredSchemasGroupPathAtStaysInRange(t *testing.T) {
	for _, info := range ListRegisteredSchemas() {
		id := info.ID
		t.Run(fmt.Sprintf("%s_%s_%s_%s", id.Type, id.Version, id.Release, id.Agency), func(t *testing.T) {
			rs := schemaRegistry[id]
			segs := segsFromTags(minimalMandatoryTags(rs.Schema.Nodes))
			totalGroups := countGroups(rs.Schema.Nodes)
			for i := range segs {
				for _, n := range GroupPathAt(rs.Schema, segs, i) {
					if n < 1 || n > totalGroups {
						t.Errorf("segment %d (%q): group number %d out of range [1,%d]", i, segs[i].Tag, n, totalGroups)
					}
				}
			}
		})
	}
}

func countGroups(nodes []SchemaNode) int {
	n := 0
	for _, node := range nodes {
		if node.Group != nil {
			n++
			n += countGroups(node.Group)
		}
	}
	return n
}
