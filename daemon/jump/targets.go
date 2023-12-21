package hyprjmp

import (
	"encoding/json"
	"log"
)

type Target interface {
}

type TmuxTarget struct {
	Session string
	Window  string
}

type FirefoxTarget struct {
	Tab string
}

type I3Target struct {
	Workspace string
}

func ParseJumpRequest(req string) (Target, error) {
	var result map[string]interface{}
	json.Unmarshal([]byte(req), &result)

	switch result["type"] {
	case "tmux":
		var tmuxTarget TmuxTarget
		if err := json.Unmarshal([]byte(req), &tmuxTarget); err != nil {
			log.Fatal(err)
			return nil, err
		}
		return tmuxTarget, nil
	}
	return nil, nil
}
