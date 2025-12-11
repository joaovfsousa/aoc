package aoc

import (
	"errors"
	"fmt"
	"time"

	"github.com/charmbracelet/log"
)

// Day is the interface each day's file implements.
type Day interface {
	Year() int
	Day() int
	Part1(inputPath string) (any, error)
	Part2(inputPath string) (any, error)
}

// registry holds registered days
var registry = make(map[string]Day)

func key(year, day int) string { return fmt.Sprintf("%04d-%02d", year, day) }

// RegisterDay is called from each solution file init()
func RegisterDay(d Day) {
	registry[key(d.Year(), d.Day())] = d
}

// RunOptions config
type RunOptions struct {
	Year      int
	Day       int // 0 => all
	Part      int // 0 => both
	InputPath string
	Example   bool
}

// Run runs the requested days/parts
func Run(opts RunOptions) error {
	found := false
	for k, d := range registry {
		if d.Year() != opts.Year {
			continue
		}
		if opts.Day != 0 && d.Day() != opts.Day {
			continue
		}
		found = true

		inputPath, err := GetInputPath(opts.Year, d.Day(), opts.InputPath, opts.Example)
		if err != nil {
			return err
		}

		runOne := func(part int) error {
			start := time.Now()
			var res any
			var runErr error
			if part == 1 {
				res, runErr = d.Part1(inputPath)
			} else {
				res, runErr = d.Part2(inputPath)
			}
			dur := time.Since(start)
			if runErr != nil {
				return fmt.Errorf("%s part %d failed: %w", k, part, runErr)
			}
			log.Infof("%s part %d → %v    (%.3fms)", k, part, res, float64(dur.Microseconds())/1000)
			return nil
		}

		if opts.Part == 0 || opts.Part == 1 {
			if err := runOne(1); err != nil {
				return err
			}
		}
		if opts.Part == 0 || opts.Part == 2 {
			if err := runOne(2); err != nil {
				return err
			}
		}
	}

	if !found {
		return errors.New("no days found for that year/day")
	}
	return nil
}
