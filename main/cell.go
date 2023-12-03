package main

import (
	"fmt"
)

type Cell struct {
	x int
	y int
}

func (cell *Cell) String() string {
	if cell == nil {
		return ""
	}
	return fmt.Sprintf("(%d,%d)", cell.x, cell.y)
}

func NewCell(x int, y int) *Cell {
	return &Cell{x, y}
}
