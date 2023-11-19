package main

import (
	"fmt"
	"os"
)

// main function,
// scan inputs and run the game loop,
// print played actions
func main() {
	stateBuilder := NewStateBuilder(os.Stdin)
	state := stateBuilder.buildInitState()
	round := 0
	for {
		round++
		state := stateBuilder.buildTurnState(state)

		// fmt.Fprintln(os.Stderr, "Debug messages...")
		cellToPlay := state.availableActions[0]
		fmt.Println(cellToPlay.row, cellToPlay.column) // Write action to stdout
	}
}

/*
	Original main

	package main

	import "fmt"

	func main() {
		for {
			var opponentRow, opponentCol int
			fmt.Scan(&opponentRow, &opponentCol)

			var validActionCount int
			fmt.Scan(&validActionCount)

			for i := 0; i < validActionCount; i++ {
				var row, col int
				fmt.Scan(&row, &col)
			}

			// fmt.Fprintln(os.Stderr, "Debug messages...")
			fmt.Println("0 0")// Write action to stdout
		}
	}
*/
