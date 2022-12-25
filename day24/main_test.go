package main

import "testing"


var testInputStr  = []string{
    "#.######",
    "#>>.<^<#",
    "#.<..<<#",
    "#>v.><>#",
    "#<^v^^>#",
    "######.#",
}


func TestSolvePart1(t *testing.T) {

    parsedInput := parseInput(testInputStr)
    testResult := solvePart1(parsedInput)

    if testResult != 18 {
        t.Errorf("Expected: %d got %d", 18, testResult)
    }
}

func TestSolvePart2(t *testing.T) {
    parsedInput := parseInput(testInputStr)
    testResult := solvePart2(parsedInput)

    if testResult != 54 {
        t.Errorf("Expected: %d got %d", 54, testResult)
    }
}
