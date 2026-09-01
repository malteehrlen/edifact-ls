package edifact

import "testing"

func TestSegmentDescriptionKnownTags(t *testing.T) {
	cases := []struct {
		tag      string
		wantName string
	}{
		{"UNH", "Message header"},
		{"BGM", "Beginning of message"},
		{"CTA", "Contact information"},
	}
	for _, c := range cases {
		info, ok := SegmentDescription(c.tag)
		if !ok {
			t.Errorf("SegmentDescription(%q) not found", c.tag)
			continue
		}
		if info.Name != c.wantName {
			t.Errorf("SegmentDescription(%q).Name = %q, want %q", c.tag, info.Name, c.wantName)
		}
		if info.Description == "" {
			t.Errorf("SegmentDescription(%q).Description is empty", c.tag)
		}
	}
}

func TestSegmentDescriptionCoversKnownServiceSegments(t *testing.T) {
	for tag := range knownServiceSegmentTags {
		if _, ok := SegmentDescription(tag); !ok {
			t.Errorf("service segment %q has no description", tag)
		}
	}
}

func TestSegmentDescriptionUnknownTag(t *testing.T) {
	if _, ok := SegmentDescription("ZZZ"); ok {
		t.Error("SegmentDescription(\"ZZZ\") = ok, want not found")
	}
}
