package ui

import (
	lipgloss "github.com/charmbracelet/lipgloss"
)

var (
	DoneStyle    = lipgloss.NewStyle().Foreground(lipgloss.Color("2")).Strikethrough(true)
	PendingStyle = lipgloss.NewStyle()
	TitleStyle = lipgloss.NewStyle().Bold(true).Underline(true)
)