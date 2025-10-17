package slide

import (
	"math/rand"
	"slices"
)

type Slot uint

const (
	Empty Slot = iota
)

func newBoard(rowsNum, colsNum int) Board {
	if rowsNum > 10 || colsNum > 10 {
		panic("board works with up to `10 * 10`, not more")
	}

	piecesNum := rowsNum * colsNum

	board := make(Board, 0, piecesNum)
	for i := range piecesNum - 1 {
		board = append(board, Slot(i+1))
	}
	board = append(board, Empty)

	return board
}

/*
Board is a 2-d arrays, that's represented in a 1-d array.
Given that 0 is the top left corner and going to higher index means right/down,
these are true:

i = (cols_num * row) + col
col = mod(i, cols_num)
row = floor(i / cols_num)

Given that we are in a spot, and want to check the slots of moving, these are true.

cn = colsNum = BoardSideSize

-cn-1 -cn -cn+1

-1    0    1

cn-1   cn  cn+1
*/
type Board = []Slot

func NewGame(rowsNum, colsNum int) Game {
	game := Game{
		RowsNum: rowsNum,
		ColsNum: colsNum,
		B:       newBoard(rowsNum, colsNum),
		EmptyI:  rowsNum*colsNum - 1,
	}
	return game
}

// Represent a Sliding-Puzzle game
type Game struct {
	RowsNum, ColsNum int
	B                Board
	EmptyI           int
}

type BoardMovement uint8

const (
	MoveUpToEmpty BoardMovement = iota
	MoveDownToEmpty
	MoveLeftToEmpty
	MoveRightToEmpty
)

// TODO: make it combined with MoveOnBoard
func (g Game) legalMove(move BoardMovement) bool {
	var emptyIChange int
	switch move {
	case MoveUpToEmpty:
		// if I'm moving up to the empty,
		// then I want to switch the empty down a row
		// etc
		emptyIChange = g.ColsNum
	case MoveDownToEmpty:
		emptyIChange = -g.ColsNum
	case MoveLeftToEmpty:
		emptyIChange = 1
	case MoveRightToEmpty:
		emptyIChange = -1
	}

	iToSwitchWithEmpty := g.EmptyI + emptyIChange

	// row = floor(i / cols_num)
	emptyRow := g.EmptyI / g.ColsNum
	toSwitchRow := iToSwitchWithEmpty / g.ColsNum

	movingHorizontally :=
		move == MoveLeftToEmpty ||
			move == MoveRightToEmpty

	if emptyRow != toSwitchRow && movingHorizontally {
		// if not in same row,
		// means moved right/left on an edge and shifted over to next row
		// which should not be possible
		return false
	}

	if iToSwitchWithEmpty < 0 || iToSwitchWithEmpty >= g.RowsNum*g.ColsNum {
		// if out of bounds
		return false
	}

	// // now switch the empty slot with the slot that should be in the empty now
	// g.B[g.EmptyI], g.B[iToSwitchWithEmpty] = g.B[iToSwitchWithEmpty], g.B[g.EmptyI]
	// g.EmptyI = iToSwitchWithEmpty

	return true
}

// applies the move, return bool whether was successful (was the move legal)
func (g *Game) MoveOnBard(movement BoardMovement) {
	var emptyIChange int
	switch movement {
	case MoveUpToEmpty:
		// if I'm moving up to the empty,
		// then I want to switch the empty down a row
		// etc
		emptyIChange = g.ColsNum
	case MoveDownToEmpty:
		emptyIChange = -g.ColsNum
	case MoveLeftToEmpty:
		emptyIChange = 1
	case MoveRightToEmpty:
		emptyIChange = -1
	}

	iToSwitchWithEmpty := g.EmptyI + emptyIChange

	// row = floor(i / cols_num)
	emptyRow := g.EmptyI / g.ColsNum
	toSwitchRow := iToSwitchWithEmpty / g.ColsNum

	movingHorizontally :=
		movement == MoveLeftToEmpty ||
			movement == MoveRightToEmpty

	if emptyRow != toSwitchRow && movingHorizontally {
		// if not in same row,
		// means moved right/left on an edge and shifted over to next row
		// which should not be possible
		return
	}

	if iToSwitchWithEmpty < 0 || iToSwitchWithEmpty >= g.RowsNum*g.ColsNum {
		// if out of bounds
		return
	}

	// now switch the empty slot with the slot that should be in the empty now
	g.B[g.EmptyI], g.B[iToSwitchWithEmpty] = g.B[iToSwitchWithEmpty], g.B[g.EmptyI]
	g.EmptyI = iToSwitchWithEmpty
}

func (g *Game) Mix() {
	// ? is there a better way to mix a board?
	stepsNum := 1111 // seems to be enough even for a 10*10
	moveOpts := [...]BoardMovement{MoveUpToEmpty, MoveDownToEmpty, MoveLeftToEmpty, MoveRightToEmpty}
	moveOptsNum := len(moveOpts)

	for range stepsNum {
		i := rand.Intn(moveOptsNum)
		g.MoveOnBard(moveOpts[i])
	}
}

func (g Game) Won() bool {
	return slices.Equal(g.B, newBoard(g.RowsNum, g.ColsNum))
}
