package main

import (
	"os"
)

type TicTacToeRunner struct {
	game     Game[State, Action]
	strategy Strategy[State, Action]
}

func (runner TicTacToeRunner) runFromInputStream(inputStream *os.File, quit chan bool) *State {
	inputs := make(chan *InputData)
	go func() {
		for {
			inputs <- ReadInputData(inputStream)
		}
	}()
	return runner.run(inputs, quit)
}

func (runner TicTacToeRunner) run(inputs chan *InputData, quit chan bool) *State {
	state := runner.game.Start()
	round := 0
	for {
		select {
		case input := <-inputs:
			round++
			if input.opponentAction != nil {
				state = runner.game.Play(state, input.opponentAction)
			}
			playerAction := runner.strategy.findAction(state, 1)
			ValidateOutput(playerAction, input)
			state = runner.game.Play(state, playerAction)
			WriteDebug("State:", state)
			WriteOutput(playerAction)
		case <-quit:
			return state
		}
	}
}

func NewTicTacToeRunner(game Game[State, Action], strategy Strategy[State, Action]) TicTacToeRunner {
	return TicTacToeRunner{game: game, strategy: strategy}
}
