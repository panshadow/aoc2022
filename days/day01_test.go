package day01

import (
	"testing"

	"github.com/panshdw/aoc2022/utils"
)

var (
	input = utils.SplitText(`1000
2000
3000

4000

5000
6000

7000
8000
9000

10000`)
)

func TestTask01(t *testing.T) {
	expected := "24000"
	result := Task01(input)
	if result != expected {
		t.Fatalf("%s != %s", result, expected)
	}
}

func TestPutInSorted(t *testing.T) {
	top := []int{0, 0, 0}
	utils.InsertIntoSlice(top, 0, 1)
	utils.InsertIntoSlice(top, 0, 2)
	utils.InsertIntoSlice(top, 0, 3)

	expect := []int{3, 2, 1}
	if !utils.EqSlice(top, expect) {
		t.Fatalf("%v != %v", top, expect)
	}
}

func TestTask02(t *testing.T) {
	expected := "45000"
	result := Task02(input)
	if result != expected {
		t.Fatalf("%s != %s", result, expected)
	}

}
