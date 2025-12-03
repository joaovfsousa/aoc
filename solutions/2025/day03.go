package solutions

import (
	"fmt"
	"strconv"

	"github.com/charmbracelet/log"

	"github.com/joaovfsousa/aoc/pkg/aoc"
)

type Day3 struct{}

func (d Day3) Year() int { return 2025 }
func (d Day3) Day() int  { return 3 }

func runeToInt(r rune) int {
	i, err := strconv.Atoi(string(r))
	if err != nil {
		panic(fmt.Sprintf("Failed to convert %v to int", r))
	}

	return i
}

func (d Day3) Part1(inputPath string) (any, error) {
	total := 0
	for l := range aoc.IterLines(inputPath) {
		lineLen := len(l)
		firstB := runeToInt(rune(l[0]))

		secondB := runeToInt(rune(l[1]))

		for i, nAsRune := range l {
			n := runeToInt(nAsRune)
			if i == 0 {
				continue
			}

			if n > firstB && i < lineLen-1 {
				firstB = n
				secondB = 0
				continue
			}

			if n > secondB {
				secondB = n
			}
		}

		jottage, err := strconv.Atoi(fmt.Sprintf("%v%v", firstB, secondB))
		if err != nil {
			panic(err)
		}

		log.Debugf("jottage = %v", jottage)

		total += jottage
	}

	return total, nil
}

func (d Day3) Part2(inputPath string) (any, error) {
	return nil, nil
}

func init() {
	aoc.RegisterDay(Day3{})
}
