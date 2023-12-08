package main

import (
	"flag"
	"fmt"
	"os"
	"regexp"
	"strings"
	// "unicode"
)

var inputFile = flag.String("inputFile", "test_input.txt", "Relative file path to use as input.")
var debug = flag.Bool("debug", false, "Print debug info")

func main() {
	flag.Parse()
	bytes, err := os.ReadFile(*inputFile)
	if err != nil {
		return
	}
	contents := string(bytes)
	lines := strings.Split(contents, "\n")
	directions := []int{}
	dirMap := map[string]int{
		"L": 0,
		"R": 1,
	}
	for _, d := range lines[0] {
		directions = append(directions, dirMap[string(d)])
	}
	re := regexp.MustCompile(`[A-Z]{3}`)
	network := make(map[string]map[int]string, len(lines)-2)
	for _, node := range lines[1:] {
		if node == "" {
			continue
		}
		res := re.FindAllString(node, -1)
		leaves := map[int]string{
			0: res[1],
			1: res[2],
		}
		network[res[0]] = leaves
	}
	nextStep := "AAA"
	steps := 0
	for i:=0; true; i++ {
		if i == len(directions) {
			i = 0
		}
		direction := directions[i]
		nextStep = network[nextStep][direction]
		steps++
		if nextStep == "ZZZ" {
			break
		}
	}
	fmt.Printf("[Part 1] The number of steps to arrive at ZZZ is: %d\n", steps)
}
