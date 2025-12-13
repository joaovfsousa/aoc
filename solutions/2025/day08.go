package solutions

import (
	"fmt"
	"slices"
	"strings"

	"github.com/joaovfsousa/aoc/pkg/aoc"
	"github.com/joaovfsousa/aoc/pkg/convert"
	"github.com/joaovfsousa/aoc/pkg/ds"
)

type Day8 struct{}

func (d Day8) Year() int { return 2025 }
func (d Day8) Day() int  { return 8 }

type Box struct {
	x, y, z int
}

func (b Box) dist(b1 Box) int {
	return ((b.x - b1.x) * (b.x - b1.x)) + ((b.y - b1.y) * (b.y - b1.y)) + ((b.z - b1.z) * (b.z - b1.z))
}

type Dist struct {
	v    int
	a, b *Box
}

func (d Day8) Part1(inputPath string) (any, error) {
	boxes := []*Box{}

	for line := range aoc.IterLines(inputPath) {
		parts := strings.Split(line, ",")

		box := Box{
			x: convert.StringToInt(parts[0]),
			y: convert.StringToInt(parts[1]),
			z: convert.StringToInt(parts[2]),
		}

		boxes = append(boxes, &box)
	}

	ds := ds.NewDisjointSet(boxes...)

	dists := []Dist{}

	for i := 0; i < len(boxes)-1; i++ {
		boxA := boxes[i]
		for j := i + 1; j < len(boxes); j++ {
			boxB := boxes[j]

			dists = append(dists, Dist{
				v: boxA.dist(*boxB),
				a: boxA,
				b: boxB,
			})
		}
	}

	slices.SortFunc(dists, func(d1, d2 Dist) int {
		return d1.v - d2.v
	})

	// Example has to run 10 times, real input 1000
	numEdges := 1000
	if len(boxes) <= 20 {
		numEdges = 10
	}

	for i := range numEdges {
		a, b := dists[i].a, dists[i].b

		ds.Union(a, b)
	}

	sizes := []int{}

	for _, s := range ds.AllSetSizes() {
		sizes = append(sizes, s)
	}

	slices.SortFunc(sizes, func(s1, s2 int) int {
		return s2 - s1
	})

	if len(sizes) < 3 {
		return 0, fmt.Errorf("expected at least 3 circuits, got %d", len(sizes))
	}

	result := 1

	for _, s := range sizes[:3] {
		result *= s
	}

	return result, nil
}

func (d Day8) Part2(inputPath string) (any, error) {
	// TODO: implement
	// Use one of the readers from pkg/aoc/readers.go
	return nil, nil
}

func init() {
	aoc.RegisterDay(Day8{})
}
