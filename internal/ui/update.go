package ui

import (
	"fmt"
	"path/filepath"
	"strings"
	"wizard-tutorial/internal/helpers"
	"wizard-tutorial/internal/ollama"
	"wizard-tutorial/internal/types"

	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/google/uuid"
)

func getBotResponseCmd(messages []types.Message) tea.Cmd {
	return func() tea.Msg {
		response := ollama.GetBotResponse(messages)
		return botResponseMsg{response: response, err: nil}
	}
}

func (m Main) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	chathistory, _ := helpers.GetHistoryFromJSON("chathistory.json")
	m.chatHistory = chathistory

	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		sideBarWidth := 35
		m.width = msg.Width
		m.height = msg.Height
		m.viewport.Width = msg.Width - sideBarWidth - 4
		m.viewport.Height = msg.Height - 4
		if !m.ready {
			m.viewport = viewport.New(m.viewport.Width, m.viewport.Height)
			m.ready = true
		}
		m.updateViewportContent()
		return m, nil

	case botResponseMsg:
		m.isLoading = false
		if msg.err != nil {
			m.err = msg.err
		} else {
			m.messages = append(m.messages, types.Message{
				Role:    "assistant",
				Content: msg.response,
			})
			m.err = nil
		}
		m.updateViewportContent()
		m.viewport.GotoBottom()

		return m, nil
	case tea.KeyMsg:
		if m.focus == 3 {
			switch msg.String() {
			case "up", "k":
				m.viewport.LineUp(2)
				return m, nil
			case "down", "j":
				m.viewport.LineDown(2)
				return m, nil
			case "pgup":
				m.viewport.HalfViewUp()
				return m, nil
			case "pgdown":
				m.viewport.HalfViewDown()
				return m, nil
			case "home":
				m.viewport.GotoTop()
				return m, nil
			case "end":
				m.viewport.GotoBottom()
				return m, nil
			}
		}
		switch msg.String() {

		case "tab":
			// Toggle focus between input and messages
			m.focus++

			if m.focus == 4 {
				m.input.Focus()
			} else {
				m.input.Blur()
				if m.focus > 4 {
					m.focus = 1
				}
			}

			return m, nil

		case "ctrl+c", "q":
			return m, tea.Quit
		case "enter":
			userMsg := strings.TrimSpace(m.input.Value())
			if userMsg == "" {
				return m, nil
			}
			m.messages = append(m.messages, types.Message{Role: "user", Content: userMsg})
			m.isLoading = true
			m.input.Reset()
			m.updateViewportContent()
			m.viewport.GotoBottom()
			return m, getBotResponseCmd(m.messages)

		default:
			m.input, cmd = m.input.Update(msg)

			return m, cmd
		}

	}
	return m, nil
}

func (m *Main) updateViewportContent() {
	if len(m.messages) == 0 && !m.isLoading {
		m.viewport.SetContent("Welcome! Start a conversation by typing a message below.\n\nTips:\n• Press Tab to scroll through messages\n• Press Esc to return to input\n• Press Ctrl+C to quit")
		// loadedMessages, err := helpers.GetMessagesFromJSON("file.json")

		m.viewport.GotoBottom()
		// if err != nil {
		// 	fmt.Println("Error loading messages:", err)
		// 	// handle error appropriately
		// } else {
		// 	m.messages = append(m.messages, loadedMessages...)
		// 	m.updateViewportContent()
		//
		// }
		return
	}
	if len(m.messages) == 1 {

		u := uuid.New()
		m.conversationId = u.String()
		m.chatHistory = append([]types.ChatHistory{{
			ConversationId: m.conversationId,
			Title:          m.messages[0].Content,
		}}, m.chatHistory...)
		helpers.SaveHistoryToJSON("chathistory.json", m.chatHistory)

	}

	var rendered []string
	filename := filepath.Join("conversations", m.conversationId+".json")
	helpers.SaveMessagesToJSON(filename, m.messages)
	for _, chat := range m.messages {
		if chat.Role == "assistant" {
			rendered = append(rendered, m.renderAssistantMessage(chat.Content))
		} else {
			rendered = append(rendered, m.renderUserMessage(chat.Content))
		}
	}

	if m.isLoading {
		loadingStyle := lipgloss.NewStyle().Foreground(lipgloss.Color("242")).Italic(true).MarginLeft(3).MarginTop(2)
		rendered = append(rendered, loadingStyle.Render("🤖 Thinking..."))
	}

	if m.err != nil {
		errorStyle := lipgloss.NewStyle().Foreground(lipgloss.Color("197")).Bold(true).MarginTop(2)
		rendered = append(rendered, errorStyle.Render(fmt.Sprintf("Error: %v", m.err)))
	}

	m.viewport.SetContent(strings.Join(rendered, "\n\n"))
}
