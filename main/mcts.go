package main

// Monte Carlo Tree Search algorithm

import (
	"math/rand"
	"time"
)

type MCTS struct {
	exploreParam int
	game         Game[State, Action]
	tree         map[string]*MCTSNode
}

func (mcts *MCTS) makeNode(state *State) *MCTSNode {
	if mcts == nil || state == nil {
		return nil
	}
	hash := state.Hash()
	if mcts.tree[hash] != nil {
		return mcts.tree[hash]
	}
	mcts.tree[hash] = NewMCTSNode(nil, state, nil, mcts.game.GetAvailableActions(state, 1))
	return mcts.tree[hash]
}

func (mcts *MCTS) selectNode(root *MCTSNode) *MCTSNode {
	node := root
	for mcts.isNodeFullyExpanded(node) && !mcts.isNodeLeaf(node) {
		// todo
	}
	return node
}

func (mcts *MCTS) expandNode(node *MCTSNode) *MCTSNode {
	unexploredActions := node.GetUnexploredActions()
	randomIndex := rand.Intn(len(unexploredActions))
	action := unexploredActions[randomIndex]
	childState := mcts.game.Play(node.state, action)
	childPossibleAction := mcts.game.GetAvailableActions(childState, mcts.game.GetNextPlayer(action))
	childNode := NewMCTSNode(action, childState, node, childPossibleAction)
	node.AddChild(childNode)
	return childNode
}

func (mcts *MCTS) simulate(node *MCTSNode) int {
	// todo
	return 0
}

func (mcts *MCTS) backPropagateResult(node *MCTSNode, winner int) {
	// todo
}

func (mcts *MCTS) isNodeFullyExpanded(node *MCTSNode) bool {
	return len(node.GetUnexploredActions()) == 0
}

func (mcts *MCTS) isNodeLeaf(node *MCTSNode) bool {
	return len(node.GetPossibleActions()) == 0
}

func (mcts *MCTS) Search(state *State, maxTime time.Time) {
	if mcts == nil {
		return
	}
	root := mcts.makeNode(state)
	node := mcts.selectNode(root)
	winner := mcts.game.Winner(node.state)
	if !node.IsLeaf() && winner == 0 {
		node = mcts.expandNode(node)
		winner = mcts.simulate(node)
	}
	mcts.backPropagateResult(node, winner)
	for maxTime.After(time.Now()) {
		time.Sleep(5 * time.Millisecond)
	}
}

func (mcts *MCTS) GetBestAction(state *State) *Action {
	if mcts == nil || state == nil {
		return nil
	}
	node := mcts.makeNode(state)
	var action *Action
	nbPlays := -1
	for _, child := range node.children {
		if child != nil && child.nbPlays > nbPlays {
			action = child.action
		}
	}
	return action
}

func NewMCTS(game Game[State, Action]) *MCTS {
	return &MCTS{
		exploreParam: 2,
		game:         game,
		tree:         map[string]*MCTSNode{},
	}
}
