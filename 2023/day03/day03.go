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

func isPart(lines []string, posX int, posY int, b Bounds) (bool, int) {
	for i := posX - 1; i <= posX+1; i++ {
		if i < b.minX || i > b.maxX {
			continue
		}
		for j := posY - 1; j <= posY+1; j++ {
			if j < b.minY || j > b.maxY || (i == posX && j == posY) {
				continue
			}
			if !unicode.IsDigit(rune(lines[i][j])) && lines[i][j] != '.' {
				if lines[i][j] == '*' {
					return true, i*1000 + j // This part `could` be part of a gear
				} else {
					return true, -1 // This part is not part of a gear
				}
			}
		}
	}
	return false, -1
}

func main() {
	flag.Parse()
	bytes, err := os.ReadFile(*inputFile)
	if err != nil {
		return
	}
	contents := string(bytes)
	lines := strings.Split(contents, "\n")
	bounds := Bounds{0, len(lines) - 2, 0, len(lines[0]) - 1} // The x boundaries needs to also remove the last empty line
	buffer := []string{}
	save := false
	var partNumber, gear, partsSum int
	gearCandidates := make(map[int][]int)
	// part 1
	for x, line := range lines {
		if len(line) < 1 {
			break
		}
		for y, r := range line {
			if unicode.IsDigit(r) {
				buffer = append(buffer, string(r))
				// Check if there is a special character above, below, left, right or diagonal from this digit
				if !save {
					save, gear = isPart(lines, x, y, bounds)
				}
			} else {
				if save {
					partNumber, _ = strconv.Atoi(strings.Join(buffer, ""))
					partsSum += partNumber
					save = false
					fmt.Printf("Part number %d saved.\n", partNumber)
					if gear >= 0 {
						gearCandidates[gear] = append(gearCandidates[gear], partNumber)
						fmt.Printf("Gear candidate %d saved.\n", gear)
					}
				}
				buffer = []string{}
			}
		}
	}
	fmt.Printf("The sum of the parts is %d\n", partsSum)
	fmt.Printf("The gear candidates are %v\n", gearCandidates)
	gearRatio := 0
	for _, parts := range gearCandidates {
		if len(parts) == 2 {
			gearRatio += parts[0] * parts[1]
		}
	}
	fmt.Printf("The gear ratio is %d\n", gearRatio)
}
