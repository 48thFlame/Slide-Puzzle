package main

import (
	"github.com/48thFlame/Slide-Puzzle/slide"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

func newModel(rowsNum, colsNum int) model {
	return model{game: slide.NewGame(rowsNum, colsNum)}
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
	windowView := ViewWindow(m)
	windowHeight := lipgloss.Height(windowView)

	aiView := lipgloss.NewStyle().
		Border(lipgloss.ThickBorder()).
		BorderForeground(purpleColor).
		Foreground(whiteColor).
		Height(windowHeight - 2). // -2 for the border
		Render(slide.AiOutput(m.game))

	return lipgloss.JoinHorizontal(lipgloss.Right, windowView, aiView)
}
