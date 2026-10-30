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

func (r *ListingRepository) GetAll(ctx context.Context) ([]Listing, error) {
	rows, err := r.db.Query(ctx, "SELECT id, title, price FROM listings ORDER BY id")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var listings []Listing
	for rows.Next() {
		var l Listing
		if err := rows.Scan(&l.ID, &l.Title, &l.Price); err != nil {
			return nil, err
		}
		listings = append(listings, l)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return listings, nil
}

func (r *ListingRepository) GetByID(ctx context.Context, id int) (*Listing, error) {
	var l Listing
	query := `SELECT id, title, price FROM listings WHERE id = $1`

	err := r.db.QueryRow(ctx, query, id).Scan(&l.ID, &l.Title, &l.Price)
	if err != nil {
		return nil, err
	}
	return &l, nil
}
