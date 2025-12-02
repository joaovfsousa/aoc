package aoc

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
)

// DownloadInput downloads puzzle input for a given year/day and writes it to inputs/<year>/dayXX.txt.
// It requires a valid session token (the value of the "session" cookie) — provide it via the AOC_SESSION env var.
// NOTE: network calls will run when you execute this in your environment.
func DownloadInput(year, day int, session string) error {
	if session == "" {
		return errors.New("empty session token")
	}
	url := fmt.Sprintf("https://adventofcode.com/%d/day/%d/input", year, day)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return err
	}
	req.AddCookie(&http.Cookie{Name: "session", Value: session})
	req.Header.Set("User-Agent", "aoc-go-framework - github.com/joaovfsousa/aoc")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return fmt.Errorf("adventofcode returned %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	dir := filepath.Join("inputs", fmt.Sprintf("%04d", year))
	if err := os.MkdirAll(dir, 0o755); err != nil {
		return err
	}
	path := filepath.Join(dir, fmt.Sprintf("day%02d.txt", day))
	if err := os.WriteFile(path, body, 0o644); err != nil {
		return err
	}
	return nil
}
