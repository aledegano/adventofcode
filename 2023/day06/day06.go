package main

import (
	"flag"
	"fmt"
	"math"
	"os"
	"regexp"
	"strings"
	"strconv"
	// "unicode"
)

var inputFile = flag.String("inputFile", "test_input.txt", "Relative file path to use as input.")
var debug = flag.Bool("debug", false, "Print debug info")

func QuadraticFormula(time int, lenght int) (float64, float64) {
	t := float64(time)
	l := float64(lenght)
	discriminant := t*t - 4*l
	if discriminant < 0 {
		return 0, 0
	}
	longPress := (t + math.Sqrt(discriminant)) / 2
	shortPress := (t - math.Sqrt(discriminant)) / 2
	return shortPress, longPress
}

func main() {
	flag.Parse()
	bytes, err := os.ReadFile(*inputFile)
	if err != nil {
		return
	}
	contents := string(bytes)
	lines := strings.Split(contents, "\n")
	re := regexp.MustCompile(`\d+`)
	timesStr := re.FindAllString(lines[0], -1)
	lengthsStr := re.FindAllString(lines[1], -1)
	times := []int{}
	lengths := []int{}
	for _, t := range timesStr {
		t, _ := strconv.Atoi(t)
		times = append(times, t)
	}
	for _, l := range lengthsStr {
		l, _ := strconv.Atoi(l)
		lengths = append(lengths, l)
	}
	// part 1
	// find the shortest and longest press with some physics and math :D
	waysToWin := 1
	for i, t := range times {
		shortPress, longPress := QuadraticFormula(t, lengths[i])
		// all the integers between shortPress and longPress are ways to win the boat race and counting them gives the number of ways to win
		waysToWin *= int(math.Ceil(longPress) - math.Floor(shortPress))-1
	}
	fmt.Printf("[Part 1] Ways to win: %d\n", waysToWin)
	// part 2
	timePt2Str := ""
	lenghtPt2Str := ""
	for _, t := range timesStr {
		timePt2Str += t
	}
	timePt2, _ := strconv.Atoi(timePt2Str)
	for _, l := range lengthsStr {
		lenghtPt2Str += l
	}
	lenghtPt2, _ := strconv.Atoi(lenghtPt2Str)
	shortPress, longPress := QuadraticFormula(timePt2, lenghtPt2)
	waysToWin = int(math.Ceil(longPress) - math.Floor(shortPress))-1
	fmt.Printf("[Part 2] Ways to win: %d\n", waysToWin)
}
