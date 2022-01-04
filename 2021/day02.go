package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

var inputFile = flag.String("inputFile", "inputs/day02.input", "Relative file path to use as input.")

func main() {
	flag.Parse()
	bytes, err := ioutil.ReadFile(*inputFile)
	if err != nil {
		fmt.Println("Input file invalid:", err)
		return
	}
	contents := string(bytes)
	split := strings.Split(contents, "\n")

	// part 1
	x, y := 0, 0
	for _, s := range split {
		direction := strings.Split(s, " ")
		switch direction[0] {
		case "forward":
			deltax, _ := strconv.Atoi(direction[1])
			x += deltax
		case "up":
			deltay, _ := strconv.Atoi(direction[1])
			y -= deltay
		case "down":
			deltay, _ := strconv.Atoi(direction[1])
			y += deltay
		}
	}
	fmt.Printf("[PART 1] The product of the final coordinates is: %d\n", x*y)

	// part 2
	x, y = 0, 0
	aim := 0
	for _, s := range split {
		direction := strings.Split(s, " ")
		switch direction[0] {
		case "forward":
			deltax, _ := strconv.Atoi(direction[1])
			x += deltax
			y += deltax * aim
		case "up":
			deltay, _ := strconv.Atoi(direction[1])
			aim -= deltay
		case "down":
			deltay, _ := strconv.Atoi(direction[1])
			aim += deltay
		}
	}
	fmt.Printf("[PART 2] The product of the final coordinates is: %d\n", x*y)
}
