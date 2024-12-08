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
	nodesPt2 := map[Position]bool{}
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
				// for part 2 whenever we have a couple of antennas their position is also an antinode
				// so we add it to the list of nodes
				nodesPt2[positions[i]] = true
				nodesPt2[positions[j]] = true
				if ! outsideMap(node1, mapSize) {
					nodes[node1] = true
					mapVis[node1.y][node1.x] = "#"
					// For part 2 keep adding the distance to the node found until we are outside the map
					resonantNode := node1
					for ! outsideMap(resonantNode, mapSize) {
						nodesPt2[resonantNode] = true
						mapVis[resonantNode.y][resonantNode.x] = "#"
						resonantNode = Position{resonantNode.x+distance.x, resonantNode.y+distance.y}
					}
				}
				node2 := Position{positions[j].x-distance.x, positions[j].y-distance.y}
				if ! outsideMap(node2, mapSize) {
					nodes[node2] = true
					nodesPt2[node2] = true
					mapVis[node2.y][node2.x] = "#"
					// For part 2 keep adding the distance to the node found until we are outside the map
					resonantNode := node2
					for ! outsideMap(resonantNode, mapSize) {
						nodesPt2[resonantNode] = true
						mapVis[resonantNode.y][resonantNode.x] = "#"
						resonantNode = Position{resonantNode.x-distance.x, resonantNode.y-distance.y}
					}
				}
			}
		}
	}
	printMap(mapVis)
	fmt.Printf("Part 1, found %d antinodes.\n", len(nodes))
	fmt.Printf("Part 2, found %d resonant antinodes.\n", len(nodesPt2))
}
