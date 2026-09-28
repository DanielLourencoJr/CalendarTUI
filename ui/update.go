package ui

import (
	"calendartui/task"
	tea "charm.land/bubbletea/v2"
	"time"
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

func (m Model) updateListing(msg tea.KeyMsg) (tea.Model, tea.Cmd) {
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
		if len(m.Tasks) == 0 {
			return m, nil
		}
		m.Tasks[m.Cursor].IsCompleted = !m.Tasks[m.Cursor].IsCompleted

		m.Tasks[m.Cursor].UpdatedAt = time.Now()
	case "ctrl+o":
		m.ShowDetails = !m.ShowDetails
	case "a":
		m.ViewMode = ModeAdding
	case "ctrl+d":
		if len(m.Tasks) == 0 {
			return m, nil
		}
		taskId := m.Tasks[m.Cursor].Id
		err := m.Store.DeleteTask(taskId)
		if err != nil {
			return m, nil
		}
		var newTasks []task.Task
		for i, _ := range m.Tasks {
			if i == m.Cursor {
				continue
			} else {
				newTasks = append(newTasks, m.Tasks[i])
			}
		}
		if m.Cursor == len(m.Tasks)-1 && m.Cursor != 0 {
			m.Cursor = len(newTasks) - 1
		}
		m.Tasks = newTasks
	}
	return m, nil
}

func (m Model) updateAdding(msg tea.KeyMsg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	switch msg.String() {
	case "ctrl+c":
		return m, tea.Quit
	case "esc":
		m.ViewMode = ModeList
		m.clearInputs()
	case "enter", "tab", "shift-tab", "up", "down":
		s := msg.String()
		if s == "enter" && m.FocusIndex == len(m.Inputs)-1 {
			m.saveTask()
		} else if (s == "up" || s == "shift-tab") && m.FocusIndex > 0 {
			m.FocusIndex--
		} else if (s == "down" || s == "tab" || s == "enter") && m.FocusIndex < len(m.Inputs)-1 {
			m.FocusIndex++
		}

		cmds := make([]tea.Cmd, len(m.Inputs))
		for i := range m.Inputs {
			if i == m.FocusIndex {
				cmds[i] = m.Inputs[i].Focus()
				continue
			}
			m.Inputs[i].Blur()
		}

		return m, tea.Batch(cmds...)

	case "ctrl+s":
		m.saveTask()
	default:
		cmds := make([]tea.Cmd, len(m.Inputs))
		for i := range m.Inputs {
			m.Inputs[i], cmd = m.Inputs[i].Update(msg)
			cmds[i] = cmd
		}
		return m, tea.Batch(cmds...)
	}
	return m, nil
}
