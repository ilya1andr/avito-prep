.PHONY: build run test clean

build:
	go build -o csv-cli .

run:
	go run . data.csv

test:
	go test -v ./...

clean:
	del /Q csv-cli

db-up:
	docker run --name pg -e POSTGRES_PASSWORD=secret -p 5432:5432 -d postgres:16

db-stop:
	docker stop pg

db-start:
	docker start pg

db-shell:
	docker exec -it pg psql -U postgres

stats:
	go run . --stats