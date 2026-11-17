package service

import (
	"context"
	"errors"
	"strings"

	"github.com/ilya1andr/avito-prep/projects/listings-api/repository"
)

type ListingService struct {
	repo repository.ListingRepositoryInterface
}

func NewListingService(repo repository.ListingRepositoryInterface) *ListingService {
	return &ListingService{repo: repo}
}

func (s *ListingService) CreateListing(ctx context.Context, title string, price int) (*repository.Listing, error) {
	title = strings.TrimSpace(title)
	if title == "" {
		return nil, errors.New("название не может быть пустым")
	}
	if len(title) > 100 {
		return nil, errors.New("название не может быть длиннее 100 символов")
	}
	if price <= 0 {
		return nil, errors.New("цена должна быть больше нуля")
	}
	if price > 10000000 {
		return nil, errors.New("цена не может быть больше 10 000 000")
	}

	return s.repo.Create(ctx, title, price)
}

func (s *ListingService) CreateListingWithTx(ctx context.Context, title string, price int) (*repository.Listing, error) {
	title = strings.TrimSpace(title)
	if title == "" {
		return nil, errors.New("название не может быть пустым")
	}
	if len(title) > 100 {
		return nil, errors.New("название не может быть длиннее 100 символов")
	}
	if price <= 0 {
		return nil, errors.New("цена должна быть больше нуля")
	}
	if price > 10000000 {
		return nil, errors.New("цена не может быть больше 10 000 000")
	}

	tx, err := s.repo.BeginTx(ctx)
	if err != nil {
		return nil, err
	}

	defer tx.Rollback(ctx)

	listing, err := s.repo.CreateWithTx(ctx, tx, title, price)
	if err != nil {
		return nil, err
	}
	if err := tx.Commit(ctx); err != nil {
		return nil, err
	}

	return listing, nil
}

func (s *ListingService) GetListingByID(ctx context.Context, id int) (*repository.Listing, error) {
	if id <= 0 {
		return nil, errors.New("ID должен быть положительным")
	}
	return s.repo.GetByID(ctx, id)
}

func (s *ListingService) GetListings(ctx context.Context, minPrice *int, maxPrice *int) ([]repository.Listing, error) {
	if minPrice != nil && maxPrice != nil && *minPrice > *maxPrice {
		return nil, errors.New("минимальная цена не может быть больше максимальной")
	}

	return s.repo.GetAll(ctx, minPrice, maxPrice)
}

func (s *ListingService) UpdateListing(ctx context.Context, id int, title string, price int) (*repository.Listing, error) {
	if id <= 0 {
		return nil, errors.New("id должен быть положительный")
	}
	if title == "" {
		return nil, errors.New("title не может быть пустым")
	}
	if price <= 0 {
		return nil, errors.New("price должна быть больше 0")
	}
	return s.repo.Update(ctx, id, title, price)
}

func (s *ListingService) DeleteListing(ctx context.Context, id int) error {
	if id <= 0 {
		return errors.New("id должен быть положительный")
	}
	return s.repo.Delete(ctx, id)
}
