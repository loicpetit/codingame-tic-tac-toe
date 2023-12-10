package main

// Monte Carlo Tree Search algorithm

import (
	"time"
)

type MCTS struct {
	availableActions []*Action
	game             Game[State, Action]
}

func (mcts *MCTS) search(state *State, maxTime time.Time) {
	mcts.availableActions = mcts.game.GetAvailableActions(state, 1)
	for maxTime.After(time.Now()) {
		time.Sleep(5 * time.Millisecond)
	}
}

func (mcts *MCTS) getBestPlay(state *State) *Action {
	if len(mcts.availableActions) == 0 {
		return nil
	}
	return mcts.availableActions[0]
}

func NewMCTS(game Game[State, Action]) *MCTS {
	return &MCTS{
		game: game,
	}
}
