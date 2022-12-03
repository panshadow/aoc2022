package day02

import (
	"fmt"
	"strings"

	"github.com/panshdw/aoc2022/utils"
)

func init() {
	utils.RegisterTask("02/01", Task01, "02/01.twitter")
	utils.RegisterTask("02/02", Task02, "02/01.twitter")
}

var (
	Lost = map[string]string{
		"A": "Z",
		"C": "Y",
		"B": "X",
	}

	Draw = map[string]string{
		"A": "X",
		"B": "Y",
		"C": "Z",
	}

	Win = map[string]string{
		"A": "Y",
		"B": "Z",
		"C": "X",
	}

	ScoresMap = map[string]int{
		"A":    1, // Rock
		"B":    2, // Paper
		"C":    3, // Scissors
		"X":    1,
		"Y":    2,
		"Z":    3,
		"win":  6,
		"draw": 3,
		"lost": 0,
	}
)

func Round(a, b string) string {
	switch {
	case Draw[a]==b:
		return "draw"
	case Lost[a] == b:
		return "lost"
	default:
		return "win"
	}
}

func Task01(input []string) string {
	var score int
	for _, round := range input {
		shapes := strings.Split(round, " ")
		result := Round(shapes[0], shapes[1])
		fmt.Println("round:",round, "shapes:", shapes, "res:", result)
		score += ScoresMap[shapes[1]] + ScoresMap[result]
	}
	return fmt.Sprintf("%d", score)
}

func Task02(input []string) string {
	var score int
	for _, round := range input {
		shapes := strings.Split(round, " ")
		var turn string
		switch shapes[1] {
		case "X":
			turn = Lost[shapes[0]]
		case "Y":
			turn = Draw[shapes[0]]
		case "Z":
			turn = Win[shapes[0]]
		}
		result := Round(shapes[0], turn)
		fmt.Println("round:",round, "shapes:", shapes, "my turn:", turn, "res:", result)
		score += ScoresMap[turn] + ScoresMap[result]
	}
	return fmt.Sprintf("%d", score)
}
