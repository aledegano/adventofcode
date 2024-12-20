package main

import (
	"flag"
	"fmt"
	"os"
	// "regexp"
	"sort"
	"strings"
	"strconv"
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
	var calories []int
	currentCalories := 0
	for _, line := range lines {
		if line == "" {
			calories = append(calories, currentCalories)
			currentCalories = 0
			continue
		}
		cal, _ := strconv.Atoi(line)
		currentCalories = currentCalories + cal
	}
	// Sort in descending order
	sort.Sort(sort.Reverse(sort.IntSlice(calories)))

	fmt.Println("Part 1: ", calories[0])
	fmt.Println("Part 2: ", calories[0]+calories[1]+calories[2])
}
