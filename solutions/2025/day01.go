package solutions

import (
	"errors"
	"strconv"
	"strings"

	"github.com/joaovfsousa/aoc/pkg/aoc"
)

type Day01 struct{}

func (Day01) Year() int { return 2025 }
func (Day01) Day() int  { return 1 }

func parseLines(input string) ([]int, error) {
	lines := strings.Fields(strings.TrimSpace(input))
	out := make([]int, 0, len(lines))
	for _, l := range lines {
		v, err := strconv.Atoi(l)
		if err != nil {
			return nil, err
		}
		out = append(out, v)
	}
	return out, nil
}

func (Day01) Part1(inputPath string) (any, error) {
	input, err := aoc.ReadEntireFile(inputPath)
	if err != nil {
		return nil, err
	}
	nums, err := parseLines(input)
	if err != nil {
		return nil, err
	}
	if len(nums) == 0 {
		return nil, errors.New("no input")
	}
	sum := 0
	for _, v := range nums {
		sum += v
	}
	return sum, nil
}

func (Day01) Part2(inputPath string) (any, error) {
	// different logic
	return 0, nil
}

func init() { aoc.RegisterDay(Day01{}) }
