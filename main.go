package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
)

func newModel(rowsNum, colsNum int) model {
	return model{game: NewGame(rowsNum, colsNum)}
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
	var res strings.Builder

	res.WriteString(m.game.String())
	res.WriteString(fmt.Sprintf("Did win: %v\n", m.game.Won()))

	return res.String()
}

func main() {
	var rowsNum, colsNum int

	nOfArgs := len(os.Args)
	switch nOfArgs {
	case 1:
		n := 4
		rowsNum, colsNum = n, n
	case 2:
		n, err := strconv.Atoi(os.Args[1])
		if err != nil {
			panic(fmt.Errorf("That's NaN! '%v', please provide (or don't) a number(s) for num of rows/cols", os.Args[1]))
		}
		rowsNum, colsNum = n, n
	case 3:
		rn, err := strconv.Atoi(os.Args[1])
		if err != nil {
			panic(fmt.Errorf("That's NaN! '%v', please provide (or don't) a number for num of rows", os.Args[1]))
		}

		cn, err := strconv.Atoi(os.Args[2])
		if err != nil {
			panic(fmt.Errorf("That's NaN! '%v', please provide (or don't) a number for num of cols", os.Args[2]))
		}

		rowsNum, colsNum = rn, cn
	}

	p := tea.NewProgram(newModel(rowsNum, colsNum))
	_, err := p.Run()
	if err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}
}
