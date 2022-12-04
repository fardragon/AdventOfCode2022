package main

import (
	"fardragon/aoc2022/v2/common"
	"fmt"
	"strconv"
	"strings"
)

type elf struct {
	start int
	end   int
}

type elfPair struct {
	a elf
	b elf
}

func parseElf(input string) elf {
	dashSplit := strings.Split(input, "-")
	if len(dashSplit) != 2 {
		panic("Invalid input")
	}

	min, minErr := strconv.Atoi(dashSplit[0])
	max, maxErr := strconv.Atoi(dashSplit[1])

	if minErr != nil || maxErr != nil {
		panic("Invalid input")
	}

	return elf{
		min,
		max,
	}
}
func parseInput(input []string) []elfPair {
	var result []elfPair
	for _, line := range input {
		commaSplit := strings.Split(line, ",")
		if len(commaSplit) != 2 {
			panic("Invalid input")
		}
		result = append(result, elfPair{
			a: parseElf(commaSplit[0]),
			b: parseElf(commaSplit[1]),
		})
	}

	return result
}

func solvePart1(pairs []elfPair) int {
	result := 0

	for _, pair := range pairs {
		if ((pair.a.start >= pair.b.start) && (pair.a.end <= pair.b.end)) ||
			((pair.b.start >= pair.a.start) && (pair.b.end <= pair.a.end)) {
            result += 1
		}
	}

	return result
}

func solvePart2(pairs []elfPair) int {
    result := 0

    for _, pair := range pairs {
        if (pair.a.start <= pair.b.end) && (pair.a.end >= pair.b.start)  {
            result += 1
            }
    }

    return result
}

func main() {
	input := common.ReadInput("input.txt")
	parsedInput := parseInput(input)

	result1 := solvePart1(parsedInput)
	result2 := solvePart2(parsedInput)

	fmt.Printf("Part 1 result: %d\r\n", result1)
	fmt.Printf("Part 2 result: %d\r\n", result2)
}
