package main

import (
	"fardragon/aoc2022/v2/common"
	"fmt"
    "math"
	"sort"
)

type point struct {
	x int
	y int
}
type heightmap struct {
	heights [][]int
	height  int
	width   int
}

type puzzle struct {
	hMap  heightmap
	start point
	end   point
}

func parseInput(input []string) puzzle {

	width := len(input[0])
	height := len(input)

	result := heightmap{
		width:  width,
		height: height,
	}

	start := point{}
	end := point{}

	for y := 0; y < height; y++ {
		row := input[y]
		if len(row) != width {
			panic("Invalid input")
		}
		result.heights = append(result.heights, []int{})

		for x := 0; x < width; x++ {
			if row[x] == 'S' {
				start = point{x, y}
				result.heights[y] = append(result.heights[y], 0)
			} else if row[x] == 'E' {
				end = point{x, y}
				result.heights[y] = append(result.heights[y], int('z'-'a'))
			} else {
				val := row[x] - 'a'
				result.heights[y] = append(result.heights[y], int(val))
			}
		}
	}

	return puzzle{
		result,
		start,
		end,
	}
}

func findMinDist(dist map[point]int, q common.Set[point]) *point {

    minDist := math.MaxInt

    result := new(point)
    found := false

    for key := range q {
        if dist[key] < minDist {
            minDist = dist[key]
            *result = point{key.x, key.y}
            found = true
        }
    }

    if !found {
        return nil
    }
    return result
}

func getNeighbours(s point, hMap heightmap) []point{
    possible := []point{
        {s.x + 1, s.y},
        {s.x - 1, s.y},
        {s.x, s.y + 1},
        {s.x, s.y - 1},
    }

    var neighbours []point
    for _, p := range possible {
        if (p.x >= 0) && (p.x < hMap.width) && (p.y >= 0) && (p.y < hMap.height) {
            if (hMap.heights[s.y][s.x] - hMap.heights[p.y][p.x]) >= -1 {
                neighbours = append(neighbours, p)
            }
        }
    }
    return neighbours
}

func findShortestPath(input puzzle) []point {

	dist := map[point]int{}
	prev := map[point]point{}

    q := common.Set[point]{}

	for y := 0; y < input.hMap.height; y++ {
		for x := 0; x < input.hMap.width; x++ {
            v := point{x, y}
            dist[v] = math.MaxInt
            q.Add(v)
		}
	}
    dist[input.start] = 0

    for !q.Empty() {
        u := findMinDist(dist, q)

        if u == nil {
            // Path not found
            return make([]point, input.hMap.height * input.hMap.width)
        }

        q.Remove(*u)

        if *u == input.end {
            var s []point
            for *u != input.start {
                s = append([]point{*u}, s...)
                *u = prev[*u]
            }
            return s
        }



        for _, v := range getNeighbours(*u, input.hMap) {
            if q.Contains(v) {
                alt := dist[*u] + 1
                if alt < dist[v] {
                    dist[v] = alt
                    prev[v] = *u
                }
            }
        }
    }
    panic("Unreachable")
}
func solvePart1(input puzzle) int {

    shortestPath := findShortestPath(input)
	return len(shortestPath)
}

func solvePart2(input puzzle) int {

    var results []int

    for y:=0; y < input.hMap.height; y++{
        for x:=0; x < input.hMap.width; x++ {
            if input.hMap.heights[y][x] == 0 {
                input.start = point{x, y}
                results = append(results, len(findShortestPath(input)))
            }
        }
    }

    sort.Ints(results)
	return results[0]
}

func main() {
	input := common.ReadInput("input.txt")
	parsedInput := parseInput(input)

	result1 := solvePart1(parsedInput)
	result2 := solvePart2(parsedInput)

	fmt.Printf("Part 1 result: %d\r\n", result1)
	fmt.Printf("Part 2 result: %d\r\n", result2)
}
