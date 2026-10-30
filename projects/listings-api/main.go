package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"github.com/ilya1andr/avito-prep/projects/listings-api/handler"
	"github.com/ilya1andr/avito-prep/projects/listings-api/repository"
	"github.com/ilya1andr/avito-prep/projects/listings-api/service"
	"github.com/jackc/pgx/v5"
)

func main() {

	ctx := context.Background()
	connString := "postgres://postgres:secret@localhost:5432/postgres"
	conn, err := pgx.Connect(ctx, connString)
	if err != nil {
		log.Fatalf("Unable to connect to database: %v\n", err)
	}
	defer conn.Close(ctx)

	repo := repository.NewListingRepository(conn)
	svc := service.NewListingService(repo)
	h := handler.NewListingHandler(svc)

	mux := http.NewServeMux()

	mux.HandleFunc("POST /listings", h.CreateListing)

	mux.HandleFunc("GET /listings", h.GetListingsHandler)

	mux.HandleFunc("GET /listings/{id}", h.GetListingHandlerByID)

	mux.HandleFunc("GET /health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]string{"status": "ok"})
	})

	log.Println("Starting server on :8080")
	log.Fatal(http.ListenAndServe(":8080", mux))
}
