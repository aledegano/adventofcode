package main

import (
	"flag"
	"fmt"
	"os"
	// "regexp"
	"strings"
	// "unicode"
)

var inputFile = flag.String("inputFile", "test_input.txt", "Relative file path to use as input.")
var debug = flag.Bool("debug", false, "Print debug info")

type Pipe struct {
	connections [][]int
}

type Coordinates struct {
	x int
	y int
}

func main() {
	flag.Parse()
	bytes, err := os.ReadFile(*inputFile)
	if err != nil {
		return
	}
	contents := string(bytes)
	lines := strings.Split(contents, "\n")
	runeToCoord := map[rune][][]int{
		'|': [][]int{[]int{0,-1},[]int{0,1}},
		'-': [][]int{[]int{-1,0},[]int{1,0}},
		'L': [][]int{[]int{0,1},[]int{1,0}},
		'J': [][]int{[]int{0,1}, []int{-1,0}},
		'7': [][]int{[]int{0,-1}, []int{-1,0}},
		'F': [][]int{[]int{0,-1}, []int{1,0}},
		'S': [][]int{[]int{0,0}, []int{0,0}},
	}
	pipesField := map[Coordinates][]Coordinates{}
	start := Coordinates{0,0}
	// I want the coord to be 0,0 in the left-bottom corner, thus I read the lines from BOTTOM to TOP and from LEFT to RIGHT
	for i, line := range lines {
		if line == "" {
			continue
		}
		y := len(lines) -2 - i
		for x, r := range line {
			if r == 'S' {
				start = Coordinates{x,y}
			}
			delta := runeToCoord[rune(r)]
			if len(delta) == 0 {
				continue
			}
			connectionOne := Coordinates{x+delta[0][0],y+delta[0][1]}
			connectionTwo := Coordinates{x+delta[1][0],y+delta[1][1]}
			pipesField[Coordinates{x,y}] = []Coordinates{connectionOne, connectionTwo}
		}
	}
	currStep := start
	nextStep := Coordinates{start.x+1,start.y}
	steps := 0
	for {
		fmt.Printf("Current step: %v, next step: %v, pipe: %v", currStep, nextStep, pipesField[nextStep])
		if pipesField[nextStep][0].x == currStep.x && pipesField[nextStep][0].y == currStep.y {
			currStep = nextStep
			nextStep = Coordinates{pipesField[nextStep][1].x, pipesField[nextStep][1].y}
		} else {
			currStep = nextStep
			nextStep = Coordinates{pipesField[nextStep][0].x, pipesField[nextStep][0].y}
		}
		fmt.Println(nextStep)
		if nextStep.x == start.x && nextStep.y == start.y {
			break
		}
		steps++
		if steps > 100000 {
			fmt.Println("[ERROR] Too many steps")
			break
		}
	}
	fmt.Println("Total steps: ", steps)
	fmt.Println("Greatest distance: ", int(steps/2)+1)
}
