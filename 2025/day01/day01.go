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
	position := 50
	for _, line := range lines {
		if len(line) < 1 {
			break
		}
		diff, _ := strconv.Atoi(line[1:])
		fmt.Println(diff)
		if string(line[0]) == "L" {
			position = position - diff
		} else {
			position = position + diff
		}
		if position < 0 {
			position = 100 - position
		}
		if position > 99 {
			position = position - 100
		}
		fmt.Println(position)
	}
}
