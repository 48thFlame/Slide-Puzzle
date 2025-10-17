package slide

import (
	"fmt"
	"slices"
)

func getLegalMoves(g Game) []BoardMovement {
	allMoves := []BoardMovement{MoveUpToEmpty, MoveDownToEmpty, MoveLeftToEmpty, MoveRightToEmpty}
	return slices.DeleteFunc(
		allMoves,
		func(m BoardMovement) bool {
			fmt.Println(g)

			return g.legalMove(m)
		})
}

// import main "github.com/48thFlame/Slide-Puzzle"

func AiOutput(g Game) string {
	// fmt.
	// return fmt.Sprint(g)
	fmt.Println(g)

	fmt.Println(getLegalMoves(g))

	// fmt.Println("Hi test")
	return "Hello World!"
}
