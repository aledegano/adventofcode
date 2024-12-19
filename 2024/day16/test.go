package main

import (
	"fmt"
	"image"
	"iter"
	"math"

	"github.com/fzipp/astar"
)

func main() {
	maze := floorPlan{
		"###############",
		"#   # #     # #",
		"# ### ### ### #",
		"#   # # #   # #",
		"### # # # ### #",
		"# # #         #",
		"# # ### ### ###",
		"#   # # # #   #",
		"### # # # # ###",
		"# #       # # #",
		"# # ######### #",
		"#         #   #",
		"# ### # # ### #",
		"#   # # #     #",
		"###############",
	}
	start := image.Pt(1, 13) // Bottom left corner
	dest := image.Pt(13, 1)  // Top right corner

	// Find the shortest path
	path := astar.FindPath[image.Point](maze, start, dest, nodeDist, nodeDist)

	fmt.Println(path)
	// Mark the path with dots before printing
	for _, p := range path {
		maze.put(p, '.')
	}
	maze.print()
}

// nodeDist is our cost function. We use points as nodes, so we
// calculate their Euclidean distance.
func nodeDist(p, q image.Point) float64 {
	d := q.Sub(p)
	return math.Sqrt(float64(d.X*d.X + d.Y*d.Y))
}

type floorPlan []string

var offsets = [...]image.Point{
	image.Pt(0, -1), // North
	image.Pt(1, 0),  // East
	image.Pt(0, 1),  // South
	image.Pt(-1, 0), // West
}

// Neighbours implements the astar.Graph[Node] interface (with Node = image.Point).
func (f floorPlan) Neighbours(p image.Point) iter.Seq[image.Point] {
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

func (f floorPlan) isFreeAt(p image.Point) bool {
	return f.isInBounds(p) && f[p.Y][p.X] == ' '
}

func (f floorPlan) isInBounds(p image.Point) bool {
	return (0 <= p.X && p.X < len(f[p.Y])) && (0 <= p.Y && p.Y < len(f))
}

func (f floorPlan) put(p image.Point, c rune) {
	f[p.Y] = f[p.Y][:p.X] + string(c) + f[p.Y][p.X+1:]
}

func (f floorPlan) print() {
	for _, row := range f {
		fmt.Println(row)
	}
}
