package main

import (
	"flag"
	"fmt"
	"math"
	"os"
	"image"
	"iter"
	"strings"

	"github.com/fzipp/astar"
)

var inputFile = flag.String("inputFile", "test_input.txt", "Relative file path to use as input.")
var debug = flag.Bool("debug", false, "Print debug info")

type ReindeerMaze []string

var offsets = [...]image.Point{
	image.Pt(0, -1), // North
	image.Pt(1, 0),  // East
	image.Pt(0, 1),  // South
	image.Pt(-1, 0), // West
}

// Neighbours implements the astar.Graph[Node] interface (with Node = image.Point).
func (f ReindeerMaze) Neighbours(p image.Point) iter.Seq[image.Point] {
	return func(yield func(image.Point) bool) {
		for _, off := range offsets {
			q := p.Add(off)
			if f.isFreeAt(q) {
				if !yield(q) {
					return
				}
			}
		}
	}
}

func (f ReindeerMaze) isFreeAt(p image.Point) bool {
	return f.isInBounds(p) && f[p.Y][p.X] == '.'
}

func (f ReindeerMaze) isInBounds(p image.Point) bool {
	return (0 <= p.X && p.X < len(f[p.Y])) && (0 <= p.Y && p.Y < len(f))
}

func nodeDist(p, q image.Point) float64 {
	d := q.Sub(p)
	return math.Sqrt(float64(d.X*d.X + d.Y*d.Y))
}

func (f ReindeerMaze) put(p image.Point, c rune) {
	f[p.Y] = f[p.Y][:p.X] + string(c) + f[p.Y][p.X+1:]
}

func (f ReindeerMaze) print() {
	for _, row := range f {
		fmt.Println(row)
	}
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
	maze := ReindeerMaze{}
	start, end := image.Point{}, image.Point{}
	for y, line := range lines {
		if line == "" {
			continue
		}
		maze = append(maze, line)
		for x, c := range line {
			if string(c) == "S"	{
				start = image.Pt(x,y)
				maze[y] = strings.ReplaceAll(maze[y], "S", ".")
			}
			if string(c) == "E"	{
				end = image.Pt(x,y)
				maze[y] = strings.ReplaceAll(maze[y], "E", ".")
			}
		}
	}
	path := astar.FindPath[image.Point](maze, start, end, nodeDist, nodeDist)
	for _, p := range path {
		maze.put(p, '*')
	}
	maze.print()
}
