package main

import (
	"fardragon/aoc2022/v2/common"
	"fmt"
	"sort"
	"strconv"
)

type data interface {
	isData()
}

type integerData struct {
	data int
}

func (integerData) isData() {}

type listData struct {
	data []data
}

func (listData) isData() {}

type packet struct {
	data listData
}

type packetPair struct {
	first  packet
	second packet
}

type orderStatus int8

const (
	inOrder orderStatus = iota
	notInOrder
	equal
)

func parseListData(packetStr string, index *int) listData {

	*index++ // skip [

	result := listData{}

	if packetStr[*index] == ']' {
		*index++
		return result
	}

loop:
	for {
		result.data = append(result.data, parseData(packetStr, index))
		switch packetStr[*index] {
		case ',':
			*index++
		case ']':
			*index++
			break loop
		default:
			panic("Parsing error")
		}
	}

	return result
}

func parseIntegerData(packetStr string, index *int) integerData {
	integerToParse := ""

	for packetStr[*index] >= '0' && packetStr[*index] <= '9' {
		integerToParse += string(packetStr[*index])
		*index++
	}

	val, err := strconv.Atoi(integerToParse)
	if err != nil {
		panic(err)
	}
	return integerData{val}
}

func parseData(packetStr string, index *int) data {
	if packetStr[*index] == '[' {
		return parseListData(packetStr, index)
	} else {
		return parseIntegerData(packetStr, index)
	}
}

func parseInput(input []string) []packetPair {

	index := 0

	var result []packetPair

	for index < len(input) {
		lineIndex := 0
		packet1 := parseListData(input[index], &lineIndex)
		index++

		lineIndex = 0
		packet2 := parseListData(input[index], &lineIndex)
		index++

		index++

		result = append(result, packetPair{packet{packet1}, packet{packet2}})
	}

	return result
}

func listInOrder(left listData, right listData) orderStatus {

	index := 0
	for (index < len(left.data)) && (index < len(right.data)) {
		elementOrder := dataInOrder(left.data[index], right.data[index])

		if elementOrder != equal {
			return elementOrder
		}
		index++
	}

	if len(left.data) == len(right.data) {
		return equal
	} else if len(left.data) < len(right.data) {
		return inOrder
	} else {
		return notInOrder
	}
}

func dataInOrder(left data, right data) orderStatus {

	_, leftList := left.(listData)
	_, rightList := right.(listData)

	if leftList && rightList {
		return listInOrder(left.(listData), right.(listData))
	} else if !leftList && !rightList {

		lInt := left.(integerData).data
		rInt := right.(integerData).data
		if lInt == rInt {
			return equal
		} else if lInt < rInt {
			return inOrder
		} else {
			return notInOrder
		}
	} else {
		if leftList {
			return listInOrder(left.(listData), listData{[]data{right.(integerData)}})
		} else {
			return listInOrder(listData{[]data{left.(integerData)}}, right.(listData))
		}
	}

}
func solvePart1(input []packetPair) int {

	result := 0
	for ix, pair := range input {
		if dataInOrder(pair.first.data, pair.second.data) == inOrder {
			result += ix + 1
		}
	}
	return result
}

func solvePart2(input []packetPair) int {

	divider1 := packet{listData{[]data{listData{[]data{integerData{2}}}}}}
	divider2 := packet{listData{[]data{listData{[]data{integerData{6}}}}}}

	packetList := []packet{
		divider1,
		divider2,
	}
	for _, pair := range input {
		packetList = append(packetList, []packet{pair.first, pair.second}...)
	}

	sort.Slice(packetList, func(i int, j int) bool {
		return dataInOrder(packetList[i].data, packetList[j].data) == inOrder
	})

	result := 1

	for ix, p := range packetList {
		if (dataInOrder(p.data, divider1.data) == equal) || (dataInOrder(p.data, divider2.data) == equal) {
			result *= ix + 1
		}
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
