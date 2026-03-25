package main

import (
	"fmt"
	"os"
	"strconv"

	tea "github.com/charmbracelet/bubbletea"
)

func getRowsColsNumArgs(args []string) (rowsNum, colsNum int, e error) {
	nOfArgs := len(args)

	switch nOfArgs {

	case 0:
		n := 4
		rowsNum, colsNum = n, n

	case 1:
		n, err := strconv.Atoi(os.Args[1])
		if err != nil {
			e = fmt.Errorf("That's NaN! '%v', please provide (or don't) a number(s) for num of rows/cols", os.Args[1])
			return
		}
		rowsNum, colsNum = n, n

	case 2:
		rn, err := strconv.Atoi(os.Args[1])
		if err != nil {
			e = fmt.Errorf("That's NaN! '%v', please provide (or don't) a number for num of rows", os.Args[1])
			return
		}

		cn, err := strconv.Atoi(os.Args[2])
		if err != nil {
			e = fmt.Errorf("That's NaN! '%v', please provide (or don't) a number for num of cols", os.Args[2])
			return
		}

		rowsNum, colsNum = rn, cn
	}
	return
}

func main() {
	// starting from the first user provided arg
	rowsNum, colsNum, err := getRowsColsNumArgs(os.Args[1:])
	if err != nil {
		panic(err)
	}

	p := tea.NewProgram(newModel(rowsNum, colsNum))
	_, err = p.Run()
	if err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}
}
