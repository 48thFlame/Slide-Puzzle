package slide

// Breadth First Search Recursive
func BFSRecur(g Game) []BoardMovement {
	if isSafeToBFS(g) {
		sol := doBFS([]bfsNode{{g: g}}, make(map[string]struct{}))
		return sol.moveSeq
	} else {
		return nil
	}
}

// Breadth First Search Not Recursive (sadly no tail call optimization in go)
func BFSNotRecur(g Game) []BoardMovement {
	if isSafeToBFS(g) {
		sol := doBFSNotRecursive(g)
		return sol.moveSeq
	} else {
		return nil
	}
}

func isSafeToBFS(g Game) bool {
	return g.RowsNum*g.ColsNum <= 10
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

func doBFS(toVisit []bfsNode, seenAlready map[string]struct{}) bfsNode {
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

	return doBFS(toVisit, seenAlready)
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
