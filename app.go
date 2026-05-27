package main

import (
	tea "github.com/charmbracelet/bubbletea"
)

func runLocalPortfolio() {
	p := tea.NewProgram(NewModel(nil), tea.WithAltScreen())
	if err := p.Start(); err != nil {
		panic(err)
	}
}
