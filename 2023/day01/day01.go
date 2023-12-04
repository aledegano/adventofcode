package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"regexp"
	"strconv"
	"strings"
	"unicode"
)

var inputFile = flag.String("inputFile", "test_input.txt", "Relative file path to use as input.")
var debug = flag.Bool("debug", false, "Print debug info")

func main() {
	flag.Parse()
	bytes, err := ioutil.ReadFile(*inputFile)
	if err != nil {
		return
	}
	// part 1
	calibration := 0
	contents := string(bytes)
	lines := strings.Split(contents, "\n")
	for _, line := range lines {
		firstDigit := -1
		lastDigit := 0
		for _, r := range line {
			if unicode.IsDigit(r) {
				if firstDigit < 0 {
					firstDigit = int(r - '0')
				}
				lastDigit = int(r - '0')
			}
		}
		if firstDigit < 0 {
			break
		}
		calibration += firstDigit*10 + lastDigit
	}
	fmt.Printf("The calibration is %d\n", calibration)
	// end part 1

	// part 2 This is wrong, it does not handle the case where a character is shared between two numbers spelled: e.g. oneight
	spellOut := map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
		"four":  4,
		"five":  5,
		"six":   6,
		"seven": 7,
		"eight": 8,
		"nine":  9,
	}
	calibration = 0
	re := regexp.MustCompile(`one|two|three|four|five|six|seven|eight|nine|\d`)
	for i, line := range lines {
		firstDigit := 0
		lastDigit := 0
		found := re.FindAllStringSubmatch(line, -1)
		if len(found) == 0 {
			continue
		}

		if len(found[0]) == 1 {
			firstDigit, err = strconv.Atoi(found[0])
			if err != nil {
				fmt.Printf("Error converting %s to int\n", found[0])
			}
		} else {
			firstDigit = spellOut[found[0]]
		}

		if len(found[len(found)-1]) == 1 {
			lastDigit, err = strconv.Atoi(found[len(found)-1])
			if err != nil {
				fmt.Printf("Error converting %s to int\n", found[0])
			}
		} else {
			lastDigit = spellOut[found[len(found)-1]]
		}
		if *debug {
			fmt.Printf("[DEBUG] line: %s, firstDigit: %d, lastDigit: %d, calibration: %d\n", line, firstDigit, lastDigit, firstDigit*10+lastDigit)
		}
		fmt.Printf("%d: %d\n", i, firstDigit*10+lastDigit)
		calibration += firstDigit*10 + lastDigit
	}
	fmt.Printf("The calibration is %d\n", calibration)
}
