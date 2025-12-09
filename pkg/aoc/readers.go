package aoc

import (
	"bufio"
	"io"
	"iter"
	"os"
	"strings"

	"github.com/charmbracelet/log"
)

func handleCloseError(fn func() error, path string) {
	err := fn()
	if err != nil {
		log.Errorf("Failed to close file %v: %v", path, err)
	}
}

// ReadEntireFile reads the entire file into memory as a string.
// This is the most common pattern for AoC solutions.
func ReadEntireFile(path string) (string, error) {
	b, err := os.ReadFile(path)
	if err != nil {
		return "", err
	}
	return string(b), nil
}

// ReadLines reads the file line by line and returns all lines as a slice.
// Empty lines are included. Trailing newlines are removed.
func ReadLines(path string) []string {
	f, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer handleCloseError(f.Close, path)

	var lines []string
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines
}

func IterLines(path string) iter.Seq2[string, error] {
	return func(yield func(string, error) bool) {
		f, err := os.Open(path)
		if err != nil {
			yield("", err)

			log.Error(err)
			return
		}

		defer handleCloseError(f.Close, path)

		scanner := bufio.NewScanner(f)

		for scanner.Scan() {
			line := scanner.Text()

			stopIteration := yield(line, nil)

			if !stopIteration {
				break
			}
		}
	}
}

// ReadLinesNonEmpty reads the file line by line and returns only non-empty lines.
// Trailing newlines are removed.
func ReadLinesNonEmpty(path string) ([]string, error) {
	lines := ReadLines(path)

	var result []string
	for _, line := range lines {
		if line != "" {
			result = append(result, line)
		}
	}
	return result, nil
}

// ReadCharByChar reads the file character by character, calling the provided function for each character.
// The function receives the character (as a rune) and its byte position in the file.
// Returns an error if the callback returns an error or if there's an I/O error.
func ReadCharByChar(path string, fn func(r rune, pos int) error) error {
	f, err := os.Open(path)
	if err != nil {
		return err
	}
	defer handleCloseError(f.Close, path)

	reader := bufio.NewReader(f)
	pos := 0
	for {
		r, size, err := reader.ReadRune()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}
		if err := fn(r, pos); err != nil {
			return err
		}
		pos += size
	}
	return nil
}

func IterCharByChar(path string) iter.Seq2[string, error] {
	return func(yield func(string, error) bool) {
		f, err := os.Open(path)
		if err != nil {
			yield("", err)

			log.Error(err)
			return
		}

		defer handleCloseError(f.Close, path)

		reader := bufio.NewReader(f)

		for {
			r, _, err := reader.ReadRune()
			if err == io.EOF {
				break
			}
			if err != nil {
				yield("", err)
				break
			}

			yield(string(r), nil)
		}
	}
}

func IterBySeparator(path, separator string) iter.Seq2[string, error] {
	return func(yield func(string, error) bool) {
		accumulator := strings.Builder{}

		for c := range IterCharByChar(path) {
			if c == separator {
				string := accumulator.String()
				if len(string) != 0 {
					yield(accumulator.String(), nil)
					accumulator = strings.Builder{}
				}
				continue
			}

			if c == "\n" {
				yield(accumulator.String(), nil)
				return
			}

			accumulator.WriteString(c)
		}
	}
}

// ReadBytes reads the entire file as a byte slice.
func ReadBytes(path string) ([]byte, error) {
	return os.ReadFile(path)
}
