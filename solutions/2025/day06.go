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

// TODO: refactor to loop only once
func (d Day6) Part2(inputPath string) (any, error) {
	lines := aoc.ReadLines(inputPath)

	input := []string{}

	isFirstColumn := true

	for c := 0; c < len(lines[0]); c++ {
		cvb := strings.Builder{}

		for _, line := range lines {
			cvb.WriteByte(line[c])
		}

		colValue := strings.TrimSpace(cvb.String())

		if len(colValue) == 0 {
			isFirstColumn = true
		} else {
			if isFirstColumn {
				isFirstColumn = false

				op := colValue[len(colValue)-1]
				value := strings.TrimSpace(colValue[:len(colValue)-1])

				input = append(input, string(op))
				input = append(input, value)
			} else {
				input = append(input, colValue)
			}
		}
	}

	total := 0

	accumulated := 0
	var op string

	for _, item := range input {
		if item == "*" {
			total += accumulated
			accumulated = 1
			op = item
			continue
		}

		if item == "+" {
			total += accumulated
			accumulated = 0
			op = item
			continue
		}

		if op == "*" {
			accumulated *= str.StringToInt(item)
			continue
		}

		accumulated += str.StringToInt(item)
	}

	total += accumulated

	return total, nil
}

func init() {
	aoc.RegisterDay(Day6{})
}
