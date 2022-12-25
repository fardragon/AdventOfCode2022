package main

import (
	"fardragon/aoc2022/v2/common"
	"fmt"
    "strconv"
    "strings"
)


type intPair struct {
    first int
    second int
}
type bluerpint struct {
    oreRobotCost int
    clayRobotCost int
    obsidianRobotCost intPair
    geodeRobotCost intPair
}

type gameState struct {
    ore int
    clay int
    obisdian int
    geodes int
    oreRobots int
    clayRobots int
    obsidianRobots int
    geodeRobots int
    time int
}

func (state *gameState) work() {
    state.ore += state.oreRobots
    state.clay += state.clayRobots
    state.obisdian += state.obsidianRobots
    state.geodes += state.geodeRobots
    state.time += 1
}

func parseBlueprint(input string) bluerpint {

    split := strings.Split(input, " ")

    oreCost, _ := strconv.Atoi(split[6])
    clayCost, _ := strconv.Atoi(split[12])
    obsidianCost1, _ := strconv.Atoi(split[18])
    obsidianCost2, _ := strconv.Atoi(split[21])
    geodeCost1, _ := strconv.Atoi(split[27])
    geodeCost2, _ := strconv.Atoi(split[30])

    return bluerpint{
        oreRobotCost:  oreCost,
        clayRobotCost: clayCost,
        obsidianRobotCost: intPair{
            first:  obsidianCost1,
            second: obsidianCost2,
        },
        geodeRobotCost: intPair{
            first:  geodeCost1,
            second: geodeCost2,
        },
    }
}
func parseInput(input []string) []bluerpint {
    var result []bluerpint
    for _, line := range input {
        result = append(result, parseBlueprint(line))
    }
	return result
}

func solveBlueprint(b bluerpint, maxTime int) int {

    startingState := gameState{
        oreRobots:      1,
    }

    statesToSolve := []gameState{startingState}
    memory := common.Set[gameState]{}

    bestGeodes := 0
    maxOreCost := common.Max(b.oreRobotCost, b.clayRobotCost, b.obsidianRobotCost.first, b.geodeRobotCost.first)

    for len(statesToSolve) > 0 {
        state := statesToSolve[0]
        statesToSolve = statesToSolve[1:]

        bestGeodes = common.Max(bestGeodes, state.geodes)

        cacheState := state
        cacheState.time = 0

        if (state.geodes < (bestGeodes - 2)) || memory.Contains(cacheState) || (state.time == maxTime) {
            continue
        }

        memory.Add(cacheState)

        if (state.ore >= b.geodeRobotCost.first) && (state.obisdian >= b.geodeRobotCost.second) {
            nextState := state
            nextState.ore -= b.geodeRobotCost.first
            nextState.obisdian -= b.geodeRobotCost.second
            nextState.work()
            nextState.geodeRobots += 1
            statesToSolve = append(statesToSolve, nextState)
        } else {
            if (state.ore >= b.oreRobotCost) && (state.oreRobots < maxOreCost) {
                nextState := state
                nextState.ore -= b.oreRobotCost
                nextState.work()
                nextState.oreRobots += 1
                statesToSolve = append(statesToSolve, nextState)
            }

            if (state.ore >= b.clayRobotCost) && (state.clayRobots < b.obsidianRobotCost.second) {
                nextState := state
                nextState.ore -= b.clayRobotCost
                nextState.work()
                nextState.clayRobots += 1
                statesToSolve = append(statesToSolve, nextState)
            }

            if (state.ore >= b.obsidianRobotCost.first) && (state.clay >= b.obsidianRobotCost.second) {
                nextState := state
                nextState.ore -= b.obsidianRobotCost.first
                nextState.clay -= b.obsidianRobotCost.second
                nextState.work()
                nextState.obsidianRobots += 1
                statesToSolve = append(statesToSolve, nextState)
            }

            {
                nextState := state
                nextState.work()
                statesToSolve = append(statesToSolve, nextState)
            }
}
}
    return bestGeodes
}
func solvePart1(blueprints []bluerpint) int {
    result := 0

    for ix, b := range blueprints {
        result += solveBlueprint(b, 24) * (ix + 1)
    }
    return result
}
func solvePart2(blueprints []bluerpint) int {

    result := 1

    for i:=0; (i < 3) && (i < len(blueprints)); i++ {
        geodes := solveBlueprint(blueprints[i], 32)
        result *= geodes
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
