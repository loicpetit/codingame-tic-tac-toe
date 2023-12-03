package main

type SimpleStrategy struct {
	availableCells []*Cell
}

func (strategy SimpleStrategy) findAction(state *State, player int) *Action {
	if state == nil {
		panic("State cannot be nil")
	}
	return NewAction(player, strategy.availableCells[0].x, strategy.availableCells[0].y)
}

func NewSimpleStrategy(availableCells []*Cell) Strategy {
	return SimpleStrategy{
		availableCells: availableCells,
	}
}
