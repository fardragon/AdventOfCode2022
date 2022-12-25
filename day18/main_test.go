package main

import "testing"

var testInput = []point{
	{2, 2, 2},
	{1, 2, 2},
	{3, 2, 2},
	{2, 1, 2},
	{2, 3, 2},
	{2, 2, 1},
	{2, 2, 3},
	{2, 2, 4},
	{2, 2, 6},
	{1, 2, 5},
	{3, 2, 5},
	{2, 1, 5},
	{2, 3, 5},
}

func TestSolvePart1(t *testing.T) {

	testResult := solvePart1(testInput)

	if testResult != 64 {
		t.Errorf("Expected: %d got %d", 64, testResult)
	}
}

func TestSolvePart2(t *testing.T) {

	testResult := solvePart2(testInput)

	if testResult != 58 {
		t.Errorf("Expected: %d got %d", 58, testResult)
	}
}
