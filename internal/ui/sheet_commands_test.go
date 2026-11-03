package ui

import (
	"strings"
	"testing"

	"github.com/gammons/slk/internal/ui/help"
)

// Every ":" command must be findable in the cheat sheet, like
// TestHelpListsEveryCommand pins for the keymap.
func TestSheetListsEveryCommand(t *testing.T) {
	aliases := map[string]bool{"on": true}
	var keys []string
	for _, e := range help.Sheet() {
		keys = append(keys, e.Key)
	}
	listed := strings.Join(keys, " ")
	for name := range commands {
		if aliases[name] {
			continue
		}
		if !strings.Contains(listed, ":"+name) {
			t.Errorf(":%s has no row in the cheat sheet", name)
		}
	}
}
