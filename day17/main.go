package main

import (
	"fmt"
)
import "fardragon/aoc2022/v2/common"

type point struct {
	x int64
	y int64
}

type shape struct {
	points []point
}

type direction int8

const (
	left direction = iota
	right
)

type fieldType int8

const (
	wall fieldType = iota
	floor
	air
	rock
	stationaryRock
)

type cave struct {
	fields [][9]fieldType
	height int64
}

type block struct {
	points []point
}

func newCave() cave {
	return cave{
		fields: [][9]fieldType{
			{floor, floor, floor, floor, floor, floor, floor, floor, floor},
			//            {wall, air, air, air, air, air, air, air, wall},
			//            {wall, air, air, air, air, air, air, air, wall},
			//            {wall, air, air, air, air, air, air, air, wall},
		},
		height: 0,
	}
}

func parseInput(input []string) []direction {
	var result []direction

	for _, char := range input[0] {
		if char == '<' {
			result = append(result, left)
		} else if char == '>' {
			result = append(result, right)
		} else {
			panic("Invalid input")
		}
	}
	return result
}

func getNextBlock(index int64, height int64) block {

	blockType := index % 5

	if blockType == 0 {
		return block{
			points: []point{{3, height}, {4, height}, {5, height}, {6, height}},
		}
	} else if blockType == 1 {
		return block{
            points: []point{{4, height}, {3, height + 1}, {4, height + 1}, {5, height + 1}, {4, height + 2}},
		}
	} else if blockType == 2 {
		return block{
			points: []point{{3, height}, {4, height}, {5, height}, {5, height + 1}, {5, height + 2}},
		}
	} else if blockType == 3 {
		return block{
			points: []point{{3, height}, {3, height + 1}, {3, height + 2}, {3, height + 3}},
		}
	} else {
		return block{
			points: []point{{3, height}, {3, height + 1}, {4, height}, {4, height + 1}},
		}
	}

}

func canMoveHorizontally(b block, dir direction, c cave) bool {
	if dir == left {
		for _, p := range b.points {
			if (c.fields[p.y][p.x-1] == wall) || (c.fields[p.y][p.x-1] == stationaryRock) {
				return false
			}
		}
	} else if dir == right {
		for _, p := range b.points {
			if (c.fields[p.y][p.x+1] == wall) || (c.fields[p.y][p.x+1] == stationaryRock) {
				return false
			}
		}
	} else {
		panic("Unreachable")
	}
	return true
}

func moveHorizontally(b *block, dir direction, c *cave) {

	for _, p := range b.points {
		(*c).fields[p.y][p.x] = air
	}

	if dir == left {
		for ix := range b.points {
			b.points[ix].x -= 1
		}
	} else if dir == right {
		for ix := range b.points {
			b.points[ix].x += 1
		}
	} else {
		panic("Unreachable")
	}

	for _, p := range b.points {
		(*c).fields[p.y][p.x] = rock
	}
}

func canMoveDown(b block, c cave) bool {

	for _, p := range b.points {
		if (c.fields[p.y-1][p.x] == wall) || (c.fields[p.y-1][p.x] == stationaryRock) || (c.fields[p.y-1][p.x] == floor) {
			return false
		}
	}
	return true
}

func moveDown(b *block, c *cave) {

	for _, p := range b.points {
		(*c).fields[p.y][p.x] = air
	}

	for ix := range b.points {
		b.points[ix].y -= 1
	}

	for _, p := range b.points {
		(*c).fields[p.y][p.x] = rock
	}
}

func makeStationary(b block, c *cave) {

    for _, p := range b.points {
        (*c).fields[p.y][p.x] = stationaryRock
    }
}
func getTop(c cave) int64 {

    for i := int64(len(c.fields)) - 1; i >= 0; i-- {
        for _, row := range c.fields[i] {
            if row == stationaryRock {
                return i
            }
        }
    }
    panic("Unreachable")
}

