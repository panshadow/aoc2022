package day09

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	. "github.com/panshadow/aoc2022/utils"
)

func init() {
	RegisterTask("09/01", Task01, "09/01.twitter")
	RegisterTask("09/02", Task02, "09/01.twitter")
}

var (
	MOVES = map[string]struct{
		R int
		C int
	}{
		"U": {1,0},
		"R": {0,1},
		"D": {-1,0},
		"L": {0,-1},
	}
)



type KnotPos struct {
	S string
	R int
	C int
}

func (k *KnotPos) String() string {
	return fmt.Sprintf("%s {%-02d:%-02d}", k.S, k.R, k.C)
}

func NewKnot(s string, r,c int) *KnotPos {
	k := new(KnotPos)
	k.S = s
	k.R = r
	k.C = c
	return k
}

func (k *KnotPos) Distance(k2 *KnotPos) int {
	dr := Abs(k2.R - k.R)
	dc := Abs(k2.C - k.C)
	if dr > dc {
		return dr
	}
	return dc
}

func (k *KnotPos) Move(dr,dc int) {
	fmt.Print(k)
	k.R += dr
	k.C += dc
	fmt.Println(" -> ", k)
}

func (k *KnotPos) PullTower(k2 *KnotPos) {
	fmt.Printf("Pull %s Tower %s:\n", k, k2)
	for k.Distance(k2)>1 {
		dr := Ord(k2.R,k.R)
		dc := Ord(k2.C,k.C)
		k.Move(dr,dc)
	}
	fmt.Println("-")
}


func Task01(input []string) string {
	H := NewKnot("H",0,0)
	T := NewKnot("T",0,0)
	Path := make(map[KnotPos]bool)
	for i,line := range input {
		fields := strings.Fields(line)
		num,err := strconv.Atoi(fields[1])
		if err != nil {
			log.Fatalf("Invalid command on line %d: %s: %v\n", i, line, err)
		}

		for i:=0;i<num;i++ {
			d := MOVES[fields[0]]
			H.Move(d.R, d.C)
			T.PullTower(H)
			Path[*T] = true
		}
	}

	return fmt.Sprint(len(Path))
}

func Task02(input []string) string {
	N := 10
	Rope := make([]*KnotPos,N)
	for i := range Rope {
		Rope[i] = NewKnot(fmt.Sprintf("K%d",i), 0, 0)
	}
	Path := make(map[KnotPos]bool)
	for i,line := range input {
		fields := strings.Fields(line)
		num,err := strconv.Atoi(fields[1])
		if err != nil {
			log.Fatalf("Invalid command on line %d: %s: %v\n", i, line, err)
		}

		for i:=0;i<num;i++ {
			d := MOVES[fields[0]]
			Rope[0].Move(d.R, d.C)
			for j:=1;j<N;j++ {
				Rope[j].PullTower(Rope[j-1])
			}
			Path[*Rope[N-1]] = true
		}
	}

	return fmt.Sprint(len(Path))
}
