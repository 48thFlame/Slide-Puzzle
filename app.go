package main

import (
	"github.com/48thFlame/Slide-Puzzle/slide"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

func newModel(g slide.Game) model {
	return model{game: g}
}

type model struct {
	game slide.Game
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
			m.game.MoveOnBard(slide.MoveUpToEmpty)
		case "down", "s":
			m.game.MoveOnBard(slide.MoveDownToEmpty)
		case "left", "a":
			m.game.MoveOnBard(slide.MoveLeftToEmpty)
		case "right", "d":
			m.game.MoveOnBard(slide.MoveRightToEmpty)
		}
	}

	return m, nil
}

func (m model) View() string {
	windowView := viewWindow(m)
	windowWidth := lipgloss.Width(windowView)
	windowHeight := lipgloss.Height(windowView) - 2 // -2 for the border

	aiView := viewAi(m.game, windowWidth, windowHeight)

	return lipgloss.JoinHorizontal(lipgloss.Right, windowView, aiView)
}
