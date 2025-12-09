package solutions

import (
	"strings"

	"github.com/joaovfsousa/aoc/pkg/aoc"
	"github.com/joaovfsousa/aoc/pkg/aoc/str"
)

type Day6 struct{}

func (d Day6) Year() int { return 2025 }
func (d Day6) Day() int  { return 6 }

func (d Day6) Part1(inputPath string) (any, error) {
	input := [][]string{}

	ln := 0
	for _, line := range aoc.ReadLines(inputPath) {
		strBuilder := strings.Builder{}
		input = append(input, []string{})
		for _, char := range line {
			if char == ' ' {
				string := strBuilder.String()
				if len(string) > 0 {
					input[ln] = append(input[ln], string)

					strBuilder = strings.Builder{}
				}

				continue
			}

			strBuilder.WriteRune(char)
		}

		string := strBuilder.String()
		if len(string) > 0 {
			input[ln] = append(input[ln], string)

			strBuilder = strings.Builder{}
		}

		ln++
	}

	total := 0

	for c := 0; c < len(input[0]); c++ {
		op := input[len(input)-1][c]
		columnTotal := 0

		if op == "*" {
			columnTotal = 1
		}

		for r := 0; r < len(input)-1; r++ {
			val := str.StringToInt(input[r][c])
			if op == "*" {
				columnTotal *= val
				continue
			}

			columnTotal += val
		}

		total += columnTotal
	}

	return total, nil
}

func (d Day6) Part2(inputPath string) (any, error) {
	// TODO: implement
	// Use one of the readers from pkg/aoc/readers.go
	return nil, nil
}

func init() {
	aoc.RegisterDay(Day6{})
}
