package day04

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	. "github.com/panshadow/aoc2022/utils"
)

func init() {
	RegisterTask("04/01", Task01, "04/01.twitter")
	RegisterTask("04/02", Task02, "04/01.twitter")
}

type Section struct {
	From int
	To int
}

func (s *Section) Size() int {
	return (s.To - s.From + 1)
}

func (s *Section) GE(s2 *Section) bool {
	return s.Size() >= s2.Size()
}

func (s *Section) Included(s2 *Section) bool {
	return (s.From <= s2.From && s.To >= s2.To)
}

func (s *Section) Before(s2 *Section) bool {
	return s.From <= s2.From && s.To <= s2.To
}

func (s *Section) Overlaped(s2 *Section) bool {
	if s.Before(s2) {
		return s.To >= s2.From
	} else {
		return s2.To >= s.From
	}
}

func (s *Section) String() string {
	return fmt.Sprintf("%d - %d (%d)", s.From, s.To, s.Size())
}


func parseSection(pair string) *Section {
	bound := strings.Split(pair, "-")
	from, err := strconv.Atoi(bound[0])
	if err != nil {
		log.Fatalf("Can't convert left bound %s of section: %s: %v", pair, bound[0], err)
	}
	to, err := strconv.Atoi(bound[1])
	if err != nil {
		log.Fatalf("Can't convert right bound %s of section: %s: %v", pair, bound[1], err)
	}
	return &Section{from, to}
}

func parsePairs(line string) (*Section,*Section) {
	p := strings.Split(line, ",")
	return parseSection(p[0]), parseSection(p[1])
}

func Task01(input []string) string {
	var result int
	for _, line := range input {
		s1, s2 := parsePairs(line)
		fmt.Printf("\n%s\n%s\n", s1, s2)
		s1ContainS2 := s1.GE(s2) && s1.Included(s2)
		s2ContainS1 := s2.GE(s1) && s2.Included(s1)
		if  s1ContainS2 || s2ContainS1 {
			result++
			fmt.Printf("Detect overlap: +1 = %d\n",result)
		}
	}
	return fmt.Sprint(result)
}

func Task02(input []string) string {
	var result int
	for _, line := range input {
		s1, s2 := parsePairs(line)
		fmt.Printf("\n%s\n%s\n", s1, s2)
		if  s1.Overlaped(s2) {
			result++
			fmt.Printf("Detect overlap: +1 = %d\n",result)
		}
	}
	return fmt.Sprint(result)
}
