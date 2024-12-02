package main

import (
	"flag"
	"fmt"
	"math"
	"os"
	"regexp"
	"strconv"
	"strings"
)

var inputFile = flag.String("inputFile", "test_input.txt", "Relative file path to use as input.")
var debug = flag.Bool("debug", false, "Print debug info")

func safeReport(report []int) bool {
	for i:=1; i<len(report); i++ {
		if i == len(report)-1 { // if we reached the last element then the report is safe
			return true
		}
		// Check that levels are monotonic increasing OR decreasing
		if (report[i] > report[i-1]) && (report[i+1] < report[i]) { // non-monotonic
			return false
		}
		if (report[i] < report[i-1]) && (report[i+1] > report[i]) { // non-monotonic
			return false
		}
		deltaPrevious := math.Abs(float64(report[i] - report[i-1]))
		deltaNext := math.Abs(float64(report[i+1] - report[i]))
		if (deltaPrevious == 0 || deltaPrevious > 3) || (deltaNext == 0 || deltaNext > 3) {
			return false
		}
	}
	return false
}

func RemoveIndex(s []int, index int) []int {
	return append(s[:index], s[index+1:]...)
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
	safeReportsCountPt1 := 0
	safeReportsCountPt2 := 0
	for _, line := range lines {
		if len(line) < 1 {
			break
		}
		report := []int{}
		re := regexp.MustCompile(`\d+`)
		l := re.FindAllString(line, -1)
		for _, n := range l {
			level, _ := strconv.Atoi(n)
			report = append(report, level)
		}
		// part 1, stop at first error
		safe := safeReport(report)
		if safe {
			safeReportsCountPt1++
			safeReportsCountPt2++
		} else {
		// part 2, try removing one level and see if it's safe
			for i:=0; i<len(report); i++ {
				newReport := make([]int, len(report))
				copy(newReport, report)
				newReport = RemoveIndex(newReport, i)
				safe = safeReport(newReport)
				if safe {
					safeReportsCountPt2++
					break
				}
			}
		}
	}
	fmt.Printf("The number of safe reports in part 1 is %d\n", safeReportsCountPt1)
	fmt.Printf("The number of safe reports in part 2 is %d\n", safeReportsCountPt2)
}
