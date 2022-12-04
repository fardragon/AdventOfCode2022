package main

import "testing"

var testPairs = []elfPair{
    {elf{2, 4}, elf{6, 8}},
    {elf{2, 3}, elf{4, 5}},
    {elf{5, 7}, elf{7, 9}},
    {elf{2, 8}, elf{3, 7}},
    {elf{6, 6}, elf{4, 6}},
    {elf{2, 6}, elf{4, 8}},
}

func TestSolvePart1(t *testing.T) {
    testResult := solvePart1(testPairs)

    if testResult != 2 {
        t.Errorf("Expected: %d got %d", 2, testResult)
    }
}

func TestSolvePart2(t *testing.T) {
    testResult := solvePart2(testPairs)

    if testResult != 4 {
        t.Errorf("Expected: %d got %d", 4, testResult)
    }
}
