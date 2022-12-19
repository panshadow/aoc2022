package day16

import (
	"fmt"
	"log"
	"strconv"

	. "github.com/panshadow/aoc2022/utils"
)

func init() {
	RegisterTask("16/01", Task01, "16/01.twitter")
	RegisterTask("16/02", Task02, "16/01.twitter")
}

type Valve struct {
	ID       string
	FlowRate int
	IsOpen   bool
	Linked   []*Valve
}

func NewValve(id string, flowRate int) *Valve {
	return &Valve{
		ID:       id,
		FlowRate: flowRate,
	}
}

func (v *Valve) Link(links ...*Valve) {
	v.Linked = make([]*Valve, len(links))
	copy(v.Linked, links)
}

func (v *Valve) SetFlowRate(flowRate int) *Valve {
	v.FlowRate = flowRate
	return v
}

func (v *Valve) Open() *Valve {
	v.IsOpen = true

	return v
}

func (v *Valve) ReleasePressure() int {
	if v.IsOpen {
		return v.FlowRate
	}
	return 0
}

type Cave struct {
	Valves   map[string]*Valve
	Timer    int
	MaxTime  int
	Pressure int
	Curr     *Valve
}

func parseLine(line string) (string, int, []string) {
	fields := Tokens(line, " =;,")
	flowRate, err := strconv.Atoi(fields[5])
	if err != nil {
		log.Fatalf("Invalid flow rate: %s in line %s: %v", fields[5], line, err)
	}

	return fields[1], flowRate, fields[10:]
}

func NewCave(input []string, maxTime int, initial string) *Cave {
	cave := &Cave{
		Valves: make(map[string]*Valve),
		MaxTime: maxTime,
		Timer: maxTime,
	}

	for i, line := range input {
		id, flowRate, links := parseLine(line)
		Debugf("%d) Create Valve(ID=%s, FlowRate=%d) linked to (%s)\n", i, id, flowRate, JoinListString(links, ", "))
		vlinks := make([]*Valve, len(links))
		for j, vid := range links {
			if v, exists := cave.Valves[vid]; exists {
				vlinks[j] = v
			} else {
				v = NewValve(vid, 0)
				cave.Valves[vid] = v
				vlinks[j] = v
			}
		}
		if v, exists := cave.Valves[id]; exists {
			v.SetFlowRate(flowRate).Link(vlinks...)
		} else {
			v = NewValve(id, flowRate)
			v.Link(vlinks...)
			cave.Valves[id] = v
		}
	}
	curr, exists := cave.Valves[initial]
	if !exists {
		log.Fatalf("Not found initial valve: %s", initial)
	}
	cave.Curr = curr

	return cave
}

func (c *Cave) ReleasePressure() int {
	cnt := 0
	openValves := make([]string, 0, len(c.Valves))
	for _, v := range c.Valves {
		if v.IsOpen {
			openValves = append(openValves, v.ID)
			cnt += v.ReleasePressure()
		}
	}
	ovNum := len(openValves)
	switch {
	case ovNum > 2:
		Debugf("Valves %s, and %s are open, releasing %d pressure.\n", JoinListString(openValves[:ovNum-1], ", "), openValves[ovNum-1], cnt)
	case ovNum > 1:
		Debugf("Valves %s and %s are open, releasing %d pressure.\n", openValves[0], openValves[1], cnt)
	case ovNum == 1:
		Debugf("Valves %s is open, releasing %d pressure.\n", openValves[0], cnt)
	}
	return cnt
}

func (c *Cave) MoveTo(next string) {
	for _, v := range c.Curr.Linked {
		if v.ID == next {
			Debugf("You move from %s to %s\n", c.Curr.ID, v.ID)
			c.Curr = v
		}
	}
}

func (c *Cave) OpenValve() {
	c.Curr.Open()
	Debugf("You open valve %s\n", c.Curr.ID)
}

func (c *Cave) TimeOver() bool {
	return c.Timer == 0
}

func (c *Cave) Loop(step func(c *Cave)) {
	if !c.TimeOver() {
		Debugf("== Minute %d ==\n", c.MaxTime-c.Timer+1)
		c.Timer--
		c.Pressure += c.ReleasePressure()
		step(c)
	} else {
		Debugln("== Time is over ==")
	}
}

func Task01(input []string) string {
	cave := NewCave(input, 30, "AA")
	Debugln("Start from ",cave.Curr.ID, len(cave.Curr.Linked))
	for !cave.TimeOver() {
		cave.Loop(func(c *Cave) {
			if c.Curr.IsOpen || c.Curr.FlowRate==0 {
				var foundNext *Valve
				maxFlowRate := 0
				for _,next := range c.Curr.Linked {
					Debugln("<=>",maxFlowRate,next.FlowRate)
					if !next.IsOpen && next.FlowRate > maxFlowRate {
						foundNext = next
						maxFlowRate = foundNext.FlowRate
					}
				}
				if foundNext != nil {
					Debugf("Found max next: %s(%d)\n",foundNext.ID,foundNext.FlowRate)
					cave.MoveTo(foundNext.ID)
				} else {
					Debugf("Move the last linked")
					cave.MoveTo(c.Curr.Linked[len(c.Curr.Linked)-1].ID)
				}
			} else {
				cave.OpenValve()
			}
		})
	}

	return fmt.Sprint(cave.Pressure)
}

func Task02(input []string) string {
	return ""
}
