package main

import (
	"flag"
	"fmt"
	"math"
	"os"
	"regexp"
	// "strconv"
	"strings"
	// "unicode"
)

var inputFile = flag.String("inputFile", "test_input.txt", "Relative file path to use as input.")
var debug = flag.Bool("debug", false, "Print debug info")

func main() {
	flag.Parse()
	bytes, err := os.ReadFile(*inputFile)
	if err != nil {
		return
	}
	contents := string(bytes)
	lines := strings.Split(contents, "\n")
	//part1
	stack_value := 0
	re := regexp.MustCompile(`\d{1,2}`)
	for _, line := range lines {
		if len(line) < 1 {
			break
		}
		numbers := strings.Split(line, ":")[1]
		tickets := strings.Split(numbers, "|")
		winning_numbers := -1
		card_value := 0
		winning_ticket := re.FindAllString(tickets[0], -1)
		owned_ticket := re.FindAllString(tickets[1], -1)
		for _, own_nmbr := range owned_ticket {
			for _, win_nmbr := range winning_ticket {
				if own_nmbr == win_nmbr {
					winning_numbers++
					break
				}
			}
		}
		if winning_numbers >= 0 {
			card_value = int(math.Pow(2, float64(winning_numbers)))
		}
		stack_value += int(card_value)
	}
	fmt.Printf("The value of the stack is: %d\n", stack_value)
}
