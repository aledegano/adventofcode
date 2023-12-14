package main

import (
	"flag"
	"fmt"
	"os"
	"regexp"
	"strconv"
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
	springs := map[int]string{}
	errors := map[int][]int{}
	for i, line := range lines {
		if line == "" {
			continue
		}
		springs[i] = strings.Split(line, " ")[0]
		for _,e := range strings.Split(strings.Split(line, " ")[1], ",") {
			err, _ := strconv.Atoi(e)
			errors[i] = append(errors[i], err)
		}
	}
	fmt.Printf("The springs are %v and errors: %v\n", springs, errors)
	re := regexp.MustCompile(`\#*\?*`)
	// part 1
	for i, spring := range springs {
		// start by removing unanbiguous groups of damaged springs
		groups := re.FindAllString(spring, -1)
		fmt.Printf("Groups: %v errors: %v\n", groups, errors[i])
		for _, group := range groups {
			fmt.Printf("Group: %s\n", group)
		}
	}
}
