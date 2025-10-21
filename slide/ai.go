package slide

import "fmt"

func AiOutput(g Game) string {
	solution := BFSNotRecur(g)
	if len(solution) == 0 {
		return "Solved"
	} else {
		return fmt.Sprintln(solution[0])
	}
}
