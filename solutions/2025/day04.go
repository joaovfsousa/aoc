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

// TODO: refactor to process with 3 lines instead of loading everything to memory
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
				x := c + d[1]
				y := r + d[0]

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

func replaceInStringMatrix(m []string, i [2]int) []string {
	rm := []string{}

	r, c := i[0], i[1]

	rowToReplace := m[r]

	prevRows := m[:r]
	nextRows := m[r+1:]

	prevCols := rowToReplace[:c]
	nextCols := rowToReplace[c+1:]

	rowAfterReplace := prevCols + "x" + nextCols

	rm = append(rm, prevRows...)
	rm = append(rm, rowAfterReplace)
	rm = append(rm, nextRows...)

	return rm
}

func (d Day4) Part2(inputPath string) (any, error) {
	rows := []string{}
	totalRemoved := 0

	for l := range aoc.IterLines(inputPath) {
		rows = append(rows, l)
	}

	rowCount := len(rows)
	colCount := len(rows[0])

	var removedCount int

	for ok := true; ok; ok = (removedCount != 0) {
		toRemove := [][2]int{}

		for r := range rowCount {
			for c := range colCount {
				totalRolls := 0

				if string(rows[r][c]) != "@" {
					continue
				}

				for _, d := range adjacentDirs {
					x := c + d[1]
					y := r + d[0]

					if x >= 0 && x < colCount && y >= 0 && y < rowCount {
						if string(rows[y][x]) == "@" {
							totalRolls++
						}
					}
				}

				if totalRolls < 4 {
					log.Debugf("x=%v, y=%v", c, r)
					toRemove = append(toRemove, [2]int{r, c})
				}
			}
		}

		removedCount = len(toRemove)

		log.Debugf("Removed %v rolls", removedCount)

		totalRemoved += removedCount

		for _, tr := range toRemove {
			rows = replaceInStringMatrix(rows, tr)
		}
	}

	return totalRemoved, nil
}

func init() {
	aoc.RegisterDay(Day4{})
}
