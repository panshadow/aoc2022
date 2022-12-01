package utils

import (
	"fmt"
	"io/ioutil"
	"strings"
)

type TaskSolutionFn func([]string) string

type TaskSolution struct {
	Code       TaskSolutionFn
	InputAlias string
}

var (
	TaskRegistry = map[string]TaskSolution{}
)

func SplitText(text string) []string {
	return strings.Split(text, "\n")
}

func LoadInput(fname string) ([]string, error) {
	var out []string
	body, err := ioutil.ReadFile(fname)

	if err != nil {
		return out, fmt.Errorf("Can't open file: %s: %v", fname, err)
	}
	return SplitText(string(body)), nil
}

func InsertIntoSlice(xs []int, pos int, x int) {
	for i := len(xs) - 1; i > pos; i-- {
		xs[i] = xs[i-1]
	}
	xs[pos] = x
}

func EqSlice[T comparable](a []T, b []T) bool {
	if len(a) != len(b) {
		return false
	}

	for i,xa := range a {
		if xa != b[i] {
			return false
		}
	}

	return true
}
