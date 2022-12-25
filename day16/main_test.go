package main

import "testing"

var testInput = []room{
    {
        name:     "AA",
        pressure: 0,
        exits:    []string{"DD", "II", "BB"},
    },
    {
        name:     "BB",
        pressure: 13,
        exits:    []string{"CC", "AA"},
    },
    {
        name:     "CC",
        pressure: 2,
        exits:    []string{"DD", "BB"},
    },
    {
        name:     "DD",
        pressure: 20,
        exits:    []string{"CC", "AA", "EE"},
    },
    {
        name:     "EE",
        pressure: 3,
        exits:    []string{"FF", "DD"},
    },
    {
        name:     "FF",
        pressure: 0,
        exits:    []string{"EE", "GG"},
    },
    {
        name:     "GG",
        pressure: 0,
        exits:    []string{"FF", "HH"},
    },
    {
        name:     "HH",
        pressure: 22,
        exits:    []string{"GG"},
    },
    {
        name:     "II",
        pressure: 0,
        exits:    []string{"AA", "JJ"},
    },
    {
        name:     "JJ",
        pressure: 21,
        exits:    []string{"II"},
    },
}

func TestSolvePart1(t *testing.T) {
    testResult := solvePart1(testInput)

    if testResult != 1651 {
        t.Errorf("Expected: %d got %d", 1651, testResult)
	}
}

func TestSolvePart2(t *testing.T) {
    testResult := solvePart2(testInput)

    if testResult != 1707 {
        t.Errorf("Expected: %d got %d", 1707, testResult)
	}
}
