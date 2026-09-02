// Package edifact lexes and parses UN/EDIFACT interchanges into a
// positioned AST, and validates interchange/message envelope structure.
package edifact

import "fmt"

// Position is a location in EDIFACT source text.
type Position struct {
	Offset int // 0-based byte offset
	Line   int // 1-based
	Column int // 1-based byte column
}

// Delimiters is the set of structural characters an interchange uses,
// either the documented defaults or as redefined by a UNA service string
// advice segment.
type Delimiters struct {
	Component  byte // default ':'
	Element    byte // default '+'
	Decimal    byte // default '.'
	Release    byte // default '?'
	Reserved   byte // default ' ' (unused structurally, reserved by the spec)
	Terminator byte // default '\''
}

// DefaultDelimiters returns the delimiter set used when no UNA service
// string advice segment is present.
func DefaultDelimiters() Delimiters {
	return Delimiters{
		Component:  ':',
		Element:    '+',
		Decimal:    '.',
		Release:    '?',
		Reserved:   ' ',
		Terminator: '\'',
	}
}

// Severity classifies a diagnostic produced while parsing, validating, or
// linting.
type Severity int

const (
	SeverityError Severity = iota
	SeverityWarning
	SeverityInfo
)

func (s Severity) String() string {
	switch s {
	case SeverityWarning:
		return "warning"
	case SeverityInfo:
		return "info"
	default:
		return "error"
	}
}

// Validate runs the full validation pipeline this package offers -- Parse,
// then ValidateEnvelopes, Lint, ValidateMessageSchemas, and
// ValidateSegmentContent layered on top of its result -- and returns the
// parsed Interchange alongside every diagnostic collected along the way,
// in that order. This is the single place that sequence is assembled,
// shared by the LSP server's publishDiagnostics path and the
// `edifact-ls check` CLI command, so a future check added to the
// pipeline can't land in one and be forgotten in the other.
func Validate(src string) (*Interchange, ErrorList) {
	ic, errs := Parse(src)
	errs = append(errs, ValidateEnvelopes(ic)...)
	errs = append(errs, Lint(ic, src)...)
	errs = append(errs, ValidateMessageSchemas(ic)...)
	errs = append(errs, ValidateSegmentContent(ic)...)
	return ic, errs
}

// Error is a structured, positioned problem found while lexing, parsing, or
// validating an interchange.
type Error struct {
	Pos      Position
	Severity Severity
	Message  string

	// Code is a stable, machine-readable identifier for this diagnostic's
	// kind (e.g. "redundant-una", "envelope-count-mismatch"), used by
	// clients to filter/match on without parsing Message's prose. Empty
	// when there's nothing more stable to offer than the message text.
	Code string

	// Fix is the mechanical edit that resolves this diagnostic, if one
	// can be derived without guessing at user intent. Nil when the
	// diagnostic isn't mechanically fixable.
	Fix *Fix
}

// Fix is a single, exact text replacement that resolves an Error.
type Fix struct {
	Title   string   // human-readable, e.g. "Remove redundant UNA service string advice"
	Pos     Position // start of the exact span to replace
	OldText string   // sanity-checked against current buffer text before applying
	NewText string   // "" for a deletion
}

func (e *Error) Error() string {
	return fmt.Sprintf("%d:%d: %s: %s", e.Pos.Line, e.Pos.Column, e.Severity, e.Message)
}

// ErrorList accumulates Errors in the order they were found.
type ErrorList []*Error

// Add appends a new formatted error with no Code or Fix.
func (el *ErrorList) Add(pos Position, severity Severity, format string, args ...any) {
	*el = append(*el, &Error{Pos: pos, Severity: severity, Message: fmt.Sprintf(format, args...)})
}

// AddFixable appends a new formatted error carrying a stable Code and,
// when the caller can derive one, a mechanical Fix.
func (el *ErrorList) AddFixable(pos Position, severity Severity, code string, fix *Fix, format string, args ...any) {
	*el = append(*el, &Error{Pos: pos, Severity: severity, Message: fmt.Sprintf(format, args...), Code: code, Fix: fix})
}

// HasErrors reports whether the list contains any SeverityError entries.
func (el ErrorList) HasErrors() bool {
	for _, e := range el {
		if e.Severity == SeverityError {
			return true
		}
	}
	return false
}
