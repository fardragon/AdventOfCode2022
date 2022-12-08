package main

import (
	"fardragon/aoc2022/v2/common"
	"fmt"
	"sort"
	"strconv"
	"strings"
)

type file struct {
	name string
	size uint64
}

type directory struct {
	name           string
	subdirectories map[string]directory
	files          []file
	size           uint64
}

type directoryEntry struct {
	name string
	size uint64
}

func parseDirectory(commands *[]string, name string) directory {
	result := directory{name, make(map[string]directory), []file{}, 0}

	if (*commands)[0] != "$ ls" {
		panic("Invalid input")
	}
	*commands = (*commands)[1:]

	for len(*commands) > 0 && (*commands)[0][0] != '$' {
		command := (*commands)[0]
		if strings.HasPrefix(command, "dir") {
			dirname := strings.Split(command, " ")[1]
			result.subdirectories[dirname] = directory{}
		} else {
			fileSplit := strings.Split(command, " ")
			fileSize, err := strconv.ParseUint(fileSplit[0], 10, 0)
			if err != nil {
				panic(err)
			}
			result.files = append(result.files, file{
				name: fileSplit[1],
				size: fileSize,
			})
			result.size += fileSize
		}
		*commands = (*commands)[1:]
	}

	for len(*commands) > 0 {
		command := (*commands)[0]
		*commands = (*commands)[1:]
		if strings.HasPrefix(command, "$ cd") {
			cdTarget := strings.Split(command, " ")[2]
			if cdTarget == ".." {
				break
			} else {
				if _, ok := result.subdirectories[cdTarget]; ok {
					result.subdirectories[cdTarget] = parseDirectory(commands, cdTarget)
					result.size += result.subdirectories[cdTarget].size
				} else {
					panic(fmt.Sprintf("Unknown directory: %s", cdTarget))
				}
			}
		} else {
			panic(fmt.Sprintf("Unexpected command: %s", command))
		}
	}

	return result
}
func parseInput(input []string) directory {
	if input[0] != "$ cd /" {
		panic("Invalid input")
	}
	input = input[1:]
	return parseDirectory(&input, "/")
}

func walkDirectories(root directory) []directoryEntry {

	result := []directoryEntry{
		{
			root.name,
			root.size,
		},
	}

	for _, dir := range root.subdirectories {
		result = append(result, walkDirectories(dir)...)
	}

	return result
}
func solvePart1(root directory) uint64 {
	directories := walkDirectories(root)

	result := uint64(0)
	for _, dir := range directories {
		if dir.size <= 100000 {
			result += dir.size
		}
	}

	return result
}

func solvePart2(root directory) uint64 {

	directories := walkDirectories(root)

	sort.Slice(directories, func(i, j int) bool {
		return directories[i].size < directories[j].size
	})

	totalSpace := uint64(70000000)
	neededSpace := uint64(30000000)
	unusedSpace := totalSpace - directories[len(directories)-1].size

	for _, dir := range directories {
		if (unusedSpace + dir.size) >= neededSpace {
			return dir.size
		}
	}
	panic("Did not find a solution")
}

func main() {
	input := common.ReadInput("input.txt")
	parsedInput := parseInput(input)

	result1 := solvePart1(parsedInput)
	result2 := solvePart2(parsedInput)

	fmt.Printf("Part 1 result: %d\r\n", result1)
	fmt.Printf("Part 2 result: %d\r\n", result2)
}
