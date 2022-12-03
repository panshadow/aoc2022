package day01

import (
	"fmt"
	"log"
	"strconv"

	"github.com/panshdw/aoc2022/utils"
)

const (
	TOP_SIZE = 3
)

func init() {
	utils.RegisterTask("01/01", Task01, "01/01.twitter")
	utils.RegisterTask("01/02", Task02, "01/01.twitter")
}

func calcCallories(input []string) []int {
	out := []int{}
	var next int
	for i, line := range input {
		if line != "" {
			val, err := strconv.Atoi(line)

			if err != nil {
				log.Fatalf("Wrong number at line %d: %v ", i, err)
			}
			next += val
		} else {
			out = append(out, next)
			next = 0
		}
	}

	if next > 0 {
		out = append(out, next)
	}

	return out

}

func Task01(input []string) string {
	elves := calcCallories(input)
	max := elves[0]
	for _, val := range elves {
		if val > max {
			max = val
		}
	}
	return fmt.Sprintf("%d", max)
}

func Task02(input []string) string {
	elves := calcCallories(input)
	max := make([]int, TOP_SIZE)
	for _, val := range elves {
		placed := false
		for i := range max {
			if !placed && val > max[i] {
				placed = true
				utils.InsertIntoSlice(max, i, val)
			}
		}
	}
	var sum int
	for _, val := range max {
		sum += val
	}
	return fmt.Sprintf("%d", sum)
}
