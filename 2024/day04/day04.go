package main

import (
	"flag"
	"fmt"
	"os"
	"regexp"
	"slices"
	"strings"
)

var inputFile = flag.String("inputFile", "test_input.txt", "Relative file path to use as input.")
var debug = flag.Bool("debug", false, "Print debug info")

func findXMAS(wordSearch []string) int {
	concatSearch := strings.Join(wordSearch, "")
	found := regexp.MustCompile("XMAS").FindAllStringIndex(concatSearch, -1)
	return len(found)
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
	wordSearch := make([][]string, len(lines)+1)
	for i, line := range lines {
		if len(line) == 0 {
			continue
		}
		wordSearch[i+1] = make([]string, len(line)+2)
		wordSearch[i+1][0] = "."
		wordSearch[i+1][len(line)+1] = "."
		for j, char := range line {
			wordSearch[i+1][j+1] = string(char)
		}
	}
	// Add padding to the top and bottom of the matrix
	wordSearch[0] = make([]string, len(lines)+1)
	for i := 0; i < len(wordSearch); i++ {
		wordSearch[0][i] = "."
	}
	wordSearch[len(lines)] = make([]string, len(lines)+1)
	for i := 0; i < len(wordSearch); i++ {
		wordSearch[len(lines)][i] = "."
	}
	totalWordsFound := 0
	// Flatten matrix to search forward and backward
	flatWordSearch := make([]string, len(wordSearch)*len(wordSearch[0]))
	for i, row := range wordSearch {
		for j, char := range row {
			flatWordSearch[i*len(wordSearch[0])+j] = char
		}
	}
	// Search XMAS forward
	totalWordsFound = findXMAS(flatWordSearch)
	// Search XMAS backward
	slices.Reverse(flatWordSearch)
	totalWordsFound += findXMAS(flatWordSearch)
	// Flatten the matrix to search up and down
	flatWordSearch = make([]string, len(wordSearch)*len(wordSearch[0]))
	for i, row := range wordSearch {
		for j, char := range row {
			flatWordSearch[j*len(wordSearch)+i] = char
		}
	}
	// Search XMAS down
	totalWordsFound += findXMAS(flatWordSearch)
	// Search XMAS up
	slices.Reverse(flatWordSearch)
	totalWordsFound += findXMAS(flatWordSearch)
	// Diagonalize the matrix NW-SE
	flatWordSearch = []string{}
	for i := -len(wordSearch); i < len(wordSearch); i++ {
		for j := 0; j < len(wordSearch); j++ {
			if i+j > len(wordSearch)-1 || i+j < 0 {
				continue
			}
			flatWordSearch = append(flatWordSearch, wordSearch[j][i+j])
		}
	}
	// Search XMAS NW->SE
	totalWordsFound += findXMAS(flatWordSearch)
	// Search XMAS SE->NW
	slices.Reverse(flatWordSearch)
	totalWordsFound += findXMAS(flatWordSearch)
	// Diagonalize the matrix NE-SW
	flatWordSearch = []string{}
	for i := 0; i < len(wordSearch)*2; i++ {
		for j := 0; j < len(wordSearch); j++ {
			if i-j < 0 || i-j >= len(wordSearch) {
				continue
			}
			flatWordSearch = append(flatWordSearch, wordSearch[j][i-j])
		}
	}
	// Search XMAS NE->SW
	totalWordsFound += findXMAS(flatWordSearch)
	// Search XMAS SW->NE
	slices.Reverse(flatWordSearch)
	totalWordsFound += findXMAS(flatWordSearch)
	fmt.Println("Part 1:", totalWordsFound)
	//part 2 find the X shaped MAS
	xShapedMAS := 0
	for i := 2; i < len(wordSearch); i++ {
		for j := 2; j < len(wordSearch); j++ {
			if wordSearch[i][j] == "A" {
				if (wordSearch[i-1][j-1] == "M" && wordSearch[i+1][j+1] == "S") || (wordSearch[i-1][j-1] == "S" && wordSearch[i+1][j+1] == "M") {
					if (wordSearch[i-1][j+1] == "M" && wordSearch[i+1][j-1] == "S") || (wordSearch[i-1][j+1] == "S" && wordSearch[i+1][j-1] == "M") {
						xShapedMAS++
					}
				}
			}
		}
	}
	fmt.Println("Part 2:", xShapedMAS)
}
