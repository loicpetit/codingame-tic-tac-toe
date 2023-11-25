package main

import (
	"fmt"
	"strings"
)

type Grid struct {
	cells  [][]int
	height int
	width  int
}

func (grid *Grid) String() string {
	if grid == nil {
		return ""
	}
	var s strings.Builder
	s.WriteRune('{')
	for x := 0; x < grid.width; x++ {
		for y := 0; y < grid.height; y++ {
			if x != 0 || y != 0 {
				s.WriteRune(';')
			}
			s.WriteString(fmt.Sprintf("(%d,%d)=%d", x, y, grid.cells[x][y]))
		}
	}
	s.WriteRune('}')
	return s.String()
}

func (grid *Grid) SetCell(x int, y int, value int) *Grid {
	if grid == nil {
		return nil
	}
	if x < 0 || x >= grid.width {
		panic(fmt.Sprintf("x %d is an invalid value for a grid of width %d", x, grid.width))
	}
	if y < 0 || y >= grid.height {
		panic(fmt.Sprintf("y %d is an invalid value for a grid of height %d", y, grid.height))
	}
	newGrid := Grid{
		cells:  copyCells(grid.cells),
		height: grid.height,
		width:  grid.width,
	}
	newGrid.cells[x][y] = value
	return &newGrid
}

func copyCells(cells [][]int) [][]int {
	width := len(cells)
	newCells := make([][]int, width)
	for x := range cells {
		height := len(cells[x])
		newCells[x] = make([]int, height)
		copy(newCells[x], cells[x])
	}
	return newCells
}

func NewGrid(width int, height int) *Grid {
	if width < 0 {
		panic(fmt.Sprintf("Invalid width %d", width))
	}
	if height < 0 {
		panic(fmt.Sprintf("Invalid heigth %d", height))
	}
	cells := make([][]int, width)
	for i := range cells {
		cells[i] = make([]int, height)
	}
	return &Grid{cells: cells, height: height, width: width}
}
