package main

import (
	"fmt"
	"os"
	"testing"
)

func TestWriter(t *testing.T) {
	// capture stdout
	stdout := os.Stdout
	reader, writer, _ := os.Pipe()
	os.Stdout = writer
	defer func() {
		os.Stdout = stdout
	}()
	defer writer.Close()
	defer reader.Close()
	// test writer
	gameWriter := NewWriter()
	gameWriter.Write(NewAction(1, 2, 0))
	var row, column int // x = column, y = row
	fmt.Fscan(reader, &row, &column)
	if row != 0 {
		t.Errorf("Expected row 0 but was %d", row)
	}
	if column != 2 {
		t.Errorf("Expected column 2 but was %d", column)
	}
}
