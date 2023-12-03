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
	contents := string(bytes)
	lines := strings.Split(contents, "\n")
	specialChars := []rune{'*','+','$','#'}
	// part 1
	for x, line := range lines {
		for y, r := range line {
			if unicode.IsDigit(r) {
				// Check if there is a special character above, below, left, right or diagonal from this digit
				
			}
		}
	}
}
