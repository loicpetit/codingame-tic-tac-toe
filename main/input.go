package main

import (
	"fmt"
	"os"
)

type Input struct {
	availableCells []*Cell
	opponentAction *Action
}

func NewInput(availableCells []*Cell, opponentAction *Action) Input {
	return Input{
		availableCells: availableCells,
		opponentAction: opponentAction,
	}
}

type Reader struct {
	game   Game[State, Action]
	stream *os.File
}

// Read implements InputReader.
func (reader Reader) Read() chan Input {
	inputs := make(chan Input)
	go func() {
		for {
			// opponent action
			var opponentRow, opponentColumn int
			fmt.Fscan(reader.stream, &opponentRow, &opponentColumn)
			var action *Action
			if opponentColumn >= 0 && opponentRow >= 0 {
				action = NewAction(2, opponentColumn, opponentRow)
			}
			// avalaible cells
			var nbAvailableCells int
			fmt.Fscan(reader.stream, &nbAvailableCells)
			if nbAvailableCells < 0 {
				nbAvailableCells = 0
			}
			availableCells := make([]*Cell, nbAvailableCells)
			for i := 0; i < nbAvailableCells; i++ {
				var row, col int
				fmt.Fscan(reader.stream, &row, &col)
				availableCells[i] = NewCell(col, row)
			}
			inputs <- NewInput(availableCells, action)
		}
	}()
	return inputs
}

// UpdateState implements InputReader.
func (reader Reader) UpdateState(state *State, input Input) *State {
	if state == nil {
		panic("Cannot update a nil state")
	}
	if input.opponentAction != nil {
		state = reader.game.Play(state, input.opponentAction)
	}
	return state
}

// ValidateAction implements InputReader.
func (Reader) ValidateAction(action *Action, input Input) {
	if action == nil {
		panic("Cannot validate output without action")
	}
	if input.availableCells == nil {
		panic("Cannot validate output without input data available cells")
	}
	for _, cell := range input.availableCells {
		if cell == nil {
			continue
		}
		if action.x == cell.x && action.y == cell.y {
			return
		}
	}
	panic("Action cell is not in input available cells")
}

func NewReader(stream *os.File, game Game[State, Action]) InputReader[Input, State, Action] {
	return Reader{
		game:   game,
		stream: stream,
	}
}
