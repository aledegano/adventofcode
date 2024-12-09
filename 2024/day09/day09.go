package main

import (
	"flag"
	"fmt"
	// "math"
	"os"
	// "regexp"
	"strconv"
	"strings"
	// "unicode"
)

var inputFile = flag.String("inputFile", "test_input.txt", "Relative file path to use as input.")
var debug = flag.Bool("debug", false, "Print debug info")

func printDisk(disk []int) {
	for _, id := range disk {
		if id == -1 {
			fmt.Print(".")
		} else {
			fmt.Print(id)
		}
	}
	fmt.Println()
}

func main() {
	flag.Parse()
	bytes, err := os.ReadFile(*inputFile)
	if err != nil {
		fmt.Printf("Error reading file %s: %v\n", *inputFile, err)
		return
	}
	contents := string(bytes)
	// remove all newlines from the input
	contents = strings.ReplaceAll(contents, "\n", "")
	id := 0
	disk := []int{}
	for i, digitS := range contents {
		digit, _ := strconv.Atoi(string(digitS))
		if i % 2 == 0 {
			for j := 0; j < digit; j++ {
				disk = append(disk, id)
			}
			id++
		} else {
			for j := 0; j < digit; j++ {
				disk = append(disk, -1)
			}
		}
	}
	// printDisk(disk)
	//part 1 replace the first empty element with the last non-empty element
	for i, id := range disk {
		// printDisk(disk)
		if id == -1 {
			for j := len(disk)-1; j > i; j-- {
				if disk[j] != -1 {
					disk[i] = disk[j]
					disk[j] = -1
					break
				}
			}
		}
	}
	printDisk(disk)
	checksum := 0
	for i, id := range disk {
		if id == -1 {
			break
		}
		checksum += i * id
	}
	fmt.Println("Part 1, checksum: ", checksum)
}
