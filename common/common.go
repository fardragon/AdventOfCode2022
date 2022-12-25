package common

import (
	"bufio"
	"os"
    "sort"
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

func (s Set[T]) Remove(elem T) {
    delete(s, elem)
}

func (s Set[T]) Len() int {
    return len(s)
}

func (s Set[T]) Empty() bool {
    return len(s) == 0
}

func (s Set[T]) Clone() Set[T] {
    result := NewSet[T]()
    for val := range s {
        result.Add(val)
    }
    return result
}

type OrderedStringSet []string
func (s *OrderedStringSet) Add(elem string) {

    if s.Contains(elem) {
        return
    }
    *s = append(*s, elem)
    sort.Strings(*s)
}

func (s *OrderedStringSet) Contains(elem string) bool {

    for _, item := range *s {
        if item == elem {
            return true
        }
    }
    return false
}

func (s *OrderedStringSet) Remove(elem string) {

    index := -1
    for ix, item := range *s {
        if item == elem {
            index = ix
            break
        }
    }
    if index == -1 {
        return
    }

    ret := make([]string, 0)
    ret = append(ret, (*s)[:index]...)
    ret = append(ret, (*s)[index+1:]...)

    *s = ret
}

func (s *OrderedStringSet) Len() int {
    return len(*s)
}

func (s *OrderedStringSet) Empty() bool {
    return s.Len() == 0
}

func Abs(a int) int {
    if a < 0 {
        return -a
    } else {
        return a
    }
}

func Abs64(a int64) int64 {
    if a < 0 {
        return -a
    } else {
        return a
    }
}

func Max(first int, rest ...int) int {
    m := first
    for _, num := range rest {
        if num > m {
            m = num
        }
    }
    return m
}
