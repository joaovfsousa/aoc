package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/joaovfsousa/aoc/pkg/aoc"
)

func main() {
	year := flag.Int("year", 0, "AoC year (required)")
	day := flag.Int("day", 0, "AoC day (required)")
	flag.Parse()
	if *year == 0 || *day == 0 {
		log.Fatal("year and day are required")
	}
	session := os.Getenv("AOC_SESSION")
	if session == "" {
		log.Fatal("AOC_SESSION environment variable not set")
	}
	if err := aoc.DownloadInput(*year, *day, session); err != nil {
		log.Fatalf("download failed: %v", err)
	}
	fmt.Printf("downloaded inputs/%04d/day%02d.txt", *year, *day)
}
