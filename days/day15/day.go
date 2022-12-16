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

func (sb *SensorBeacon) CoverRange(y int, includeBeacon bool) [2]int {
	w := sb.Radius() - sb.Distance(y)
	from := sb[0].X-w
	to := sb[0].X+w
	if includeBeacon && sb[1].Y == y {
		if from == sb[1].X {
			from++
		} else if to == sb[1].X {
			to--
		}
	}
	return [2]int{from, to}
}

func (sb *SensorBeacon) CoverRow(y int) []Point {
	crange := sb.CoverRange(y, true)
	from := crange[0]
	to := crange[1]
	out := make([]Point,0,to-from+1)
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

func (p Point) Addr(size int) int {
	return p.X*size + p.Y
}

func AppendRange(ranges [][2]int, crange [2]int) [][2]int {
	out := make([][2]int,0,len(ranges)+1)
	for i:=0; i<len(ranges); i++ {
		switch {
		case crange[1] < ranges[i][0]:
			out = append(out,crange)
			for j:=i;j<len(ranges);j++ {
				out = append(out, ranges[j])
			}
			return out
		case crange[0] <= ranges[i][0] && crange[1] >= ranges[i][0]:
			crange[1] = ranges[i][1]
			out = append(out, crange)
			for j:=i+1;j<len(ranges);j++ {
				out = append(out, ranges[j])
			}
			return out
		case crange[0] <= ranges[i][1] && crange[1] >= ranges[i][1]:
			crange[0] = ranges[i][0]
			out = append(out, crange)
			for j:=i+1;j<len(ranges);j++ {
				out = append(out, ranges[j])
			}
			return out
		default:
			out = append(out, ranges[i])
		}
	}
	out = append(out, crange)
	return out
}

func Solution2(input []string, size int) int {
	beacons := make([]*SensorBeacon, len(input))
	ranges := make([][][2]int,size+1)
	Debugln("Allocate memory")
	for i := range ranges {
		ranges[i] = make([][2]int,0, len(input))
	}
	// Min := Point{}
	// Max := Point{}
	Debugln("Collect ranges")
	for i,line := range input {
		Debugln(line)
		beacon := parseSensor(line)
		beacons[i] = beacon
		for y := 0; y<=size; y++ {
			if beacon.IsCover(y) {
				crange := beacon.CoverRange(y, false)
				if crange[1] >= 0  || crange[0] <= size {
					crange[0] = Max(crange[0],0)
					crange[1] = Min(crange[1],size)
					ranges[y] = append(ranges[y], crange)
				}
			}
		}
	}

	Debugln("Start scanning distress beacon")


	for y:=0;y<=size;y++ {
		// for i:=0; !found && i<len(ranges[y]); i++ {
		// 	if x >= ranges[y][i][0] && x <= ranges[y][i][1] {
		// 		found = true
		// 	}
		// }
		// if !found {
		// 	return x*size + y
		// }
		if len(ranges[y]) >= len(input) {
			return -2
		}
	}
	return -1
}

func Task01(input []string) string {
	return fmt.Sprint(Solution(input, 2000000))
}

func Task02(input []string) string {
	return fmt.Sprint(Solution2(input, 4000000))
}
