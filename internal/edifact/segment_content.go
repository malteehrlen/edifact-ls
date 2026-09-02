package edifact

// SegmentElementSchema describes one segment tag's data element structure
// -- intrinsic to the tag, the same regardless of which message type uses
// it (confirmed against the real UN/EDIFACT Segment Directory: a
// segment's definition page lists one structure, then the full list of
// message types that share it). Independent of Schema/SchemaNode, which
// describe segment/group structure *within* a message instead.
type SegmentElementSchema struct {
	Elements []ElementSchema
}

// ElementSchema describes one data element position within a segment.
// Components holds one entry for a simple data element, or several for a
// composite -- each with its own Mandatory flag, since a composite's
// components are independently mandatory/conditional.
type ElementSchema struct {
	Name       string
	Mandatory  bool // whether the element itself may be omitted entirely
	Components []ComponentSchema
}

// ComponentSchema describes one component of a data element (the element
// itself, for a simple data element with exactly one component).
type ComponentSchema struct {
	Name      string
	Mandatory bool

	// CodeList, if non-empty, is the UN Trade Data Element Directory
	// data-element number of the code list this component's actual
	// value is drawn from (e.g. "1225" for BGM's message function
	// code) -- see codelist.go. Empty for a component whose value isn't
	// coded, or one that is but whose code list hasn't been sourced yet
	// (see edifact-ls-6xaz's scope note on this being registered
	// incrementally, not exhaustively).
	CodeList string
}

// ValidateSegmentElements checks seg's actual elements/components against
// schema, reporting a missing mandatory element (the element itself
// absent) or a missing mandatory component (the element present but a
// required component within it missing or empty).
func ValidateSegmentElements(schema SegmentElementSchema, seg Segment) ErrorList {
	var errs ErrorList

	for i, es := range schema.Elements {
		el := seg.Element0(i)
		if el == nil {
			if es.Mandatory {
				errs.Add(seg.Pos, SeverityError, "segment %q is missing its mandatory element %d (%s)", seg.Tag, i+1, es.Name)
			}
			continue
		}

		for ci, cs := range es.Components {
			if !cs.Mandatory {
				continue
			}
			if ci >= len(el.Components) || el.Components[ci].Raw == "" {
				if len(es.Components) == 1 {
					errs.Add(el.Pos, SeverityError, "segment %q is missing its mandatory element %d (%s)", seg.Tag, i+1, es.Name)
				} else {
					errs.Add(el.Pos, SeverityError, "segment %q element %d (%s) is missing its mandatory component %d (%s)", seg.Tag, i+1, es.Name, ci+1, cs.Name)
				}
			}
		}
	}

	return errs
}

// segmentElementSchemas holds every SegmentElementSchema registered via
// RegisterSegmentElementSchema, keyed by segment tag. Unlike
// schemaRegistry (keyed by a message's full MessageID), this is keyed by
// the plain tag: a segment's element structure is intrinsic to the tag,
// not per-message-type -- see edifact-ls-9ger's design note.
var segmentElementSchemas = map[string]SegmentElementSchema{}

// RegisterSegmentElementSchema registers schema as tag's element
// structure. Intended to be called once, typically from a package-level
// init in a file dedicated to one segment's real data (see
// segment_elements_data.go) -- adding a new segment's content validation
// should never require touching the registry or ValidateSegmentContent,
// only a new registration call.
func RegisterSegmentElementSchema(tag string, schema SegmentElementSchema) {
	segmentElementSchemas[tag] = schema
}

// SegmentElementSchemaFor returns the registered SegmentElementSchema for
// tag, if any -- for callers outside this package that need a segment's
// element/component structure directly, e.g. hover's coded-value lookup
// (see edifact-ls-6xaz).
func SegmentElementSchemaFor(tag string) (SegmentElementSchema, bool) {
	schema, ok := segmentElementSchemas[tag]
	return schema, ok
}

// ValidateSegmentContent checks every segment in ic against a registered
// SegmentElementSchema for its tag, if one exists. Segments whose tag has
// no registered schema produce no diagnostic.
func ValidateSegmentContent(ic *Interchange) ErrorList {
	var errs ErrorList
	for _, seg := range ic.Segments {
		if schema, ok := segmentElementSchemas[seg.Tag]; ok {
			errs = append(errs, ValidateSegmentElements(schema, seg)...)
		}
	}
	return errs
}
