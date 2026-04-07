package slide

import "slices"

// solvable - checks if the given board position is solvable
// not all possible arrangements of a grid of numbers can be slid back to solved position
func Solvable(g Game) bool {
	// a solvable board needs to have all the pieces
	if !allPieces(g) {
		return false
	}

	// these 2 functions, cyclesNum and mhtDist
	// each take turns changing by one every switch on the board (every turn)
	// however mhtDist starts at 0 and cyclesNum 1
	// therefore one if them should always be odd and one even
	// if not - means pos is unsolvable
	// see https://www.lukelavalva.com/theoryofsliding
	md := mhtDist(g)
	noc, _ := numOfCycles(g)

	return (md+noc)%2 != 0
}

func allPieces(g Game) bool {
	solvedBoard := newBoard(g.RowsNum, g.ColsNum)

	if len(g.B) != len(solvedBoard) {
		return false
	}

	for _, piece := range solvedBoard {
		if !slices.Contains(g.B, piece) {
			return false
		}
	}

	return true
}

// mht dist - how far the empty slot is far from starting pos
func mhtDist(g Game) (dist int) {
	/*
		row = floor(i / cols_num)
		col = mod(i, cols_num)
	*/
	rowI := g.EmptyI / g.ColsNum
	colI := g.EmptyI % g.ColsNum

	rowsDist := g.RowsNum - 1 - rowI
	colsDist := g.ColsNum - 1 - colI

	return colsDist + rowsDist
}

// takes a starting index returns a slice of all indices of the permutation
func permutation(acc []int, startingI int, board Board) []int {
	nextI := int(board[startingI])
	if slices.Contains(acc, nextI) {
		return acc
	} else {
		return permutation(append(acc, nextI), nextI, board)
	}

}

func numOfCycles(g Game) (int, [][]int) {
	perms := make([][]int, 0)

	seenAlready := make(map[int]struct{})

	for i := range g.B {
		if _, exists := seenAlready[i]; !exists {
			p := permutation(make([]int, 0), i, g.B)

			for _, v := range p {
				seenAlready[v] = struct{}{}
			}

			perms = append(perms, p)
		}
	}

	return len(perms), perms
}
