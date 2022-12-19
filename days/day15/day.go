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
			Debugf("Insert range [%d:%d] before #%d [%d:%d]\n",crange[0],crange[1],i,ranges[i][0], ranges[i][1])
			out = append(out,crange)
			for j:=i;j<len(ranges);j++ {
				out = append(out, ranges[j])
			}
			for _,r := range out {
				Debugf("[%d:%d] - ",r[0],r[1])
			}
			Debugln("*")
			return out
		case crange[0] <= ranges[i][0] && crange[1] >= ranges[i][1]:
			Debugf("Merge range [%d:%d] with #%d [%d:%d] outter\n",crange[0],crange[1],i,ranges[i][0], ranges[i][1])
			out = append(out, crange)
			for j:=i+1;j<len(ranges);j++ {
				out = AppendRange(out, ranges[j])
			}
			for _,r := range out {
				Debugf("[%d:%d] - ",r[0],r[1])
			}
			Debugln("*")
			return out
		case crange[0] >= ranges[i][0] && crange[1] <= ranges[i][1]:
			Debugf("Merge range [%d:%d] with #%d [%d:%d] inner\n",crange[0],crange[1],i,ranges[i][0], ranges[i][1])
			out = append(out, ranges[i])
			for j:=i+1;j<len(ranges);j++ {
				out = AppendRange(out, ranges[j])
			}
			for _,r := range out {
				Debugf("[%d:%d] - ",r[0],r[1])
			}
			Debugln("*")
			return out
		case crange[0] < ranges[i][0] && crange[1] >= ranges[i][0]-1 && crange[1] <= ranges[i][1]:
			Debugf("Merge range [%d:%d] with #%d [%d:%d] left\n",crange[0],crange[1],i,ranges[i][0], ranges[i][1])
			crange[1] = ranges[i][1]
			out = append(out, crange)
			for j:=i+1;j<len(ranges);j++ {
				out = AppendRange(out, ranges[j])
			}
			for _,r := range out {
				Debugf("[%d:%d] - ",r[0],r[1])
			}
			Debugln("*")
			return out
		case ranges[i][0] < crange[0] && ranges[i][1] >= crange[0]-1 && ranges[i][1] <= crange[1]:
			Debugf("Merge range [%d:%d] with #%d [%d:%d] right\n",crange[0],crange[1],i,ranges[i][0], ranges[i][1])
			crange[0] = ranges[i][0]
			out = append(out, crange)
			for j:=i+1;j<len(ranges);j++ {
				out = AppendRange(out, ranges[j])
			}
			for _,r := range out {
				Debugf("[%d:%d] - ",r[0],r[1])
			}
			Debugln("*")
			return out
		default:
			Debugf("Save range #%d [%d:%d]\n", i, ranges[i][0],ranges[i][1])
			out = append(out, ranges[i])
		}
	}
	Debugf("Append range [%d:%d]\n", crange[0], crange[1])

	out = append(out, crange)
	for _,r := range out {
		Debugf("[%d:%d] - ",r[0],r[1])
	}
	Debugln("*")
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
			Debugf("Y=%d, line=%d\n",y,i)
			if beacon.IsCover(y) {
				crange := beacon.CoverRange(y, false)
				Debugf("range: %d:%d\n",crange[0], crange[1])
				if crange[1] >= 0  || crange[0] <= size {
					crange[0] = Max(crange[0],0)
					crange[1] = Min(crange[1],size)
					Debugf("corrected range: %d:%d\n",crange[0], crange[1])
					ranges[y] = AppendRange(ranges[y], crange)
				}
			}
		}
	}

	Debugln("Start scanning distress beacon")


	for y:=0;y<=size;y++ {
		Debugf("y=%d ",y)
		for _,r := range ranges[y] {
			Debugf("[%d:%d] - ",r[0],r[1])
		}
		Debugln("*")
		var x int
		if len(ranges[y]) > 1 {
			for _,r := range ranges[y] {
				Debugf("[%d:%d] - ",r[0],r[1])
			}
			Debugln("*")
			for j := 0; j < len(ranges[y])-1; j++ {
				if ranges[y][j][1]+1 < ranges[y][j+1][0] {
					x = ranges[y][j][1]+1
					Debugln("Found x=",x," y=",y)
					fmt.Printf("%+v, y=%d, len(ranges[%d])=%d\n",ranges[y],y,y,len(ranges[y]))
					return x*size+y
				}
			}
		} else if len(ranges[y])==1 && len(ranges[y][0]) < size {
			Debugln("This one")
			if ranges[y][0][0] > 0 {
				x = ranges[y][0][0]
			} else if ranges[y][0][1] < size {
				x = ranges[y][0][1]
			}
			Debugln("Found x=",x," y=",y)
			return x*size+y
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
