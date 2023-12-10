package main

// Monte Carlo Tree Search algorithm

import (
	"time"
)

type MCTS struct {
	exploreParam int
	game         Game[State, Action]
	tree         map[string]*Action
}

func (mcts *MCTS) search(state *State, maxTime time.Time) {
	if mcts == nil {
		return
	}
	mcts.tree[state.Hash()] = mcts.game.GetAvailableActions(state, 1)[0]
	for maxTime.After(time.Now()) {
		time.Sleep(5 * time.Millisecond)
	}
}

func (mcts *MCTS) getBestPlay(state *State) *Action {
	if mcts == nil || state == nil {
		return nil
	}
	return mcts.tree[state.Hash()]
}

func NewMCTS(game Game[State, Action]) *MCTS {
	return &MCTS{
		exploreParam: 2,
		game:         game,
		tree:         map[string]*Action{},
	}
}
