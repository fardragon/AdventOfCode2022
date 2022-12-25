package main

import "testing"

var testInput = []int{
    1,
    2,
    -3,
    3,
    -2,
    0,
    4,
}
func TestSolvePart1(t *testing.T) {
    testResult := solvePart1(testInput)

	if testResult != 3 {
		t.Errorf("Expected: %d got %d", 3, testResult)
	}
}

func TestSolvePart2(t *testing.T) {
    testResult := solvePart2(testInput)

    if testResult != 1623178306 {
        t.Errorf("Expected: %d got %d", 1623178306, testResult)
	}
}
