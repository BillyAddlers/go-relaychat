package main

import (
	_ "embed"
	"flag"
	"github.com/gorilla/websocket"
	"html/template"
)

var (
	addr     = flag.String("addr", "localhost:8080", "http service address")
	upgrader = websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	}
	//go:embed template/index.html
	templateFile  string
	indexTemplate = template.Must(template.New("serveIndex").Parse(templateFile))

	//go:embed template/home.html
	templateHomeFile string
	homeTemplate     = template.Must(template.New("serveHome").Parse(templateHomeFile))
)
