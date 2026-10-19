package main

import (
	"context"
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/jackc/pgx/v5"
)

func parseCSV(r io.Reader) ([][]string, error) {
	reader := csv.NewReader(r)
	return reader.ReadAll()
}

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintln(os.Stderr, "usage: csv-cli <file.csv>")
		os.Exit(1)
	}

	path := os.Args[1]

	file, err := os.Open(path)
	if err != nil {
		log.Fatalf("open file: %v", err)
	}
	defer file.Close()

	rows, err := parseCSV(file)
	if err != nil {
		log.Fatalf("parse csv: %v", err)
	}

	ctx := context.Background()
	host := os.Getenv("DB_HOST")
	if host == "" {
		host = "localhost"
	}
	user := os.Getenv("DB_USER")
	if user == "" {
		user = "postgres"
	}
	pass := os.Getenv("DB_PASSWORD")
	if pass == "" {
		pass = "secret"
	}
	dbname := os.Getenv("DB_NAME")
	if dbname == "" {
		dbname = "postgres"
	}

	connString := fmt.Sprintf("postgres://%s:%s@%s:5432/%s", user, pass, host, dbname)
	conn, err := pgx.Connect(ctx, connString)
	if err != nil {
		log.Fatalf("Unable to connect to database: %v\n", err)
	}
	defer conn.Close(ctx)

	log.Println("Подключение к БД успешно!")

	for i, row := range rows {
		if i == 0 {
			continue
		}

		_, err := conn.Exec(ctx, "INSERT INTO people (name, age, city) VALUES ($1, $2, $3)", row[0], row[1], row[2])
		if err != nil {
			log.Printf("Ошибка вставки строки %d: %v", i, err)
		}
	}

	log.Println("Все данные успешно загружены в БД!")
}
