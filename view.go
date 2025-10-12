package main

import (
	"strconv"

	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/lipgloss/table"
)

const (
	whiteColor  = lipgloss.Color("#dddddd")
	purpleColor = lipgloss.Color("#5f5fff")
	footerColor = lipgloss.Color("#444444")
)

func (s Slot) String() string {
	if s == Empty {
		return "  "
	}

	si := int(s)
	sis := strconv.Itoa(si)

	if si >= 10 {
		return sis
	}

	return " " + sis
}

func viewSlot(s Slot) string {
	return lipgloss.NewStyle().
		Bold(true).
		Foreground(whiteColor).
		Render(s.String())
}

// creates a lipgloss table - which is the grid of the sliding puzzle
func createBoardGrid(g Game) *table.Table {
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

func ViewGame(g Game) string {
	bg := createBoardGrid(g)
	return bg.String()
}

func ViewWindow(m model) string {
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
		// Background(whiteColor).
		Render("q: quit • m: mix")

	window := border.Render(
		lipgloss.JoinVertical(
			lipgloss.Center,
			title,
			ViewGame(m.game),
			footer,
		),
	)

	return window
}
