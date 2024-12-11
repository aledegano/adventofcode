package main

import (
	"flag"
	"fmt"
	// "math"
	"os"
	// "regexp"
	"strconv"
	"strings"
	// "unicode"
)

var inputFile = flag.String("inputFile", "test_input.txt", "Relative file path to use as input.")
var debug = flag.Bool("debug", false, "Print debug info")

type Position struct {
	x,y int
}

var directions = map[int]Position{
	0: Position{0,1},
	1: Position{1,0},
	2: Position{0,-1},
	3: Position{-1,0},
}

func nextStep(pos Position, topoMap map[Position]int) Position {
	for i:=0; i<4; i++ {
		newPos := Position{pos.x + directions[i].x, pos.y + directions[i].y}
		if height, ok := topoMap[newPos]; ok {
			if height == topoMap[pos]+1 {
				return newPos
			}
		}
	}
	return Position{-1,-1}
}

func main() {
	flag.Parse()
	bytes, err := os.ReadFile(*inputFile)
	if err != nil {
		fmt.Printf("Error reading file %s: %v\n", *inputFile, err)
		return
	}
	contents := string(bytes)
	lines := strings.Split(contents, "\n")
	topoMap := map[Position]int{} // maps a position to the height
	trailHeads := []Position{} // list of positions that are the start of a trail
	for y, line := range lines {
		if len(line) == 0 {
			continue
		}
		for x, c := range line {
			height, _ := strconv.Atoi(string(c))
			topoMap[Position{x,y}] = height
			if height == 0 {
				trailHeads = append(trailHeads, Position{x,y})
			}
		}
	}
	fmt.Printf("TopoMap: %v\n", topoMap)
	fmt.Printf("TrailHeads: %v\n", trailHeads)
	trailScores := 0
	for _, trailHead := range trailHeads {
		pos := trailHead
		for {
			nextPos := nextStep(pos, topoMap)
			if nextPos.x == -1 || nextPos.y == -1 {
				break
			}
			if topoMap[nextPos] == 9 {
				trailScores++
				break
			}
			pos = nextPos
		}
	}
	fmt.Printf("TrailScores: %d\n", trailScores)
}
