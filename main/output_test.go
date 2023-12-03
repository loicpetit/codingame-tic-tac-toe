package main

import (
	"testing"
)

func TestValidateOutput(t *testing.T) {
	action := NewAction(1, 1, 1)
	inputData := &InputData{availableCells: []*Cell{NewCell(2, 2), NewCell(1, 1)}}
	ValidateOutput(action, inputData)
}

func TestValidateOutputWithoutAction(t *testing.T) {
	defer func() {
		if recover() == nil {
			t.Error("Panic is expected")
		}
	}()
	inputData := &InputData{availableCells: []*Cell{NewCell(2, 2), NewCell(1, 1)}}
	ValidateOutput(nil, inputData)
}

func TestValidateOutputWithoutInputData(t *testing.T) {
	defer func() {
		if recover() == nil {
			t.Error("Panic is expected")
		}
	}()
	action := NewAction(1, 1, 1)
	ValidateOutput(action, nil)
}

func TestValidateOutputWithoutAvailableActions(t *testing.T) {
	defer func() {
		if recover() == nil {
			t.Error("Panic is expected")
		}
	}()
	action := NewAction(1, 1, 1)
	inputData := &InputData{}
	ValidateOutput(action, inputData)
}

func TestValidateOutputWithoutValidCell(t *testing.T) {
	defer func() {
		if recover() == nil {
			t.Error("Panic is expected")
		}
	}()
	action := NewAction(1, 2, 1)
	inputData := &InputData{availableCells: []*Cell{NewCell(2, 2), NewCell(1, 1)}}
	ValidateOutput(action, inputData)
}
