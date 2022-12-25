package main

import (
    "fardragon/aoc2022/v2/common"
	"fmt"
    "strconv"
)


type codePoint struct {
    value int
    initialPosition int
}

func parseInput(input []string) []int {
	var result []int

	for _, line := range input {
		val, err := strconv.Atoi(line)
        if err != nil {
            panic(err)
        }
        result = append(result, val)
	}
	return result
}

func buildCode(input []int) []codePoint {
    var result []codePoint

    for ix, val := range input {
        result = append(result, codePoint{
            value:           val,
            initialPosition: ix,
        })
    }
    return result
}

func findElelementWithInitialPosition(code []codePoint, initialPosition int) int {
    for ix, val := range code {
        if val.initialPosition == initialPosition {
            return ix
        }
    }
    panic("Unrechable")
}

func moveElement(code []codePoint, index int, offset int) []codePoint {
    // Make a copy of the slice
    newSlice := make([]codePoint, len(code))
    copy(newSlice, code)

    // Remove the element at the specified index
    element := newSlice[index]
    newSlice = append(newSlice[:index], newSlice[index+1:]...)

    // Insert the element at the new index
    newIndex := (index + offset) % len(newSlice)
    if newIndex <= 0 {
        newIndex = len(newSlice) + newIndex
    }

    newSlice = append(newSlice[:newIndex], append([]codePoint{element}, newSlice[newIndex:]...)...)

    return newSlice
}

func mixCode(code *[]codePoint) {
    for i := 0; i < len(*code); i++ {
        currentPosition := findElelementWithInitialPosition(*code, i)
        *code = moveElement(*code, currentPosition, (*code)[currentPosition].value)
    }
}
func solvePart1(input []int) int {
    code := buildCode(input)

    mixCode(&code)
    startingPosition := -1
    for ix, val := range code {
        if val.value == 0 {
            startingPosition = ix
        }
    }

    if startingPosition == -1 {
        panic("Solving error")
    }

    first := code[(startingPosition + 1000) % len(code)].value
    second := code[(startingPosition + 2000) % len(code)].value
    third := code[(startingPosition + 3000) % len(code)].value

    return first + second + third
}

func solvePart2(input []int) int {
    code := buildCode(input)

    //Apply key
    for ix := range code {
        code[ix].value = code[ix].value * 811589153
    }

    for i := 0; i < 10; i++ {
        mixCode(&code)
    }

    startingPosition := -1
    for ix, val := range code {
        if val.value == 0 {
            startingPosition = ix
        }
    }

    if startingPosition == -1 {
        panic("Solving error")
    }

    first := code[(startingPosition + 1000) % len(code)].value
    second := code[(startingPosition + 2000) % len(code)].value
    third := code[(startingPosition + 3000) % len(code)].value

    return first + second + third
}

func main() {
	input := common.ReadInput("input.txt")
	parsedInput := parseInput(input)

	result1 := solvePart1(parsedInput)
	result2 := solvePart2(parsedInput)

	fmt.Printf("Part 1 result: %d\r\n", result1)
	fmt.Printf("Part 2 result: %d\r\n", result2)

}
