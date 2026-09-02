package edifact

import (
	"fmt"
	"sort"
	"strings"
)

// MessageID identifies a UN/EDIFACT message type by the four components
// UNH's S009 composite carries (data elements 0065/0052/0054/0051): type,
// version, release, and controlling agency -- e.g. {"IFTMCS", "D", "21A",
// "UN"}. A Schema is registered against one exact MessageID, since a
// message type's branching diagram is release-specific.
type MessageID struct {
	Type    string
	Version string
	Release string
	Agency  string
}

func (id MessageID) versionString() string {
	return fmt.Sprintf("%s:%s:%s", id.Version, id.Release, id.Agency)
}

// messageIDOf extracts the message identifier from a UNH segment's S009
// composite (element index 1).
func messageIDOf(unh *Segment, d Delimiters) MessageID {
	return MessageID{
		Type:    unh.ComponentN(1, 0, d),
		Version: unh.ComponentN(1, 1, d),
		Release: unh.ComponentN(1, 2, d),
		Agency:  unh.ComponentN(1, 3, d),
	}
}

// registeredSchema pairs a Schema with the source it was transcribed
// from, so provenance travels as queryable data (for ListRegisteredSchemas
// and everything built on it) instead of living only in a Go doc comment.
type registeredSchema struct {
	Schema Schema
	Source string
}

// schemaRegistry holds every Schema registered via RegisterSchema, keyed
// by the exact MessageID it applies to.
var schemaRegistry = map[MessageID]registeredSchema{}

// RegisterSchema registers schema as the structural specification for
// messages self-reporting id in their UNH, citing source (the canonical
// spec URL it was transcribed from). Intended to be called once
// (typically from a package-level init in a file dedicated to one message
// type's schema data, e.g. an IFTMCS-D-21A source file) -- adding a new
// message type should never require touching the registry, the validator,
// or diagnostics wiring, only a new schema-data file that calls this.
func RegisterSchema(id MessageID, schema Schema, source string) {
	schemaRegistry[id] = registeredSchema{Schema: schema, Source: source}
}

// SchemaInfo pairs a registered MessageID with the source it was
// transcribed from, for listing/documentation purposes (the `edifact-ls
// schemas` CLI command and the generated SUPPORTED_MESSAGES.md both build
// on this).
type SchemaInfo struct {
	ID     MessageID
	Source string
}

// ListRegisteredSchemas returns every registered schema's identity and
// source, sorted by Type, then Version, then Release, then Agency, for
// deterministic output.
func ListRegisteredSchemas() []SchemaInfo {
	infos := make([]SchemaInfo, 0, len(schemaRegistry))
	for id, rs := range schemaRegistry {
		infos = append(infos, SchemaInfo{ID: id, Source: rs.Source})
	}
	sort.Slice(infos, func(i, j int) bool {
		a, b := infos[i].ID, infos[j].ID
		if a.Type != b.Type {
			return a.Type < b.Type
		}
		if a.Version != b.Version {
			return a.Version < b.Version
		}
		if a.Release != b.Release {
			return a.Release < b.Release
		}
		return a.Agency < b.Agency
	})
	return infos
}

// registeredVersionsOf lists the version:release:agency of every schema
// registered for msgType, sorted for deterministic output.
func registeredVersionsOf(msgType string) []string {
	var versions []string
	for id := range schemaRegistry {
		if id.Type == msgType {
			versions = append(versions, id.versionString())
		}
	}
	sort.Strings(versions)
	return versions
}

// ValidateMessageSchemas checks each UNH..UNT message in ic against a
// registered Schema for its self-reported MessageID, if one exists. Each
// resulting violation is decorated with which message type/version and
// which UNH it was validated against (e.g. "...as specified by message
// type IFTMCS D:21A:UN on line 2"), so that's clear even in an interchange
// with several messages, without a separate diagnostic just to announce
// recognition. A message whose type is registered under one or more
// different versions/releases gets an informational diagnostic naming
// what's available instead of being silently skipped or validated against
// the wrong release; a type with no registered schema at all produces no
// diagnostic. A message missing its UNT is skipped here -- ValidateEnvelopes
// already reports that, and without a UNT the body's extent is ambiguous.
func ValidateMessageSchemas(ic *Interchange) ErrorList {
	var errs ErrorList
	d := ic.Delimiters

	var unh *Segment
	bodyStart := -1
	for i := range ic.Segments {
		seg := &ic.Segments[i]
		if seg.Tag == "UNH" {
			unh = seg
			bodyStart = i + 1
			continue
		}
		if seg.Tag == "UNT" && unh != nil {
			body := ic.Segments[bodyStart:i]
			errs = append(errs, validateOneMessage(unh, body, seg.Pos, d)...)
			unh = nil
			bodyStart = -1
		}
	}
	return errs
}

func validateOneMessage(unh *Segment, body []Segment, untPos Position, d Delimiters) ErrorList {
	var errs ErrorList
	id := messageIDOf(unh, d)

	if rs, ok := schemaRegistry[id]; ok {
		violations := ValidateSchema(rs.Schema, body, untPos)
		for _, v := range violations {
			v.Message = fmt.Sprintf("%s, as specified by message type %s %s on line %d", v.Message, id.Type, id.versionString(), unh.Pos.Line)
		}
		errs = append(errs, violations...)
		return errs
	}

	if alts := registeredVersionsOf(id.Type); len(alts) > 0 {
		errs.Add(unh.Pos, SeverityInfo,
			"no message specification registered for %q version %s release %s (agency %s); structural validation skipped -- registered for %q: %s",
			id.Type, id.Version, id.Release, id.Agency, id.Type, strings.Join(alts, ", "))
	}
	return errs
}
