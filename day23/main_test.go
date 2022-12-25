package main

import "testing"

var testInput = []string{
	"..............",
	"..............",
	".......#......",
	".....###.#....",
	"...#...#.#....",
	"....#...##....",
	"...#.###......",
	"...##.#.##....",
	"....#..#......",
	"..............",
	"..............",
	"..............",
}

func TestSolvePart1(t *testing.T) {
    parsedInout := parseInput(testInput)
    testResult := solvePart1(parsedInout)

    if testResult != 110 {
        t.Errorf("Expected: %d got %d", 110, testResult)
	}
}

func TestSolvePart2(t *testing.T) {
    parsedInout := parseInput(testInput)
    testResult := solvePart2(parsedInout)

    if testResult != 20 {
        t.Errorf("Expected: %d got %d", 20, testResult)
	}
}
