package day08

import (
	"fmt"

	. "github.com/panshadow/aoc2022/utils"
)

func init() {
	RegisterTask("08/01", Task01, "08/01.twitter")
	RegisterTask("08/02", Task02, "08/01.twitter")
}

func parseRow(row string) []int {
	out := make([]int, len(row))
	for i, ch := range row {
		out[i] = int(ch - '0')
	}
	return out
}

func parseMap(rows []string) [][]int {
	out := make([][]int, len(rows))
	for i, row := range rows {
		out[i] = parseRow(row)
	}
	return out
}

func newMap(rows, cols int) [][]int {
	out := make([][]int, rows)
	for i := range out {
		out[i] = make([]int, cols)
	}
	return out
}

func isVisible(treeMap [][]int, r, c int) func(int, int) bool {
	H := len(treeMap)
	W := len(treeMap[0])
	return func(dr, dc int) bool {
		if dr != 0 {
			for i := r + dr; i >= 0 && i < H; i += dr {
				if treeMap[i][c] >= treeMap[r][c] {
					return false
				}
			}
		} else if dc != 0 {
			for j := c + dc; j >= 0 && j < W; j += dc {
				if treeMap[r][j] >= treeMap[r][c] {
					return false
				}
			}

		}
		fmt.Printf("Tree %d (%d:%d) from %d:%d is visible\n",treeMap[r][c],r,c,dr,dc)
		return true
	}
}

func visDistance(treeMap [][]int, r, c int) func(int, int) int {
	H := len(treeMap)
	W := len(treeMap[0])
	return func(dr, dc int) int {
		var d int
		if dr != 0 {
			for i := r + dr; i >= 0 && i < H; i += dr {
				d++
				if treeMap[i][c] >= treeMap[r][c] {
					fmt.Printf("Tree %d (%d:%d) from %d:%d blocked at (%d:%d) and has visDistance %d\n",treeMap[r][c],r,c,i,c,dr,dc,d)
					return d
				}
			}
		} else if dc != 0 {
			for j := c + dc; j >= 0 && j < W; j += dc {
				d++
				if treeMap[r][j] >= treeMap[r][c] {
					fmt.Printf("Tree %d (%d:%d) from %d:%d blocked at (%d:%d) and has visDistance %d\n",treeMap[r][c],r,c,r,j,dr,dc,d)
					return d
				}
			}

		}
		fmt.Printf("Tree %d (%d:%d) from %d:%d has visDistance %d\n",treeMap[r][c],r,c,dr,dc,d)
		return d
	}
}


func Task01(input []string) string {
	treeMap := parseMap(input)
	fmt.Print(treeMap)
	var count int

	for i, row := range treeMap {
		for j := range row {
			visible := isVisible(treeMap, i, j)
			if visible(-1, 0) || visible(1, 0) || visible(0, -1) || visible(0, 1) {
				count++
			}
		}
	}

	return fmt.Sprint(count)

}

func Task02(input []string) string {
	treeMap := parseMap(input)
	fmt.Print(treeMap)
	var max int

	for i, row := range treeMap {
		for j := range row {
			visdis := visDistance(treeMap, i, j)
			score := visdis(-1, 0)*visdis(1, 0)*visdis(0, -1)*visdis(0, 1)
			if score > max {
				max = score
			}
		}
	}

	return fmt.Sprint(max)
}
