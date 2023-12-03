package main

import (
	"os"
)

func runFromInputStream(inputStream *os.File, quit chan bool) *State {
	inputs := make(chan *InputData)
	go func() {
		for {
			inputs <- ReadInputData(inputStream)
		}
	}()
	return run(inputs, quit)
}

func run(inputs chan *InputData, quit chan bool) *State {
	game := NewTicTacToeGame()
	state := game.Start()
	round := 0
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
		case <-quit:
			return state
		}
	}
}
