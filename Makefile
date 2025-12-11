.PHONY: run scaffold download test zip

RUN_GO := go run ./cmd/runner
SCAFFOLD_GO := go run ./cmd/scaffold

e ?= 0

run:
	$(RUN_GO) -e=$(e)

run-day:
	$(RUN_GO) -year=$(year) -day=$(day) -e=$(e)

scaffold:
	$(SCAFFOLD_GO) -year=$(year) -day=$(day)

download:
	go run ./cmd/downloader -year=$(year) -day=$(day)

test:
	go test ./...
