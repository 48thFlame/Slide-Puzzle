package slide_test

import (
	"fmt"
	"testing"

	"github.com/48thFlame/Slide-Puzzle/slide"
)

// `go test slide/ai_test.go -v`
func Test(t *testing.T) {
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
	g := slide.Game{
		RowsNum: 2,
		ColsNum: 3,
		B:       []slide.Slot{1, 0, 3, 4, 2, 5},
		EmptyI:  1,
	}
	// fmt.Println("g:")
	fmt.Println(g)

	// fmt.Print()
	// t.Log(slide.AiOutput(g))
	fmt.Println(slide.AiOutput(g))
	// g.MoveOnBard(slide.MoveRightToEmpty)
	// fmt.Println(g)
	// fmt.Println(slide.AiOutput(g))

}
