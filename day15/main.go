package main

import (
	"fmt"
    "sort"
    "strconv"
	"strings"
)
import "fardragon/aoc2022/v2/common"

type point struct {
	x int64
	y int64
}

type interval struct {
    start int64
    end int64
}
type sensor struct {
    position point
    sensorRange int64
    closestBeacon point
}

func parsePoint(input string) point {

    split := strings.Split(input, ", ")

    x, xErr := strconv.ParseInt(split[0][2:], 10, 64)
    if xErr != nil {
        panic(xErr)
    }

    y, yErr := strconv.ParseInt(split[1][2:], 10, 64)
    if yErr != nil {
        panic(yErr)
    }

    return point{x, y}
}

func sensorRange(sensor point, closestBeacon point) int64 {
    return common.Abs64(sensor.x - closestBeacon.x) + common.Abs64(sensor.y - closestBeacon.y)
}
func parseInput(input []string) []sensor {
	var result []sensor

	for _, line := range input {
        initialSplit := strings.Split(line, ": closest beacon is at ")

        pos := parsePoint(initialSplit[0][10:])
        beaconPos := parsePoint(initialSplit[1])

        result = append(result, sensor{
            position: pos,
            sensorRange: sensorRange(pos, beaconPos),
            closestBeacon: beaconPos,
        })
	}
	return result
}

func sensorRowInfluence(input sensor, row int64) []point {

    rowDiff := common.Abs64(input.position.y - row)

    var result []point

    for x := input.position.x - input.sensorRange + rowDiff; x <= input.position.x + input.sensorRange - rowDiff; x++{
        result = append(result, point{x, row})
    }


    return result
}

func solvePart1(input []sensor, targetRow int64) int {

    influencedPoints := common.Set[point]{}

    for _, sensor := range input {
        inf := sensorRowInfluence(sensor, targetRow)

        for _, p := range inf {
            influencedPoints.Add(p)
        }
    }

    // Remove beacons from set
    for _, sensor := range input {
        if influencedPoints.Contains(sensor.closestBeacon) {
            influencedPoints.Remove(sensor.closestBeacon)
        }
    }

	return influencedPoints.Len()
}


func calculateSensorRowInterval(input sensor, row int64) *interval {

    rowDiff := common.Abs64(input.position.y - row)

    start := input.position.x - input.sensorRange + rowDiff
    end := input.position.x + input.sensorRange - rowDiff

    if start <= end {
        return &interval{start, end}
    } else {
        return nil
    }
}

func max64(a int64, b int64) int64 {
    if a > b {
        return a
    } else {
        return b
    }
}
func mergeIntervals(input []interval) []interval {
    sort.Slice(input, func(i int, j int) bool {
        return input[i].start < input[j].start
    })

    index := 0
    for i := 1; i < len(input); i++ {
        if input[index].end >= (input[i].start - 1) {
            input[index].end = max64(input[index].end, input[i].end)
        } else {
            index++
            input[index] = input[i]
        }
    }

    return input[0:index+1]
}


func solvePart2(input []sensor, bound int64) int64 {
    for y := int64(0); y <= bound; y++ {
        var rowIntervals []interval
        for _, s := range input {
            sensorRowInterval := calculateSensorRowInterval(s, y)
            if sensorRowInterval != nil {
                rowIntervals = append(rowIntervals, *sensorRowInterval)
            }
        }
        mergedIntervals := mergeIntervals(rowIntervals)
        if len(mergedIntervals) == 1 {
            continue
        } else if len(mergedIntervals) == 2 {
            x := (mergedIntervals[1].start + mergedIntervals[0].end) / 2
            return x * 4000000 + y
        } else {
            panic("Solving error")
        }

    }

    panic("Unreachable")
}

func main() {
	input := common.ReadInput("input.txt")
	parsedInput := parseInput(input)

    result1 := solvePart1(parsedInput, 2000000)
    result2 := solvePart2(parsedInput, 4000000)

	fmt.Printf("Part 1 result: %d\r\n", result1)
	fmt.Printf("Part 2 result: %d\r\n", result2)

}
