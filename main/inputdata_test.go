package main

import (
	"os"
	"testing"
)

func TestReadInputData(t *testing.T) {
	// prepare
	read, write, _ := os.Pipe()
	defer read.Close()
	defer write.Close()
	done := make(chan bool)
	var input1, input2 *InputData
	go func() {
		input1 = ReadInputData(read)
		input2 = ReadInputData(read)
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
	// check
	if input1 == nil {
		// t.Error("Opponent should not be nil")
		t.Fatal("Input 1 should exist")
	}
	if input2 == nil {
		t.Fatal("Input 2 should exist")
	}
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
