package ui

import (
	tea "charm.land/bubbletea/v2"
	"fmt"
	"strings"
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
				if m.ShowDetails {
					s += fmt.Sprintf("    %s", t.Description)

					s += fmt.Sprintf("    %s\n", t.DueTime.Format("02/01/2006"))
				}

			}
		}
		s += "\nPress q to quit\n"
		return tea.NewView(s)
	case ModeAdding:
		var b strings.Builder
		for i, _ := range m.Inputs {
			b.WriteString(m.Inputs[i].View())
			if i < len(m.Inputs)-1 {
				b.WriteRune('\n')
			}
		}
		return tea.NewView(b.String())
	}
	return tea.NewView("")
}

func (m Model) TaskTitleView() string {
	return "Create a new task:\n"
}

func (m Model) footerView() string { return "\n(esc to quit)" }
