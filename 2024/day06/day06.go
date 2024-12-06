package main

import (
	"flag"
	"fmt"
	// "math"
	"os"
	// "regexp"
	// "strconv"
	"strings"
	"time"
	// "unicode"
)

var inputFile = flag.String("inputFile", "test_input.txt", "Relative file path to use as input.")
var debug = flag.Bool("debug", false, "Print debug info")

func main() {
	flag.Parse()
	bytes, err := os.ReadFile(*inputFile)
	if err != nil {
		fmt.Printf("Error reading file %s: %v\n", *inputFile, err)
		return
	}
	contents := string(bytes)
	lines := strings.Split(contents, "\n")
	labMap := make([][]string, len(lines))
	type Position struct {
		x, y int
	}
	guardPosition := Position{0, 0}
	move := map[string]Position{
		"<": Position{0, -1},
		">": Position{0, 1},
		"^": Position{-1, 0},
		"v": Position{1, 0},
	}
	rotateMove := map[string]string{
		"<": "^",
		"^": ">",
		">": "v",
		"v": "<",
	}
	uniqueVisitedPositions := make(map[Position]bool)
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
	for {
		newGuardPosition = Position{
			guardPosition.x + move[labMap[guardPosition.y][guardPosition.x]].x,
			guardPosition.y + move[labMap[guardPosition.y][guardPosition.x]].y,
		}
		fmt.Printf("New Guard position, and direction: %d, %d, %s\n", newGuardPosition.x, newGuardPosition.y, labMap[newGuardPosition.y][newGuardPosition.x])
		time.Sleep(1 * time.Second)
		if labMap[newGuardPosition.y][newGuardPosition.x] == "#" { // Rotate 90 degrees to the right
			newGuardPosition = Position{
				guardPosition.x + move[rotateMove[labMap[guardPosition.y][guardPosition.x]]].x,
				guardPosition.y + move[rotateMove[labMap[guardPosition.y][guardPosition.x]]].y,
			}
			fmt.Printf("New Guard position, and direction: %d, %d, %s\n", newGuardPosition.x, newGuardPosition.y, labMap[newGuardPosition.y][newGuardPosition.x])
			time.Sleep(1 * time.Second)
		}
		// if the guard exited the map break the loop
		if newGuardPosition.x < 0 || newGuardPosition.x >= len(labMap[0]) || newGuardPosition.y < 0 || newGuardPosition.y >= len(labMap) {
			break
		}
		guardPosition = newGuardPosition
		fmt.Printf("Guard position and direction: %d, %d, %s\n", guardPosition.x, guardPosition.y, labMap[guardPosition.y][guardPosition.x])
		time.Sleep(1 * time.Second)
		uniqueVisitedPositions[newGuardPosition] = true
	}
	fmt.Printf("The number of unique visited positions is %d\n", len(uniqueVisitedPositions))
}
