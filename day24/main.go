package main

import (
	"fardragon/aoc2022/v2/common"
	"fmt"
)


type point struct {
    y int
    x int
}

type blizzard struct {
    position point
    direction point
}

type grove struct {
    width int
    height int
    start point
    end point
    walls []point
    blizzards []blizzard
}

type searchState struct {
    time int
    next point
}

func mod(x, y int) int {
    if x < 0 {
        return (x % y + y) % y
    }
    return x % y
}
func parseInput(input []string) grove {
    result := grove{
        width:  len(input[0]) - 2,
        height: len(input) - 2,
        start:  point{},
        end:    point{},
        walls:  []point{},
        blizzards: []blizzard{},
    }


	for row, line := range input {
		for column, c := range line {
			if c == '.' {
                if row == 0 {
                    result.start = point {row, column}
                } else if row == (len(input) - 1) {
                    result.end = point {row, column}
                }
            } else if c == '#' {
                result.walls = append(result.walls, point{row, column})
            } else if c == '^' {
                result.blizzards = append(result.blizzards, blizzard{point{row, column}, point{-1, 0}})
            } else if c == 'v' {
                result.blizzards = append(result.blizzards, blizzard{point{row, column}, point{1, 0}})
            } else if c == '<' {
                result.blizzards = append(result.blizzards, blizzard{point{row, column}, point{0, -1}})
            } else if c == '>' {
                result.blizzards = append(result.blizzards, blizzard{point{row, column}, point{0, 1}})
            }
		}

	}

    result.walls = append(result.walls, point{result.start.y - 1, result.start.x})
	return result
}

func precomputeBlockedSquares(input grove) []common.Set[point] {

    combinations := input.width * input.height
    result := make([]common.Set[point], combinations)

    for time := 0; time < combinations; time++ {
        result[time] = common.NewSet[point]()

        for _, b := range input.blizzards {
            result[time].Add(point{
                1 + mod(b.position.y - 1 + b.direction.y * time, input.height),
                1 + mod(b.position.x - 1 + b.direction.x * time, input.width),
                })
        }

        for _, sq := range input.walls {
            result[time].Add(sq)
        }
    }

    return result



}
func moveCandidates(current point, blocked common.Set[point]) []point {

    candidates := []point{
        {current.y + 1, current.x},
        {current.y, current.x + 1},
        {current.y, current.x - 1},
        {current.y - 1, current.x},
        current,
    }

    var result []point
    for _, c := range candidates {
        if !blocked.Contains(c) {
            result = append(result, c)
        }
    }
    return result
}

func findPath(start point, end point, startTime int, blockedSquares []common.Set[point]) int {

    explored := common.NewSet[searchState]()
    queue := []searchState{{startTime, start}}

    for len(queue) > 0 {
        time := queue[0].time
        square := queue[0].next
        queue = queue[1:]

        time += 1

        candidates := moveCandidates(square, blockedSquares[time % len(blockedSquares)])

        for _, next := range candidates {
            state := searchState{time, next}
            if !explored.Contains(state) {
                if next == end {
                    return time
                }
                explored.Add(state)
                queue = append(queue, state)
            }
        }
    }
    panic("Unreachable")
}
func solvePart1(input grove) int {

    blocked := precomputeBlockedSquares(input)
    return findPath(input.start, input.end, 0, blocked)
}

func solvePart2(input grove) int {
    blocked := precomputeBlockedSquares(input)

    p1 := findPath(input.start, input.end, 0, blocked)
    p2 := findPath(input.end, input.start, p1, blocked)
    p3 := findPath(input.start, input.end, p2, blocked)

    return p3
}

func main() {
	input := common.ReadInput("input.txt")
	parsedInput := parseInput(input)

	result1 := solvePart1(parsedInput)
	result2 := solvePart2(parsedInput)

	fmt.Printf("Part 1 result: %d\r\n", result1)
	fmt.Printf("Part 2 result: %d\r\n", result2)
}
