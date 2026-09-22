package main

import (
	tea "charm.land/bubbletea/v2"
	"fmt"
	lipgloss "github.com/charmbracelet/lipgloss"
	"time"
)

type Task struct {
	Name        string
	Description string
	CreatedAt   time.Time
	UpdatedAt   time.Time
	IsCompleted bool
}

type Model struct {
	Tasks           []Task
	Cursor          int
	ShowDescription bool
}

func (m Model) Init() tea.Cmd {
	return nil
}

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

var (
	doneStyle    = lipgloss.NewStyle().Foreground(lipgloss.Color("2")).Strikethrough(true)
	pendingStyle = lipgloss.NewStyle()
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
			row = doneStyle.Render(row)
		} else {
			row = pendingStyle.Render(row)
		}
		s += row + "\n"
		if m.ShowDescription {
			s += fmt.Sprintf("    %s\n", t.Description)
		}

	}
	s += "\nPress q to quit\n"
	return tea.NewView(s)
}

func InitialModel() Model {
	return Model{
		Tasks: []Task{
			Task{
				"Study Go",
				"Learning a new programming language",
				time.Now(),
				time.Now(),
				false,
			},
			Task{
				"Make chapter exercises",
				"So I don't reprove in exams",
				time.Now(),
				time.Now(),
				false,
			},
			Task{
				"Keep reading that book",
				"It's very interessting. I shouldn't wait so much to resume reading.",
				time.Now(),
				time.Now(),
				false,
			},
		},
		Cursor: 0,
	}
}
