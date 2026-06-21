.PHONY: all test test-race bench tidy run-fundamentals run-goroutines lint help

all: test

test:
	go test ./...

test-race:
	go test -race ./...

bench:
	go test -bench=. -benchmem ./...

tidy:
	go mod tidy

run-fundamentals:
	go run ./01-fundamentals/examples/

run-goroutines:
	go run ./49-goroutines/examples/

run-linked-lists:
	go run ./18-linked-lists/examples/

run-sorting:
	go run ./31-sorting/examples/

help:
	@echo "Targets: test, test-race, bench, tidy, run-fundamentals, run-goroutines"
