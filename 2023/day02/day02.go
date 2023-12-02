package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
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

	// part 1
	possible_games := 0
	max_cubes := map[string]int{
		"red":   12,
		"green": 13,
		"blue":  14,
	}
	set_cube_power := 0
	for _, line := range lines {
		min_cubes := map[string]int{
			"red":   1,
			"green": 1,
			"blue":  1,
		}
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
				if count > min_cubes[b[2]] {
					min_cubes[b[2]] = count
				}
				if count > max_cubes[b[2]] {
					fmt.Printf("Game %d is not possible because there are too many %s cubes: %d\n", id, b[2], count)
					possible_game = false
				}
			}
		}
		if possible_game {
			possible_games += id
		}
		set_cube_power += min_cubes["red"] * min_cubes["green"] * min_cubes["blue"]
	}
	fmt.Printf("The possible games are %d\n", possible_games)
	fmt.Printf("The set cube power is %d\n", set_cube_power)
}
