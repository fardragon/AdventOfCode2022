package main

import (
	"fardragon/aoc2022/v2/common"
	"fmt"
)


func parseInput(input []string) string {
    if len(input) != 1 {
        panic("Invalid input")
    }
    return input[0]
}
func solvePart1(input string) int {

    for i:=3; i < len(input); i++ {
        unique := make(map[uint8]bool)

        unique[input[i - 0]] = true
        unique[input[i - 1]] = true
        unique[input[i - 2]] = true
        unique[input[i - 3]] = true
        if len(unique) == 4 {
            return i + 1
        }
    }

    panic("No solution")
}

func solvePart2(input string) int {
    for i := 13; i < len(input); i++ {
        unique := make(map[uint8]bool)

        for j := 0; j < 14; j++ {
            unique[input[i-j]] = true
        }
        if len(unique) == 14 {
            return i + 1
        }
    }

    panic("No solution")
}

func main() {
	input := common.ReadInput("input.txt")
	parsedInput := parseInput(input)
    parsedInput2 := parseInput(input)


	result1 := solvePart1(parsedInput)
    result2 := solvePart2(parsedInput2)

	fmt.Printf("Part 1 result: %d\r\n", result1)
	fmt.Printf("Part 2 result: %d\r\n", result2)
}
