package main

import (
	"flag"
	"fmt"
	"log"

	_ "github.com/panshdw/aoc2022/days"
	"github.com/panshdw/aoc2022/utils"
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
		if task.InputAlias != "" {
			taskID = task.InputAlias
		}
		inputFile := fmt.Sprintf("%s%s", DATA_DIR, taskID)
		input, err := utils.LoadInput(inputFile)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("Result: ", task.Code(input))
	}
}
