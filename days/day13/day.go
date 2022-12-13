package day13

import (
	"fmt"
	"log"
	"sort"
	"strconv"
	"strings"

	. "github.com/panshadow/aoc2022/utils"
)

func init() {
	RegisterTask("13/01", Task01, "13/01.twitter")
	RegisterTask("13/02", Task02, "13/01.twitter")
}

func cmp(x, y int) int {
	switch {
	case x < y:
		return -1
	case x > y:
		return 1
	}
	return 0
}

func CompareList(xi, yi interface{}) int {

	xs, xlist := xi.([]interface{})
	ys, ylist := yi.([]interface{})

	x, xval := xi.(int)
	y, yval := yi.(int)

	switch {
	case xval && yval:
		Debugln("both values are integers", x, y)
		return cmp(x, y)
	case xlist && yval:
		Debugln("left value is list and right is integer")
		return CompareList(xi, []interface{}{y})
	case xval && ylist:
		Debugln("left value is integer and right is list")
		return CompareList([]interface{}{x}, yi)
	case xlist && ylist:
		Debugln("both values are lists")
		switch {
		case len(xs) == 0 && len(ys) == 0:
			Debugln("both lists are empty")
			return 0
		case len(xs) == 0:
			Debugln("first list is empty")
			return -1
		case len(ys) == 0:
			Debugln("second list is empty")
			return 1
		default:
			Debugln("Compare items in lists")
			for i := range xs {
				if i < len(ys) {
					if res := CompareList(xs[i], ys[i]); res != 0 {
						return res
					}
				} else {
					Debugln("right list runs out of items first", i, len(ys))
					return 1
				}
			}
			if len(xs) < len(ys) {
				Debugln("left list runs out of items first")
				return -1
			}
		}
	default:
		Debugf("STRANGE VALS: %T %T\n", xi, yi)
	}
	return 0
}

func ParseList(list string) []interface{} {
	tokens := []string{}
	if len(list) > 0 {
		from := 0
		level := 0
		for i, ch := range list {
			switch ch {
			case ',':
				if level == 0 {
					tokens = append(tokens, list[from:i])
					from = i + 1
				}
			case '[':
				level++
				if level == 1 {
					from = i
				}
			case ']':
				level--
			}
		}
		tokens = append(tokens, list[from:])
	}
	out := make([]interface{}, len(tokens))
	for i, token := range tokens {
		if strings.HasPrefix(token, "[") && strings.HasSuffix(token, "]") {
			out[i] = ParseList(token[1:len(token)-1])
		} else {
			val, err := strconv.Atoi(token)
			if err != nil {
				log.Fatalf("Invalid token %s in list %s: %v ", token, list, err)
			}
			out[i] = val
		}
	}
	return out
}

func GetLists(input []string) [][]interface{} {
	lists := make([][]interface{},0,len(input))
	for _, line := range input {
		if line != "" {
			lists = append(lists, ParseList(line[1:len(line)-1]))
		}
	}
	return lists
}

func Task01(input []string) string {
	lists := GetLists(input)
	num := 1
	result := 0
	for i:=0; i<len(lists); i+=2 {
		if CompareList(lists[i], lists[i+1]) == -1 {
			result += num
		}
		num++
	}
	return fmt.Sprint(result)
}

func ShowList(list []interface{}) string {
	out := make([]string, len(list))
	for i,item := range list {
		switch v:=item.(type) {
		case int:
			out[i] = fmt.Sprint(v)
		case []interface{}:
			out[i] = ShowList(v)
		default:
			out[i] = fmt.Sprintf("[%T: %v]", item, item)
		}
	}

	return fmt.Sprintf("[%s]",strings.Join(out, ","))
}

func Task02(input []string) string {
	lists := GetLists(input)
	lists = append(lists, ParseList("[2]"))
	lists = append(lists, ParseList("[6]"))
	sort.Slice(lists, func(i, j int) bool {
		return CompareList(lists[i], lists[j]) == -1
	})
	result := 1
	for i := range lists {
		slist := ShowList(lists[i])
		if slist == "[[2]]" || slist == "[[6]]" {
			result *= i+1
		}
	}
	return fmt.Sprint(result)

}
