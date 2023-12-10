package main

import (
	"time"
)

type MctsStrategy struct {
	mcts *MCTS
}

func (strategy MctsStrategy) findAction(state *State, player int, maxTime time.Time) *Action {
	if state == nil {
		panic("State cannot be nil")
	}
	strategy.mcts.Search(state, maxTime)
	return strategy.mcts.GetBestAction(state)
}

func NewMctsStrategy(game Game[State, Action]) Strategy[State, Action] {
	WriteDebug("MCTS strategy")
	return MctsStrategy{
		mcts: NewMCTS(game),
	}
}
