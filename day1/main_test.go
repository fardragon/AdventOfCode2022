package main

import "testing"

var testElves = []elf{{calories: []int{1000, 2000, 3000}}, {calories: []int{4000}}, {calories: []int{5000, 6000}}, {calories: []int{7000, 8000, 9000}}, {calories: []int{10000}}}

func TestSolvePart1(t *testing.T) {
    testResult := solvePart1(testElves)

    if testResult != 24000 {
        t.Errorf("Expected: %d got %d", 24000, testResult)
    }
}

func TestSolvePart2(t *testing.T) {
    testResult := solvePart2(testElves)

    if testResult != 45000 {
        t.Errorf("Expected: %d got %d", 45000, testResult)
    }
}
