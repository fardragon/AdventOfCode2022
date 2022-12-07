package main

import "testing"
var testInputRaw = []string{
    "$ cd /",
    "$ ls",
    "dir a",
    "14848514 b.txt",
    "8504156 c.dat",
    "dir d",
    "$ cd a",
    "$ ls",
    "dir e",
    "29116 f",
    "2557 g",
    "62596 h.lst",
    "$ cd e",
    "$ ls",
    "584 i",
    "$ cd ..",
    "$ cd ..",
    "$ cd d",
    "$ ls",
    "4060174 j",
    "8033020 d.log",
    "5626152 d.ext",
    "7214296 k",
}

var testInput = parseInput(testInputRaw)

func TestSolvePart1(t *testing.T) {

    testResult := solvePart1(testInput)

    if testResult != 95437 {
        t.Errorf("Expected: %d got %d", 95437, testResult)
    }
}

func TestSolvePart2(t *testing.T) {
    testResult := solvePart2(testInput)

    if testResult != 24933642 {
        t.Errorf("Expected: %d got %d", 24933642, testResult)
    }
}
