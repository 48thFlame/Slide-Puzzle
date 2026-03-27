package slide

// I found that bigger takes too much time and memory to brute force the solution
// 10 cells which is also (5*2)!/2 ~= 1.8 million positions seems still reasonable
func isSafeToBFS(g Game) bool {
	return g.RowsNum*g.ColsNum <= cellsLimitForBFS
}

type bfsNode struct {
	g       Game
	moveSeq []BoardMovement
}

func (node bfsNode) children() []bfsNode {
	return mapFunc(
		getLegalMoves(node.g),
		func(m BoardMovement) bfsNode {
			updG := node.g.copyOfGame()
			updG.MoveOnBard(m)

			// copy the moveSeq so its not appending
			// to the parent slice (and so will all the children)
			updMovSeq := make([]BoardMovement, len(node.moveSeq))
			copy(updMovSeq, node.moveSeq)
			return bfsNode{g: updG, moveSeq: append(updMovSeq, m)}
		})
}

func _doBFSRec(toVisit []bfsNode, seenAlready map[string]struct{}) bfsNode {
	// pop the first to visit
	checking := toVisit[0]

	// check if found solution
	if checking.g.Won() {
		return checking
	}

	// remove the item from the queue
	toVisit = toVisit[1:]

	bsv := boardToStringValue(checking.g.B)

	if _, exists := seenAlready[bsv]; !exists {
		// if did not previously check
		// mark as checked
		seenAlready[bsv] = struct{}{}

		// add children of node to options to visit
		toVisit = append(toVisit, checking.children()...)
	} else {
	}

	return _doBFSRec(toVisit, seenAlready)
}

func doBFSNotRecursive(g Game) bfsNode {
	toVisit := []bfsNode{{g: g, moveSeq: make([]BoardMovement, 0)}} // need to initialize the moveSeq so won't be `nil`
	seenAlready := make(map[string]struct{})

	for {
		// pop the first to visit
		checking := toVisit[0]

		// check if found solution
		if checking.g.Won() {
			return checking
		}

		// remove the item from the queue
		toVisit = toVisit[1:]

		bsv := boardToStringValue(checking.g.B)

		if _, exists := seenAlready[bsv]; !exists {
			// if did not previously check
			// mark as checked
			seenAlready[bsv] = struct{}{}

			// add children of node to options to visit
			toVisit = append(toVisit, checking.children()...)
		} else {
			continue
		}
	}
}
