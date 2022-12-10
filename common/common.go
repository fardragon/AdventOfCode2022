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

type Set[T comparable] map[T]struct{}
func NewSet[T comparable]() Set[T] {
    return make(Set[T])
}

func (s Set[T]) Add(elem T) {
    s[elem] = struct{}{}
}

func (s Set[T]) Contains(elem T) bool {
    _, ok := s[elem]
    return ok
}

func (s Set[T]) Len() int {
    return len(s)
}

func Abs(a int) int {
    if a < 0 {
        return -a
    } else {
        return a
    }
}