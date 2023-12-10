package main

import (
	"testing"
)

func TestMctsNodeNewWithNilValues(t *testing.T) {
	node := NewMCTSNode(nil, nil, nil, nil)
	if node == nil {
		t.Fatal("Node should node be nil")
	}
	if node.action != nil {
		t.Error("Action is expected nil")
	}
	if node.state != nil {
		t.Error("State is expected nil")
	}
	if node.nbPlays != 0 {
		t.Error("Nb plays is expected to be 0")
	}
	if node.nbWins != 0 {
		t.Error("Nb wins is expected to be 0")
	}
	if node.parent != nil {
		t.Error("Parent is expected nil")
	}
	if node.children == nil {
		t.Error("Children should not be nil")
	}
	if len(node.children) != 0 {
		t.Error("Children size should be 0")
	}
}

func TestMctsNodeNew(t *testing.T) {
	game := NewTicTacToeGame()
	rootState := NewState().SetHeight(3).SetWidth(3)
	possibleRootActions := game.GetAvailableActions(rootState, 1)
	action := possibleRootActions[0]
	state := game.Play(rootState, action)
	possibleActions := game.GetAvailableActions(state, 2)
	root := NewMCTSNode(nil, rootState, nil, possibleRootActions)
	node := NewMCTSNode(action, state, root, possibleActions)
	if node == nil {
		t.Fatal("Node should node be nil")
	}
	if node.action != action {
		t.Error("Bad action")
	}
	if node.state != state {
		t.Error("Bad state")
	}
	if node.nbPlays != 0 {
		t.Error("Nb plays is expected to be 0")
	}
	if node.nbWins != 0 {
		t.Error("Nb wins is expected to be 0")
	}
	if node.parent != root {
		t.Error("Bad parent")
	}
	if node.children == nil {
		t.Error("Children should not be nil")
	}
	if len(node.children) != len(possibleActions) {
		t.Error("Bad children size")
	}
}

func TestMctsNodeGetChild(t *testing.T) {
	possiblesActions := make([]*Action, 3)
	possiblesActions[0] = NewAction(1, 0, 0)
	possiblesActions[1] = NewAction(2, 0, 1)
	possiblesActions[2] = NewAction(1, 0, 2)
	action := possiblesActions[1]
	tree := NewMCTSNode(nil, nil, nil, possiblesActions)
	expectedChild := NewMCTSNode(action, nil, nil, nil)
	tree.children[action] = expectedChild
	dataSet := []struct {
		testName      string
		action        *Action
		expectedChild *MCTSNode
	}{
		{"Nil action", nil, nil},
		{"Action not found", NewAction(2, 2, 2), nil},
		{"Same action reference", action, expectedChild},
		{"Same action values", NewAction(2, 0, 1), expectedChild},
	}
	for _, data := range dataSet {
		t.Run(data.testName, func(t *testing.T) {
			child := tree.GetChild(data.action)
			if child != data.expectedChild {
				t.Error("Bad child result")
			}
		})
	}
}
