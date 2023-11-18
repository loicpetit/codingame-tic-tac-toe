package main

import (
	"fmt"
	"os"
	"testing"
)

func TestScan(t *testing.T) {
	read, write, _ := os.Pipe()
	defer read.Close()
	defer write.Close()
	var val1, val2, val3 int
	done := make(chan bool)
	go func() {
		fmt.Fscan(read, &val1)
		fmt.Fscan(read, &val2, &val3)
		done <- true
	}()
	write.WriteString("1\n")
	write.WriteString("2 3\n")
	<-done
	expectedVal1 := 1
	expectedVal2 := 2
	expectedVal3 := 3
	if expectedVal1 != val1 {
		t.Errorf("val1: expected %d but was %d", expectedVal1, val1)
	}
	if expectedVal2 != val2 {
		t.Errorf("val2: expected %d but was %d", expectedVal2, val2)
	}
	if expectedVal3 != val3 {
		t.Errorf("val3: expected %d but was %d", expectedVal3, val3)
	}
}
