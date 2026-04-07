package slide

import "slices"

type searchNode struct {
	g       Game
	moveSeq []BoardMovement
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

func (node searchNode) children() []searchNode {
	return mapFunc(
		getLegalMoves(node.g),
		func(m BoardMovement) searchNode {
			updG := node.g.copyOfGame()
			updG.MoveOnBard(m)

			// copy the moveSeq so its not appending
			// to the parent slice (and so will all the children)
			updMovSeq := make([]BoardMovement, len(node.moveSeq))
			copy(updMovSeq, node.moveSeq)
			return searchNode{g: updG, moveSeq: append(updMovSeq, m)}
		})
}
