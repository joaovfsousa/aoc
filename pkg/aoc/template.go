package aoc

import (
	"fmt"
	"os"
	"path/filepath"
	"text/template"
)

const dayTemplate = `package solutions

import "github.com/joaovfsousa/aoc/pkg/aoc"

// Day{{.Day}} implements aoc.Day for {{.Year}} day {{.Day}}
type Day{{.Day}} struct{}

func (d Day{{.Day}}) Year() int { return {{.Year}} }
func (d Day{{.Day}}) Day() int  { return {{.Day}} }

func (d Day{{.Day}}) Part1(inputPath string) (any, error) {
	// TODO: implement
	// Use aoc.ReadEntireFile(inputPath) to read the entire file
	// Or aoc.ReadLines(inputPath) for line-by-line reading
	// Or aoc.ReadCharByChar(inputPath, ...) for character-by-character reading
	return nil, nil
}

func (d Day{{.Day}}) Part2(inputPath string) (any, error) {
	// TODO: implement
	// Use aoc.ReadEntireFile(inputPath) to read the entire file
	// Or aoc.ReadLines(inputPath) for line-by-line reading
	// Or aoc.ReadCharByChar(inputPath, ...) for character-by-character reading
	return nil, nil
}

func init() {
	aoc.RegisterDay(Day{{.Day}}{})
}
`

func ScaffoldDay(root string, year, day int) error {
	dir := filepath.Join(root, fmt.Sprintf("%04d", year))
	if err := os.MkdirAll(dir, 0o755); err != nil {
		return err
	}
	path := filepath.Join(dir, fmt.Sprintf("day%02d.go", day))
	if _, err := os.Stat(path); err == nil {
		return fmt.Errorf("file already exists: %s", path)
	}

	f, err := os.Create(path)
	if err != nil {
		return err
	}
	defer f.Close()

	t := template.Must(template.New("day").Parse(dayTemplate))
	data := map[string]int{"Year": year, "Day": day}
	return t.Execute(f, data)
}
