package main

import (
	"fmt"
	"strconv"
	"strings"
)
import "fardragon/aoc2022/v2/common"

type point struct {
	x int
	y int
}

type boardElement int8

const (
	air boardElement = iota
	rock
	sand
)

type board struct {
	elements [][]boardElement
	bottom   int
}
type instruction []point
type instructions []instruction

func parsePoint(pointStr string) point {
	split := strings.Split(pointStr, ",")

	if len(split) != 2 {
		panic("Parsing error")
	}

	x, errX := strconv.Atoi(split[0])
	if errX != nil {
		panic(errX)
	}

	y, errY := strconv.Atoi(split[1])
	if errY != nil {
		panic(errY)
	}
	return point{x, y}
}

func determineBoardDimensions(ins instructions) int {

	maxDepth := 0

	for _, instruction := range ins {
		for _, point := range instruction {
			if point.y > maxDepth {
				maxDepth = point.y
			}
		}
	}

	return maxDepth
}

func min(a int, b int) int {
    if a < b {
        return a
    } else {
        return b
    }
}

func max(a int, b int) int {
    if a > b {
        return a
    } else {
        return b
    }
}
func getPointsBetween(a point, b point) []point {

    var result []point
    if a.x == b.x {
        for i:=min(a.y, b.y); i <= max(a.y, b.y); i++ {
            result = append(result, point{a.x, i})
        }

    } else if a.y == b.y {
        for i:=min(a.x, b.x); i <= max(a.x, b.x); i++ {
            result = append(result, point{i, a.y})
        }
    } else {
        panic("Invalid instruction")
    }
    return result
}

func buildBoard(ins instructions) board {

	depth := determineBoardDimensions(ins)

	result := board{
		elements: make([][]boardElement, depth + 3),
		bottom:   depth,
	}

	for ix := range result.elements {
		result.elements[ix] = make([]boardElement, 1000)
		for i := range result.elements[ix] {
			result.elements[ix][i] = air
		}
	}

	for _, in := range ins {
		for i := 0; i < len(in)-1; i++ {
			points := getPointsBetween(in[i], in[i+1])
            for _, point := range points {
                result.elements[point.y][point.x] = rock
            }
		}
	}

	return result
}

func parseInput(input []string) instructions {
	var result instructions

	for _, line := range input {
		split := strings.Split(line, " -> ")

		var ins instruction
		for _, insStr := range split {
			ins = append(ins, parsePoint(insStr))
		}
		result = append(result, ins)
	}
	return result
}

func moveSand(pos point, b board) point {

    if b.elements[pos.y + 1][pos.x]  == air {
        return point{pos.x, pos.y + 1}
    } else if b.elements[pos.y + 1][pos.x - 1] == air {
        return point{pos.x - 1, pos.y + 1}
    } else if b.elements[pos.y + 1][pos.x + 1] == air {
        return point{pos.x + 1, pos.y + 1}
    } else {
        return pos
    }

}
func solvePart1(ins instructions) int {

	b := buildBoard(ins)

    result := 0

    outsideLoop:
    for {
        sandPos := point{500, 0}
        b.elements[sandPos.y][sandPos.x] = sand

        for {
            newSandPos := moveSand(sandPos, b)
            if newSandPos == sandPos {
                result++
                break
            } else if newSandPos.y >= b.bottom {
                break outsideLoop
            } else {
                b.elements[sandPos.y][sandPos.x] = air
                b.elements[newSandPos.y][newSandPos.x] = sand
                sandPos = newSandPos
            }
        }
    }

	return result
}



func solvePart2(ins instructions) int {

    b := buildBoard(ins)

    for i := 0; i < 1000; i++ {
        b.elements[b.bottom + 2][i] = rock
    }


	result := 0

    outsideLoop:
        for {
            sandPos := point{500, 0}
            b.elements[sandPos.y][sandPos.x] = sand

            for {
                newSandPos := moveSand(sandPos, b)
                if newSandPos == sandPos {
                    result++
                    if newSandPos == (point{500, 0}) {
                        break outsideLoop
                    } else {
                        break
                    }
                } else if newSandPos.y >= b.bottom + 2 {
                    panic("Unreachable")
                } else {
                    b.elements[sandPos.y][sandPos.x] = air
                    b.elements[newSandPos.y][newSandPos.x] = sand
                    sandPos = newSandPos
                }
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
