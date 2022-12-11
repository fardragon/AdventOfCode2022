package main

import "testing"

var testMonkeys = []monkey{
	{
		items:     []uint{79, 98},
		op:        multiplyIntInstruction{19},
		divisor:   23,
		receivers: [2]uint{2, 3},
	},
	{
		items:     []uint{54, 65, 75, 74},
		op:        addIntInstruction{6},
		divisor:   19,
		receivers: [2]uint{2, 0},
	},
	{
		items:     []uint{79, 60, 97},
		op:        squareInstruction{},
		divisor:   13,
		receivers: [2]uint{1, 3},
	},
	{
		items:     []uint{74},
		op:        addIntInstruction{3},
		divisor:   17,
		receivers: [2]uint{0, 1},
	},
}

func TestSolvePart1(t *testing.T) {

    testInput := make([]monkey, len(testMonkeys))
    copy(testInput, testMonkeys)

    testResult := solvePart1(testInput)

	if testResult != 10605 {
		t.Errorf("Expected: %d got %d", 10605, testResult)
	}
}

func TestSolvePart2(t *testing.T) {

    testInput := make([]monkey, len(testMonkeys))
    copy(testInput, testMonkeys)

    testResult := solvePart2(testInput)

    if testResult != 2713310158 {
        t.Errorf("Expected: %d got %d", 2713310158, testResult)
	}
}
