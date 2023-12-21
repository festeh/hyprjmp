package hyprjmp

import (
	"log"
	"testing"
)

func TestGetTmuxInfo(t *testing.T) {
	info, err := GetTmuxSessions()
	if err != nil {
		t.Error(err)
	}
	log.Println(info)
	if len(info) == 0 {
		t.Error("Expected at least one tmux session")
	}
}

func TestParseWindow(t *testing.T) {
	windowInfos := []string{"0: zsh* (1 panes) [80x23] [layout 1a,80x23,0,0,0]",
		"2: zsh (1 panes) [96x24] [layout b5a6,96x24,0,0,9] @9"}
	expectedWindows := []Window{{0, "zsh*"}, {2, "zsh"}}
	for i, windowInfo := range windowInfos {
		if window := parseWindow(windowInfo); window != expectedWindows[i] {
			t.Errorf("Expected %v, got %v", expectedWindows[i], window)
		}
	}

}
