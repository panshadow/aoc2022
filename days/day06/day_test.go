package day06

import (
	"testing"

	. "github.com/panshadow/aoc2022/utils"
)

var (
	inputMap = map[string]string{
		"mjqjpqmgbljsphdztnvjfqwrcgsmlb":    "7",
		"bvwbjplbgvbhsrlpgdmjqwftvncz":      "5",
		"nppdvjthqldpwncqszvftbrmjlhg":      "6",
		"nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg": "10",
		"zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw":  "11",
	}
	expected01 = ""
	expected02 = ""
)

func TestTask01(t *testing.T) {
	for input, expected := range inputMap {
		Is(t, Task01([]string{ input }), expected)
	}
}

func TestTask02(t *testing.T) {
	Is(t, Task02([]string{}), expected02)
}
