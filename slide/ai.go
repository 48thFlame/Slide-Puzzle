package slide

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

// using a pointer as a "maybe" type
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
