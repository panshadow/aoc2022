package day07

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	. "github.com/panshadow/aoc2022/utils"
)

func init() {
	RegisterTask("07/01", Task01, "07/01.twitter")
	RegisterTask("07/02", Task02, "07/01.twitter")
}

type System struct {
	Path []string
	Dirs map[string]int
}

func NewSystem() *System {
	sys := new(System)
	sys.Path = []string{}
	sys.Dirs = make(map[string]int)
	return sys
}

func (sys *System) handleCD(dir string) {
	fmt.Printf("[/%s] # cd %s\n",strings.Join(sys.Path, ""), dir)
	switch dir {
	case "/":
		sys.Path = []string{}
	case "..":
		sys.Path = sys.Path[:len(sys.Path)-1]
	default:
		sys.Path = append(sys.Path, fmt.Sprintf("%s/",dir))
	}
	fmt.Printf("[/%s] #\n", strings.Join(sys.Path, ""))
}

func (sys *System) handleLS(line string) {
	path := fmt.Sprintf("/%s", strings.Join(sys.Path,""))
	fields := strings.Split(line, " ")
	if fields[0] == "dir" {
		dirpath := fmt.Sprintf("%s%s/", path, fields[1])
		if size, exists := sys.Dirs[dirpath]; !exists {
			fmt.Printf("Found dir %s %d\n", dirpath, size)
			sys.Dirs[dirpath] = size
		}
	} else {
		filepath := fmt.Sprintf("%s%s", path, fields[1])
		size,err := strconv.Atoi(fields[0])
		if err != nil {
			log.Fatalf("Invalid size (%s) of file %s in %s: %v ", fields[0], fields[1], path, err)
		}
		if _,exists := sys.Dirs[filepath]; !exists {
			fmt.Printf("Found file %s %d\n", filepath, size)
			sys.Dirs[filepath] = size
			curpath := "/"
			sys.Dirs[curpath] += size
			fmt.Printf("Update dir %s +%d = %d\n", curpath, size, sys.Dirs[curpath])
			for _,dir := range sys.Path {
				curpath = fmt.Sprintf("%s%s", curpath, dir)
				sys.Dirs[curpath] += size
				fmt.Printf("Update dir %s +%d = %d\n", curpath, size, sys.Dirs[curpath])
			}
		}
	}
}

func (sys *System) DirSize(path string) int {
	size, found := sys.Dirs[path]
	if found {
		return size
	}
	return -1
}

func (sys *System) Show() {
	for path,size := range sys.Dirs {
		fmt.Println(path, size)
	}
}

func (sys *System) runScript(input []string) {
	listing := false
	for _,line := range input {
		if strings.HasPrefix(line, "$ ") {
			listing = false
			fields := strings.Fields(line)
			if fields[1] == "cd" {
				sys.handleCD(fields[2])
			} else if fields[1] == "ls" {
				fmt.Printf("[/%s] # ls\n",strings.Join(sys.Path,""))
				listing = true
			}
		} else if listing {
			sys.handleLS(line)
		}
	}
}

func Task01(input []string) string {
	sys := NewSystem()
	sys.runScript(input)
	sys.Show()
	result := 0
	for path, size := range sys.Dirs {
		if strings.HasSuffix(path, "/") && size < 100000 {
			fmt.Println("FOUND ",path,size)
			result += size
		}
	}
	return fmt.Sprint(result)
}

func Task02(input []string) string {
	sys := NewSystem()
	sys.runScript(input)
	sys.Show()

	total := 70000000
	required := 30000000
	used := sys.DirSize("/")
	free := total - used
	need := required - free
	result := used
	for path, size := range sys.Dirs {
		if strings.HasSuffix(path, "/") {
			if size < result && size >=  need {
				result = size
				fmt.Println("FOUND ",path,size)
			}
		}
	}

	return fmt.Sprint(result)
}
