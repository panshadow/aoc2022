package day13

import (
	"testing"

	. "github.com/panshadow/aoc2022/utils"
)

var (
	input = SplitText(`# day13
[1,1,3,1,1]
[1,1,5,1,1]

[[1],[2,3,4]]
[[1],4]

[9]
[[8,7,6]]

[[4,4],4,4]
[[4,4],4,4,4]

[7,7,7,7]
[7,7,7]

[]
[3]

[[[]]]
[[]]

[1,[2,[3,[4,[5,6,7]]]],8,9]
[1,[2,[3,[4,[5,6,0]]]],8,9]`)
	expected01 = "13"
	expected02 = "140"
)

func LST(xs ...interface{}) []interface{} {
	return xs
}

func TestCompareList(t *testing.T) {
	Is(t, CompareList(LST(1, 1, 3, 1, 1), LST(1, 1, 5, 1, 1)), -1)
	Is(t, CompareList(LST(), LST()), 0)
	Is(t, CompareList(LST(LST(1), LST(2, 3, 4)), LST(LST(1), 4)), -1)
	Is(t, CompareList(LST(9), LST(LST(8, 7, 6))), 1)
	Is(t, CompareList(LST(LST(4, 4), 4, 4), LST(LST(4, 4), 4, 4, 4)), -1)
	Is(t, CompareList(LST(7, 7, 7, 7), LST(7, 7, 7)), 1)
	Is(t, CompareList(LST(), LST(3)), -1)
	Is(t, CompareList(LST(LST(LST())), LST(LST())), 1)
	Is(t, CompareList(LST(1, LST(2, LST(3, LST(4, LST(5, 6, 7)))), 8, 9), LST(1, LST(2, LST(3, LST(4, LST(5, 6, 0)))), 8, 9)), 1)

}

func TestParseList(t *testing.T) {
	Is(t, CompareList(ParseList("1,2"), LST(1, 2)), 0)
	Is(t, CompareList(ParseList("1,2,[]"), LST(1, 2, LST())), 0)
	Is(t, CompareList(ParseList("1,[2,[3,[4,[5,6,7]]]],8,9"), LST(1, LST(2, LST(3, LST(4, LST(5, 6, 7)))), 8, 9)), 0)
}

func TestShowList(t *testing.T) {
	Is(t, ShowList(ParseList("1,[2,[3,[4,[5,6,7]]]],8,9")), "[1,[2,[3,[4,[5,6,7]]]],8,9]")
}

func TestTask01(t *testing.T) {
	Is(t, Task01(input), expected01)
}

func TestTask02(t *testing.T) {
	Is(t, Task02(input), expected02)
}
