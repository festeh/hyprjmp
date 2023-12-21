package hyprjmp

import (
	"fmt"
	"log"
	"os/exec"

	// "os/exec"
	"strings"

	"go.i3wm.org/i3/v4"
)

func childCmp(node *i3.Node, name string) *i3.Node {
	return node.FindChild(func(n *i3.Node) bool {
		return n.Type == i3.WorkspaceNode && n.Name == name
	})
}

func makeCommand(cmd string) {
	res, err := i3.RunCommand(cmd)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(res)
}

func getI3WS() error {
	tree, err := i3.GetTree()
	if err != nil {
		return err
	}
	ws, err := i3.GetWorkspaces()
	if err != nil {
		return err
	}
	for _, w := range ws {
		if node := childCmp(tree.Root, w.Name); node != nil {
			fmt.Printf("%s %d %d \n", w.Name, w.ID, w.Num)
			node.FindChild(func(n *i3.Node) bool {
				fmt.Printf("  %s %d \n", n.Name, n.ID)
				// Print window properties
				fmt.Printf("		%s %s %s %s \n", n.WindowProperties.Class, n.WindowProperties.Instance, n.WindowProperties.Role, n.WindowProperties.Title)
				found := strings.Contains(n.Name, "tmux")
				return found
			})
		}
	}
	return nil
}

func findFocusedWorkspace() string {
	ws, err := i3.GetWorkspaces()
	if err != nil {
		log.Fatal(err)
		return ""
	}
	for _, w := range ws {
		if w.Focused {
			log.Printf("Focused workspace: %s", w.Name)
			return w.Name
		}
	}
	log.Fatal("No focused workspace found")
	return ""
}

func MakeTmuxTarget(session string, window string) string {
	windowId := strings.Split(window, ":")[0]
	return session + ":" + windowId
}

func OpenTmuxSessionWindow(workspace string, target string) {
	if !strings.Contains(findFocusedWorkspace(), workspace) {
		makeCommand(fmt.Sprintf("workspace number %s", workspace))
	}
	exec.Command("tmux", "switch-client", "-t", target).Run()
}
