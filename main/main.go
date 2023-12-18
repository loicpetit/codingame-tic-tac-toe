package main

import "os"

func main() {
	game := NewTicTacToeGame()
	strategy := NewStrategyFromArgs(os.Args, game)
	runner := NewTicTacToeRunner(os.Stdin, game, strategy)
	runner.Run(nil)
}
