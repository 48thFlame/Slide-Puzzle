package main

import (
	"fmt"

	"github.com/48thFlame/Slide-Puzzle/slide"
)

func _test() {
	// g := slide.NewGame(3, 3)
	// g := slide.Game{
	// 	RowsNum: 3,
	// 	ColsNum: 3,
	// 	B:       []slide.Slot{1, 2, 3, 4, 6, 5, 7, 0, 8},
	// 	EmptyI:  8,
	// }
	// g.Mix()
	g := slide.NewGame(5, 5)
	g.MoveOnBard(slide.MoveDownToEmpty)
	g.MoveOnBard(slide.MoveDownToEmpty)
	g.MoveOnBard(slide.MoveRightToEmpty)
	g.MoveOnBard(slide.MoveRightToEmpty)
	g.MoveOnBard(slide.MoveDownToEmpty)
	// g.MoveOnBard(slide.MoveDownToEmpty)

	fmt.Println(g)
	a, b := slide.PuzzleSize(g)
	fmt.Println(a, b)

	// g := slide.Game{
	// 	RowsNum: 2,
	// 	ColsNum: 3,
	// 	B:       []slide.Slot{1, 3, 0, 4, 2, 5},
	// 	EmptyI:  2,
	// }
	// g := slide.NewGame(3, 3)
	// g.Mix()
	// g := slide.Game{
	// 	RowsNum: 3,
	// 	ColsNum: 3,
	// 	B:       []slide.Slot{1, 2, 3, 4, 5, 7, 7, 7, 0},
	// 	EmptyI:  8,
	// }
	// g := slide.Game{
	// 	RowsNum: 2,
	// 	ColsNum: 2,
	// 	B:       []slide.Slot{2, 1, 3, 0},
	// 	EmptyI:  3,
	// }
	// fmt.Println(g)
	// fmt.Println("Solvable?", slide.Solvable(g))

	// fmt.Print("AiOutput: ")
	// fmt.Print(slide.AiOutput(g))
	// fmt.Println("")
	// fmt.Println("mhtDist:", slide.MhtDist(g))
	// nc, cycles := slide.NumOfCycles(g)
	// fmt.Println("perms:", cycles)
	// fmt.Println("numOfPerm:", nc)

}
