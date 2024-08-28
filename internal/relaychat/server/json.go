package main

import (
	"encoding/json"
	"log"
)

func parseJSON(msg []byte) Message {
	var m Message
	if err := json.Unmarshal(msg, &m); err != nil {
		log.Println("Error parsing JSON: ", err)
	}
	return m
}
