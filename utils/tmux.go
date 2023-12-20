package hyprjmp

import (
	"fmt"
	"os/exec"
	"strings"
)

func GetTmuxSessions() ([]string, error) {
	sessions := []string{}
	sessionsStr, err := exec.Command("tmux", "list-sessions").Output()
	if err != nil {
		return sessions, err
	}
	for _, sessionInfo := range strings.Split(string(sessionsStr), "\n") {
		session := strings.Split(sessionInfo, ":")[0]
		sessions = append(sessions, session)
	}
	return sessions, err
}

type Window struct {
	Id  int
	Name string
}

func GetWindowsSession(session string) ([]Window, error) {
	windows := []Window{}
	windowsStr, err := exec.Command("tmux", "list-windows", "-t", session).Output()
	if err != nil {
		return windows, err
	}
	for _, windowInfo := range strings.Split(string(windowsStr), "\n") {
		if windowInfo == "" {
			continue
		}
		window := parseWindow(windowInfo)
		windows = append(windows, window)
	}
	return windows, err
}

func parseWindow(windowInfo string) Window {
	window := Window{}
	fmt.Sscanf(windowInfo, "%d: %s", &window.Id, &window.Name)
	return window
}

type Session struct {
	Name    string
	Windows []Window
}

type TmuxInfo struct {
	Sessions []Session
}

func GetTmuxInfo() (TmuxInfo, error) {
	tmuxInfo := TmuxInfo{}
	sessions, err := GetTmuxSessions()
	if err != nil {
		return tmuxInfo, err
	}
	for _, session := range sessions {
		windows, err := GetWindowsSession(session)
		if err != nil {
			return tmuxInfo, err
		}
		tmuxInfo.Sessions = append(tmuxInfo.Sessions, Session{session, windows})
	}
	return tmuxInfo, err
}
