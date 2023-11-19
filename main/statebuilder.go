package main

import (
	"fmt"
	"os"
)

type StateBuilder struct {
	inputStream *os.File
}

func (_ StateBuilder) buildInitState() *State {
	return NewState().SetWidth(3).SetHeight(3)
}

func (builder StateBuilder) buildTurnState(previousState *State) *State {
	var opponentRow, opponentColumn int
	fmt.Fscan(builder.inputStream, &opponentRow, &opponentColumn)
	var nbAvailableActions int
	fmt.Fscan(builder.inputStream, &nbAvailableActions)
	availableActions := make([]*Cell, nbAvailableActions)
	for i := 0; i < nbAvailableActions; i++ {
		var row, col int
		fmt.Fscan(builder.inputStream, &row, &col)
		availableActions[i] = NewCell(col, row)
	}
	return previousState.SetOpponent(NewCell(opponentColumn, opponentRow)).SetAvailableActions(availableActions)
}

func NewStateBuilder(inputStream *os.File) StateBuilder {
	return StateBuilder{inputStream}
}
