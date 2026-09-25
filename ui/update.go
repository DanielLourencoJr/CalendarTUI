package ui

import (
	tea "charm.land/bubbletea/v2"
	"time"
	task "calendartui/task"
)

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch m.ViewMode {
		case ModeList:
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
			case "a":
				m.ViewMode = 1
			}
		case ModeAdding:
			switch msg.String() {
			case "ctrl+c":
				return m, tea.Quit
			case "esc":
				m.ViewMode = ModeList
				m.TaskTitleInput.SetValue("")
				return m, nil
			case "enter":
				newTask := task.Task{
					Name: m.TaskTitleInput.Value(),
					CreatedAt: time.Now(),
					UpdatedAt: time.Now(),
				}
				m.Tasks = append(m.Tasks, newTask)
				m.TaskTitleInput.SetValue("")
				m.ViewMode = ModeList
				return m,nil
			default:
				m.TaskTitleInput, cmd = m.TaskTitleInput.Update(msg)
				return m, cmd
			}
		case ModeEditing:
		}
	}
	return m, nil
}