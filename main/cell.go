package main

import (
	"fmt"
)

type Cell struct {
	column int
	row    int
}

func (cell *Cell) String() string {
	if cell == nil {
		return ""
	}
	return fmt.Sprintf("(%d,%d)", cell.column, cell.row)
}

func NewCell(column int, row int) *Cell {
	return &Cell{column, row}
}
