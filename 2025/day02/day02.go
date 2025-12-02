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
	idRanges := strings.Split(contents, ",")
	invalidIdSum := 0
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
				invalidIdSum += i
			}
		}
	}
	fmt.Println("Part 1:", invalidIdSum)
}
