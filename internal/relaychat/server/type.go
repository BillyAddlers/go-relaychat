package main

import "github.com/gorilla/websocket"

type Message struct {
	Username string `json:"username"`
	Message  string `json:"message"`
}

// Hub maintains the set of active clients and broadcasts messages to the websocket.
type Hub struct {
	// Registered Clients.
	Clients map[*Client]bool

	// Inbound messages from the clients.
	broadcast chan []byte

	// register requests from the clients.
	register chan *Client

	// unregister requests from clients.
	unregister chan *Client
}

// Client is a middleman between the websocket connection and the hub.
type Client struct {
	Hub *Hub

	// The websocket connection.
	conn *websocket.Conn

	// Buffered channel of outbound messages.
	send chan []byte
}
