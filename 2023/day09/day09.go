package main

import (
	"flag"
	"fmt"
	// "math"
	"os"
	"regexp"
	"strings"
	"strconv"
	// "unicode"
)

var inputFile = flag.String("inputFile", "test_input.txt", "Relative file path to use as input.")
var debug = flag.Bool("debug", false, "Print debug info")

func main() {
	flag.Parse()
	bytes, err := os.ReadFile(*inputFile)
	if err != nil {
		return
	}
	contents := string(bytes)
	lines := strings.Split(contents, "\n")
	// part 1
	series := [][]int{}
	re := regexp.MustCompile(`-?\d+`)
	for _, line := range lines {
		if line == "" {
			continue
		}
		serieStr := re.FindAllString(line,-1)
		serie := []int{}
		for _, s := range serieStr {
			i, _ := strconv.Atoi(s)
			serie = append(serie, i)
		}
		series = append(series, serie)
	}
	nextValues := 0
	prevValues := 0
	for _, serie := range series {
		thisDeriv := serie
		firstValues := []int{}
		dance:
			for {
				allZero := true
				firstValues = append(firstValues, thisDeriv[0])
				nextDeriv := make([]int, len(thisDeriv)-1)
				for i := 0; i < len(thisDeriv)-1; i++ {
					delta := (thisDeriv[i+1]-thisDeriv[i])
					if delta != 0 {
						allZero = false
					}
					nextDeriv[i] = delta
				}
				nextValues += thisDeriv[len(thisDeriv)-1]
				if allZero {
					break dance
				}
				thisDeriv = nextDeriv
			}
			newFirstValue := 0
			for i:= len(firstValues)-2; i >= 0; i-- {
				newFirstValue = firstValues[i] - firstValues[i+1]
			}
			fmt.Println("First value:", newFirstValue)
			prevValues += newFirstValue
	}
	fmt.Println("[Part 1] Sum of all next values:", nextValues)
	fmt.Println("[Part 2] Sum of all prev values:", prevValues)
}
