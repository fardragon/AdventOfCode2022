package main

import "testing"



var testCases = map[snafu]int {
    "1=-0-2": 1747,
    "12111": 906,
    "2=0=": 198,
    "21": 11,
    "2=01": 201,
    "111": 31,
    "20012": 1257,
    "112": 32,
    "1=-1=": 353,
    "1-12": 107,
    "12": 7,
    "1=": 3,
    "122": 37,
}



func TestSnafuToDecimal(t *testing.T) {
    for input, expectedResult := range testCases {
        testResult := snafuToDecimal(input)
        if testResult != expectedResult {
            t.Errorf("Expected: %d got %d", expectedResult, testResult)
        }
    }
}

func TestDecimalToSnafu(t *testing.T) {
    for expectedResult, input := range testCases {
        testResult := decimalToSnafu(input)
        if testResult != expectedResult {
            t.Errorf("Expected: %s got %s", expectedResult, testResult)
        }
    }
}

func TestSolve(t *testing.T) {
    var input []snafu
    for s := range testCases {
        input = append(input, s)
    }
    testResult := solve(input)
    if testResult != "2=-1=0" {
        t.Errorf("Expected: %s got %s", "2=-1=0", testResult)
    }

}


