package day03

import (
	"testing"

	. "github.com/panshadow/aoc2022/utils"
)

var (
	input = SplitText(`#
vJrwpWtwJgWrhcsFMMfFFhFp
jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL
PmmdzqPrVvPwwTWBwg
wMqvLMZHhHMvwLHjbvcjnnSBnvTQFn
ttgJtRGJQctTZtZT
CrZsJsPPZsGzwwsLwLmpwMDw
`)
)

func TestPriority(t *testing.T) {
	Is(t, priority('a'), 1)
	Is(t, priority('A'), 27)
	Is(t, priority('z'), 26)
	Is(t, priority('Z'), 52)
}

func TestTask01(t *testing.T) {
	Is(t, Task01(input), "157")
}

func TestTask02(t *testing.T) {
	Is(t, Task02(input), "70")
}
