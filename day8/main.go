package main

import (
	"fardragon/aoc2022/v2/common"
	"fmt"
	"sort"
	"strconv"
)

type forest struct {
	trees  [][]int
	width  int
	height int
}

type direction struct {
	x int
	y int
}

var up = direction{0, -1}
var down = direction{0, 1}
var left = direction{-1, 0}
var right = direction{1, 0}

func parseInput(input []string) forest {

	width := len(input[0])
	height := len(input)

	result := forest{
		width:  width,
		height: height,
	}

	for y := 0; y < height; y++ {
		row := input[y]
		if len(row) != width {
			panic("Invalid input")
		}
		result.trees = append(result.trees, []int{})

		for x := 0; x < width; x++ {
			val, err := strconv.Atoi(string(row[x]))
			if err != nil {
				panic(err)
			}
			result.trees[y] = append(result.trees[y], val)
		}
	}

	return result
}

func isVisibleInDirection(x int, y int, dir direction, forest forest) bool {

	isVisible := true
	treeHeight := forest.trees[y][x]

	xx := x + dir.x
	yy := y + dir.y

	for (xx >= 0) && (xx < forest.width) && (yy >= 0) && (yy < forest.height) {
		if forest.trees[yy][xx] >= treeHeight {
			isVisible = false
			break
		}
		xx += dir.x
		yy += dir.y
	}

	return isVisible
}
func solvePart1(input forest) int {

	circumference := 2*input.height + 2*input.width - 4

	innerVisible := 0
	for y := 1; y < input.height-1; y++ {
		for x := 1; x < input.width-1; x++ {
			if isVisibleInDirection(x, y, up, input) {
				innerVisible += 1
				continue
			} else if isVisibleInDirection(x, y, down, input) {
				innerVisible += 1
				continue
			} else if isVisibleInDirection(x, y, left, input) {
				innerVisible += 1
				continue
			} else if isVisibleInDirection(x, y, right, input) {
				innerVisible += 1
				continue
			}
		}
	}

	return circumference + innerVisible
}

func viewDistanceInDirection(x int, y int, dir direction, forest forest) int {

	viewDistance := 0
	treeHeight := forest.trees[y][x]

	xx := x + dir.x
	yy := y + dir.y

	for (xx >= 0) && (xx < forest.width) && (yy >= 0) && (yy < forest.height) {
		viewDistance += 1
		if forest.trees[yy][xx] >= treeHeight {
			break
		}
		xx += dir.x
		yy += dir.y
	}

	return viewDistance
}

func solvePart2(input forest) int {

	var results []int

	for y := 0; y < input.height; y++ {
		for x := 0; x < input.width; x++ {

			score := viewDistanceInDirection(x, y, up, input)
			score *= viewDistanceInDirection(x, y, down, input)
			score *= viewDistanceInDirection(x, y, left, input)
			score *= viewDistanceInDirection(x, y, right, input)
			results = append(results, score)
		}
	}
	sort.Ints(results)
	return results[len(results)-1]
}

func main() {
	input := common.ReadInput("input.txt")
	parsedInput := parseInput(input)

	result1 := solvePart1(parsedInput)
	result2 := solvePart2(parsedInput)

	fmt.Printf("Part 1 result: %d\r\n", result1)
	fmt.Printf("Part 2 result: %d\r\n", result2)
}
