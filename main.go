package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/48thFlame/Slide-Puzzle/slide"
	tea "github.com/charmbracelet/bubbletea"
)

const (
	defaultRowsNum = 4
)

func isFlagPassed(name string) bool {
	found := false
	flag.Visit(func(f *flag.Flag) {
		if f.Name == name {
			found = true
		}
	})
	return found
}

func getGameFromArgs() (slide.Game, error) {
	// the first arg is the program name, so we ignore it
	numOfArgs := len(os.Args) - 1

	if numOfArgs == 0 {
		return slide.NewGame(defaultRowsNum, defaultRowsNum), nil
	}

	rowsNum := flag.Int("rn", defaultRowsNum, "Number of rows in the puzzle")
	colsNum := flag.Int("cn", defaultRowsNum, "Number of cols in the puzzle")

	boardFlag := flag.String("board", "", "`board` should be a space separated string of the values for the board, for example: \"1 2 3 4 5 6 7 8 0\", make sure to set `cn` and `rn`")

	flag.Parse()

	if isFlagPassed("board") {
		// the user set a manual board, so we need to validate it and parse it

		// we need the flags to be set
		if !isFlagPassed("rn") || !isFlagPassed("cn") {
			return slide.Game{},
				fmt.Errorf("if you want to set the board, you must also set the number of rows and cols with `rn` and `cn` flags respectively")
		}

		boardSplitInput := strings.Split(*boardFlag, " ")

		// the length of the board should be good
		if len(boardSplitInput) != (*rowsNum * *colsNum) {
			return slide.Game{},
				fmt.Errorf("the number of values in the board should be equal to rowsNum * colsNum, but board is size %v, and rowsNum * colsNum is %v", len(boardSplitInput), (*rowsNum * *colsNum))
		}

		b := make(slide.Board, 0, numOfArgs)
		emptyI := -1
		for i, v := range boardSplitInput {
			n, err := strconv.Atoi(v)
			if err != nil {
				return slide.Game{},
					fmt.Errorf("That's NaN! '%v', please provide a valid number for all cells in the board", v)
			}
			if n == 0 {
				emptyI = i
			}
			b = append(b, slide.Slot(n))
		}

		if emptyI == -1 {
			// meaning the board didn't contain an empty slot, because the -1 never changed
			return slide.Game{},
				fmt.Errorf("the board should contain an empty slot, which is represented by 0, but it wasn't found in the provided board")
		}

		return slide.NewGameManual(*rowsNum, *colsNum, b, emptyI), nil
	}

	return slide.NewGame(*rowsNum, *colsNum), nil
}

func main() {
	// _test()
	// return
	// starting from the first user provided arg
	g, err := getGameFromArgs()
	if err != nil {
		fmt.Printf("Alas, there's been an error: %v\n", err)
		os.Exit(2)
	}

	p := tea.NewProgram(newModel(g))
	_, err = p.Run()
	if err != nil {
		fmt.Printf("Alas, there's been an error: %v\n", err)
		os.Exit(1)
	}
}
