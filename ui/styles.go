package ui

import (
	lipgloss "charm.land/lipgloss/v2"
)

var (
	DoneStyle    = lipgloss.NewStyle().Foreground(lipgloss.Color("2")).Strikethrough(true)
	PendingStyle = lipgloss.NewStyle()
	TitleStyle   = lipgloss.NewStyle().Bold(true).Underline(true)
)
