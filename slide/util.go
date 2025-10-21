package slide

import (
	"slices"
	"strconv"
	"strings"
)

func mapFunc[A, B any](a []A, f func(A) B) []B {
	b := make([]B, 0, len(a))
	for _, v := range a {
		b = append(b, f(v))
	}
	return b
}

func (g Game) copyOfGame() (ng Game) {
	ng = g

	// need to just really copy the board, because nothing else is a "pointer type"
	ng.B = make([]Slot, len(g.B))
	copy(ng.B, g.B)
	return
}

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

	return true
}

func getLegalMoves(g Game) []BoardMovement {
	allMoves := []BoardMovement{
		MoveUpToEmpty,
		MoveDownToEmpty,
		MoveLeftToEmpty,
		MoveRightToEmpty}

	return slices.DeleteFunc(
		allMoves,
		func(m BoardMovement) bool {
			return !g.legalMove(m)
		})
}

// used to store boards in lookup tables (like in duplication checking in BFS)
func boardToStringValue(b Board) string {
	sb := strings.Builder{}

	for _, slot := range b {
		sb.WriteString(strconv.Itoa(int(slot)))
	}

	return sb.String()
}
