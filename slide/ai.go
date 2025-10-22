package slide

type AiSolMove struct {
	Move   BoardMovement
	NumOfM int
}

// using a pointer as a "maybe" type
func AiOutput(g Game) *AiSolMove {
	solution := BFSNotRecur(g)
	lenSol := len(solution)
	if lenSol == 0 {
		return nil
	} else {
		return &AiSolMove{Move: solution[0], NumOfM: lenSol}
	}
}
