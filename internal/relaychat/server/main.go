package main

import (
	_ "embed"
	"flag"
	"log"
	"net/http"
)

func main() {
	flag.Parse()
	hub := NewHub()
	go hub.Run()
	log.SetFlags(0)
	http.HandleFunc("/project", serveIndex)
	http.HandleFunc("/", serveHome)
	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		serveWs(hub, w, r)
	})
	http.HandleFunc("/echo", echo)
	log.Fatal(http.ListenAndServe(*addr, nil))
}
