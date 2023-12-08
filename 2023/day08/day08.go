package main

import (
	"flag"
	"fmt"
	"os"
	"regexp"
	"slices"
	"strings"
)

var inputFile = flag.String("inputFile", "test_input.txt", "Relative file path to use as input.")
var debug = flag.Bool("debug", false, "Print debug info")

// greatest common divisor (GCD) via Euclidean algorithm
func GCD(a, b int) int {
	for b != 0 {
					t := b
					b = a % b
					a = t
	}
	return a
}

// find Least Common Multiple (LCM) via GCD
func LCM(a, b int, integers ...int) int {
	result := a * b / GCD(a, b)
	for i := 0; i < len(integers); i++ {
					result = LCM(result, integers[i])
	}
	return result
}

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
	// part 2
	startingNodes := []string{}
	endingNodes := []string{}
	startRegex := regexp.MustCompile(`[A-Z][A-Z][A]`)
	endRegex := regexp.MustCompile(`[A-Z][A-Z][Z]`)
	for k, _ := range network {
		if startRegex.MatchString(k) {
			startingNodes = append(startingNodes, k)
		}
		if endRegex.MatchString(k) {
			endingNodes = append(endingNodes, k)
		}
	}
	fmt.Printf("Starting nodes: %v\nEnding nodes: %v\n", startingNodes, endingNodes)
	part2Steps := []int{}
	for _, start := range startingNodes {
		nextStep := start
		steps := 0
		for i:=0; true; i++ {
			if i == len(directions) {
				i = 0
			}
			direction := directions[i]
			nextStep = network[nextStep][direction]
			steps++
			if slices.Contains(endingNodes, nextStep) {
				part2Steps = append(part2Steps, steps)
				break
			}
		}
	}
	fmt.Printf("Steps: %v\n", part2Steps)
	fmt.Printf("[Part 2] LCM: %d\n", LCM(part2Steps[0], part2Steps[1], part2Steps...))
}
