package day17

import (
	"testing"

	. "github.com/panshadow/aoc2022/utils"
)

var (

	input = SplitText(`# day17
>>><<><>><<<>><>>><<<>>><<<><<<>><>><<>>`)
	expected01 = "3068"
	expected02 = "1514285714288"

)

func TestTask01(t *testing.T) {
	Is(t, Task01(input), expected01)
}

func TestTask02(t *testing.T) {
	Is(t, Task02(input), expected02)
}
