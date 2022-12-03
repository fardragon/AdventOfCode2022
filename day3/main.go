package main

import (
	"fardragon/aoc2022/v2/common"
	"fmt"
)

type compartment []byte

type rucksack struct {
	compartmentA compartment
	compartmentB compartment
}

func parseInput(input []string) []rucksack {
	var result []rucksack

	for _, line := range input {
		size := len(line)
		if (size % 2) != 0 {
			panic("Invalid input")
		}

		result = append(result, rucksack{
			compartmentA: []byte(line[:size/2]),
			compartmentB: []byte(line[size/2:]),
		})
	}

	return result
}

func sliceToMap(input []byte) map[byte]bool {

    result := map[byte]bool{}
    for _, element := range input {
        result[element] = true
    }
    return result
}
func compartmentsIntersection(input rucksack) []byte {
	var result []byte

	set := sliceToMap(input.compartmentA)

	resultSet := map[byte]bool{}

	for _, element := range input.compartmentB {
		if set[element] {
			resultSet[element] = true
		}
	}

	for key := range resultSet {
		result = append(result, key)
	}

	return result
}
func getPriority(input byte) int {
    if (input >= 'a') && (input <= 'z') {
        return int(input - 'a') + 1
    } else if (input >= 'A') && (input <= 'Z') {
        return int(input - 'A') + 27
    } else {
        panic("Invalid input")
    }
}



func solvePart1(sacks []rucksack) int {
    result := 0
	for _, sack := range sacks {
		intersection := compartmentsIntersection(sack)
		if len(intersection) != 1 {
			panic("Invalid input")
		}
        result += getPriority(intersection[0])
	}

	return result
}

func solvePart2(sacks []rucksack) int {

    sackCount := len(sacks)

    if sackCount % 3 != 0 {
        panic("Invalid input")
    }

    result := 0
    for i := 0; i < sackCount / 3; i++ {
        setA := sliceToMap(append(sacks[3 * i].compartmentA, sacks[3 * i].compartmentB...))
        setB := sliceToMap(append(sacks[3 * i + 1].compartmentA, sacks[3 * i + 1].compartmentB...))
        setC := sliceToMap(append(sacks[3 * i + 2].compartmentA, sacks[3 * i + 2].compartmentB...))

        resultSet := map[byte]bool{}
        for element := range setA {
            if setB[element] && setC[element] {
                resultSet[element] = true
            }
        }

        var possbileBadges []byte
        for key := range resultSet {
            possbileBadges = append(possbileBadges, key)
        }

        if len(possbileBadges) != 1 {
            panic("Invalid input")
        }
        result += getPriority(possbileBadges[0])

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
