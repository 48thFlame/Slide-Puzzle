package main

import (
	"fmt"
	"os"
	"strconv"

	tea "github.com/charmbracelet/bubbletea"
)

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
