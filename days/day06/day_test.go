package day06

import (
	"fmt"
	"testing"

	. "github.com/panshadow/aoc2022/utils"
)

var (
	inputMap1 = map[string]string{
		"mjqjpqmgbljsphdztnvjfqwrcgsmlb":    "7",
		"bvwbjplbgvbhsrlpgdmjqwftvncz":      "5",
		"nppdvjthqldpwncqszvftbrmjlhg":      "6",
		"nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg": "10",
		"zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw":  "11",
	}
	inputMap2 = map[string]string{
		"mjqjpqmgbljsphdztnvjfqwrcgsmlb":    "19",
		"bvwbjplbgvbhsrlpgdmjqwftvncz":      "23",
		"nppdvjthqldpwncqszvftbrmjlhg":      "23",
		"nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg": "29",
		"zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw":  "26",
	}
)

func TestIsMarker(t *testing.T) {
	fmt.Println("Testing marker detector")
	Is(t, isMarker("abcd"), true)
	Is(t, isMarker("abca"), false)
	Is(t, isMarker("aaaa"), false)
}

func TestTask01(t *testing.T) {
	for input, expected := range inputMap1 {
		Is(t, Task01([]string{input}), expected)
	}
}

func TestTask02(t *testing.T) {
	for input, expected := range inputMap2 {
		Is(t, Task02([]string{input}), expected)
	}
}
