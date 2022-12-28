package day18

import (
	"fmt"

	. "github.com/panshadow/aoc2022/utils"
)

func init() {
	RegisterTask("18/01", Task01, "18/01.twitter")
	RegisterTask("18/02", Task02, "18/01.twitter")
}

type Dot struct {
	X int
	Y int
	Z int
}

var (
	Sides = []Dot{
		{0, 0, -1},
		{0, 0, 1},
		{0, -1, 0},
		{0, 1, 0},
		{-1, 0, 0},
		{1, 0, 0},
	}
	LAVA  = '#'
	WATER = '.'
	STEAM = 'o'
)

func (d Dot) Next(dd Dot) Dot {
	return Dot{d.X + dd.X, d.Y + dd.Y, d.Z + dd.Z}
}

func (d Dot) String() string {
	return fmt.Sprintf("%d,%d,%d", d.X, d.Y, d.Z)
}

func ParseDot(line string) Dot {
	coord := IntSlice(Tokens(line, ","))

	return Dot{coord[0], coord[1], coord[2]}
}

func Task01(input []string) string {
	world := make(map[Dot]rune)
	for _, line := range input {
		dot := ParseDot(line)
		world[dot] = LAVA
	}
	result := 0
	for cube := range world {
		surface := 6
		for _, side := range Sides {
			if _, exists := world[cube.Next(side)]; exists {
				surface--
			}
		}
		Debugf("Dot %s uncovered surface is %d\n", cube, surface)
		result += surface
	}
	return fmt.Sprint(result)
}

func FillWithWater(scene map[Dot]rune, min, max, cur Dot) int {
	if cur.X >= min.X && cur.Y >= min.Y && cur.Z >= min.Z && cur.X <= max.X && cur.Y <= max.Y && cur.Z <= max.Z {
		matter, exists := scene[cur]
		if !exists {
			scene[cur] = WATER
			result := 0
			for _, side := range Sides {
				result += FillWithWater(scene, min, max, cur.Next(side))
			}
			return result
		} else if matter == LAVA {
			return 1
		}
	}

	return 0
}

func Task02(input []string) string {
	world := make(map[Dot]rune)
	var min, max Dot
	for _, line := range input {
		dot := ParseDot(line)

		if dot.X < min.X {
			min.X = dot.X
		}
		if dot.X > max.X {
			max.X = dot.X
		}
		if dot.Y < min.Y {
			min.Y = dot.Y
		}
		if dot.Y > max.Y {
			max.Y = dot.Y
		}
		if dot.Z < min.Z {
			min.Z = dot.Z
		}
		if dot.Z > max.Z {
			max.Z = dot.Z
		}

		world[dot] = LAVA
	}
	min = min.Next(Dot{-1,-1,-1})
	max = max.Next(Dot{1,1,1})

	result := FillWithWater(world, min, max, min)
	return fmt.Sprint(result)
}
