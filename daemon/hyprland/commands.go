package hyprland

import (
	"fmt"
	"strings"
)

func parseWorkspaceOutput(output string, id string) (string, error) {
	found := false
	// iterate over lines
	for _, line := range strings.Split(output, "\n") {
		if strings.Contains(line, fmt.Sprintf("workspace ID %s", id)) {
			found = true
		}
		if found {
			// check if string has lastwindow:XXXX in it, extract windowd address
			if strings.Contains(line, "lastwindow:") {
				info := strings.Split(line, "lastwindow:")[1]
				trimmed := strings.TrimSpace(info)
				return trimmed, nil
			}
		}
	}

	return "", fmt.Errorf("Not found")
}

// func ClearWorkspace(id string) {
// res := exec.Command("hyprctl", "workspaces").Run()
// }
