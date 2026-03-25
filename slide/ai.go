package slide

const (
	numOfStepsToMixBoard = 2000
	cellsLimitForBFS     = 10
)

type AiOutFlags uint8

const (
	Solved AiOutFlags = iota
	CantSolve
	SolMove
)

type AiOut struct {
	Move   BoardMovement
	NumOfM int
}

// func AiOutput(g Game) int
// 	return misplacedTiles(g){
// }

// using a pointer as a "maybe" type (this way can use nil as a value)
func AiOutput(g Game) (AiOutFlags, *AiOut) {
	solution := BFSNotRecur(g)

	if solution == nil {
		// means board is too big
		return CantSolve, nil
	}

	lenSol := len(solution)
	if lenSol == 0 {
		return Solved, nil
	} else {
		return SolMove, &AiOut{Move: solution[0], NumOfM: lenSol}
	}
}
