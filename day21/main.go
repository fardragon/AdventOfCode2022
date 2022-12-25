package main

import (
	"fardragon/aoc2022/v2/common"
	"fmt"
	"strconv"
	"strings"
)

type expressionMap map[string]expression

type expression interface {
	getName() string
	getValue(expressions expressionMap) int64
	solutionContainsHuman(expressions expressionMap) bool
	reverseSolve(expressions expressionMap, target int64) int64
}

type intExpression struct {
	name  string
	value int64
}

func (e intExpression) getName() string {
	return e.name
}

func (e intExpression) getValue(_ expressionMap) int64 {
	return e.value
}

func (e intExpression) solutionContainsHuman(_ expressionMap) bool {
	return false
}

func (e intExpression) reverseSolve(_ expressionMap, _ int64) int64 {
	panic("Unreachable")
}

type operation int8

const (
	add operation = iota
	subtract
	multiply
	divide
)

type operatorExpression struct {
	name         string
	operator     operation
	leftOperand  string
	rightOperand string
}

func (e operatorExpression) getName() string {
	return e.name
}

func (e operatorExpression) getValue(expressions expressionMap) int64 {

	lVal := expressions[e.leftOperand].getValue(expressions)
	rVal := expressions[e.rightOperand].getValue(expressions)

	switch e.operator {
	case add:
		return lVal + rVal
	case subtract:
		return lVal - rVal
	case multiply:
		return lVal * rVal
	case divide:
		return lVal / rVal
	default:
		panic("Unrechable")
	}
}

func (e operatorExpression) solutionContainsHuman(expressions expressionMap) bool {

	if e.leftOperand == "humn" || e.rightOperand == "humn" {
		return true
	} else {
		return expressions[e.leftOperand].solutionContainsHuman(expressions) || expressions[e.rightOperand].solutionContainsHuman(expressions)
	}
}

func (e operatorExpression) reverseSolve(expressions expressionMap, target int64) int64 {
	if expressions[e.leftOperand].getName() == "humn" ||
		expressions[e.leftOperand].solutionContainsHuman(expressions) {
		right := expressions[e.rightOperand].getValue(expressions)
		var newTarget int64

		switch e.operator {
		case add:
			newTarget = target - right
		case subtract:
			newTarget = target + right
		case multiply:
			newTarget = target / right
		case divide:
			newTarget = target * right
		default:
			panic("Unrechable")
		}
        if expressions[e.leftOperand].getName() == "humn" {
            return newTarget
        } else {
            return expressions[e.leftOperand].reverseSolve(expressions, newTarget)
        }

	} else if expressions[e.rightOperand].getName() == "humn" ||
        expressions[e.rightOperand].solutionContainsHuman(expressions) {
		left := expressions[e.leftOperand].getValue(expressions)

		var newTarget int64

		switch e.operator {
		case add:
			newTarget = target - left
		case subtract:
			newTarget = left - target
		case multiply:
			newTarget = target / left
		case divide:
			newTarget = left / target
		default:
			panic("Unrechable")
		}
        if expressions[e.rightOperand].getName() == "humn" {
		  return newTarget
        } else {
            return expressions[e.rightOperand].reverseSolve(expressions, newTarget)
        }

	} else {
        panic("Unrechable")
	}
}

func parseIntExpression(input string) intExpression {
	split := strings.Split(input, " ")

	name := strings.TrimSuffix(split[0], ":")
	value, err := strconv.Atoi(split[1])
	if err != nil {
		panic(err)
	}

	return intExpression{
		name:  name,
		value: int64(value),
	}
}

func parseOperatorExpression(input string) operatorExpression {
	split := strings.Split(input, " ")

	name := strings.TrimSuffix(split[0], ":")

	var op operation
	switch split[2] {
	case "+":
		op = add
	case "-":
		op = subtract
	case "*":
		op = multiply
	case "/":
		op = divide
	default:
		panic("Unreachable")
	}

	return operatorExpression{
		name:         name,
		operator:     op,
		leftOperand:  split[1],
		rightOperand: split[3],
	}
}

func parseInput(input []string) []expression {
	var result []expression
	for _, line := range input {

		if strings.ContainsAny(line, "+-*/") {
			result = append(result, parseOperatorExpression(line))
		} else {
			result = append(result, parseIntExpression(line))
		}

	}
	return result
}

func solvePart1(instructions []expression) int64 {

	exps := expressionMap{}

	for _, val := range instructions {
		exps[val.getName()] = val
	}

	return exps["root"].getValue(exps)
}
func solvePart2(instructions []expression) int64 {

	exps := expressionMap{}

	for _, val := range instructions {
		exps[val.getName()] = val
	}

	root := exps["root"].(operatorExpression)

	right := exps[root.rightOperand].getValue(exps)

	solution := exps[root.leftOperand].reverseSolve(exps, right)

	return solution
}

func main() {
	input := common.ReadInput("input.txt")
	parsedInput := parseInput(input)

	result1 := solvePart1(parsedInput)
	result2 := solvePart2(parsedInput)

	fmt.Printf("Part 1 result: %d\r\n", result1)
	fmt.Printf("Part 2 result: %d\r\n", result2)
}
