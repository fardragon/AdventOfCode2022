package main

import "testing"

var testSacks = []rucksack{
    {[]byte("vJrwpWtwJgWr"), []byte("hcsFMMfFFhFp")},
    {[]byte("jqHRNqRjqzjGDLGL"), []byte("rsFMfFZSrLrFZsSL")},
    {[]byte("PmmdzqPrV"), []byte("vPwwTWBwg")},
    {[]byte("wMqvLMZHhHMvwLH"), []byte("jbvcjnnSBnvTQFn")},
    {[]byte("ttgJtRGJ"), []byte("QctTZtZT")},
    {[]byte("CrZsJsPPZsGz"), []byte("wwsLwLmpwMDw")},
}

func TestSolvePart1(t *testing.T) {
    testResult := solvePart1(testSacks)

    if testResult != 157 {
        t.Errorf("Expected: %d got %d", 157, testResult)
    }
}

func TestSolvePart2(t *testing.T) {
    testResult := solvePart2(testSacks)

    if testResult != 70 {
        t.Errorf("Expected: %d got %d", 70, testResult)
    }
}
