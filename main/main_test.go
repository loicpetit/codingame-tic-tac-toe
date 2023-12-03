package main

import (
	"os"
	"testing"
	"time"
)

func TestMainNoTurn(t *testing.T) {
	// prepare
	quit := make(chan bool)
	read, write, _ := os.Pipe()
	defer read.Close()
	defer write.Close()
	var finalState *State
	go func() {
		quit <- true
	}()
	finalState = mainFrom(read, quit)
	// assert
	if finalState == nil {
		t.Fatal("Final state should not be nil")
	}
	if finalState.lastPlayer != 0 {
		t.Error("Last player should be 0")
	}
	if finalState.grid.width != 3 {
		t.Error("Grid width should be 3")
	}
	if finalState.grid.height != 3 {
		t.Error("Grid height should be 3")
	}
	if finalState.grid.cells[0][0] != 0 {
		t.Errorf("Grid (%d, %d) should be %d but was %d", 0, 0, 0, finalState.grid.cells[0][0])
	}
	if finalState.grid.cells[0][1] != 0 {
		t.Errorf("Grid (%d, %d) should be %d but was %d", 0, 1, 0, finalState.grid.cells[0][1])
	}
	if finalState.grid.cells[0][2] != 0 {
		t.Errorf("Grid (%d, %d) should be %d but was %d", 0, 2, 0, finalState.grid.cells[0][2])
	}
	if finalState.grid.cells[1][0] != 0 {
		t.Errorf("Grid (%d, %d) should be %d but was %d", 1, 0, 0, finalState.grid.cells[1][0])
	}
	if finalState.grid.cells[1][1] != 0 {
		t.Errorf("Grid (%d, %d) should be %d but was %d", 1, 1, 0, finalState.grid.cells[1][1])
	}
	if finalState.grid.cells[1][2] != 0 {
		t.Errorf("Grid (%d, %d) should be %d but was %d", 1, 2, 0, finalState.grid.cells[1][2])
	}
	if finalState.grid.cells[2][0] != 0 {
		t.Errorf("Grid (%d, %d) should be %d but was %d", 2, 0, 0, finalState.grid.cells[2][0])
	}
	if finalState.grid.cells[2][1] != 0 {
		t.Errorf("Grid (%d, %d) should be %d but was %d", 2, 1, 0, finalState.grid.cells[2][1])
	}
	if finalState.grid.cells[2][2] != 0 {
		t.Errorf("Grid (%d, %d) should be %d but was %d", 2, 2, 0, finalState.grid.cells[2][2])
	}
}

func TestMain(t *testing.T) {
	// prepare
	quit := make(chan bool)
	read, write, _ := os.Pipe()
	defer read.Close()
	defer write.Close()
	go func() {
		// turn 1
		write.WriteString("-1 -1\n") // opponent move row col (y, x)
		write.WriteString("9\n")     // nb available actions
		write.WriteString("0 0\n")   // available action, row col (y, x)
		write.WriteString("0 1\n")   // available action, row col (y, x)
		write.WriteString("0 2\n")   // available action, row col (y, x)
		write.WriteString("1 0\n")   // available action, row col (y, x)
		write.WriteString("1 1\n")   // available action, row col (y, x)
		write.WriteString("1 2\n")   // available action, row col (y, x)
		write.WriteString("2 0\n")   // available action, row col (y, x)
		write.WriteString("2 1\n")   // available action, row col (y, x)
		write.WriteString("2 2\n")   // available action, row col (y, x)
		time.Sleep(1000 * time.Millisecond)
		// turn 2
		write.WriteString("1 1\n") // opponent move row col (y, x)
		write.WriteString("7\n")   // nb available actions
		write.WriteString("0 1\n") // available action, row col (y, x)
		write.WriteString("0 2\n") // available action, row col (y, x)
		write.WriteString("1 0\n") // available action, row col (y, x)
		write.WriteString("1 2\n") // available action, row col (y, x)
		write.WriteString("2 0\n") // available action, row col (y, x)
		write.WriteString("2 1\n") // available action, row col (y, x)
		write.WriteString("2 2\n") // available action, row col (y, x)
		time.Sleep(100 * time.Millisecond)
		// stop
		quit <- true
	}()
	finalState := mainFrom(read, quit)
	// assert
	if finalState == nil {
		t.Fatal("Final state should not be nil")
	}
	if finalState.lastPlayer != 1 {
		t.Error("Last player should be 1")
	}
	if finalState.grid.width != 3 {
		t.Error("Grid width should be 3")
	}
	if finalState.grid.height != 3 {
		t.Error("Grid height should be 3")
	}
	if finalState.grid.cells[0][0] != 1 {
		t.Errorf("Grid (%d, %d) should be %d but was %d", 0, 0, 1, finalState.grid.cells[0][0])
	}
	if finalState.grid.cells[0][1] != 0 {
		t.Errorf("Grid (%d, %d) should be %d but was %d", 0, 1, 0, finalState.grid.cells[0][1])
	}
	if finalState.grid.cells[0][2] != 0 {
		t.Errorf("Grid (%d, %d) should be %d but was %d", 0, 2, 0, finalState.grid.cells[0][2])
	}
	if finalState.grid.cells[1][0] != 1 {
		t.Errorf("Grid (%d, %d) should be %d but was %d", 1, 0, 0, finalState.grid.cells[1][0])
	}
	if finalState.grid.cells[1][1] != 2 {
		t.Errorf("Grid (%d, %d) should be %d but was %d", 1, 1, 2, finalState.grid.cells[1][1])
	}
	if finalState.grid.cells[1][2] != 0 {
		t.Errorf("Grid (%d, %d) should be %d but was %d", 1, 2, 0, finalState.grid.cells[1][2])
	}
	if finalState.grid.cells[2][0] != 0 {
		t.Errorf("Grid (%d, %d) should be %d but was %d", 2, 0, 0, finalState.grid.cells[2][0])
	}
	if finalState.grid.cells[2][1] != 0 {
		t.Errorf("Grid (%d, %d) should be %d but was %d", 2, 1, 0, finalState.grid.cells[2][1])
	}
	if finalState.grid.cells[2][2] != 0 {
		t.Errorf("Grid (%d, %d) should be %d but was %d", 2, 2, 1, finalState.grid.cells[2][2])
	}
}
