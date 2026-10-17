.PHONY: build run test clean

build:
	go build -o csv-cli .

run:
	go run . data.csv

test:
	go test -v ./...

clean:
	rm -f csv-cli