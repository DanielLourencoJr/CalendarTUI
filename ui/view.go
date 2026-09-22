package ui

import (
	tea "charm.land/bubbletea/v2"
	"fmt"
)

func (m Model) View() tea.View {
	s := "Your Tasks:\n\n"
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
	s += "\nPress q to quit\n"
	return tea.NewView(s)
}