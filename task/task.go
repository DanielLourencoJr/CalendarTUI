package task

import "time"

type Task struct {
	Name        string
	Description string
	CreatedAt   time.Time
	UpdatedAt   time.Time
	IsCompleted bool
}
