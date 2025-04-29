package app

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
)

func Init() error {
	m := InitialModel()
	p := tea.NewProgram(m)
	if _, err := p.Run(); err != nil {
		fmt.Println("Error running program:", err)
		os.Exit(1)
	}

	return nil
}
