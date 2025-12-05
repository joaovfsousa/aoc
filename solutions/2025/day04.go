package solutions

import (
	"github.com/charmbracelet/log"

	"github.com/joaovfsousa/aoc/pkg/aoc"
)

var adjacentDirs = [8][2]int{
	{-1, -1},
	{-1, 0},
	{-1, 1},
	{0, -1},
	{0, 1},
	{1, -1},
	{1, 0},
	{1, 1},
}

type Day4 struct{}

func (d Day4) Year() int { return 2025 }
func (d Day4) Day() int  { return 4 }

func (d Day4) Part1(inputPath string) (any, error) {
	rows := []string{}
	totalAccessible := 0

	for l := range aoc.IterLines(inputPath) {
		rows = append(rows, l)
	}

	rowCount := len(rows)
	colCount := len(rows[0])

	for r := range rowCount {
		for c := range colCount {
			totalRolls := 0

			if string(rows[r][c]) != "@" {
				continue
			}

			for _, d := range adjacentDirs {
				x := c + d[0]
				y := r + d[1]

				if x >= 0 && x < colCount && y >= 0 && y < rowCount {
					if string(rows[y][x]) == "@" {
						totalRolls++
					}
				}
			}

			if totalRolls < 4 {
				log.Debugf("x=%v, y=%v", c, r)
				totalAccessible++
			}
		}
	}

	return totalAccessible, nil
}

func (d Day4) Part2(inputPath string) (any, error) {
	// TODO: implement
	// Use one of the readers from pkg/aoc/readers.go
	return nil, nil
}

func init() {
	aoc.RegisterDay(Day4{})
}
