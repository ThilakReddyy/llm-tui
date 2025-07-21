package ui

import "github.com/charmbracelet/lipgloss"

var (
	SidebarWidth    = 35
	SidebarGptTitle = lipgloss.NewStyle().Align(lipgloss.Center).Bold(true).Width(SidebarWidth - 4).Border(lipgloss.RoundedBorder())
	SidebarHeader   = lipgloss.NewStyle().Italic(true).PaddingTop(2).Foreground(lipgloss.Color("#afafaf"))
	SidebarEmpty    = lipgloss.NewStyle().Align(lipgloss.Center).Width(SidebarWidth - 3).MarginTop(2)
	SidebarHistory  = lipgloss.NewStyle().
			Align(lipgloss.Left).
			Background(lipgloss.Color("#afafaf")).
			Foreground(lipgloss.Color("#000")).
			Width(SidebarWidth - 3).
			MarginTop(1).
			Padding(1)
	SidebarStyle     = lipgloss.NewStyle().Height(0).Padding(1).Width(SidebarWidth).BorderRight(true).BorderStyle(lipgloss.RoundedBorder())
	ErrorTextStyle   = lipgloss.NewStyle().Foreground(lipgloss.Color("197")).Bold(true).MarginTop(2)
	LoadingTextStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("242")).Italic(true).MarginLeft(3).MarginTop(2)
)

func ViewportStyle(mwidth int, mheight int) lipgloss.Style {
	return lipgloss.NewStyle().Height(mheight - 3).Width(mwidth - SidebarWidth - 4)
}

func InputStyle(mwidth int, mheight int) lipgloss.Style {
	return lipgloss.NewStyle().Border(lipgloss.RoundedBorder()).BorderForeground(lipgloss.Color("241")).Padding(0, 1).Width(mwidth - SidebarWidth - 4)
}
