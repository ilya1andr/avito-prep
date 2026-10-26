package main

import (
	"context"
	"encoding/csv"
	"flag"
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
	statsFlag := flag.Bool("stats", false, "Show statistics from the database")
	flag.Parse()

	args := flag.Args()
	if len(args) < 1 && !*statsFlag {
		fmt.Fprintln(os.Stderr, "usage: csv-cli <file.csv> or csv-cli --stats")
		os.Exit(1)
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

	if *statsFlag {
		fmt.Println("--- Статистика по городам ---")
		rows, err := conn.Query(ctx, "SELECT city, COUNT(*), AVG(age) FROM people GROUP BY city ORDER BY city")
		if err != nil {
			log.Fatalf("Query error: %v", err)
		}
		defer rows.Close()

		for rows.Next() {
			var city string
			var count int
			var avgAge float64
			if err := rows.Scan(&city, &count, &avgAge); err != nil {
				log.Fatal(err)
			}
			fmt.Printf("Город: %-10s | Людей: %d | Средний возраст: %.1f\n", city, count, avgAge)
		}
		return
	}

	path := args[0]
	file, err := os.Open(path)
	if err != nil {
		log.Fatalf("open file: %v", err)
	}
	defer file.Close()

	rows, err := parseCSV(file)
	if err != nil {
		log.Fatalf("parse csv: %v", err)
	}

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
