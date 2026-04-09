package slide

// I found that bigger takes too much time and memory to brute force the solution
// 10 cells which is also (5*2)!/2 ~= 1.8 million positions seems still reasonable
func isSafeToBFS(g Game) bool {
	if g.RowsNum*g.ColsNum <= boardCellsNumLimitForBFS {
		return true
	} else {
		rm, cm := PuzzleSize(g)
		return rm*cm <= cellsMixedLimitForBFS
	}
}

func _doBFSRec(toVisit []searchNode, seenAlready map[string]struct{}) searchNode {
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

func doBFSNotRecursive(g Game) searchNode {
	toVisit := []searchNode{{g: g, moveSeq: make([]BoardMovement, 0)}} // need to initialize the moveSeq so won't be `nil`
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
