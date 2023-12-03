package main

import (
	"os"
)

// main function,
// scan inputs and run the game loop,
// print played actions
func main() {
	game := NewTicTacToeGame()
	strategy := NewSimpleStrategy(game)
	runner := NewTicTacToeRunner(game, strategy)
	runner.runFromInputStream(os.Stdin, nil)
}
