package day12

import (
	"fmt"

	. "github.com/panshadow/aoc2022/utils"
)

func init() {
	RegisterTask("12/01", Task01, "12/01.twitter")
	RegisterTask("12/02", Task02, "12/01.twitter")
}

type Point struct {
	R int
	C int
}

func (p Point) String() string {
	return fmt.Sprintf("{%d:%d}", p.R, p.C)
}

func (p Point) Next(dr, dc int) Point {
	return Point{p.R + dr, p.C + dc}
}

func (p Point) Valid(w, h int) bool {
	return p.R >= 0 && p.R < w && p.C >= 0 && p.C < h
}

func (p Point) Elevation(m [][]rune) rune {
	return m[p.R][p.C]
}

type LinkedPoints struct {
	Val   Point
	Index int
	Link  *LinkedPoints
}

func (path *LinkedPoints) AddToPath(p Point) *LinkedPoints {
	node := new(LinkedPoints)
	node.Val = p
	node.Link = path
	if path != nil {
		node.Index = path.Index + 1
	} else {
		node.Index = 0
	}

	return node
}

func (path *LinkedPoints) Exists(p Point) bool {
	cur := path

	for cur != nil {
		if cur.Val == p {
			return true
		}
		cur = cur.Link
	}
	return false
}

type HeightMap struct {
	Map   [][]rune
	Start Point
	Starts []Point
	Goal  Point
}

type void struct{}

var (
	Void   = struct{}{}
	Orders = map[string]struct {
		Dr int
		Dc int
	}{
		"Up":    {-1, 0},
		"Right": {0, 1},
		"Down":  {1, 0},
		"Left":  {0, -1},
	}
)

func NewHeightMap(input []string) *HeightMap {
	hm := new(HeightMap)

	hm.Map = make([][]rune, len(input))
	hm.Starts = make([]Point,0,len(input)*len(input[0]))
	for r, line := range input {
		hm.Map[r] = make([]rune, len(line))
		for c, ch := range line {
			switch ch {
			case 'a':
				hm.Map[r][c] = 'a'
				hm.Starts = append(hm.Starts, Point{r, c})
			case 'S':
				hm.Map[r][c] = 'a'
				hm.Start = Point{r, c}
				hm.Starts = append(hm.Starts, hm.Start)
			case 'E':
				hm.Map[r][c] = 'z'
				hm.Goal = Point{r, c}
			default:
				hm.Map[r][c] = ch
			}
		}
	}
	return hm
}

func FindPath(m [][]rune, start Point, goal Point) int {
	mapw := len(m)
	maph := len(m[0])

	var root *LinkedPoints

	paths := []*LinkedPoints{root.AddToPath(start)}
	iteration := 0
	visited := make(map[Point]void)
	for {
		if len(paths) == 0 {
			return 0
		}
		if iteration%10 == 0 {
			fmt.Println("Iteration:",iteration, "Paths: ",len(paths), "Index:", paths[0].Index)

		}
		iteration++
		nextPaths := []*LinkedPoints{}
		for i, path := range paths {
			Debugf("%d) Find next square for %s #%d\n", i, path.Val, path.Index)
			elevation := path.Val.Elevation(m)
			ways := make([]Point, 0, 4)
			for ordKey, ord := range Orders {
				Debugf("Try next square %s\n", ordKey)
				way := path.Val.Next(ord.Dr, ord.Dc)
				if way.Valid(mapw, maph) {
					if _,wasHere := visited[way]; !wasHere {
						if way.Elevation(m) <= elevation+1 {
							Debugf("Square at %s: %s\n", ordKey, way)
							visited[way] = Void
							ways = append(ways, way)
						} else {
							Debugf("Impossible elevation %d %s for %d %s\n", way.Elevation(m),way, elevation, path.Val)
						}
					} else {
						Debugf("Point %s already in path\n", way)
					}
				} else {
					Debugf("Square %s isn't valid\n",way)
				}
			}
			for _, way := range ways {
				if way == goal {
					return path.Index + 1
				}
				nextPaths = append(nextPaths, path.AddToPath(way))
			}
		}
		paths = nextPaths
	}
}


func Solution(input []string, single bool) int {
	hm := NewHeightMap(input)
	pathLength := -1
	if single {
		pathLength = FindPath(hm.Map, hm.Start, hm.Goal)
	} else {
		for i, start := range hm.Starts {
			fmt.Printf("Finding path for %d/%d %s\n", i,len(hm.Starts), start)
			pl := FindPath(hm.Map, start, hm.Goal)
			if pathLength==-1 || (pl != 0 && pl < pathLength)  {
				pathLength = pl
			}
		}
	}
	return pathLength
}

func Task01(input []string) string {
	return fmt.Sprint(Solution(input, true))
}

func Task02(input []string) string {
	return fmt.Sprint(Solution(input, false))
}
