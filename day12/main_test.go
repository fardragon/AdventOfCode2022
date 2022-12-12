package main

import "testing"

var testPuzzle = puzzle{
	hMap: heightmap{
		heights: [][]int{
			{0, 0, 1, 16, 15, 14, 13, 12},
			{0, 1, 2, 17, 24, 23, 23, 11},
			{0, 2, 2, 18, 25, 25, 23, 10},
			{0, 2, 2, 19, 20, 21, 22, 9},
			{0, 1, 3, 4, 5, 6, 7, 8},
		},
		height: 5,
		width:  8,
	},
	start: point{0, 0},
	end:   point{5, 2},
}

func TestSolvePart1(t *testing.T) {
    testResult := solvePart1(testPuzzle)

	if testResult != 31 {
		t.Errorf("Expected: %d got %d", 31, testResult)
	}
}

func TestSolvePart2(t *testing.T) {
    testResult := solvePart2(testPuzzle)

    if testResult != 29 {
        t.Errorf("Expected: %d got %d", 29, testResult)
    }
}
