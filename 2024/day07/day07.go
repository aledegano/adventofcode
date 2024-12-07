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
	"github.com/mowshon/iterium"
)

var inputFile = flag.String("inputFile", "test_input.txt", "Relative file path to use as input.")
var debug = flag.Bool("debug", false, "Print debug info")

func checkPermutations(result int, factors []int) bool {
	permutations, _ := iterium.Product([]bool{true, false}, len(factors)-1).Slice()
	// fmt.Printf("%d operators, %d permutations.\n", len(factors)-1, len(permutations))
	// fmt.Printf("%d: %v\n", result, factors)
	// fmt.Printf("Permutations: %v\n", permutations)
	for _, perm := range permutations {
		perResult := factors[0]
		for i, p := range perm {
			if p {
				perResult += factors[i+1]
			} else {
				perResult *= factors[i+1]
			}
		}
		if perResult == result {
			// fmt.Printf("Found a match: %v\n", perm)
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
	re := regexp.MustCompile(`(\d+)`)
	calibrations := [][]int{} // the first element is the left side of the equation, the following elements are the right side
	for line := range lines {
		if len(lines[line]) == 0 {
			continue
		}
		calibrations = append(calibrations, []int{})
		calibrationsString := re.FindAllString(lines[line], -1)
		for i := range calibrationsString {
			calibration, _ := strconv.Atoi(calibrationsString[i])
			calibrations[line] = append(calibrations[line], calibration)
		}
	}
	// part 1, see if the right side of the equation can be summed or multiplied to obtain the left side
	goodCalibrations := 0
	rejectedCalibrations := 0
	for _, calibration := range calibrations {
		// check the smallest number
		min := 0
		for i:=1; i<len(calibration); i++ {
			if calibration[i] == 1 { // if one factor is 1 then the minimum would be to multiply by it, hence we skip which is the same
				continue
			}
			min +=  calibration[i]
		}
		if min == calibration[0] {
			// fmt.Printf("Found a match: %v\n", calibration)
			goodCalibrations += calibration[0]
			continue
		}
		if min > calibration[0] {
			fmt.Printf("Reject %v with min: %d\n", calibration, min)
			rejectedCalibrations++
			continue
		}
		// check the largest number
		max := 1
		for i:=1; i<len(calibration); i++ {
			if calibration[i] == 1 { // if one factor is 1 then the maximum is obtained by adding it
				max +=  calibration[i]
				continue
			}
			max = max * calibration[i]
		}
		if max == calibration[0] {
			goodCalibrations += calibration[0]
			// fmt.Printf("Found a match: %v\n", calibration)
			continue
		}
		if max < calibration[0] {
			fmt.Printf("Reject %v with max: %d\n", calibration, max)
			rejectedCalibrations++
			continue
		}
		if checkPermutations(calibration[0], calibration[1:]){
			goodCalibrations += calibration[0]
			continue
			// fmt.Printf("Found a match: %v\n", calibration)
		}
		rejectedCalibrations++
		fmt.Printf("Reject %v with no permutation working, min:%d, max:%d.\n", calibration, min, max)
	}
	fmt.Printf("Good calibrations: %v\n", goodCalibrations)
	fmt.Printf("Rejected calibrations: %v\n", rejectedCalibrations)
}
