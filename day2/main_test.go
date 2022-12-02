package main

import "testing"

var testRounds = []round{{A, Y}, {B, X}, {C, Z}}

func TestSolvePart1(t *testing.T) {
    testResult := solvePart1(testRounds)

    if testResult != 15 {
        t.Errorf("Expected: %d got %d", 15, testResult)
    }
}

func TestSolvePart2(t *testing.T) {
    testResult := solvePart2(testRounds)

    if testResult != 12 {
        t.Errorf("Expected: %d got %d", 12, testResult)
    }
}
