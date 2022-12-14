package day14

import (
	"testing"

	. "github.com/panshadow/aoc2022/utils"
)

var (

	input = SplitText(`# day14
498,4 -> 498,6 -> 496,6
503,4 -> 502,4 -> 502,9 -> 494,9`)
	expected01 = "24"
	expected02 = ""
)

func TestTask01(t *testing.T) {
	Is(t, Task01(input), expected01)
}

func TestTask02(t *testing.T) {
	Is(t, Task02(input), expected02)
}
