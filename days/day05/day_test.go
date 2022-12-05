package day05

import (
	"testing"

	. "github.com/panshadow/aoc2022/utils"
)

var (
	input = SplitText(`# day05
    [D]
[N] [C]
[Z] [M] [P]
 1   2   3

move 1 from 2 to 1
move 3 from 1 to 3
move 2 from 2 to 1
move 1 from 1 to 2
`)
)

func TestTask01(t *testing.T) {
	Is(t, Task01(input), "CMZ")
}

func TestTask02(t *testing.T) {
	Is(t, Task02(input), "MCD")
}
