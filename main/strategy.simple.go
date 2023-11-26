package main

type SimpleStrategy struct{}

func (strategy SimpleStrategy) findAction(state *State) *Cell {
	if state == nil {
		panic("State cannot be nil")
	}
	return state.availableActions[0]
}

func NewSimpleStrategy() Strategy {
	return SimpleStrategy{}
}
