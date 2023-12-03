package main

import "fmt"

type State struct {
	lastPlayer int
	grid       *Grid
}

func (state *State) String() string {
	if state == nil {
		return ""
	}
	return fmt.Sprintf("{width: %d, height: %d, lastPlayer: %d, grid: %v}", state.grid.width, state.grid.height, state.lastPlayer, state.grid)
}

func (state *State) SetWidth(width int) *State {
	if state == nil {
		state = NewState()
	}
	return &State{
		lastPlayer: state.lastPlayer,
		grid:       NewGrid(width, state.grid.height),
	}
}

func (state *State) SetHeight(height int) *State {
	if state == nil {
		state = NewState()
	}
	return &State{
		lastPlayer: state.lastPlayer,
		grid:       NewGrid(state.grid.width, height),
	}
}

func (state *State) SetCell(x int, y int, player int) *State {
	if state == nil {
		return nil
	}
	return &State{
		lastPlayer: player,
		grid:       state.grid.SetCell(x, y, player),
	}
}

func NewState() *State {
	return &State{
		lastPlayer: 0,
		grid:       NewGrid(0, 0),
	}
}
