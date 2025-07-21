package ui

import (
	"strings"

	"github.com/charmbracelet/glamour"
	"github.com/charmbracelet/glamour/ansi"
	"github.com/charmbracelet/lipgloss"
)

func strPtr(s string) *string { return &s }

func (m Main) renderUserMessage(content string) string {
	userBorder := lipgloss.Border{
		Top: "─", Bottom: "─", Left: "│", Right: "│",
		TopLeft: "╭", TopRight: "╮", BottomLeft: "╰", BottomRight: "╯",
	}

	bubble := lipgloss.NewStyle().Border(userBorder).BorderForeground(lipgloss.Color("42")).Padding(0, 1).MarginLeft(11).Align(lipgloss.Left)
	if len(content) > 102 {
		bubble.Width(122)
	}
	return lipgloss.NewStyle().Width(m.width - 42).Align(lipgloss.Right).Render(bubble.Render("You: " + content))
}

func (m Main) renderAssistantMessage(content string) string {
	botBorder := lipgloss.Border{
		Top: "─", Bottom: "─", Left: "│", Right: "│",
		TopLeft: "┌", TopRight: "╮", BottomLeft: "╰", BottomRight: "╯",
	}
	width := min(len(content)+12, 130)

	if width > m.width-SidebarWidth-9 {
		width = m.width - SidebarWidth - 11
	}

	renderer, err := glamour.NewTermRenderer(
		glamour.WithStyles(ansi.StyleConfig{
			CodeBlock: ansi.StyleCodeBlock{
				StyleBlock: ansi.StyleBlock{
					StylePrimitive: ansi.StylePrimitive{BackgroundColor: strPtr("239"), Color: strPtr("253")},
				},
			},
			Code: ansi.StyleBlock{
				StylePrimitive: ansi.StylePrimitive{BackgroundColor: strPtr("241"), Color: strPtr("253")},
			},
		}),
		glamour.WithAutoStyle(),
		glamour.WithWordWrap(width),
		glamour.WithStandardStyle("dark"),
	)

	var renderedContent string
	if err != nil {
		renderedContent = content
	} else {
		rendered, err := renderer.Render(content)
		if err != nil {
			renderedContent = content
		} else {
			renderedContent = strings.TrimSpace(rendered)
		}
	}

	botStyle := lipgloss.NewStyle().Border(botBorder).BorderForeground(lipgloss.Color("244")).Padding(0, 1).MarginLeft(3).Align(lipgloss.Left)
	if width > 201 {
		botStyle.Width(201)
	}
	if m.width < width {
		botStyle.Width(m.width - 50)
	}
	botStyle.Width(50)
	return botStyle.Render("Assistant:\n" + renderedContent)
}
