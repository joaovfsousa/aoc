.PHONY: run scaffold download test zip

RUN_GO := go run ./cmd/runner
SCAFFOLD_GO := go run ./cmd/scaffold

run:
	$(RUN_GO)

scaffold:
	$(SCAFFOLD_GO) -year=$(year) -day=$(day)

download:
	@echo "download requires AOC_SESSION env var set (your session cookie)"
	go run ./cmd/downloader -year=$(year) -day=$(day)

test:
	go test ./...

