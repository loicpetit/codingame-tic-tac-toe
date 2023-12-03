package main

import (
	"fmt"
	"os"
)

type InputData struct {
	availableCells []*Cell
	opponentAction *Action
}

func ReadInputData(inputStream *os.File) *InputData {
	// opponent action
	var opponentRow, opponentColumn int
	fmt.Fscan(inputStream, &opponentRow, &opponentColumn)
	var action *Action
	if opponentColumn >= 0 && opponentRow >= 0 {
		action = NewAction(2, opponentColumn, opponentRow)
	}
	// avalaible cells
	var nbAvailableCells int
	fmt.Fscan(inputStream, &nbAvailableCells)
	if nbAvailableCells < 0 {
		nbAvailableCells = 0
	}
	availableCells := make([]*Cell, nbAvailableCells)
	for i := 0; i < nbAvailableCells; i++ {
		var row, col int
		fmt.Fscan(inputStream, &row, &col)
		availableCells[i] = NewCell(col, row)
	}
	return &InputData{
		availableCells: availableCells,
		opponentAction: action,
	}
}
