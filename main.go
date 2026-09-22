package main

import (
	tea "charm.land/bubbletea/v2"
	"fmt"
)

func main() {
	p := tea.NewProgram(InitialModel())
	if _, err := p.Run(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
	}
}
