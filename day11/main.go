package main

import (
	"fardragon/aoc2022/v2/common"
	"fmt"
	"sort"
	"strconv"
	"strings"
)

type operation interface {
	execute(old uint) uint
}

type addIntInstruction struct {
	operand uint
}

func (ins addIntInstruction) execute(old uint) uint {
	return old + ins.operand
}

type multiplyIntInstruction struct {
	operand uint
}

func (ins multiplyIntInstruction) execute(old uint) uint {
	return old * ins.operand
}

type squareInstruction struct{}

func (ins squareInstruction) execute(old uint) uint {
	return old * old
}

type monkey struct {
	items     []uint
	op        operation
	divisor   uint
	receivers [2]uint
}

func parseMonkey(input []string, index *int) monkey {
	result := monkey{}

	if !strings.HasPrefix(input[*index], "Monkey ") {
		panic("Parsing error")
	}
	*index++

	if !strings.HasPrefix(input[*index], "  Starting items:") {
		panic("Parsing error")
	}
	itemsStr := strings.TrimPrefix(input[*index], "  Starting items:")
	itemsStr = strings.ReplaceAll(itemsStr, " ", "")
	itemsSplit := strings.Split(itemsStr, ",")
	for _, item := range itemsSplit {
		parsedItem, err := strconv.ParseUint(item, 10, 0)
		if err != nil {
			panic(err)
		}
		result.items = append(result.items, uint(parsedItem))
	}
	*index++

	if !strings.HasPrefix(input[*index], "  Operation: new = ") {
		panic("Parsing error")
	}
	operationStr := strings.TrimPrefix(input[*index], "  Operation: new = ")
	operationSplit := strings.Split(operationStr, " ")
	if len(operationSplit) != 3 {
		panic("Parsing error")
	}
	if operationSplit[1] == "+" {
		parsedOperand, err := strconv.ParseUint(operationSplit[2], 10, 0)
		if err != nil {
			panic(err)
		}
		result.op = addIntInstruction{uint(parsedOperand)}
	} else if operationSplit[1] == "*" {
		if operationSplit[2] == "old" {
			result.op = squareInstruction{}
		} else {
			parsedOperand, err := strconv.ParseUint(operationSplit[2], 10, 0)
			if err != nil {
				panic(err)
			}
			result.op = multiplyIntInstruction{uint(parsedOperand)}
		}
	} else {
		panic("Parsing error")
	}
	*index++

	if !strings.HasPrefix(input[*index], "  Test: divisible by ") {
		panic("Parsing error")
	}
	divisorStr := strings.TrimPrefix(input[*index], "  Test: divisible by ")
	parsedDivisor, err := strconv.ParseUint(divisorStr, 10, 0)
	if err != nil {
		panic(err)
	}
	result.divisor = uint(parsedDivisor)
	*index++

	if !strings.HasPrefix(input[*index], "    If true: throw to monkey ") {
		panic("Parsing error")
	}
	receiver1Str := strings.TrimPrefix(input[*index], "    If true: throw to monkey ")
	prasedReceiver1, err := strconv.ParseUint(receiver1Str, 10, 0)
	if err != nil {
		panic(err)
	}
	result.receivers[0] = uint(prasedReceiver1)
	*index++

	if !strings.HasPrefix(input[*index], "    If false: throw to monkey ") {
		panic("Parsing error")
	}
	receiver2Str := strings.TrimPrefix(input[*index], "    If false: throw to monkey ")
	prasedReceiver2, err := strconv.ParseUint(receiver2Str, 10, 0)
	if err != nil {
		panic(err)
	}
	result.receivers[1] = uint(prasedReceiver2)
	*index++

	//Skip empty line that separates monkeys
	*index++

	return result
}

func parseInput(input []string) []monkey {

	var result []monkey

	i := 0
	for i < len(input) {
		result = append(result, parseMonkey(input, &i))
	}
	return result
}

func proceesMonkey(monkeys []monkey, index int, worryDivisor uint, worryModulo uint) int {

	monkey := &monkeys[index]
	for _, elem := range monkey.items {
		newElem := monkey.op.execute(elem)
        if worryDivisor > 1 {
            newElem /= worryDivisor
        } else {
            newElem %= worryModulo
        }
		if newElem%monkey.divisor == 0 {
            monkeys[monkey.receivers[0]].items = append(monkeys[monkey.receivers[0]].items, newElem)
		} else {
            monkeys[monkey.receivers[1]].items = append(monkeys[monkey.receivers[1]].items, newElem)
		}
	}

	result := len(monkey.items)
	monkey.items = []uint{}

	return result
}
func solvePart1(input []monkey) int {
    activity := make([]int, len(input))
    monkeys := make([]monkey, len(input))
    copy(monkeys, input)

	for i := 0; i < 20; i++ {
        for index := range monkeys {
            inspectedItems := proceesMonkey(monkeys, index, 3, 0)
			activity[index] += inspectedItems
		}
	}

	sort.Ints(activity)
	return activity[len(activity)-1] * activity[len(activity)-2]
}

func solvePart2(input []monkey) int {
    activity := make([]int, len(input))
    monkeys := make([]monkey, len(input))
    copy(monkeys, input)

    modulo := uint(1)
    for _, m := range monkeys {
        modulo *= m.divisor
    }

    for i := 0; i < 10000; i++ {
        for index := range monkeys {
            inspectedItems := proceesMonkey(monkeys, index, 1, modulo)
            activity[index] += inspectedItems
        }
    }

    sort.Ints(activity)
    return activity[len(activity)-1] * activity[len(activity)-2]
}

func main() {
	input := common.ReadInput("input.txt")
	parsedInput := parseInput(input)

	result1 := solvePart1(parsedInput)
	result2 := solvePart2(parsedInput)

	fmt.Printf("Part 1 result: %d\r\n", result1)
	fmt.Printf("Part 2 result: %d\r\n", result2)
}
