package solutions

import (
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

func (r Range) contains(i int) bool {
	return i >= r.Min && i <= r.Max
}

func (d Day5) Part1(inputPath string) (any, error) {
	hasFinishedRanges := false
	freshCount := 0

	ranges := []Range{}

	for l := range aoc.IterLines(inputPath) {
		if hasFinishedRanges {

			id := str.StringToInt(l)

			for _, r := range ranges {
				if r.contains(id) {
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

		ranges = append(ranges, Range{Min: str.StringToInt(l[:sepIndex]), Max: str.StringToInt(l[sepIndex+1:])})
	}

	return 0, nil
}

func init() {
	aoc.RegisterDay(Day5{})
}
