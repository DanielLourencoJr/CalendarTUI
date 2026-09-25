package ui

import (
	tea "charm.land/bubbletea/v2"
)

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch m.ViewMode {
		case ModeList:
			return m.updateListing(msg)
		case ModeAdding:
			return m.updateAdding(msg)
		case ModeEditing:
		}
	}
	return m, nil
}
