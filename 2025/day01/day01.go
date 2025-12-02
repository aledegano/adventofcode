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
	newPosition := 0
	password := 0
	across := 0
	for _, line := range lines {
		if len(line) < 1 {
			break
		}
		diff, _ := strconv.Atoi(line[1:])
		if string(line[0]) == "L" {
			newPosition = position - diff
		} else {
			newPosition = position + diff
		}
		if *debug {
			fmt.Println("Old position:", position, "New position:", newPosition)
		}
		if int(math.Abs(float64(newPosition))) > 100 {
			across = across + int(math.Floor(math.Abs(float64(newPosition))/100))
			if *debug {
				fmt.Println("* Multiple turns, across is now", across)
			}
			if (newPosition%100 < 0 && position > 0) || (newPosition%100 > 0 && position < 0) {
				across = across + 1
				if *debug {
					fmt.Println("* Crossed zero, across is now", across)
				}
			}
		position = newPosition
		position = position % 100
		if position == 0 {
			password = password + 1
			if *debug {
				fmt.Println("On zero", password)
			}
		}
	}
	fmt.Println("Part 1 solution:", password)
	fmt.Println("Part 2 solution:", password+across)
}
