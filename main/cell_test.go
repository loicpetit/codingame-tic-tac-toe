package main

import "testing"

func TestNewCell(t *testing.T) {
	cell := NewCell(1, 2)
	if nil == cell {
		t.Fatal("No cell returned")
	}
	expectedColumn := 1
	if expectedColumn != cell.column {
		t.Errorf("Expected column %d but was %d", expectedColumn, cell.column)
	}
	expectedRow := 2
	if expectedRow != cell.row {
		t.Errorf("Expected row %d but was %d", expectedRow, cell.row)
	}
}
