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

	// part 1
	last := -1
	increments := 0
	for _, s := range split {
		i, _ := strconv.Atoi(s)
		if last != -1 && i > last {
			increments++
		}
		last = i
	}
	fmt.Println(increments)
}
