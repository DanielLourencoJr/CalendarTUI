package task

import "time"

type Task struct {
	Id          int
	Name        string
	Description string
	DueTime     time.Time
	HasDueTime  bool
	CreatedAt   time.Time
	UpdatedAt   time.Time
	IsCompleted bool
}
