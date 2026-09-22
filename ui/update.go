package ui

import (
	tea "charm.land/bubbletea/v2"
	"time"
)

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		case "up", "k":
			if m.Cursor > 0 {
				m.Cursor--
			}
		case "down", "j":
			if m.Cursor < len(m.Tasks)-1 {
				m.Cursor++
			}
		case "space", "enter":
			m.Tasks[m.Cursor].IsCompleted = !m.Tasks[m.Cursor].IsCompleted

			m.Tasks[m.Cursor].UpdatedAt = time.Now()
		case "ctrl+o":
			m.ShowDescription = !m.ShowDescription
		}
	}
	return m, nil
}