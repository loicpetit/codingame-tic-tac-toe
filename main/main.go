package main

import (
	"fmt"
	"os"
)

// main function,
// scan inputs and run the game loop
func main() {

	inputStream := os.Stdin

	for {
		var opponentRow, opponentCol int
		fmt.Fscan(inputStream, &opponentRow, &opponentCol)

		var validActionCount int
		fmt.Fscan(inputStream, &validActionCount)

		for i := 0; i < validActionCount; i++ {
			var row, col int
			fmt.Fscan(inputStream, &row, &col)
		}

		// fmt.Fprintln(os.Stderr, "Debug messages...")
		fmt.Println("0 0") // Write action to stdout
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
