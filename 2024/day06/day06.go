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
	guardPosition := Position{0, 0}

	uniqueVisitedPositions := make(map[Position]string)
	for y, line := range lines {
		if len(line) == 0 {
			continue
		}
		labMap[y] = make([]string, len(line))
		for x, char := range line {
			labMap[y][x] = string(char)
			if string(char) == "<" || string(char) == ">" || string(char) == "^" || string(char) == "v" {
				guardPosition.x = x
				guardPosition.y = y
			}
		}
	}
	newGuardPosition := guardPosition
	loops := 0
	for {
		direction := labMap[guardPosition.y][guardPosition.x]
		newGuardPosition = moveGuard(guardPosition, direction)
		// if the guard exited the map break the loop
		if outsideMap(newGuardPosition, labMap) {
			break
		}
		if labMap[newGuardPosition.y][newGuardPosition.x] == "#" { // Rotate 90 degrees to the right
			direction = rotateMove[labMap[guardPosition.y][guardPosition.x]]
			newGuardPosition = moveGuard(guardPosition, direction)
		}
		labMap[guardPosition.y][guardPosition.x] = "."
		labMap[newGuardPosition.y][newGuardPosition.x] = direction
		guardPosition = newGuardPosition
		uniqueVisitedPositions[newGuardPosition] = direction
		// part 2, check if a legal rotation would place the guard in a position already visited with the same direction it had the first time
		potentialLoopDirection := rotateMove[direction]
		potentialLoopPosition := moveGuard(guardPosition, potentialLoopDirection)
		for {
			if _, ok := uniqueVisitedPositions[potentialLoopPosition]; ok { // the potential position is a visited position
				if uniqueVisitedPositions[potentialLoopPosition] == potentialLoopDirection { // the direction is the same as the first time
					fmt.Printf("Found a loop! The position is %v and the direction is %s\n", potentialLoopPosition, potentialLoopDirection)
					loops++
					break
				}
			}
			// if the potential position is not a visited position, continue in the same direction and try again until a visited position is found or the guard exits the map
			potentialLoopPosition = moveGuard(potentialLoopPosition, potentialLoopDirection)
			if outsideMap(potentialLoopPosition, labMap) {
				break
			}
		}
	}
	fmt.Printf("Part 1. The number of unique visited positions is %d\n", len(uniqueVisitedPositions))
	fmt.Printf("Part 2. The number of loops is %d\n", loops)
}
