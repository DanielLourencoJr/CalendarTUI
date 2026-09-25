package ui

import (
	"calendartui/task"
	"charm.land/bubbles/v2/textinput"
	tea "charm.land/bubbletea/v2"
	"time"
)

type Mode int

const (
	ModeList Mode = iota
	ModeAdding
	ModeEditing
)

func (m Mode) String() string {
	switch m {
	case ModeList:
		return "list"
	case ModeAdding:
		return "adding"
	case ModeEditing:
		return "editing"
	default:
		return "unknown"
	}
}

type Model struct {
	Tasks           []task.Task
	Cursor          int
	ShowDescription bool
	ViewMode        Mode
	TaskTitleInput  textinput.Model
}

func (m Model) Init() tea.Cmd {
	return nil
}

func InitialModel() Model {
	TextInput := textinput.New()
	TextInput.Placeholder = "Task Name"
	TextInput.SetVirtualCursor(false)
	TextInput.Focus()
	TextInput.CharLimit = 156
	TextInput.SetWidth(20)
	return Model{
		Tasks: []task.Task{
			{
				Name:        "Study Go",
				Description: "Learning a new programming language",
				CreatedAt:   time.Now(),
				UpdatedAt:   time.Now(),
				IsCompleted: false,
			},
			{
				Name:        "Make chapter exercises",
				Description: "So I don't reprove in exams",
				CreatedAt:   time.Now(),
				UpdatedAt:   time.Now(),
				IsCompleted: false,
			},
			{
				Name:        "Keep reading that book",
				Description: "It's very interessting. I shouldn't wait so much to resume reading.",
				CreatedAt:   time.Now(),
				UpdatedAt:   time.Now(),
				IsCompleted: false,
			},
		},
		ViewMode:       0,
		Cursor:         0,
		TaskTitleInput: TextInput,
	}
}

func (m Model) updateAdding(msg tea.KeyMsg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	switch msg.String() {
	case "ctrl+c":
		return m, tea.Quit
	case "esc":
		m.ViewMode = ModeList
		m.TaskTitleInput.SetValue("")
	case "enter":
		newTask := task.Task{
			Name:      m.TaskTitleInput.Value(),
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		}
		m.Tasks = append(m.Tasks, newTask)
		m.TaskTitleInput.SetValue("")
		m.ViewMode = ModeList
	default:
		m.TaskTitleInput, cmd = m.TaskTitleInput.Update(msg)
		return m, cmd
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
		m.Tasks[m.Cursor].IsCompleted = !m.Tasks[m.Cursor].IsCompleted

		m.Tasks[m.Cursor].UpdatedAt = time.Now()
	case "ctrl+o":
		m.ShowDescription = !m.ShowDescription
	case "a":
		m.ViewMode = ModeAdding
	}
	return m, nil
}
