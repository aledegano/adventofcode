package main

import (
	"flag"
	"fmt"
	"math"
	"os"
	// "regexp"
	"strconv"
	"strings"
	// "unicode"
)

var inputFile = flag.String("inputFile", "test_input.txt", "Relative file path to use as input.")
var debug = flag.Bool("debug", false, "Print debug info")

func isOrderRespected(orderRules map[int][]int, currentPage int, nextPage int) bool {
	for _, next := range orderRules[currentPage] {
		if next == nextPage {
			return true
		}
	}
	return false
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
	middlePageSum := 0
	incorrectUpdates := [][]int{}
	for _, update := range pageUpdates {
		updateRespectsOrder := true
		out:
		for i, page := range update {
			for j:=i+1; j<len(update); j++ {
				if !isOrderRespected(orderRules, page, update[j]) {
					updateRespectsOrder = false
					incorrectUpdates = append(incorrectUpdates, update)
					break out
				}
			}
		}
		if updateRespectsOrder {
			middlePageSum += update[int(math.Floor(float64(len(update))/2))]
		}
	}
	fmt.Printf("Part 1: %d\n", middlePageSum)
	//Part2, fix the incorrect updates
	middlePageSum = 0
	for _, incorrectUpdate := range incorrectUpdates {
		fixThisUpdate:
		for i, page := range incorrectUpdate {
			for j:=i+1; j<len(incorrectUpdate); j++ {
				if !isOrderRespected(orderRules, page, incorrectUpdate[j]) {
					incorrectUpdate[i], incorrectUpdate[j] = incorrectUpdate[j], incorrectUpdate[i]
					goto fixThisUpdate
				}
			}
		}
		middlePageSum += incorrectUpdate[int(math.Floor(float64(len(incorrectUpdate))/2))]
	}
	fmt.Printf("Part 2: %d\n", middlePageSum)
}
