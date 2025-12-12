package solutions

import (
	"github.com/charmbracelet/log"

	"github.com/joaovfsousa/aoc/pkg/aoc"
)

type Day7 struct{}

func (d Day7) Year() int { return 2025 }
func (d Day7) Day() int  { return 7 }

func (d Day7) Part1(inputPath string) (any, error) {
	input := aoc.ReadLines(inputPath)

	beams := make(map[int]int)

	for i, c := range input[0] {
		if c == 'S' {
			beams[i] = 1
		}
	}

	total := 0

	for _, l := range input[1:] {
		for i, c := range l {
			if c == '^' && beams[i] > 0 {
				beams[i-1] += beams[i]
				beams[i+1] += beams[i]
				delete(beams, i)
				total++
			}
		}
	}

	log.Debug(beams)

	return total, nil
}

func (d Day7) Part2(inputPath string) (any, error) {
	// input := aoc.ReadLines(inputPath)
	return nil, nil
}

func init() {
	aoc.RegisterDay(Day7{})
}
