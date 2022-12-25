package main

import (
    "fardragon/aoc2022/v2/common"
	"fmt"
    "sort"
    "strconv"
	"strings"
)

type room struct {
    name string
    pressure int
    exits []string
}

type path struct {
    a string
    b string
}

func parseInput(input []string) []room {
    var result []room

	for _, line := range input {
        initialSplit := strings.Split(line, " ")
        name := initialSplit[1]
        pressureStr := initialSplit[4]
        pressureStr = strings.TrimPrefix(pressureStr, "rate=")
        pressureStr = strings.TrimSuffix(pressureStr, ";")
        pressure, err := strconv.Atoi(pressureStr)
        if err != nil {
            panic(err)
        }
        exitsSlice := initialSplit[9:]
        for ix, exit := range exitsSlice {
            exitsSlice[ix] = strings.TrimSuffix(exit, ",")
        }

        result = append(result, room{
            name:     name,
            pressure: pressure,
            exits:    exitsSlice,
        })
	}
	return result
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



func findShortestPaths(rooms []room) map[path]int {
    result := map[path]int{}

    for _, u := range rooms {
        result[path{u.name,u.name}] = 0
        for _, v := range u.exits {
            result[path{u.name, v}] = 1
        }
    }

    for _, k := range rooms {
        for _, i := range rooms {
            if _, present := result[path{i.name, k.name}]; !present {
                continue
            }
            for _, j := range rooms {
                if _, present := result[path{k.name, j.name}]; !present {
                    continue
                }

                if _, present := result[path{i.name, j.name}]; present {
                    result[path{i.name, j.name}]  = min(
                        result[path{i.name, j.name}],
                        result[path{i.name, k.name}] + result[path{k.name, j.name}],
                        )
                } else {
                    result[path{i.name, j.name}] = result[path{i.name, k.name}] + result[path{k.name, j.name}]
                }
            }
        }
    }

    return result
}

func findPath(timeLimit int, currentRoom string, time int, openedValves common.Set[string], pressurizedRooms []room, shortestPaths map[path]int) int {

    if openedValves.Len() == len(pressurizedRooms) {
        return 0
    }

    pressure := 0
    for _, destination := range pressurizedRooms {
        if openedValves.Contains(destination.name) {
            continue
        }

        timeToOpen := time + shortestPaths[path{currentRoom, destination.name}] + 1

        if timeToOpen <= timeLimit {
            releasePresure := destination.pressure * (timeLimit - timeToOpen)
            openedValves.Add(destination.name)

            subPressure := findPath(timeLimit, destination.name, timeToOpen, openedValves, pressurizedRooms, shortestPaths)

            openedValves.Remove(destination.name)

            pressure = max(pressure, releasePresure + subPressure)
        }
    }
    return pressure
}
func solvePart1(input []room) int {

    shortestPaths := findShortestPaths(input)

    var pressurizedRooms []room
    for _, r := range input {
        if r.pressure > 0 {
            pressurizedRooms = append(pressurizedRooms, r)
        }
    }

    return findPath(30, "AA", 0, common.Set[string]{}, pressurizedRooms, shortestPaths)
}

func pow(n uint, m uint) uint {
    if m == 0 {
        return 1
    }
    result := n
    for i := uint(2); i <= m; i++ {
        result *= n
    }
    return result
}

func partition(input []room) [][][]room {
    var partitions [][][]room

    for i := uint(0); i < pow(2, uint(len(input))); i++ {

        var subsetA []room
        var subsetB []room

        for j := 0; j < len(input); j++ {
            bitmask := uint(1) << j

            if (i & bitmask) != 0 {
                subsetA = append(subsetA, input[j])
            } else {
                subsetB = append(subsetB, input[j])
            }
        }
        partitions = append(partitions, [][]room{subsetA, subsetB})
    }
    return partitions
}

func solvePart2(input []room) int {

    shortestPaths := findShortestPaths(input)

    var pressurizedRooms []room

    for _, r := range input {
        if r.pressure > 0 {
            pressurizedRooms = append(pressurizedRooms, r)
        }
    }


    sort.Slice(pressurizedRooms, func(i int, j int) bool {
        return pressurizedRooms[i].name < pressurizedRooms[j].name
    })


    maxPressure := 0
    for _, p := range partition(pressurizedRooms) {
        pressure := findPath(26, "AA", 0, common.Set[string]{}, p[0], shortestPaths) +
            findPath(26, "AA", 0, common.Set[string]{}, p[1], shortestPaths)

        maxPressure = max(maxPressure, pressure)
    }

    return maxPressure
}

func main() {
	input := common.ReadInput("input.txt")
	parsedInput := parseInput(input)

    result1 := solvePart1(parsedInput)
    result2 := solvePart2(parsedInput)

	fmt.Printf("Part 1 result: %d\r\n", result1)
	fmt.Printf("Part 2 result: %d\r\n", result2)

}
