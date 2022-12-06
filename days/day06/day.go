package day06

import (
	"fmt"

	. "github.com/panshadow/aoc2022/utils"
)

func init() {
	RegisterTask("06/01", Task01, "06/01.twitter")
	RegisterTask("06/02", Task02, "06/01.twitter")
}

func isMarker(marker string) bool {
	for i:=0;i<len(marker)-1;i++ {
		for j:=i+1;j<len(marker);j++ {
			if marker[i]==marker[j] {
				return false
			}
		}
	}
	return true
}

func detectMarker(size int, input string) int {
	fmt.Println("DS: ",input)
	cur := size
	fmt.Printf("After %d marker: %s\n", cur, input[cur-size:cur])
	if isMarker(input[cur-size:cur]) {
		fmt.Println("Found in ", cur)
		return cur
	}
	for cur < len(input) {
		oldCh := input[cur-size]
		nextCh := input[cur]
		cur++
		marker := input[cur-size:cur]
		fmt.Printf("%d %c <- [%s:%c]\n", cur, oldCh, marker[:size-1], nextCh)
		if isMarker(marker) {
			fmt.Println("Found in ", cur)
			return cur
		}
	}
	return -1
}

func Task01(input []string) string {
	return fmt.Sprint(detectMarker(4, input[0]))
}

func Task02(input []string) string {
	return fmt.Sprint(detectMarker(14, input[0]))
}
