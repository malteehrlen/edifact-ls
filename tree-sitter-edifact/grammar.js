// EDIFACT tree-sitter grammar.
//
// Scope: default delimiters only (component ':', element '+', release '?',
// terminator '\''), matching the vast majority of real-world interchanges.
// A `UNA` service string advice segment is recognized and highlighted as
// its own node, but -- unlike the LSP server's own Go parser -- this
// grammar does NOT honor a custom delimiter set it might define: tree-
// sitter grammars are static (compiled ahead of time), so adapting lexing
// to delimiters read from the file's own content would require an external
// C scanner. That's out of scope here; a file using a non-default UNA will
// highlight incorrectly (documented limitation, tracked as follow-up work
// if it turns out to matter in practice).
module.exports = grammar({
  name: 'edifact',

  // Segments are commonly placed one-per-line purely for human
  // readability; EDIFACT itself has no concept of significant newlines
  // (mirrors internal/edifact/lexer.go's skipInterSegmentWhitespace).
  extras: $ => [/\r?\n/],

  rules: {
    source_file: $ => seq(
      optional($.una_advice),
      repeat($.segment),
    ),

    // Fixed 9-byte layout: "UNA" + 6 delimiter-definition characters, with
    // no terminator of its own.
    una_advice: $ => /UNA.{6}/,

    // Each '+' introduces an element; a wholly-empty element (nothing at
    // all between two '+'s, or between '+' and the terminator) leaves no
    // characters to build a node from, so it's simply absent from the tree
    // -- there's nothing to highlight there either way. The tag may also
    // carry ':'-separated control-number components directly (no leading
    // '+'), e.g. "GDS:1+..." -- "explicit representation" of a repeating
    // segment, per section 9.5.1 of
    // https://unece.org/DAM/trade/untdid/texts/d423.htm. Rare in modern
    // usage but part of the formal syntax (mirrors
    // internal/edifact/parser.go's consumeTagControlNumbers).
    segment: $ => seq(
      field('tag', $.segment_tag),
      repeat(seq(':', field('control_number', $.data))),
      repeat(seq('+', optional(field('element', $.element)))),
      $.terminator,
    ),

    segment_tag: $ => /[A-Z]{3}/,

    // An element is one or more ':'-separated components. Tree-sitter
    // disallows a rule that can match the empty string, so -- unlike a
    // wholly-empty element (handled above) -- an element that's merely
    // missing its first component (e.g. ":B") still needs to consume at
    // least the ':' character; expressed as two alternatives so every
    // derivation is guaranteed non-empty. Note this means an empty
    // component produces no node at all (not a zero-width one), so a
    // component count derived from this tree undercounts empty slots -- a
    // non-issue for highlighting (nothing to color in an empty component
    // anyway), but not a source of truth for semantic analysis; that's
    // internal/edifact's job.
    element: $ => choice(
      seq($.component, repeat(seq(':', optional($.component)))),
      seq(':', optional($.component), repeat(seq(':', optional($.component)))),
    ),

    component: $ => $.data,

    // A run of characters up to (but not including) the next unescaped
    // delimiter/terminator/newline, where '?' followed by any character
    // escapes that character (mirrors the release-character handling in
    // internal/edifact/lexer.go).
    data: $ => /([^:+'?\r\n]|\?.)+/,

    terminator: $ => "'",
  },
});
