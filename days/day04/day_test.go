package day04

import (
	"testing"

	. "github.com/panshadow/aoc2022/utils"
)

var (
	input = SplitText(`# day04 test input
2-4,6-8
2-3,4-5
5-7,7-9
2-8,3-7
6-6,4-6
2-6,4-8
`)
)

func TestTask01(t *testing.T) {
	Is(t, Task01(input), "2")
}

func TestTask02(t *testing.T) {
	Is(t, Task02(input), "4")
}
