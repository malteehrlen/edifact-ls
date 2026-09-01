; Highlight query for the EDIFACT grammar. Uses standard/conventional
; capture names so nvim's default theme renders something sensible with no
; extra user configuration.

(segment_tag) @keyword

; Service segments (tags starting with "UN", reserved for that purpose per
; section 6.2 of https://unece.org/DAM/trade/untdid/texts/d423.htm -- e.g.
; UNB/UNZ/UNG/UNE/UNH/UNT/UNS) are visually distinguished from ordinary
; user data segment tags. This is a query-level distinction (text-pattern
; matching via #match?), not a grammar-level one -- the grammar itself has
; no notion of which tags are reserved for service use, and doesn't need
; one just for highlighting purposes.
((segment_tag) @keyword.directive
  (#match? @keyword.directive "^UN"))

(una_advice) @keyword.directive

(data) @string

"+" @punctuation.delimiter
":" @punctuation.delimiter
(terminator) @punctuation.delimiter
