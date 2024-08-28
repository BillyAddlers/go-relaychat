package utils

import (
	"time"
)

const (
	// WriteWait Time allowed to write a message to the peer.
	WriteWait = 10 * time.Second

	// PongWait Time allowed to read the next pong message from the peer.
	PongWait = 60 * time.Second

	// PingPeriod Send pings to peer with this period. Must be less than PongWait.
	PingPeriod = (PongWait * 9) / 10

	// MaxMessageSize Maximum message size allowed from peer.
	MaxMessageSize = 512
)

var (
	// Newline and space are used to sanitize incoming messages.
	Newline = []byte{'\n'}

	// Space is used to sanitize incoming messages.
	Space = []byte{' '}
)
