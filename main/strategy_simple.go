package main

import (
	"time"
)

type SimpleStrategy struct {
	game Game[State, Action]
}

func (strategy SimpleStrategy) findAction(state *State, player int, maxTime time.Time) *Action {
	if state == nil {
		panic("State cannot be nil")
	}
	availableActions := strategy.game.GetAvailableActions(state, player)
	for maxTime.After(time.Now()) {
		time.Sleep(5 * time.Millisecond)
	}
	if len(availableActions) == 0 {
		return nil
	}
	return availableActions[0]
}

func NewSimpleStrategy(game Game[State, Action]) Strategy[State, Action] {
	WriteDebug("Simple strategy")
	return SimpleStrategy{
		game: game,
	}
}
