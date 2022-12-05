package day05

import (
	"log"
	"fmt"
	"strconv"
	"strings"

	. "github.com/panshadow/aoc2022/utils"
)

func init() {
	RegisterTask("05/01", Task01, "05/01.twitter")
	RegisterTask("05/02", Task02, "05/01.twitter")
}

type Stack struct {
	Buf []string
	Cur int
}

type Command struct {
	Num int
	Src int
	Dst int
}

func NewStack(size int) *Stack {
	buf := make([]string, size)
	return &Stack{buf, -1}
}

func (s *Stack) Push(val string) {
	s.Cur++
	s.Buf[s.Cur] = val
}

func (s *Stack) Pop() string {
	val := s.Buf[s.Cur]
	s.Cur--
	return val
}

func (s *Stack) PushN(vals []string) {
	for _,val := range vals {
		s.Push(val)
	}
}

func (s *Stack) PopN(n int) []string {
	fr := s.Cur-n+1
	to := s.Cur+1
	val := s.Buf[fr:to]
	s.Cur -= n
 	return val
}



func (s *Stack) String() string {
	return strings.Join(s.Buf[:s.Cur+1], ", ")
}

func parseStackRow(ss []*Stack, row string) {
	for i,s := range ss {
		from := i*4+1
		to := from+1
		if to <= len(row) {
			val := row[from:to]
			if val != " " {
				s.Push(val)
			}
		}
	}
}

func parseCommandRow(row string) *Command {
	fields := strings.Fields(row)
	num, err := strconv.Atoi(fields[1])
	if err != nil {
		log.Fatalf("Can't parse num field: %s\ncommand row: %s\nerror: %v", fields[1], row, err)
	}
	src, err := strconv.Atoi(fields[3])
	if err != nil {
		log.Fatalf("Can't parse src field: %s\ncommand row: %s\nerror: %v", fields[3], row, err)
	}
	dst, err := strconv.Atoi(fields[5])
	if err != nil {
		log.Fatalf("Can't parse num field: %s\ncommand row: %s\nerror: %v", fields[5], row, err)
	}

	return &Command{num, src-1, dst-1}
}

func Task01(input []string) string {
	var empty int
	for empty=0;input[empty] != ""; empty++ {
	}
	labels := strings.Fields(input[empty-1])
	stacks := make([]*Stack,len(labels))
	stackSize := len(labels)*(empty-1)
	for i := range labels {
		stacks[i] = NewStack(stackSize)
	}
	for i := empty-2; i>=0; i-- {
		parseStackRow(stacks, input[i])
	}
	for i,s := range stacks {
		fmt.Println(labels[i],s.String())
	}

	for i := empty+1; i<len(input); i++ {
		cmd := parseCommandRow(input[i])
		for j := 0; j<cmd.Num; j++ {
			val := stacks[cmd.Src].Pop()
			fmt.Printf("Move [%s]: %s -> %s\n",val,labels[cmd.Src], labels[cmd.Dst])
			stacks[cmd.Dst].Push(val)
		}
	}
	for i := range labels {
		labels[i] = stacks[i].Pop()
	}
	return strings.Join(labels, "")
}

func Task02(input []string) string {
	var empty int
	for empty=0;input[empty] != ""; empty++ {
	}
	labels := strings.Fields(input[empty-1])
	stacks := make([]*Stack,len(labels))
	stackSize := len(labels)*(empty-1)
	for i := range labels {
		stacks[i] = NewStack(stackSize)
	}
	for i := empty-2; i>=0; i-- {
		parseStackRow(stacks, input[i])
	}
	for i,s := range stacks {
		fmt.Println(labels[i],s.String())
	}

	for i := empty+1; i<len(input); i++ {
		cmd := parseCommandRow(input[i])
		vals := stacks[cmd.Src].PopN(cmd.Num)
		fmt.Printf("Move [%s]: %s -> %s\n",vals,labels[cmd.Src], labels[cmd.Dst])
		stacks[cmd.Dst].PushN(vals)
	}
	for i := range labels {
		labels[i] = stacks[i].Pop()
	}
	return strings.Join(labels, "")
}
