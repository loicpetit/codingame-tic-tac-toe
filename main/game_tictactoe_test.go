package main

import (
	"fmt"
	"testing"
)

func TestGameStart(t *testing.T) {
	state := NewTicTacToeGame().Start()
	if state == nil {
		t.Fatal("Start state should not be nil")
	}
	expectedWidth := 3
	if state.grid.width != expectedWidth {
		t.Errorf("Expected grid width %d but was %d", expectedWidth, state.grid.width)
	}
	expectedHeight := 3
	if state.grid.height != expectedHeight {
		t.Errorf("Expected grid height %d but was %d", expectedHeight, state.grid.height)
	}
	for x := 0; x < expectedWidth; x++ {
		for y := 0; y < expectedWidth; y++ {
			if state.grid.cells[x][y] != 0 {
				t.Errorf("Expected grid cell (%d,%d) = %d but was %d", x, y, 0, state.grid.cells[x][y])
			}
		}
	}
	expectedLastPlayer := 0
	if state.lastPlayer != expectedLastPlayer {
		t.Errorf("Expected last player %d but was %d", expectedLastPlayer, state.lastPlayer)
	}
}

func TestGameGetAvailableActionsStartState(t *testing.T) {
	game := NewTicTacToeGame()
	actions := game.GetAvailableActions(game.Start(), 1)
	if actions == nil {
		t.Fatal("Actions should not be nil")
	}
	expectedNbAvailableActions := 9
	nbActions := len(actions)
	if nbActions != expectedNbAvailableActions {
		t.Fatalf("Expected %d availables actions but was %d", expectedNbAvailableActions, nbActions)
	}
	checkAction := func(index, expectedPlayer, expectedX, expectedY int) {
		if actions[index].player != expectedPlayer {
			t.Errorf("Expected player %d for action %d but was %d", expectedPlayer, index, actions[index].player)
		}
		if actions[index].x != expectedX {
			t.Errorf("Expected x %d for action %d but was %d", expectedX, index, actions[index].x)
		}
		if actions[index].y != expectedY {
			t.Errorf("Expected y %d for action %d but was %d", expectedY, index, actions[index].y)
		}

	}
	checkAction(0, 1, 0, 0)
	checkAction(1, 1, 0, 1)
	checkAction(2, 1, 0, 2)
	checkAction(3, 1, 1, 0)
	checkAction(4, 1, 1, 1)
	checkAction(5, 1, 1, 2)
	checkAction(6, 1, 2, 0)
	checkAction(7, 1, 2, 1)
	checkAction(8, 1, 2, 2)
}

func TestGameGetAvailableActionsFullState(t *testing.T) {
	game := NewTicTacToeGame()
	actions := game.GetAvailableActions(
		game.Start().
			SetCell(0, 0, 1).
			SetCell(1, 0, 2).
			SetCell(2, 0, 1).
			SetCell(0, 1, 1).
			SetCell(1, 1, 2).
			SetCell(2, 1, 1).
			SetCell(0, 2, 2).
			SetCell(1, 2, 1).
			SetCell(2, 2, 2),
		1,
	)
	if actions == nil {
		t.Fatal("Actions should not be nil")
	}
	expectedNbAvailableActions := 0
	nbActions := len(actions)
	if nbActions != expectedNbAvailableActions {
		t.Fatalf("Expected %d availables actions but was %d", expectedNbAvailableActions, nbActions)
	}
}

func TestGameGetAvailableActionsBadPlayer(t *testing.T) {
	defer func() {
		if recover() == nil {
			t.Error("Panic is expected")
		}
	}()
	game := NewTicTacToeGame()
	game.GetAvailableActions(game.Start(), 0)
}

func TestGameGetNextPayer(t *testing.T) {
	dataSet := []struct {
		testName       string
		action         *Action
		expectedPlayer int
	}{
		{"Nil action", nil, 1},
		{"Action player 1", NewAction(1, 1, 1), 2},
		{"Action player 2", NewAction(2, 2, 2), 1},
	}
	game := NewTicTacToeGame()
	for _, data := range dataSet {
		t.Run(data.testName, func(t *testing.T) {
			player := game.GetNextPlayer(data.action)
			if player != data.expectedPlayer {
				t.Error("Bad player result")
			}
		})
	}
}

