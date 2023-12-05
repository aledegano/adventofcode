package main

import (
	"flag"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
	"unicode"
)

var inputFile = flag.String("inputFile", "test_input.txt", "Relative file path to use as input.")
var debug = flag.Bool("debug", false, "Print debug info")

type transform struct {
	source int
	dest int
	rng int
}

func WalkSteps(steps [][]transform, seeds []int) int {
	nearerLocation := 0
	for i, seed := range seeds {
		for _, step := range steps {
			for _, t := range step {
				if seed >= t.source && seed <= t.source + t.rng -1 {
					seed += t.dest - t.source
					break
				}
			}
		}
		if seed < nearerLocation || i == 0 {
			nearerLocation = seed
		}
	}
	return nearerLocation
}

func main() {
	flag.Parse()
	bytes, err := os.ReadFile(*inputFile)
	if err != nil {
		return
	}
	contents := string(bytes)
	lines := strings.Split(contents, "\n")
	// parse seeds
	re := regexp.MustCompile(`\d+`)
	seedsStr := re.FindAllString(lines[0], -1)
	seeds := make([]int, len(seedsStr))
	for i, seed := range seedsStr {
		s, _ := strconv.Atoi(seed)
		seeds[i] = s
	}
	// each step is a list of transforms, each transform is valid for a range of _source_ values
	steps := [][]transform{}
	stepTransforms := []transform{}
	locations := []int{}
	recordLocations := false
	for _, line := range lines[2:] {
		if len(line) == 0 {
			steps = append(steps, stepTransforms) // every time we hit a blank line, we're done with the step, add it to the list of steps
			continue
		}
		// each time a line starts with a alphabetic character, it's a new step
		if unicode.IsLetter(rune(line[0])) {
			if line[0] == 'h' { // humidity-to-location map, to record all the locations for part 2
				recordLocations = true
			}
			stepTransforms = []transform{}
			continue
		}
		tStr := re.FindAllString(line, -1)
		dest, _ := strconv.Atoi(tStr[0])
		source, _ := strconv.Atoi(tStr[1])
		rng, _ := strconv.Atoi(tStr[2])
		// stepTransforms = append(stepTransforms, transform{source, source + rng -1, dest-source})
		stepTransforms = append(stepTransforms, transform{source, dest, rng})
		if recordLocations {
			locations = append(locations, dest)
		}
	}
	// run each seed through the steps
	nearerLocation := WalkSteps(steps, seeds)
	fmt.Printf("[Part 1] The nearer location of the seeds is %v\n", nearerLocation)
	// part 2, instead of evaluating all the seeds, we reverse the steps and check if each location corresponds to a _range_ of seeds
	rangeSeeds := [][]int{}
	for i := 0; i < len(seeds); i+=2 {
		r := []int{seeds[i], seeds[i]+seeds[i+1]}
		rangeSeeds = append(rangeSeeds, r)
	}
	fmt.Printf("the seeds ranges are %v\n", rangeSeeds)
	// reverse the steps and for each location compute which seed would map to it
	fmt.Printf("the locations are %v\n", locations)
}
