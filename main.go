package main

import (
	"calendartui/storage"
	"calendartui/ui"
	tea "charm.land/bubbletea/v2"
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	databasePath, err := databasePath()
	if err != nil {
		fmt.Printf("%v", err)
		return
	}
	store, err := storage.Open(databasePath)
	if err != nil {
		fmt.Printf("%v", err)
		return
	}
	defer store.Close()

	tasks, err := store.LoadAllTasks()
	if err != nil {
		fmt.Printf("%v", err)
		return
	}

	p := tea.NewProgram(ui.InitialModel(tasks, store))
	if _, err := p.Run(); err != nil {
		fmt.Printf("%v\n", err)
	}
}

func databasePath() (string, error) {
	dataHome := os.Getenv("XDG_DATA_HOME")

	if dataHome == "" || !filepath.IsAbs(dataHome) {
		home, err := os.UserHomeDir()
		if err != nil {
			return "", err
		}
		dataHome = filepath.Join(home, ".local", "share")
	}

	appDir := filepath.Join(dataHome, "calendartui")
	if err := os.MkdirAll(appDir, 0o700); err != nil {
		return "", err
	}

	return filepath.Join(appDir, "calendar.db"), nil
}
