package main

import (
	"flag"
	"fmt"
	// "math"
	"os"
	"regexp"
	"strconv"
	"strings"
	// "unicode"
)

var inputFile = flag.String("inputFile", "test_input.txt", "Relative file path to use as input.")
var debug = flag.Bool("debug", false, "Print debug info")

func multiply(instructions []string) int {
	result := 0
	reFactors := regexp.MustCompile(`(\d+)`)
	for _, instruction := range instructions {
		factors := reFactors.FindAllString(instruction, -1)
		firstFactor, _ := strconv.Atoi(factors[0])
		secondFactor, _ := strconv.Atoi(factors[1])
		result += firstFactor * secondFactor
	}
	return result
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
	reInstruction := regexp.MustCompile(`mul\((\d+),(\d+)\)`)
	reInstructionStatement := regexp.MustCompile(`do\(\)|don't\(\)|mul\((\d+),(\d+)\)`)
	result := 0
	instructionStatements := []string{}
	for _, line := range lines {
		instructions := reInstruction.FindAllString(line, -1)
		instructionStatements = append(instructionStatements, reInstructionStatement.FindAllString(line, -1)...)
		result += multiply(instructions)
	}
	fmt.Printf("Part 1 result: %d\n", result)
	// part 2
	remove := false
	keptInstructions := []string{}
	for _, instruction := range instructionStatements {
		if (instruction == "don't()") {
			if !remove {
				remove = true
			}
			continue
		}
		if (instruction == "do()") {
			if remove {
				remove = false
			}
			continue
		}
		if remove {
			continue
		}
		// if we are here we want to keep this instruction
		keptInstructions = append(keptInstructions, instruction)
	}
	result = multiply(keptInstructions)
	fmt.Printf("Part 2 result: %d\n", result)
}
