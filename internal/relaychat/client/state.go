package main

import (
	"flag"
	"os"
)

var (
	defAddr = "localhost:8080"
)

func getAddr() string {
	if value, exists := os.LookupEnv("RELAYCHAT_URI"); exists {
		defAddr = value
	}
	return *flag.String("addr", defAddr, "http service address")
}
