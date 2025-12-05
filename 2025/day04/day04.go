package main

import (
	"flag"
	"fmt"

	// "math"
	"os"
	// "regexp"
	// "strconv"
	"strings"
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
	floorGrid := [][]string{}
	newLine := make([]string, len(lines[0])+2)
	for i := 0; i < len(lines[0])+2; i++ {
		newLine[i] = "."
	}
	floorGrid = append(floorGrid, newLine)
	for _, line := range lines {
		if line == "" {
			break
		}
		newLine := make([]string, len(line)+2)
		newLine[0] = "."
		newLine[len(newLine)-1] = "."
		for i, char := range line {
			newLine[i+1] = string(char)
		}
		floorGrid = append(floorGrid, newLine)
	}
	newLine = make([]string, len(lines[0])+2)
	for i := 0; i < len(lines[0])+2; i++ {
		newLine[i] = "."
	}
	floorGrid = append(floorGrid, newLine)
	accessibleFloorGrid := [][]string{}
	for y := 0; y < len(floorGrid); y++ {
		newLine := make([]string, len(floorGrid[0]))
		for x := 0; x < len(floorGrid[0]); x++ {
			newLine[x] = string(floorGrid[y][x])
		}
		accessibleFloorGrid = append(accessibleFloorGrid, newLine)
	}

	accessibleRolls := 0
	for y := 1; y < len(floorGrid)-1; y++ {
		for x := 1; x < len(floorGrid[0])-1; x++ {
			adjacentRolls := 0
			if floorGrid[y][x] != "@" {
				continue
			}
			for adjY := -1; adjY <= 1; adjY++ {
				for adjX := -1; adjX <= 1; adjX++ {
					if adjY == 0 && adjX == 0 {
						continue
					}
					if floorGrid[y+adjY][x+adjX] == "@" {
						adjacentRolls++
					}
				}
			}
			if adjacentRolls < 4 {
				accessibleRolls++
				accessibleFloorGrid[y][x] = "x"
			}
		}
	}
	fmt.Printf("Accessible Rolls: %d\n", accessibleRolls)
	if *debug {
		fmt.Println("Accessible Floor Grid:")
		for _, line := range accessibleFloorGrid {
			fmt.Println(strings.Join(line, ""))
		}
	}
}
