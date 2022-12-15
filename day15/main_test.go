package main

import "testing"

var testBeacons = []sensor{
	{
		position:      point{2, 18},
        sensorRange: 7,
		closestBeacon: point{-2, 15},
	},
    {
        position:      point{9, 16},
        sensorRange: 1,
        closestBeacon: point{10, 16},
    },
    {
        position:      point{13, 2},
        sensorRange: 3,
        closestBeacon: point{15, 3},
    },
    {
        position:      point{12, 14},
        sensorRange: 4,
        closestBeacon: point{10, 16},
    },
    {
        position:      point{10, 20},
        sensorRange: 4,
        closestBeacon: point{10, 16},
    },
    {
        position:      point{14, 17},
        sensorRange: 5,
        closestBeacon: point{10, 16},
    },
    {
        position:      point{8, 7},
        sensorRange: 9,
        closestBeacon: point{2, 10},
    },
    {
        position:      point{2, 0},
        sensorRange: 10,
        closestBeacon: point{2, 10},
    },
    {
        position:      point{0, 11},
        sensorRange: 3,
        closestBeacon: point{2, 10},
    },
    {
        position:      point{20, 14},
        sensorRange: 8,
        closestBeacon: point{25, 17},
    },
    {
        position:      point{17, 20},
        sensorRange: 6,
        closestBeacon: point{21, 22},
    },
    {
        position:      point{16, 7},
        sensorRange: 5,
        closestBeacon: point{15, 3},
    },
    {
        position:      point{14, 3},
        sensorRange: 1,
        closestBeacon: point{15, 3},
    },
    {
        position:      point{20, 1},
        sensorRange: 7,
        closestBeacon: point{15, 3},
    },
}

func TestSolvePart1(t *testing.T) {
    testResult := solvePart1(testBeacons, 10)

	if testResult != 26 {
		t.Errorf("Expected: %d got %d", 26, testResult)
	}
}

func TestSolvePart2(t *testing.T) {
    testResult := solvePart2(testBeacons, 20)

    if testResult != 56000011 {
        t.Errorf("Expected: %d got %d", 56000011, testResult)
	}
}
