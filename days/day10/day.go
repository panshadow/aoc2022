package day10

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	. "github.com/panshadow/aoc2022/utils"
)

func init() {
	RegisterTask("10/01", Task01, "10/01.twitter")
	RegisterTask("10/02", Task02, "10/01.twitter")
}

type CPU struct {
	RegX int
	CurCycle int
	Watchers map[int]int
	CRT []rune
}

const (
	CRT_W = 40
	CRT_H = 6
)

var (
	WatchCycles1 = [...]int{20,60,100,140,180,220}
)

func NewCPU() *CPU {
	cpu := new(CPU)
	cpu.RegX = 1
	cpu.CurCycle = 1
	cpu.Watchers = make(map[int]int)
	cpu.CRT = make([]rune,CRT_W*CRT_H)

	return cpu
}

func (c *CPU) WatchCycle(cycle int) {
	c.Watchers[cycle] = 0
}

func (c *CPU) SignalStrength(cycle int) int {
	if val, found := c.Watchers[cycle]; found && cycle < c.CurCycle {
		return val
	}
	return 0
}

func (c *CPU) Cycle() {
	Debugf("(%04d) ::",c.CurCycle)
	if _, watch := c.Watchers[c.CurCycle]; watch {
		signalStrength := c.CurCycle * c.RegX
		c.Watchers[c.CurCycle] = signalStrength
		Debug("W", signalStrength)
	}

	spriteMid := c.RegX % CRT_W
	crtPos := c.CurCycle-1
	if crtPos%CRT_W >= spriteMid-1 && crtPos%CRT_W <=spriteMid+1 {
		c.CRT[crtPos] = '#'
	} else {
		c.CRT[crtPos] = '.'
	}
	c.CurCycle++
	Debugln("::")
}

func (c *CPU) handleAddX(val int) {
	c.Cycle()
	c.Cycle()
	c.RegX += val
}

func (c *CPU) handleNoop() {
	c.Cycle()
}

func parseLine(cpu *CPU, line string) {
	fields := strings.Fields(line)
	cur := cpu.CurCycle
	switch fields[0] {
	case "noop":
		cpu.handleNoop()
	case "addx":
		if len(fields)<2 {
			log.Fatalf("Missed arguments: %s", line)
		}

		val, err := strconv.Atoi(fields[1])
		if err != nil {
			log.Fatalf("Invalid line: %s: %s: %v", line, fields[1], err)
		}

		cpu.handleAddX(val)
	}
	Debugf("[%04d]: %10s X:[%-4d]\n",cur, line, cpu.RegX)
}

func (c *CPU) ShowCRT() string {
	out := make([]string,CRT_H+1)
	for r := 0; r<CRT_H; r++ {
		out[r+1] = string(c.CRT[r*CRT_W:(r+1)*CRT_W])
	}
	return strings.Join(out, "\n")
}

func Solution(cpu *CPU, input []string) {
	for i,line := range input {
		Debugf("Parse %d row\n", i)
		parseLine(cpu, line)
	}
}

func Task01(input []string) string {
	cpu := NewCPU()

	for _,cycle := range WatchCycles1 {
		cpu.WatchCycle(cycle)
	}
	Solution(cpu, input)

	var result int
	for _, cycle := range WatchCycles1 {
		result += cpu.SignalStrength(cycle)
	}

	return fmt.Sprint(result)
}

func Task02(input []string) string {
	cpu := NewCPU()

	Solution(cpu, input)

	return cpu.ShowCRT()
}
