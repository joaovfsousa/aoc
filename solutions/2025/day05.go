package solutions

import (
	"slices"
	"strings"

	"github.com/joaovfsousa/aoc/pkg/aoc"
	"github.com/joaovfsousa/aoc/pkg/aoc/str"
)

type Day5 struct{}

func (d Day5) Year() int { return 2025 }
func (d Day5) Day() int  { return 5 }

type Range struct {
	Min, Max int
}

func (r Range) containsElem(i int) bool {
	return i >= r.Min && i <= r.Max
}

func (r *Range) extendBy(r2 Range) {
	r.Min = min(r.Min, r2.Min)
	r.Max = max(r.Max, r2.Max)
}

func (r Range) countElem() int {
	return r.Max - r.Min + 1
}

func (d Day5) Part1(inputPath string) (any, error) {
	hasFinishedRanges := false
	freshCount := 0

	ranges := []Range{}

	for l := range aoc.IterLines(inputPath) {
		if hasFinishedRanges {

			id := str.StringToInt(l)

			for _, r := range ranges {
				if r.containsElem(id) {
					freshCount++
					break
				}
			}

			continue
		}

		if len(l) == 0 {
			hasFinishedRanges = true
			continue
		}

		sepIndex := strings.Index(l, "-")

		ranges = append(ranges, Range{Min: str.StringToInt(l[:sepIndex]), Max: str.StringToInt(l[sepIndex+1:])})
	}

	return freshCount, nil
}

func (d Day5) Part2(inputPath string) (any, error) {
	ranges := []Range{}

	for l := range aoc.IterLines(inputPath) {
		if len(l) == 0 {
			break
		}

		sepIndex := strings.Index(l, "-")

		min := str.StringToInt(l[:sepIndex])
		max := str.StringToInt(l[sepIndex+1:])

		ranges = append(ranges, Range{Min: min, Max: max})
	}

	slices.SortFunc(ranges, func(r1, r2 Range) int {
		return r1.Min - r2.Min
	})

	prev := ranges[0]

	total := 0

	for _, curr := range ranges {
		if prev.Max >= curr.Min {
			prev.extendBy(curr)
		} else {
			total += prev.countElem()
			prev = curr
		}
	}

	total += prev.countElem()

	return total, nil
}

func init() {
	aoc.RegisterDay(Day5{})
}
