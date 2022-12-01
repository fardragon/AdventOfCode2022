package main

import (
    "fmt"
    "sort"
)
import "fardragon/aoc2022/v2/common"
import "strconv"


type elf struct {
    calories []int
}

func parseInput(input []string) []elf {
    var result []elf

    var currentElf elf
    for _, line := range input {
        if len(line) == 0 {
            result = append(result, currentElf)
            currentElf = elf{}
        } else {
            value, err := strconv.Atoi(line)
            if err != nil {
                panic(err)
            }
            currentElf.calories = append(currentElf.calories, value)
        }
    }

    return result
}

func sumElf(elf elf) int {
    sum := 0
    for _, calories := range elf.calories {
        sum += calories
    }
    return sum
}

func solvePart1(input []elf) int {

    result := 0
    for _, elf := range input {
        sum := sumElf(elf)
        if sum > result {
            result = sum
        }
    }
    return result
}

func solvePart2(input []elf) int {


    sort.Slice(input, func(i int, j int) bool {
        return sumElf(input[i]) > sumElf(input[j])
    })

    return sumElf(input[0]) + sumElf(input[1]) + sumElf(input[2])
}


func main() {
    fmt.Println("hello world")

    input := common.ReadInput("input.txt")
    parsedInput := parseInput(input)

    result1 := solvePart1(parsedInput)
    result2 := solvePart2(parsedInput)


    fmt.Printf("Part 1 result: %d\r\n", result1)
    fmt.Printf("Part 2 result: %d\r\n", result2)

}
