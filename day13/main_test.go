package main

import "testing"

var testInput = []packetPair{
	{
		first:  packet{data: listData{data: []data{integerData{data: 1}, integerData{data: 1}, integerData{data: 3}, integerData{data: 1}, integerData{data: 1}}}},
		second: packet{data: listData{data: []data{integerData{data: 1}, integerData{data: 1}, integerData{data: 5}, integerData{data: 1}, integerData{data: 1}}}}},
	{
		first:  packet{data: listData{data: []data{listData{data: []data{integerData{data: 1}}}, listData{data: []data{integerData{data: 2}, integerData{data: 3}, integerData{data: 4}}}}}},
		second: packet{data: listData{data: []data{listData{data: []data{integerData{data: 1}}}, integerData{data: 4}}}}},
	{
		first:  packet{data: listData{data: []data{integerData{data: 9}}}},
		second: packet{data: listData{data: []data{listData{data: []data{integerData{data: 8}, integerData{data: 7}, integerData{data: 6}}}}}}},
	{
		first:  packet{data: listData{data: []data{listData{data: []data{integerData{data: 4}, integerData{data: 4}}}, integerData{data: 4}, integerData{data: 4}}}},
		second: packet{data: listData{data: []data{listData{data: []data{integerData{data: 4}, integerData{data: 4}}}, integerData{data: 4}, integerData{data: 4}, integerData{data: 4}}}}},
	{
		first:  packet{data: listData{data: []data{integerData{data: 7}, integerData{data: 7}, integerData{data: 7}, integerData{data: 7}}}},
		second: packet{data: listData{data: []data{integerData{data: 7}, integerData{data: 7}, integerData{data: 7}}}}},
	{
		first:  packet{data: listData{data: []data(nil)}},
		second: packet{data: listData{data: []data{integerData{data: 3}}}}},
	{
		first:  packet{data: listData{data: []data{listData{data: []data{listData{data: []data(nil)}}}}}},
		second: packet{data: listData{data: []data{listData{data: []data(nil)}}}}},
	{
		first:  packet{data: listData{data: []data{integerData{data: 1}, listData{data: []data{integerData{data: 2}, listData{data: []data{integerData{data: 3}, listData{data: []data{integerData{data: 4}, listData{data: []data{integerData{data: 5}, integerData{data: 6}, integerData{data: 7}}}}}}}}}, integerData{data: 8}, integerData{data: 9}}}},
		second: packet{data: listData{data: []data{integerData{data: 1}, listData{data: []data{integerData{data: 2}, listData{data: []data{integerData{data: 3}, listData{data: []data{integerData{data: 4}, listData{data: []data{integerData{data: 5}, integerData{data: 6}, integerData{data: 0}}}}}}}}}, integerData{data: 8}, integerData{data: 9}}}}},
}

func TestSolvePart1(t *testing.T) {
	testResult := solvePart1(testInput)

	if testResult != 13 {
		t.Errorf("Expected: %d got %d", 13, testResult)
	}
}

func TestSolvePart2(t *testing.T) {
	testResult := solvePart2(testInput)

	if testResult != 140 {
		t.Errorf("Expected: %d got %d", 140, testResult)
	}
}
