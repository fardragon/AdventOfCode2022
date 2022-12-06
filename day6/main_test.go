package main

import "testing"

type testCase struct {
	input       string
	part1Result int
	part2Result int
}

var testInput = []testCase{
	{
		"mjqjpqmgbljsphdztnvjfqwrcgsmlb",
		7,
		19,
	},
	{
		"bvwbjplbgvbhsrlpgdmjqwftvncz",
		5,
		23,
	},

	{
		"nppdvjthqldpwncqszvftbrmjlhg",
		6,
		23,
	},
	{
		"nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg",
		10,
		29,
	},
	{
		"zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw",
		11,
		26,
	},
}

func TestSolvePart1(t *testing.T) {
	for _, test := range testInput {
		testResult := solvePart1(test.input)

		if testResult != test.part1Result {
			t.Errorf("Expected: %d got %d", test.part1Result, testResult)
		}
	}
}

func TestSolvePart2(t *testing.T) {
	for _, test := range testInput {
		testResult := solvePart2(test.input)

		if testResult != test.part2Result {
			t.Errorf("Expected: %d got %d", test.part2Result, testResult)
		}
	}
}
