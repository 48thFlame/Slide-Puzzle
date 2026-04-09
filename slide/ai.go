package slide

const (
	numOfStepsToMixBoard     = 2000
	boardCellsNumLimitForBFS = 10 // see comment in isSafeToBFS()
	// even if the board is too big but the mixed area is small
	// - it's probably still safe to BFS
	mixedPercentCellsLimitForBFS = 0.5
	cellsMixedLimitForBFS        = boardCellsNumLimitForBFS * mixedPercentCellsLimitForBFS
)

type AiOutFlags uint8

const (
	Solved AiOutFlags = iota
	Unsolvable
	TooHardCantSolve
	BfsSol
)

type BfsAiOut struct {
	Move   BoardMovement
	NumOfM int
}

// using a pointer as a "maybe" type (this way can use nil as a value)
func AiOutput(g Game) (AiOutFlags, *BfsAiOut, [2]int) {
	if !Solvable(g) {
		return Unsolvable, nil, [2]int{-1, -1}
	}

	puzSizeR, puzSizeC := PuzzleSize(g)
	if isSafeToBFS(g) {
		// sol := _doBFSRec([]bfsNode{{g: g}}, make(map[string]struct{}))
		solutionBFSNode := doBFSNotRecursive(g)
		solution := solutionBFSNode.moveSeq
		lenSol := len(solution)

		if lenSol == 0 {
			return Solved, nil, [2]int{0, 0}

		} else {
			return BfsSol,
				&BfsAiOut{Move: solution[0], NumOfM: lenSol},
				[2]int{puzSizeR, puzSizeC}
		}
	}

	return TooHardCantSolve, nil, [2]int{puzSizeR, puzSizeC}
}
