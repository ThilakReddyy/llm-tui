package ui

import (
	"github.com/charmbracelet/lipgloss"
)

func (m Main) View() string {
	if m.width < 51 || m.height < 10 {
		return "Terminal too small. Please resize."
	}

	var sidebarHistory []string
	for _, history := range m.chatHistory {
		sidebarHistory = append(sidebarHistory, SidebarHistory.Render(history.Title))
	}
	sidebar := SidebarStyle.Height(m.height).Render(
		lipgloss.JoinVertical(lipgloss.Top,
			SidebarGptTitle.Render("ARCH GPT"),
			SidebarHeader.Render("History"),
			SidebarEmpty.Render("No history to be found"),
		),
	)

	if len(m.chatHistory) > 0 {
		sidebar = SidebarStyle.Height(m.height).Render(
			lipgloss.JoinVertical(lipgloss.Top,
				append([]string{
					SidebarGptTitle.Render("ARCH GPT"),
					SidebarHeader.Render("History"),
				}, sidebarHistory...)...,
			),
		)
	}

	viewportStyle := ViewportStyle(m.width, m.height)
	inputStyle := InputStyle(m.width, m.height)

	return lipgloss.JoinHorizontal(lipgloss.Top,
		sidebar,
		lipgloss.JoinVertical(lipgloss.Top,
			viewportStyle.Render(m.viewport.View()),
			inputStyle.Render(m.input.View()),
		),
	)
}
