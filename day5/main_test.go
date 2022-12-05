package main

import "testing"

//	type state struct {
//	   stacks []stack
//	   instructions []instruction
//	}
var testState = state{
	stacks: []stack{
		{
			'Z',
			'N',
		},
		{
			'M',
			'C',
			'D',
		},
		{
			'P',
		},
	},
	instructions: []instruction{
		{
			count: 1,
			from:  2,
			to:    1,
		},
		{
			count: 3,
			from:  1,
			to:    3,
		},
		{
			count: 2,
			from:  2,
			to:    1,
		},
		{
			count: 1,
			from:  1,
			to:    2,
		},
	},
}

func TestSolvePart1(t *testing.T) {
	testResult := solvePart1(testState)

	if testResult != "CMZ" {
		t.Errorf("Expected: %s got %s", "CMZ", testResult)
	}
}

func TestSolvePart2(t *testing.T) {
	testResult := solvePart2(testState)

	if testResult != "MCD" {
        t.Errorf("Expected: %s got %s", "MCD", testResult)
	}
}
