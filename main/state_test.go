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
	expectedWidth := 0
	if expectedWidth != state.width {
		t.Errorf("Expected width %d but was %d", expectedWidth, state.width)
	}
	expectedHeight := 0
	if expectedHeight != state.height {
		t.Errorf("Expected height %d but was %d", expectedHeight, state.height)
	}
	if nil != state.opponent {
		t.Errorf("Expected opponent nil but was %v", state.opponent)
	}
	expectedNbAvailableActions := 0
	nbAvailableActions := len(state.availableActions)
	if expectedNbAvailableActions != nbAvailableActions {
		t.Errorf("Expected nb available actions %d but was %d", expectedNbAvailableActions, nbAvailableActions)
	}
	if state.grid == nil {
		t.Error("Expected an initialized grid")
	} else {
		if state.grid.width != expectedWidth {
			t.Errorf("Expected grid width %d but was %d", expectedWidth, state.grid.width)
		}
		if state.grid.height != expectedHeight {
			t.Errorf("Expected grid height %d but was %d", expectedHeight, state.grid.height)
		}
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
	expectedWidth := 0
	if expectedWidth != state.width {
		t.Errorf("Expected width %d but was %d", expectedWidth, state.width)
	}
	expectedNewWidth := 2
	if expectedNewWidth != newState.width {
		t.Errorf("Expected width %d but was %d", expectedNewWidth, newState.width)
	}
	if state.height != newState.height {
		t.Errorf("Height changed")
	}
	if state.opponent != newState.opponent {
		t.Errorf("Opponent changed")
	}
	assertCellsAreEqual(t, state.availableActions, newState.availableActions)
	if state.grid == newState.grid {
		t.Error("Expected a new grid")
	}
	if newState.grid.width != expectedNewWidth {
		t.Errorf("Expected new grid width to be %d but was %d", expectedNewWidth, newState.grid.width)
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
	expectedHeight := 0
	if expectedHeight != state.height {
		t.Errorf("Expected height %d but was %d", expectedHeight, state.height)
	}
	expectedNewHeight := 2
	if expectedNewHeight != newState.height {
		t.Errorf("Expected height %d but was %d", expectedNewHeight, newState.height)
	}
	if state.width != newState.width {
		t.Errorf("Width changed")
	}
	if state.opponent != newState.opponent {
		t.Errorf("Opponent changed")
	}
	assertCellsAreEqual(t, state.availableActions, newState.availableActions)
	if state.grid == newState.grid {
		t.Error("Expected a new grid")
	}
	if newState.grid.height != expectedNewHeight {
		t.Errorf("Expected new grid height to be %d but was %d", expectedNewHeight, newState.grid.height)
	}
}

func TestSetStateOpponent(t *testing.T) {
	state := NewState().SetWidth(3).SetHeight(3)
	newState := state.SetOpponent(&Cell{2, 1})
	if nil == newState {
		t.Fatal("No state returned")
	}
	if state == newState {
		t.Fatal("Should return a new state")
	}
	if state.width != newState.width {
		t.Errorf("Width changed")
	}
	if state.height != newState.width {
		t.Errorf("Width changed")
	}
	if state.opponent == newState.opponent {
		t.Errorf("Opponent should change")
	}
	if newState.grid.cells[2][1] != 2 {
		t.Error("Grid (2,1) should be 2")
	}
	assertCellsAreEqual(t, state.availableActions, newState.availableActions)
}

func TestSetStateOutOfRangeOpponent(t *testing.T) {
	defer func() {
		if recover() == nil {
			t.Error("Panic expected")
		}
	}()
	NewState().SetOpponent(&Cell{1, 1})
}

func TestSetStatePlayedCellOpponent(t *testing.T) {
	defer func() {
		if recover() == nil {
			t.Error("Panic expected")
		}
	}()
	NewState().
		SetWidth(3).
		SetHeight(3).
		SetOpponent(&Cell{1, 1}).
		SetOpponent(&Cell{1, 1})
}

func TestSetStatePlayer(t *testing.T) {
	state := NewState().SetWidth(3).SetHeight(3)
	newState := state.SetPlayer(&Cell{2, 1})
	if nil == newState {
		t.Fatal("No state returned")
	}
	if state == newState {
		t.Fatal("Should return a new state")
	}
	if state.width != newState.width {
		t.Errorf("Width changed")
	}
	if state.height != newState.width {
		t.Errorf("Width changed")
	}
	if state.opponent != newState.opponent {
		t.Errorf("Opponent changed")
	}
	if newState.grid.cells[2][1] != 1 {
		t.Error("Grid (2,1) should be 1")
	}
	assertCellsAreEqual(t, state.availableActions, newState.availableActions)
}

func TestSetStateOutOfRangePlayer(t *testing.T) {
	defer func() {
		if recover() == nil {
			t.Error("Panic expected")
		}
	}()
	NewState().SetPlayer(&Cell{1, 1})
}

func TestSetStatePlayedCellPlayer(t *testing.T) {
	defer func() {
		if recover() == nil {
			t.Error("Panic expected")
		}
	}()
	NewState().
		SetWidth(3).
		SetHeight(3).
		SetOpponent(&Cell{1, 1}).
		SetPlayer(&Cell{1, 1})
}

func TestSetStateAvailableActions(t *testing.T) {
	state := NewState()
	newState := state.SetAvailableActions([]*Cell{{1, 1}, {1, 2}})
	if nil == newState {
		t.Fatal("No state returned")
	}
	if state == newState {
		t.Fatal("Should return a new state")
	}
	if state.width != newState.width {
		t.Errorf("Width changed")
	}
	if state.height != newState.width {
		t.Errorf("Width changed")
	}
	if state.opponent != newState.opponent {
		t.Errorf("Opponent changed")
	}
	assertCellsAreNotEqual(t, state.availableActions, newState.availableActions)
}

func TestStateSetters(t *testing.T) {
	state := NewState().
		SetWidth(3).
		SetHeight(2).
		SetPlayer(&Cell{0, 0}).
		SetOpponent(&Cell{1, 0}).
		SetPlayer(&Cell{1, 1}).
		SetAvailableActions([]*Cell{{2, 0}, {0, 1}})
	if state.width != 3 {
		t.Error("Bad width")
	}
	if state.height != 2 {
		t.Error("Bad height")
	}
	if state.opponent.column != 1 {
		t.Error("Bad opponent column")
	}
	if state.opponent.row != 0 {
		t.Error("Bad opponent row")
	}
	if len(state.availableActions) != 2 {
		t.Error("Bad nb available actions")
	}
	if state.availableActions[0].column != 2 {
		t.Error("Bad available actions 0 column")
	}
	if state.availableActions[0].row != 0 {
		t.Error("Bad available actions 0 row")
	}
	if state.availableActions[1].column != 0 {
		t.Error("Bad available actions 1 column")
	}
	if state.availableActions[1].row != 1 {
		t.Error("Bad available actions 1 row")
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

func assertCellsAreEqual(t *testing.T, array1 []*Cell, array2 []*Cell) {
	equal, msg := areCellsEqual(t, array1, array2)
	if !equal {
		t.Error(msg)
	}
}

func assertCellsAreNotEqual(t *testing.T, array1 []*Cell, array2 []*Cell) {
	equal, _ := areCellsEqual(t, array1, array2)
	if equal {
		t.Error("Arrays are equal")
	}
}

func areCellsEqual(t *testing.T, array1 []*Cell, array2 []*Cell) (bool, string) {
	nbCell1 := len(array1)
	nbCell2 := len(array2)
	if nbCell1 != nbCell2 {
		return false, fmt.Sprintf("First array has length %d bu second has length %d", nbCell1, nbCell2)
	}
	for i := 0; i < nbCell1; i++ {
		cell1 := array1[i]
		cell2 := array2[i]
		if cell1 != cell2 {
			return false, fmt.Sprintf("Cell %d is different", i)
		}
	}
	return true, ""
}
