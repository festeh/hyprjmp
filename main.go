package main

import (
	"encoding/json"
	"hyprjmp/utils"
	"log"
	"net/http"
	"sync"
	"time"
)

var (
	tmuxInfo  hyprjmp.TmuxInfo
	cacheLock sync.RWMutex
)

func tmuxWorker() {
	for {
		info, err := hyprjmp.GetTmuxInfo()
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

func main() {
	go tmuxWorker()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Add("Access-Control-Allow-Headers", "Accept, Content-Type, Access-Control-Allow-Headers, Authorization, X-Requested-With")
		w.Header().Add("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		cacheLock.RLock()
		// marshal tmuxInfo to json
		bytes, err := json.Marshal(tmuxInfo)
		if err != nil {
			log.Fatal(err)
		}
		cacheLock.RUnlock()
		w.Write(bytes)
	})

	http.ListenAndServe(":8999", nil)
}
