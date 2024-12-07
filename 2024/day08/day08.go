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
	antennas := make(map[string][][]int)
	mapSize := []int{len(lines[0]), len(lines)-1}
	for y, line := range lines {
		if line == "" {
			continue
		}
		idx := re.FindAllStringSubmatchIndex(line, -1)
		for _, match := range idx {
			frequency := string(line[match[0]])
			antennas[frequency] = append(antennas[frequency], []int{match[0], y})
		}
	}
	fmt.Println(antennas)
	nodes := [][]int{}
	for _, positions := range antennas {
		visited := map[int]bool{}
		for i := range positions {
			visited[i] = true
			for j:=i+1; j<len(positions); j++ {
				if visited[j] {
					continue
				}
				node1 := []int{positions[i][0]-positions[j][0], positions[i][1]-positions[j][1]}
				if node1[0] >= 0 && node1[1] >= 0 && node1[0] < mapSize[0] && node1[1] < mapSize[1] {
					nodes = append(nodes, node1)
				}
				node2 := []int{positions[j][0]-positions[i][0], positions[j][1]-positions[i][1]}
				if node2[0] >= 0 && node2[1] >= 0 && node2[0] < mapSize[0] && node2[1] < mapSize[1] {
					nodes = append(nodes, node2)
				}
				fmt.Println(node1, node2)
			}
		}
	}
	fmt.Println("Nodes: ", nodes)
}
