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
	fmt.Println(initialPhoto)
	fmt.Println(emptyRows)
	fmt.Println(emptyCols)
	for i, coord := range initialPhoto {
		for row, empty := range emptyRows {
			if empty && coord.x > row {
				coord.x++
			}
		}
		for col, empty := range emptyCols {
			if empty && coord.y > col {
				coord.y++
			}
		}
		initialPhoto[i] = coord
	}
	fmt.Println(initialPhoto)
}
