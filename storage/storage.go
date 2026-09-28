package storage

import (
	"calendartui/task"
	"database/sql"
	_ "modernc.org/sqlite"
	"time"
)

func Open(path string) (*Store, error) {
	db, err := sql.Open("sqlite", path)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		db.Close()
		return nil, err
	}

	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS tasks (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			name TEXT NOT NULL,
			description TEXT NOT NULL DEFAULT '',
			due_time TEXT,
			has_due_time BOOLEAN NOT NULL DEFAULT FALSE,
			is_completed BOOLEAN NOT NULL DEFAULT FALSE,
			created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
			updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
		)
	`)
	if err != nil {
		db.Close()
		return nil, err
	}

	return &Store{db: db}, nil
}

type Store struct {
	db *sql.DB
}

func (s *Store) LoadAllTasks() ([]task.Task, error) {
	rows, err := s.db.Query(`
		SELECT id, name, description, due_time, has_due_time, is_completed, created_at, updated_at FROM tasks
	`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tasks []task.Task
	for rows.Next() {
		var id int
		var name, description string
		var dueTime sql.NullString
		var completed, hasDueTime bool
		var createdAt, updatedAt time.Time

		if err := rows.Scan(
			&id,
			&name,
			&description,
			&dueTime,
			&hasDueTime,
			&completed,
			&createdAt,
			&updatedAt,
		); err != nil {
			return nil, err
		}

		var parsedDueTime time.Time
		if dueTime.Valid {
			parsed, err := time.Parse("2006-01-02 15:04", dueTime.String)
			if err != nil {
				return nil, err
			}
			parsedDueTime = parsed
		}

		task := task.Task{
			Id:          id,
			Name:        name,
			Description: description,
			DueTime:     parsedDueTime,
			IsCompleted: completed,
			CreatedAt:   createdAt,
			UpdatedAt:   updatedAt,
		}
		tasks = append(tasks, task)
	}
	if rows.Err() != nil {
		return nil, rows.Err()
	}
	return tasks, nil
}

func (s *Store) CreateTask(newTask task.Task) error {
	var dueTime any

	if !newTask.DueTime.IsZero() {
		dueTime = newTask.DueTime.Format("2006-01-02 15:04")
	}
	_, err := s.db.Exec(`
		INSERT INTO tasks (name, description, due_time, has_due_time) VALUES (?, ?, ?, ?)
	`, newTask.Name, newTask.Description, dueTime, newTask.HasDueTime)
	return err
}

func (s *Store) DeleteTask(taskId int) error {
	_, err := s.db.Exec(`
		DELETE FROM tasks WHERE id = ?
	`, taskId)
	if err != nil {
		return err
	}
	return nil
}

func (s *Store) Close() error {
	return s.db.Close()
}
