package edifact

import "strings"

// RenderSupportedMessagesDoc renders every registered structural schema
// as a Markdown table -- the single generated source for
// docs/SUPPORTED_MESSAGES.md (produced via `make docs` / tools/gendocs).
// TestSupportedMessagesDocIsUpToDate compares this function's output
// against that checked-in file, so the two can never silently drift: the
// registry (via ListRegisteredSchemas) is always the source of truth.
func RenderSupportedMessagesDoc() string {
	var b strings.Builder
	b.WriteString("# Supported message specifications\n\n")
	b.WriteString("Generated from the schemas registered in `internal/edifact` -- see `make docs`.\n")
	b.WriteString("Do not hand-edit; regenerate instead.\n\n")
	b.WriteString("Structural (segment/group presence, order, cardinality) validation is\n")
	b.WriteString("available for exactly the message identities listed below -- matched on the\n")
	b.WriteString("full (type, version, release, agency) tuple a message declares in its own\n")
	b.WriteString("UNH. A message declaring a recognized type under a different version,\n")
	b.WriteString("release, or agency than any row here gets an informational diagnostic\n")
	b.WriteString("naming what IS registered, not silence and not a false match.\n\n")
	b.WriteString("| Type | Version | Release | Agency | Source |\n")
	b.WriteString("| --- | --- | --- | --- | --- |\n")

	for _, info := range ListRegisteredSchemas() {
		b.WriteString("| ")
		b.WriteString(info.ID.Type)
		b.WriteString(" | ")
		b.WriteString(info.ID.Version)
		b.WriteString(" | ")
		b.WriteString(info.ID.Release)
		b.WriteString(" | ")
		b.WriteString(info.ID.Agency)
		b.WriteString(" | ")
		b.WriteString(info.Source)
		b.WriteString(" |\n")
	}

	return b.String()
}
