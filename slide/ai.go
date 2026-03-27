package slide

const (
	numOfStepsToMixBoard = 2000
	cellsLimitForBFS     = 10 // see comment in isSafeToBFS()
)

type AiOutFlags uint8

const (
	Solved AiOutFlags = iota
	Unsolvable
	TooHardCantSolve
	SolMove
)

type AiOut struct {
	Move   BoardMovement
	NumOfM int
}

// using a pointer as a "maybe" type (this way can use nil as a value)
func AiOutput(g Game) (AiOutFlags, *AiOut) {
	if !Solvable(g) {
		return Unsolvable, nil
	}
	if !isSafeToBFS(g) {
		return TooHardCantSolve, nil
	}

	// sol := _doBFSRec([]bfsNode{{g: g}}, make(map[string]struct{}))
	solutionBFSNode := doBFSNotRecursive(g)
	solution := solutionBFSNode.moveSeq

	lenSol := len(solution)
	if lenSol == 0 {
		return Solved, nil
	} else {
		return SolMove, &AiOut{Move: solution[0], NumOfM: lenSol}
	}
}
