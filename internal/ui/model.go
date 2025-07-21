package ui

import (
	"wizard-tutorial/internal/types"

	"github.com/charmbracelet/bubbles/textinput"
	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
)

type botResponseMsg struct {
	response string
	err      error
}

type Main struct {
	width           int
	height          int
	input           textinput.Model
	messages        []types.Message
	viewport        viewport.Model
	conversationId  string
	chatHistory     []types.ChatHistory
	ready           bool
	focusOnMessages bool
	focus           int
	scrollOffset    int
	isLoading       bool
	err             error
}

func InitialModel() Main {
	ti := textinput.New()
	ti.Placeholder = "Ask anything..."
	ti.Focus()
	ti.CharLimit = 1000
	ti.Width = 50

	vp := viewport.New(78, 20)
	vp.SetContent("")

	return Main{
		input:    ti,
		messages: []types.Message{},
		viewport: vp,
		focus:    4,
	}
}

func (m Main) Init() tea.Cmd {
	return textinput.Blink
}
