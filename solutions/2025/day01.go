package solutions

import (
	"strconv"

	"github.com/charmbracelet/log"

	"github.com/joaovfsousa/aoc/pkg/aoc"
)

type Day01 struct{}

func (Day01) Year() int { return 2025 }
func (Day01) Day() int  { return 1 }

func (Day01) Part1(inputPath string) (any, error) {
	dial := 50
	zeros := 0

	for l := range aoc.IterLines(inputPath) {
		side := l[:1]
		clicksStr := l[1:]

		clicks, err := strconv.Atoi(clicksStr)
		if err != nil {
			log.Errorf("Error transforming %v to integer: %v", clicksStr, err)
			return -1, nil
		}

		if side == "L" {
			dial = (dial - clicks) % 100
		} else {
			dial = (dial + clicks) % 100
		}

		if dial < 0 {
			dial += 100
		}

		if dial == 0 {
			zeros++
		}
	}

	return zeros, nil
}

func (Day01) Part2(inputPath string) (any, error) {
	return 0, nil
}

func init() { aoc.RegisterDay(Day01{}) }
