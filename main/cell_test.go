package main

import "testing"

func TestNewCell(t *testing.T) {
	cell := NewCell(1, 2)
	if nil == cell {
		t.Fatal("No cell returned")
	}
	expectedX := 1
	if expectedX != cell.x {
		t.Errorf("Expected column %d but was %d", expectedX, cell.x)
	}
	expectedY := 2
	if expectedY != cell.y {
		t.Errorf("Expected row %d but was %d", expectedY, cell.y)
	}
}
