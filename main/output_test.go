package main

import (
	"testing"
)

func TestWriter(t *testing.T) {
	interceptor := NewStdoutInterceptor()
	defer interceptor.Close()
	interceptor.Intercept()
	gameWriter := NewWriter()
	gameWriter.Write(NewAction(1, 2, 0))
	var row, column int // x = column, y = row
	interceptor.Scan(&row, &column)
	if row != 0 {
		t.Errorf("Expected row 0 but was %d", row)
	}
	if column != 2 {
		t.Errorf("Expected column 2 but was %d", column)
	}
}
