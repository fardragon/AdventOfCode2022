package main

import "testing"

var testInput = []bluerpint{
    {4, 2, intPair{3, 14}, intPair{2, 7}},
    {2, 3, intPair{3, 8},  intPair{3, 12}},
}

func TestSolvePart1(t *testing.T) {

	testResult := solvePart1(testInput)

	if testResult != 33 {
		t.Errorf("Expected: %d got %d", 33, testResult)
	}
}

func TestSolvePart2(t *testing.T) {

	testResult := solvePart2(testInput)

	if testResult != 3472 {
        t.Errorf("Expected: %d got %d", 3472, testResult)
	}
}
