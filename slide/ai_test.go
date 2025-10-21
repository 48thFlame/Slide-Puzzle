package slide_test

import (
	"fmt"
	"testing"

	"github.com/48thFlame/Slide-Puzzle/slide"
)

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
	g := slide.NewGame(3, 3)
	// g.Mix()
	// g := slide.Game{
	// 	RowsNum: 2,
	// 	ColsNum: 3,
	// 	B:       []slide.Slot{1, 0, 3, 4, 2, 5},
	// 	EmptyI:  2,
	// }
	// fmt.Println("g:")
	// fmt.Println(g)

	fmt.Print(slide.AiOutput(g))

}
