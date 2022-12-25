package main

import "testing"

var testInput = ">>><<><>><<<>><>>><<<>>><<<><<<>><>><<>>"

func TestSolvePart1(t *testing.T) {

    testJest := parseInput([]string{testInput})
    testResult := solvePart1(testJest)

	if testResult != 3068 {
		t.Errorf("Expected: %d got %d", 3068, testResult)
	}
}

func TestSolvePart2(t *testing.T) {
    testJest := parseInput([]string{testInput})
    testResult := solvePart2(testJest)

    if testResult != 1514285714288 {
        t.Errorf("Expected: %d got %d", 1514285714288, testResult)
	}
}
