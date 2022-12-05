package main

import (
	"fardragon/aoc2022/v2/common"
	"fmt"
    "math"
    "strconv"
    "strings"
)

type stack []rune

func (s *stack) isEmpty() bool {
    return len(*s) == 0
}

func (s *stack) push(val rune) {
    *s = append(*s, val)
}

func (s *stack) pop() (rune, bool) {
    if s.isEmpty() {
        return ' ', false
    } else {
        index := len(*s) - 1
        element := (*s)[index]
        *s = (*s)[:index]
        return element, true
    }
}

func (s *stack) peek() (rune, bool) {
    if s.isEmpty() {
        return ' ', false
    } else {
        index := len(*s) - 1
        element := (*s)[index]
        return element, true
    }
}

type instruction struct {
    count int
    from int
    to int
}

type state struct {
    stacks []stack
    instructions []instruction
}

func parseInput(input []string) state {

    maxInitialStackSize := -1

    for index, line := range input {
        if (len(line) >= 2) && (line[1] == '1') {
            maxInitialStackSize = index
            break
        }
    }
    if maxInitialStackSize == -1 {
        panic("Invalid input")
    }

    stackDefinition := input[maxInitialStackSize]
    stackDefinitionSplit := strings.Split(stackDefinition, "")
    stackCount, err := strconv.Atoi(stackDefinitionSplit[len(stackDefinitionSplit) - 1])
    if err != nil {
        panic(err)
    }

    result := state{
        stacks: make([]stack, stackCount),
        instructions: make([]instruction, 0),
    }

    for i := maxInitialStackSize - 1; i >= 0; i-- {
        activeStacks := int(math.Ceil(float64(len(input[i]))  / 4.0))

        for j := 0; j < activeStacks; j++ {
            char := rune(input[i][j * 4 + 1])
            if char != ' ' {
                result.stacks[j].push(char)
            }
        }
    }

    if len(input[maxInitialStackSize + 1]) != 0 {
        panic("Invalid input")
    }

    for i := maxInitialStackSize + 2; i < len(input); i++ {
        line := input[i]
        lineSplit := strings.Split(line, " ")
        if len(lineSplit) != 6 {
            panic("Invalid input")
        }

        count, err := strconv.Atoi(lineSplit[1])
        if err != nil {
            panic(err)
        }

        from, err := strconv.Atoi(lineSplit[3])
        if err != nil {
            panic(err)
        }

        to, err := strconv.Atoi(lineSplit[5])
        if err != nil {
            panic(err)
        }
        result.instructions = append(result.instructions, instruction{
            count: count,
            from: from ,
            to: to,
        })
    }

	return result
}

func executeInstruction9000(stacks *[]stack, in instruction) {
    for i := 0; i < in.count; i++ {
        char, _ := (*stacks)[in.from - 1].pop()
        (*stacks)[in.to - 1].push(char)
    }
}
func solvePart1(input state) string {

    s := input.stacks
    for _, in := range input.instructions {
        executeInstruction9000(&s, in)
    }

    var result string
    for _, resStack := range s {
        char, _ := resStack.peek()
        result += string(char)
    }

	return result
}

func executeInstruction9001(stacks *[]stack, in instruction) {

    var buffer []rune
    for i := 0; i < in.count; i++ {
        char, present := (*stacks)[in.from - 1].pop()
        if !present {
            panic("Invalid input")
        }
        buffer=append(buffer, char)
    }

    for i := in.count - 1; i >= 0 ; i-- {
        (*stacks)[in.to - 1].push(buffer[i])
    }
}

func solvePart2(input state) string {
    s := input.stacks
    for _, in := range input.instructions {
        executeInstruction9001(&s, in)
    }

    var result string
    for _, resStack := range s {
        char, _ := resStack.peek()
        result += string(char)
    }

    return result
}

func main() {
	input := common.ReadInput("input.txt")
	parsedInput := parseInput(input)
    parsedInput2 := parseInput(input)


	result1 := solvePart1(parsedInput)
    result2 := solvePart2(parsedInput2)

	fmt.Printf("Part 1 result: %s\r\n", result1)
	fmt.Printf("Part 2 result: %s\r\n", result2)
}
