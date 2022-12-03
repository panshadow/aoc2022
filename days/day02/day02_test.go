package day02

import (
	"testing"

	"github.com/panshdw/aoc2022/utils"
)

var (
	input = utils.SplitText(`#
A Y
B X
C Z
`)
)

func TestTask01(t *testing.T) {
	if fail := utils.RunTest(input, Task01, "15"); fail != "" {
		t.Fatal(fail)
	}
}

func TestTask02(t *testing.T) {
	if fail := utils.RunTest(input, Task02, "12"); fail != "" {
		t.Fatal(fail)
	}
}
