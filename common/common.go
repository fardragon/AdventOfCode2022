package common

import (
	"bufio"
	"os"
)

func ReadInput(path string) []string {
	file, error := os.Open(path)
	check(error)
    defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

    var result []string
	for scanner.Scan() {
        result = append(result, scanner.Text())
	}
    return result
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}
