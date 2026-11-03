package help

import "testing"

func TestSheetHeadersHideWhileFiltering(t *testing.T) {
	m := New()
	m.SetEntries(Sheet())
	m.Open()
	headers := 0
	for _, e := range m.VisibleEntries() {
		if e.Header {
			headers++
		}
	}
	if headers < 10 {
		t.Fatalf("unfiltered sheet shows %d section headers, want at least 10", headers)
	}
	m.HandleKey("/")
	for _, r := range "thread" {
		m.HandleKey(string(r))
	}
	vis := m.VisibleEntries()
	if len(vis) == 0 {
		t.Fatal("filter 'thread' matched nothing")
	}
	for _, e := range vis {
		if e.Header {
			t.Fatalf("header %q shown while filtering", e.Section)
		}
	}
	// Section names match too: an entry under "Threads" whose text
	// never says "thread" still comes up.
	found := false
	for _, e := range vis {
		if e.Key == "S" && e.Section == "Threads" {
			found = true
		}
	}
	if !found {
		t.Error("filtering by section name did not surface the 'S' save-thread row")
	}
}
