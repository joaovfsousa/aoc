package main

import (
	"flag"
	"log"
	"os"
	"time"

	"github.com/joaovfsousa/aoc/pkg/aoc"
)

func main() {
	year := flag.Int("year", time.Now().Year(), "AoC year")
	day := flag.Int("day", 0, "AoC day (0 = all registered for the year)")
	part := flag.Int("part", 0, "Part to run (0 = both)")
	inputPath := flag.String("input", "", "Path to input file (overrides builtin loader)")
	flag.Parse()

	opts := aoc.RunOptions{
		Year:      *year,
		Day:       *day,
		Part:      *part,
		InputPath: *inputPath,
	}

	if err := aoc.Run(opts); err != nil {
		log.Fatalf("run failed: %v", err)
	}
	os.Exit(0)
}
