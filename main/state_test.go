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
}

func TestSetStateOpponent(t *testing.T) {
	state := NewState()
	newState := state.SetOpponent(&Cell{1, 1})
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
	assertCellsAreEqual(t, state.availableActions, newState.availableActions)
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
		SetOpponent(&Cell{1, 0}).
		SetAvailableActions([]*Cell{{0, 0}, {2, 0}})
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
	if state.availableActions[0].column != 0 {
		t.Error("Bad available actions 0 column")
	}
	if state.availableActions[0].row != 0 {
		t.Error("Bad available actions 0 row")
	}
	if state.availableActions[1].column != 2 {
		t.Error("Bad available actions 1 column")
	}
	if state.availableActions[1].row != 0 {
		t.Error("Bad available actions 1 row")
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
