package main

import (
	"os"
)

// main function,
// scan inputs and run the game loop,
// print played actions
func main() {
	mainFrom(os.Stdin)
}

func mainFrom(inputStream *os.File) {
	game := NewTicTacToeGame()
	state := game.Start()
	round := 0
	for {
		round++
		inputData := ReadInputData(inputStream)
		if inputData.opponentAction != nil {
			state = game.Play(state, inputData.opponentAction)
		}
		strategy := NewSimpleStrategy(inputData.availableCells)
		playerAction := strategy.findAction(state, 1)
		ValidateOutput(playerAction, inputData)
		state = game.Play(state, playerAction)
		WriteDebug("State:", state)
		WriteOutput(playerAction)
	}
}
