package main

import (
	"os"
	"time"
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
	timer := NewTimer()

	state := runner.game.Start()
	round := 0
	for {
		select {
		case input := <-inputs:
			round++
			maxTime := startTimer(timer, round)
			if input.opponentAction != nil {
				state = runner.game.Play(state, input.opponentAction)
			}
			playerAction := runner.strategy.findAction(state, 1, maxTime)
			ValidateOutput(playerAction, input)
			state = runner.game.Play(state, playerAction)
			WriteDebug("State:", state)
			WriteOutput(playerAction)
			endTimer(timer, round)
			WriteDebug("Timer:", timer)
		case <-quit:
			return state
		}
	}
}

func startTimer(timer *Timer, round int) time.Time {
	if round == 1 {
		timer.startInit()
		return timer.initStart.Add(990 * time.Millisecond)
	} else {
		timer.startRound()
		return timer.roundStart.Add(95 * time.Millisecond)
	}
}

func endTimer(timer *Timer, round int) {
	if round == 1 {
		timer.endInit()
	} else {
		timer.endRound()
	}
}

func NewTicTacToeRunner(game Game[State, Action], strategy Strategy[State, Action]) TicTacToeRunner {
	return TicTacToeRunner{game: game, strategy: strategy}
}
