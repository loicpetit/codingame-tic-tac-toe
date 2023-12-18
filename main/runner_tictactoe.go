package main

import (
	"os"
)

func NewTicTacToeRunner(
	stream *os.File,
	game Game[State, Action],
	strategy Strategy[State, Action],
) Runner[Input, State, Action] {
	return NewRunner[Input, State, Action](
		game,
		NewReader(stream, game),
		strategy,
		NewWriter(),
	)
}
