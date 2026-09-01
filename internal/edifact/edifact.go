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

// Error is a structured, positioned problem found while lexing, parsing, or
// validating an interchange.
type Error struct {
	Pos      Position
	Severity Severity
	Message  string
}

func (e *Error) Error() string {
	return fmt.Sprintf("%d:%d: %s: %s", e.Pos.Line, e.Pos.Column, e.Severity, e.Message)
}

// ErrorList accumulates Errors in the order they were found.
type ErrorList []*Error

// Add appends a new formatted error.
func (el *ErrorList) Add(pos Position, severity Severity, format string, args ...any) {
	*el = append(*el, &Error{Pos: pos, Severity: severity, Message: fmt.Sprintf(format, args...)})
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
