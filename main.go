package main

import (
	"calendartui/ui"
	tea "charm.land/bubbletea/v2"
	"fmt"
)

func main() {
	p := tea.NewProgram(ui.InitialModel())
	if _, err := p.Run(); err != nil {
		fmt.Printf("%v\n", err)
	}
}
