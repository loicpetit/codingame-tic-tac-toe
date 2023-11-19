package main

import (
	"os"
	"testing"
)

func TestStateBuilderInit(t *testing.T) {
	builder := NewStateBuilder(nil)
	state := builder.buildInitState()
	if nil == state {
		t.Fatal("No state returned")
	}
	expectedWidth := 3
	if expectedWidth != state.width {
		t.Errorf("Expected width %d but was %d", expectedWidth, state.width)
	}
	expectedHeight := 3
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

func TestStateBuilderTurn(t *testing.T) {

	read, write, _ := os.Pipe()
	defer read.Close()
	defer write.Close()
	builder := NewStateBuilder(read)
	state := builder.buildInitState()
	done := make(chan bool)
	go func() {
		state = builder.buildTurnState(state)
		state = builder.buildTurnState(state)
		done <- true
	}()
	// turn 1
	write.WriteString("-1 -1\n") // opponent move, -1 -1 if I play first
	write.WriteString("2\n")     // nb available actions
	write.WriteString("2 1\n")   // available action, row col
	write.WriteString("2 2\n")   // available action, row col
	// turn 2
	write.WriteString("2 2\n") // opponent move, -1 -1 if I play first
	write.WriteString("1\n")   // nb available actions
	write.WriteString("2 1\n") // available action, row col
	<-done
	if state.opponent == nil {
		t.Error("Opponent should not be nil")
	} else {
		if state.opponent.column != 2 {
			t.Error("Opponent column should be 2")
		}
		if state.opponent.row != 2 {
			t.Error("Opponent row should be 2")
		}
	}
	if len(state.availableActions) != 1 {
		t.Error("Nb available actions should be 1")
	} else {
		if state.availableActions[0].column != 1 {
			t.Error("Available action 0 column should be 1")
		}
		if state.availableActions[0].row != 2 {
			t.Error("Available action 0 row should be 2")
		}
	}
}
