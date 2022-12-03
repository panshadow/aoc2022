package day03

import (
	"fmt"
	"strings"

	. "github.com/panshadow/aoc2022/utils"
)

func init() {
	RegisterTask("03/01", Task01, "03/01.twitter")
	RegisterTask("03/02", Task02, "03/01.twitter")
}

func priority(item rune) int {
	if item >= 'A' && item <= 'Z' {
		return int(27+item - 'A' )
	} else {
		return int(1+item-'a')
	}
}

func getCommonItems(c1, c2 string, firstOnly bool) []rune {
	out := []rune{}
	for _, r := range []rune(c2) {
		if strings.IndexRune(c1, r) > -1 {
			if firstOnly {
				return []rune{r}
			} else if strings.IndexRune(string(out), r) == -1 {
				out = append(out, r)
			}
		}
	}
	return out
}

func Task01(input []string) string {
	var sum int
	for i, rucksack := range input {
		size := len(rucksack)
		half := size >> 1
		fmt.Printf("%d rucksacks contains %d (%s %d + %s %d) items\n", i, size, rucksack[:half], half, rucksack[half:], half)
		items := getCommonItems(rucksack[:half], rucksack[half:], true)
		if items != nil {
			itemp := priority(items[0])
			sum += itemp
			fmt.Printf("Item %c appears in both comparments: +%d = %d\n", items[0], itemp, sum)
		}
	}
	return fmt.Sprintf("%d", sum)
}

func getBadgeOfGroup(rucksacks []string) rune {
	common := getCommonItems(rucksacks[0], rucksacks[1], false)
	items := getCommonItems(rucksacks[2], string(common), true)
	return items[0]
}

func Task02(input []string) string {
	var sum int
	for i:=0; i<len(input); i+=3 {
		fmt.Printf("Group:\n %s\n",strings.Join(input[i:i+3],"\n"))
		item := getBadgeOfGroup(input[i:i+3])
		itemp := priority(item)
		sum += itemp
		fmt.Printf("Common item in group: %c +%d = %d\n", item, itemp, sum )
	}
	return fmt.Sprintf("%d", sum)
}
