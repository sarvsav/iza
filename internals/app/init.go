package app

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
)

func Init() error {
	// Initialize the application
	// This function will be called when the application starts
	// You can add your initialization code here
	fmt.Println("init called")
	fmt.Println("Check for the connection")
	fmt.Println("Check for overriding the default values")
	m := InitialModel()
	p := tea.NewProgram(m)
	if _, err := p.Run(); err != nil {
		fmt.Println("Error running program:", err)
		os.Exit(1)
	}

	return nil
}
