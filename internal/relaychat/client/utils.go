package main

import "strings"

func splitMessage(message string) (string, string) {
	parts := strings.SplitN(message, ":", 2)
	if len(parts) < 2 {
		return parts[0], ""
	}
	return parts[0], parts[1]
}
