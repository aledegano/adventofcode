package main

import (
	"flag"
	"fmt"
	// "math"
	"os"
	// "regexp"
	"strconv"
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
	orderRules := make(map[int][]int)
	toggleUpdateParse := false
	pageUpdates := [][]int{}
	for _, line := range lines {
		if len(line) == 0 {
			toggleUpdateParse = true
			continue
		} 
		if !toggleUpdateParse {
			strings := strings.Split(line, "|")
			leftNum, _ := strconv.Atoi(strings[0])
			rightNum, _ := strconv.Atoi(strings[1])
			orderRules[leftNum] = append(orderRules[leftNum], rightNum)
		} else {
			pageUpdate := []int{}
			strings := strings.Split(line, ",")
			for _, pageStr := range strings {
				page, _ := strconv.Atoi(pageStr)
				pageUpdate = append(pageUpdate, page)
			}
			pageUpdates = append(pageUpdates, pageUpdate)
		}
	}
	fmt.Println(orderRules)
	fmt.Println(pageUpdates)
}
