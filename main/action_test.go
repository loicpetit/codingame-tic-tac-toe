package main

import (
	"testing"
)

func TestNewAction(t *testing.T) {
	action := NewAction(1, 1, 2)
	if nil == action {
		t.Fatal("No action returned")
	}
	expectedPlayer := 1
	if expectedPlayer != action.player {
		t.Errorf("Expected player %d but was %d", expectedPlayer, action.player)
	}
	expectedX := 1
	if expectedX != action.x {
		t.Errorf("Expected cell column %d but was %d", expectedX, action.x)
	}
	expectedY := 2
	if expectedY != action.y {
		t.Errorf("Expected cell row %d but was %d", expectedY, action.y)
	}
}
