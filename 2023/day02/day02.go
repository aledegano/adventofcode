package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
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

	// part 1
	possible_games := 0
	max_balls := map[string]int{
		"red":   12,
		"green": 13,
		"blue":  14,
	}
	for _, line := range lines {
		if len(line) < 1 {
			break
		}
		possible_game := true
		games := strings.Split(line, ":")
		id, _ := strconv.Atoi(strings.Split(games[0], " ")[1])
		draws := strings.Split(games[1], ";")
		for _, d := range draws {
			draw := strings.Split(d, ",")
			for _, ball := range draw {
				b := strings.Split(ball, " ")
				count, _ := strconv.Atoi(b[1])
				if count > max_balls[b[2]] {
					fmt.Printf("Game %d is not possible because there are too many %s balls: %d\n", id, b[2], count)
					possible_game = false
					break
				}
			}
			if !possible_game {
				break
			}
		}
		if possible_game {
			possible_games += id
		}
	}
	fmt.Printf("The possible games are %d\n", possible_games)
}
