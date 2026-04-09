package slide

// rbr - Row By Row

import (
	"slices"
)

func PuzzleSize(g Game) (rowMixedNum, colMixedNum int) {
	solvedBoard := newBoard(g.RowsNum, g.ColsNum)
	var solvedRowNum, solvedColNum int

	for rowI := 0; rowI < g.RowsNum; rowI++ {
		// go through each row and compare to the solved
		row := g.B[rowI*g.ColsNum : (rowI+1)*g.ColsNum]
		solvedRow := solvedBoard[rowI*g.ColsNum : (rowI+1)*g.ColsNum]
		// fmt.Printf("row: %v\n", row)
		// fmt.Printf("solvedRow: %v\n", solvedRow)

		if !slices.Equal(row, solvedRow) {
			// starting from the top, as soon as there's a bad row -
			// all the ret are by definition mixed, so we jut break
			break
		}
		solvedRowNum++
	}

	for colI := 0; colI < g.ColsNum; colI++ {
		// building the column is slightly more complex than the rows
		col := make([]Slot, 0, g.RowsNum)
		solvedCol := make([]Slot, 0, g.RowsNum)

		for rowI := 0; rowI < g.RowsNum; rowI++ {
			col = append(col, g.B[rowI*g.ColsNum+colI])
			solvedCol = append(solvedCol, solvedBoard[rowI*g.ColsNum+colI])
		}
		// fmt.Printf("col: %v\n", col)
		// fmt.Printf("solvedCol: %v\n", solvedCol)

		if !slices.Equal(col, solvedCol) {
			break
		}
		solvedColNum++
	}

	return g.RowsNum - solvedRowNum, g.ColsNum - solvedColNum
}

// func doBFSNotRecursive(g Game) searchNode {
// 	toVisit := []searchNode{{g: g, moveSeq: make([]BoardMovement, 0)}} // need to initialize the moveSeq so won't be `nil`
// 	seenAlready := make(map[string]struct{})

// 	for {
// 		// pop the first to visit
// 		checking := toVisit[0]

// 		// check if found solution
// 		if checking.g.Won() {
// 			return checking
// 		}

// 		// remove the item from the queue
// 		toVisit = toVisit[1:]

// 		bsv := boardToStringValue(checking.g.B)

// 		if _, exists := seenAlready[bsv]; !exists {
// 			// if did not previously check
// 			// mark as checked
// 			seenAlready[bsv] = struct{}{}

// 			// add children of node to options to visit
// 			toVisit = append(toVisit, checking.children()...)
// 		} else {
// 			continue
// 		}
// 	}
// }
