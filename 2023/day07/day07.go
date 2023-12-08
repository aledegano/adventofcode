package main

import (
	"flag"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
	// "unicode"
)

var inputFile = flag.String("inputFile", "test_input.txt", "Relative file path to use as input.")
var debug = flag.Bool("debug", false, "Print debug info")

type Hand struct {
	cards string
	bid int
	handType int
}

func SortCards(cards string) string {
	mapHandStrength = map[rune]int{
		"A" : 0,
		"K" : 1,
		"Q" : 2,
		"J" : 3,
		"T" : 4,
		"9" : 5,
		"8" : 6,
		"7" : 7,
		"6" : 8,
		"5" : 9,
		"4" : 10,
		"3" : 11,
		"2" : 12,
	}
	cardsInt := []int{}
	for i := 0; i < len(cards); i ++ {
		cardsInt = append(cardsInt, mapHandStrength[cards[i]])
	}
	sort.Ints(cardsInt)
}

func main() {
	flag.Parse()
	bytes, err := os.ReadFile(*inputFile)
	if err != nil {
		return
	}
	contents := string(bytes)
	lines := strings.Split(contents, "\n")
	hands := []Hand{}
	re := regexp.MustCompile(`\w+`)
	for _, line := range lines {
		if line == "" {
			continue
		}
		bid, _ := strconv.Atoi(re.FindAllString(line, -1)[1])
		hand := Hand{re.FindAllString(line, -1)[0], bid}
		hands = append(hands, hand)
	}
	fmt.Println(hands)
}
