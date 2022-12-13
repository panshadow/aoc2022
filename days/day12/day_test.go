package day12

import (
	"testing"

	. "github.com/panshadow/aoc2022/utils"
)

var (

	input = SplitText(`# day12
Sabqponm
abcryxxl
accszExk
acctuvwj
abdefghi`)
	expected01 = "31"
	expected02 = "29"
)

func TestTask01(t *testing.T) {
	Is(t, Task01(input), expected01)
}

func TestTask02(t *testing.T) {
	Is(t, Task02(input), expected02)
}
