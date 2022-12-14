package day14

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	. "github.com/panshadow/aoc2022/utils"
)

func init() {
	RegisterTask("14/01", Task01, "14/01.twitter")
	RegisterTask("14/02", Task02, "14/01.twitter")
}

type Point struct {
	R int
	C int
}

type Direction struct {
	Dr int
	Dc int
}

const (
	Rock = '#'
	Sand = 'o'
	Air  = '.'
	Edge = 'e'
)

var (
	Directions = []Direction{
		{0, 1},
		{-1, 1},
		{1, 1},
	}
)

func (p Point) Next(d Direction) Point {
	return Point{p.R + d.Dr, p.C + d.Dc}
}

func (p Point) String() string {
	return fmt.Sprintf("{%d:%d}", p.R, p.C)
}

func DrawWorld(world map[Point]rune, input []string) {
	Rmin := -1
	Rmax := -1
	Cmax := -1
	for i, line := range input {
		points := strings.Split(line, " -> ")
		ps := make([]Point, len(points))
		for j, sPoint := range points {
			nums := strings.Split(sPoint, ",")
			R, err := strconv.Atoi(nums[0])
			if err != nil {
				log.Fatalf("Can't parse number: %s at line %d:%d: %s: %v\n", nums[0], i, j, sPoint, err)
			}
			if Rmin == -1 || R < Rmin {
				Rmin = R
			}
			if Rmax == -1 || R > Rmax {
				Rmax = R
			}
			C, err := strconv.Atoi(nums[1])
			if err != nil {
				log.Fatalf("Can't parse number: %s at line %d:%d: %s: %v\n", nums[1], i, j, sPoint, err)
			}
			if Cmax == -1 || C > Cmax {
				Cmax = C
			}
			ps[j] = Point{R, C}
		}
		for j := 0; j < len(ps)-1; j++ {
			dr := Ord(ps[j+1].R, ps[j].R)
			dc := Ord(ps[j+1].C, ps[j].C)
			if dr != 0 {
				for r := ps[j].R; r != ps[j+1].R; r += dr {
					world[Point{r, ps[j].C}] = Rock
				}
			} else {
				for c := ps[j].C; c != ps[j+1].C; c += dc {
					world[Point{ps[j].R, c}] = Rock
				}
			}
		}
	}
	Debugf("Rmin=%d,Rmax=%d,Cmax=%d\n",Rmin,Rmax,Cmax)
	for j := Rmin; j<=Rmax; j++ {
		world[Point{j,Cmax+1}] = Edge
	}
	for j := 0; j < Cmax+1; j++ {
		world[Point{Rmin - 1, j}] = Edge
		world[Point{Rmax + 1, j}] = Edge
	}
}

func FallUnit(world map[Point]rune, source Point) bool {
	cur := source
	Debugf("Put sand to %s\n", source)
	for {
		moved := false
		for i := 0; !moved && i < len(Directions); i++ {
			next := cur.Next(Directions[i])
			Debugf("Try direction {%d:%d} -> %s", Directions[i].Dr, Directions[i].Dc, next)
			tile, exists := world[next]
			if !exists {
				Debugln("It's awailable. Move!")
				cur = next
				moved = true
			} else if tile == Edge {
				Debugln("It's edge. Moving endless!")
				return true
			} else {
				Debugf("It's impossible: %c\n", tile)
			}

		}
		if !moved {
			Debugf("Placed sand to %s\n\n", cur)
			world[cur] = Sand
			return false
		}
	}
}

func Task01(input []string) string {
	world := make(map[Point]rune)
	DrawWorld(world, input)
	counter := 0
	for !FallUnit(world, Point{500, 0}) {
		counter++
	}
	return fmt.Sprint(counter)
}

func Task02(input []string) string {
	return ""
}
