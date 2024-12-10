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

func calculateChecksum(disk []int) int {
	checksum := 0
	for i, id := range disk {
		if id == -1 {
			break
		}
		checksum += i * id
	}
	return checksum
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
	blockPos := map[int][]int{} // id -> [start, size]
	for i, digitS := range contents {
		digit, _ := strconv.Atoi(string(digitS))
		if i%2 == 0 {
			blockPos[id] = []int{len(disk), digit}
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
	origDisk := make([]int, len(disk))
	copy(origDisk, disk)
	// printDisk(disk)
	//part 1 replace the first empty element with the last non-empty element
	for i, id := range disk {
		// printDisk(disk)
		if id == -1 {
			for j := len(disk) - 1; j > i; j-- {
				if disk[j] != -1 {
					disk[i] = disk[j]
					disk[j] = -1
					break
				}
			}
		}
	}
	// printDisk(disk)
	fmt.Println("Part 1, checksum: ", calculateChecksum(disk))
	// reset the disk
	copy(disk, origDisk)
	printDisk(disk)
	// part 2, now the replacement has to be done in blocks: for each empty _block_ replace with a block of non-empty elements that can fit starting from the end of the disk
	for i := len(blockPos) - 1; i > 0; i-- { // consider the blocks from the last to the first
		// loop over the disk to find the leftmost empty block that can fit the current block
		for j := 0; j < len(disk); j++ {
			emptySize := 0
			if disk[j] == -1 {
				emptySize++
				if emptySize == blockPos[i][1] && j < blockPos[i][0] { // move only if the empty space is left to the block
					// replace the block
					for k := 0; k < blockPos[i][1]; k++ {
						disk[j+k] = blockPos[i][0]
					}
					break
				}
			} else {
				emptySize = 0
			}
		}
	}
	printDisk(disk)
	fmt.Println("Part 2, checksum: ", calculateChecksum(disk))
}
