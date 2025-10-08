package main

import (
	"strconv"
	"strings"
)

type Slot uint

const (
	Empty Slot = iota
)

func (s Slot) String() string {
	if s == Empty {
		return "  "
	}

	si := int(s)
	sis := strconv.Itoa(si)

	if si >= 10 {
		return sis
	}

	return " " + sis
}

func newBoard(rowsNum, colsNum int) (Board, int) {
	piecesNum := rowsNum * colsNum

	board := make(Board, 0, piecesNum)
	for i := range piecesNum - 1 {
		board = append(board, Slot(i+1))
	}
	board = append(board, Empty)

	return board, piecesNum
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
	b, size := newBoard(rowsNum, colsNum)
	return Game{
		RowsNum: rowsNum,
		ColsNum: colsNum,
		B:       b,
		emptyI:  size - 1,
	}
}

// Represent a Sliding-Puzzle game
type Game struct {
	RowsNum, ColsNum int
	B                Board
	emptyI           int
}

func (g Game) String() string {
	var builder strings.Builder

	for ri := range g.RowsNum {
		builder.WriteRune('\n')
		for ci := range g.ColsNum {
			slot := g.B[(g.ColsNum*(ri))+(ci)]

			builder.WriteString(slot.String())
			builder.WriteRune(' ')
		}
	}

	return builder.String()
}

type BoardMovement uint8

const (
	MoveUpToEmpty BoardMovement = iota
	MoveDownToEmpty
	MoveLeftToEmpty
	MoveRightToEmpty
)

func (g *Game) MoveOnBard(movement BoardMovement) {
	var emptyIChange int
	switch movement {
	case MoveUpToEmpty:
		// if I'm moving up to the empty, then I want to switch the empty down a row
		emptyIChange = g.ColsNum
	case MoveDownToEmpty:
		emptyIChange = -g.ColsNum
	case MoveLeftToEmpty:
		emptyIChange = 1
	case MoveRightToEmpty:
		emptyIChange = -1
	}

	iToSwitchWithEmpty := g.emptyI + emptyIChange

	if iToSwitchWithEmpty < 0 || iToSwitchWithEmpty >= g.RowsNum*g.ColsNum {
		return
	}

	// now switch the empty slot with the slot that should be in the empty now
	g.B[g.emptyI], g.B[iToSwitchWithEmpty] = g.B[iToSwitchWithEmpty], g.B[g.emptyI]
	g.emptyI = iToSwitchWithEmpty
}
