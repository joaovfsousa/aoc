package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"os/exec"
	"os/signal"
	"path/filepath"
	"syscall"
	"time"

	"github.com/charmbracelet/log"

	"github.com/joaovfsousa/aoc/pkg/aoc"
	_ "github.com/joaovfsousa/aoc/solutions/2025"
)

func rebuildAndRestart(year, day, part int, inputPath string) error {
	tmpDir := os.TempDir()
	tmpBinary := filepath.Join(tmpDir, fmt.Sprintf("aoc-runner-%d", os.Getpid()))

	buildCmd := exec.Command("go", "build", "-o", tmpBinary, "./cmd/runner")
	buildCmd.Stdout = os.Stdout
	buildCmd.Stderr = os.Stderr
	if err := buildCmd.Run(); err != nil {
		return fmt.Errorf("build failed: %w", err)
	}

	args := []string{tmpBinary}
	if year != time.Now().Year() {
		args = append(args, "-year", fmt.Sprintf("%d", year))
	}
	if day != 0 {
		args = append(args, "-day", fmt.Sprintf("%d", day))
	}
	if part != 0 {
		args = append(args, "-part", fmt.Sprintf("%d", part))
	}
	if inputPath != "" {
		args = append(args, "-input", inputPath)
	}

	// Execute the new binary, replacing this process
	return syscall.Exec(tmpBinary, args, os.Environ())
}

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

		log.Infof("year=%v, day=%v, part=%v", *year, *day, *part)

		if err := aoc.Run(opts); err != nil {
			log.Infof("run failed: %v", err)
		}

		fmt.Print("\nPress 'enter' to rerun, 'ctrl-c' to quit")
		input, _, _ := reader.ReadRune()

		if string(input) != "\n" {
			break
		}

		// Rebuild and restart
		fmt.Println("\nRebuilding...")
		if err := rebuildAndRestart(*year, *day, *part, *inputPath); err != nil {
			log.Errorf("rebuild failed: %v", err)
			fmt.Println("Press 'enter' to continue...")
			reader.ReadRune()
			continue
		}
		// If rebuildAndRestart succeeds, it will have replaced this process
		break
	}
}
