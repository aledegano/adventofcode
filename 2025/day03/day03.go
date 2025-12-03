package main

import (
	"flag"
	"fmt"
	"math"

	// "math"
	"os"
	// "regexp"
	"strconv"
	"strings"
	// "unicode"
)

var inputFile = flag.String("inputFile", "test_input.txt", "Relative file path to use as input.")
var debug = flag.Bool("debug", false, "Print debug info")

func mostSignificantDigit(bank []int, firstBattery int, lastBattery int) (digit int, position int) {
	if *debug {
		fmt.Println("Finding most significant digit between", firstBattery, "and", lastBattery)
	}
	for i := firstBattery; i < lastBattery; i++ {
		if bank[i] > digit {
			digit = bank[i]
			position = i
		}
	}
	return digit, position
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
	joltagePart1 := 0
	joltagePart2 := 0
	for _, line := range lines {
		if len(line) < 1 {
			break
		}
		// Each line goes in a slice of integers
		bank := []int{}
		for i := 0; i < len(line); i++ {
			bankString, _ := strconv.Atoi(string(line[i]))
			bank = append(bank, bankString)
		}
		// Chose the first bankString as the maximum value between the first element and the last-1 element
		battery := 0
		firstBattery := 0
		firstPosition := 0
		secondBattery := 0
		for j := 0; j < len(bank)-1; j++ {
			if bank[j] > firstBattery {
				firstBattery = bank[j]
				firstPosition = j
			}
		}
		for k := firstPosition + 1; k < len(bank); k++ {
			if bank[k] > secondBattery {
				secondBattery = bank[k]
			}
		}
		battery = firstBattery*10 + secondBattery
		joltagePart1 = joltagePart1 + battery
		// Part 2
		currentDigit := 0
		position := -1
		battery = 0
		for digit := 11; digit >= 0; digit-- {
			if *debug {
				fmt.Println("Digit:", digit, "Position:", position)
			}
			currentDigit, position = mostSignificantDigit(bank, position+1, len(bank)-digit)
			if *debug {
				fmt.Println("Current Digit:", currentDigit, "at position", position)
			}
			if *debug {
				fmt.Println("Multiplying", currentDigit, "*", int(math.Pow(10, float64(digit))))
			}
			battery = currentDigit*int(math.Pow(10, float64(digit))) + battery
			if *debug {
				fmt.Println("Battery now:", battery)
			}
		}
		joltagePart2 = joltagePart2 + battery
	}
	fmt.Println("Joltage Part1:", joltagePart1)
	fmt.Println("Joltage Part2:", joltagePart2)
}
