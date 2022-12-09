package main

import "testing"

var testInstructions = []instruction{
	{
		right,
		4,
	},
	{
		up,
		4,
	},
	{
		left,
		3,
	},
	{
		down,
		1,
	},
	{
		right,
		4,
	},
	{
		down,
		1,
	},
	{
		left,
		5,
	},
	{
		right,
		2,
	},
}

var testInstructionsExtendeed = []instruction{
	{
		right,
		5,
	},
	{
		up,
		8,
	},
	{
		left,
		8,
	},
	{
		down,
		3,
	},
	{
		right,
		17,
	},
	{
		down,
		10,
	},
	{
		left,
		25,
	},
	{
		up,
		20,
	},
}

func TestSolvePart1(t *testing.T) {
	testResult := solvePart1(testInstructions)

	if testResult != 13 {
		t.Errorf("Expected: %d got %d", 13, testResult)
	}
}

func TestSolvePart2(t *testing.T) {
	testResult := solvePart2(testInstructions)

	if testResult != 1 {
		t.Errorf("Expected: %d got %d", 1, testResult)
	}

	testResultExtended := solvePart2(testInstructionsExtendeed)

	if testResultExtended != 36 {
		t.Errorf("Expected: %d got %d", 36, testResult)
	}
}
