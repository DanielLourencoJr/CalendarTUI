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
	Tasks       []task.Task
	Cursor      int
	ShowDetails bool
	ViewMode    Mode
	Inputs      []textinput.Model
	FocusIndex  int
}

func (m Model) Init() tea.Cmd {
	return nil
}

func InitialModel() Model {
	TitleInput := textinput.New()
	TitleInput.Placeholder = "Task Name"
	TitleInput.SetVirtualCursor(false)
	TitleInput.Focus()
	TitleInput.CharLimit = 156
	TitleInput.SetWidth(20)

	DescriptionInput := textinput.New()
	DescriptionInput.Placeholder = "Task Description"
	DescriptionInput.SetVirtualCursor(false)
	DescriptionInput.CharLimit = 1024
	DescriptionInput.SetWidth(50)

	DueTimeInput := textinput.New()
	DueTimeInput.Placeholder = "Task Due Time"
	DueTimeInput.SetVirtualCursor(false)
	DueTimeInput.CharLimit = 156
	DueTimeInput.SetWidth(20)

	var inputs []textinput.Model
	inputs = append(inputs, TitleInput)
	inputs = append(inputs, DescriptionInput)
	inputs = append(inputs, DueTimeInput)

	return Model{
		Tasks:      []task.Task{},
		ViewMode:   ModeList,
		Cursor:     0,
		Inputs:     inputs,
		FocusIndex: 0,
	}
}

func (m *Model) clearInputs() {
	for i := range m.Inputs {
		m.Inputs[i].SetValue("")
	}
}

func (m *Model) saveTask() {
	dueTime, err := time.Parse("02/01/2006", m.Inputs[2].Value())
	var newTask task.Task
	if err != nil {
		newTask = task.Task{
			Name:        m.Inputs[0].Value(),
			Description: m.Inputs[1].Value(),
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		}
	} else {
		newTask = task.Task{
			Name:        m.Inputs[0].Value(),
			Description: m.Inputs[1].Value(),
			DueTime:     dueTime,
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		}
	}
	m.Tasks = append(m.Tasks, newTask)
	m.FocusIndex = 0
	m.clearInputs()
	m.ViewMode = ModeList
}
