# aoc-go

Reusable Advent of Code framework in Go.

- "Framework" code live in `pkg/aoc/`.
- Put per-day solutions in `solutions/<year>/dayXX.go`.
- Cached inputs live in `inputs/<year>/dayXX.txt`.
- Use `AOC_SESSION` env var (your session cookie) to download inputs.

Examples:

```
  make scaffold year=2025 day=1
  AOC_SESSION=xxxx make download year=2025 day=1
  make run-day year=2025 day=1
```

Press enter to recompile and run or ctrl-c to close.
