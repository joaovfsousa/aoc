package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/charmbracelet/log"

	"github.com/joaovfsousa/aoc/pkg/aoc"
	_ "github.com/joaovfsousa/aoc/solutions/2025"
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

	// Handle Ctrl-C gracefully
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-sigChan
		fmt.Println("\nExiting...")
		os.Exit(0)
	}()

	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("\033[H\033[2J")
		if err := aoc.Run(opts); err != nil {
			log.Infof("run failed: %v", err)
		}

		fmt.Print("\nPress 'enter' to rerun, 'ctrl-c' to quit")
		input, _, _ := reader.ReadRune()

		if string(input) != "\n" {
			break
		}

		fmt.Println()
	}
}
