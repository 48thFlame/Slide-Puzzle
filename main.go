package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
)

func newModel() model {
	return model{game: NewGame(3, 3)}
}

type model struct {
	game Game
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		case "m":
			m.game.Mix()
		case "up", "w":
			m.game.MoveOnBard(MoveUpToEmpty)
		case "down", "s":
			m.game.MoveOnBard(MoveDownToEmpty)
		case "left", "a":
			m.game.MoveOnBard(MoveLeftToEmpty)
		case "right", "d":
			m.game.MoveOnBard(MoveRightToEmpty)
		}
	}

	return m, nil
}

func (m model) View() string {
	return m.game.String()
}

func main() {
	p := tea.NewProgram(newModel())
	_, err := p.Run()
	if err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}
}
