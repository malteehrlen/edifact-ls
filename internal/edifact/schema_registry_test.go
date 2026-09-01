package edifact

import "testing"

func TestMessageIDOfExtractsS009Components(t *testing.T) {
	ic, errs := Parse("UNH+1+IFTMCS:D:21A:UN'")
	if errs.HasErrors() {
		t.Fatalf("unexpected parse errors: %v", errs)
	}
	got := messageIDOf(&ic.Segments[0], ic.Delimiters)
	want := MessageID{Type: "IFTMCS", Version: "D", Release: "21A", Agency: "UN"}
	if got != want {
		t.Fatalf("messageIDOf = %+v, want %+v", got, want)
	}
}

func TestValidateMessageSchemasAppliesRegisteredSchema(t *testing.T) {
	id := MessageID{Type: "TESTMSG", Version: "D", Release: "1A", Agency: "UN"}
	RegisterSchema(id, Schema{Nodes: []SchemaNode{
		{Segment: "BGM", Mandatory: true, MaxRepeat: 1},
	}})
	defer delete(schemaRegistry, id)

	src := "UNB+UNOA:1+S+R+201001:1200+1'UNH+1+TESTMSG:D:1A:UN'UNT+2+1'UNZ+1+1'"
	ic, errs := Parse(src)
	if errs.HasErrors() {
		t.Fatalf("unexpected parse errors: %v", errs)
	}

	got := ValidateMessageSchemas(ic)
	if len(got) != 1 {
		t.Fatalf("got %d diagnostics, want 1: %v", len(got), got)
	}
	if got[0].Severity != SeverityError || !containsMessage(got, "missing mandatory") || !containsMessage(got, `"BGM"`) {
		t.Errorf("diagnostic = %v, want an error about missing mandatory BGM", got[0])
	}
}

func TestValidateMessageSchemasCleanMessagePasses(t *testing.T) {
	id := MessageID{Type: "TESTMSG2", Version: "D", Release: "1A", Agency: "UN"}
	RegisterSchema(id, Schema{Nodes: []SchemaNode{
		{Segment: "BGM", Mandatory: true, MaxRepeat: 1},
	}})
	defer delete(schemaRegistry, id)

	src := "UNB+UNOA:1+S+R+201001:1200+1'UNH+1+TESTMSG2:D:1A:UN'BGM+220'UNT+3+1'UNZ+1+1'"
	ic, errs := Parse(src)
	if errs.HasErrors() {
		t.Fatalf("unexpected parse errors: %v", errs)
	}
	if got := ValidateMessageSchemas(ic); len(got) != 0 {
		t.Fatalf("unexpected diagnostics for a conformant message: %v", got)
	}
}

func TestValidateMessageSchemasUnmatchedVersionProducesInfo(t *testing.T) {
	id := MessageID{Type: "TESTMSG3", Version: "D", Release: "1A", Agency: "UN"}
	RegisterSchema(id, Schema{Nodes: []SchemaNode{
		{Segment: "BGM", Mandatory: true, MaxRepeat: 1},
	}})
	defer delete(schemaRegistry, id)

	// Same type, different release -- no exact match, but not unknown
	// either.
	src := "UNB+UNOA:1+S+R+201001:1200+1'UNH+1+TESTMSG3:D:2B:UN'BGM+220'UNT+3+1'UNZ+1+1'"
	ic, errs := Parse(src)
	if errs.HasErrors() {
		t.Fatalf("unexpected parse errors: %v", errs)
	}

	got := ValidateMessageSchemas(ic)
	if len(got) != 1 {
		t.Fatalf("got %d diagnostics, want 1: %v", len(got), got)
	}
	if got[0].Severity != SeverityInfo {
		t.Errorf("severity = %v, want info", got[0].Severity)
	}
	if !containsMessage(got, `"TESTMSG3"`) || !containsMessage(got, "D:1A:UN") {
		t.Errorf("message = %q, want it to name the registered alternative D:1A:UN", got[0].Message)
	}
}

func TestValidateMessageSchemasUnknownTypeProducesNoDiagnostic(t *testing.T) {
	src := "UNB+UNOA:1+S+R+201001:1200+1'UNH+1+TOTALLYUNKNOWN:D:1A:UN'BGM+220'UNT+3+1'UNZ+1+1'"
	ic, errs := Parse(src)
	if errs.HasErrors() {
		t.Fatalf("unexpected parse errors: %v", errs)
	}
	if got := ValidateMessageSchemas(ic); len(got) != 0 {
		t.Fatalf("unexpected diagnostics for a wholly unregistered message type: %v", got)
	}
}

// TestValidateMessageSchemasPlugAndPlaySecondType registers two unrelated
// message types back to back and validates a message against each,
// proving the registry/dispatch needs no per-type code -- only a
// registration call -- per the plug-and-play acceptance criterion.
func TestValidateMessageSchemasPlugAndPlaySecondType(t *testing.T) {
	orders := MessageID{Type: "TESTORDERS", Version: "D", Release: "96A", Agency: "UN"}
	RegisterSchema(orders, Schema{Nodes: []SchemaNode{
		{Segment: "BGM", Mandatory: true, MaxRepeat: 1},
		{Segment: "DTM", Mandatory: false, MaxRepeat: 5},
	}})
	defer delete(schemaRegistry, orders)

	iftmcsLike := MessageID{Type: "TESTIFTMCS", Version: "D", Release: "21A", Agency: "UN"}
	RegisterSchema(iftmcsLike, Schema{Nodes: []SchemaNode{
		{Segment: "CTA", Mandatory: true, MaxRepeat: 1},
	}})
	defer delete(schemaRegistry, iftmcsLike)

	ordersSrc := "UNB+UNOA:1+S+R+201001:1200+1'UNH+1+TESTORDERS:D:96A:UN'BGM+220'UNT+3+1'UNZ+1+1'"
	ic, errs := Parse(ordersSrc)
	if errs.HasErrors() {
		t.Fatalf("unexpected parse errors: %v", errs)
	}
	if got := ValidateMessageSchemas(ic); len(got) != 0 {
		t.Fatalf("unexpected diagnostics validating against TESTORDERS: %v", got)
	}

	iftmcsSrc := "UNB+UNOA:1+S+R+201001:1200+1'UNH+1+TESTIFTMCS:D:21A:UN'UNT+2+1'UNZ+1+1'"
	ic, errs = Parse(iftmcsSrc)
	if errs.HasErrors() {
		t.Fatalf("unexpected parse errors: %v", errs)
	}
	got := ValidateMessageSchemas(ic)
	if len(got) != 1 || !containsMessage(got, `"CTA"`) {
		t.Fatalf("validating against TESTIFTMCS: got %v, want one missing-mandatory-CTA error", got)
	}
}
