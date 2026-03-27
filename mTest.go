package main

import (
	"fmt"

	"github.com/48thFlame/Slide-Puzzle/slide"
)

func _test() {
	// g := slide.NewGame(2, 3)
	// g := slide.Game{
	// 	RowsNum: 2,
	// 	ColsNum: 3,
	// 	B:       []slide.Slot{1, 2, 3, 4, 5, 0},
	// 	EmptyI:  5,
	// }

	// g := slide.Game{
	// 	RowsNum: 2,
	// 	ColsNum: 3,
	// 	B:       []slide.Slot{1, 3, 0, 4, 2, 5},
	// 	EmptyI:  2,
	// }
	// g := slide.NewGame(3, 3)
	// g.Mix()
	// g := slide.Game{
	// 	RowsNum: 2,
	// 	ColsNum: 3,
	// 	B:       []slide.Slot{1, 0, 3, 4, 2, 5},
	// 	EmptyI:  1,
	// }
	g := slide.Game{
		RowsNum: 2,
		ColsNum: 2,
		B:       []slide.Slot{2, 1, 3, 0},
		EmptyI:  3,
	}
	fmt.Println(g)
	fmt.Println("Solvable?", slide.Solvable(g))

	// fmt.Print()
	// t.Log(slide.AiOutput(g))
	fmt.Println(slide.AiOutput(g))
	// fmt.Println("mhtDist", slide.MhtDist(g))
	// nc, cycles := slide.NumOfCycles(g)
	// fmt.Println("perms", cycles)
	// fmt.Println("numOfPerm", nc)
	// g.MoveOnBard(slide.MoveRightToEmpty)
	// fmt.Println(g)
	// fmt.Println(slide.AiOutput(g))

}
