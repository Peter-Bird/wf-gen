// wf-gen/main.go
package main

import (
	"log"
	"net/http"
	"os"
	"path/filepath"
	"wf-gen/ep"

	"github.com/Peter-Bird/ws"
)

func init() {
	// Set log prefix based on application name from args
	appName := filepath.Base(os.Args[0]) // Removes leading "./" if present
	log.SetPrefix("[" + appName + "] ")
	//log.SetFlags(0) // Optional: removes default date and time from log output
}

func main() {
	ep.List()
	run()
}

func run() {
	http.HandleFunc("/", ws.Handler)

	log.Println("Starting server on :8082...")
	err := http.ListenAndServe(":8082", nil)
	if err != nil {
		log.Fatal("Server failed to start:", err)
	}
}
