package main

import (
	"fardragon/aoc2022/v2/common"
	"fmt"
)

type symbol int

const (
	A symbol = iota
	B
	C
	X
	Y
	Z
)

type round struct {
	opponentSymbol symbol
	mySymbol       symbol
}

type roundResult int

const (
	win roundResult = iota
	draw
	lose
)

type hand int

const (
	rock hand = iota
	paper
	scissors
)

func parseInput(input []string) []round {

	var result []round
	for _, line := range input {
		if len(line) != 3 {
			panic("Invalid input")
		}

		var opponentSymbol symbol
		switch line[0] {
		case 'A':
			{
				opponentSymbol = A
			}
		case 'B':
			{
				opponentSymbol = B
			}
		case 'C':
			{
				opponentSymbol = C
			}
		default:
			{
				panic("Invalid input")
			}
		}

		var mySymbol symbol
		switch line[2] {
		case 'X':
			{
				mySymbol = X
			}
		case 'Y':
			{
				mySymbol = Y
			}
		case 'Z':
			{
				mySymbol = Z
			}
		default:
			{
				panic("Invalid input")
			}
		}
		result = append(result, round{opponentSymbol: opponentSymbol, mySymbol: mySymbol})
	}
	return result
}

func solveRound(opponent hand, me hand) roundResult {
	if opponent == me {
		return draw
	} else {
		switch opponent {
		case rock:
			{
				if me == paper {
					return win
				} else {
					return lose
				}
			}
		case paper:
			{
				if me == scissors {
					return win
				} else {
					return lose
				}
			}
		case scissors:
			{
				if me == rock {
					return win
				} else {
					return lose
				}
			}
		default:
			{
				panic("")
			}
		}
	}
}

func countPoints(result roundResult, hand hand) int {
    roundPoints := 0
    if result == win {
        roundPoints += 6
    } else if result == draw {
        roundPoints += 3
    }

    if hand == rock {
        roundPoints += 1
    } else if hand == paper {
        roundPoints += 2
    } else {
        roundPoints += 3
    }
    return roundPoints
}

func solvePart1(rounds []round) int {

	mapping := map[symbol]hand{
		A: rock,
		B: paper,
		C: scissors,
		X: rock,
		Y: paper,
		Z: scissors,
	}

	totalPoints := 0
	for _, r := range rounds {
		opponentsHand := mapping[r.opponentSymbol]
		myHand := mapping[r.mySymbol]
		result := solveRound(opponentsHand, myHand)
		totalPoints += countPoints(result, myHand)
	}
	return totalPoints
}

func solvePart2(rounds []round) int {

	mapping := map[symbol]hand{
		A: rock,
		B: paper,
		C: scissors,
	}

	totalPoints := 0
	for _, r := range rounds {
		opponentsHand := mapping[r.opponentSymbol]

		var myHand hand
		var result roundResult
		switch r.mySymbol {
		case X:
			{
				result = lose
				if opponentsHand == rock {
					myHand = scissors
				} else if opponentsHand == paper {
					myHand = rock
				} else {
					myHand = paper
				}
			}
		case Y:
			{
				result = draw
                myHand = opponentsHand
			}
		case Z:
			{
				result = win
				if opponentsHand == rock {
					myHand = paper
				} else if opponentsHand == paper {
					myHand = scissors
				} else {
					myHand = rock
				}
			}

		}
        totalPoints += countPoints(result, myHand)
	}
	return totalPoints
}

func main() {
	input := common.ReadInput("input.txt")
	parsedInput := parseInput(input)

	result1 := solvePart1(parsedInput)
    result2 := solvePart2(parsedInput)

	fmt.Printf("Part 1 result: %d\r\n", result1)
    fmt.Printf("Part 2 result: %d\r\n", result2)
}
