package repository

import (
	"context"

	"github.com/jackc/pgx/v5"
)

type Listing struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
	Price int    `json:"price"`
}

type ListingRepository struct {
	db *pgx.Conn
}

func NewListingRepository(db *pgx.Conn) *ListingRepository {
	return &ListingRepository{db: db}
}

func (r *ListingRepository) Create(ctx context.Context, title string, price int) (*Listing, error) {
	var l Listing
	query := `INSERT INTO listings (title, price) VALUES ($1, $2) RETURNING id`

	err := r.db.QueryRow(ctx, query, title, price).Scan(&l.ID)
	if err != nil {
		return nil, err
	}

	l.Title = title
	l.Price = price
	return &l, nil
}
