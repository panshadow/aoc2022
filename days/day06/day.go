package day06

import (
	"fmt"

	. "github.com/panshadow/aoc2022/utils"
)

func init() {
	RegisterTask("06/01", Task01, "06/01.twitter")
	RegisterTask("06/02", Task02, "06/01.twitter")
}

func detectMarker(size int, input string) int {
	fmt.Println("DS: ",input)
	chMap := make(map[byte]int)
	cur := 0
	for ; cur < size; cur ++ {
		chMap[input[cur]]++
	}
	fmt.Printf("After %d marker: %s\n", cur+1, input[cur-size+1:cur+1])
	if len(input[cur-size+1:cur+1])==len(chMap) {
		fmt.Println("Found in ", cur+1)
		return cur
	}
	for cur < len(input) {
		delete(chMap, input[cur-size])
		cur++
		nextCh := input[cur]
		chMap[nextCh]++
		fmt.Printf("After %d and %c marker: %s\n", cur+1, nextCh, input[cur-size+1:cur+1])
		if len(chMap) == size {
			fmt.Println("Found in ", cur+1)
			return cur
		}
	}
	return -1
}

func Task01(input []string) string {
	return fmt.Sprint(detectMarker(4, input[0]))
}

func Task02(input []string) string {
	return ""
}
