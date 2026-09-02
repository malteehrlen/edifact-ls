package edifact

// CodedValue is one entry in a UN/EDIFACT code list (a "UNCL", identified
// by its UN Trade Data Element Directory data-element number, e.g. "1225"
// for message function code) -- the meaning of one specific coded value a
// data element can take, e.g. code "9" in list "1225" means "Original".
type CodedValue struct {
	Name        string
	Description string
}

// codeLists holds every code list registered via RegisterCodeList, keyed
// by its data-element number, then by the code itself.
var codeLists = map[string]map[string]CodedValue{}

// RegisterCodeList registers values as the code list identified by id (a
// UN Trade Data Element Directory data-element number). Intended to be
// called once, from a package-level init in a file dedicated to one real
// code list's data (see codelist_1225.go) -- adding a new code list
// should never require touching this file, only a new data file that
// calls this.
func RegisterCodeList(id string, values map[string]CodedValue) {
	codeLists[id] = values
}

// LookupCode returns the CodedValue for code within code list id, if both
// the list and the code within it are registered.
func LookupCode(id, code string) (CodedValue, bool) {
	list, ok := codeLists[id]
	if !ok {
		return CodedValue{}, false
	}
	cv, ok := list[code]
	return cv, ok
}
