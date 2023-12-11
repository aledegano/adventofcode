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
				deltaX++
			}
		}
		for col, empty := range emptyCols {
			if empty && coord.y > col {
				deltaY++
			}
		}
		initialPhoto[i] = Coord{coord.x + deltaX, coord.y + deltaY}
	}
	// evaluate distances the brute way
	part1Result := 0
	for i, coord := range initialPhoto {
		for j:=i+1; j<len(initialPhoto); j++ {
			part1Result += int(math.Abs(float64((initialPhoto[j].x - coord.x)))) + int(math.Abs(float64(coord.y - initialPhoto[j].y)))
		}
	}
	fmt.Printf("[Part 1] Sum of the shortest distance: %d\n", part1Result)
}
