package main

type MCTSNode struct {
	action  *Action //action to get to that state
	state   *State
	nbPlays int
	nbWins  int
	parent  *MCTSNode
	// possible actions from that state
	// if node is nil it is still not expanded
	children map[*Action]*MCTSNode
}

func (node *MCTSNode) AddChild(child *MCTSNode) {
	if node == nil || child == nil {
		return
	}
	node.children[child.action] = child
}

func (node *MCTSNode) GetChild(action *Action) *MCTSNode {
	if node == nil || action == nil {
		return nil
	}
	for key, child := range node.children {
		if key == action || (key.player == action.player &&
			key.x == action.x &&
			key.y == action.y) {
			return child
		}
	}
	return nil
}

func (node *MCTSNode) GetPossibleActions() []*Action {
	actions := make([]*Action, 0)
	for action, _ := range node.children {
		actions = append(actions, action)
	}
	return actions
}

func (node *MCTSNode) GetUnexploredActions() []*Action {
	actions := make([]*Action, 0)
	for action, child := range node.children {
		if child == nil {
			actions = append(actions, action)
		}
	}
	return actions
}

func (node *MCTSNode) IsLeaf() bool {
	return len(node.children) == 0
}

func NewMCTSNode(
	action *Action,
	state *State,
	parent *MCTSNode,
	possibleActions []*Action,
) *MCTSNode {
	children := make(map[*Action]*MCTSNode)
	for _, action := range possibleActions {
		children[action] = nil
	}
	return &MCTSNode{
		action:   action,
		state:    state,
		nbPlays:  0,
		nbWins:   0,
		parent:   parent,
		children: children,
	}
}
