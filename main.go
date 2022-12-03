package main

import (
	"flag"
	"fmt"
	"log"

	_ "github.com/panshadow/aoc2022/days/day01"
	_ "github.com/panshadow/aoc2022/days/day02"
	_ "github.com/panshadow/aoc2022/days/day03"
	"github.com/panshadow/aoc2022/utils"
)
const (
	DATA_DIR = "./data/"
)

func main() {
	flag.Parse()
	fmt.Println("Advent Of Code 2022!")
	fmt.Printf("Found %d solutions\n", len(utils.TaskRegistry))

	if flag.NArg() > 0 {
		taskID := flag.Arg(0)
		task, found := utils.TaskRegistry[taskID]
		if !found {
			log.Fatalf("Task %s not found\n", taskID)
		}
		inputFile := fmt.Sprintf("%s%s", DATA_DIR, task.Input)
		input, err := utils.LoadInput(inputFile)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("Result: ", task.Code(input))
	}
}
