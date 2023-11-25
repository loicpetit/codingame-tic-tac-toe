package main

import (
	"fmt"
	"testing"
)

func TestNewGridValidSizes(t *testing.T) {
	sizes := []struct {
		width, height int
	}{
		{0, 0},
		{0, 1},
		{1, 0},
		{3, 3},
	}
	for _, size := range sizes {
		testName := fmt.Sprintf("w%d,h%d", size.width, size.height)
		t.Run(testName, func(t *testing.T) {
			grid := NewGrid(size.width, size.height)
			if grid.width != size.width {
				t.Errorf("Expected grid width %d but was %d", size.width, grid.width)
			}
			if grid.height != size.height {
				t.Errorf("Expected grid height %d but was %d", size.height, grid.height)
			}
			width := len(grid.cells)
			if width != size.width {
				t.Errorf("Expected cells width %d but was %d", size.width, width)
			}
			for x, column := range grid.cells {
				height := len(column)
				if height != size.height {
					t.Errorf("Expected cells height %d for column %d but was %d", size.height, x, height)
				}
				for y, value := range column {
					if value != 0 {
						t.Errorf("Expected value 0 at (%d,%d)", x, y)
					}
				}
			}
		})
	}
}

func TestNewGridInvalidSizes(t *testing.T) {
	sizes := []struct {
		width, height int
	}{
		{-1, 0},
		{0, -1},
		{-1, -1},
	}
	for _, size := range sizes {
		testName := fmt.Sprintf("w%d,h%d", size.width, size.height)
		t.Run(testName, func(t *testing.T) {
			defer func() {
				if recover() == nil {
					t.Error("Panic is expected")
				}
			}()
			NewGrid(size.width, size.height)
		})
	}
}

func TestGridSetValidCell(t *testing.T) {
	dataSet := []struct {
		x, y, value int
	}{
		{0, 0, 1},
		{2, 2, -1},
		{1, 0, 2},
	}
	originalGrid := NewGrid(3, 3).SetCell(0, 0, 2)
	// fmt.Println("Original grid (before):", originalGrid)
	for _, data := range dataSet {
		testName := fmt.Sprintf("x%d,y%d,v%d", data.x, data.y, data.value)
		t.Run(testName, func(t *testing.T) {
			grid := originalGrid.SetCell(data.x, data.y, data.value)
			if grid == originalGrid {
				t.Fatal("Expected a new grid")
			}
			if grid.width != originalGrid.width {
				t.Errorf("Expected grid width %d but was %d", originalGrid.width, grid.width)
			}
			if grid.height != originalGrid.height {
				t.Errorf("Expected grid height %d but was %d", originalGrid.height, grid.height)
			}
			width := len(grid.cells)
			if width != originalGrid.width {
				t.Errorf("Expected cells width %d but was %d", originalGrid.width, width)
			}
			for x, column := range grid.cells {
				height := len(column)
				if height != originalGrid.height {
					t.Errorf("Expected cells height %d for column %d but was %d", originalGrid.height, x, height)
				}
				for y, value := range column {
					if x == data.x && y == data.y {
						if value != data.value {
							t.Errorf("Expected value %d at (%d,%d) but was %d", data.value, x, y, value)
						}
						if originalGrid.cells[x][y] == value {
							t.Errorf("Original grid changed at (%d,%d)", x, y)
						}
					} else if value != originalGrid.cells[x][y] {
						t.Errorf("Expected value %d at (%d,%d) but was %d", originalGrid.cells[x][y], x, y, value)
					}
				}
			}
		})
	}
	// fmt.Println("Original grid (after):", originalGrid)
}

func TestGridSetInvalidCell(t *testing.T) {
	dataSet := []struct {
		x, y, value int
	}{
		{-2, 0, 1},
		{2, -2, -1},
		{-1, -1, 2},
	}
	originalGrid := NewGrid(3, 3).SetCell(0, 0, 2)
	for _, data := range dataSet {
		testName := fmt.Sprintf("x%d,y%d", data.x, data.y)
		t.Run(testName, func(t *testing.T) {
			defer func() {
				if recover() == nil {
					t.Error("Panic is expected")
				}
			}()
			originalGrid.SetCell(data.x, data.y, data.value)
		})
	}
}
