package main

import "fmt"

type State struct {
	width            int
	height           int
	opponent         *Cell
	availableActions []*Cell
	grid             *Grid
}

func (state *State) String() string {
	if state == nil {
		return ""
	}
	return fmt.Sprintf("{width: %d, height: %d, opponent: %v, availableActions: %v, grid: %v}", state.width, state.height, state.opponent, state.availableActions, state.grid)
}

func (state *State) SetWidth(width int) *State {
	if state == nil {
		state = NewState()
	}
	return &State{
		width:            width,
		height:           state.height,
		opponent:         state.opponent,
		availableActions: state.availableActions,
		grid:             NewGrid(width, state.height),
	}
}

func (state *State) SetHeight(height int) *State {
	if state == nil {
		state = NewState()
	}
	return &State{
		width:            state.width,
		height:           height,
		opponent:         state.opponent,
		availableActions: state.availableActions,
		grid:             NewGrid(state.width, height),
	}
}

func (state *State) SetPlayer(player *Cell) *State {
	if state == nil {
		state = NewState()
	}
	if state.grid.cells[player.column][player.row] != 0 {
		panic(fmt.Sprintf("Cell (%d,%d) already used by %d", player.column, player.row, state.grid.cells[player.column][player.row]))
	}
	return &State{
		width:            state.width,
		height:           state.height,
		opponent:         state.opponent,
		availableActions: state.availableActions,
		grid:             state.grid.SetCell(player.column, player.row, 1),
	}
}

func (state *State) SetOpponent(opponent *Cell) *State {
	if state == nil {
		state = NewState()
	}
	if state.grid.cells[opponent.column][opponent.row] != 0 {
		panic(fmt.Sprintf("Cell (%d,%d) already used by %d", opponent.column, opponent.row, state.grid.cells[opponent.column][opponent.row]))
	}
	return &State{
		width:            state.width,
		height:           state.height,
		opponent:         opponent,
		availableActions: state.availableActions,
		grid:             state.grid.SetCell(opponent.column, opponent.row, 2),
	}
}

func (state *State) SetAvailableActions(availableActions []*Cell) *State {
	if state == nil {
		state = NewState()
	}
	return &State{
		width:            state.width,
		height:           state.height,
		opponent:         state.opponent,
		availableActions: availableActions,
		grid:             state.grid,
	}
}

func NewState() *State {
	return &State{
		width:  0,
		height: 0,
		grid:   NewGrid(0, 0),
	}
}
