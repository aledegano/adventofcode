package main

import (
	"flag"
	"fmt"
	"os"
	// "regexp"
	"strings"
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
	
	rpsMatrix := make([][]int, 3)
	rpsMatrix[0] = []int{3+1, 6+2, 0+3} // rock-rock, rock-paper, rock-scissors
	rpsMatrix[1] = []int{0+1, 3+2, 6+3} // paper-rock, paper-paper, paper-scissors
	rpsMatrix[2] = []int{6+1, 0+2, 3+3} // scissors-rock, scissors-paper, scissors-scissors
	decryptMap := map[string]int{"A": 0, "B": 1, "C": 2, "X": 0, "Y": 1, "Z": 2}
	gameSum := 0
	for _, line := range lines {
		if len(line) == 0 {
			continue
		}
		game := rpsMatrix[decryptMap[string(line[0])]][decryptMap[string(line[2])]]
		gameSum = gameSum + game
	}
	fmt.Println("Part 1: ", gameSum)
	// Part 2
	resultToChoice := make([][]int, 3)
	resultToChoice[0] = []int{2, 0, 1} // rock-lose, rock-draw, rock-win
	resultToChoice[1] = []int{0, 1, 2} // paper-lose, paper-draw, paper-win
	resultToChoice[2] = []int{1, 2, 0} // scissors-lose, scissors-draw, scissors-win
	gameSum = 0
	for _, line := range lines {
		if len(line) == 0 {
			continue
		}
		game := rpsMatrix[decryptMap[string(line[0])]][resultToChoice[decryptMap[string(line[0])]][decryptMap[string(line[2])]]]
		gameSum = gameSum + game
	}
	fmt.Println("Part 2: ", gameSum)
}
