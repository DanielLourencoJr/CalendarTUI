package ui

import (
	"calendartui/task"
	"time"
	tea "charm.land/bubbletea/v2"
	"charm.land/bubbles/v2/textinput"
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
	ViewMode Mode
	TaskTitleInput textinput.Model
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
				Name: "Study Go",
				Description: "Learning a new programming language",
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
				IsCompleted: false,
			},
			{
				Name: "Make chapter exercises",
				Description: "So I don't reprove in exams",
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
				IsCompleted: false,
			},
			{
				Name: "Keep reading that book",
				Description: "It's very interessting. I shouldn't wait so much to resume reading.",
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
				IsCompleted: false,
			},
		},
		ViewMode: 0,
		Cursor: 0,
		TaskTitleInput: TextInput,
	}
}
