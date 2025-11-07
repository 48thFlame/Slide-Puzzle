package main

import (
	"fmt"
	"strconv"

	"github.com/48thFlame/Slide-Puzzle/slide"
	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/lipgloss/table"
)

const (
	whiteColor  = lipgloss.Color("#dddddd")
	purpleColor = lipgloss.Color("#5f5fff")
	footerColor = lipgloss.Color("#444444")
)

func slotToString(s slide.Slot) string {
	if s == slide.Empty {
		return "  "
	}

	si := int(s)
	sis := strconv.Itoa(si)

	if si >= 10 {
		return sis
	}

	return " " + sis
}

func viewSlot(s slide.Slot) string {
	return lipgloss.NewStyle().
		Bold(true).
		Foreground(whiteColor).
		Render(slotToString(s))
}

// creates a lipgloss table - which is the grid of the sliding puzzle
func createBoardGrid(g slide.Game) *table.Table {
	boardTableGrid := table.New().
		Border(lipgloss.DoubleBorder()).
		BorderRow(true).
		BorderStyle(
			lipgloss.NewStyle().Foreground(whiteColor)).
		StyleFunc(func(row, col int) lipgloss.Style {
			return lipgloss.NewStyle().Padding(0, 1)
		})

	for ri := range g.RowsNum {
		r := make([]string, 0, g.ColsNum)
		for ci := range g.ColsNum {
			slot := g.B[(g.ColsNum*(ri))+(ci)]
			r = append(r, viewSlot(slot))
		}
		boardTableGrid.Row(r...)
	}

	return boardTableGrid
}

func moveToString(m slide.BoardMovement) (s string) {
	switch m {
	case slide.MoveUpToEmpty:
		s = "Move Up   "
	case slide.MoveDownToEmpty:
		s = "Move Down "
	case slide.MoveLeftToEmpty:
		s = "Move Left "
	case slide.MoveRightToEmpty:
		s = "Move Right"
	}

	return
}

func viewWindow(m model) string {
	border := lipgloss.NewStyle().
		Border(lipgloss.ThickBorder()).
		BorderForeground(purpleColor).
		Padding(0, 1)

	title := lipgloss.NewStyle().
		Bold(true).
		Foreground(whiteColor).
		Background(purpleColor).
		Render("Sliding Puzzle")

	footer := lipgloss.NewStyle().
		Foreground(footerColor).
		Render("q: quit • m: mix")

	game := createBoardGrid(m.game).String()

	window := border.Render(
		lipgloss.JoinVertical(
			lipgloss.Center,
			title,
			game,
			footer,
		),
	)

	return window
}

func viewAi(g slide.Game, width, height int) string {
	border := lipgloss.NewStyle().
		Border(lipgloss.ThickBorder()).
		BorderForeground(purpleColor).
		Width(width).
		Height(height-2). // -2 for the border
		AlignHorizontal(lipgloss.Center).
		Padding(0, 1)

	title := lipgloss.NewStyle().
		Bold(true).
		Foreground(whiteColor).
		Background(purpleColor).
		Render("Ai Solution")

	var aiStr []string

	aiOutFlag, aiOut := slide.AiOutput(g)
	switch aiOutFlag {
	case slide.SolMove:
		aiStr = []string{
			"",
			"Do:",
			moveToString(aiOut.Move),
			"",
			"Length:",
			fmt.Sprint(aiOut.NumOfM)}
	case slide.Solved:
		aiStr = []string{"", "", "Solved!"}
	case slide.CantSolve:
		aiStr = []string{"", "Sorry", "Not", "Attempting", "That"}
	}

	return border.Render(
		lipgloss.JoinVertical(lipgloss.Center,
			append([]string{title}, aiStr...)...))
}
