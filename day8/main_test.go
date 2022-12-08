package main

import "testing"

var testForest = forest {
    trees: [][]int{
       {3, 0, 3, 7, 3},
       {2, 5, 5, 1, 2},
       {6, 5, 3, 3, 2},
       {3, 3, 5, 4, 9},
       {3, 5, 3, 9, 0},
    },
    width: 5,
    height: 5,
}

func TestSolvePart1(t *testing.T) {
    testResult := solvePart1(testForest)

    if testResult != 21 {
        t.Errorf("Expected: %d got %d", 21, testResult)
    }
}

func TestSolvePart2(t *testing.T) {
    testResult := solvePart2(testForest)

    if testResult != 8 {
        t.Errorf("Expected: %d got %d", 8, testResult)
    }
}
