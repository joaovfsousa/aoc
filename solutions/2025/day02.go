package solutions

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/charmbracelet/log"

	"github.com/joaovfsousa/aoc/pkg/aoc"
)

type Day2 struct{}

func (d Day2) Year() int { return 2025 }
func (d Day2) Day() int  { return 2 }

type Range struct {
	divider int
	start   int
	end     int
}

func loop(inputPath string, calculateInvalidSum func(i int) int) (int, error) {
	invalidIdsSum := 0
	for r := range aoc.IterBySeparator(inputPath, ",") {
		sepIndex := strings.Index(r, "-")

		min, err := strconv.Atoi(r[0:sepIndex])
		if err != nil {
			return 0, err
		}

		max, err := strconv.Atoi(r[sepIndex+1:])
		if err != nil {
			return 0, err
		}

		for i := min; i < max+1; i++ {
			invalidIdsSum += calculateInvalidSum(i)
		}
	}

	return invalidIdsSum, nil
}

func calculateInvalidSumPart1(i int) int {
	numAsStr := fmt.Sprintf("%d", i)

	strSize := len(numAsStr)

	if strSize%2 == 1 {
		return 0
	}

	if numAsStr[:(strSize/2)] == numAsStr[(strSize/2):] {
		return i
	}

	return 0
}

func getRanges(i int) [][]Range {
	ranges := [][]Range{}

	for j := 1; j < i; j++ {
		if i%j != 0 { // not a divider
			continue
		}

		drs := []Range{}

		for k := 0; k < i; k += j {
			drs = append(drs, Range{
				divider: j,
				start:   k,
				end:     k + j,
			})
		}

		if len(drs) > 1 {
			ranges = append(ranges, drs)
		}
	}

	return ranges
}

func calculateInvalidSumPart2(i int) int {
	numAsStr := fmt.Sprintf("%d", i)

	strSize := len(numAsStr)

	dividers := getRanges(strSize)

	log.Debugf("num = %v, dividers = %v", i, dividers)

	for _, d := range dividers {
		isEqual := true
		lastChunk := ""

		for _, r := range d {
			chunk := numAsStr[r.start:r.end]

			log.Debugf("i = %v, lastChunk = %v, chunk = %v", i, lastChunk, chunk)

			if lastChunk != "" && chunk != lastChunk {
				isEqual = false
				break
			}

			lastChunk = chunk
		}

		if isEqual {
			log.Debugf("Invalid Id = %v, divider = %v", i, d)
			return i
		}
	}

	return 0
}

func (d Day2) Part1(inputPath string) (any, error) {
	return loop(inputPath, calculateInvalidSumPart1)
}

func (d Day2) Part2(inputPath string) (any, error) {
	return loop(inputPath, calculateInvalidSumPart2)
}

func init() {
	aoc.RegisterDay(Day2{})
}
