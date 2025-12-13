package solutions

import (
	"github.com/charmbracelet/log"

	"github.com/joaovfsousa/aoc/pkg/aoc"
	"github.com/joaovfsousa/aoc/pkg/convert"
)

type Day01 struct{}

func (Day01) Year() int { return 2025 }
func (Day01) Day() int  { return 1 }

func modulo100(a int) int {
	b := a % 100

	if b < 0 {
		b += 100
	}

	return b
}

func (Day01) Part1(inputPath string) (any, error) {
	dial := 50
	zeros := 0

	for l := range aoc.IterLines(inputPath) {
		side := l[:1]
		clicksStr := l[1:]

		clicks := convert.StringToInt(clicksStr)

		if side == "L" {
			dial = modulo100(dial - clicks)
		} else {
			dial = modulo100(dial + clicks)
		}

		if dial == 0 {
			zeros++
		}
	}

	return zeros, nil
}

func (Day01) Part2(inputPath string) (any, error) {
	dial := 50
	zeros := 0

	log.Debugf("Dial=%02d  zeros=%02d", dial, zeros)

	for l := range aoc.IterLines(inputPath) {
		prevDial := dial
		side := l[:1]
		clicksStr := l[1:]

		clicks := convert.StringToInt(clicksStr)

		distanceToZero := 0

		if side == "L" {
			distanceToZero = dial
			dial = modulo100(dial - clicks)
		} else {
			distanceToZero = 100 - dial

			dial = modulo100(dial + clicks)
		}

		mc := modulo100(clicks)

		add := 0

		// If prevDial == 0, distanceToZero will always be <= mc.
		if mc >= distanceToZero && prevDial != 0 {
			add = 1
		}

		toAdd := ((clicks / 100) + (add))

		zeros += toAdd

		log.Debugf("movement=% 3s  dial=%02d  zeros=%02d  toAdd=%02d  d2z=%02d  mc=%02d", l, dial, zeros, toAdd, distanceToZero, mc)
	}

	return zeros, nil
}

func init() { aoc.RegisterDay(Day01{}) }
