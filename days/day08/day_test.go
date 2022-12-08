package day08

import (
	"testing"

	. "github.com/panshadow/aoc2022/utils"
)

var (

	input = SplitText(`# day08
30373
25512
65332
33549
35390
`)
	expected01 = "21"
	expected02 = "8"
)

func TestTask01(t *testing.T) {
	Is(t, Task01(input), expected01)
}

func TestTask02(t *testing.T) {
	Is(t, Task02(input), expected02)
}
