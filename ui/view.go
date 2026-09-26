package ui

import (
	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"
	"fmt"
)

func (m Model) View() tea.View {
	switch m.ViewMode {
	case ModeList:
		s := "Your Tasks:\n\n"

		if len(m.Tasks) == 0 {
			s += "No tasks yet. Press a to add one.\n"
		} else {
			for i, t := range m.Tasks {
				Cursor := " "
				if m.Cursor == i {
					Cursor = ">"
				}
				row := fmt.Sprintf("%s %s", Cursor, t.Name)

				if t.IsCompleted {
					row = DoneStyle.Render(row)
				} else {
					row = PendingStyle.Render(row)
				}
				s += row + "\n"
				if m.ShowDescription {
					s += fmt.Sprintf("    %s\n", t.Description)
				}

			}
		}
		s += "\nPress q to quit\n"
		return tea.NewView(s)
	case ModeAdding:
		var c *tea.Cursor
		if !m.TaskTitleInput.VirtualCursor() {
			if c != nil {
				c = m.TaskTitleInput.Cursor()
				c.Y += lipgloss.Height(m.TaskTitleView())
			}
		}

		str := lipgloss.JoinVertical(lipgloss.Top, m.TaskTitleView(), m.TaskTitleInput.View(), m.footerView())

		v := tea.NewView(str)
		v.Cursor = c
		return v
	}
	return tea.NewView("")
}

func (m Model) TaskTitleView() string {
	return "Create a new task:\n"
}

func (m Model) footerView() string { return "\n(esc to quit)" }
