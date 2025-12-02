package aoc

import (
	"fmt"
	"os"
	"path/filepath"
)

// GetInputPath returns the path to the input file.
// It checks overridePath first, then ./inputs/<year>/dayXX.txt, then ./inputs/dayXX.txt.
// If file is missing and AOC_SESSION is set, it will try to download it.
func GetInputPath(year, day int, overridePath string) (string, error) {
	if overridePath != "" {
		if _, err := os.Stat(overridePath); err == nil {
			return overridePath, nil
		}
		return "", fmt.Errorf("input file not found: %s", overridePath)
	}

	inputsDir := "inputs"
	filename := filepath.Join(inputsDir, fmt.Sprintf("%04d", year), fmt.Sprintf("day%02d.txt", day))
	if _, err := os.Stat(filename); err == nil {
		return filename, nil
	}

	// If file not found, look for ./inputs/dayXX.txt
	filename = filepath.Join(inputsDir, fmt.Sprintf("day%02d.txt", day))
	if _, err := os.Stat(filename); err == nil {
		return filename, nil
	}

	// Try downloader if session present
	session := os.Getenv("AOC_SESSION")
	if session != "" {
		if err := DownloadInput(year, day, session); err == nil {
			// try again
			filename = filepath.Join(inputsDir, fmt.Sprintf("%04d", year), fmt.Sprintf("day%02d.txt", day))
			if _, err := os.Stat(filename); err == nil {
				return filename, nil
			}
		}
	}

	return "", fmt.Errorf("input file not found: tried %s", filename)
}