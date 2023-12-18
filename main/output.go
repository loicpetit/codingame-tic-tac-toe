package main

import (
	"fmt"
)

type Writer struct{}

func (Writer) Write(action *Action) {
	if action == nil {
		panic("No action to write")
	}
	WriteOutput(fmt.Sprintf("%d %d", action.y, action.x))
}

func NewWriter() OutputWriter[Action] {
	return Writer{}
}
