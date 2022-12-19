package day17

import (
	"fmt"

	. "github.com/panshadow/aoc2022/utils"
)

func init() {
	RegisterTask("17/01", Task01, "17/01.twitter")
	RegisterTask("17/02", Task02, "17/01.twitter")
}

type Point struct {
	R int
	C int
}

type Rock struct {
	Rocks []Point
}

type void struct{}

var (
	Void        = struct{}{}
	ShapeHor    = Rock{Rocks: []Point{{0, 0}, {0, 1}, {0, 2}, {0, 3}}}
	ShapeCross  = Rock{Rocks: []Point{{2, 1}, {1, 0}, {1, 1}, {1, 2}, {0, 1}}}
	ShapeAngle  = Rock{Rocks: []Point{{2, 2}, {1, 2}, {0, 2}, {0, 1}, {0, 0}}}
	ShapePipe   = Rock{Rocks: []Point{{3, 0}, {2, 0}, {1, 0}, {0, 0}}}
	ShapeSquare = Rock{Rocks: []Point{{1, 0}, {1, 1}, {0, 1}, {0, 0}}}
)

func NewRock(pattern []Point, pos Point) *Rock {
	r := Rock{}
	r.Rocks = make([]Point, len(pattern))
	copy(r.Rocks, pattern)
	for i := range r.Rocks {
		r.Rocks[i].C += pos.C
		r.Rocks[i].R += pos.R
	}

	return &r
}

func checkBit(rocks []uint, r, c int) bool {
	if r < len(rocks) {
		mask := uint(1) << c
		return rocks[r]&mask != 0
	}
	return false
}

func (r *Rock) Move(m []uint, width, dr, dc int) bool {
	next := make([]Point, len(r.Rocks))
	for i, p := range r.Rocks {
		np := Point{p.R + dr, p.C + dc}
		if np.C >= 0 && np.C < width && np.R >= 0 {

			if !checkBit(m, np.R, np.C) {
				next[i] = np
			} else {
				return false
			}
		} else {
			return false
		}
	}
	copy(r.Rocks, next)
	return true
}

func (tc *TallChamber) Put(r *Rock) int {
	maxH := r.Rocks[0].R + 1
	if len(tc.Rocks) < maxH {
		add := make([]uint, maxH-len(tc.Rocks))
		tc.Rocks = append(tc.Rocks, add...)
	}
	for _, p := range r.Rocks {
		Debugln("Put ", p.R, p.C, tc.Rocks[p.R])
		tc.Rocks[p.R] |= 1 << p.C
		Debugf("New %04X\n", tc.Rocks[p.R])
	}
	return maxH
}

type TallChamber struct {
	Rocks      []uint
	Height     int
	Width      int
	Counter    int
	JetPattern string
	JPCur      int
	Shapes     []Rock
	ShapeCur   int
}

func NewTallChamber(width int, jetPattern string, shapes []Rock) *TallChamber {
	tc := &TallChamber{
		Rocks:      make([]uint, 0, 1024*1024),
		Width:      width,
		JetPattern: jetPattern,
		Shapes:     shapes,
	}

	return tc
}

func (tc *TallChamber) Show() {
	r0 := 0
	Debugln("After ", tc.Counter, "Rocks, Height is ", tc.Height)
	for r := 12; r >= 0; r-- {
		row := make([]rune, tc.Width)
		for c := 0; c < tc.Width; c++ {
			exists := checkBit(tc.Rocks, r+r0, c)
			if exists {
				row[c] = '#'
			} else {
				row[c] = '.'
			}
		}
		var rr uint
		if r+r0 < len(tc.Rocks) {
			rr = tc.Rocks[r+r0]
		}

		Debugf("%4d |%s| %04X\n", r+r0, string(row), rr)
	}
	Debugln("----------")
}

func DetectCycles(xs []uint) {
	i0 := len(xs) - 1
	for i := len(xs) - 2; i >= 0; i-- {
		if xs[i] == xs[i0] {
			j := 0
			for ; xs[i-j] == xs[i0-j] && j < i0-i; j++ {

			}
			fmt.Printf("Found repeat in len %d [%d:%d] & [%d:%d]\n", j, i0, i0-j, i, i-j)
		}
	}
}

func (tc *TallChamber) RockRound() {
	r := NewRock(tc.Shapes[tc.ShapeCur].Rocks, Point{tc.Height + 3, 2})
	stopped := false
	for !stopped {
		Debugln(r.Rocks)
		move := tc.JetPattern[tc.JPCur]
		if move == '<' {
			Debugln(tc.Counter, "Left")
			r.Move(tc.Rocks, tc.Width, 0, -1)
		} else if move == '>' {
			Debugln(tc.Counter, "Right")
			r.Move(tc.Rocks, tc.Width, 0, 1)
		}
		Debugln(tc.Counter, "Down")
		stopped = !r.Move(tc.Rocks, tc.Width, -1, 0)
		if stopped {
			Debugln(tc.Counter, "Stopped")
		}
		tc.JPCur = (tc.JPCur + 1) % len(tc.JetPattern)
	}
	maxR := tc.Put(r)
	if maxR > tc.Height {
		tc.Height = maxR
	}
	if tc.JPCur == 0 {
		fmt.Println("Counter is ",tc.Counter)
		DetectCycles(tc.Rocks)
	}
	tc.Counter++
	tc.ShapeCur = (tc.ShapeCur + 1) % len(tc.Shapes)
}

func Task01(input []string) string {
	tc := NewTallChamber(7, input[0], []Rock{
		ShapeHor,
		ShapeCross,
		ShapeAngle,
		ShapePipe,
		ShapeSquare,
	})
	for i := 0; i < 2022; i++ {
		tc.RockRound()
		tc.Show()
	}
	return fmt.Sprint(tc.Height)
}

func Task02(input []string) string {
	tc := NewTallChamber(7, input[0], []Rock{
		ShapeHor,
		ShapeCross,
		ShapeAngle,
		ShapePipe,
		ShapeSquare,
	})
	// for i := 0; i < 1000000000000; i++ {
	for i := 0; i < 1000000000000; i++ {
		tc.RockRound()
	}
	return fmt.Sprint(tc.Height)
}
