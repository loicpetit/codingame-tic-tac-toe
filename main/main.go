package main

import (
	"os"
)

// main function,
// scan inputs and run the game loop,
// print played actions
func main() {
	mainFrom(os.Stdin, nil)
}

func mainFrom(inputStream *os.File, quit chan bool) *State {
	game := NewTicTacToeGame()
	state := game.Start()
	inputs := make(chan *InputData)
	round := 0
	var read = func() {
		inputs <- ReadInputData(inputStream)
	}
	go read()
	for {
		select {
		case input := <-inputs:
			round++
			if input.opponentAction != nil {
				state = game.Play(state, input.opponentAction)
			}
			strategy := NewSimpleStrategy(input.availableCells)
			playerAction := strategy.findAction(state, 1)
			ValidateOutput(playerAction, input)
			state = game.Play(state, playerAction)
			WriteDebug("State:", state)
			WriteOutput(playerAction)
			go read()
		case <-quit:
			return state
		}
	}
}
