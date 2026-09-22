package ui

import (
	"calendartui/task"
	"time"
	tea "charm.land/bubbletea/v2"
)

type Model struct {
	Tasks           []task.Task
	Cursor          int
	ShowDescription bool
}

func (m Model) Init() tea.Cmd {
	return nil
}

func InitialModel() Model {
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
		Cursor: 0,
	}
}
