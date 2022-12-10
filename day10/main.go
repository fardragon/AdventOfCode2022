package main

import (
	"fardragon/aoc2022/v2/common"
	"fmt"
	"strconv"
	"strings"
)

type instruction interface {
	isInstruction()
}

type noopInstruction struct{}

func (noopInstruction) isInstruction() {}

type addInstruction struct {
	x int
}

func (addInstruction) isInstruction() {}

func parseInput(input []string) []instruction {

	var result []instruction

	for _, line := range input {
		if strings.HasPrefix(line, "noop") {
			result = append(result, noopInstruction{})
		} else if strings.HasPrefix(line, "addx") {
			split := strings.Split(line, " ")
			if len(split) != 2 {
				panic(fmt.Sprintf("Invalid input: %s", line))
			}

			val, err := strconv.Atoi(split[1])
			if err != nil {
				panic(err)
			}
			result = append(result, addInstruction{val})
		} else {
			panic(fmt.Sprintf("Invalid input: %s", line))
		}
	}

	return result
}

func executeInstructions(input []instruction) []int {
	x := 1
	var cycles []int

	for _, ins := range input {
		switch i := ins.(type) {
		case noopInstruction:
			{
				//                fmt.Println("noopInstruction{},")
				cycles = append(cycles, x)
			}
		case addInstruction:
			{
				//                fmt.Printf("addInstruction{%d},\r\n", i.x)
				cycles = append(cycles, x)
				cycles = append(cycles, x)
				x += i.x
			}
		}
	}
	return cycles
}
func solvePart1(input []instruction) int {

	cycles := executeInstructions(input)

	if len(cycles) < 220 {
		panic("Not enough cycles")
	}

	signal := 0
	for i := 0; i < 6; i++ {
		mCycle := 20 + i*40
		signal += mCycle * cycles[mCycle-1]
	}

	return signal
}

func solvePart2(input []instruction) []string {

	cycles := executeInstructions(input)

	var result []string

	for row := 0; row < 6; row++ {
		line := ""

		for pixel := 0; pixel < 40; pixel++ {
			cCycle := row*40 + pixel
			spritePosition := cycles[cCycle]

			if common.Abs(spritePosition-pixel) <= 1 {
				line += "#"
			} else {
				line += "."
			}
		}
		result = append(result, line)
	}

	return result
}

func main() {
	input := common.ReadInput("input.txt")
	parsedInput := parseInput(input)

	result1 := solvePart1(parsedInput)
	result2 := solvePart2(parsedInput)

	fmt.Printf("Part 1 result: %d\r\n", result1)

	fmt.Println("Part 2 result:")
	for _, line := range result2 {
		fmt.Println(line)
	}
}
