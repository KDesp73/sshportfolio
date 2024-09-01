package main

import (
	"log"
	"sshportfolio/internal/tui"
	tea "github.com/charmbracelet/bubbletea"
)


func main() {
	f, err := tea.LogToFile("debug.log", "debug")
	if err != nil {
		log.Fatalf("ERRO: %v", err)
	}
	defer f.Close()

	p := tea.NewProgram(tui.NewModel(), tea.WithAltScreen())
	if _, err := p.Run(); err != nil {
		log.Fatalf("%v\n", err)
	}
}

