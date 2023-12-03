package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)

var inputFile = flag.String("inputFile", "test_input.txt", "Relative file path to use as input.")
var debug = flag.Bool("debug", false, "Print debug info")

type Bounds struct {
	minX, maxX, minY, maxY int
}

func isPart(lines []string, posX int, posY int, b Bounds) bool {
	for i:=posX-1; i<=posX+1; i++ {
		if i < b.minX || i > b.maxX {
			continue
		}
		for j:=posY-1; j<=posY+1; j++ {
			if j < b.minY || j > b.maxY || (i == posX && j == posY) {
				continue
			}
			if ! unicode.IsDigit(rune(lines[i][j])) && lines[i][j] != '.' {
				return true
			}
		}
	}
	return false
}

func main() {
	flag.Parse()
	bytes, err := os.ReadFile(*inputFile)
	if err != nil {
		return
	}
	contents := string(bytes)
	lines := strings.Split(contents, "\n")
	bounds := Bounds{0, len(lines)-2, 0, len(lines[0])-1} // The x boundaries needs to also remove the last empty line
	buffer := []string{}
	save := false
	partNumber := 0
	partsSum := 0
	// part 1
	for x, line := range lines {
		if len(line) < 1 {
			break
		}
		for y, r := range line {
			if unicode.IsDigit(r) {
				buffer = append(buffer, string(r))
				// Check if there is a special character above, below, left, right or diagonal from this digit
				if ! save {
					save = isPart(lines, x, y, bounds)
				}
				if save {
					fmt.Printf("Digit %c at (%d, %d) is a part number\n", r, x, y)
				}
			} else {
				if save {
					partNumber, _ = strconv.Atoi(strings.Join(buffer, ""))
					partsSum += partNumber
					save = false
					fmt.Printf("Part number %d saved.\n", partNumber)
				}
				buffer = []string{}
			}
		}
	}
	fmt.Printf("The sum of the parts is %d\n", partsSum)
}
