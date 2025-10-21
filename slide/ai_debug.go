package slide

import (
	"strconv"

	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/lipgloss/table"
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

func (g Game) String() string {
	bg := table.New().
		Border(lipgloss.DoubleBorder()).
		BorderRow(true).
		StyleFunc(func(row, col int) lipgloss.Style {
			return lipgloss.NewStyle().Padding(0, 1)
		})

	for ri := range g.RowsNum {
		r := make([]string, 0, g.ColsNum)
		for ci := range g.ColsNum {
			slot := g.B[(g.ColsNum*(ri))+(ci)]
			r = append(r, slot.String())
		}
		bg.Row(r...)
	}
	return bg.String()
}

func (m BoardMovement) String() (s string) {
	switch m {
	case MoveUpToEmpty:
		s = "MoveUpToEmpty"
	case MoveDownToEmpty:
		s = "MoveDownToEmpty"
	case MoveLeftToEmpty:
		s = "MoveLeftToEmpty"
	case MoveRightToEmpty:
		s = "MoveRightToEmpty"
	}

	return
}
