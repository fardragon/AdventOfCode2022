package main

import (
	"fardragon/aoc2022/v2/common"
	"fmt"
    "math"
)

type position struct {
	y int
	x int
}

type bounds struct {
    left int
    right int
    top int
    bottom int
}

type moveCheckFunc func(pos position, positions common.Set[position]) bool
type moveFunc func(pos position) position


func parseInput(input []string) common.Set[position] {
	result := common.Set[position]{}
	for row, line := range input {
		for column, c := range line {
			if c == '#' {
				result.Add(position{row, column})
			}
		}

	}
	return result
}

func canMoveNorth(pos position, positions common.Set[position]) bool {
	return !(positions.Contains(position{pos.y - 1, pos.x - 1}) ||
		positions.Contains(position{pos.y - 1, pos.x}) ||
		positions.Contains(position{pos.y - 1, pos.x + 1}))
}

func moveNorth(pos position) position {
    return position{pos.y - 1, pos.x }
}

func canMoveSouth(pos position, positions common.Set[position]) bool {
	return !(positions.Contains(position{pos.y + 1, pos.x - 1}) ||
		positions.Contains(position{pos.y + 1, pos.x}) ||
		positions.Contains(position{pos.y + 1, pos.x + 1}))
}

func moveSouth(pos position) position {
    return position{pos.y + 1, pos.x }
}
func canMoveEast(pos position, positions common.Set[position]) bool {
	return !(positions.Contains(position{pos.y + 1, pos.x + 1}) ||
		positions.Contains(position{pos.y, pos.x + 1}) ||
		positions.Contains(position{pos.y - 1, pos.x + 1}))
}

func moveEast(pos position) position {
    return position{pos.y, pos.x + 1}
}
func canMoveWest(pos position, positions common.Set[position]) bool {
	return !(positions.Contains(position{pos.y + 1, pos.x - 1}) ||
		positions.Contains(position{pos.y, pos.x - 1}) ||
		positions.Contains(position{pos.y - 1, pos.x - 1}))
}

func moveWest(pos position) position {
    return position{pos.y, pos.x - 1}
}

func getMoveCheckFunctions() ([4]moveCheckFunc,[4]moveFunc) {
	return [4]moveCheckFunc{
		canMoveNorth,
		canMoveSouth,
		canMoveWest,
		canMoveEast,
	},
    [4]moveFunc{
        moveNorth,
        moveSouth,
        moveWest,
        moveEast,
    }
}

func shouldMove(pos position, positions common.Set[position]) bool {
	return !(canMoveNorth(pos, positions) && canMoveSouth(pos, positions) &&
		canMoveWest(pos, positions) && canMoveEast(pos, positions))
}

func findBounds(positions common.Set[position]) bounds {

    left := math.MaxInt
    right := math.MinInt
    top := math.MaxInt
    bottom := math.MinInt

    for pos := range positions {
        if pos.x < left {
            left = pos.x
        } else if pos.x > right {
            right = pos.x
        }

        if pos.y < top {
            top = pos.y
        } else if pos.y > bottom {
            bottom = pos.y
        }
    }

    return bounds{
        left:   left,
        right:  right,
        top:    top,
        bottom: bottom,
    }
}

func solvePart1(input common.Set[position]) int {
    elves := input.Clone()
    moveCheckFuncs, moveFuncs := getMoveCheckFunctions()

    for round := 0; round < 10; round++ {

        proposedMoves := make(map[position][]position)

        for elf := range elves {
            if !(shouldMove(elf, elves)) {
                continue
            }

            moveCheckLoop:
            for j := 0; j < len(moveCheckFuncs); j++ {
                ix := (j + round) % len(moveCheckFuncs)
                if moveCheckFuncs[ix](elf, elves) {
                    newPos := moveFuncs[ix](elf)

                    _, contains := proposedMoves[newPos]
                    if contains {
                        proposedMoves[newPos] = append(proposedMoves[newPos], elf)
                    } else {
                        proposedMoves[newPos] = []position{elf}
                    }
                    break moveCheckLoop
                }
            }
        }

        for dst, src := range proposedMoves {
            if len(src) > 1 {
                continue
            }
            elves.Remove(src[0])
            elves.Add(dst)
        }
    }

    b := findBounds(elves)

    area := (b.right - b.left + 1) * (b.bottom - b.top + 1)

    return area - len(elves)
}

func solvePart2(input common.Set[position]) int {
    elves := input.Clone()
    moveCheckFuncs, moveFuncs := getMoveCheckFunctions()

    round := 0
    for {

        proposedMoves := make(map[position][]position)

        for elf := range elves {
            if !(shouldMove(elf, elves)) {
                continue
            }

            moveCheckLoop:
                for j := 0; j < len(moveCheckFuncs); j++ {
                    ix := (j + round) % len(moveCheckFuncs)
                    if moveCheckFuncs[ix](elf, elves) {
                        newPos := moveFuncs[ix](elf)

                        _, contains := proposedMoves[newPos]
                        if contains {
                            proposedMoves[newPos] = append(proposedMoves[newPos], elf)
                        } else {
                            proposedMoves[newPos] = []position{elf}
                        }
                        break moveCheckLoop
                    }
                }
        }

        if len(proposedMoves) == 0 {
            break
        }

        for dst, src := range proposedMoves {
            if len(src) > 1 {
                continue
            }
            elves.Remove(src[0])
            elves.Add(dst)
        }

        round++
    }

    return round + 1
}

func main() {
	input := common.ReadInput("input.txt")
	parsedInput := parseInput(input)

	result1 := solvePart1(parsedInput)
	result2 := solvePart2(parsedInput)

	fmt.Printf("Part 1 result: %d\r\n", result1)
	fmt.Printf("Part 2 result: %d\r\n", result2)
}
