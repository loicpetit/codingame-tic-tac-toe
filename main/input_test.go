package main

import (
	"os"
	"testing"
)

func TestReaderRead(t *testing.T) {
	// prepare
	read, write, _ := os.Pipe()
	defer read.Close()
	defer write.Close()
	reader := NewReader(read, NewTicTacToeGame())
	go func() {
		// turn 1
		write.WriteString("-1 -1\n") // opponent move, -1 -1 if I play first
		write.WriteString("2\n")     // nb available actions
		write.WriteString("2 1\n")   // available action, row col
		write.WriteString("2 2\n")   // available action, row col
		// turn 2
		write.WriteString("2 2\n") // opponent move, -1 -1 if I play first
		write.WriteString("1\n")   // nb available actions
		write.WriteString("2 1\n") // available action, row col
	}()
	inputs := reader.Read()
	input1 := <-inputs
	input2 := <-inputs
	// check
	if input1.opponentAction != nil {
		t.Error("Input 1 action should be null")
	}
	if input1.opponentAction != nil {
		t.Error("Input 1 action should be null")
	}
	if len(input1.availableCells) != 2 {
		t.Error("Input 1 should have 2 available cells")
	} else {
		if input1.availableCells[0].x != 1 {
			t.Errorf("Input 1 available cell 0 x should be 1 but is %d", input1.availableCells[0].x)
		}
		if input1.availableCells[0].y != 2 {
			t.Errorf("Input 1 available cell 0 y should be 2 but is %d", input1.availableCells[0].y)
		}
		if input1.availableCells[1].x != 2 {
			t.Errorf("Input 1 available cell 1 x should be 2 but is %d", input1.availableCells[1].x)
		}
		if input1.availableCells[1].y != 2 {
			t.Errorf("Input 1 available cell 1 y should be 2 but is %d", input1.availableCells[1].y)
		}
	}
	if len(input2.availableCells) != 1 {
		t.Error("Input 2 should have 1 available cell")
	} else {
		if input2.availableCells[0].x != 1 {
			t.Errorf("Input 2 available cell 0 x should be 1 but is %d", input2.availableCells[0].x)
		}
		if input2.availableCells[0].y != 2 {
			t.Errorf("Input 2 available cell 0 y should be 2 but is %d", input2.availableCells[0].y)
		}
	}
}

func TestReaderUpdateState(t *testing.T) {
	state := NewReader(nil, NewTicTacToeGame()).UpdateState(NewState().SetHeight(3).SetWidth(3), NewInput(nil, NewAction(2, 0, 0)))
	if state == nil {
		t.Fatal("State should not be nil")
	}
	if state.lastPlayer != 2 {
		t.Errorf("Last player should be 2 but is %d", state.lastPlayer)
	}
	if state.grid.cells[0][0] != 2 {
		t.Errorf("Cell (0,0) should be 2 but is %d", state.grid.cells[0][0])
	}
}

func TestReaderUpdateStateWithNilState(t *testing.T) {
	defer func() {
		if recover() == nil {
			t.Error("Panic is expected")
		}
	}()
	NewReader(nil, NewTicTacToeGame()).UpdateState(nil, NewInput(nil, NewAction(2, 0, 0)))
}

func TestReaderUpdateStateWithNilOpponentAction(t *testing.T) {
	state := NewState().SetHeight(3).SetWidth(3)
	newState := NewReader(nil, NewTicTacToeGame()).UpdateState(state, NewInput(nil, nil))
	if state != newState {
		t.Error("State should not change if no action")
	}
}

func TestReaderValidateAction(t *testing.T) {
	action := NewAction(1, 1, 1)
	input := NewInput([]*Cell{NewCell(2, 2), NewCell(1, 1)}, nil)
	NewReader(nil, NewTicTacToeGame()).ValidateAction(action, input)
}

func TestReaderValidateActionWithoutAction(t *testing.T) {
	defer func() {
		if recover() == nil {
			t.Error("Panic is expected")
		}
	}()
	input := NewInput([]*Cell{NewCell(2, 2), NewCell(1, 1)}, nil)
	NewReader(nil, NewTicTacToeGame()).ValidateAction(nil, input)
}

func TestReaderValidateActionWithoutAvailableActions(t *testing.T) {
	defer func() {
		if recover() == nil {
			t.Error("Panic is expected")
		}
	}()
	action := NewAction(1, 1, 1)
	input := NewInput(nil, nil)
	NewReader(nil, NewTicTacToeGame()).ValidateAction(action, input)
}

func TestReaderValidateActionWithoutValidCell(t *testing.T) {
	defer func() {
		if recover() == nil {
			t.Error("Panic is expected")
		}
	}()
	action := NewAction(1, 2, 1)
	input := NewInput([]*Cell{NewCell(2, 2), NewCell(1, 1)}, nil)
	NewReader(nil, NewTicTacToeGame()).ValidateAction(action, input)
}
