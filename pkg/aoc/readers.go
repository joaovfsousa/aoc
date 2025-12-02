package aoc

import (
	"bufio"
	"io"
	"os"
)

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
func ReadLines(path string) ([]string, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	var lines []string
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

// ReadLinesNonEmpty reads the file line by line and returns only non-empty lines.
// Trailing newlines are removed.
func ReadLinesNonEmpty(path string) ([]string, error) {
	lines, err := ReadLines(path)
	if err != nil {
		return nil, err
	}
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
	defer f.Close()

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

// ReadBytes reads the entire file as a byte slice.
func ReadBytes(path string) ([]byte, error) {
	return os.ReadFile(path)
}

