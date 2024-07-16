build:
	go build -o bin/main main.go

run:
	go run main.go

test:
	go test

test_coverage:
	go test -coverprofile=coverage.out

docker b:
	docker build . -t atoyegbe/sre-bootcamp

docker r:
	docker run -p80:8000 atoyegbe/sre-bootcamp

all: build run