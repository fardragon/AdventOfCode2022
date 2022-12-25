package main

import "testing"

var testInput = []expression{
    operatorExpression{
        name:         "root",
        operator:     add,
        leftOperand:  "pppw",
        rightOperand: "sjmn",
    },
    intExpression{
        name:  "dbpl",
        value: 5,
    },
    operatorExpression{
        name:         "cczh",
        operator:     add,
        leftOperand:  "sllz",
        rightOperand: "lgvd",
    },
    intExpression{
        name:  "zczc",
        value: 2,
    },
    operatorExpression{
        name:         "ptdq",
        operator:     subtract,
        leftOperand:  "humn",
        rightOperand: "dvpt",
    },
    intExpression{
        name:  "dvpt",
        value: 3,
    },
    intExpression{
        name:  "lfqf",
        value: 4,
    },
    intExpression{
        name:  "humn",
        value: 5,
    },
    intExpression{
        name:  "ljgn",
        value: 2,
    },
    operatorExpression{
        name:         "sjmn",
        operator:     multiply,
        leftOperand:  "drzm",
        rightOperand: "dbpl",
    },
    intExpression{
        name:  "sllz",
        value: 4,
    },
    operatorExpression{
        name:         "pppw",
        operator:     divide,
        leftOperand:  "cczh",
        rightOperand: "lfqf",
    },
    operatorExpression{
        name:         "lgvd",
        operator:     multiply,
        leftOperand:  "ljgn",
        rightOperand: "ptdq",
    },
    operatorExpression{
        name:         "drzm",
        operator:     subtract,
        leftOperand:  "hmdt",
        rightOperand: "zczc",
    },
    intExpression{
        name:  "hmdt",
        value: 32,
    },
}

func TestSolvePart1(t *testing.T) {

	testResult := solvePart1(testInput)

	if testResult != 152 {
		t.Errorf("Expected: %d got %d", 152, testResult)
	}
}

func TestSolvePart2(t *testing.T) {

	testResult := solvePart2(testInput)

    if testResult != 301 {
        t.Errorf("Expected: %d got %d", 301, testResult)
	}
}
