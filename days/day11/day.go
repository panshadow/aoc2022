package day11

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	. "github.com/panshadow/aoc2022/utils"
)

func init() {
	RegisterTask("11/01", Task01, "11/01.twitter")
	RegisterTask("11/02", Task02, "11/01.twitter")
}

type Monkey struct {
	ID string
	Items []uint64
	Op func(old uint64) uint64
	TestCond uint64
	TestY string
	TestN string
	ICounter int
}

func Operations(op string, sArgs ...string) func(uint64) uint64 {
	return func(old uint64) uint64 {
		args := make([]uint64, len(sArgs))
		for i,a := range sArgs{
			if a == "old" {
				args[i] = old
			} else {
				val, err := strconv.Atoi(a)
				if err != nil {
					log.Fatalf("Invalid value in operations: %s(%s): %v", op, strings.Join(sArgs, ", "), err)
				}
				args[i] = uint64(val)
			}
		}
		switch op {
		case "*":
			return args[0]*args[1]
		case "+":
			return args[0]+args[1]
		}
		log.Fatalf("Unknown operation %s", op)
		return 0
	}
}

func NewMonkey(input []string) *Monkey {
	m := new(Monkey)
	for i,line := range input {
		fields := strings.FieldsFunc(line, func(ch rune) bool {
			return (ch == ' ' || ch == ',' || ch == ':')
		})
		switch  {
		case fields[0] == "Monkey":
			m.ID = fields[1]
		case fields[0] == "Starting":
			m.Items = make([]uint64, len(fields)-2)
			for j, sItem := range fields[2:] {
				val, err := strconv.Atoi(sItem)
				if err != nil {
					log.Fatalf("Invalid starting item: %s at line %d for %s: %v",
								sItem,
								i,
								input[0],
								err)
				}
				m.Items[j] = uint64(val)
			}
		case fields[0] == "Operation":
			m.Op = Operations(fields[4], fields[3], fields[5])
		case fields[0] == "Test":
			val, err := strconv.Atoi(fields[3])
			if err != nil {
				log.Fatalf("Invalid divisible value: %s at line %d: %v", fields[3], i, err)
			}
			m.TestCond = uint64(val)
		case fields[0] == "If" && fields[1] == "true":
			m.TestY = fields[5]
		case fields[0] == "If" && fields[1] == "false":
			m.TestN = fields[5]
		}
	}

	return m
}

func (m *Monkey) ThrowItem(level uint64) {
	m.Items = append(m.Items, level)
}

func (m *Monkey) String() string {
	return fmt.Sprintf("Monkey %s (%04d): %v", m.ID, m.ICounter, m.Items)
}

func Inspect(reliefScale uint64, monkeys map[string]*Monkey, id string, ) {
	Debugf("Monkey %s", id)
	m := monkeys[id]
	for _, item := range m.Items {
		Debugf("  Monkey inspects an item with a worry level of %d.\n",item)
		newLevel := m.Op(item)
		Debugf("    Worry level changed to %d.\n", newLevel)
		if reliefScale != 1 {
			newLevel = newLevel/reliefScale
			Debugf("    Monkey gets bored with item. Worry level is divided by %d to %d.\n", reliefScale, newLevel)
		} else {
			var prodAllTests uint64 = 1
			for mid := range monkeys {
				prodAllTests *= monkeys[mid].TestCond
			}
			newLevel = newLevel % prodAllTests

		}
		var throwTo string
		if newLevel % m.TestCond == 0 {
			Debugf("    Current worry level is divisible by %d.\n", m.TestCond)
			throwTo = m.TestY
		} else {
			Debugf("    Current worry level is not divisible by %d.\n", m.TestCond)
			throwTo = m.TestN
		}
		Debugf("   Item with worry level %d is thrown to monkey %s.\n", newLevel, throwTo)
		monkeys[throwTo].ThrowItem(newLevel)
	}
	m.ICounter+=len(m.Items)
	m.Items = []uint64{}
}

func Solution(input []string, reliefScale uint64, rounds int) int {
	monkeys := make(map[string]*Monkey)
	mids := []string{}
	from := 0
	for i,line := range input {
		if line == "" && from != i {
			m := NewMonkey(input[from:i])
			Debugf("Added monkey %s\n",m)
			from = i+1
			mids = append(mids, m.ID)
			monkeys[m.ID] = m
		}
	}
	if from < len(input) {
		m := NewMonkey(input[from:])
		Debugf("Added monkey %s",m)
		mids = append(mids, m.ID)
		monkeys[m.ID] = m
	}

	for round:=1;round <= rounds; round++ {
		Debugln("Round ",round)
		for _,mid := range mids {
			Inspect(reliefScale, monkeys, mid)
		}
		for _,mid := range mids {
			Debugln(monkeys[mid])
		}
	}
	max := make([]int, 2)
	for _, m := range monkeys {
		placed := false
		for i := range max {
			if !placed && m.ICounter > max[i] {
				placed = true
				InsertIntoSlice(max, i, m.ICounter)
			}
		}
	}
	return max[0]*max[1]
}

func Task01(input []string) string {
	return fmt.Sprint(Solution(input, 3, 20))
}

func Task02(input []string) string {
	return fmt.Sprint(Solution(input, 1, 10000))
}
