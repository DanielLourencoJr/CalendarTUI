package ui

import (
	"calendartui/storage"
	"calendartui/task"
	"charm.land/bubbles/v2/textinput"
	tea "charm.land/bubbletea/v2"
	"strings"
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
	Store       *storage.Store
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

func InitialModel(tasks []task.Task, store *storage.Store) Model {
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
	DueTimeInput.Placeholder = "DD/MM/YYYY hh:mm(optional)"
	DueTimeInput.SetVirtualCursor(false)
	DueTimeInput.CharLimit = 156
	DueTimeInput.SetWidth(20)

	var inputs []textinput.Model
	inputs = append(inputs, TitleInput)
	inputs = append(inputs, DescriptionInput)
	inputs = append(inputs, DueTimeInput)

	return Model{
		Store:      store,
		Tasks:      tasks,
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

	dueTimeInfo, err := parseDueTime(m.Inputs[2].Value())
	if err != nil {
		return
	}
	dueTime := dueTimeInfo.dueTime
	hasDueTime := dueTimeInfo.hasDueTime
	newTask := task.Task{
		Name:        m.Inputs[0].Value(),
		Description: m.Inputs[1].Value(),
		DueTime:     dueTime,
		HasDueTime:  hasDueTime,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
	m.createTask(newTask)
	m.FocusIndex = 0
	m.clearInputs()
	m.ViewMode = ModeList
}

func (m *Model) createTask(newTask task.Task) {
	err := m.Store.CreateTask(newTask)
	if err != nil {
		err.Error()
	}
	tasks, err := m.Store.LoadAllTasks()
	if err != nil {
		return
	}
	m.Tasks = tasks
}

type DueTimeInfo struct {
	dueTime    time.Time
	hasDueTime bool
}

func parseDueTime(input string) (DueTimeInfo, error) {
	input = strings.TrimSpace(input)
	var dueTimeInfo DueTimeInfo
	if input == "" {
		return dueTimeInfo, nil
	}

	layouts := []string{
		"02/01/2006 15:04",
		"02/01/2006",
	}

	var lastErr error
	var hasDueTime bool
	for _, layout := range layouts {
		dueTime, err := time.ParseInLocation(layout, input, time.Local)
		if err == nil {
			if strings.Contains(layout, "15") {
				hasDueTime = true
			}
			dueTimeInfo.dueTime = dueTime
			dueTimeInfo.hasDueTime = hasDueTime
			return dueTimeInfo, nil
		}
		lastErr = err
	}
	return dueTimeInfo, lastErr
}
