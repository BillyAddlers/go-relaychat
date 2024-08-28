package main

import tea "github.com/charmbracelet/bubbletea"

func readWebSocketMessages() tea.Cmd {
	return func() tea.Msg {
		for {
			_, message, err := conn.ReadMessage()
			if err != nil {
				return OnMessage{Message: "Error receiving Message: " + err.Error()}
			}
			return OnMessage{Message: string(message)}
		}
	}
}
