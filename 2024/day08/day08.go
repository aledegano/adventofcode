package main

import (
	"flag"
	"fmt"
	// "math"
	"os"
	"regexp"
	// "strconv"
	"strings"
	// "unicode"
)

var inputFile = flag.String("inputFile", "test_input.txt", "Relative file path to use as input.")
var debug = flag.Bool("debug", false, "Print debug info")

type Position struct {
	x, y int
}

func outsideMap(position Position, mapSize Position) bool {
	return position.x < 0 || position.x >= mapSize.x || position.y < 0 || position.y >= mapSize.y
}

func printMap(m [][]string) {
	for _, row := range m {
		for _, cell := range row {
			fmt.Printf(cell)
		}
		fmt.Println()
	}
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
	re := regexp.MustCompile(`\w`)
	antennas := make(map[string][]Position)
	mapSize := Position{len(lines[0]), len(lines)-1}
	// for visualization purposes only
	mapVis := make([][]string, len(lines)-1)
	for y, line := range lines {
		if line == "" {
			continue
		}
		idx := re.FindAllStringSubmatchIndex(line, -1)
		for _, match := range idx {
			frequency := string(line[match[0]])
			antennas[frequency] = append(antennas[frequency], Position{match[0], y})
		}
		for _, c := range line {
			mapVis[y] = append(mapVis[y], string(c))
		}
	}
	// Part 1
	nodes := map[Position]bool{}
	for _, positions := range antennas {
		visited := map[int]bool{}
		for i := range positions {
			visited[i] = true
			for j:=i+1; j<len(positions); j++ {
				if visited[j] {
					continue
				}
				distance := Position{
					positions[i].x-positions[j].x,
					positions[i].y-positions[j].y,
				}
				node1 := Position{positions[i].x+distance.x, positions[i].y+distance.y}
				if ! outsideMap(node1, mapSize) {
					nodes[node1] = true
					mapVis[node1.y][node1.x] = "#"
				}
				node2 := Position{positions[j].x-distance.x, positions[j].y-distance.y}
				if ! outsideMap(node2, mapSize) {
					nodes[node2] = true
					mapVis[node2.y][node2.x] = "#"
				}
			}
		}
	}
	fmt.Printf("Part 1, found %d antinodes.\n", len(nodes))
}
