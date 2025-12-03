package solutions

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/joaovfsousa/aoc/pkg/aoc"
)

type Day2 struct{}

func (d Day2) Year() int { return 2025 }
func (d Day2) Day() int  { return 2 }

func (d Day2) Part1(inputPath string) (any, error) {
	invalidIdsSum := 0
	for r := range aoc.IterBySeparator(inputPath, ",") {
		sepIndex := strings.Index(r, "-")

		min, err := strconv.Atoi(r[0:sepIndex])
		if err != nil {
			return nil, err
		}

		max, err := strconv.Atoi(r[sepIndex+1:])
		if err != nil {
			return nil, err
		}

		for i := min; i < max+1; i++ {
			numAsStr := fmt.Sprintf("%d", i)

			strSize := len(numAsStr)

			if strSize%2 == 1 {
				continue
			}

			if numAsStr[:(strSize/2)] == numAsStr[(strSize/2):] {
				invalidIdsSum += i
			}

		}
	}

	return invalidIdsSum, nil
}

func (d Day2) Part2(inputPath string) (any, error) {
	return nil, nil
}

func init() {
	aoc.RegisterDay(Day2{})
}
