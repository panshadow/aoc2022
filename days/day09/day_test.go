package day09

import (
	"testing"

	. "github.com/panshadow/aoc2022/utils"
)

var (

	input = SplitText(`# day09
R 4
U 4
L 3
D 1
R 4
D 1
L 5
R 2
`)
	input2 = SplitText(`# day09
R 5
U 8
L 8
D 3
R 17
D 10
L 25
U 20
`)
	expected01 = "13"
	expected02 = "36"
)

func TestTask01(t *testing.T) {
	Is(t, Task01(input), expected01)
}

func TestTask02(t *testing.T) {
	Is(t, Task02(input2), expected02)
}
