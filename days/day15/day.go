package day15

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	. "github.com/panshadow/aoc2022/utils"
)

func init() {
	RegisterTask("15/01", Task01, "15/01.twitter")
	RegisterTask("15/02", Task02, "15/01.twitter")
}

type Point struct {
	X int
	Y int
}


type SensorBeacon [2]Point

func (sb *SensorBeacon) Radius() int {
	return Abs(sb[0].X-sb[1].X) + Abs(sb[0].Y-sb[1].Y)
}

func (sb *SensorBeacon) Distance(y int) int {
	return Abs(sb[0].Y - y)
}

func (sb *SensorBeacon) CoverRow(y int) []Point {
	w := sb.Radius() - sb.Distance(y)
	from := sb[0].X-w
	to := sb[0].X+w
	out := make([]Point,0,to-from+1)
	if sb[1].Y == y {
		if from == sb[1].X {
			from++
		} else if to == sb[1].X {
			to--
		}
	}
	for x := from; x<=to; x++ {
		out = append(out, Point{x,y})
	}
	return out
}

func (sb *SensorBeacon) IsCover(y int) bool {
	return sb.Distance(y) <= sb.Radius()
}

func parseSensor(line string) *SensorBeacon {
	fields := strings.FieldsFunc(line, func(ch rune) bool {
		return strings.IndexRune(" ,:=", ch) != -1
	})
	sx, err := strconv.Atoi(fields[3])
	if err != nil {
		log.Fatalf("Invalid sensor posX: %s in line: %s: %v", fields[3], line, err)
	}
	sy, err := strconv.Atoi(fields[5])
	if err != nil {
		log.Fatalf("Invalid sensor posY: %s in line: %s: %v", fields[5], line, err)
	}
	bx, err := strconv.Atoi(fields[11])
	if err != nil {
		log.Fatalf("Invalid beacon posX: %s in line: %s: %v", fields[11], line, err)
	}
	by, err := strconv.Atoi(fields[13])
	if err != nil {
		log.Fatalf("Invalid beacon posY: %s in line: %s: %v", fields[13], line, err)
	}

	return &SensorBeacon{{sx,sy}, {bx, by}}
}

func Solution(input []string, y int) int {
	row := make(map[Point]rune)
	for _,line := range input {
		beacon := parseSensor(line)
		if beacon.IsCover(y) {
			for _,p := range beacon.CoverRow(y) {
				row[p] = '#'
			}
		}
	}
	return len(row)
}

func Task01(input []string) string {
	return fmt.Sprint(Solution(input, 2000000))
}

func Task02(input []string) string {
	return ""
}
