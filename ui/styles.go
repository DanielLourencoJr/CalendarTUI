package ui

import (
	lipgloss "charm.land/lipgloss/v2"
)

var (
	DoneTaskStyle    = lipgloss.NewStyle().Foreground(lipgloss.Color("2")).Strikethrough(true)
	DoneDetailsStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("8"))
	PendingStyle = lipgloss.NewStyle()
	TitleStyle   = lipgloss.NewStyle().Bold(true).Underline(true)
)
