package main

import (
	"fmt"
)

type TicTacToeGame struct{}

func (game TicTacToeGame) Start() *State {
	return NewState().SetWidth(3).SetHeight(3)
}

func (game TicTacToeGame) validatePlayer(player int) {
	if player < 1 || player > 2 {
		panic("Player should be 1 or 2")
	}
}

func (game TicTacToeGame) GetAvailableActions(state *State, player int) []*Action {
	game.validatePlayer(player)
	actions := make([]*Action, 0)
	if state == nil {
		return actions
	}
	for x := 0; x < state.grid.width; x++ {
		for y := 0; y < state.grid.height; y++ {
			if state.grid.cells[x][y] <= 0 {
				actions = append(actions, NewAction(player, x, y))
			}
		}
	}
	return actions
}

func (game TicTacToeGame) Play(state *State, action *Action) *State {
	if state == nil {
		panic("State should not be nil to play")
	}
	if action == nil {
		panic("Action should not be nil to play")
	}
	game.validatePlayer(action.player)
	if state.lastPlayer == action.player {
		panic(fmt.Sprintf("Player %d cannot play twice in a row", action.player))
	}
	if state.grid.cells[action.x][action.y] != 0 {
		panic(fmt.Sprintf("Cannot play at (%d, !%d), already played by %d", action.x, action.y, state.grid.cells[action.x][action.y]))
	}
	return state.SetCell(action.x, action.y, action.player)
}

func (game TicTacToeGame) Winner(state *State) int {
	if state == nil {
		panic("State should not be nil to search a winner")
	}
	if state.grid.width != state.grid.height {
		panic("State grid should be a square to search a winner")
	}
	size := state.grid.width
	winner := 0
	var checkRow = func(y int) int {
		if winner != 0 {
			return winner
		}
		if size == 0 {
			return 0
		}
		player := state.grid.cells[0][y]
		for x := 1; x < size; x++ {
			if player != state.grid.cells[x][y] {
				return 0
			}
		}
		return player
	}
	var checkColumn = func(x int) int {
		if winner != 0 {
			return winner
		}
		if size == 0 {
			return 0
		}
		player := state.grid.cells[x][0]
		for y := 1; y < size; y++ {
			if player != state.grid.cells[x][y] {
				return 0
			}
		}
		return player
	}
	var checkDiagonale1 = func() int {
		if winner != 0 {
			return winner
		}
		if size == 0 {
			return 0
		}
		player := state.grid.cells[0][0]
		for i := 1; i < size; i++ {
			if player != state.grid.cells[i][i] {
				return 0
			}
		}
		return player
	}
	var checkDiagonale2 = func() int {
		if winner != 0 {
			return winner
		}
		if size == 0 {
			return 0
		}
		maxIndex := size - 1
		player := state.grid.cells[maxIndex][0]
		for i := 1; i < size; i++ {
			if player != state.grid.cells[maxIndex-i][0+i] {
				return 0
			}
		}
		return player
	}
	for i := 0; i < size; i++ {
		winner = checkRow(i)
		winner = checkColumn(i)
	}
	winner = checkDiagonale1()
	winner = checkDiagonale2()
	return winner
}

func NewTicTacToeGame() Game[State, Action] {
	return TicTacToeGame{}
}
