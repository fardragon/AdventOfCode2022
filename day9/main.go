package main

import (
	"fardragon/aoc2022/v2/common"
	"fmt"
    "strconv"
	"strings"
)
type point struct {
	x int
	y int
}

var up = point{0, -1}
var down = point{0, 1}
var left = point{-1, 0}
var right = point{1, 0}

type instruction struct {
    direction point
    steps uint64
}

func parseInput(input []string) []instruction {

    var result []instruction

    for _, val := range input {
        split := strings.Split(val, " ")
        if len(split) != 2 {
            panic("Invalid input")
        }

        var ins instruction

        if split[0] == "U" {
            ins.direction = up
        } else if split[0] == "D" {
            ins.direction = down
        } else if split[0] == "L" {
            ins.direction = left
        } else if split[0] == "R" {
            ins.direction = right
        } else {
            panic("Invalid input")
        }

        steps, err := strconv.ParseUint(split[1], 10, 64)
        if err != nil {
            panic(err)
        }

        ins.steps = steps
        result = append(result, ins)
    }

    return result
}


func abs(a int) int {
    if a < 0 {
        return -a
    } else {
        return a
    }
}
func nodeChase(head point, tail point) point {

    distance := point{head.x - tail.x, head.y - tail.y}

    if abs(distance.x) <= 1 && abs(distance.y) <= 1 {
        //head and tail are adjacent, tail doesn't move
        return tail
    }

    direction := point{0, 0}
    if distance.x != 0 {
        direction.x = distance.x / abs(distance.x)
    }

    if distance.y != 0 {
        direction.y = distance.y / abs(distance.y)
    }

    return point{tail.x + direction.x, tail.y + direction.y}
}
func solvePart1(input []instruction) int {

    visitedPoints := common.NewSet[point]()


    head := point{0,0}
    tail := point{0,0}

    for _, instruction := range input {

        for i := uint64(0); i < instruction.steps; i++ {
//            fmt.Printf("Head: %+v Tail: %+v\r\n", head, tail)
            visitedPoints.Add(tail)

            head.x += instruction.direction.x
            head.y += instruction.direction.y

            tail = nodeChase(head, tail)
        }
    }
    visitedPoints.Add(tail)

    return visitedPoints.Len()
}

type longRope struct {
    nodes [10]point
}

func (rope *longRope) moveInDirection(direction point) {

    newNodes := rope.nodes

    //move head
    newNodes[0] = point{rope.nodes[0].x + direction.x, rope.nodes[0].y + direction.y}
    for i := 1; i < 10; i++ {
        newNodes[i] = nodeChase(newNodes[i - 1], newNodes[i])
    }


    rope.nodes = newNodes
}

func solvePart2(input []instruction) int {
    visitedPoints := common.NewSet[point]()


    rope := longRope{}


    for _, instruction := range input {

        for i := uint64(0); i < instruction.steps; i++ {
            visitedPoints.Add(rope.nodes[9])
            rope.moveInDirection(instruction.direction)
        }
    }
    visitedPoints.Add(rope.nodes[9])

    return visitedPoints.Len()
}

func main() {
	input := common.ReadInput("input.txt")
	parsedInput := parseInput(input)

	result1 := solvePart1(parsedInput)
	result2 := solvePart2(parsedInput)

	fmt.Printf("Part 1 result: %d\r\n", result1)
	fmt.Printf("Part 2 result: %d\r\n", result2)
}
