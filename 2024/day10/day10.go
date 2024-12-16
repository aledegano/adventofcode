package main

import (
	"flag"
	"fmt"
	// "math"
	"os"
	// "regexp"
	"strconv"
	"strings"
	// "time"
	// "unicode"
)

var inputFile = flag.String("inputFile", "test_input.txt", "Relative file path to use as input.")
var debug = flag.Bool("debug", false, "Print debug info")

type Position struct {
	x, y int
}

type Trail struct {
	start, end Position
}

var directions = map[int]Position{
	0: Position{0, 1},
	1: Position{1, 0},
	2: Position{0, -1},
	3: Position{-1, 0},
}

func nextStep(pos Position, topoMap map[Position]int) []Position {
	toVisit := []Position{}
	for i := 0; i < 4; i++ {
		newPos := Position{pos.x + directions[i].x, pos.y + directions[i].y}
		if height, ok := topoMap[newPos]; ok {
			if height == topoMap[pos]+1 {
				toVisit = append(toVisit, newPos)
			}
		}
	}
	if len(toVisit) == 0 {
		toVisit = append(toVisit, Position{-1, -1})
	}
	return toVisit
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
	trailHeads := []Position{}    // list of positions that are the start of a trail
	for y, line := range lines {
		if len(line) == 0 {
			continue
		}
		for x, c := range line {
			if c != '.' {
				height, _ := strconv.Atoi(string(c))
				topoMap[Position{x, y}] = height
				if height == 0 {
					trailHeads = append(trailHeads, Position{x, y})
				}
			}
		}
	}
	uniqueTrailScores := map[Trail]bool{}
	trailScores := map[Trail]int{}
	totalScore := 0
	posList := []Position{}
	for _, trailHead := range trailHeads {
		posList = []Position{}
		posList = append(posList, trailHead)
		for {
			// time.Sleep(1 * time.Second)
			if len(posList) == 0 {
				break // no more trails to follow for this trailhead
			}
			if topoMap[posList[0]] == 9 {
				trailScores[Trail{trailHead, posList[0]}]++ // reached the end of a trail and scored a point
				if _, ok := uniqueTrailScores[Trail{trailHead, posList[0]}]; !ok {
					uniqueTrailScores[Trail{trailHead, posList[0]}] = true
					posList = posList[1:]                       // continue with the next fork in the trail
					continue
				}
			}
			posList = append(posList, nextStep(posList[0], topoMap)...)
			posList = posList[1:] // remove the current position from the list
			if posList[0].x == -1 || posList[0].y == -1 {
				posList = posList[1:] // reached a dead end, continue with the next fork in the trail
				continue
			}
		}
	}
	fmt.Printf("Part1: UniqueTrailScores: %v\n", len(uniqueTrailScores))
	for _, score := range trailScores {
		totalScore += score
	}
	fmt.Printf("Part2: Total trailScores: %d\n", totalScore)
}
