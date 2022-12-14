package main

import "testing"

var testInstructions = instructions{
	instruction{
		{498, 4},
		{498, 6},
		{496, 6},
	},
	instruction{
		{503, 4},
		{502, 4},
		{502, 9},
		{494, 9},
	},
}

func TestSolvePart1(t *testing.T) {
    testResult := solvePart1(testInstructions)

	if testResult != 24 {
		t.Errorf("Expected: %d got %d", 24, testResult)
	}
}

func TestSolvePart2(t *testing.T) {
    testResult := solvePart2(testInstructions)

	if testResult != 93 {
		t.Errorf("Expected: %d got %d", 93, testResult)
	}
}
