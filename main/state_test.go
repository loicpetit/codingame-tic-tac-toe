package main

import (
	"fmt"
	"testing"
)

func TestNewState(t *testing.T) {
	state := NewState()
	if nil == state {
		t.Fatal("No state returned")
	}
	if nil == state.grid {
		t.Fatal("No grid in state")
	}
	expectedWidth := 0
	if expectedWidth != state.grid.width {
		t.Errorf("Expected grid width %d but was %d", expectedWidth, state.grid.width)
	}
	expectedHeight := 0
	if expectedHeight != state.grid.height {
		t.Errorf("Expected grid height %d but was %d", expectedHeight, state.grid.height)
	}
	expectedPlayer := 0
	if expectedPlayer != state.lastPlayer {
		t.Errorf("Expected last player %d but was %d", expectedPlayer, state.lastPlayer)
	}
}

func TestSetStateWidth(t *testing.T) {
	state := NewState()
	newState := state.SetWidth(2)
	if nil == newState {
		t.Fatal("No state returned")
	}
	if state == newState {
		t.Fatal("Should return a new state")
	}
	if nil == newState.grid {
		t.Fatal("No grid in new state")
	}
	if state.grid == newState.grid {
		t.Error("Expected a new grid")
	}
	expectedWidth := 0
	if expectedWidth != state.grid.width {
		t.Errorf("Expected old grid width %d but was %d", expectedWidth, state.grid.width)
	}
	expectedNewWidth := 2
	if expectedNewWidth != newState.grid.width {
		t.Errorf("Expected new grid width %d but was %d", expectedNewWidth, newState.grid.width)
	}
	if state.grid.height != newState.grid.height {
		t.Errorf("Grid height changed")
	}
	if state.lastPlayer != newState.lastPlayer {
		t.Errorf("Last player changed")
	}
}

func TestSetStateHeight(t *testing.T) {
	state := NewState()
	newState := state.SetHeight(2)
	if nil == newState {
		t.Fatal("No state returned")
	}
	if state == newState {
		t.Fatal("Should return a new state")
	}
	if nil == newState.grid {
		t.Fatal("No grid in new state")
	}
	if state.grid == newState.grid {
		t.Error("Expected a new grid")
	}
	expectedHeight := 0
	if expectedHeight != state.grid.height {
		t.Errorf("Expected old grid height %d but was %d", expectedHeight, state.grid.height)
	}
	expectedNewHeight := 2
	if expectedNewHeight != newState.grid.height {
		t.Errorf("Expected new grid height %d but was %d", expectedNewHeight, newState.grid.height)
	}
	if state.grid.width != newState.grid.width {
		t.Errorf("Grid width changed")
	}
	if state.lastPlayer != newState.lastPlayer {
		t.Errorf("Last player changed")
	}
}

// todo set cell
func TestSetCell(t *testing.T) {
	state := NewState().SetWidth(3).SetHeight(3)
	newState := state.SetCell(2, 1, 1)
	if nil == newState {
		t.Fatal("No state returned")
	}
	if state == newState {
		t.Fatal("Should return a new state")
	}
	if nil == newState.grid {
		t.Fatal("No grid in new state")
	}
	if state.grid == newState.grid {
		t.Error("Expected a new grid")
	}
	expectedHeight := 3
	if expectedHeight != newState.grid.height {
		t.Errorf("Expected grid height %d but was %d", expectedHeight, state.grid.height)
	}
	expectedWidth := 3
	if expectedWidth != newState.grid.width {
		t.Errorf("Expected grid width %d but was %d", expectedWidth, newState.grid.width)
	}
	for x := 0; x < expectedWidth; x++ {
		for y := 0; y < expectedHeight; y++ {
			if x == 2 && y == 1 {
				if newState.grid.cells[x][y] != 1 {
					t.Errorf("Expected cell (2,1) to have value 1 but was %d", newState.grid.cells[x][y])
				}
			} else if newState.grid.cells[x][y] != 0 {
				t.Errorf("Expected cell (2,1) to have value 0 but was %d", newState.grid.cells[x][y])
			}
		}
	}
	expectedOldLastPlayer := 0
	if expectedOldLastPlayer != state.lastPlayer {
		t.Errorf("Expected old last player %d but was %d", expectedOldLastPlayer, state.lastPlayer)
	}
	expectedNewLastPlayer := 1
	if expectedNewLastPlayer != newState.lastPlayer {
		t.Errorf("Expected new last player %d but was %d", expectedNewLastPlayer, newState.lastPlayer)
	}
}

func TestStateSetters(t *testing.T) {
	state := NewState().
		SetWidth(3).
		SetHeight(2).
		SetCell(0, 0, 1).
		SetCell(1, 0, 2).
		SetCell(1, 1, 1)
	if state.grid.width != 3 {
		t.Error("Bad width")
	}
	if state.grid.height != 2 {
		t.Error("Bad height")
	}
	for x, column := range state.grid.cells {
		for y, value := range column {
			if x == 0 && y == 0 {
				if value != 1 {
					t.Error("Bad value at (0,0)")
				}
			} else if x == 1 && y == 0 {
				if value != 2 {
					t.Error("Bad value at (1,0)")
				}
			} else if x == 1 && y == 1 {
				if value != 1 {
					t.Error("Bad value at (1,1)")
				}
			} else if value != 0 {
				t.Errorf("Bad value at (%d,%d)", x, y)
			}
		}
	}
}

func TestStateHash(t *testing.T) {
	dataSet := []struct {
		actions []struct {
			x, y, value int
		}
		expectedHash string
	}{
		{
			expectedHash: "0-000000000",
		},
		{
			actions: []struct {
				x     int
				y     int
				value int
			}{
				{0, 0, 1},
				{1, 1, 2},
				{2, 2, 1},
			},
			expectedHash: "1-100020001",
		},
		{
			actions: []struct {
				x     int
				y     int
				value int
			}{
				{0, 0, 2},
				{1, 0, 1},
				{2, 0, 2},
				{0, 1, 1},
				{1, 1, 2},
				{2, 1, 1},
				{0, 2, 2},
				{1, 2, 1},
				{2, 2, 2},
			},
			expectedHash: "2-212121212",
		},
	}
	for i, data := range dataSet {
		testName := fmt.Sprintf("Set%d", i+1)
		t.Run(testName, func(t *testing.T) {
			state := NewState().SetWidth(3).SetHeight(3)
			for _, action := range data.actions {
				state = state.SetCell(action.x, action.y, action.value)
			}
			hash := state.Hash()
			if hash != data.expectedHash {
				t.Errorf("Expected hash %s but was %s", data.expectedHash, hash)
			}
		})
	}
}
