package main

import (
	"encoding/json"
	"fmt"
	hyprjmp "hyprjmp/daemon/jump"
	utils "hyprjmp/daemon/utils"
	"io"
	"log"
	"net/http"

	// "net/http/httputil"
	"sync"
	"time"
)

var (
	tmuxInfo  utils.TmuxInfo
	cacheLock sync.RWMutex
)

func tmuxWorker() {
	for {
		info, err := utils.GetTmuxInfo()
		if err != nil {
			log.Fatal(err)
		} else {
			cacheLock.Lock()
			tmuxInfo = info
			cacheLock.Unlock()
		}
		time.Sleep(1 * time.Second)
	}
}
func setupCORSHeaders(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Add("Access-Control-Allow-Headers", "Accept, Content-Type, Access-Control-Allow-Headers, Authorization, X-Requested-With")
	w.Header().Add("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
}

func main() {
	go tmuxWorker()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		setupCORSHeaders(w)
		cacheLock.RLock()
		// marshal tmuxInfo to json
		bytes, err := json.Marshal(tmuxInfo)
		if err != nil {
			log.Fatal(err)
		}
		cacheLock.RUnlock()
		w.Write(bytes)
	})

	http.HandleFunc("/jump", func(w http.ResponseWriter, r *http.Request) {
		setupCORSHeaders(w)
		defer r.Body.Close()
		buf, err := io.ReadAll(r.Body)
		if err != nil {
			log.Printf("Error reading request body %s", err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		target, err := hyprjmp.ParseJumpRequest(string(buf))
		if err != nil {
			log.Printf("Error parsing jump request: %s", err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		switch target.(type) {
		case hyprjmp.TmuxTarget:
			tmuxTarget := target.(hyprjmp.TmuxTarget)
			log.Printf("Jumping to tmux session: %s window: %s", tmuxTarget.Session, tmuxTarget.Window)
			tmuxTargetStr := utils.MakeTmuxTarget(tmuxTarget.Session, tmuxTarget.Window)
			log.Printf("Jumping to tmux target: %s", tmuxTargetStr)
			utils.OpenTmuxSessionWindow("2", tmuxTargetStr)
		}
	})

	fmt.Println("Starting server on port 8999")
	http.ListenAndServe(":8999", nil)
}
