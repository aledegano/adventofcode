package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
)

var inputFile = flag.String("inputFile", "test_input.txt", "Relative file path to use as input.")
var debug = flag.Bool("debug", false, "Print debug info")

type Position struct {
	x, y int
}

var move = map[string]Position{
	"<": Position{-1, 0},
	">": Position{1, 0},
	"^": Position{0, -1},
	"v": Position{0, 1},
}

var rotateMove = map[string]string{
	"<": "^",
	"^": ">",
	">": "v",
	"v": "<",
}

func moveGuard(currentPosition Position, direction string) Position {
	return Position{
		currentPosition.x + move[direction].x,
		currentPosition.y + move[direction].y,
	}
}

func outsideMap(position Position, labMap [][]string) bool {
	return position.x < 0 || position.x >= len(labMap[0]) || position.y < 0 || position.y >= len(labMap)
}

func walkTheGuard(initialPosition Position, initialDirection string, labMap [][]string) (map[Position]string, bool) {
	path := make(map[Position]string)
	currentPosition := initialPosition
	currentDirection := initialDirection
	for {
		nextPosition := moveGuard(currentPosition, currentDirection)
		// if the guard exited the map break the loop
		if outsideMap(nextPosition, labMap) {
			path[currentPosition] = currentDirection
			return path, false
		}
		// if there is an obstacle in the next position, rotate 90 degrees to the right
		if labMap[nextPosition.y][nextPosition.x] == "#" { // Rotate 90 degrees to the right
			currentDirection = rotateMove[currentDirection]
			nextPosition = moveGuard(currentPosition, currentDirection)
		}
		path[currentPosition] = currentDirection
		// fmt.Printf("Position: %v, Direction: %s\n", currentPosition, currentDirection)
		// if the guard is in a loop, return.
		// the guard is in a loop if the next position is already visited and the direction is the same as the first time
		if _, ok := path[nextPosition]; ok {
			if path[nextPosition] == currentDirection {
				return path, true
			}
		}
		currentPosition = nextPosition
	}
}

func main() {
	flag.Parse()
	bytes, err := os.ReadFile(*inputFile)
	if err != nil {
		//fmt.Printf("Error reading file %s: %v\n", *inputFile, err)
		return
	}
	contents := string(bytes)
	lines := strings.Split(contents, "\n")
	labMap := make([][]string, len(lines)-1)
	initialPosition := Position{0, 0}
	initialDirection := ""

	for y, line := range lines {
		if len(line) == 0 {
			continue
		}
		labMap[y] = make([]string, len(line))
		for x, char := range line {
			labMap[y][x] = string(char)
			if string(char) == "<" || string(char) == ">" || string(char) == "^" || string(char) == "v" {
				initialPosition = Position{x, y}
				initialDirection = string(char)
			}
		}
	}
	// part 1
	// walk the guard and save the unique visited positions
	path, loop := walkTheGuard(initialPosition, initialDirection, labMap)
	fmt.Printf("Part 1. The number of unique visited positions is %d, found a loop? %t\n", len(path), loop)

	// part 2, check if a legal rotation would place the guard in a position already visited with the same initialDirection it had the first time
	obstructionPositions := make(map[Position]bool)
	for currentPosition, currentDirection := range path {
		obstructedLabMap := make([][]string, len(labMap))
		copy(obstructedLabMap, labMap)
		obstruction := moveGuard(currentPosition,currentDirection) // put an obstruction where the guard would be next
		if outsideMap(obstruction, labMap) {
			continue
		}
		if _, ok := obstructionPositions[obstruction]; ok {
			continue
		}
		if obstruction == initialPosition { // as per instructions the obstruction cannot be placed at the guard initial position
			continue
		}
		obstructedLabMap[obstruction.y][obstruction.x] = "#"
		// walk the guard and save the unique visited positions
		_, loop := walkTheGuard(currentPosition, currentDirection, obstructedLabMap)
		if loop {
			obstructionPositions[obstruction] = true
		}
		// remove the obstruction because it is faster to type than to deep-copy every slice of slices of labMap
		obstructedLabMap[obstruction.y][obstruction.x] = "."
	}
	fmt.Printf("Part 2. The number of obstructions creating a loop is %d\n", len(obstructionPositions))
}
