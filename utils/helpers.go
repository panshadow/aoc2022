package utils

import (
	"fmt"
	"io/ioutil"
	"strings"
	"golang.org/x/exp/constraints"
)


func SplitText(text string) []string {
	rows := strings.Split(text, "\n")
	from := 0
	to := len(rows)
	if strings.HasPrefix(rows[0], "#")  {
		from = 1
	}
	if rows[to-1] == "" {
		to--
	}
	return rows[from:to]
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

func Ord[T constraints.Integer](x,y T) T {
	var one T = 1
	switch {
	case x>y:
		return one
	case x<y:
		return -one
	default:
		return 0
	}
}

func Abs[T constraints.Integer](x T) T {
	if x<0 {
		return -x
	}
	return x
}

func Max[T constraints.Integer](x T, xs ...T) T {
	max := x
	for i := range xs {
		if xs[i] > max {
			max = xs[i]
		}
	}
	return max
}

func Min[T constraints.Integer](x T, xs ...T) T {
	min := x
	for i := range xs {
		if xs[i] < min {
			min = xs[i]
		}
	}
	return min
}
