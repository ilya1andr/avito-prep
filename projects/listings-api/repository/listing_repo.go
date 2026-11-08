package repository

import (
	"context"
	"errors"
	"fmt"
	"strings"

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

func (r *ListingRepository) GetAll(ctx context.Context, minPrice *int, maxPrice *int) ([]Listing, error) {
	query := "SELECT id, title, price FROM listings"

	var args []interface{}
	var conditions []string

	if minPrice != nil {
		args = append(args, *minPrice)
		conditions = append(conditions, fmt.Sprintf("price >= $%d", len(args)))
	}

	if maxPrice != nil {
		args = append(args, *maxPrice)
		conditions = append(conditions, fmt.Sprintf("price <= $%d", len(args)))
	}

	if len(conditions) > 0 {
		query += " WHERE " + strings.Join(conditions, " AND ")
	}

	query += " ORDER BY id"

	rows, err := r.db.Query(ctx, query, args...)
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

func (r *ListingRepository) Update(ctx context.Context, id int, title string, price int) (*Listing, error) {
	query := `UPDATE listings SET title = $1, price = $2 WHERE id = $3 RETURNING id, title, price`

	var l Listing
	err := r.db.QueryRow(ctx, query, title, price, id).Scan(&l.ID, &l.Title, &l.Price)
	if err != nil {
		return nil, err
	}
	return &l, nil
}

func (r *ListingRepository) Delete(ctx context.Context, id int) error {
	query := `DELETE FROM listings WHERE id = $1`

	cmdTag, err := r.db.Exec(ctx, query, id)
	if err != nil {
		return err
	}

	if cmdTag.RowsAffected() == 0 {
		return errors.New("not found")
	}
	return nil
}

type ListingRepositoryInterface interface {
	Create(ctx context.Context, title string, price int) (*Listing, error)
	GetAll(ctx context.Context, minPrice *int, maxPrice *int) ([]Listing, error)
	GetByID(ctx context.Context, id int) (*Listing, error)
	Update(ctx context.Context, id int, title string, price int) (*Listing, error)
	Delete(ctx context.Context, id int) error
}
