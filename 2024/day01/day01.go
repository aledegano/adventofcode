package main

import (
	"flag"
	"fmt"
	"math"
	"os"
	"regexp"
	"strings"
	"strconv"
	"sort"
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
	leftList := make([]int, 0)
	rightList := make([]int, 0)
	rightHistogram := make(map[int]int)
	re := regexp.MustCompile(`\d+`)
	for _, line := range lines {
		if len(line) < 1 {
			break
		}
		numbers := re.FindAllString(line, -1)
		left, _ := strconv.Atoi(numbers[0])
		right, _ := strconv.Atoi(numbers[1])
		leftList = append(leftList, left)
		rightList = append(rightList, right)
		rightHistogram[right] = rightHistogram[right] + 1
		sort.Ints(leftList)
		sort.Ints(rightList)
	}
	// part 1
	deltaSum := 0 
	for i := 0; i < len(leftList); i++ {
		deltaSum = deltaSum + int(math.Abs(float64(rightList[i]) - float64(leftList[i])))
	}
	fmt.Printf("The sum of the deltas is %d\n", deltaSum)
	// part 2
	similarityScore := 0
	for _, left := range leftList {
		similarityScore = similarityScore + left * rightHistogram[left]
	}
	fmt.Printf("The similarity score is %d\n", similarityScore)
}
