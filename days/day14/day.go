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
	SAFE_DEEP = 10
)

var (
	Directions = []Direction{
		{0, 1},
		{-1, 1},
		{1, 1},
	}
)

type Board struct {
	World map[Point]rune
	Min Point
	Max Point
}

func (p Point) Next(d Direction) Point {
	return Point{p.R + d.Dr, p.C + d.Dc}
}

func (p Point) String() string {
	return fmt.Sprintf("{%d:%d}", p.R, p.C)
}

func NewBoard(input []string) *Board {
	board := new(Board)
	board.World = make(map[Point]rune)
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
				for r := ps[j].R; r != ps[j+1].R+dr; r += dr {
					board.World[Point{r, ps[j].C}] = Rock
				}
			} else {
				for c := ps[j].C; c != ps[j+1].C+dc; c += dc {
					board.World[Point{ps[j].R, c}] = Rock
				}
			}
		}
	}
	board.Min = Point{Rmin,0}
	board.Max = Point{Rmax,Cmax}
	Debugf("Rmin=%d,Rmax=%d,Cmax=%d\n",Rmin,Rmax,Cmax)

	return board
}

func (b *Board) WrapEdge() {
	b.Min.R--
	b.Max.R++
	b.Max.C++
	for r := b.Min.R; r<= b.Max.R; r++ {
		b.World[Point{r,b.Max.C}] = Edge
	}
	for c := 0; c < b.Max.C; c++ {
		b.World[Point{b.Min.R, c}] = Edge
		b.World[Point{b.Max.R, c}] = Edge
	}
}

func (b *Board) AddBottomFloor(offset int) {
	b.Max.C += offset-1
}

func (b *Board) FallUnit(source Point) bool {
	cur := source
	if tile,found := b.World[cur]; found {
		Debugf("Can't put sand unit. Blocked by %c at %s\n",tile,source)
		return false
	}
	Debugf("Put sand to %s\n", source)
	for {
		if cur.C >= b.Max.C+SAFE_DEEP {
			Debugf("It seems like overflow: %d over %d (+%d)\n", cur.C, b.Max.C+SAFE_DEEP,SAFE_DEEP)
			return false
		}
		moved := false
		for i := 0; !moved && i < len(Directions); i++ {
			next := cur.Next(Directions[i])
			Debugf("Try direction {%d:%d} -> %s", Directions[i].Dr, Directions[i].Dc, next)
			tile, exists := b.World[next]
			if !exists && next.C <= b.Max.C {
				Debugln("It's awailable. Move!")
				cur = next
				moved = true
			} else if tile == Edge {
				Debugln("It's edge. Moving endless!")
				return false
			} else {
				Debugf("It's impossible: %c\n", tile)
			}

		}
		if !moved {
			Debugf("Placed sand to %s\n\n", cur)
			b.World[cur] = Sand
			if cur.R < b.Min.R {
				b.Min.R = cur.R
			}
			if cur.R > b.Max.R {
				b.Max.R = cur.R
			}
			return true
		}
	}
}

func (b *Board) Show() string {
	bw := b.Max.R-b.Min.R+1
	out := make([]string, b.Max.C+1)
	out[0] = fmt.Sprintf("%d - %d", b.Min.R, b.Max.R)
	for c := 0; c<=b.Max.C; c++ {
		row := make([]rune, bw)
		for r := 0; r < bw; r++ {
			tile, found := b.World[Point{r+b.Min.R,c}]
			if found {
				row[r] = tile
			} else {
				row[r] = '.'
			}
		}
		out[c] = fmt.Sprintf("%02d %s",c, string(row))
	}
	return strings.Join(out, "\n")
}

func Task01(input []string) string {
	board := NewBoard(input)
	board.WrapEdge()
	counter := 0
	for board.FallUnit(Point{500, 0}) {
		Debugln(board.Show())
		counter++
	}
	return fmt.Sprint(counter)
}

func Task02(input []string) string {
	board := NewBoard(input)
	board.AddBottomFloor(2)
	counter := 0
	for board.FallUnit(Point{500, 0}) {
		Debugln(board.Show())
		counter++
	}
	return fmt.Sprint(counter)
}
