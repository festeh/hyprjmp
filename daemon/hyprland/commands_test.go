package hyprland

import (
	"testing"
)

func TestParseWorkspaceOutput(t *testing.T) {
	input := `
	workspace ID 3 (3) on monitor DP-3:
	monitorID: 1
	windows: 1
	hasfullscreen: 0
	lastwindow: 0x55d447f4a820
	lastwindowtitle: ~

workspace ID 4 (4) on monitor DP-3:
	monitorID: 1
	windows: 1
	hasfullscreen: 0
	lastwindow: 0x55d447f54a30
	lastwindowtitle: hyprctl workspaces

workspace ID 6 (6) on monitor DP-3:
	monitorID: 1
	windows: 1
	hasfullscreen: 0
	lastwindow: 0x55d447f37270
	lastwindowtitle: Mozilla Firefox
	`
	expected := "0x55d447f37270"
	output, err := parseWorkspaceOutput(input, "6")
	if err != nil {
		t.Error(err)
	}
	if output != expected {
		t.Errorf("Expected %s, got %s", expected, output)
	}
}
