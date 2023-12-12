package main

import (
	"testing"
	"time"
)

func TestMctsGetBestActionWithState(t *testing.T) {
	state := NewState().SetWidth(3).SetHeight(3)
	mcts := NewMCTS(NewTicTacToeGame())
	mcts.Search(state, time.Now().Add(70*time.Millisecond))
	for _, child := range mcts.tree["0-000000000"].children {
		t.Log("MCTS", child)
	}
	action := mcts.GetBestAction(state)
	if action == nil {
		t.Fatal("Action should not be nil")
	}
	if action.player != 1 {
		t.Errorf("Expected player 1 but was %d", action.player)
	}
	if action.x != 1 {
		t.Errorf("Expected x 1 but was %d", action.x)
	}
	if action.y != 1 {
		t.Errorf("Expected y 1 but was %d", action.y)
	}
}

func TestMctsGetBestPlayWithoutState(t *testing.T) {
	mcts := NewMCTS(nil)
	action := mcts.GetBestAction(nil)
	if action != nil {
		t.Fatal("Action should be nil")
	}
}
