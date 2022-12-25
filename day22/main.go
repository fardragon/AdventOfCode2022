package main

import (
	"fardragon/aoc2022/v2/common"
	"fmt"
	"strconv"
	"strings"
)

type coord struct {
	y int
	x int
}

type instructionType int8

const (
	rotateLeft instructionType = iota
	rotateRight
	forward
)

type instruction struct {
	insType instructionType
	value   int
}
type direction int8

const (
	right direction = iota
	down
	left
	up
)

type state struct {
	c   coord
	dir direction
}

type edge struct {
	face     string
	rotation int
}

type topology struct {
	edges     map[string][4]edge
	blockSize int
	shape     map[string]coord
}

func parseInput(input []string) ([]string, []instruction) {

	var resultMap []string
	var resultIns []instruction

	ix := 0

	for len(input[ix]) != 0 {
		resultMap = append(resultMap, input[ix])
		ix++
	}

	// skip separator
	ix++

	instructionString := strings.ReplaceAll(strings.ReplaceAll(input[ix], "L", " L "), "R", " R ")

	instructionSplit := strings.Split(instructionString, " ")

	for _, s := range instructionSplit {
		if s == "L" {
			resultIns = append(resultIns, instruction{rotateLeft, 0})
		} else if s == "R" {
			resultIns = append(resultIns, instruction{rotateRight, 0})
		} else {
			val, err := strconv.Atoi(s)
			if err != nil {
				panic(err)
			}
			resultIns = append(resultIns, instruction{forward, val})
		}

	}
	return resultMap, resultIns
}

func solve(monkeyMap []string, instructions []instruction, t topology) int {

	s := state{
		c:   coord{1, 51},
		dir: right,
	}

	for _, ins := range instructions {
		switch ins.insType {
		case rotateLeft:
			s.dir = (s.dir + 3) % 4
		case rotateRight:
			s.dir = (s.dir + 1) % 4
		case forward:
			for i := 0; i < ins.value; i++ {
				nextState := step(s, t)
				if monkeyMap[nextState.c.y][nextState.c.x] == '.' {
					s = nextState
				} else {
					break
				}
			}
		}
	}

	return 1000*(s.c.y+1) + 4*(s.c.x+1) + int(s.dir)
}

func wrapsAround(c coord, blockSize int) bool {
	return c.y < 0 || c.y >= blockSize || c.x < 0 || c.x >= blockSize
}
func step(s state, t topology) state {

	srcBlock := ""
	for key, val := range t.shape {
		localC := coord{s.c.y - val.y, s.c.x - val.x}
		if !wrapsAround(localC, t.blockSize) {
			srcBlock = key
			break
		}
	}

	if srcBlock == "" {
		panic("Could not determine src block")
	}

	dstBlock := srcBlock

	c := s.c
	dir := s.dir

	c.y -= t.shape[srcBlock].y
	c.x -= t.shape[srcBlock].x

	switch dir {
	case left:
		c.x -= 1
	case right:
		c.x += 1
	case up:
		c.y -= 1
	case down:
		c.y += 1
	default:
		panic("Invalid direction")
	}

	if wrapsAround(c, t.blockSize) {
		part := t.edges[srcBlock]
		neighbour := part[dir]

		dstBlock = neighbour.face
		rotation := neighbour.rotation

		c.y = (c.y + t.blockSize) % t.blockSize
		c.x = (c.x + t.blockSize) % t.blockSize

		for i := 0; i < rotation; i++ {
			newC := coord{
				y: c.x,
				x: t.blockSize - c.y - 1,
			}
			dir = (dir + 1) % 4
			c = newC
		}
	}

	c.y += t.shape[dstBlock].y
	c.x += t.shape[dstBlock].x
	return state{c, dir}
}

func main() {
	input := common.ReadInput("input.txt")
	monkeyMap, instructions := parseInput(input)

	inputBlockSize := 50
	var inputShape = map[string]coord{
		"A": {0, inputBlockSize},
		"B": {0, 2 * inputBlockSize},
		"C": {inputBlockSize, inputBlockSize},
		"D": {2 * inputBlockSize, 0},
		"E": {2 * inputBlockSize, inputBlockSize},
		"F": {3 * inputBlockSize, 0},
	}

	part1Topology := topology{
		edges: map[string][4]edge{
			"A": {{"B", 0}, {"C", 0}, {"B", 0}, {"E", 0}},
			"B": {{"A", 0}, {"B", 0}, {"A", 0}, {"B", 0}},
			"C": {{"C", 0}, {"E", 0}, {"C", 0}, {"A", 0}},
			"D": {{"E", 0}, {"F", 0}, {"E", 0}, {"F", 0}},
			"E": {{"D", 0}, {"A", 0}, {"D", 0}, {"C", 0}},
			"F": {{"F", 0}, {"D", 0}, {"F", 0}, {"D", 0}},
		},
		blockSize: inputBlockSize,
		shape:     inputShape,
	}

	part2Topology := topology{
		edges: map[string][4]edge{
			"A": {{"B", 0}, {"C", 0}, {"D", 2}, {"F", 1}},
			"B": {{"E", 2}, {"C", 1}, {"A", 0}, {"F", 0}},
			"C": {{"B", 3}, {"E", 0}, {"D", 3}, {"A", 0}},
			"D": {{"E", 0}, {"F", 0}, {"A", 2}, {"C", 1}},
			"E": {{"B", 2}, {"F", 1}, {"D", 0}, {"C", 0}},
			"F": {{"E", 3}, {"B", 0}, {"A", 3}, {"D", 0}},
		},
		blockSize: inputBlockSize,
		shape:     inputShape,
	}

	result1 := solve(monkeyMap, instructions, part1Topology)
	result2 := solve(monkeyMap, instructions, part2Topology)

	fmt.Printf("Part 1 result: %d\r\n", result1)
	fmt.Printf("Part 2 result: %d\r\n", result2)
}
