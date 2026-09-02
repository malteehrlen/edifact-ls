package edifact

import "testing"

func TestLookupCodeKnownList(t *testing.T) {
	cv, ok := LookupCode("1225", "9")
	if !ok {
		t.Fatal("LookupCode(1225, 9) not found, want the real 'Original' entry")
	}
	if cv.Name != "Original" {
		t.Errorf("Name = %q, want %q", cv.Name, "Original")
	}
	if cv.Description == "" {
		t.Error("Description is empty, want the real UN/EDIFACT description")
	}
}

func TestLookupCodeUnknownCodeInKnownList(t *testing.T) {
	if _, ok := LookupCode("1225", "999999"); ok {
		t.Error("LookupCode(1225, 999999) found something, want not-ok for a code that doesn't exist")
	}
}

func TestLookupCodeUnknownList(t *testing.T) {
	if _, ok := LookupCode("9999999", "1"); ok {
		t.Error("LookupCode for an unregistered list found something, want not-ok")
	}
}

func TestCodeList1225RegisteredWithRealEntryCount(t *testing.T) {
	// 69 real entries per the source's own "Code Values:" section -- a
	// stand-in for "the extraction didn't silently drop or corrupt
	// data" without re-fetching the source on every test run.
	if got := len(codeLists["1225"]); got != 69 {
		t.Errorf("len(codeLists[1225]) = %d, want 69", got)
	}
}

func TestCodeList3139RegisteredWithRealEntryCount(t *testing.T) {
	if got := len(codeLists["3139"]); got != 103 {
		t.Errorf("len(codeLists[3139]) = %d, want 103", got)
	}
}

func TestBGMMessageFunctionCodeIsCoded(t *testing.T) {
	schema := segmentElementSchemas["BGM"]
	cs := schema.Elements[2].Components[0] // "Message function code"
	if cs.CodeList != "1225" {
		t.Fatalf("BGM message function code's CodeList = %q, want %q", cs.CodeList, "1225")
	}
}

func TestCTAContactFunctionCodeIsCoded(t *testing.T) {
	schema := segmentElementSchemas["CTA"]
	cs := schema.Elements[0].Components[0] // "Contact function code"
	if cs.CodeList != "3139" {
		t.Fatalf("CTA contact function code's CodeList = %q, want %q", cs.CodeList, "3139")
	}
}
