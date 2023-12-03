package main

import (
	"fmt"
	"os"
)

func ValidateOutput(action *Action, inputData *InputData) {
	if action == nil {
		panic("Cannot validate output without action")
	}
	if inputData == nil {
		panic("Cannot validate output without input data")
	}
	if inputData.availableCells == nil {
		panic("Cannot validate output without input data available cells")
	}
	for _, cell := range inputData.availableCells {
		if cell == nil {
			continue
		}
		if action.x == cell.x && action.y == cell.y {
			return
		}
	}
	panic("Action cell is not in input data available cells")
}

func WriteOutput(action *Action) {
	fmt.Println(action.y, action.x)
}

func WriteDebug(a ...any) {
	fmt.Fprintln(os.Stderr, a...)
}