func printCave(c cave) {
    fmt.Println("________________________________")
    for i := len(c.fields) - 1; i >= 0; i-- {
        row := c.fields[i]
        for _, field := range row {
            if field == air {
                fmt.Print(".")
            } else if field == wall {
                fmt.Print("|")
            } else if field == rock {
                fmt.Print("@")
            } else if field == stationaryRock {
                fmt.Print("#")
            } else if field == floor {
                fmt.Print("-")
            } else {
                panic("Unrechable")
            }
        }
        fmt.Println("")
    }
}

func solvePart1(input []direction) int64 {

	c := newCave()
	patternIndex := 0

    for i := int64(0); i < 2022; i++ {
//        fmt.Printf("Processing block: %d\r\n", i)

		// place new block
		b := getNextBlock(i, c.height + 3 + 1)
		for _, p := range b.points {
            blockExpand := p.y - int64(len(c.fields)) + 1
            for j := int64(0); j < blockExpand; j++ {
				c.fields = append(c.fields, [9]fieldType{wall, air, air, air, air, air, air, air, wall})
			}
			c.fields[p.y][p.x] = rock
		}

//        printCave(c)

		// start moving block
		for {
			if canMoveHorizontally(b, input[patternIndex], c) {
				moveHorizontally(&b, input[patternIndex], &c)
			}

			patternIndex = (patternIndex + 1) % len(input)

			if canMoveDown(b, c) {
				moveDown(&b, &c)
			} else {
				break
			}
		}

        makeStationary(b, &c)

        // get new height
        c.height = getTop(c)
	}

	return c.height
}

type caveState struct {
    nextBlock int64
    nextJet int64
    heightProfile [7]int64
}

type roundState struct {
    index int64
    height int64
}

func getHeightProfile(c cave) [7] int64{

    profile := [7]int64{}

    for column := int64(0); column <= 6; column++{
        for row := int64(len(c.fields)) - 1; row >= 0; row-- {
            if c.fields[row][column + 1] != air {
                profile[column] = int64(len(c.fields)) - row
                break
            }
        }
    }
    return profile
}
func solvePart2(input []direction) int64 {
    c := newCave()
    patternIndex := int64(0)

    memory := map[caveState]roundState{}
    skippedHeight := int64(0)
    skipped := false

    for i := int64(0); i < 1_000_000_000_000; i++ {

        if !skipped {
            state := caveState{
                nextBlock: int64(i) % 5,
                nextJet: patternIndex,
                heightProfile: getHeightProfile(c),
                }
                if cachedState, contains := memory[state]; contains{
                    cycleSize := i - cachedState.index
                    cycleHeight := c.height - cachedState.height

                    remainingBlocks := 1_000_000_000_000 - i - 1
                    skipCycles := remainingBlocks / cycleSize
                    skippedHeight = skipCycles * cycleHeight

                    i += skipCycles * cycleSize
                    skipped = true
                } else {
                    memory[state] = roundState{i, c.height}
                }
        }


        // place new block
        b := getNextBlock(i, c.height + 3 + 1)
        for _, p := range b.points {
            blockExpand := p.y - int64(len(c.fields)) + 1
            for j := int64(0); j < blockExpand; j++ {
                c.fields = append(c.fields, [9]fieldType{wall, air, air, air, air, air, air, air, wall})
            }
            c.fields[p.y][p.x] = rock
        }

        //        printCave(c)

        // start moving block
        for {
            if canMoveHorizontally(b, input[patternIndex], c) {
                moveHorizontally(&b, input[patternIndex], &c)
            }

            patternIndex = (patternIndex + 1) % int64(len(input))

            if canMoveDown(b, c) {
                moveDown(&b, &c)
            } else {
                break
            }
        }

        makeStationary(b, &c)

        // get new height
        c.height = getTop(c)
    }

    return c.height + skippedHeight
}

func main() {
	input := common.ReadInput("input.txt")
	parsedInput := parseInput(input)

	result1 := solvePart1(parsedInput)
	result2 := solvePart2(parsedInput)

	fmt.Printf("Part 1 result: %d\r\n", result1)
	fmt.Printf("Part 2 result: %d\r\n", result2)

}
