package task

import "time"

type Task struct {
	Name        string
	Description string
	DueTime     time.Time
	CreatedAt   time.Time
	UpdatedAt   time.Time
	IsCompleted bool
}
