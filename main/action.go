package main

import (
	"fmt"
)

type Action struct {
	player int
	x      int
	y      int
}

func (action *Action) String() string {
	if action == nil {
		return ""
	}
	return fmt.Sprintf("{player: %d, x: %d, y: %d}", action.player, action.x, action.y)
}

func NewAction(player int, x int, y int) *Action {
	return &Action{player: player, x: x, y: y}
}
