package main

type SimpleStrategy struct {
	game Game[State, Action]
}

func (strategy SimpleStrategy) findAction(state *State, player int) *Action {
	if state == nil {
		panic("State cannot be nil")
	}
	availableActions := strategy.game.GetAvailableActions(state, player)
	if len(availableActions) == 0 {
		return nil
	}
	return availableActions[0]
}

func NewSimpleStrategy(game Game[State, Action]) Strategy[State, Action] {
	return SimpleStrategy{
		game: game,
	}
}
