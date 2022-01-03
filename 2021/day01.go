package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

var inputFile = flag.String("inputFile", "inputs/day01.input", "Relative file path to use as input.")

func main() {
	flag.Parse()
	bytes, err := ioutil.ReadFile(*inputFile)
	if err != nil {
		return
	}
	contents := string(bytes)
	split := strings.Split(contents, "\n")
	depths := make([]int, len(split))
	for i, s := range split {
		depths[i], _ = strconv.Atoi(s)
	}

	// part 1
	last := -1
	increments := 0
	for _, s := range depths {
		if last != -1 && s > last {
			increments++
		}
		last = s
	}
	fmt.Printf("The number of increments is %d\n", increments)

	// part 2
	last = -1
	increments = 0
	for i := range depths {
		if i >= len(depths)-2 {
			break
		}
		window := depths[i] + depths[i+1] + depths[i+2]
		if last != -1 && window > last {
			increments++
		}
		last = window
	}
	fmt.Printf("The number of window increments is %d\n", increments)
}
