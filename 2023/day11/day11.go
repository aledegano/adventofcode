package main

import (
	"flag"
	"fmt"
	"math"
	"os"
	"strings"
)

var inputFile = flag.String("inputFile", "test_input.txt", "Relative file path to use as input.")
var debug = flag.Bool("debug", false, "Print debug info")
var part = flag.Int("part", 1, "Which part of the puzzle to solve")

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
	initialPhoto := []Coord{}
	emptyRows := make(map[int]bool, len(lines[0]))
	for i := 0; i<len(lines[0]); i++ {
		emptyRows[i] = true
	}
	emptyCols := make(map[int]bool, len(lines)-1)
	for i := 0; i<len(lines)-1; i++ {
		emptyCols[i] = true
	}
	for i, line := range lines {
		if len(line) == 0 {
			continue
		}
		for j, char := range line {
			if char == '#' {
				initialPhoto = append(initialPhoto, Coord{i, j})
				emptyCols[j] = false
				emptyRows[i] = false
			}
		}
	}
	for i, coord := range initialPhoto {
		deltaX := 0
		deltaY := 0
		for row, empty := range emptyRows {
			if empty && coord.x > row {
				if *part == 1 {
					deltaX++
				} else {
					deltaX+=999999 // we need to replace 1 row with 1M rows, so we need to add 999999
				}
			}
		}
		for col, empty := range emptyCols {
			if empty && coord.y > col {
				if *part == 1 {
					deltaY++
				} else {
					deltaY+=999999
				}
			}
		}
		initialPhoto[i] = Coord{coord.x + deltaX, coord.y + deltaY}
	}
	// evaluate distances between all pair of galaxies
	result := 0
	for i, coord := range initialPhoto {
		for j:=i+1; j<len(initialPhoto); j++ {
			result += int(math.Abs(float64((initialPhoto[j].x - coord.x)))) + int(math.Abs(float64(coord.y - initialPhoto[j].y)))
		}
	}
	fmt.Printf("[Part %d] Sum of the shortest distance: %d\n", *part, result)
}
