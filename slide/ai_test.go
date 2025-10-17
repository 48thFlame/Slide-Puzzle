package slide_test

import (
	"testing"

	"github.com/48thFlame/Slide-Puzzle/slide"
	// "github.com/48thFlame/Slide-Puzzle/slide"
)

func Test(t *testing.T) {
	g := slide.NewGame(4, 4)
	slide.AiOutput(g)
}
