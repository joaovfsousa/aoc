package solutions

import (
	"fmt"

	"github.com/charmbracelet/log"

	"github.com/joaovfsousa/aoc/pkg/aoc"
	"github.com/joaovfsousa/aoc/pkg/convert"
)

type Day3 struct{}

func (d Day3) Year() int { return 2025 }
func (d Day3) Day() int  { return 3 }

func (d Day3) Part1(inputPath string) (any, error) {
	total := 0
	for l := range aoc.IterLines(inputPath) {
		lineLen := len(l)
		firstB := convert.StringToInt(string(l[0]))

		secondB := convert.StringToInt(string(l[1]))

		for i, nAsRune := range l {
			n := convert.StringToInt(string(nAsRune))
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

		jottage := convert.StringToInt(fmt.Sprintf("%v%v", firstB, secondB))

		log.Debugf("jottage = %v", jottage)

		total += jottage
	}

	return total, nil
}

func (d Day3) Part2(inputPath string) (any, error) {
	total := 0

	for l := range aoc.IterLines(inputPath) {
		lenLine := len(l)
		num := l[:12]

		for evalIndex := 12; evalIndex < lenLine; evalIndex++ {
			eval := l[evalIndex]
			mmax := num
			for toRemove := range 12 {
				firstPart := num[:toRemove]
				secondPart := num[toRemove+1:]
				newNum := firstPart + secondPart + string(eval)

				log.Debugf("evalIndex= %v, toRemove= %v, num= %v, firstPart= %v, secondPart= %v, eval= %v", evalIndex, toRemove, num, firstPart, secondPart, eval)

				if convert.StringToInt(newNum) > convert.StringToInt(mmax) {
					mmax = newNum
				}
			}

			if convert.StringToInt(mmax) > convert.StringToInt(num) {
				num = mmax
			}
		}

		log.Debugf("Max = %v", num)

		total += convert.StringToInt(num)
	}

	return total, nil
}

func init() {
	aoc.RegisterDay(Day3{})
}
