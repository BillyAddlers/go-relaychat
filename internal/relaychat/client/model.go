package main

import (
	"fmt"
	"github.com/charmbracelet/bubbles/cursor"
	"github.com/charmbracelet/bubbles/textarea"
	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/gorilla/websocket"
	"os"
	"strings"
)

var (
	name username
	conn *websocket.Conn
)

func Run(c *websocket.Conn, initModel model) {
	// Initialize values
	name.name = "Anonymous"
	name.set = false
	conn = c

	p := tea.NewProgram(initModel)
	go func() {
		for {
			p.Send(readWebSocketMessages()())
		}
	}()
	if _, err := p.Run(); err != nil {
		fmt.Fprintf(os.Stderr, "Oof: %v\n", err)
	}
}

func initialModel() model {
	ta := textarea.New()
	ta.Placeholder = "Send a Message..."
	ta.Focus()

	ta.Prompt = "┃ "
	ta.CharLimit = 280

	ta.SetWidth(30)
	ta.SetHeight(3)

	// Remove cursor line styling
	ta.FocusedStyle.CursorLine = lipgloss.NewStyle()

	ta.ShowLineNumbers = false

	vp := viewport.New(30, 10)
	vp.SetContent(`Welcome to the chat room!
Please type your username and press Enter to send.`)

	ta.KeyMap.InsertNewline.SetEnabled(false)

	return model{
		textarea:    ta,
		messages:    []string{},
		viewport:    vp,
		senderStyle: lipgloss.NewStyle().Foreground(lipgloss.Color("5")),
		err:         nil,
	}
}

func (m model) Init() tea.Cmd {
	return textarea.Blink
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case OnMessage:
		sender, message := splitMessage(msg.Message)
		sender = m.senderStyle.Render(sender + ":")
		m.messages = append(m.messages, sender+message)
		m.viewport.SetContent(strings.Join(m.messages, "\n"))
		m.viewport.GotoBottom()
		return m, nil
	case tea.WindowSizeMsg:
		m.viewport.Width = msg.Width
		m.textarea.SetWidth(msg.Width)
		return m, nil
	case tea.KeyMsg:
		switch msg.String() {
		case "esc", "ctrl+c":
			// Quit.
			fmt.Println(m.textarea.Value())
			return m, tea.Quit
		case "enter":
			v := m.textarea.Value()

			if v == "" {
				// Don't send empty messages.
				return m, nil
			}

			if !name.set {
				name.name = v
				name.set = true
				m.viewport.SetContent("Welcome, " + name.name + "!")
				m.textarea.Reset()
				m.viewport.GotoBottom()
				return m, nil
			}

			// Simulate sending a Message. In your application you'll want to
			// also return a custom command to send the Message off to
			// a server.
			err := conn.WriteMessage(websocket.TextMessage, []byte(name.name+": "+v))
			if err != nil {
				m.messages = append(m.messages, "Error sending Message: "+err.Error())
			}
			m.textarea.Reset()
			m.viewport.GotoBottom()
			return m, nil
		default:
			// Send all other keypresses to the textarea.
			var cmd tea.Cmd
			m.textarea, cmd = m.textarea.Update(msg)
			return m, cmd
		}

	case cursor.BlinkMsg:
		// Textarea should also process cursor blinks.
		var cmd tea.Cmd
		m.textarea, cmd = m.textarea.Update(msg)
		return m, cmd

	default:
		return m, nil
	}
}

func (m model) View() string {
	return fmt.Sprintf(
		"%s\n\n%s",
		m.viewport.View(),
		m.textarea.View(),
	) + "\n\n"
}
