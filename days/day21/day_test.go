package day21

import (
	"testing"

	. "github.com/panshadow/aoc2022/utils"
)

var (

	input = SplitText(`# day21
`)
	expected01 = ""
	expected02 = ""
)

func TestTask01(t *testing.T) {
	Is(t, Task01(input), expected01)
}

func TestTask02(t *testing.T) {
	Is(t, Task02(input), expected02)
}
