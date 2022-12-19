package main

import (
	"flag"
	"fmt"
	"log"
	"sort"

	_ "github.com/panshadow/aoc2022/days/day01"
	_ "github.com/panshadow/aoc2022/days/day02"
	_ "github.com/panshadow/aoc2022/days/day03"
	_ "github.com/panshadow/aoc2022/days/day04"
	_ "github.com/panshadow/aoc2022/days/day05"
	_ "github.com/panshadow/aoc2022/days/day06"
	_ "github.com/panshadow/aoc2022/days/day07"
	_ "github.com/panshadow/aoc2022/days/day08"
	_ "github.com/panshadow/aoc2022/days/day09"
	_ "github.com/panshadow/aoc2022/days/day10"
	_ "github.com/panshadow/aoc2022/days/day11"
	_ "github.com/panshadow/aoc2022/days/day12"
	_ "github.com/panshadow/aoc2022/days/day13"
	_ "github.com/panshadow/aoc2022/days/day14"
	_ "github.com/panshadow/aoc2022/days/day15"
	_ "github.com/panshadow/aoc2022/days/day16"
	_ "github.com/panshadow/aoc2022/days/day17"
	"github.com/panshadow/aoc2022/utils"
)
const (
	DATA_DIR = "./data/"
)

func main() {
	var inputFile string
	flag.StringVar(&inputFile, "i", "", "custom input file")
	flag.Parse()

	fmt.Println("Advent Of Code 2022!")
	fmt.Printf("Found %d solutions\n", len(utils.TaskRegistry))
	utils.Debugln("INPUT FILE:",inputFile)

	if flag.NArg() > 0 {
		taskID := flag.Arg(0)
		task, found := utils.TaskRegistry[taskID]
		if !found {
			log.Fatalf("Task %s not found\n", taskID)
		}

		if inputFile == "" {
			inputFile = fmt.Sprintf("%s%s", DATA_DIR, task.Input)
		}
		input, err := utils.LoadInput(inputFile)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("Result: ", task.Code(input))
	} else {
		allTasks := []string{}
		for taskID := range utils.TaskRegistry {
			allTasks = append(allTasks, taskID)
		}
		sort.Strings(allTasks)
		for _,taskID := range allTasks {
			fmt.Println(" - ",taskID)
		}
	}
}