func TestGamePlay(t *testing.T) {
	game := NewTicTacToeGame()
	state := game.Start()
	state = game.Play(game.Start(), NewAction(1, 0, 2))
	state = game.Play(state, NewAction(2, 1, 2))
	state = game.Play(state, NewAction(1, 0, 1))
	state = game.Play(state, NewAction(2, 1, 1))
	state = game.Play(state, NewAction(1, 0, 0))
}

func TestGamePlayWithoutState(t *testing.T) {
	defer func() {
		if recover() == nil {
			t.Error("Panic is expected")
		}
	}()
	game := NewTicTacToeGame()
	game.Play(nil, NewAction(1, 0, 2))
}

func TestGamePlayWithoutAction(t *testing.T) {
	defer func() {
		if recover() == nil {
			t.Error("Panic is expected")
		}
	}()
	game := NewTicTacToeGame()
	game.Play(game.Start(), nil)
}

func TestGamePlayBadPlayer(t *testing.T) {
	defer func() {
		if recover() == nil {
			t.Error("Panic is expected")
		}
	}()
	game := NewTicTacToeGame()
	game.Play(game.Start(), NewAction(0, 0, 2))
}

func TestGamePlaySamePlayer(t *testing.T) {
	defer func() {
		if recover() == nil {
			t.Error("Panic is expected")
		}
	}()
	game := NewTicTacToeGame()
	state := game.Play(game.Start(), NewAction(1, 0, 2))
	game.Play(state, NewAction(1, 0, 1))
}

func TestGamePlayPlayedCell(t *testing.T) {
	defer func() {
		if recover() == nil {
			t.Error("Panic is expected")
		}
	}()
	game := NewTicTacToeGame()
	state := game.Play(game.Start(), NewAction(1, 0, 2))
	game.Play(state, NewAction(2, 0, 2))
}

func TestGameWinner(t *testing.T) {
	game := NewTicTacToeGame()
	dataSet := []struct {
		winner int
		state  *State
	}{
		/* 0  */ {0, game.Start()},
		/* 1  */ {0, game.Start().SetCell(0, 0, 1)},
		/* 2  */ {0, game.Start().SetCell(0, 0, 1).SetCell(1, 0, 2)},
		/* 3  */ {0, game.Start().SetCell(0, 0, 1).SetCell(1, 0, 2).SetCell(0, 1, 1)},
		/* 4  */ {0, game.Start().SetCell(0, 0, 1).SetCell(1, 0, 2).SetCell(0, 1, 1).SetCell(1, 1, 2)},
		/* 5  */ {1, game.Start().SetCell(0, 0, 1).SetCell(1, 0, 2).SetCell(0, 1, 1).SetCell(1, 1, 2).SetCell(0, 2, 1)},
		/* 6  */ {2, game.Start().SetCell(1, 0, 2).SetCell(1, 1, 2).SetCell(1, 2, 2)},
		/* 7  */ {1, game.Start().SetCell(2, 0, 1).SetCell(2, 1, 1).SetCell(2, 2, 1)},
		/* 8  */ {1, game.Start().SetCell(0, 0, 1).SetCell(1, 0, 1).SetCell(2, 0, 1)},
		/* 9  */ {1, game.Start().SetCell(0, 1, 1).SetCell(1, 1, 1).SetCell(2, 1, 1)},
		/* 10 */ {1, game.Start().SetCell(0, 2, 1).SetCell(1, 2, 1).SetCell(2, 2, 1)},
		/* 11 */ {1, game.Start().SetCell(0, 0, 1).SetCell(1, 1, 1).SetCell(2, 2, 1)},
		/* 12 */ {1, game.Start().SetCell(2, 0, 1).SetCell(1, 1, 1).SetCell(0, 2, 1)},
	}
	for i, data := range dataSet {
		testName := fmt.Sprintf("State%d", i)
		t.Run(testName, func(t *testing.T) {
			winner := game.Winner(data.state)
			if winner != data.winner {
				t.Errorf("Expected winner %d for state %d but was %d", data.winner, i, winner)
			}
		})
	}
}
