package main

import (
	"os"
)

// main function,
// scan inputs and run the game loop,
// print played actions
func main() {
	runFromInputStream(os.Stdin, nil)
}
