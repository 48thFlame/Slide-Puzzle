package main

import (
	"fmt"
)

// func newModel() model {
// 	return model{count: 0}
// }

// type model struct {
// 	count int
// }

// func (m model) Init() tea.Cmd {
// 	return nil
// }

// func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
// 	switch msg := msg.(type) {
// 	case tea.KeyMsg:
// 		switch msg.String() {
// 		case "ctrl+c", "q":
// 			return m, tea.Quit
// 		case "a":
// 			m.count--

// 		case "s":
// 			m.count++
// 		}
// 	}

// 	return m, nil
// }

// func (m model) View() string {
// 	return strconv.Itoa(m.count)
// }

func main() {
	g := NewGame(4, 4)
	fmt.Printf("g: %v\n", g)
	fmt.Println("--")
	g.MoveOnBard(MoveDownToEmpty)
	fmt.Printf("g: %v\n", g)
	fmt.Println("--")
	g.MoveOnBard(MoveUpToEmpty)
	fmt.Printf("g: %v\n", g)
	fmt.Println("--")
	g.MoveOnBard(MoveRightToEmpty)
	fmt.Printf("g: %v\n", g)
	fmt.Println("--")
	g.MoveOnBard(MoveLeftToEmpty)
	fmt.Printf("g: %v\n", g)
	fmt.Println("--")

	// p := tea.NewProgram(newModel())
	// if _, err := p.Run(); err != nil {
	// 	fmt.Printf("Alas, there's been an error: %v", err)
	// 	os.Exit(1)
	// }
}
