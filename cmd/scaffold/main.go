package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/joaovfsousa/aoc/pkg/aoc"
)

func main() {
	year := flag.Int("year", 0, "AoC year (required)")
	day := flag.Int("day", 0, "AoC day (required)")
	flag.Parse()

	if *year == 0 || *day == 0 {
		log.Fatal("year and day are required")
	}

	if err := aoc.ScaffoldDay("solutions", *year, *day); err != nil {
		log.Fatalf("scaffold failed: %v", err)
	}
	fmt.Printf("scaffolded solutions/%04d/day%02d.go", *year, *day)
}
