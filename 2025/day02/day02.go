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

func divisors(a int) []int {
	divs := []int{}
	for i := 2; i <= a; i++ { // for this problem we ignore divisibility by 1
		if a%i == 0 {
			divs = append(divs, i)
		}
	}
	return divs
}

func main() {
	flag.Parse()
	bytes, err := os.ReadFile(*inputFile)
	if err != nil {
		fmt.Printf("Error reading file %s: %v\n", *inputFile, err)
		return
	}
	contents := string(bytes)
	idRanges := strings.Split(contents, ",")
	invalidIdSumPt1 := 0
	invalidIdSumPt2 := 0
	for _, idRange := range idRanges {
		startEnd := strings.Split(idRange, "-")
		start, _ := strconv.Atoi(startEnd[0])
		end, _ := strconv.Atoi(startEnd[1])
		for i := start; i <= end; i+=1 {
			id := strconv.Itoa(i)
			if (len(id) % 2) != 0 {
				continue
			}
			if id[:len(id)/2] == id[len(id)/2:] {
				invalidIdSumPt1 += i
			}
		}
		// Part2
		for i := start; i <= end; i+=1 {
			id := strconv.Itoa(i)
			// First determine in how many groups the id strings can be divided, i.e. the MCP of the length of the slice
			divisors := divisors(len(id))
			for _, d := range divisors {
				groupSize := len(id) / d
				groups := []string{}
				for g := 0; g < d; g++ {
					groups = append(groups, id[g*groupSize:(g+1)*groupSize])
				}
				allEqual := true
				for g := 1; g < len(groups); g++ {
					if groups[g] != groups[0] {
						allEqual = false
						break
					}
				}
				if allEqual && groupSize < len(id) {
					invalidIdSumPt2 += i
					break
				}
			}
		}
	}
	fmt.Println("Part 1:", invalidIdSumPt1)
	fmt.Println("Part 2:", invalidIdSumPt2)
}
