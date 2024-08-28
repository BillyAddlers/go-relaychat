package main

import (
	"log"
	"net/http"
)

func serveIndex(w http.ResponseWriter, r *http.Request) {
	indexTemplate.Execute(w, "ws://"+r.Host+"/echo")
}

func serveHome(w http.ResponseWriter, r *http.Request) {
	log.Println(r.URL)
	if r.URL.Path != "/" {
		http.Error(w, "Not found", http.StatusNotFound)
		return
	}
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	homeTemplate.Execute(w, "ws://"+r.Host+"/ws")
}
