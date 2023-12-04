package main

import (
	"flag"
	"fmt"
	"math"
	"os"
	"regexp"
	"strings"
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
	cards_score := []int{}
	cards_copies := []int{}
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
			cards_score = append(cards_score, winning_numbers+1)
		} else {
			cards_score = append(cards_score, 0)
		}
		cards_copies = append(cards_copies, 1)
		stack_value += int(card_value)
	}
	fmt.Printf("[Part 1] The value of the stack is: %d\n", stack_value)
	//part2
	for id := 0; id < len(cards_score); id++ {
		if cards_score[id] == 0 {
			continue
		}
		for i := id + 1; i <= id+cards_score[id]; i++ { // create n copies of the next score cards
			cards_copies[i] += cards_copies[id]
		}
	}
	total_copies := 0
	for id := 0; id < len(cards_copies); id++ {
		total_copies += cards_copies[id]
	}
	fmt.Printf("[Part 2] Total copies: %d\n", total_copies)
}
