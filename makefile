build:
	go build -o bin/main main.go

run:
	go run main.go

test:
	go test

test_coverage:
	go test ./... -coverprofile=coverage.out

all: build run