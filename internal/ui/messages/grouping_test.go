package messages

import (
	"strings"
	"testing"

	"github.com/charmbracelet/x/ansi"
)

func TestGroupMessagesCollapsesSameAuthorRun(t *testing.T) {
	msgs := []MessageItem{
		{TS: "1700000000.000100", UserID: "U1", UserName: "alice", Text: "one", Timestamp: "10:00"},
		{TS: "1700000060.000100", UserID: "U1", UserName: "alice", Text: "two", Timestamp: "10:01"},
		{TS: "1700000120.000100", UserID: "U2", UserName: "bob", Text: "three", Timestamp: "10:02"},
		{TS: "1700000900.000100", UserID: "U2", UserName: "bob", Text: "four", Timestamp: "10:15"}, // >5 min: new header
	}
	m := New(msgs, "general")
	m.SetGroupMessages(true)
	got := ansi.Strip(m.View(30, 60))
	if n := strings.Count(got, "alice"); n != 1 {
		t.Errorf("alice header count = %d, want 1 (second message should be a continuation)\n%s", n, got)
	}
	if n := strings.Count(got, "bob"); n != 2 {
		t.Errorf("bob header count = %d, want 2 (15-minute gap breaks the run)\n%s", n, got)
	}
	for _, body := range []string{"one", "two", "three", "four"} {
		if !strings.Contains(got, body) {
			t.Errorf("body %q missing from view\n%s", body, got)
		}
	}
	m.SetGroupMessages(false)
	if n := strings.Count(ansi.Strip(m.View(30, 60)), "alice"); n != 2 {
		t.Errorf("grouping off: alice header count = %d, want 2", n)
	}
}
