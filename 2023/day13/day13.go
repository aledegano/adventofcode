package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
	// "unicode"
)

var inputFile = flag.String("inputFile", "test_input.txt", "Relative file path to use as input.")
var debug = flag.Bool("debug", false, "Print debug info")

type Coord struct {
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

	mirrors := []map[Coord]bool{}
	mapsBoundaries := []Coord{}
	mirrorMap := map[Coord]bool{}
	y := 0
	maxX := 0
	for _, line := range lines {
		for x, char := range line {
			if char == '#' {
				mirrorMap[Coord{x, y}] = true
			}
			maxX = len(line)
		}
		if line == "" {
			mirrors = append(mirrors, mirrorMap)
			mirrorMap = map[Coord]bool{}
			mapsBoundaries = append(mapsBoundaries, Coord{maxX, y})
			y = 0
			continue
		}
		y++
	}
	fmt.Println(mirrors)
	fmt.Println(mapsBoundaries)
	//part 1
	for i, mirrorMap := range mirrors {
		// first loop every row and check above and below if there is a match
		dance:
			for row := 1; row < mapsBoundaries[i].y; row++ {
				depth := maxY - row
				for x := 0; x < maxX; x++ {
					for y := 0; y < depth; y++ {
						if mirrorMap[Coord{x, row+y}] != mirrorMap[Coord{x, row-y}] {
							break dance
						}
					}
				}
			}
				fmt.Println("Found a match at row", row)
	}
}
