package main

type Cell struct {
	column int
	row    int
}

func NewCell(column int, row int) *Cell {
	return &Cell{column, row}
}
