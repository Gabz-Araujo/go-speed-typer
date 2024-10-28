package main

import (
	"fmt"

	"github.com/GabzAraujo/go-speed-typer/internal/game"
	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	engine := game.NewGameEngine()
	p := tea.NewProgram(engine)
	if _, err := p.Run(); err != nil {
		fmt.Printf("Error running program: %v\n", err)
	}
}
