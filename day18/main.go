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
	z int
}
type direction int8

const (
	left direction = iota
	right
	up
	down
	forward
	backward
)

var directions = []direction{
	left,
	right,
	up,
	down,
	forward,
	backward,
}

type face struct {
	location point
	dir      direction
}

func parseCoord(input string) int {
	coord, err := strconv.Atoi(input)
	if err != nil {
		panic(err)
	}
	return coord
}
func parseInput(input []string) []point {
	var result []point

	for _, line := range input {

		split := strings.Split(line, ",")
		if len(split) != 3 {
			panic("Parsing error")
		}

		// Add 1 to every coordinate to move cubes away from 0,
		// this doesn't affect the result and makes it easier to implement flood fill in part 2
		result = append(result, point{
			x: parseCoord(split[0]) + 1,
			y: parseCoord(split[1]) + 1,
			z: parseCoord(split[2]) + 1,
		})
	}
	return result
}

func buildFaces(cubes []point) common.Set[face] {

	result := common.Set[face]{}

	for _, cube := range cubes {
		for _, dir := range directions {
			result.Add(face{
				location: cube,
				dir:      dir,
			})
		}
	}

	return result
}

func getOppositeFace(f face) face {
	switch f.dir {
	case left:
		return face{
			location: point{
				x: f.location.x - 1,
				y: f.location.y,
				z: f.location.z,
			},
			dir: right,
		}
	case right:
		return face{
			location: point{
				x: f.location.x + 1,
				y: f.location.y,
				z: f.location.z,
			},
			dir: left,
		}
	case up:
		return face{
			location: point{
				x: f.location.x,
				y: f.location.y + 1,
				z: f.location.z,
			},
			dir: down,
		}
	case down:
		return face{
			location: point{
				x: f.location.x,
				y: f.location.y - 1,
				z: f.location.z,
			},
			dir: up,
		}
	case forward:
		return face{
			location: point{
				x: f.location.x,
				y: f.location.y,
				z: f.location.z + 1,
			},
			dir: backward,
		}
	case backward:
		return face{
			location: point{
				x: f.location.x,
				y: f.location.y,
				z: f.location.z - 1,
			},
			dir: forward,
		}
	default:
		panic("Unrechable")
	}
}
func solvePart1(cubes []point) int {
	faces := buildFaces(cubes)

	var result []face

	for f := range faces {
		opposite := getOppositeFace(f)
		if !faces.Contains(opposite) {
			result = append(result, f)
		}
	}
	return len(result)
}
func findMaxDimension(cubes []point) int {
	m := 0
	for _, c := range cubes {
		m = common.Max(m, c.x, c.y, c.z)
	}
	return m
}

func newDroplet(size int) [][][]bool {
	droplet := make([][][]bool, size)

	for i := 0; i < size; i++ {
		droplet[i] = make([][]bool, size)
		for j := 0; j < size; j++ {
			droplet[i][j] = make([]bool, size)
			for k := 0; k < size; k++ {
				droplet[i][j][k] = false
			}
		}
	}
	return droplet
}
func solvePart2(cubes []point) int {

	maxDimension := findMaxDimension(cubes) + 2

	droplet := newDroplet(maxDimension)

	for _, c := range cubes {
		droplet[c.z][c.y][c.x] = true
	}

	if droplet[0][0][0] == true {
		panic("Invalid assumption")
	}

	encounteredFaces := common.Set[face]{}
	visitedPoints := make(map[point]bool)
	queue := []point{{0, 0, 0}}

	for len(queue) > 0 {
		p := queue[0]
		queue = queue[1:]

        if visitedPoints[p] {
            continue
        }

		visitedPoints[p] = true

		// check up direction
		if p.y+1 < maxDimension {
			if droplet[p.z][p.y+1][p.x] {
				encounteredFaces.Add(face{
					location: point{p.x, p.y + 1, p.z},
					dir:      down,
				})
			} else {
				queue = append(queue, point{p.x, p.y + 1, p.z})
			}
		}

		// check down direction
		if p.y-1 >= 0 {
			if droplet[p.z][p.y-1][p.x] {
				encounteredFaces.Add(face{
					location: point{p.x, p.y - 1, p.z},
					dir:      up,
				})
			} else {
				queue = append(queue, point{p.x, p.y - 1, p.z})
			}
		}

		// check right direction
		if p.x+1 < maxDimension {
			if droplet[p.z][p.y][p.x+1] {
				encounteredFaces.Add(face{
					location: point{p.x + 1, p.y, p.z},
					dir:      left,
				})
			} else {
				queue = append(queue, point{p.x + 1, p.y, p.z})
			}
		}

		// check left direction
		if p.x-1 >= 0 {
			if droplet[p.z][p.y][p.x-1] {
				encounteredFaces.Add(face{
					location: point{p.x - 1, p.y, p.z},
					dir:      right,
				})
			} else {
				queue = append(queue, point{p.x - 1, p.y, p.z})
			}
		}

		// check right direction
		if p.x+1 < maxDimension {
			if droplet[p.z][p.y][p.x+1] {
				encounteredFaces.Add(face{
					location: point{p.x + 1, p.y, p.z},
					dir:      left,
				})
			} else {
				queue = append(queue, point{p.x + 1, p.y, p.z})
			}
		}

		// check left direction
		if p.x-1 >= 0 {
			if droplet[p.z][p.y][p.x-1] {
				encounteredFaces.Add(face{
					location: point{p.x - 1, p.y, p.z},
					dir:      right,
				})
			} else {
				queue = append(queue, point{p.x - 1, p.y, p.z})
			}
		}

		// check forward direction
		if p.z+1 < maxDimension {
			if droplet[p.z+1][p.y][p.x] {
				encounteredFaces.Add(face{
					location: point{p.x, p.y, p.z + 1},
					dir:      backward,
				})
			} else {
				queue = append(queue, point{p.x, p.y, p.z + 1})
			}
		}

		// check backward direction
		if p.z-1 >= 0 {
			if droplet[p.z-1][p.y][p.x] {
				encounteredFaces.Add(face{
					location: point{p.x, p.y, p.z - 1},
					dir:      forward,
				})
			} else {
				queue = append(queue, point{p.x, p.y, p.z - 1})
			}
		}
	}

    return encounteredFaces.Len()
}

func main() {
	input := common.ReadInput("input.txt")
	parsedInput := parseInput(input)

	result1 := solvePart1(parsedInput)
	result2 := solvePart2(parsedInput)

	fmt.Printf("Part 1 result: %d\r\n", result1)
	fmt.Printf("Part 2 result: %d\r\n", result2)

}
