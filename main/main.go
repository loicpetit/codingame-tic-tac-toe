package main

import (
	"os"
)

// main function,
// scan inputs and run the game loop,
// print played actions
func main() {
	game := NewTicTacToeGame()
	strategy := NewStrategyFromArgs(os.Args, game)
	runner := NewTicTacToeRunner(game, strategy)
	runner.runFromInputStream(os.Stdin, nil)
}
